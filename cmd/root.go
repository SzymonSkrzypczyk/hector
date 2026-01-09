package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hector",
	Short: "Static CV generator inspired by Hugo",
	Long: `Hector is a static CV generator that allows users to create
professional CVs using predefined themes and structured data files.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
