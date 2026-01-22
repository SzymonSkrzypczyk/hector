package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// for now predefined and hardcoded 
const (
	version = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hector",
	Long:  `All software has versions. This is Hector's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hector v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
