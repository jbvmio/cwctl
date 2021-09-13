package connectwise

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

// RawRestID performs a Raw REST request using the given ID.
func (C *Client) RawRestID(ep EP, body string, params *Parameters) ([]byte, error) {
	if !epAvailable(ep) {
		return []byte{}, fmt.Errorf("endpoint id %d is not defined or unavailable", ep)
	}
	req, err := makeReq(C.baseURL, C.clientID, ep, params, body)
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

// RawRestPath performs a Raw REST request using the given path.
func (C *Client) RawRestPath(path, method, body string, params *Parameters) ([]byte, error) {
	if !strings.HasPrefix(path, `/`) {
		path = `/` + path
	}
	U := C.baseURL + path
	if params != nil {
		U += params.Build().Encode()
	}
	pl, err := ioReaderOrNil(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, U, pl)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add(`Content-Type`, `application/json`)
	req.Header.Add(`ClientId`, C.clientID)
	req.Header.Add(`Authorization`, `Bearer `+C.Token.AccessToken)
	resp, err := C.c.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
