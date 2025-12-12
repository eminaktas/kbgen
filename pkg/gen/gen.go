package gen

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
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
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
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

	// Use a queue to process schemas recursively, ensuring a deterministic order
	schemasQueue := make([]string, 0, len(schemaMap))
	for name := range schemaMap {
		schemasQueue = append(schemasQueue, name)
	}
	sort.Strings(schemasQueue) // Ensure consistent processing order

	structData := make(map[string][]string) // Temporary storage for struct definitions
	imports := make(map[string]string)      // Track imports separately

	for len(schemasQueue) > 0 {
		// Always process in sorted order
		sort.Strings(schemasQueue)
		name := schemasQueue[0]
		schemasQueue = schemasQueue[1:]
		schema := schemaMap[name]

		// Skip if already processed
		if processed[name] {
			continue
		}
		processed[name] = true

		baseFileName := schema.GetFilename()
		if baseFileName == "" && schema.Item != nil {
			baseFileName = schema.Item.GetFilename()
		}
		filename := strings.TrimSuffix(filepath.Base(baseFileName), filepath.Ext(baseFileName))

		// Check if this schema should be generated as a map type.
		if indexSignature := schema.GetIndexSignature(); indexSignature != nil {
			keyGoType, err := gr.kclTypeToGoType(indexSignature.GetKey())
			if err != nil {
				return nil, err
			}
			valGoType, err := gr.kclTypeToGoType(indexSignature.Val)
			if err != nil {
				return nil, err
			}
			var mapTypeDef string
			if schema.GetSchemaName() == name {
				mapTypeDef = fmt.Sprintf(
					"type %s map[%s]%s\n",
					name,
					keyGoType,
					valGoType,
				)
			}
			structData[filename] = append(structData[filename], mapTypeDef)
			continue // Skip further processing for this schema type
		}

		// Retrieve schema properties
		var fields map[string]*api.KclType
		if schema.GetProperties() != nil {
			fields = schema.GetProperties()
		} else if schema.Item != nil && schema.Item.GetProperties() != nil {
			fields = schema.Item.GetProperties()
		} else {
			return nil, fmt.Errorf("unknown schema type for %s", name)
		}

		// Handle BaseSchema
		baseSchema := schema.BaseSchema
		if baseSchema != nil {
			if _, exists := schemaMap[baseSchema.SchemaName]; !exists {
				schemaMap[baseSchema.SchemaName] = baseSchema
				schemasQueue = append(schemasQueue, baseSchema.SchemaName)
			}
			// Remove baseSchema properties from the struct
			if baseSchema.Properties != nil {
				for baseField := range baseSchema.Properties {
					delete(fields, baseField)
				}
			}
		}

		// Collect field names and sort them for deterministic order
		fieldNames := make([]string, 0, len(fields))
		for fieldName := range fields {
			if !strings.HasPrefix(fieldName, "_") { // Ignore non-exported fields
				fieldNames = append(fieldNames, fieldName)
			}
		}
		sort.Strings(fieldNames) // Ensure consistent order

		// Build struct definition
		var structDefBuilder strings.Builder
		structDefBuilder.WriteString(fmt.Sprintf("type %s struct {\n", name))

		for _, fieldName := range fieldNames {
			fieldType := fields[fieldName]

			// Process nested schemas
			if isSchema, nestedTypes := isTypeSchema(fieldType); isSchema {
				// Determine whether this is an index signature case.
				possibleIndexSignature := true
				if nestedTypes == nil {
					if fieldType.Item != nil {
						if _, exists := schemaMap[fieldType.Item.SchemaName]; !exists {
							schemaMap[fieldType.Item.SchemaName] = proto.Clone(fieldType).(*api.KclType)
							schemasQueue = append(schemasQueue, fieldType.Item.SchemaName)
							possibleIndexSignature = false
						}
					} else {
						if _, exists := schemaMap[fieldType.SchemaName]; !exists {
							schemaMap[fieldType.SchemaName] = proto.Clone(fieldType).(*api.KclType)
							schemasQueue = append(schemasQueue, fieldType.SchemaName)
							if fieldType.Item != nil || fieldType.Properties != nil {
								schemaMap[fieldType.SchemaName] = proto.Clone(fieldType).(*api.KclType)
								possibleIndexSignature = false
							}
						}
					}
				} else {
					for nestedName, nestedSchema := range nestedTypes {
						if _, exists := schemaMap[nestedName]; !exists {
							schemaMap[nestedName] = nestedSchema
							possibleIndexSignature = false
							schemasQueue = append(schemasQueue, nestedName)
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

			goType, err := gr.kclTypeToGoType(fieldType)
			if err != nil {
				return nil, err
			}

			// Ensure necessary imports are included (store in separate map)
			if strings.Contains(goType, gr.config.CustomAnyType.Type) {
				if _, exists := imports[filename]; !exists {
					imports[filename] = fmt.Sprintf(importTemplate, gr.config.CustomAnyType.Import)
				}
			}

			// Append field comment and kubebuilder validation markers
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

			// Append the field definition
			structDefBuilder.WriteString(fmt.Sprintf("\t%s %s `json:\"%s%s\" yaml:\"%s%s\"`\n",
				capitalize(fieldName), goType, fieldName, omitEmpty, fieldName, omitEmpty))
		}

		// Handle BaseSchema inclusion
		if baseSchema != nil {
			structDefBuilder.WriteString(fmt.Sprintf("\n\t%s `json:\",inline\" yaml:\",inline\"`\n", baseSchema.SchemaName))
		}

		structDefBuilder.WriteString("}\n")
		structData[filename] = append(structData[filename], structDefBuilder.String())

		// Append DeepCopy function if needed
		if isSchema, _ := isTypeSchema(schema); isSchema {
			if deepCopyFunc, err := generateDeepCopyFuncs(name); err == nil {
				structData[filename] = append(structData[filename], deepCopyFunc)
			} else {
				return nil, err
			}
		}
	}

	// Ensure deterministic ordering of the final result (excluding imports)
	for filename, structs := range structData {
		sort.Strings(structs) // Sort generated struct contents

		// Insert import at the top if it exists
		if importStmt, exists := imports[filename]; exists {
			result[filename] = append([]string{importStmt}, structs...)
		} else {
			result[filename] = structs
		}
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
