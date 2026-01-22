package render

import (
	"fmt"
	hlog "hector/internal/log"
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

	srcDir := filepath.Dir(absSrc)
	srcFile := filepath.Base(absSrc)

	hlog.Debug("Copying asset: %s to %s\n", srcFile, outputDir)

	rootDir, err := os.OpenRoot(srcDir)
	if err != nil {
		return fmt.Errorf("failed to open source directory %s: %v", srcDir, err)
	}
	defer func() {
		if err := rootDir.Close(); err != nil {
			log.Printf("failed to close source directory: %v", err)
		}
	}()

	if _, err := rootDir.Stat(srcFile); err != nil {
		return fmt.Errorf("asset not found: %s", absSrc)
	}

	in, err := rootDir.Open(srcFile)
	if err != nil {
		return fmt.Errorf("asset not found: %s", absSrc)
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(in)

	outRoot, err := os.OpenRoot(outputDir)
	if err != nil {
		return fmt.Errorf("failed to open output directory %s: %v", outputDir, err)
	}
	defer func() {
		if err := outRoot.Close(); err != nil {
			log.Printf("failed to close output directory: %v", err)
		}
	}()

	out, err := outRoot.Create(copiedPhotoName)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s in %s: %v", copiedPhotoName, outputDir, err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(out)

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("failed to copy asset from %s to %s: %v", absSrc, filepath.Join(outputDir, copiedPhotoName), err)
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
	hlog.Debug("Template directory: %s\n", templateDirectory)

	htmlFile := filepath.Join(templateDirectory, details.HTMLTemplate)
	cssFile := filepath.Join(templateDirectory, details.CSSFile)

	dir, err := os.ReadDir(outputDirectory)
	if err != nil || dir == nil {
		err := os.MkdirAll(outputDirectory, 0750)
		if err != nil {
			log.Fatalf("Failed to create output directory: %s", err)
		}
	}

	outRoot, err := os.OpenRoot(outputDirectory)
	if err != nil {
		log.Fatalf("Failed to open output directory: %s", err)
	}
	defer func() {
		if err := outRoot.Close(); err != nil {
			log.Printf("failed to close output directory: %v", err)
		}
	}()

	hf, err := outRoot.Create(outputFile)
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
	srcRoot, err := os.OpenRoot(templateDirectory)
	if err != nil {
		log.Fatalf("warning: could not open theme directory (%s): %v", templateDirectory, err)
	}
	defer func() {
		if err := srcRoot.Close(); err != nil {
			log.Printf("warning: failed to close theme directory: %v", err)
		}
	}()

	hlog.Debug("Copying CSS file: %s from %s\n", details.CSSFile, templateDirectory)
	src, err := srcRoot.Open(details.CSSFile)
	if err != nil {
		log.Fatalf("warning: could not open theme style.css (%s): %v", cssFile, err)
	} else {
		defer func(src *os.File) {
			err := src.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(src)

		dstRoot, err := os.OpenRoot(outputDirectory)
		if err != nil {
			log.Fatalf("warning: could not open output directory (%s): %v", outputDirectory, err)
		}
		defer func() {
			if err := dstRoot.Close(); err != nil {
				log.Printf("warning: failed to close output directory: %v", err)
			}
		}()

		dst, err := dstRoot.Create(details.CSSFile)
		if err != nil {
			dstPath := filepath.Join(outputDirectory, details.CSSFile)
			log.Fatalf("warning: could not create destination css file (%s): %v", dstPath, err)
		} else {
			defer func() {
				if err := dst.Close(); err != nil {
					dstPath := filepath.Join(outputDirectory, details.CSSFile)
					log.Fatalf("warning: could not close destination css file (%s): %v", dstPath, err)
				}
			}()
			if _, err := io.Copy(dst, src); err != nil {
				dstPath := filepath.Join(outputDirectory, details.CSSFile)
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

	hlog.Info("HTML resume generated at: %s\n", filepath.Join(outputDirectory, outputFile))

	return filepath.Join(outputDirectory, outputFile)
}
