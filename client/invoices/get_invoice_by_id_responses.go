package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// GetInvoiceByIDReader is a Reader for the GetInvoiceByID structure.
type GetInvoiceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetInvoiceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInvoiceByIDOK creates a GetInvoiceByIDOK with default headers values
func NewGetInvoiceByIDOK() *GetInvoiceByIDOK {
	return &GetInvoiceByIDOK{}
}

/*GetInvoiceByIDOK handles this case with default header values.

success
*/
type GetInvoiceByIDOK struct {
	Payload *models.InvoicePagedMetadata
}

func (o *GetInvoiceByIDOK) Error() string {
	return fmt.Sprintf("[GET /invoices/{invoice-ID}][%d] getInvoiceByIdOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoicePagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceByIDDefault creates a GetInvoiceByIDDefault with default headers values
func NewGetInvoiceByIDDefault(code int) *GetInvoiceByIDDefault {
	return &GetInvoiceByIDDefault{
		_statusCode: code,
	}
}

/*GetInvoiceByIDDefault handles this case with default header values.

error
*/
type GetInvoiceByIDDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the get invoice by ID default response
func (o *GetInvoiceByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetInvoiceByIDDefault) Error() string {
	return fmt.Sprintf("[GET /invoices/{invoice-ID}][%d] getInvoiceByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetInvoiceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
