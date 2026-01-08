package main

import (
	"fmt"
	"hector/internal/parser"
)

func main() {
	exampleStruct := parser.ParseNormal("examples/example.yaml")

	fmt.Println(exampleStruct)
}
