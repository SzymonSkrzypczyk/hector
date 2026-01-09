package pipeline

import (
	"hector/internal/parser"
	"hector/internal/pdf"
	"hector/internal/render"
	"hector/internal/theme"
)

// it will be hardcoded for now to use normal schema
// ultimately it will be passed by CLI arg
const (
	defaultOutputDirectory = "output/"
)

func Pipeline(selectedThemeName, valuesPath, outputPath string) {
	selectedTheme := theme.SelectTheme(selectedThemeName)
	exampleStruct := parser.Parse(valuesPath, selectedTheme)
	rendered := render.Render(exampleStruct)
	if outputPath != "" {
		pdf.GeneratePDF(rendered, exampleStruct.PDFDetails(), outputPath)
	} else {
		pdf.GeneratePDF(rendered, exampleStruct.PDFDetails(), defaultOutputDirectory)
	}
}
