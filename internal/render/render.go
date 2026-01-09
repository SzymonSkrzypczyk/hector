package render

import (
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
	details := ExtractFiles(values)
	templateDirectory := filepath.Join(themePath, details.ThemeDirectory)

	htmlFile := filepath.Join(templateDirectory, details.HTMLTemplate)
	cssFile := filepath.Join(templateDirectory, details.CSSFile)

	dir, err := os.ReadDir(outputDirectory)
	if err != nil || dir == nil {
		err := os.MkdirAll(outputDirectory, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}

	hf, err := os.Create(filepath.Join(outputDirectory, outputFile))
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if cerr := hf.Close(); cerr != nil {
			log.Printf("warning: failed to close file %s: %v", filepath.Join(outputDirectory, outputFile), cerr)
		}
	}()

	tmpl, err := template.ParseFiles(htmlFile)
	if err != nil {
		log.Fatalln(err)
	}

	if err := tmpl.Execute(hf, values); err != nil {
		log.Fatalln(err)
	}

	src, err := os.Open(cssFile)
	if err != nil {
		log.Printf("warning: could not open theme style.css (%s): %v", cssFile, err)
	} else {
		defer src.Close()
		dstPath := filepath.Join(outputDirectory, details.CSSFile)
		dst, err := os.Create(dstPath)
		if err != nil {
			log.Fatalln(err)
		} else {
			defer func() {
				if err := dst.Close(); err != nil {
					log.Fatalln(err)
				}
			}()
			if _, err := io.Copy(dst, src); err != nil {
				log.Fatalln(err)
			}
		}
	}

	log.Println("HTML resume generated at:", filepath.Join(outputDirectory, outputFile))

	return filepath.Join(outputDirectory, outputFile)
}
