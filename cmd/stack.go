package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage Performance Engineering stacks",
}

var listStackCmd = &cobra.Command{
	Use:   "list",
	Short: "List available stacks",
	Run: func(cmd *cobra.Command, args []string) {
		services.ListStacks()
	},
}

func init() {
	stackCmd.AddCommand(listStackCmd)
	rootCmd.AddCommand(stackCmd)
}