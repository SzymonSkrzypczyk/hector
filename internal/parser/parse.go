package parser

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"hector/internal/schemas"
)

func Parse(filePath string, schema schemas.ThemeSchema) schemas.ThemeSchema {
	dir := filepath.Dir(filePath)
	base := filepath.Base(filePath)

	root, err := os.OpenRoot(dir)
	if err != nil {
		log.Fatalf("Error opening directory %s: %s", dir, err)
	}
	defer root.Close()

	file, err := root.Open(base)
	if err != nil {
		log.Fatalf("Error opening file %s: %s", base, err)
	}
	defer file.Close()

	f, err := io.ReadAll(file)
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
