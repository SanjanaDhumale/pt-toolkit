package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage JMeter plugins",
}

var listPluginCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed plugins",
	Run: func(cmd *cobra.Command, args []string) {
		services.ListPlugins()
	},
}

func init() {
	pluginCmd.AddCommand(listPluginCmd)
	rootCmd.AddCommand(pluginCmd)
}