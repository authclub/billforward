package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// UpsertMetadataForAccountReader is a Reader for the UpsertMetadataForAccount structure.
type UpsertMetadataForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpsertMetadataForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpsertMetadataForAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpsertMetadataForAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpsertMetadataForAccountOK creates a UpsertMetadataForAccountOK with default headers values
func NewUpsertMetadataForAccountOK() *UpsertMetadataForAccountOK {
	return &UpsertMetadataForAccountOK{}
}

/*UpsertMetadataForAccountOK handles this case with default header values.

success
*/
type UpsertMetadataForAccountOK struct {
	Payload models.DynamicMetadata
}

func (o *UpsertMetadataForAccountOK) Error() string {
	return fmt.Sprintf("[PUT /accounts/{account-ID}/metadata][%d] upsertMetadataForAccountOK  %+v", 200, o.Payload)
}

func (o *UpsertMetadataForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpsertMetadataForAccountDefault creates a UpsertMetadataForAccountDefault with default headers values
func NewUpsertMetadataForAccountDefault(code int) *UpsertMetadataForAccountDefault {
	return &UpsertMetadataForAccountDefault{
		_statusCode: code,
	}
}

/*UpsertMetadataForAccountDefault handles this case with default header values.

error
*/
type UpsertMetadataForAccountDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the upsert metadata for account default response
func (o *UpsertMetadataForAccountDefault) Code() int {
	return o._statusCode
}

func (o *UpsertMetadataForAccountDefault) Error() string {
	return fmt.Sprintf("[PUT /accounts/{account-ID}/metadata][%d] upsertMetadataForAccount default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertMetadataForAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
