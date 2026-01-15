package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hector/internal/parser"
	"hector/internal/schemas"
	valtheme "hector/internal/theme"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const defaultTemplateValidation = "normal"

var (
	validationDataFile string
	validationTemplate string
)

func humanizeKind(kind string) string {
	return strings.Title(strings.ReplaceAll(kind, "_", " "))
}

func humanizeFieldPath(path string) string {
	path = strings.ReplaceAll(path, ".", " → ")
	path = regexp.MustCompile(`\[(\d+)\]`).ReplaceAllStringFunc(path, func(m string) string {
		i, _ := strconv.Atoi(m[1 : len(m)-1])
		return fmt.Sprintf(" (item %d)", i+1)
	})
	return path
}

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
				log.Printf("\t- %s:\n", humanizeKind(kind))
				for _, field := range fields {
					log.Printf("\t\t• %s", humanizeFieldPath(field.Name))
					log.Printf("\t\t  Problem: %s", field.Reason)
					log.Printf("\t\t  Fix: %s", field.Suggestion)
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
