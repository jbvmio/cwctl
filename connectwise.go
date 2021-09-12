package cwctl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client communicates with the CW API.
type Client struct {
	c        http.Client
	clientID string
	baseURL  string
	Token    *Token
}

// NewClient returns a new Client for the CW API.
func NewClient(baseURL, clientID string, auth Auth) (*Client, error) {
	switch a := auth.(type) {
	case *Token:
		return clientFromToken(baseURL, clientID, a)
	case *Credentials:
		return clientFromLogin(baseURL, clientID, a)
	default:
		return &Client{}, fmt.Errorf("%T is not a valid AuthType", a)
	}
}

// RawRest performs a Raw REST request.
func (C *Client) RawRest(URI, method, body string) error {
	var uri, params string
	path := strings.Split(URI, `?`)
	switch len(path) {
	case 1:
		uri = URI
	case 2:
		uri = path[0]
		params = path[1]
	default:
		return fmt.Errorf("too many ?'s")
	}
	if params != "" {
		urlParams := url.Values{}
		kvs := strings.Split(params, `&`)
		for _, p := range kvs {
			kv := strings.Split(p, `=`)
			switch len(kv) {
			case 2:
				urlParams.Add(kv[0], kv[1])
			default:
				return fmt.Errorf("error in the matrix")
			}
		}
		uri = uri + `?` + urlParams.Encode()
	}
	var pl io.Reader
	switch body {
	case "":
		pl = nil
	default:
		pl = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, C.baseURL+uri, pl)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add(`Content-Type`, `application/json`)
	req.Header.Add(`ClientId`, C.clientID)
	req.Header.Add(`Authorization`, `Bearer `+C.Token.AccessToken)
	resp, err := C.c.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading request: %w", err)
	}
	fmt.Printf("%s\n", b)
	return nil
}

// RefreshToken refreshes the current token being used by the Client.
func (C *Client) RefreshToken() error {
	return C.Token.Refresh(C.baseURL, C.clientID, &C.c)
}

// SaveToken saves the current token to the given path.
func (C *Client) SaveToken(path string) error {
	return C.Token.Save(path)
}

func clientFromToken(baseURL, clientID string, token *Token) (*Client, error) {
	E, err := token.Expires()
	if err != nil {
		return &Client{}, fmt.Errorf("error obtaining token expiration: %w", err)
	}
	if time.Now().UTC().After(E) {
		return &Client{}, fmt.Errorf("Token Expired")
	}
	return &Client{
		c:        http.Client{},
		clientID: clientID,
		Token:    token,
		baseURL:  baseURL,
	}, nil
}

func clientFromLogin(baseURL, clientID string, creds *Credentials) (*Client, error) {
	c := http.Client{}
	req, err := makeReq(baseURL, clientID, EPToken, nil, creds)
	if err != nil {
		return &Client{}, fmt.Errorf("error creating request: %w", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		return &Client{}, fmt.Errorf("error on login: %w", err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return &Client{}, fmt.Errorf("401 Unauthorized")
	case http.StatusOK:
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Client{}, fmt.Errorf("error reading login response: %w", err)
	}
	var token Token
	err = json.Unmarshal(b, &token)
	if err != nil {
		return &Client{}, fmt.Errorf("error unmarshaling token: %w", err)
	}
	return &Client{
		c:        c,
		clientID: clientID,
		Token:    &token,
		baseURL:  baseURL,
	}, nil
}

func makeReq(baseURL, clientID string, ep EP, params *Parameters, v interface{}) (*http.Request, error) {
	method := getMethod(ep)
	U := baseURL + ep.String()
	if params != nil {
		U += params.Build().Encode()
	}
	pl, err := ioReaderOrNil(v)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, U, pl)
	if err != nil {
		return req, err
	}
	req.Header.Add(`Content-Type`, `application/json`)
	req.Header.Add(`ClientId`, clientID)
	return req, nil
}

func ioReaderOrNil(v interface{}) (io.Reader, error) {
	var pl io.Reader
	switch v {
	case nil:
		pl = nil
	default:
		j, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		pl = bytes.NewBuffer(j)
	}
	return pl, nil
}
