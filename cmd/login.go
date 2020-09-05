package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/volteuropa/volt-admin/internal"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := http.Client{}
		req, err := internal.RequestWithoutToken("GET", "/login", nil)
		if err != nil {
			return err
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		authURL := string([]byte(body))
		fmt.Printf("Go to the following link in your browser then paste the authorization code in this shell and press enter: \n%v\n", authURL)

		// Waiting for the auth code.
		var authCode string
		if _, err := fmt.Scan(&authCode); err != nil {
			return err
		}
		values := map[string]string{"code": authCode}

		jsonValue, err := json.Marshal(values)
		if err != nil {
			return err
		}

		resp, err = http.Post(internal.ServerURL()+"/login", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			return err
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
		tokenPath, err := internal.TokenPath()
		if err != nil {
			return err
		}

		f, err := os.OpenFile(tokenPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatalf("Unable to cache oauth token: %v", err)
		}
		defer f.Close()

		_, err = f.WriteString(string(body))
		if err != nil {
			log.Fatalf("Unable to save token: %v", err)
		}
		fmt.Println("You token has been saved in " + tokenPath)

		return nil
	},
}
