/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/SanjanaDhumale/pt-toolkit/cmd/release"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ptctl",
	Short: "Performance Testing Toolkit CLI",
	Long: `PT Toolkit

	A Performance Engineering Toolkit 
	for managing performance testing environments.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(release.ReleaseCmd)
}
