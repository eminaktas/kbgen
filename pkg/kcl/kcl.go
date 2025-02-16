package kcl

import (
	"fmt"

	"kcl-lang.io/kcl-go/pkg/spec/gpyrpc"
	"kcl-lang.io/lib/go/api"
	"kcl-lang.io/lib/go/native"
)

// KCL encapsulates the KCL service client and its base configuration.
type KCL struct {
	Client      api.ServiceClient // Service client for making KCL API calls.
	programPath string            // Base directory or path for the KCL program.
}

// NewKCL returns a new KCL instance with a native service client.
func NewKCL() *KCL {
	return &KCL{
		Client: native.NewNativeServiceClient(),
	}
}

// NewKCLWithPath returns a new KCL instance with the given program path.
func NewKCLWithPath(path string) *KCL {
	k := NewKCL()
	k.programPath = path
	return k
}

// updateDependencies updates KCL module dependencies based on the programPath.
// It returns a list of external packages needed by the KCL program.
func (k *KCL) updateDependencies() ([]*api.ExternalPkg, error) {
	// Call the native service client to update dependencies.
	resp, err := k.Client.UpdateDependencies(&gpyrpc.UpdateDependencies_Args{
		ManifestPath: k.programPath,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update dependencies: %w", err)
	}
	return resp.GetExternalPkgs(), nil
}

// GetSchemaTypeMapping retrieves the schema type mapping for the provided KCL files.
// It first updates the dependencies and then fetches the schema mapping using the service client.
func (k *KCL) GetSchemaTypeMapping(kFiles []string) (map[string]*api.KclType, error) {
	// Update external package dependencies.
	externalPkgs, err := k.updateDependencies()
	if err != nil {
		return nil, err
	}

	// Get the schema type mapping from the service client.
	result, err := k.Client.GetSchemaTypeMapping(&api.GetSchemaTypeMapping_Args{
		ExecArgs: &gpyrpc.ExecProgram_Args{
			KFilenameList: kFiles,
			WorkDir:       k.programPath,
			ExternalPkgs:  externalPkgs,
		},
	})
	if err != nil {
		return nil, err
	}

	return result.SchemaTypeMapping, nil
}
