package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// NewCreateSubscriptionV2Params creates a new CreateSubscriptionV2Params object
// with the default values initialized.
func NewCreateSubscriptionV2Params() *CreateSubscriptionV2Params {
	var ()
	return &CreateSubscriptionV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubscriptionV2ParamsWithTimeout creates a new CreateSubscriptionV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSubscriptionV2ParamsWithTimeout(timeout time.Duration) *CreateSubscriptionV2Params {
	var ()
	return &CreateSubscriptionV2Params{

		timeout: timeout,
	}
}

// NewCreateSubscriptionV2ParamsWithContext creates a new CreateSubscriptionV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSubscriptionV2ParamsWithContext(ctx context.Context) *CreateSubscriptionV2Params {
	var ()
	return &CreateSubscriptionV2Params{

		Context: ctx,
	}
}

// NewCreateSubscriptionV2ParamsWithHTTPClient creates a new CreateSubscriptionV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSubscriptionV2ParamsWithHTTPClient(client *http.Client) *CreateSubscriptionV2Params {
	var ()
	return &CreateSubscriptionV2Params{
		HTTPClient: client,
	}
}

/*CreateSubscriptionV2Params contains all the parameters to send to the API endpoint
for the create subscription v2 operation typically these are written to a http.Request
*/
type CreateSubscriptionV2Params struct {

	/*Request*/
	Request *models.CreateSubscriptionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create subscription v2 params
func (o *CreateSubscriptionV2Params) WithTimeout(timeout time.Duration) *CreateSubscriptionV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subscription v2 params
func (o *CreateSubscriptionV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subscription v2 params
func (o *CreateSubscriptionV2Params) WithContext(ctx context.Context) *CreateSubscriptionV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subscription v2 params
func (o *CreateSubscriptionV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subscription v2 params
func (o *CreateSubscriptionV2Params) WithHTTPClient(client *http.Client) *CreateSubscriptionV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subscription v2 params
func (o *CreateSubscriptionV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create subscription v2 params
func (o *CreateSubscriptionV2Params) WithRequest(request *models.CreateSubscriptionRequest) *CreateSubscriptionV2Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create subscription v2 params
func (o *CreateSubscriptionV2Params) SetRequest(request *models.CreateSubscriptionRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request == nil {
		o.Request = new(models.CreateSubscriptionRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
