package yodeck

import "fmt"

// WebpageService handles the communication with Webpage
type WebpageService service

// Webpage represents a url.
type Webpage struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Workspace   int    `json:"workspace_id,omitempty"`
}

// Get retrieves information about a user.
func (s *WebpageService) Get(id string) (*Webpage, *Response, error) {
	u := fmt.Sprintf("/url/%s", id)
	v := new(Webpage)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new Webpage.
func (s *WebpageService) Create(webpage *Webpage) (*Webpage, *Response, error) {
	u := "/url/"
	v := new(Webpage)

	resp, err := s.client.newRequestDo("POST", u, nil, webpage, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Update updates an existing Webpage.
func (s *WebpageService) Update(id string, webpage *Webpage) (*Webpage, *Response, error) {
	u := fmt.Sprintf("/url/%s/", id)
	v := new(Webpage)

	resp, err := s.client.newRequestDo("PUT", u, nil, webpage, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete removes an existing Webpage.
func (s *WebpageService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/url/%s/", id)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}
