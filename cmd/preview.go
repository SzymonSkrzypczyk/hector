package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"hector/internal/pipeline"

	"github.com/spf13/cobra"
)

var (
	selectedThemePreview string
	valuesPathPreview    string
	noOpen               bool
	watchPreview         bool
)

func openBrowser(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", path)
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}

func checkDataChanged(originalChanges, comparedChanges string) bool {
	return originalChanges != comparedChanges
}

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview your CV in the browser",
	Run: func(cmd *cobra.Command, args []string) {
		if valuesPathPreview == "" {
			log.Fatalln("Error: --data flag is required")
		}
		htmlPath, _ := pipeline.GenerateHTML(selectedThemePreview, valuesPathPreview)
		log.Printf("Preview generated at %s\n", htmlPath)
		if noOpen {
			return
		}
		err := openBrowser(htmlPath)
		if err != nil {
			log.Printf("Failed to open browser: %v\n", err)
		}

		// variables for comparing file changes
		dir := filepath.Dir(valuesPathPreview)
		base := filepath.Base(valuesPathPreview)

		root, err := os.OpenRoot(dir)
		if err != nil {
			log.Fatalf("Error opening directory %s: %v\n", dir, err)
		}
		defer func() {
			if err := root.Close(); err != nil {
				log.Printf("Error closing directory: %v\n", err)
			}
		}()

		f, err := root.Open(base)
		if err != nil {
			log.Fatalf("Error opening file %s: %v\n", base, err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Printf("Error closing file: %v\n", err)
			}
		}()

		previousData, err := io.ReadAll(f)
		if err != nil {
			log.Fatalf("Error reading data file: %v\n", err)
		}

		if watchPreview {
			for {
				func() {
					dir := filepath.Dir(valuesPathPreview)
					base := filepath.Base(valuesPathPreview)

					root, err := os.OpenRoot(dir)
					if err != nil {
						log.Fatalf("Error opening directory %s: %v\n", dir, err)
					}
					defer func() {
						if err := root.Close(); err != nil {
							log.Printf("Error closing directory: %v\n", err)
						}
					}()

					f, err := root.Open(base)
					if err != nil {
						log.Fatalf("Error opening file %s: %v\n", base, err)
					}
					defer func() {
						if err := f.Close(); err != nil {
							log.Printf("Error closing file: %v\n", err)
						}
					}()

					currentData, err := io.ReadAll(f)
					if err != nil {
						log.Fatalf("Error reading data file: %v\n", err)
					}

					if checkDataChanged(string(previousData), string(currentData)) {
						log.Println("Changes detected, regenerating preview...")
						_, _ = pipeline.GenerateHTML(selectedThemePreview, valuesPathPreview)
						previousData = currentData
					}
				}()
			}
		}
	},
}

func init() {
	previewCmd.Flags().StringVarP(
		&selectedThemePreview,
		"theme",
		"t",
		"normal",
		"Theme to use",
	)

	previewCmd.Flags().StringVarP(
		&valuesPathPreview,
		"data",
		"d",
		"",
		"Path to CV data file (YAML or JSON)",
	)

	previewCmd.Flags().BoolVar(
		&noOpen,
		"no-open",
		false,
		"Generate preview without opening browser",
	)

	previewCmd.Flags().BoolVar(
		&watchPreview,
		"watch",
		false,
		"Watch for changes and regenerate preview automatically",
	)

	rootCmd.AddCommand(previewCmd)
}
