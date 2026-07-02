package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/services"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Show installed Docker images",
	Run: func(cmd *cobra.Command, args []string) {
		services.ShowImages()
	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}