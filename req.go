package cwctl

import (
	"fmt"
	"io/ioutil"
)

// GetClients returns Clients.
func (C *Client) GetClients(params *Parameters) ([]byte, error) {
	req, err := makeReq(C.baseURL, C.clientID, EPClients, params, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add(`Authorization`, `Bearer `+C.Token.AccessToken)
	resp, err := C.c.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
