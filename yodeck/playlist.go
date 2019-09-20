package yodeck

import "fmt"

// PlaylistService handles the communication with URL
type PlaylistService service

// Playlist represents a playlist. brackets
type Playlist struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Workspace   int    `json:"workspace_id,omitempty"`
	// AddGaps              bool             `json:"add_gaps,omitempty"`
	// KeepWebPagesInMemory bool             `json:"keep_web_pages_in_memory,omitempty"`
	Media []*PlaylistMedia `json:"media,omitempty"`
}

// PlaylistMedia is definition of media
type PlaylistMedia struct {
	Priority int `json:"priority,omitempty"`
	Media    int `json:"media,omitempty"`
	Duration int `json:"duration,omitempty"`
}

// PlaylistResponse represents a playlist.
type PlaylistResponse struct {
	ID          int                      `json:"id,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Description string                   `json:"description,omitempty"`
	Workspace   int                      `json:"workspace_id,omitempty"`
	Media       []*PlaylistMediaResponse `json:"media,omitempty"`
}

// PlaylistMediaResponse is definition of media
type PlaylistMediaResponse struct {
	Priority int                         `json:"priority,omitempty"`
	Media    *PlaylistMediaMediaResponse `json:"media,omitempty"`
	Duration int                         `json:"duration,omitempty"`
}

// PlaylistMediaMediaResponse is definition of media
type PlaylistMediaMediaResponse struct {
	ID int `json:"id,omitempty"`
}

// Get retrieves information about a user.
func (s *PlaylistService) Get(id string) (*PlaylistResponse, *Response, error) {
	u := fmt.Sprintf("/playlist/mixed/%s/", id)
	v := new(PlaylistResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new URL.
func (s *PlaylistService) Create(playlist *Playlist) (*PlaylistResponse, *Response, error) {
	u := "/playlist/mixed/"
	v := new(PlaylistResponse)

	resp, err := s.client.newRequestDo("POST", u, nil, playlist, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Update updates an existing URL.
func (s *PlaylistService) Update(id string, playlist *Playlist) (*PlaylistResponse, *Response, error) {
	u := fmt.Sprintf("/playlist/mixed/%s/", id)
	v := new(PlaylistResponse)

	resp, err := s.client.newRequestDo("PUT", u, nil, playlist, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete removes an existing URL.
func (s *PlaylistService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/playlist/mixed/%s/", id)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}

// {"add_gaps":true,
// "haspdf":false,
// "workspace_id":0,
// "resource_uri":"/api/v1/playlist/mixed/953878/",
// "hasplaylist":false,
// "created_at":"2018-10-10T07:20:30UTC",
// "description":"Closed Description",
// "uuid":"a9e7c66aedb14ba8b7b902114cb2de22",
// "hasaudio":false,
// "name":"Closed Sign",
// "last_push_count":1,
// "hasvideo":false,
// "hasurl":false,
// "last_modified":"2018-10-10T07:20:30UTC",
// "keep_web_pages_in_memory":false,
// "hasimage":false,
// "type":"mixed",
// "id":953878,
// "haswidget":false,
// "media":[{"priority":1,"media":1328464,"duration":20},{"priority":2,"media":1012175,"duration":60}]}

// add_gaps: false
// description: "desc UI"
// keep_web_pages_in_memory: false
// media: [{priority: 1, media: 1328464, duration: 20}, {priority: 2, media: 1012175, duration: 60},…]
// name: "test PL UI"
// parent_folder_id: null
// workspace_id: "1735"

// {"add_gaps":false,"name":"test PL UI","description":"desc UI","keep_web_pages_in_memory":false,"media":[{"priority":1,"media":1328464,"duration":20},{"priority":2,"media":1012175,"duration":60},{"priority":3,"media":784399,"duration":5}],"workspace_id":"1735","parent_folder_id":null}
