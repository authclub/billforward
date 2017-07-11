package profiles

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

// NewGetProfileByEmailAddressParams creates a new GetProfileByEmailAddressParams object
// with the default values initialized.
func NewGetProfileByEmailAddressParams() *GetProfileByEmailAddressParams {
	var (
		includeRetiredDefault = bool(false)
		offsetDefault         = int32(0)
		orderDefault          = string("DESC")
		orderByDefault        = string("created")
		recordsDefault        = int32(10)
	)
	return &GetProfileByEmailAddressParams{
		IncludeRetired: &includeRetiredDefault,
		Offset:         &offsetDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfileByEmailAddressParamsWithTimeout creates a new GetProfileByEmailAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfileByEmailAddressParamsWithTimeout(timeout time.Duration) *GetProfileByEmailAddressParams {
	var (
		includeRetiredDefault = bool(false)
		offsetDefault         = int32(0)
		orderDefault          = string("DESC")
		orderByDefault        = string("created")
		recordsDefault        = int32(10)
	)
	return &GetProfileByEmailAddressParams{
		IncludeRetired: &includeRetiredDefault,
		Offset:         &offsetDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,

		timeout: timeout,
	}
}

// NewGetProfileByEmailAddressParamsWithContext creates a new GetProfileByEmailAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProfileByEmailAddressParamsWithContext(ctx context.Context) *GetProfileByEmailAddressParams {
	var (
		includeRetiredDefault = bool(false)
		offsetDefault         = int32(0)
		orderDefault          = string("DESC")
		orderByDefault        = string("created")
		recordsDefault        = int32(10)
	)
	return &GetProfileByEmailAddressParams{
		IncludeRetired: &includeRetiredDefault,
		Offset:         &offsetDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,

		Context: ctx,
	}
}

// NewGetProfileByEmailAddressParamsWithHTTPClient creates a new GetProfileByEmailAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProfileByEmailAddressParamsWithHTTPClient(client *http.Client) *GetProfileByEmailAddressParams {
	var (
		includeRetiredDefault = bool(false)
		offsetDefault         = int32(0)
		orderDefault          = string("DESC")
		orderByDefault        = string("created")
		recordsDefault        = int32(10)
	)
	return &GetProfileByEmailAddressParams{
		IncludeRetired: &includeRetiredDefault,
		Offset:         &offsetDefault,
		Order:          &orderDefault,
		OrderBy:        &orderByDefault,
		Records:        &recordsDefault,
		HTTPClient:     client,
	}
}

/*GetProfileByEmailAddressParams contains all the parameters to send to the API endpoint
for the get profile by email address operation typically these are written to a http.Request
*/
type GetProfileByEmailAddressParams struct {

	/*Email
	  The email address of the profile.

	*/
	Email string
	/*IncludeRetired
	  Whether retired profiles should be returned.

	*/
	IncludeRetired *bool
	/*Offset
	  The offset from the first profile to return.

	*/
	Offset *int32
	/*Order
	  Ihe direction of any ordering, either ASC or DESC.

	*/
	Order *string
	/*OrderBy
	  Specify a field used to order the result set.

	*/
	OrderBy *string
	/*Organizations
	  A list of organizations used to restrict the scope of API calls.

	*/
	Organizations []string
	/*Records
	  The maximum number of profiles to return.

	*/
	Records *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithTimeout(timeout time.Duration) *GetProfileByEmailAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithContext(ctx context.Context) *GetProfileByEmailAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithHTTPClient(client *http.Client) *GetProfileByEmailAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithEmail(email string) *GetProfileByEmailAddressParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetEmail(email string) {
	o.Email = email
}

// WithIncludeRetired adds the includeRetired to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithIncludeRetired(includeRetired *bool) *GetProfileByEmailAddressParams {
	o.SetIncludeRetired(includeRetired)
	return o
}

// SetIncludeRetired adds the includeRetired to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetIncludeRetired(includeRetired *bool) {
	o.IncludeRetired = includeRetired
}

// WithOffset adds the offset to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithOffset(offset *int32) *GetProfileByEmailAddressParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrder adds the order to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithOrder(order *string) *GetProfileByEmailAddressParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetOrder(order *string) {
	o.Order = order
}

// WithOrderBy adds the orderBy to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithOrderBy(orderBy *string) *GetProfileByEmailAddressParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizations adds the organizations to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithOrganizations(organizations []string) *GetProfileByEmailAddressParams {
	o.SetOrganizations(organizations)
	return o
}

// SetOrganizations adds the organizations to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetOrganizations(organizations []string) {
	o.Organizations = organizations
}

// WithRecords adds the records to the get profile by email address params
func (o *GetProfileByEmailAddressParams) WithRecords(records *int32) *GetProfileByEmailAddressParams {
	o.SetRecords(records)
	return o
}

// SetRecords adds the records to the get profile by email address params
func (o *GetProfileByEmailAddressParams) SetRecords(records *int32) {
	o.Records = records
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfileByEmailAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
		return err
	}

	if o.IncludeRetired != nil {

		// query param include_retired
		var qrIncludeRetired bool
		if o.IncludeRetired != nil {
			qrIncludeRetired = *o.IncludeRetired
		}
		qIncludeRetired := swag.FormatBool(qrIncludeRetired)
		if qIncludeRetired != "" {
			if err := r.SetQueryParam("include_retired", qIncludeRetired); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Order != nil {

		// query param order
		var qrOrder string
		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {
			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}

	}

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	if o.Records != nil {

		// query param records
		var qrRecords int32
		if o.Records != nil {
			qrRecords = *o.Records
		}
		qRecords := swag.FormatInt32(qrRecords)
		if qRecords != "" {
			if err := r.SetQueryParam("records", qRecords); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
