package release

import (
	"fmt"

	internalrelease "github.com/SanjanaDhumale/pt-toolkit/internal/release"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build PT Toolkit release package",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("===================================")
		fmt.Println("      PT Toolkit Release Build")
		fmt.Println("===================================")

		return internalrelease.BuildRelease()
	},
}
