package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	teamCmd.AddCommand(teamDeleteCmd)
}

var teamDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete teams",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete teams")
	},
}
