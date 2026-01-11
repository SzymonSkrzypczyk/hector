package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

const defaultTemplateValidation = "normal"

var (
	validationDataFile string
	validationTemplate string
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the CV data file",
	Run: func(cmd *cobra.Command, args []string) {
		if validationDataFile == "" {
			log.Fatalln("Error: --data flag is required")
		}

		if validationTemplate == "" {
			validationTemplate = defaultTemplateValidation
		}

	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(
		&validationDataFile,
		"data",
		"d",
		"",
		"Data file to be validated",
	)

	validateCmd.Flags().StringVarP(
		&validationTemplate,
		"template",
		"t",
		"normal",
		"Template to validate against",
	)

	validateCmd.MarkFlagRequired("data")
}
