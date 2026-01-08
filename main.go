package main

import (
	"fmt"
	"hector/internal/parser"
	"hector/internal/pdf"
	"hector/internal/render"
)

func main() {
	exampleStruct := parser.ParseNormal("examples/example.yaml")
	rendered := render.NormalRender(exampleStruct)
	pdf.GenerateNormalPDF("output/" + rendered)

	fmt.Println(rendered)
}
