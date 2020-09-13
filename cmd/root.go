package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "volt-admin",
	Short:        "Help you manage the volt.team database",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute runs the CLI.
func Execute() {
	rootCmd.Execute()
}

func init() {
	// rootCmd.PersistentFlags().String("user", "", "username)
	// rootCmd.PersistentFlags().String("password", "", "password")
}
