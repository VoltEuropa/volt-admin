package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	memberCmd.AddCommand(memberDeleteCmd)
}

var memberDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete members",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete members")
	},
}
