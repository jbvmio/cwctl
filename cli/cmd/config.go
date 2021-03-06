package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/jbvmio/cwctl/connectwise"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

var (
	staticBaseURL  string
	staticClientID string
)

// Config holds configuration details for API ops for siemplify.
type Config struct {
	BaseURL   string `json:"baseURL" yaml:"baseURL"`
	ClientID  string `json:"clientID" yaml:"clientID"`
	TokenFile string `json:"tokenFile" yaml:"tokenFile"`
}

// staticConfig attempts to create and return a Config if static variables are set.
// Returns nil if static variables are missing.
// staticConfig always sets the TokenFile in the user's home directory using .cwTokenFile.yaml
func staticConfig() *Config {
	switch "" {
	case staticBaseURL, staticClientID:
		return nil
	}
	return &Config{
		BaseURL:   staticBaseURL,
		ClientID:  staticClientID,
		TokenFile: homeDir() + `/.cwTokenFile.yaml`,
	}
}

// GetConfig reads from the given path and returns a Config or error.
func GetConfig(path string) (*Config, error) {
	var C Config
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return &C, err
	}
	switch {
	case strings.HasSuffix(path, `yaml`) || strings.HasSuffix(path, `yml`):
		err = yaml.Unmarshal(f, &C)
	case strings.HasSuffix(path, `json`):
		err = json.Unmarshal(f, &C)
	default:
		err = yaml.Unmarshal(f, &C)
		if err != nil {
			err = json.Unmarshal(f, &C)
			if err != nil {
				err = fmt.Errorf("invalid config (niether json|yaml) at %q", path)
			}
		}
	}
	return &C, err
}

func clientFromConfig(cfg *Config) (*connectwise.Client, error) {
	switch {
	case !fileExists(cfg.TokenFile):
		return loginCW(cfg)
	default:
		token, err := connectwise.ImportToken(cfg.TokenFile)
		if err != nil {
			return &connectwise.Client{}, fmt.Errorf("error importing CW token: %w", err)
		}
		client, err := connectwise.NewClient(cfg.BaseURL, cfg.ClientID, token)
		if err != nil {
			switch err.Error() {
			case `Token Expired`, `Token Invalid`:
				return loginCW(cfg)
			default:
				return client, fmt.Errorf("error creating CW client: %w", err)
			}
		}
		S, err := token.SecondsLeft()
		if err != nil {
			return client, fmt.Errorf("error obtaining token seconds before expiry: %w", err)
		}
		if S <= 2400 {
			err := client.RefreshToken()
			if err != nil {
				return client, fmt.Errorf("error refreshing token: %w", err)
			}
			err = client.SaveToken(cfg.TokenFile)
			if err != nil {
				return client, fmt.Errorf("error saving refreshed CW token: %w", err)
			}
		}
		return client, nil
	}
}

func loginCW(cfg *Config) (*connectwise.Client, error) {
	user, pass, code := getLoginCreds()
	auth := connectwise.Credentials{
		Username:          user,
		Password:          pass,
		TwoFactorPasscode: code,
	}
	client, err := connectwise.NewClient(cfg.BaseURL, cfg.ClientID, &auth)
	if err != nil {
		return client, fmt.Errorf("error creating CW client: %w", err)
	}
	err = client.SaveToken(cfg.TokenFile)
	if err != nil {
		return client, fmt.Errorf("error saving CW token: %w", err)
	}
	return client, nil
}

func getLoginCreds() (user, pass, code string) {
	user = readResponse("Username: ")
	pass = readSecret("Password: ")
	fmt.Println()
	code = readResponse("Passcode: ")
	return
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func readResponse(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	r, err := reader.ReadBytes(byte(10))
	if err != nil {
		Failf("error reading response: %v", err)
	}
	r = bytes.TrimSpace(r)
	return string(r)
}

func readSecret(prompt string) (secret string) {
	fmt.Fprint(os.Stderr, prompt)
	byteSecret, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		Failf("Error reading secret: %v", err)
	}
	return string(byteSecret)
}
