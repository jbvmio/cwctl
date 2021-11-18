package connectwise

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Get performs GET requests, returning resources using the given endpoint and parameters.
func (C *Client) Get(ep EP, params *Parameters, args ...interface{}) ([]byte, error) {
	req, err := makeReq(C.baseURL, C.clientID, ep, params, nil, args...)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating GET request: %w", err)
	}
	req.Header.Add(`Authorization`, `Bearer `+C.Token.AccessToken)
	resp, err := C.c.Do(req)
	switch {
	case err != nil:
		return []byte{}, fmt.Errorf("error sending GET request: %w", err)
	case resp.StatusCode != 200:
		err = checkApiErr(resp.Body)
		if err == nil {
			err = fmt.Errorf("status code %d", resp.StatusCode)
		}
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Post performs POST requests, returning responses using the given endpoint and parameters.
func (C *Client) Post(ep EP, params *Parameters, body interface{}, args ...interface{}) ([]byte, error) {
	req, err := makeReq(C.baseURL, C.clientID, ep, params, body, args...)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating POST request: %w", err)
	}
	req.Header.Add(`Authorization`, `Bearer `+C.Token.AccessToken)
	resp, err := C.c.Do(req)
	switch {
	case err != nil:
		return []byte{}, fmt.Errorf("error sending POST request: %w", err)
	case resp.StatusCode != 200:
		err = checkApiErr(resp.Body)
		if err == nil {
			err = fmt.Errorf("status code %d", resp.StatusCode)
		}
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// RawRestID performs a Raw REST request using the given ID.
func (C *Client) RawRestID(ep EP, body string, params *Parameters, args ...interface{}) ([]byte, error) {
	if !epAvailable(ep) {
		return []byte{}, fmt.Errorf("endpoint id %d is not defined or unavailable", ep)
	}
	method := getMethod(ep)
	U := C.baseURL + ep.String(args...)
	if params != nil {
		U += params.Build().Encode()
	}
	var pl io.Reader = nil
	if body != "" {
		pl = strings.NewReader(body)
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

// RawRestPath performs a Raw REST request using the given path.
func (C *Client) RawRestPath(path, method, body string, params *Parameters) ([]byte, error) {
	if !strings.HasPrefix(path, `/`) {
		path = `/` + path
	}
	U := C.baseURL + path
	if params != nil {
		U += params.Build().Encode()
	}
	var pl io.Reader = nil
	if body != "" {
		pl = strings.NewReader(body)
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

func checkApiErr(body io.ReadCloser) error {
	if body == nil {
		return nil
	}
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil
	}
	tmp := struct {
		Message string
	}{}
	json.Unmarshal(b, &tmp)
	if tmp.Message == "" {
		return nil
	}
	return fmt.Errorf("%s", tmp.Message)
}
