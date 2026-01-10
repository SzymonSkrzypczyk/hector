package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

func ListDirectories(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, nil
}

var themesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Get a list of available themes",
	Run: func(cmd *cobra.Command, args []string) {
		directory := "themes/"
		themes, err := ListDirectories(directory)
		if err != nil {
			cmd.PrintErrln("Error:", err)
		}
		cmd.Println("Available themes:")
		for _, theme := range themes {
			cmd.Printf("\t- %s\n", theme)
		}
	},
}

func init() {
	rootCmd.AddCommand(themesCmd)
}
