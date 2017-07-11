package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAllPayments returns a collection of all payments by default 10 values are returned records are returned in natural order

{"nickname":"Get all payments","response":"getPaymentAll.html"}
*/
func (a *Client) GetAllPayments(params *GetAllPaymentsParams) (*GetAllPaymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPaymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllPayments",
		Method:             "GET",
		PathPattern:        "/payments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllPaymentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllPaymentsOK), nil

}

/*
GetPaymentByID returns a single payment specified by the payment ID parameter

{"nickname":"Retrieve an existing payment","response":"getPaymentByID.html"}
*/
func (a *Client) GetPaymentByID(params *GetPaymentByIDParams) (*GetPaymentByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentByID",
		Method:             "GET",
		PathPattern:        "/payments/{payment-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentByIDOK), nil

}

/*
GetPaymentByInvoiceID returns a collection of payments specified by the invoice ID parameter by default 10 values are returned records are returned in natural order

{"nickname":"Get for invoice","response":"getPaymentByInvoice.html"}
*/
func (a *Client) GetPaymentByInvoiceID(params *GetPaymentByInvoiceIDParams) (*GetPaymentByInvoiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentByInvoiceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentByInvoiceID",
		Method:             "GET",
		PathPattern:        "/payments/invoice/{invoice-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentByInvoiceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentByInvoiceIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
