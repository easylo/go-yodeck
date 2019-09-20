package yodeck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://app.yodeck.com/api/v1"
)

type service struct {
	client *Client
}

// // Auth wraps the authorisation headers required for each request
// type Auth struct {
// 	Username string
// 	Apikey   string
// }

// func (a *Auth) validate() error {
// 	e := make(ValidationError)

// 	if a.Username == "" {
// 		e["Username"] = "is required"
// 	}

// 	if a.Apikey == "" {
// 		e["Apikey"] = "is required"
// 	}

// 	if len(e) > 0 {
// 		return e
// 	}

// 	return nil
// }

// Config represents the configuration for a yodeck client
type Config struct {
	BaseURL    string
	HTTPClient *http.Client
	// Token      string
	Username  string
	Apikey    string
	UserAgent string
	Debug     bool
}

// Client manages the communication with the yodeck API
type Client struct {
	baseURL   *url.URL
	client    *http.Client
	Config    *Config
	Webpage   *WebpageService
	Workspace *WorkspaceService
	Playlist  *PlaylistService
	Show      *ShowService
	Device    *DeviceService
}

// Response is a wrapper around http.Response
type Response struct {
	*http.Response
}

// // New returns a new Client
// func New(auth Auth) (*Client, error) {
// 	if err := auth.validate(); err != nil {
// 		return nil, err
// 	}

// 	return &Client{
// 		c:        &http.Client{},
// 		username: auth.Username,
// 		apiKey:   auth.Apikey,
// 	}, nil
// }

// NewClient returns a new yodeck API client.
func NewClient(config *Config) (*Client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}

	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	config.UserAgent = "go-yodeck(terraform)"

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL: baseURL,
		client:  config.HTTPClient,
		Config:  config,
	}

	c.Webpage = &WebpageService{c}
	c.Workspace = &WorkspaceService{c}
	c.Playlist = &PlaylistService{c}
	c.Show = &ShowService{c}
	c.Device = &DeviceService{c}

	return c, nil
}

func (c *Client) newRequest(method, url string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {

		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}

	}

	if c.Config.Debug {
		log.Printf("[DEBUG] yodeck - Preparing %s request to %s with body: %s", method, url, buf)
	}

	u := c.baseURL.String() + url

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s:%s", c.Config.Username, c.Config.Apikey))
	req.Header.Add("Content-Type", "application/json")

	if c.Config.UserAgent != "" {
		req.Header.Add("User-Agent", c.Config.UserAgent)
	}

	return req, nil
}

func (c *Client) newRequestDo(method, url string, options, body, v interface{}) (*Response, error) {
	if options != nil {
		values, err := query.Values(options)
		if err != nil {
			return nil, err
		}

		if v := values.Encode(); v != "" {
			url = fmt.Sprintf("%s?%s", url, v)
		}
	}

	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return c.do(req, v)
}

func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{resp}

	if err := checkResponse(response); err != nil {
		return response, err
	}

	if v != nil {
		if err := decodeJSON(response, v); err != nil {
			return response, err
		}
	}

	return response, nil
}

// ValidateAuth validates a token against the yodeck API
// func (c *Client) ValidateAuth() error {
// 	_, _, err := c.Abilities.List()
// 	return err
// }

func decodeJSON(res *Response, v interface{}) error {
	return json.NewDecoder(res.Body).Decode(v)
}

func checkResponse(res *Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	return decodeErrorResponse(res)
}

func decodeErrorResponse(res *Response) error {
	// Try to decode error response or fallback with standard error
	v := &errorResponse{Error: &Error{ErrorResponse: res}}
	if err := decodeJSON(res, v); err != nil {
		return fmt.Errorf("%s API call to %s failed: %v", res.Request.Method, res.Request.URL.String(), res.Status)
	}

	return v.Error
}
