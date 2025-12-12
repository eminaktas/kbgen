package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	defaultImport  = "apiextensionsv1 \"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1\""
	defaultAnyType = "*apiextensionsv1.JSON"
)

// Config holds the configuration settings loaded from a YAML file.
type Config struct {
	// CustomAnyType configures a custom type and its import for handling "any" type.
	CustomAnyType struct {
		Type   string `yaml:"type,omitempty"`
		Import string `yaml:"import,omitempty"`
	} `yaml:"customAnyType,omitempty"`
}

// LoadConfig reads the YAML configuration from the provided file path and unmarshals it into a Config struct.
func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}
	var err error
	if path != "" {
		cfg, err = LoadConfigFromFile(path)
		if err != nil {
			return nil, err
		}
	}

	// Set default values for custom type if not provided.
	if cfg.CustomAnyType.Type == "" {
		cfg.CustomAnyType.Type = defaultAnyType
	}
	if cfg.CustomAnyType.Import == "" {
		cfg.CustomAnyType.Import = defaultImport
	}
	return cfg, nil
}

// LoadConfigFromFile reads the YAML configuration from the provided file path and unmarshals it into a Config struct.
func LoadConfigFromFile(filePath string) (*Config, error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(absFilePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
