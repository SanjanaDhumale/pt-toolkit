package release

import (
	"github.com/spf13/cobra"
)

var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Manage PT Toolkit releases",
}

func init() {
	ReleaseCmd.AddCommand(BuildCmd)

}
