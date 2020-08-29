package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(memberCmd)
}

var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "Manage members",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
