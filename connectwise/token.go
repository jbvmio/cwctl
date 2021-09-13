package connectwise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	dateLayout = `2006-01-02T15:04:05`
)

// Token is the token payload given by the CW API.
type Token struct {
	AccessToken                 string
	TokenType                   string
	ExpirationDate              string
	AbsoluteExpirationDate      string
	UserId                      string
	IsTwoFactorRequired         bool
	IsInternalTwoFactorRequired bool
}

// ImportToken loads a saved token from the given path.
func ImportToken(path string) (*Token, error) {
	var T Token
	t, err := ioutil.ReadFile(path)
	if err != nil {
		return &T, err
	}
	err = json.Unmarshal(t, &T)
	return &T, err
}

// Refresh refreshed the token.
func (t *Token) Refresh(baseURL, clientID string, c *http.Client) error {
	req, err := makeReq(baseURL, clientID, EPTokenRefresh, nil, t.AccessToken)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return fmt.Errorf("401 Unauthorized")
	case http.StatusOK:
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var token Token
	err = json.Unmarshal(b, &token)
	if err != nil {
		return err
	}
	t = &token
	return nil
}

// Expires returns the ExpirationDate for the Token.
func (t *Token) Expires() (time.Time, error) {
	return time.Parse(dateLayout, t.ExpirationDate)
}

// TimeLeft returns the duration remaining before the Token expires.
func (t *Token) TimeLeft() (time.Duration, error) {
	T, err := t.Expires()
	if err != nil {
		return 0, err
	}
	return T.UTC().Sub(time.Now().UTC()), nil
}

// SecondsLeft returns the seconds remaining before the Token expires.
func (t *Token) SecondsLeft() (float64, error) {
	T, err := t.Expires()
	if err != nil {
		return 0, err
	}
	return T.UTC().Sub(time.Now().UTC()).Seconds(), nil
}

// Save locally saves the Token to the given path.
func (t *Token) Save(path string) error {
	j, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, j, 0777)
}
