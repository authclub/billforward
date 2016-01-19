package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*JaxbDynamicMetadata JaxbDynamicMetadata jaxb dynamic metadata

swagger:model JaxbDynamicMetadata
*/
type JaxbDynamicMetadata struct {

	/* Values values
	 */
	Values map[string]string `json:"values,omitempty"`
}

// Validate validates this jaxb dynamic metadata
func (m *JaxbDynamicMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JaxbDynamicMetadata) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	return nil
}