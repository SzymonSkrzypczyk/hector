package cmd

import (
	"os"

	"github.com/spf13/cobra"
	hlog "hector/internal/log"
)

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "hector",
	Short: "Static CV generator inspired by Hugo",
	Long: `Hector is a static CV generator that allows users to create
professional CVs using predefined themes and structured data files.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		hlog.SetVerbose(verbose)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
}
