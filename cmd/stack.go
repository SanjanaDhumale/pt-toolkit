package cmd

import (
	"fmt"
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
	"github.com/SanjanaDhumale/pt-toolkit/internal/engine"
	"github.com/SanjanaDhumale/pt-toolkit/internal/stack"
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

var statusStackCmd = &cobra.Command{
	Use:   "status [stack]",
	Short: "Show stack status",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		switch args[0] {

		case "monitoring":
			stack.StatusMonitoring()

		default:
			fmt.Println("Unknown stack")
		}
	},
}

var downStackCmd = &cobra.Command{
	Use:   "down [stack]",
	Short: "Stop a stack",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		switch args[0] {

		case "monitoring":

			stack.DownMonitoring()

		default:

			fmt.Println("Unknown stack")
		}
	},
}

func init() {
	stackCmd.AddCommand(listStackCmd)
	rootCmd.AddCommand(stackCmd)
	stackCmd.AddCommand(installStackCmd)
	stackCmd.AddCommand(statusStackCmd)
	stackCmd.AddCommand(downStackCmd)
}