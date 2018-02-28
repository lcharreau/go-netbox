// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsolePort console port
// swagger:model ConsolePort
type ConsolePort struct {

	// Connection status
	ConnectionStatus bool `json:"connection_status,omitempty"`

	// cs port
	// Required: true
	CsPort *ConsoleServerPort `json:"cs_port"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	Name *string `json:"name"`
}

// Validate validates this console port
func (m *ConsolePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCsPort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsolePort) validateCsPort(formats strfmt.Registry) error {

	if err := validate.Required("cs_port", "body", m.CsPort); err != nil {
		return err
	}

	if m.CsPort != nil {

		if err := m.CsPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cs_port")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {

		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePort) UnmarshalBinary(b []byte) error {
	var res ConsolePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
