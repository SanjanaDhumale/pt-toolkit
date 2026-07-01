package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check PT Toolkit environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running PT Toolkit health check...")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}