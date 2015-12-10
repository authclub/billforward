package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

type UpdateProfileReader struct {
	formats strfmt.Registry
}

func (o *UpdateProfileReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result UpdateProfileOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	default:
		var result UpdateProfileDefault
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateProfile default", &result, response.Code())
	}
}

/*
success
*/
type UpdateProfileOK struct {
	Payload *models.ProfilePagedMetadata
}

func (o *UpdateProfileOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProfilePagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*
error
*/
type UpdateProfileDefault struct {
	Payload *models.BFError
}

func (o *UpdateProfileDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
