package pipeline

import (
	"hector/internal/parser"
	"hector/internal/pdf"
	"hector/internal/render"
	"hector/internal/schemas"
	"hector/internal/theme"
)

const (
	defaultOutputDirectory = "output/"
)

func GenerateHTML(selectedThemeName, valuesPath string) (string, schemas.ThemeSchema) {
	selectedTheme := theme.SelectTheme(selectedThemeName)
	exampleStruct := parser.Parse(valuesPath, selectedTheme)
	rendered := render.Render(exampleStruct)

	return rendered, exampleStruct
}

func Pipeline(selectedThemeName, valuesPath, outputPath string) {
	rendered, exampleStruct := GenerateHTML(selectedThemeName, valuesPath)
	if outputPath != "" {
		pdf.GeneratePDF(rendered, exampleStruct.PDFDetails(), outputPath)
	} else {
		pdf.GeneratePDF(rendered, exampleStruct.PDFDetails(), defaultOutputDirectory)
	}
}
