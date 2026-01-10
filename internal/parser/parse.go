package parser

import (
	"gopkg.in/yaml.v3"
	"hector/internal/schemas"
	"log"
	"os"
)

func Parse(filePath string, schema schemas.ThemeSchema) schemas.ThemeSchema {
	f, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	err = yaml.Unmarshal(f, schema)

	if err != nil {
		log.Fatalf("Error parsing YAML: %s", err)
	}

	log.Printf("Unmarshaled %v schema\n", schema.ThemeName())

	return schema
}
