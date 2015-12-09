package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new profiles API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profiles API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Returns a single profile, specified by the ID parameter.

{"nickname":"Retrieve an existing profile","response":"getProfileByID.html"}
*/
func (a *Client) GetProfile(params GetProfileParams) (*GetProfileOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getProfile",
		Params: &params,
		Reader: &GetProfileReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfileOK), nil
}

/*Update a profile

{"nickname":"Update a profile","request":"updateProfileRequest.html","response":"updateProfileResponse.html"}
*/
func (a *Client) UpdateProfile(params UpdateProfileParams) (*UpdateProfileOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "updateProfile",
		Params: &params,
		Reader: &UpdateProfileReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProfileOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
