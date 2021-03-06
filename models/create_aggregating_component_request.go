package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CreateAggregatingComponentRequest Entity for requesting that an 'aggregating component' (i.e. a component which should be re-priced upon invoice aggregation) be created.
// swagger:model CreateAggregatingComponentRequest
type CreateAggregatingComponentRequest struct {

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// {"default":"(Auto-populated using your authentication credentials)","description":"ID of the BillForward Organization within which the requested pricing component should be created. If omitted: this will be auto-populated using your authentication credentials.","verbs":["POST"]}
	OrganizationID string `json:"organizationID,omitempty"`

	// {"description":"Name of the pricing component upon which to aggregate. The subscriber to the aggregating rate plan (which contains the AggregatingComponent specified here), will consult its children at the end of each billing period, and collect from those children all charges whose pricing component matches the ID of the component identified here. Those charges' quantities will be counted, and used when calculating the price of consuming this AggregatingComponent. The aggregating subscription then raises a discount charge &mdash; to account for the more favourable price tiering that emerges when aggregating.","verbs":["POST"]}
	// Required: true
	PricingComponent *string `json:"pricingComponent"`
}

// Validate validates this create aggregating component request
func (m *CreateAggregatingComponentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePricingComponent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAggregatingComponentRequest) validatePricingComponent(formats strfmt.Registry) error {

	if err := validate.Required("pricingComponent", "body", m.PricingComponent); err != nil {
		return err
	}

	return nil
}
