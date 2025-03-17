package gen

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"slices"

	"github.com/eminaktas/kbgen/pkg/config"
	"github.com/eminaktas/kbgen/pkg/kcl"
	"google.golang.org/protobuf/proto"
	"kcl-lang.io/kcl-go/pkg/tools/gen"
	"kcl-lang.io/lib/go/api"
)

const (
	// Import-related constants.
	importTemplate = "import (\n\t%s\n)\n"
	defaultImport  = "apiextensionsv1 \"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1\""
	defaultAnyType = "*apiextensionsv1.JSON"

	// DeepCopy function template.
	deepCopyTemplate = `func (in *{{ .TypeName }}) DeepCopyInto(out *{{ .TypeName }}) {
	*out = *in
}

func (in *{{ .TypeName }}) DeepCopy() *{{ .TypeName }} {
	if in == nil {
		return nil
	}
	out := new({{ .TypeName }})
	in.DeepCopyInto(out)
	return out
}
`

	// KCL basic types.
	typSchema           = "schema"
	typDict             = "dict"
	typList             = "list"
	typStr              = "str"
	typInt              = "int"
	typFloat            = "float"
	typBool             = "bool"
	typAny              = "any"
	typUnion            = "union"
	typNumberMultiplier = "number_multiplier"

	intPointerStr = "*int"
)

// Generator generates Go struct code from KCL schemas.
type Generator struct {
	kcl         *kcl.KCL
	packageName string
	directory   string
	outputDir   string
	config      *config.Config
}

// NewGeneratorWithPath initializes a new Generator instance using the provided paths.
func NewGeneratorWithPath(packageName, programPath, directory, outputDir, configPath string) (*Generator, error) {
	absConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadConfig(absConfigPath)
	if err != nil {
		return nil, err
	}

	// Set default values for custom type if not provided.
	if cfg.CustomAnyType.Type == "" {
		cfg.CustomAnyType.Type = defaultAnyType
	}
	if cfg.CustomAnyType.Import == "" {
		cfg.CustomAnyType.Import = defaultImport
	}

	absProgramPath, err := filepath.Abs(programPath)
	if err != nil {
		return nil, err
	}
	absDirectory, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}
	absOutputDir, err := filepath.Abs(outputDir)
	if err != nil {
		return nil, err
	}

	return &Generator{
		kcl:         kcl.NewKCLWithPath(absProgramPath),
		packageName: packageName,
		directory:   absDirectory,
		outputDir:   absOutputDir,
		config:      cfg,
	}, nil
}

// Generate reads KCL files, generates Go struct definitions, and writes them to files.
func (gr *Generator) Generate() error {
	// Gather all .k files from the given directory.
	kFiles, err := gr.collectKCLFiles()
	if err != nil {
		return err
	}

	// Get mapping from schema names to KCL types.
	schemaMap, err := gr.kcl.GetSchemaTypeMapping(kFiles)
	if err != nil {
		return err
	}

	// Generate the Go struct definitions from the KCL schema map.
	processed := make(map[string]bool)
	nestedStructs, err := gr.generateGoStructs(schemaMap, processed)
	if err != nil {
		return err
	}

	// Disabled, not needed anymore
	// Remove any empty string items.
	// cleanStructs := gr.removeEmptyItems(nestedStructs)

	// Write the generated code to files.
	for filename, structDefs := range nestedStructs {
		outputPath := filepath.Join(gr.outputDir, gr.packageName, filename+".go")
		if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
			return err
		}
		// Prepend package declaration and join struct definitions.
		fileContent := fmt.Sprintf("package %s\n\n%s", gr.packageName, strings.Join(structDefs, "\n"))
		if err := os.WriteFile(outputPath, []byte(fileContent), 0644); err != nil {
			return err
		}
	}

	return nil
}

// collectKCLFiles walks the directory and returns all paths with a ".k" extension.
func (gr *Generator) collectKCLFiles() ([]string, error) {
	var kFiles []string
	err := filepath.WalkDir(gr.directory, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".k" {
			kFiles = append(kFiles, path)
		}
		return nil
	})
	return kFiles, err
}

// generateGoStructs processes the schema map and returns a mapping from file name to a slice of Go code snippets.
func (gr *Generator) generateGoStructs(
	schemaMap map[string]*api.KclType,
	processed map[string]bool,
) (map[string][]string, error) {
	result := make(map[string][]string)
	// Use a queue to process schemas recursively.
	schemasQueue := []string{}
	for name := range schemaMap {
		schemasQueue = append(schemasQueue, name)
	}

	for len(schemasQueue) > 0 {
		// Dequeue the first element.
		name := schemasQueue[0]
		schemasQueue = schemasQueue[1:]
		schema := schemaMap[name]

		// Determine the file name from the schema's filename.
		baseFileName := schema.GetFilename()
		if baseFileName == "" {
			baseFileName = schema.Item.GetFilename()
		}
		filename := strings.TrimSuffix(filepath.Base(baseFileName), filepath.Ext(baseFileName))

		// Use a string builder to construct the struct definition.
		var structDefBuilder strings.Builder

		var fields map[string]*api.KclType
		var baseSchema *api.KclType

		// Skip if already processed.
		if processed[name] {
			continue
		}
		processed[name] = true

		// Check if this schema should be generated as a map type.
		for _, mapSchema := range gr.config.MapTypeSchemas {
			if mapSchema.Name == name {
				mapTypeDef := fmt.Sprintf("type %s map[string]*%s\n", name, mapSchema.Item.SchemaName)
				result[filename] = append(result[filename], mapTypeDef)
				// Skip further processing for this schema.
				goto nextSchema
			}
		}

		// Begin struct definition.
		structDefBuilder.WriteString(fmt.Sprintf("type %s struct {\n", name))

		// Retrieve properties for the schema.
		fields = schema.Properties
		if fields == nil {
			if schema.Item != nil {
				fields = schema.Item.Properties
			} else {
				return nil, fmt.Errorf("possible index signature used in %s. "+
					"Please refer to the documentation for the workaround", name)
			}
		}

		// If current schema comes with BaseSchema, remove the existing fields
		// in the struct and include BaseSchema with inline tag.
		if baseSchema = schema.BaseSchema; baseSchema != nil {
			schemasQueue = append(schemasQueue, baseSchema.SchemaName)
			schemaMap[baseSchema.SchemaName] = baseSchema
			if baseSchema.Properties != nil {
				for bName := range baseSchema.Properties {
					delete(fields, bName)
				}
			}
		}

		// Process each field in the schema.
		for fieldName, fieldType := range fields {
			// Skip non-exported fields.
			if strings.HasPrefix(fieldName, "_") {
				continue
			}

			// Check if the field's type is a schema type and queue nested schemas if needed.
			if isSchema, nestedTypes := isTypeSchema(fieldType); isSchema {
				// Determine whether this is an index signature case.
				possibleIndexSignature := true
				if nestedTypes == nil {
					if fieldType.Item != nil {
						if _, ok := schemaMap[fieldType.Item.SchemaName]; !ok {
							schemasQueue = append(schemasQueue, fieldType.Item.SchemaName)
							schemaMap[fieldType.Item.SchemaName] = proto.Clone(fieldType).(*api.KclType)
							possibleIndexSignature = false
						}
					} else {
						if _, ok := schemaMap[fieldType.SchemaName]; !ok {
							schemasQueue = append(schemasQueue, fieldType.SchemaName)
							if fieldType.Item != nil || fieldType.Properties != nil {
								schemaMap[fieldType.SchemaName] = proto.Clone(fieldType).(*api.KclType)
								possibleIndexSignature = false
							}
						}
					}
				} else {
					for nestedName, nestedSchema := range nestedTypes {
						if _, ok := schemaMap[nestedName]; !ok {
							schemasQueue = append(schemasQueue, nestedName)
							schemaMap[nestedName] = nestedSchema
							possibleIndexSignature = false
						}
					}
				}
				// Fallback: try to retrieve the schema using its filename.
				if possibleIndexSignature {
					if fileName := fieldType.GetFilename(); fileName != "" {
						nestedMapping, err := gr.kcl.GetSchemaTypeMapping([]string{fileName})
						if err != nil {
							return nil, err
						}
						for nestedName, nestedSchema := range nestedMapping {
							if _, ok := schemaMap[nestedName]; !ok {
								schemasQueue = append(schemasQueue, nestedName)
								schemaMap[nestedName] = nestedSchema
							}
						}
					}
				}
			}

			// Convert KCL type to Go type.
			goType, err := gr.kclTypeToGoType(fieldType)
			if err != nil {
				return nil, err
			}

			// If the type uses a custom "any" type, ensure the necessary import is added.
			if strings.Contains(goType, gr.config.CustomAnyType.Type) {
				if !gr.hasImport(result[filename], gr.config.CustomAnyType.Import) {
					// Prepend the import statement.
					result[filename] = append([]string{fmt.Sprintf(
						importTemplate, gr.config.CustomAnyType.Import)}, result[filename]...)
				}
			}

			// Add field comment (if any) and kubebuilder validation markers.
			if fieldComment := getFieldComment(fieldType); fieldComment != "" {
				structDefBuilder.WriteString(fmt.Sprintf("\t%s\n", fieldComment))
			}
			omitEmpty := ",omitempty"
			if slices.Contains(schema.Required, fieldName) {
				omitEmpty = ""
				structDefBuilder.WriteString("\t// +required\n")
			} else {
				structDefBuilder.WriteString("\t// +optional\n")
			}

			// Append the field definition (capitalize to export the field).
			structDefBuilder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s%s\" yaml:\"%s%s\"`\n",
				capitalize(fieldName), goType, fieldName, omitEmpty, fieldName, omitEmpty))
		}

		// Handle BaseSchema
		if baseSchema != nil {
			structDefBuilder.WriteString(fmt.Sprintf("\n\t%s `json:\",inline\" yaml:\",inline\"`\n", baseSchema.SchemaName))
		}

		structDefBuilder.WriteString("}\n")
		result[filename] = append(result[filename], structDefBuilder.String())

		if isSchema, _ := isTypeSchema(schema); isSchema {
			// Append a DeepCopyInto function if go type is struct
			if deepCopyFunc, err := generateDeepCopyFuncs(name); err != nil {
				return nil, err
			} else {
				result[filename] = append(result[filename], deepCopyFunc)
			}
		}

	nextSchema:
		// Label used to skip additional processing when a map type is generated.
		continue
	}

	return result, nil
}

// generateDeepCopyFuncs to generate the deep copy functions using the template
func generateDeepCopyFuncs(typeName string) (string, error) {
	tmpl, err := template.New("deepCopy").Parse(deepCopyTemplate)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, map[string]string{
		"TypeName": typeName,
	})
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

// hasImport checks whether the given import string is already present.
func (gr *Generator) hasImport(codeSnippets []string, importStr string) bool {
	if len(codeSnippets) > 0 {
		expectedImport := fmt.Sprintf(importTemplate, importStr)
		return codeSnippets[0] == expectedImport
	}
	return false
}

// removeEmptyItems cleans up empty string items from the generated code map.
// func (gr *Generator) removeEmptyItems(nestedStructs map[string][]string) map[string][]string {
// 	cleaned := make(map[string][]string)
// 	for key, items := range nestedStructs {
// 		for _, item := range items {
// 			if item != "" {
// 				cleaned[key] = append(cleaned[key], item)
// 			}
// 		}
// 	}
// 	return cleaned
// }

// kclTypeToGoType converts a KCL type to its corresponding Go type.
func (gr *Generator) kclTypeToGoType(kType *api.KclType) (string, error) {
	switch kType.Type {
	case typInt:
		return intPointerStr, nil
	case typFloat:
		return "*float64", nil
	case typBool:
		return "*bool", nil
	case typStr:
		return "string", nil
	case typList:
		elemType, err := gr.kclTypeToGoType(kType.Item)
		if err != nil {
			return "", err
		}
		return "[]" + elemType, nil
	case typDict:
		keyType, err := gr.kclTypeToGoType(kType.Key)
		if err != nil {
			return "", err
		}
		valType, err := gr.kclTypeToGoType(kType.Item)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("map[%s]%s", keyType, valType), nil
	case typSchema:
		return "*" + kType.SchemaName, nil
	case typNumberMultiplier:
		return intPointerStr, nil
	case typUnion:
		uniqueTypes := make(map[string]bool)
		for _, t := range kType.UnionTypes {
			tType, err := gr.kclTypeToGoType(t)
			if err != nil {
				return "", err
			}
			uniqueTypes[tType] = true
		}
		// If there is exactly one unique type, use it; otherwise use custom any type.
		if len(uniqueTypes) == 1 {
			for t := range uniqueTypes {
				return t, nil
			}
		}
		return gr.config.CustomAnyType.Type, nil
	case typAny:
		return gr.config.CustomAnyType.Type, nil
	default:
		// Fallback: check if it is a literal type.
		if isLit, basicTyp, _ := gen.IsLitType(kType); isLit {
			switch basicTyp {
			case typBool:
				return "*bool", nil
			case typInt:
				return intPointerStr, nil
			case typFloat:
				return "*float64", nil
			case typStr:
				return "string", nil
			}
		}
		return "", fmt.Errorf("unknown KCL type: '%v'", kType.Type)
	}
}

// getFieldComment returns appropriate field comments based on the KCL type.
func getFieldComment(kType *api.KclType) string {
	const anyTypeComment = "// +kubebuilder:pruning:PreserveUnknownFields"
	switch kType.Type {
	case typUnion:
		var enumValues []string
		for _, t := range kType.UnionTypes {
			if isLiteral, _, val := gen.IsLitType(t); isLiteral {
				enumValues = append(enumValues, strings.Trim(val, "\""))
			}
		}
		if len(enumValues) > 0 {
			return fmt.Sprintf("// +kubebuilder:validation:Enum=%s", strings.Join(enumValues, ";"))
		}
		return anyTypeComment
	case typAny:
		return anyTypeComment
	case typList:
		if kType.Item.Type == typAny {
			return anyTypeComment
		}
	case typDict:
		if kType.Item.Type == typAny {
			return anyTypeComment
		}
	}
	return ""
}

// isTypeSchema determines if the provided KCL type represents a schema and,
// if applicable, returns nested schema types.
func isTypeSchema(kType *api.KclType) (bool, map[string]*api.KclType) {
	switch kType.Type {
	case typUnion:
		nestedSchemas := make(map[string]*api.KclType)
		found := false
		for _, t := range kType.UnionTypes {
			if t.Type == typSchema {
				found = true
				nestedSchemas[t.SchemaName] = proto.Clone(t).(*api.KclType)
			}
		}
		if found {
			return true, nestedSchemas
		}
	case typList:
		if kType.Item.Type == typSchema {
			return true, nil
		}
	case typDict:
		if kType.Item.Type == typSchema {
			return true, nil
		}
	case typSchema:
		return true, nil
	}
	return false, nil
}

// capitalize returns the input string with its first letter capitalized.
func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
