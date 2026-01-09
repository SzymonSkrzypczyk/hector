package cmd

import (
	"github.com/spf13/cobra"
)

var (
	selectedTheme string
	targetFile    string
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Render yaml template for a given theme",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	templateCmd.Flags().StringVarP(
		&selectedTheme,
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

	// 👇 make --theme required
	err := templateCmd.MarkFlagRequired("theme")
	if err != nil {
		return
	}
}
