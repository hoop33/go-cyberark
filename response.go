package cyberark

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response is a CyberArk response
type Response struct {
	StatusCode int
	Body       json.RawMessage
}

// NewResponse creates a CyberArk response from an HTTP response
func NewResponse(resp *http.Response) (*Response, error) {
	r := &Response{
		StatusCode: resp.StatusCode,
	}

	if resp.Body != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if len(body) > 0 {
			if err := json.Unmarshal(body, &r.Body); err != nil {
				return nil, err
			}
		}
	}
	return r, nil
}
