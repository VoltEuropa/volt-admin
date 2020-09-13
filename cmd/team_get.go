package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/volteuropa/volt-admin/internal"
)

func init() {
	teamCmd.AddCommand(teamGetCmd)
}

var teamGetCmd = &cobra.Command{
	Use:   "get <team-id>",
	Short: "Get teams",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := http.Client{}
		req, err := internal.Request("GET", "/teams/"+args[0], nil)
		if err != nil {
			return err
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("team with id '%s' not found", args[0])
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Print(string(body))

		return nil
	},
}
