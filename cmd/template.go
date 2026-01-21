package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	selectedThemeTemplate string
	targetFile            string
)

func renderTemplate(theme string, outputFile string) error {
	if theme == "" {
		return fmt.Errorf("theme must be specified")
	}

	// templatePath is filepath.Join("themes", theme, "template.yaml")
	// Scope access to "themes" directory
	themesRoot, err := os.OpenRoot("themes")
	if err != nil {
		return fmt.Errorf("failed to open themes directory: %w", err)
	}
	defer themesRoot.Close()
	
	src, err := themesRoot.Open(filepath.Join(theme, "template.yaml"))
	if err != nil {
		return fmt.Errorf("failed to open template %q: %w", filepath.Join("themes", theme, "template.yaml"), err)
	}
	defer src.Close()

	var dst io.Writer
	if outputFile == "" {
		dst = os.Stdout
	} else {
		dir := filepath.Dir(outputFile)
		base := filepath.Base(outputFile)
		
		outRoot, err := os.OpenRoot(dir)
		if err != nil {
			return fmt.Errorf("failed to open output directory %q: %w", dir, err)
		}
		defer outRoot.Close()
		
		file, err := outRoot.Create(base)
		if err != nil {
			return fmt.Errorf("failed to create output file %q: %w", outputFile, err)
		}
		defer file.Close()
		dst = file
	}

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to copy template: %w", err)
	}

	return nil
}

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Render yaml template for a given theme",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := renderTemplate(selectedThemeTemplate, targetFile)
		if err != nil {
			log.Fatalf("failed to render template: %v", err)
		}
		return nil
	},
}

func init() {
	templateCmd.Flags().StringVarP(
		&selectedThemeTemplate,
		"theme",
		"t",
		"normal",
		"Theme to use",
	)

	templateCmd.Flags().StringVarP(
		&targetFile,
		"file",
		"f",
		"",
		"File to output the template to",
	)

	rootCmd.AddCommand(templateCmd)
}
