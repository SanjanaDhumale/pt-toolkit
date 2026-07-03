package cmd

import (
	"fmt"
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
	"github.com/SanjanaDhumale/pt-toolkit/internal/engine"
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

		if err := engine.Install(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	stackCmd.AddCommand(listStackCmd)
	rootCmd.AddCommand(stackCmd)
	stackCmd.AddCommand(installStackCmd)
}