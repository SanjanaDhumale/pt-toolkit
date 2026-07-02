package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/installer"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Performance Engineering Toolkit",
	Run: func(cmd *cobra.Command, args []string) {
		installer.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}