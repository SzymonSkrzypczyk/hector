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
	themePath       = "themes/normal"
	outputDirectory = "output/"
	outputFile      = "normal_resume.html"
	templateFile    = "template.html"
	styleFile       = "style.css"
)

func NormalRender(values schemas.Normal) string {
	htmlFile := filepath.Join(themePath, templateFile)
	cssFile := filepath.Join(themePath, styleFile)

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
		dstPath := filepath.Join(outputDirectory, styleFile)
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

	return outputFile
}
