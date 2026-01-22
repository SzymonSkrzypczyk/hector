package pdf

import (
	hlog "hector/internal/log"
	"hector/internal/schemas"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func clearOutputDirectory(outputDirectory string) {
	files, err := os.ReadDir(outputDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".pdf" {
			err := os.Remove(filepath.Join(outputDirectory, file.Name()))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func GeneratePDF(htmlPath string, styleGuide schemas.PDFSchema, outputDirectory string) error {
	if err := os.MkdirAll(outputDirectory, 0750); err != nil {
		log.Fatalln(err)
	}

	absHTML, err := filepath.Abs(htmlPath)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(absHTML); os.IsNotExist(err) {
		return fmt.Errorf("HTML file not found at %s", absHTML)
	}
	hlog.Debug("Loading HTML from: %s\n", absHTML)

	fileURL := "file:///" + filepath.ToSlash(absHTML)

	l := launcher.New().
		Headless(true).
		Set("disable-gpu", "true").
		Set("no-sandbox", "true")

	url, err := l.Launch()
	if err != nil {
		return fmt.Errorf("failed to launch browser (missing dependencies?): %w", err)
	}

	hlog.Debug("Browser launched. Connecting...\n")
	browser := rod.New().ControlURL(url)
	err = browser.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to browser: %w", err)
	}
	defer browser.MustClose()

	page, err := browser.Page(proto.TargetCreateTarget{URL: fileURL})
	if err != nil {
		return fmt.Errorf("failed to open page: %w", err)
	}

	err = page.WaitLoad()
	if err != nil {
		return fmt.Errorf("failed to load page: %w", err)
	}

	targetFile := filepath.Join(outputDirectory, "result_cv.pdf")
	hlog.Info("Page loaded. Generating PDF...\n")

	pdfStream, err := page.PDF(&proto.PagePrintToPDF{
		PrintBackground: true,
		PaperWidth:      floatPtr(8.27), // A4
		PaperHeight:     floatPtr(11.7), // A4
		MarginTop:       floatPtr(styleGuide.MarginTop),
		MarginBottom:    floatPtr(styleGuide.MarginBottom),
		MarginLeft:      floatPtr(styleGuide.MarginLeft),
		MarginRight:     floatPtr(styleGuide.MarginRight),
		Scale:           floatPtr(styleGuide.Scale),
	})
	hlog.Debug("PDF Stream generated. Reading data...\n")
	if err != nil {
		return fmt.Errorf("failed to generate PDF command: %w", err)
	}

	pdfData, err := io.ReadAll(pdfStream)
	if err != nil {
		return fmt.Errorf("failed to read PDF stream: %w", err)
	}

	if err := os.WriteFile(targetFile, pdfData, 0600); err != nil {
		return fmt.Errorf("failed to save PDF file: %w", err)
	}

	clearOutputDirectory(outputDirectory)
	hlog.Info("Success! PDF generated at: %s\n", targetFile)
	return nil
}

func floatPtr(v float64) *float64 {
	return &v
}
