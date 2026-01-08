package pipeline

import (
	"hector/internal/parser"
	"hector/internal/pdf"
	"hector/internal/render"
)

// it will be hardcoded for now to use normal schema
// ultimately it will be passed by CLI arg
const (
	hardcodedExamplePath = "examples/example.yaml"
)

func Pipeline() {
	exampleStruct := parser.ParseNormal(hardcodedExamplePath)
	rendered := render.NormalRender(exampleStruct)
	pdf.GenerateNormalPDF(rendered)
}
