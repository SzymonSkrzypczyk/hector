package pdf

import (
	"hector/internal/schemas"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGeneratePDF(t *testing.T) {
	tempDir := t.TempDir()
	htmlPath := filepath.Join(tempDir, "index.html")
	htmlContent := []byte(`
		<html>
		<body style="background-color: white;">
			<h1>Test PDF Generation</h1>
			<p>This is a test.</p>
		</body>
		</html>`)

	if err := os.WriteFile(htmlPath, htmlContent, 0644); err != nil {
		t.Fatal(err)
	}

	style := schemas.PDFSchema{
		MarginTop:    0,
		MarginBottom: 0,
		MarginLeft:   0,
		MarginRight:  0,
		Scale:        1.0,
	}

	outputDir := filepath.Join(tempDir, "output")

	GeneratePDF(htmlPath, style, outputDir)

	expectedPDF := filepath.Join(outputDir, "result_cv.pdf")
	info, err := os.Stat(expectedPDF)
	if os.IsNotExist(err) {
		t.Errorf("PDF was not generated at %s", expectedPDF)
	}

	if info.Size() == 0 {
		t.Error("Generated PDF is empty")
	}
}

func TestClearOutputDirectory(t *testing.T) {
	dir := t.TempDir()

	// Create a .pdf file (should stay)
	err := os.WriteFile(filepath.Join(dir, "keep.pdf"), []byte("data"), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a .html file (should be removed)
	err = os.WriteFile(filepath.Join(dir, "remove.html"), []byte("data"), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	clearOutputDirectory(dir)

	files, _ := os.ReadDir(dir)
	if len(files) != 1 {
		t.Errorf("Expected 1 file remaining, got %d", len(files))
	}
	if files[0].Name() != "keep.pdf" {
		t.Errorf("Expected keep.pdf to remain, got %s", files[0].Name())
	}
}
