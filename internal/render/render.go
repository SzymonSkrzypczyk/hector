package render

import (
	"fmt"
	"hector/internal/schemas"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

const (
	themePath       = "themes"
	outputDirectory = "output/"
	outputFile      = "normal_resume.html"
	copiedPhotoName = "photo.jpg"
)

func copyAsset(srcPath, outputDir string) error {
	absSrc, err := filepath.Abs(srcPath)
	if err != nil {
		return err
	}

	if _, err := os.Stat(absSrc); err != nil {
		return fmt.Errorf("asset not found: %s", absSrc)
	}

	dst := filepath.Join(outputDir, copiedPhotoName)

	in, err := os.Open(absSrc)
	if err != nil {
		return fmt.Errorf("asset not found: %s", absSrc)
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {

		}
	}(in)

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %v", dst, err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("failed to copy asset from %s to %s: %v", absSrc, dst, err)
	}

	return nil
}

func GetPhotoFromTheme(exampleStruct schemas.ThemeSchema) string {
	val := reflect.ValueOf(exampleStruct)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName("Photo")
	if field.IsValid() && field.Kind() == reflect.String {
		return field.String()
	}

	return ""
}

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

	// parse HTML template
	tmpl, err := template.ParseFiles(htmlFile)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %s", err)
	}

	if err := tmpl.Execute(hf, values); err != nil {
		log.Fatalf("Failed to render HTML template: %s", err)
	}

	// copy css file
	src, err := os.Open(cssFile)
	if err != nil {
		log.Fatalf("warning: could not open theme style.css (%s): %v", cssFile, err)
	} else {
		defer func(src *os.File) {
			err := src.Close()
			if err != nil {

			}
		}(src)
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

	// copy photo if exists
	if photoPath := GetPhotoFromTheme(values); photoPath != "" {
		err := copyAsset(photoPath, outputDirectory)
		if err != nil {
			log.Printf("warning: could not copy photo asset (%s): %v", photoPath, err)
		}
	}

	log.Println("HTML resume generated at:", filepath.Join(outputDirectory, outputFile))

	return filepath.Join(outputDirectory, outputFile)
}
