package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ReviveSubscriptionRequest ReviveSubscriptionRequest
// swagger:model ReviveSubscriptionRequest
type ReviveSubscriptionRequest struct {

	// subscription ID
	// Required: true
	SubscriptionID *string `json:"subscriptionID"`
}

// Validate validates this revive subscription request
func (m *ReviveSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReviveSubscriptionRequest) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionID", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}
