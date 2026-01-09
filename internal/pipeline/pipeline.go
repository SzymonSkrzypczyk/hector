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
	hardcodedExamplePath = "examples/example.yaml"
	exampleSelectedTheme = "normal"
)

func Pipeline() {
	selectedTheme := theme.SelectTheme(exampleSelectedTheme)
	exampleStruct := parser.Parse(hardcodedExamplePath, selectedTheme)
	rendered := render.Render(exampleStruct)
	pdf.GeneratePDF(rendered, exampleStruct.PDFDetails())
}
