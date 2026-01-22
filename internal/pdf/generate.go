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

func GeneratePDF(htmlPath string, styleGuide schemas.PDFSchema, outputDirectory string) {
	if err := os.MkdirAll(outputDirectory, 0750); err != nil {
		log.Fatalln(err)
	}

	absHTML, err := filepath.Abs(htmlPath)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(absHTML); os.IsNotExist(err) {
		log.Fatalln("Error: HTML file not found at", absHTML)
	}
	hlog.Debug("Loading HTML from: %s\n", absHTML)

	fileURL := "file:///" + filepath.ToSlash(absHTML)

	l := launcher.New().
		Headless(true).
		Set("disable-gpu", "true").
		Set("no-sandbox", "true")

	url, err := l.Launch()
	if err != nil {
		hlog.Fatal("Failed to launch browser. This is often due to missing dependencies on the system.\nError: %v\nSee https://go-rod.github.io/#/compatibility?id=os for help.\n", err)
	}

	hlog.Debug("Browser launched. Connecting...\n")
	browser := rod.New().ControlURL(url)
	err = browser.Connect()
	if err != nil {
		hlog.Fatal("Failed to connect to browser: %v", err)
	}
	defer browser.MustClose()

	page, err := browser.Page(proto.TargetCreateTarget{URL: fileURL})
	if err != nil {
		hlog.Fatal("Failed to open page: %v", err)
	}
	
	err = page.WaitLoad()
	if err != nil {
		hlog.Fatal("Failed to load page: %v", err)
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
		log.Fatalln("Failed to generate PDF command:", err)
	}

	pdfData, err := io.ReadAll(pdfStream)
	if err != nil {
		log.Fatalln("Failed to read PDF stream:", err)
	}

	if err := os.WriteFile(targetFile, pdfData, 0600); err != nil {
		log.Fatalln("Failed to save PDF file:", err)
	}

	clearOutputDirectory(outputDirectory)
	hlog.Info("Success! PDF generated at: %s\n", targetFile)
}

func floatPtr(v float64) *float64 {
	return &v
}
