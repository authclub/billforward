package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// SetMetadataForSubscriptionReader is a Reader for the SetMetadataForSubscription structure.
type SetMetadataForSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMetadataForSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetMetadataForSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewSetMetadataForSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetMetadataForSubscriptionOK creates a SetMetadataForSubscriptionOK with default headers values
func NewSetMetadataForSubscriptionOK() *SetMetadataForSubscriptionOK {
	return &SetMetadataForSubscriptionOK{}
}

/*SetMetadataForSubscriptionOK handles this case with default header values.

success
*/
type SetMetadataForSubscriptionOK struct {
	Payload models.DynamicMetadata
}

func (o *SetMetadataForSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscription-ID}/metadata][%d] setMetadataForSubscriptionOK  %+v", 200, o.Payload)
}

func (o *SetMetadataForSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetMetadataForSubscriptionDefault creates a SetMetadataForSubscriptionDefault with default headers values
func NewSetMetadataForSubscriptionDefault(code int) *SetMetadataForSubscriptionDefault {
	return &SetMetadataForSubscriptionDefault{
		_statusCode: code,
	}
}

/*SetMetadataForSubscriptionDefault handles this case with default header values.

error
*/
type SetMetadataForSubscriptionDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the set metadata for subscription default response
func (o *SetMetadataForSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *SetMetadataForSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscription-ID}/metadata][%d] setMetadataForSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *SetMetadataForSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
