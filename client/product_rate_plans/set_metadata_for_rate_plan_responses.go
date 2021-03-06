package product_rate_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// SetMetadataForRatePlanReader is a Reader for the SetMetadataForRatePlan structure.
type SetMetadataForRatePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMetadataForRatePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetMetadataForRatePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewSetMetadataForRatePlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetMetadataForRatePlanOK creates a SetMetadataForRatePlanOK with default headers values
func NewSetMetadataForRatePlanOK() *SetMetadataForRatePlanOK {
	return &SetMetadataForRatePlanOK{}
}

/*SetMetadataForRatePlanOK handles this case with default header values.

success
*/
type SetMetadataForRatePlanOK struct {
	Payload models.DynamicMetadata
}

func (o *SetMetadataForRatePlanOK) Error() string {
	return fmt.Sprintf("[POST /product-rate-plans/{product-rate-plan-ID}/metadata][%d] setMetadataForRatePlanOK  %+v", 200, o.Payload)
}

func (o *SetMetadataForRatePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetMetadataForRatePlanDefault creates a SetMetadataForRatePlanDefault with default headers values
func NewSetMetadataForRatePlanDefault(code int) *SetMetadataForRatePlanDefault {
	return &SetMetadataForRatePlanDefault{
		_statusCode: code,
	}
}

/*SetMetadataForRatePlanDefault handles this case with default header values.

error
*/
type SetMetadataForRatePlanDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the set metadata for rate plan default response
func (o *SetMetadataForRatePlanDefault) Code() int {
	return o._statusCode
}

func (o *SetMetadataForRatePlanDefault) Error() string {
	return fmt.Sprintf("[POST /product-rate-plans/{product-rate-plan-ID}/metadata][%d] setMetadataForRatePlan default  %+v", o._statusCode, o.Payload)
}

func (o *SetMetadataForRatePlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
