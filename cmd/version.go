package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show PT Toolkit version",
	Run: func(cmd *cobra.Command, args []string) {
		version.Show()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}