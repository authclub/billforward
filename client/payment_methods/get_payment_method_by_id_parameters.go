package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentMethodByIDParams creates a new GetPaymentMethodByIDParams object
// with the default values initialized.
func NewGetPaymentMethodByIDParams() *GetPaymentMethodByIDParams {
	var ()
	return &GetPaymentMethodByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentMethodByIDParamsWithTimeout creates a new GetPaymentMethodByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentMethodByIDParamsWithTimeout(timeout time.Duration) *GetPaymentMethodByIDParams {
	var ()
	return &GetPaymentMethodByIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentMethodByIDParamsWithContext creates a new GetPaymentMethodByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentMethodByIDParamsWithContext(ctx context.Context) *GetPaymentMethodByIDParams {
	var ()
	return &GetPaymentMethodByIDParams{

		Context: ctx,
	}
}

// NewGetPaymentMethodByIDParamsWithHTTPClient creates a new GetPaymentMethodByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentMethodByIDParamsWithHTTPClient(client *http.Client) *GetPaymentMethodByIDParams {
	var ()
	return &GetPaymentMethodByIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentMethodByIDParams contains all the parameters to send to the API endpoint
for the get payment method by ID operation typically these are written to a http.Request
*/
type GetPaymentMethodByIDParams struct {

	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*PaymentMethodID*/
	PaymentMethodID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) WithTimeout(timeout time.Duration) *GetPaymentMethodByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) WithContext(ctx context.Context) *GetPaymentMethodByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) WithHTTPClient(client *http.Client) *GetPaymentMethodByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizations adds the organizations to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) WithOrganizations(organizations []string) *GetPaymentMethodByIDParams {
	o.SetOrganizations(organizations)
	return o
}

// SetOrganizations adds the organizations to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) SetOrganizations(organizations []string) {
	o.Organizations = organizations
}

// WithPaymentMethodID adds the paymentMethodID to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) WithPaymentMethodID(paymentMethodID string) *GetPaymentMethodByIDParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the get payment method by ID params
func (o *GetPaymentMethodByIDParams) SetPaymentMethodID(paymentMethodID string) {
	o.PaymentMethodID = paymentMethodID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentMethodByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	// path param payment-method-ID
	if err := r.SetPathParam("payment-method-ID", o.PaymentMethodID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
