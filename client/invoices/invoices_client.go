package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new invoices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ExecuteInvoice attempts payment for the outstanding value of an invoice

{"nickname":"Execute invoice","response":"executeInvoiceResponse.html","request":"ExecuteInvoiceRequest.html"}
*/
func (a *Client) ExecuteInvoice(params *ExecuteInvoiceParams) (*ExecuteInvoiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteInvoiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "executeInvoice",
		Method:             "POST",
		PathPattern:        "/invoices/{invoice-ID}/execute",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteInvoiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExecuteInvoiceOK), nil

}

/*
GetInvoiceAsPDF retrieves a single invoice specified by the ID parameter

{ "nickname" : "PDF Invoice","response" : "getInvoiceByID.pdf"}
*/
func (a *Client) GetInvoiceAsPDF(params *GetInvoiceAsPDFParams) (*GetInvoiceAsPDFOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceAsPDFParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInvoiceAsPDF",
		Method:             "GET",
		PathPattern:        "/invoices/{ID}.pdf",
		ProducesMediaTypes: []string{"application/pdf"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoiceAsPDFReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoiceAsPDFOK), nil

}

/*
GetInvoiceByID retrieves a single invoice specified by the invoice ID parameter

{ "nickname" : "Retrieve an existing invoice","response" : "getInvoiceByID.html"}
*/
func (a *Client) GetInvoiceByID(params *GetInvoiceByIDParams) (*GetInvoiceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInvoiceByID",
		Method:             "GET",
		PathPattern:        "/invoices/{invoice-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoiceByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoiceByIDOK), nil

}

/*
GetInvoicesByAccountID retrieves a collection of invoices specified by the account ID parameter by default 10 values are returned records are returned in natural order

{ "nickname" : "Retrieve by account","response" : "getInvoiceByAccountID.html"}
*/
func (a *Client) GetInvoicesByAccountID(params *GetInvoicesByAccountIDParams) (*GetInvoicesByAccountIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoicesByAccountIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInvoicesByAccountID",
		Method:             "GET",
		PathPattern:        "/invoices/account/{account-ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInvoicesByAccountIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoicesByAccountIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
