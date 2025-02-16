package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the configuration settings loaded from a YAML file.
type Config struct {
	// DeepCopyIntoNominee specifies schema names for which a DeepCopyInto function should be generated.
	DeepCopyIntoNominee []string `yaml:"deepCopyIntoNominee,omitempty"`

	// CustomAnyType configures a custom type and its import for handling "any" type.
	CustomAnyType struct {
		Type   string `yaml:"type,omitempty"`
		Import string `yaml:"import,omitempty"`
	} `yaml:"customAnyType,omitempty"`

	// MapTypeSchemas defines schemas that should be generated as map types.
	MapTypeSchemas []struct {
		// Name is the identifier for the schema.
		Name string `yaml:"name,omitempty"`
		// PkgPath indicates the package path for the schema.
		PkgPath string `yaml:"pkgPath,omitempty"`
		// Item holds details about the schema item.
		Item struct {
			// SchemaName is the name of the item schema.
			SchemaName string `yaml:"schemaName,omitempty"`
			// SchemaDoc contains documentation for the schema.
			SchemaDoc string `yaml:"schemaDoc,omitempty"`
		} `yaml:"item,omitempty"`
	} `yaml:"mapTypeSchemas,omitempty"`
}

// LoadConfig reads the YAML configuration from the provided file path and unmarshals it into a Config struct.
func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
