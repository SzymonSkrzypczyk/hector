package main

import (
	"fmt"
	"hector/internal/parser"
	"hector/internal/render"
)

func main() {
	exampleStruct := parser.ParseNormal("examples/example.yaml")
	rendered := render.NormalRender(exampleStruct)

	fmt.Println(rendered)
}
