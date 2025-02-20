package main

import (
	"fmt"
	"os"

	"github.com/eminaktas/kbgen/pkg/gen"
	"github.com/spf13/cobra"
)

var (
	outputDir      string
	packageName    string
	kclProgramPath string
	directory      string
	configPath     string
)

// generatorCmd represents the command to generate Go structs from KCL schemas.
var generatorCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Go structs from KCL schemas for Kubebuilder",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the generator with provided flags.
		generator, err := gen.NewGeneratorWithPath(packageName, kclProgramPath, directory, outputDir, configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing generator: %v\n", err)
			os.Exit(1)
		}

		// Run the generation process.
		if err := generator.Generate(); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
			os.Exit(1)
		}

		// TODO: Run betteralign to the generated code.
	},
}

func init() {
	generatorCmd.Flags().StringVar(&outputDir, "outputDir", "./", "Output directory for generated Go structs")
	generatorCmd.Flags().StringVar(&kclProgramPath, "programPath", "./", "Program path where the module file exists")
	generatorCmd.Flags().StringVar(&directory, "directory", "", "Directory containing .k files to generate schemas from")
	generatorCmd.Flags().StringVar(&packageName, "packageName", "", "Package name for the generated Go structs")
	generatorCmd.Flags().StringVar(&configPath, "configPath", "", "Path to the configuration file")
}
