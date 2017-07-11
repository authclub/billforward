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

// ReviveSubscriptionReader is a Reader for the ReviveSubscription structure.
type ReviveSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReviveSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReviveSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReviveSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReviveSubscriptionOK creates a ReviveSubscriptionOK with default headers values
func NewReviveSubscriptionOK() *ReviveSubscriptionOK {
	return &ReviveSubscriptionOK{}
}

/*ReviveSubscriptionOK handles this case with default header values.

success
*/
type ReviveSubscriptionOK struct {
	Payload *models.SubscriptionPagedMetadata
}

func (o *ReviveSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscription-ID}/revive][%d] reviveSubscriptionOK  %+v", 200, o.Payload)
}

func (o *ReviveSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReviveSubscriptionDefault creates a ReviveSubscriptionDefault with default headers values
func NewReviveSubscriptionDefault(code int) *ReviveSubscriptionDefault {
	return &ReviveSubscriptionDefault{
		_statusCode: code,
	}
}

/*ReviveSubscriptionDefault handles this case with default header values.

error
*/
type ReviveSubscriptionDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the revive subscription default response
func (o *ReviveSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *ReviveSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscription-ID}/revive][%d] reviveSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *ReviveSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
