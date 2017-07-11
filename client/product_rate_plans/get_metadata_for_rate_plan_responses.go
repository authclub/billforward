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

// GetMetadataForRatePlanReader is a Reader for the GetMetadataForRatePlan structure.
type GetMetadataForRatePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetadataForRatePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetadataForRatePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMetadataForRatePlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetadataForRatePlanOK creates a GetMetadataForRatePlanOK with default headers values
func NewGetMetadataForRatePlanOK() *GetMetadataForRatePlanOK {
	return &GetMetadataForRatePlanOK{}
}

/*GetMetadataForRatePlanOK handles this case with default header values.

success
*/
type GetMetadataForRatePlanOK struct {
	Payload models.DynamicMetadata
}

func (o *GetMetadataForRatePlanOK) Error() string {
	return fmt.Sprintf("[GET /product-rate-plans/{product-rate-plan-ID}/metadata][%d] getMetadataForRatePlanOK  %+v", 200, o.Payload)
}

func (o *GetMetadataForRatePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetadataForRatePlanDefault creates a GetMetadataForRatePlanDefault with default headers values
func NewGetMetadataForRatePlanDefault(code int) *GetMetadataForRatePlanDefault {
	return &GetMetadataForRatePlanDefault{
		_statusCode: code,
	}
}

/*GetMetadataForRatePlanDefault handles this case with default header values.

error
*/
type GetMetadataForRatePlanDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get metadata for rate plan default response
func (o *GetMetadataForRatePlanDefault) Code() int {
	return o._statusCode
}

func (o *GetMetadataForRatePlanDefault) Error() string {
	return fmt.Sprintf("[GET /product-rate-plans/{product-rate-plan-ID}/metadata][%d] getMetadataForRatePlan default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetadataForRatePlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
