package render

import (
	"fmt"
	"hector/internal/schemas"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	themePath       = "themes"
	outputDirectory = "output/"
	outputFile      = "normal_resume.html"
)

func Render(values schemas.ThemeSchema) string {
	details := values.TemplateDetails()
	templateDirectory := filepath.Join(themePath, details.ThemeDirectory)
	fmt.Println(templateDirectory)

	htmlFile := filepath.Join(templateDirectory, details.HTMLTemplate)
	cssFile := filepath.Join(templateDirectory, details.CSSFile)

	dir, err := os.ReadDir(outputDirectory)
	if err != nil || dir == nil {
		err := os.MkdirAll(outputDirectory, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create output directory: %s", err)
		}
	}

	hf, err := os.Create(filepath.Join(outputDirectory, outputFile))
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer func() {
		if err := hf.Close(); err != nil {
			log.Fatalf("warning: failed to close file %s: %v", filepath.Join(outputDirectory, outputFile), err)
		}
	}()

	tmpl, err := template.ParseFiles(htmlFile)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %s", err)
	}

	if err := tmpl.Execute(hf, values); err != nil {
		log.Fatalf("Failed to render HTML template: %s", err)
	}

	src, err := os.Open(cssFile)
	if err != nil {
		log.Fatalf("warning: could not open theme style.css (%s): %v", cssFile, err)
	} else {
		defer src.Close()
		dstPath := filepath.Join(outputDirectory, details.CSSFile)
		dst, err := os.Create(dstPath)
		if err != nil {
			log.Fatalf("warning: could not create destination css file (%s): %v", dstPath, err)
		} else {
			defer func() {
				if err := dst.Close(); err != nil {
					log.Fatalf("warning: could not close destination css file (%s): %v", dstPath, err)
				}
			}()
			if _, err := io.Copy(dst, src); err != nil {
				log.Fatalf("warning: could not copy css file (%s): %v", dstPath, err)
			}
		}
	}

	log.Println("HTML resume generated at:", filepath.Join(outputDirectory, outputFile))

	return filepath.Join(outputDirectory, outputFile)
}
