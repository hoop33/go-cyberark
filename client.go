package cyberark

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const basePath = "AIMWebService/api/"

// ClientOptionFunc is a function that configures a client
type ClientOptionFunc func(*Client) error

// Client is a client for the CyberArk Enterprise Password Vault
type Client struct {
	host                 string
	skipCertVerification bool
	timeout              time.Duration
}

// NewClient returns a new CyberArk client
//
// It is designed to be short-lived -- you fire your request and store the result.
//
// The caller can configure the new client by passing configuration options.
// Host is required.
//
// Example:
//
//   client, err := cyberark.NewClient(
//     cyberark.SetHost("cyaberark.example.com"),
//   )
func NewClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		host:                 "",
		skipCertVerification: false,
		timeout:              30,
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	if c.host == "" {
		return nil, errors.New("host is required")
	}

	c.host = canonicalize(c.host)

	return c, nil
}

// SetHost sets the host for the client
func SetHost(host string) ClientOptionFunc {
	return func(c *Client) error {
		if host == "" {
			return errors.New("host cannot be empty")
		}
		c.host = host
		return nil
	}
}

// SetSkipCertVerification sets whether the connection can be insecury (skip certificate verification)
func SetSkipCertVerification(skipCertVerification bool) ClientOptionFunc {
	return func(c *Client) error {
		c.skipCertVerification = skipCertVerification
		return nil
	}
}

// SetTimeout sets the client timeout
func SetTimeout(timeout time.Duration) ClientOptionFunc {
	return func(c *Client) error {
		c.timeout = timeout
		return nil
	}
}

// PerformRequest performs an HTTP request to CyberArk
func (c *Client) PerformRequest(method string, path string, params url.Values, body interface{}) (*Response, error) {

	urlStr := c.buildURL(path)
	if len(params) > 0 {
		urlStr = fmt.Sprintf("%s?%s", urlStr, params.Encode())
	}
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * c.timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.skipCertVerification,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return NewResponse(resp)
}

// GetPassword returns a new GetPasswordService
func (c *Client) GetPassword() *GetPasswordService {
	return newGetPasswordService(c)
}

func (c *Client) buildURL(path string) string {
	return fmt.Sprintf("%s%s%s", c.host, basePath, path)
}

func canonicalize(host string) string {
	if !(strings.HasPrefix(host, "https://") || strings.HasPrefix(host, "http://")) {
		host = fmt.Sprintf("https://%s", host)
	}

	if !strings.HasSuffix(host, "/") {
		host = fmt.Sprintf("%s/", host)
	}

	return host
}
