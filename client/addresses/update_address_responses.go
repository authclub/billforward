package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// UpdateAddressReader is a Reader for the UpdateAddress structure.
type UpdateAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAddressOK creates a UpdateAddressOK with default headers values
func NewUpdateAddressOK() *UpdateAddressOK {
	return &UpdateAddressOK{}
}

/*UpdateAddressOK handles this case with default header values.

success
*/
type UpdateAddressOK struct {
	Payload *models.AddressPagedMetadata
}

func (o *UpdateAddressOK) Error() string {
	return fmt.Sprintf("[PUT /addresses][%d] updateAddressOK  %+v", 200, o.Payload)
}

func (o *UpdateAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAddressDefault creates a UpdateAddressDefault with default headers values
func NewUpdateAddressDefault(code int) *UpdateAddressDefault {
	return &UpdateAddressDefault{
		_statusCode: code,
	}
}

/*UpdateAddressDefault handles this case with default header values.

error
*/
type UpdateAddressDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the update address default response
func (o *UpdateAddressDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAddressDefault) Error() string {
	return fmt.Sprintf("[PUT /addresses][%d] updateAddress default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
