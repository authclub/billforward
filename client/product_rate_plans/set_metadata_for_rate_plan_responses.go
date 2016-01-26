package product_rate_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

type SetMetadataForRatePlanReader struct {
	formats strfmt.Registry
}

func (o *SetMetadataForRatePlanReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
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
		return nil, result
	}
}

// NewSetMetadataForRatePlanOK creates a SetMetadataForRatePlanOK with default headers values
func NewSetMetadataForRatePlanOK() *SetMetadataForRatePlanOK {
	return &SetMetadataForRatePlanOK{}
}

/*SetMetadataForRatePlanOK

success
*/
type SetMetadataForRatePlanOK struct {
	Payload models.DynamicMetadata
}

func (o *SetMetadataForRatePlanOK) Error() string {
	return fmt.Sprintf("[POST /product-rate-plans/{product-rate-plan-ID}/metadata][%d] setMetadataForRatePlanOK  %+v", 200, o.Payload)
}

func (o *SetMetadataForRatePlanOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

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

/*SetMetadataForRatePlanDefault

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

func (o *SetMetadataForRatePlanDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
