package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

/*CreateAddressParams contains all the parameters to send to the API endpoint
for the create address operation typically these are written to a http.Request
*/
type CreateAddressParams struct {

	/*Request
	  The address object to be created.

	*/
	Request *models.CreateAddressRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAddressParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Request == nil {
		o.Request = new(models.CreateAddressRequest)
	}

	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
