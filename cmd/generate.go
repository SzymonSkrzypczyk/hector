package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	dataFile  string
	theme     string
	outputDir string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a CV from a data file",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Data:", dataFile)
		fmt.Println("Theme:", theme)
		fmt.Println("Output:", outputDir)

		return nil
	},
}

func init() {
	generateCmd.Flags().StringVarP(
		&dataFile,
		"data",
		"d",
		"",
		"Path to CV data file (YAML or JSON)",
	)

	generateCmd.Flags().StringVarP(
		&theme,
		"theme",
		"t",
		"",
		"Theme to use",
	)

	generateCmd.Flags().StringVarP(
		&outputDir,
		"output-dir",
		"o",
		"",
		"Output directory",
	)

	// 👇 make --data required
	generateCmd.MarkFlagRequired("data")
}
