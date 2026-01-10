package cmd

import (
	"fmt"
	"hector/internal/pipeline"
	"log"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	selectedThemePreview string
	valuesPathPreview    string
	noOpen               bool
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

	rootCmd.AddCommand(previewCmd)
}
