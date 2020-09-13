package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/volteuropa/volt-admin/internal"
)

func init() {
	memberCmd.AddCommand(memberGetCmd)
}

var memberGetCmd = &cobra.Command{
	Use:   "get <member-id>",
	Short: "Get members",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := http.Client{}
		req, err := internal.Request("GET", "/members/"+args[0], nil)
		if err != nil {
			return err
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("member with id '%s' not found", args[0])
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Print(string(body))

		return nil
	},
}
