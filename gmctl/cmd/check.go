package cmd

import (
	"github.com/huerni/gmitex/gmctl/envcheck"
	"github.com/spf13/cobra"
)

// gmctl check
// gmctl check
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check gmctl Environment Installation.",
	Long:  `Check gmctl Environment Installation.`,
	Args:  cobra.ExactArgs(0),
	RunE:  gmCheck,
}

func gmCheck(_ *cobra.Command, args []string) error {
	return envcheck.Prepare(true, true, true)
}

func init() {
	rootCmd.AddCommand(checkCmd)
	// 绑定flag

}
