package yodeck

import "fmt"

// WorkspaceService handles the communication with URL
type WorkspaceService service

// Workspace represents a workspace.
type Workspace struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Get retrieves information about a user.
func (s *WorkspaceService) Get(id string) (*Workspace, *Response, error) {
	u := fmt.Sprintf("/workspace/%s", id)
	v := new(Workspace)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new URL.
func (s *WorkspaceService) Create(workspace *Workspace) (*Workspace, *Response, error) {
	u := "/workspace/"
	v := new(Workspace)

	resp, err := s.client.newRequestDo("POST", u, nil, workspace, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Update updates an existing URL.
func (s *WorkspaceService) Update(id string, workspace *Workspace) (*Workspace, *Response, error) {
	u := fmt.Sprintf("/workspace/%s/", id)
	v := new(Workspace)

	resp, err := s.client.newRequestDo("PUT", u, nil, workspace, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete removes an existing URL.
func (s *WorkspaceService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/workspace/%s/", id)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}
