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

// ExtrasImageAttachmentsListOKBody extras image attachments list o k body
// swagger:model extrasImageAttachmentsListOKBody
type ExtrasImageAttachmentsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results ExtrasImageAttachmentsListOKBodyResults `json:"results"`
}

// Validate validates this extras image attachments list o k body
func (m *ExtrasImageAttachmentsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtrasImageAttachmentsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *ExtrasImageAttachmentsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("results", "body", m.Results); err != nil {
		return err
	}

	if err := m.Results.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("results")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtrasImageAttachmentsListOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtrasImageAttachmentsListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasImageAttachmentsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
