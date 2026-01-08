package parser

import (
	"gopkg.in/yaml.v3"
	"hector/internal/schemas"
	"log"
	"os"
)

func ParseNormal(filePath string) schemas.Normal {
	f, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	var normalSchema schemas.Normal

	err = yaml.Unmarshal(f, &normalSchema)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Unmarshaled Normal Schema")

	return normalSchema
}
