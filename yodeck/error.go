package yodeck

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrNoToken is returned by NewClient if a user
	// passed an empty/missing token.
	ErrNoToken = errors.New("an empty token was provided")

	// ErrAuthFailure is returned by NewClient if a user
	// passed an invalid token and failed validation against the API.
	ErrAuthFailure = errors.New("failed to authenticate using the provided token")
)

type errorResponse struct {
	Error *Error `json:"error"`
}

// Error represents an error response from the API.
type Error struct {
	ErrorResponse *Response
	Code          int         `json:"code,omitempty"`
	Errors        interface{} `json:"errors,omitempty"`
	Message       string      `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s API call to %s failed %v. Code: %d, Errors: %v, Message: %s", e.ErrorResponse.Request.Method, e.ErrorResponse.Request.URL.String(), e.ErrorResponse.Response.Status, e.Code, e.Errors, e.Message)
}

// ValidationError is a map where the key is the invalid field and the value is a message describing why the field is invalid.
type ValidationError map[string]string

func (e ValidationError) Error() string {
	var messages []string

	for k, v := range e {
		m := fmt.Sprintf("%s %s", k, v)
		messages = append(messages, m)
	}

	return strings.Join(messages, ", ")
}
