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

var installStackCmd = &cobra.Command{
	Use:   "install [stack]",
	Short: "Install a Performance Engineering stack",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		services.InstallStack(args[0])
	},
}

func init() {
	stackCmd.AddCommand(listStackCmd)
	rootCmd.AddCommand(stackCmd)
	stackCmd.AddCommand(installStackCmd)
}