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

// GetMetadataForSubscriptionReader is a Reader for the GetMetadataForSubscription structure.
type GetMetadataForSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetadataForSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetadataForSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMetadataForSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetadataForSubscriptionOK creates a GetMetadataForSubscriptionOK with default headers values
func NewGetMetadataForSubscriptionOK() *GetMetadataForSubscriptionOK {
	return &GetMetadataForSubscriptionOK{}
}

/*GetMetadataForSubscriptionOK handles this case with default header values.

success
*/
type GetMetadataForSubscriptionOK struct {
	Payload models.DynamicMetadata
}

func (o *GetMetadataForSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscription-ID}/metadata][%d] getMetadataForSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetMetadataForSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetadataForSubscriptionDefault creates a GetMetadataForSubscriptionDefault with default headers values
func NewGetMetadataForSubscriptionDefault(code int) *GetMetadataForSubscriptionDefault {
	return &GetMetadataForSubscriptionDefault{
		_statusCode: code,
	}
}

/*GetMetadataForSubscriptionDefault handles this case with default header values.

error
*/
type GetMetadataForSubscriptionDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get metadata for subscription default response
func (o *GetMetadataForSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetMetadataForSubscriptionDefault) Error() string {
	return fmt.Sprintf("[GET /subscriptions/{subscription-ID}/metadata][%d] getMetadataForSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetadataForSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
