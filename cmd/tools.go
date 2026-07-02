package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Show installed performance engineering tools",
	Run: func(cmd *cobra.Command, args []string) {
		services.ShowTools()
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)
}