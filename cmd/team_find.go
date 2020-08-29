package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	teamCmd.AddCommand(teamFindCmd)
}

var teamFindCmd = &cobra.Command{
	Use:   "find",
	Short: "Find teams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find teams")
	},
}
