package render

import (
	"hector/internal/schemas"
	"os"
	"path/filepath"
	"testing"
)

// MockTheme for Render testing
type MockTheme struct {
	Title string
	Photo string
}

// Correct Validate() signature to match schemas.ThemeSchema
func (m *MockTheme) Validate() (bool, []schemas.ValidationFlaw) {
	return true, nil
}

func (m *MockTheme) ThemeName() string { return "mock_theme" }

func (m *MockTheme) TemplateDetails() schemas.FileSchema {
	return schemas.FileSchema{
		ThemeDirectory: "mock_theme",
		HTMLTemplate:   "index.html",
		CSSFile:        "style.css",
	}
}

func (m *MockTheme) PDFDetails() schemas.PDFSchema { return schemas.PDFSchema{} }

func TestRender(t *testing.T) {
	cwd, _ := os.Getwd()
	themeDir := filepath.Join(cwd, "themes", "mock_theme")
	outputDir := filepath.Join(cwd, "output")

	defer func() {
		os.RemoveAll(filepath.Join(cwd, "themes"))
		os.RemoveAll(outputDir)
	}()

	// Prepare mock theme files
	err := os.MkdirAll(themeDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create mock theme dir: %v", err)
	}

	htmlContent := []byte(`<html><body><h1>{{.Title}}</h1></body></html>`)
	err = os.WriteFile(filepath.Join(themeDir, "index.html"), htmlContent, 0644)
	if err != nil {
		t.Fatal(err)
	}

	cssContent := []byte(`body { color: red; }`)
	err = os.WriteFile(filepath.Join(themeDir, "style.css"), cssContent, 0644)
	if err != nil {
		t.Fatal(err)
	}

	mockData := &MockTheme{Title: "Hello Hector"}

	outputPath := Render(mockData)

	absOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		t.Fatalf("Failed to resolve absolute path of output: %v", err)
	}

	expectedOutput := filepath.Join(outputDir, "normal_resume.html")

	if absOutputPath != expectedOutput {
		t.Errorf("Path mismatch.\nExpected: %s\nGot:      %s", expectedOutput, absOutputPath)
	}

	content, err := os.ReadFile(absOutputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	if string(content) != "<html><body><h1>Hello Hector</h1></body></html>" {
		t.Errorf("HTML content mismatch. Got: %s", string(content))
	}

	if _, err := os.Stat(filepath.Join(outputDir, "style.css")); os.IsNotExist(err) {
		t.Error("style.css was not copied to output directory")
	}
}
