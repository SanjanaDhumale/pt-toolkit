package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage Performance Engineering services",
}

var listServiceCmd = &cobra.Command{
	Use:   "list",
	Short: "List available services",
	Run: func(cmd *cobra.Command, args []string) {
		services.ListServices()
	},
}

var startServiceCmd = &cobra.Command{
	Use:   "start [service]",
	Short: "Start a service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		services.StartService(args[0])
	},
}

var statusServiceCmd = &cobra.Command{
	Use:   "status",
	Short: "Show service status",
	Run: func(cmd *cobra.Command, args []string) {
		services.ServiceStatus()
	},
}

func init() {
	serviceCmd.AddCommand(listServiceCmd)
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(startServiceCmd)
	serviceCmd.AddCommand(statusServiceCmd)
}

