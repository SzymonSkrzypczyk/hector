package parser

import (
	"io"
	"log"
	"os"
	"path/filepath"

	hlog "hector/internal/log"

	"gopkg.in/yaml.v3"
	"hector/internal/schemas"
)

func Parse(filePath string, schema schemas.ThemeSchema) schemas.ThemeSchema {
	dir := filepath.Dir(filePath)
	base := filepath.Base(filePath)

	hlog.Debug("Parsing data from: %s\n", filePath)

	root, err := os.OpenRoot(dir)
	if err != nil {
		log.Fatalf("Error opening directory %s: %s", dir, err)
	}
	defer func() {
		if err := root.Close(); err != nil {
			log.Printf("Error closing directory: %v\n", err)
		}
	}()

	file, err := root.Open(base)
	if err != nil {
		log.Fatalf("Error opening file %s: %s", base, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v\n", err)
		}
	}()

	f, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	err = yaml.Unmarshal(f, schema)

	if err != nil {
		log.Fatalf("Error parsing YAML: %s", err)
	}

	hlog.Info("Unmarshaled %v schema\n", schema.ThemeName())

	return schema
}
