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

// CancelSubscriptionReader is a Reader for the CancelSubscription structure.
type CancelSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCancelSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelSubscriptionOK creates a CancelSubscriptionOK with default headers values
func NewCancelSubscriptionOK() *CancelSubscriptionOK {
	return &CancelSubscriptionOK{}
}

/*CancelSubscriptionOK handles this case with default header values.

success
*/
type CancelSubscriptionOK struct {
	Payload *models.SubscriptionCancellationPagedMetadata
}

func (o *CancelSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscription-ID}/cancel][%d] cancelSubscriptionOK  %+v", 200, o.Payload)
}

func (o *CancelSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionCancellationPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelSubscriptionDefault creates a CancelSubscriptionDefault with default headers values
func NewCancelSubscriptionDefault(code int) *CancelSubscriptionDefault {
	return &CancelSubscriptionDefault{
		_statusCode: code,
	}
}

/*CancelSubscriptionDefault handles this case with default header values.

error
*/
type CancelSubscriptionDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the cancel subscription default response
func (o *CancelSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *CancelSubscriptionDefault) Error() string {
	return fmt.Sprintf("[POST /subscriptions/{subscription-ID}/cancel][%d] cancelSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *CancelSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
