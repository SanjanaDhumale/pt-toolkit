package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check PT Toolkit environment",
	Run: func(cmd *cobra.Command, args []string) {
		services.RunDoctor()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}