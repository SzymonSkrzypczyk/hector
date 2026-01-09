package cmd

import (
	"github.com/spf13/cobra"
	"hector/internal/pipeline"
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
		pipeline.Pipeline(theme, dataFile, outputDir)

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
