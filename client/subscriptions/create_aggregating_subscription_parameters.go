package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// NewCreateAggregatingSubscriptionParams creates a new CreateAggregatingSubscriptionParams object
// with the default values initialized.
func NewCreateAggregatingSubscriptionParams() *CreateAggregatingSubscriptionParams {
	var ()
	return &CreateAggregatingSubscriptionParams{}
}

/*CreateAggregatingSubscriptionParams contains all the parameters to send to the API endpoint
for the create aggregating subscription operation typically these are written to a http.Request
*/
type CreateAggregatingSubscriptionParams struct {

	/*Request*/
	Request *models.CreateAggregatingSubscriptionRequest
}

// WithRequest adds the request to the create aggregating subscription params
func (o *CreateAggregatingSubscriptionParams) WithRequest(request *models.CreateAggregatingSubscriptionRequest) *CreateAggregatingSubscriptionParams {
	o.Request = request
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAggregatingSubscriptionParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.CreateAggregatingSubscriptionRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
