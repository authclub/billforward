package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// InvoicePayment An invoice-payment specifies the amount of a particular payment used to pay part or entirety of the outstanding balance of the associated invoice.
// swagger:model InvoicePayment
type InvoicePayment struct {

	// { "description" : "The value of the payment used by payment-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	ActualAmount *float64 `json:"actualAmount"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "CRM ID of the invoice.", "verbs":["POST","PUT","GET"] }
	// Required: true
	CrmID *string `json:"crmID"`

	// { "description" : "The currency of the payment specified by a three character ISO 4217 currency code.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Currency *string `json:"currency"`

	// Payment gateway associated with the payment
	// Required: true
	Gateway *string `json:"gateway"`

	// { "description" : "ID of the payment-line.", "verbs":["POST","PUT","GET"] }
	ID string `json:"id,omitempty"`

	// { "description" : "ID of the invoice associated with the payment-line.", "verbs":["POST","PUT","GET"] }
	// Required: true
	InvoiceID *string `json:"invoiceID"`

	// { "description" : "The value that this payment-line represents.", "verbs":["POST","PUT","GET"] }
	// Required: true
	NominalAmount *float64 `json:"nominalAmount"`

	// { "description" : "ID of the organization associated with the invoice-payment.", "verbs":["POST","PUT","GET"] }
	// Required: true
	OrganizationID *string `json:"organizationID"`

	// { "description" : "ID of the payment used.", "verbs":["POST","PUT","GET"] }
	// Required: true
	PaymentID *string `json:"paymentID"`

	// {"description":"The date when the associated payment was received.","verbs":["POST","PUT","GET"]}
	// Required: true
	PaymentReceived *strfmt.DateTime `json:"paymentReceived"`

	// {"description":"The date when the associated payment was refunded.","verbs":["POST","PUT","GET"]}
	// Required: true
	RefundReceived *strfmt.DateTime `json:"refundReceived"`

	// { "description" : "The refunded amount of the invoice-payment.", "verbs":["POST","PUT","GET"] }
	// Required: true
	RefundedAmount *float64 `json:"refundedAmount"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this invoice payment
func (m *InvoicePayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCrmID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNominalAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentReceived(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundReceived(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundedAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoicePayment) validateActualAmount(formats strfmt.Registry) error {

	if err := validate.Required("actualAmount", "body", m.ActualAmount); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateCrmID(formats strfmt.Registry) error {

	if err := validate.Required("crmID", "body", m.CrmID); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

var invoicePaymentTypeGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cybersource_token","card_vault","paypal_simple","locustworld","free","coupon","credit_note","stripe","braintree","balanced","paypal","billforward_test","offline","trial","stripeACH","authorizeNet","spreedly","sagePay","trustCommerce","payvision","kash"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoicePaymentTypeGatewayPropEnum = append(invoicePaymentTypeGatewayPropEnum, v)
	}
}

const (
	// InvoicePaymentGatewayCybersourceToken captures enum value "cybersource_token"
	InvoicePaymentGatewayCybersourceToken string = "cybersource_token"
	// InvoicePaymentGatewayCardVault captures enum value "card_vault"
	InvoicePaymentGatewayCardVault string = "card_vault"
	// InvoicePaymentGatewayPaypalSimple captures enum value "paypal_simple"
	InvoicePaymentGatewayPaypalSimple string = "paypal_simple"
	// InvoicePaymentGatewayLocustworld captures enum value "locustworld"
	InvoicePaymentGatewayLocustworld string = "locustworld"
	// InvoicePaymentGatewayFree captures enum value "free"
	InvoicePaymentGatewayFree string = "free"
	// InvoicePaymentGatewayCoupon captures enum value "coupon"
	InvoicePaymentGatewayCoupon string = "coupon"
	// InvoicePaymentGatewayCreditNote captures enum value "credit_note"
	InvoicePaymentGatewayCreditNote string = "credit_note"
	// InvoicePaymentGatewayStripe captures enum value "stripe"
	InvoicePaymentGatewayStripe string = "stripe"
	// InvoicePaymentGatewayBraintree captures enum value "braintree"
	InvoicePaymentGatewayBraintree string = "braintree"
	// InvoicePaymentGatewayBalanced captures enum value "balanced"
	InvoicePaymentGatewayBalanced string = "balanced"
	// InvoicePaymentGatewayPaypal captures enum value "paypal"
	InvoicePaymentGatewayPaypal string = "paypal"
	// InvoicePaymentGatewayBillforwardTest captures enum value "billforward_test"
	InvoicePaymentGatewayBillforwardTest string = "billforward_test"
	// InvoicePaymentGatewayOffline captures enum value "offline"
	InvoicePaymentGatewayOffline string = "offline"
	// InvoicePaymentGatewayTrial captures enum value "trial"
	InvoicePaymentGatewayTrial string = "trial"
	// InvoicePaymentGatewayStripeACH captures enum value "stripeACH"
	InvoicePaymentGatewayStripeACH string = "stripeACH"
	// InvoicePaymentGatewayAuthorizeNet captures enum value "authorizeNet"
	InvoicePaymentGatewayAuthorizeNet string = "authorizeNet"
	// InvoicePaymentGatewaySpreedly captures enum value "spreedly"
	InvoicePaymentGatewaySpreedly string = "spreedly"
	// InvoicePaymentGatewaySagePay captures enum value "sagePay"
	InvoicePaymentGatewaySagePay string = "sagePay"
	// InvoicePaymentGatewayTrustCommerce captures enum value "trustCommerce"
	InvoicePaymentGatewayTrustCommerce string = "trustCommerce"
	// InvoicePaymentGatewayPayvision captures enum value "payvision"
	InvoicePaymentGatewayPayvision string = "payvision"
	// InvoicePaymentGatewayKash captures enum value "kash"
	InvoicePaymentGatewayKash string = "kash"
)

// prop value enum
func (m *InvoicePayment) validateGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, invoicePaymentTypeGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoicePayment) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	// value enum
	if err := m.validateGatewayEnum("gateway", "body", *m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateInvoiceID(formats strfmt.Registry) error {

	if err := validate.Required("invoiceID", "body", m.InvoiceID); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateNominalAmount(formats strfmt.Registry) error {

	if err := validate.Required("nominalAmount", "body", m.NominalAmount); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validatePaymentID(formats strfmt.Registry) error {

	if err := validate.Required("paymentID", "body", m.PaymentID); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validatePaymentReceived(formats strfmt.Registry) error {

	if err := validate.Required("paymentReceived", "body", m.PaymentReceived); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateRefundReceived(formats strfmt.Registry) error {

	if err := validate.Required("refundReceived", "body", m.RefundReceived); err != nil {
		return err
	}

	return nil
}

func (m *InvoicePayment) validateRefundedAmount(formats strfmt.Registry) error {

	if err := validate.Required("refundedAmount", "body", m.RefundedAmount); err != nil {
		return err
	}

	return nil
}
