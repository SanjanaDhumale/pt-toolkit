package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/workspace"
	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage Performance Engineering workspaces",
}

var workspaceCreateCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		return workspace.CreateWorkspace(args[0])
	},
}

var workspaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces",
	Run: func(cmd *cobra.Command, args []string) {

		workspace.ListWorkspaces()

	},
}

func init() {

	rootCmd.AddCommand(workspaceCmd)

	workspaceCmd.AddCommand(workspaceCreateCmd)

	workspaceCmd.AddCommand(workspaceListCmd)
}