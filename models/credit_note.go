package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CreditNote CreditNote
// swagger:model CreditNote
type CreditNote struct {

	// { "description" : "", "verbs":["POST","GET"] }
	// Required: true
	AccountID *string `json:"accountID"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "", "verbs":["GET"] }
	CreatedBy string `json:"createdBy,omitempty"`

	// { "description" : "Currency of the credit-note specified by a three character ISO 4217 currency code.", "verbs":["POST","GET"] }
	// Required: true
	Currency *string `json:"currency"`

	// { "description" : "", "verbs":["POST","GET"] }
	Description string `json:"description,omitempty"`

	// { "description" : "", "verbs":["GET"] }
	ID string `json:"id,omitempty"`

	// { "description" : "References an invoice from this credit note. This has no side-effects, such as limited scope of credit note.", "verbs":["POST","GET"] }
	InvoiceID string `json:"invoiceID,omitempty"`

	// { "description" : "", "verbs":[] }
	OrganizationID string `json:"organizationID,omitempty"`

	// { "description" : "Remaining value of the payment not used. In the case when a credit-note is used across a range of invoices, each use reducing the available blance of the credit note.", "verbs":["GET"] }
	// Required: true
	RemainingValue *float64 `json:"remainingValue"`

	// { "description" : "Subscription to apply the credit note to. By default credit notes are owned by the account an can be used on any subscription. Providing this value limits the credit-note to only being used on the specified subscription.", "verbs":["POST","GET"] }
	SubscriptionID string `json:"subscriptionID,omitempty"`

	// { "description" : "", "verbs":[] }
	// Required: true
	Type *string `json:"type"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// { "description" : " Monetary value of the credit-note", "verbs":["POST","GET"] }
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this credit note
func (m *CreditNote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemainingValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreditNote) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *CreditNote) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *CreditNote) validateRemainingValue(formats strfmt.Registry) error {

	if err := validate.Required("remainingValue", "body", m.RemainingValue); err != nil {
		return err
	}

	return nil
}

var creditNoteTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["invoice","manual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		creditNoteTypeTypePropEnum = append(creditNoteTypeTypePropEnum, v)
	}
}

const (
	// CreditNoteTypeInvoice captures enum value "invoice"
	CreditNoteTypeInvoice string = "invoice"
	// CreditNoteTypeManual captures enum value "manual"
	CreditNoteTypeManual string = "manual"
)

// prop value enum
func (m *CreditNote) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, creditNoteTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreditNote) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CreditNote) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}
