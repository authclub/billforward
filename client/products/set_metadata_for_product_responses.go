package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// SetMetadataForProductReader is a Reader for the SetMetadataForProduct structure.
type SetMetadataForProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMetadataForProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetMetadataForProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewSetMetadataForProductDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetMetadataForProductOK creates a SetMetadataForProductOK with default headers values
func NewSetMetadataForProductOK() *SetMetadataForProductOK {
	return &SetMetadataForProductOK{}
}

/*SetMetadataForProductOK handles this case with default header values.

success
*/
type SetMetadataForProductOK struct {
	Payload models.DynamicMetadata
}

func (o *SetMetadataForProductOK) Error() string {
	return fmt.Sprintf("[POST /products/{product-ID}/metadata][%d] setMetadataForProductOK  %+v", 200, o.Payload)
}

func (o *SetMetadataForProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetMetadataForProductDefault creates a SetMetadataForProductDefault with default headers values
func NewSetMetadataForProductDefault(code int) *SetMetadataForProductDefault {
	return &SetMetadataForProductDefault{
		_statusCode: code,
	}
}

/*SetMetadataForProductDefault handles this case with default header values.

error
*/
type SetMetadataForProductDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the set metadata for product default response
func (o *SetMetadataForProductDefault) Code() int {
	return o._statusCode
}

func (o *SetMetadataForProductDefault) Error() string {
	return fmt.Sprintf("[POST /products/{product-ID}/metadata][%d] setMetadataForProduct default  %+v", o._statusCode, o.Payload)
}

func (o *SetMetadataForProductDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
