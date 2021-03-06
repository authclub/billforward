package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Receipt Receipt
// swagger:model Receipt
type Receipt struct {

	// account ID
	AccountID string `json:"accountID,omitempty"`

	// card country
	CardCountry string `json:"cardCountry,omitempty"`

	// card description
	CardDescription string `json:"cardDescription,omitempty"`

	// card last four
	CardLastFour string `json:"cardLastFour,omitempty"`

	// card province
	CardProvince string `json:"cardProvince,omitempty"`

	// card type
	CardType string `json:"cardType,omitempty"`

	// cardholder name
	CardholderName string `json:"cardholderName,omitempty"`

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "CRM ID of the subscription.", "verbs":["POST","PUT","GET"] }
	// Required: true
	CrmID *string `json:"crmID"`

	// { "description" : "Currency of the invoice specified by a three character ISO 4217 currency code.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Currency *string `json:"currency"`

	// decision
	Decision string `json:"decision,omitempty"`

	// execution attempt
	ExecutionAttempt int32 `json:"executionAttempt,omitempty"`

	// gateway reference ID
	GatewayReferenceID string `json:"gatewayReferenceID,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice ID
	InvoiceID string `json:"invoiceID,omitempty"`

	// { "description" : "The type of the invoice. A subscription invoice is raised every time a subscription recurs. An amendment is created for intra-contract changes. An Adhoc invoice is created for payment that is taken out-of-band of a subscription. Finally the invoice generated for a trial period is marked as Trial.", "verbs":["POST","PUT","GET"] }
	// Required: true
	InvoiceType *string `json:"invoiceType"`

	// {"description":"IP address associated with this payment method.","verbs":["POST","PUT","GET"]}
	IPAddress string `json:"ipAddress,omitempty"`

	// {"description":"Country of the IP address associated with this payment method.","verbs":["POST","PUT","GET"]}
	IPAddressCountry string `json:"ipAddressCountry,omitempty"`

	// organization ID
	OrganizationID string `json:"organizationID,omitempty"`

	// payment gateway
	PaymentGateway string `json:"paymentGateway,omitempty"`

	// payment ID
	PaymentID string `json:"paymentID,omitempty"`

	// payment method ID
	PaymentMethodID string `json:"paymentMethodID,omitempty"`

	// raw data
	RawData []strfmt.Base64 `json:"rawData"`

	// raw reason code
	RawReasonCode string `json:"rawReasonCode,omitempty"`

	// reason code
	ReasonCode int32 `json:"reasonCode,omitempty"`

	// refund ID
	// Required: true
	RefundID *string `json:"refundID"`

	// { "description" : "Type of transaction.", "verbs":["POST","PUT","GET"] }
	// Required: true
	Type *string `json:"type"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this receipt
func (m *Receipt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrmID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDecision(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRawData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Receipt) validateCrmID(formats strfmt.Registry) error {

	if err := validate.Required("crmID", "body", m.CrmID); err != nil {
		return err
	}

	return nil
}

func (m *Receipt) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

var receiptTypeDecisionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accept","Reject","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		receiptTypeDecisionPropEnum = append(receiptTypeDecisionPropEnum, v)
	}
}

const (
	// ReceiptDecisionAccept captures enum value "Accept"
	ReceiptDecisionAccept string = "Accept"
	// ReceiptDecisionReject captures enum value "Reject"
	ReceiptDecisionReject string = "Reject"
	// ReceiptDecisionError captures enum value "Error"
	ReceiptDecisionError string = "Error"
)

// prop value enum
func (m *Receipt) validateDecisionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, receiptTypeDecisionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Receipt) validateDecision(formats strfmt.Registry) error {

	if swag.IsZero(m.Decision) { // not required
		return nil
	}

	// value enum
	if err := m.validateDecisionEnum("decision", "body", m.Decision); err != nil {
		return err
	}

	return nil
}

var receiptTypeInvoiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Subscription","Trial","Charge","FinalArrears","Amendment","Aggregated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		receiptTypeInvoiceTypePropEnum = append(receiptTypeInvoiceTypePropEnum, v)
	}
}

const (
	// ReceiptInvoiceTypeSubscription captures enum value "Subscription"
	ReceiptInvoiceTypeSubscription string = "Subscription"
	// ReceiptInvoiceTypeTrial captures enum value "Trial"
	ReceiptInvoiceTypeTrial string = "Trial"
	// ReceiptInvoiceTypeCharge captures enum value "Charge"
	ReceiptInvoiceTypeCharge string = "Charge"
	// ReceiptInvoiceTypeFinalArrears captures enum value "FinalArrears"
	ReceiptInvoiceTypeFinalArrears string = "FinalArrears"
	// ReceiptInvoiceTypeAmendment captures enum value "Amendment"
	ReceiptInvoiceTypeAmendment string = "Amendment"
	// ReceiptInvoiceTypeAggregated captures enum value "Aggregated"
	ReceiptInvoiceTypeAggregated string = "Aggregated"
)

// prop value enum
func (m *Receipt) validateInvoiceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, receiptTypeInvoiceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Receipt) validateInvoiceType(formats strfmt.Registry) error {

	if err := validate.Required("invoiceType", "body", m.InvoiceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateInvoiceTypeEnum("invoiceType", "body", *m.InvoiceType); err != nil {
		return err
	}

	return nil
}

var receiptTypePaymentGatewayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cybersource_token","card_vault","paypal_simple","locustworld","free","coupon","credit_note","stripe","braintree","balanced","paypal","billforward_test","offline","trial","stripeACH","authorizeNet","spreedly","sagePay","trustCommerce","payvision","kash"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		receiptTypePaymentGatewayPropEnum = append(receiptTypePaymentGatewayPropEnum, v)
	}
}

const (
	// ReceiptPaymentGatewayCybersourceToken captures enum value "cybersource_token"
	ReceiptPaymentGatewayCybersourceToken string = "cybersource_token"
	// ReceiptPaymentGatewayCardVault captures enum value "card_vault"
	ReceiptPaymentGatewayCardVault string = "card_vault"
	// ReceiptPaymentGatewayPaypalSimple captures enum value "paypal_simple"
	ReceiptPaymentGatewayPaypalSimple string = "paypal_simple"
	// ReceiptPaymentGatewayLocustworld captures enum value "locustworld"
	ReceiptPaymentGatewayLocustworld string = "locustworld"
	// ReceiptPaymentGatewayFree captures enum value "free"
	ReceiptPaymentGatewayFree string = "free"
	// ReceiptPaymentGatewayCoupon captures enum value "coupon"
	ReceiptPaymentGatewayCoupon string = "coupon"
	// ReceiptPaymentGatewayCreditNote captures enum value "credit_note"
	ReceiptPaymentGatewayCreditNote string = "credit_note"
	// ReceiptPaymentGatewayStripe captures enum value "stripe"
	ReceiptPaymentGatewayStripe string = "stripe"
	// ReceiptPaymentGatewayBraintree captures enum value "braintree"
	ReceiptPaymentGatewayBraintree string = "braintree"
	// ReceiptPaymentGatewayBalanced captures enum value "balanced"
	ReceiptPaymentGatewayBalanced string = "balanced"
	// ReceiptPaymentGatewayPaypal captures enum value "paypal"
	ReceiptPaymentGatewayPaypal string = "paypal"
	// ReceiptPaymentGatewayBillforwardTest captures enum value "billforward_test"
	ReceiptPaymentGatewayBillforwardTest string = "billforward_test"
	// ReceiptPaymentGatewayOffline captures enum value "offline"
	ReceiptPaymentGatewayOffline string = "offline"
	// ReceiptPaymentGatewayTrial captures enum value "trial"
	ReceiptPaymentGatewayTrial string = "trial"
	// ReceiptPaymentGatewayStripeACH captures enum value "stripeACH"
	ReceiptPaymentGatewayStripeACH string = "stripeACH"
	// ReceiptPaymentGatewayAuthorizeNet captures enum value "authorizeNet"
	ReceiptPaymentGatewayAuthorizeNet string = "authorizeNet"
	// ReceiptPaymentGatewaySpreedly captures enum value "spreedly"
	ReceiptPaymentGatewaySpreedly string = "spreedly"
	// ReceiptPaymentGatewaySagePay captures enum value "sagePay"
	ReceiptPaymentGatewaySagePay string = "sagePay"
	// ReceiptPaymentGatewayTrustCommerce captures enum value "trustCommerce"
	ReceiptPaymentGatewayTrustCommerce string = "trustCommerce"
	// ReceiptPaymentGatewayPayvision captures enum value "payvision"
	ReceiptPaymentGatewayPayvision string = "payvision"
	// ReceiptPaymentGatewayKash captures enum value "kash"
	ReceiptPaymentGatewayKash string = "kash"
)

// prop value enum
func (m *Receipt) validatePaymentGatewayEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, receiptTypePaymentGatewayPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Receipt) validatePaymentGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentGateway) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentGatewayEnum("paymentGateway", "body", m.PaymentGateway); err != nil {
		return err
	}

	return nil
}

func (m *Receipt) validateRawData(formats strfmt.Registry) error {

	if swag.IsZero(m.RawData) { // not required
		return nil
	}

	return nil
}

func (m *Receipt) validateRefundID(formats strfmt.Registry) error {

	if err := validate.Required("refundID", "body", m.RefundID); err != nil {
		return err
	}

	return nil
}

var receiptTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["credit","debit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		receiptTypeTypePropEnum = append(receiptTypeTypePropEnum, v)
	}
}

const (
	// ReceiptTypeCredit captures enum value "credit"
	ReceiptTypeCredit string = "credit"
	// ReceiptTypeDebit captures enum value "debit"
	ReceiptTypeDebit string = "debit"
)

// prop value enum
func (m *Receipt) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, receiptTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Receipt) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}
