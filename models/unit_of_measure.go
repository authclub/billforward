package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// UnitOfMeasure Units in which a "quantity of consumed PricingComponent" can be measured
// swagger:model UnitOfMeasure
type UnitOfMeasure struct {

	// { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	ChangedBy string `json:"changedBy,omitempty"`

	// { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	Created strfmt.DateTime `json:"created,omitempty"`

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	CrmID string `json:"crmID,omitempty"`

	// { "description" : "", "verbs":["GET"] }
	// Required: true
	Deleted bool `json:"deleted"`

	// { "description" : "Unit of measurement, such as users, pounds, minutes.", "verbs":["POST","PUT","GET"] }
	// Required: true
	DisplayedAs *string `json:"displayedAs"`

	// { "description" : "", "verbs":["GET"] }
	ID string `json:"id,omitempty"`

	// { "description" : "", "verbs":["POST","PUT","GET"] }
	// Required: true
	Name *string `json:"name"`

	// { "description" : "", "verbs":[] }
	// Required: true
	OrganizationID *string `json:"organizationID"`

	// { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this unit of measure
func (m *UnitOfMeasure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayedAs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnitOfMeasure) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

func (m *UnitOfMeasure) validateDisplayedAs(formats strfmt.Registry) error {

	if err := validate.Required("displayedAs", "body", m.DisplayedAs); err != nil {
		return err
	}

	return nil
}

func (m *UnitOfMeasure) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UnitOfMeasure) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}
