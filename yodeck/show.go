package yodeck

import (
	"fmt"
)

// ShowService handles the communication with URL
type ShowService service

// Show represents a show. brackets
type Show struct {
	ID        int            `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	Workspace int            `json:"workspace_id,omitempty"`
	Regions   []*ShowRegions `json:"regions,omitempty"`
}

// ShowRegions is definition of media
type ShowRegions struct {
	Left               int                     `json:"left,int"`
	Top                int                     `json:"top,int"`
	Width              int                     `json:"width,int"`
	Height             int                     `json:"height,int"`
	Playlists          []*ShowRegionsPlaylists `json:"playlists"`
	Fit                string                  `json:"fit,omitempty"`
	Zindex             int                     `json:"zindex,int"`
	EnableTransparency bool                    `json:"enable_transparency,bool"`
	IsMuted            bool                    `json:"is_muted,bool"`
	ResWidth           int                     `json:"res_width,int"`
	ResHeight          int                     `json:"res_height,int"`
	BackgroundAudio    bool                    `json:"background_audio,bool"`
}

// ShowRegionsPlaylists is definition of Playlist in ShowRegions
type ShowRegionsPlaylists struct {
	Playlist string `json:"playlist_id"`
	Duration int    `json:"duration,int"`
	Order    int    `json:"order,int"`
}

// // ShowResponse represents a show. brackets
// type ShowResponse struct {
// 	ID        int                    `json:"id,omitempty"`
// 	Name      string                 `json:"name,omitempty"`
// 	Workspace int                    `json:"workspace_id,omitempty"`
// 	Regions   []*ShowRegionsResponse `json:"regions,omitempty"`
// }

// // ShowRegionsResponse represents a playlist.
// type ShowRegionsResponse struct {
// 	Left               int                             `json:"left,int"`
// 	Top                int                             `json:"top,int"`
// 	Width              int                             `json:"width,int"`
// 	Height             int                             `json:"height,int"`
// 	Playlists          []*ShowRegionsPlaylistsResponse `json:"playlists"`
// 	Fit                string                          `json:"fit,omitempty"`
// 	Zindex             int                             `json:"zindex,int"`
// 	EnableTransparency bool                            `json:"enable_transparency,bool"`
// 	IsMuted            bool                            `json:"is_muted,bool"`
// 	ResWidth           int                             `json:"res_width,int"`
// 	ResHeight          int                             `json:"res_height,int"`
// 	BackgroundAudio    bool                            `json:"background_audio,bool"`
// }

// // ShowRegionsPlaylistsResponse is definition of media
// type ShowRegionsPlaylistsResponse struct {
// 	Playlist *ShowRegionsPlaylistsPlaylistsResponse `json:"playlist"`
// 	Duration int                                    `json:"duration,int"`
// 	Order    int                                    `json:"order,int"`
// }

// // ShowRegionsPlaylistsPlaylistsResponse is definition of media
// type ShowRegionsPlaylistsPlaylistsResponse struct {
// 	ID string `json:"id,omitempty"`
// }

// // ShowResponse represents a show.
// type ShowResponse struct {
// 	ID          int                    `json:"id,omitempty"`
// 	Name        string                 `json:"name,omitempty"`
// 	Workspace   int                    `json:"workspace_id,omitempty"`
// 	Regions     []*ShowRegionsResponse `json:"regions,omitempty"`
// }

// // ShowRegionsResponse is definition of media
// type ShowRegionsResponse struct {
// 	Priority int                       `json:"priority,omitempty"`
// 	Regions  *ShowRegionsMediaResponse `json:"regions,omitempty"`
// 	Duration int                       `json:"duration,omitempty"`
// }

// // ShowRegionsMediaResponse is definition of media
// type ShowRegionsMediaResponse struct {
// 	ID int `json:"id,omitempty"`
// }

// Get retrieves information about a user.
func (s *ShowService) Get(id string) (*Show, *Response, error) {
	u := fmt.Sprintf("/show/show/%s/", id)
	v := new(Show)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new URL.
func (s *ShowService) Create(show *Show) (*Show, *Response, error) {
	u := "/show/show/"
	v := new(Show)

	resp, err := s.client.newRequestDo("POST", u, nil, show, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Update updates an existing URL.
func (s *ShowService) Update(id string, show *Show) (*Show, *Response, error) {
	u := fmt.Sprintf("/show/show/%s/", id)
	v := new(Show)

	resp, err := s.client.newRequestDo("PUT", u, nil, show, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete removes an existing URL.
func (s *ShowService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/show/show/%s/", id)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}
