package pdf

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

const outputDirectory = "output/"

func GenerateNormalPDF(htmlPath string) {
	if err := os.MkdirAll(outputDirectory, 0755); err != nil {
		log.Fatalln(err)
	}

	absHTML, err := filepath.Abs(htmlPath)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(absHTML); os.IsNotExist(err) {
		log.Fatalln("Error: HTML file not found at", absHTML)
	}
	fmt.Println("Loading HTML from:", absHTML)

	fileURL := "file:///" + filepath.ToSlash(absHTML)

	l := launcher.New().
		Headless(true).
		Set("disable-gpu", "true").
		Set("no-sandbox", "true").
		MustLaunch()

	browser := rod.New().ControlURL(l).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(fileURL).MustWaitLoad()

	targetFile := filepath.Join(outputDirectory, "result_cv.pdf")
	fmt.Println("Page loaded. Generating PDF...")

	pdfStream, err := page.PDF(&proto.PagePrintToPDF{
		PrintBackground: true,
		PaperWidth:      floatPtr(8.27), // A4
		PaperHeight:     floatPtr(11.7), // A4
		MarginTop:       floatPtr(0.4),
		MarginBottom:    floatPtr(0.4),
		MarginLeft:      floatPtr(0.4),
		MarginRight:     floatPtr(0.4),
	})
	if err != nil {
		log.Fatalln("Failed to generate PDF command:", err)
	}

	pdfData, err := io.ReadAll(pdfStream)
	if err != nil {
		log.Fatalln("Failed to read PDF stream:", err)
	}

	if err := os.WriteFile(targetFile, pdfData, 0644); err != nil {
		log.Fatalln("Failed to save PDF file:", err)
	}

	fmt.Println("Success! PDF generated at:", targetFile)
}

func floatPtr(v float64) *float64 {
	return &v
}
