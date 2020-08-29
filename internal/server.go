package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

type googleToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}

// ServerURL returns the server URL.
func ServerURL() string {
	server, ok := os.LookupEnv("VOLT_ADMIN_SERVER")
	if ok {
		return server
	}
	return "http://volt.team:8090"
}

// TokenPath returns where is the token.json used.
func TokenPath() (string, error) {
	server, ok := os.LookupEnv("VOLT_ADMIN_TOKEN")
	if ok {
		return server, nil
	}
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".volt-admin-token.json"), nil
}

// Request returns a request after taking care of the server URL and token.
func Request(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, ServerURL()+path, body)
	if err != nil {
		return nil, err
	}

	tokenPath, err := TokenPath()
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read the token file '%s', are you logged in? Run 'volt-admin login' if that is not the case", tokenPath)
	}
	data := googleToken{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("AccessToken", data.AccessToken)
	req.Header.Add("RefreshToken", data.RefreshToken)
	req.Header.Add("Expiry", data.Expiry)

	return req, nil
}

// RequestWithoutToken returns a request set to the correct server URL but without a token set.
func RequestWithoutToken(method string, path string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, ServerURL()+path, body)
}
