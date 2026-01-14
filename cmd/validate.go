package cmd

import (
	"github.com/spf13/cobra"
	"hector/internal/parser"
	"hector/internal/schemas"
	valtheme "hector/internal/theme"
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

		selectedTheme := valtheme.SelectTheme(validationTemplate)
		exampleStruct := parser.Parse(validationDataFile, selectedTheme)
		isCorrect, missing := exampleStruct.Validate()

		if isCorrect {
			log.Println("Validation successful: No flaws detected.")
		} else {
			log.Println("Validation failed:")
			grouped := schemas.GroupByKind(missing)
			for kind, fields := range grouped {
				log.Printf("\t%s:\n", kind)
				for _, field := range fields {
					log.Printf("\t\t- %s\n", field)
				}
			}
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
		"theme",
		"t",
		"normal",
		"Template to validate against",
	)

	validateCmd.MarkFlagRequired("data")
}
