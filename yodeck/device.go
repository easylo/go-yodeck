package yodeck

import "fmt"

// DeviceService handles the communication with URL
type DeviceService service

// Device represents a device.
type Device struct {
	ID          int               `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Workspace   int               `json:"workspace_id,omitempty"`
	DefaultShow DeviceDefaultShow `json:"default_show,omitempty"`
	WifiSSID    string            `json:"wifi_ssid,omitempty"`
	WifiKey     string            `json:"wifi_key,omitempty"`
	WifiMode    string            `json:"wifi_mode,omitempty"`
}

// DeviceDefaultShow represent default show
type DeviceDefaultShow struct {
	SourceID   int    `json:"source_id"`
	SourceType string `json:"source_type"`
}

// Get retrieves information about a user.
func (s *DeviceService) Get(id string) (*Device, *Response, error) {
	u := fmt.Sprintf("/device/%s", id)
	v := new(Device)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new URL.
func (s *DeviceService) Create(device *Device) (*Device, *Response, error) {
	u := "/device/"
	v := new(Device)

	resp, err := s.client.newRequestDo("POST", u, nil, device, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Update updates an existing URL.
func (s *DeviceService) Update(id string, device *Device) (*Device, *Response, error) {
	u := fmt.Sprintf("/device/%s/", id)
	v := new(Device)

	resp, err := s.client.newRequestDo("PUT", u, nil, device, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete removes an existing URL.
func (s *DeviceService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/device/%s/", id)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
	// return nil, nil
}
