package cmd

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/execution"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run performance tests",
}

var jmeterCmd = &cobra.Command{
	Use:   "jmeter [script]",
	Short: "Run a JMeter script",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		return execution.RunJMeter(args[0])

	},
}

func init() {

	rootCmd.AddCommand(testCmd)

	testCmd.AddCommand(jmeterCmd)

}