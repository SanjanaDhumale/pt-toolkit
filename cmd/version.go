package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display PT Toolkit version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("===================================")
		fmt.Println("PT Toolkit")
		fmt.Println("Version : v0.1.0")
		fmt.Println("Author  : Sanjana Dhumale")
		fmt.Println("===================================")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}