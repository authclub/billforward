package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
StripeAuthCaptureRequest stripe auth capture request

swagger:model StripeAuthCaptureRequest
*/
type StripeAuthCaptureRequest struct {
	AuthCaptureRequest

	/* {"description":"ID of the captured Card in Stripe. This can be provided as well as &mdash; or instead of &mdash; the one-use `stripeToken`, to lead BillForward to the card tokenized within the Stripe vault.","verbs":["POST"]}
	 */
	CardID string `json:"cardID,omitempty"`

	/* {"description":"Single-use token <a href=\"https://stripe.com/docs/stripe.js\">provided by Stripe's client-side card capture SDK</a>, in response to your capturing a card into the Stripe vault. This token will be used by BillForward to find the tokenized card within the Stripe vault &mdash; precursory to linking a BillForward PaymentMethod to that tokenized card.","verbs":["POST"]}

	Required: true
	*/
	StripeToken string `json:"stripeToken,omitempty"`
}

// Validate validates this stripe auth capture request
func (m *StripeAuthCaptureRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.AuthCaptureRequest.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StripeAuthCaptureRequest) validateStripeToken(formats strfmt.Registry) error {

	if err := validate.Required("stripeToken", "body", string(m.StripeToken)); err != nil {
		return err
	}

	return nil
}
