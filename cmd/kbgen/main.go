package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	goversion "sigs.k8s.io/release-utils/version"
)

var (
	commit    string // Injected at build time via -ldflags
	treeState string // Injected at build time via -ldflags
	date      string // Injected at build time via -ldflags
	version   string // Injected at build time via -ldflags
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildVersion())
	},
}

var rootCmd = &cobra.Command{
	Use:   "kbgen",
	Short: "kbgen is a Go struct generator for Kubebuilder from KCL schemas",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generatorCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// buildVersion returns a populated string with custom metadata
func buildVersion() string {
	// Initialize with default version info
	goVersion := goversion.GetVersionInfo()

	// Customize version info
	goVersion.Name = "kbgen"
	goVersion.Description = "kbgen is a Go struct generator for Kubebuilder from KCL schemas"
	goVersion.ASCIIName = "true"    // Flag for ASCII art representation
	goVersion.FontName = "standard" // Font used for ASCII art

	// Override fields with build-time injected variables
	goVersion.GitCommit = commit
	goVersion.GitTreeState = treeState
	goVersion.BuildDate = date
	goVersion.GitVersion = version

	return goVersion.String()
}
