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

// WritableProvider writable provider
// swagger:model WritableProvider
type WritableProvider struct {

	// Account number
	// Max Length: 30
	Account string `json:"account,omitempty"`

	// Admin contact
	AdminContact string `json:"admin_contact,omitempty"`

	// ASN
	// Maximum: 4.294967295e+09
	// Minimum: 1
	Asn int64 `json:"asn,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	Name *string `json:"name"`

	// NOC contact
	NocContact string `json:"noc_contact,omitempty"`

	// Portal
	// Max Length: 200
	PortalURL strfmt.URI `json:"portal_url,omitempty"`

	// Slug
	// Required: true
	// Max Length: 50
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`
}

// Validate validates this writable provider
func (m *WritableProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAsn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePortalURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableProvider) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if err := validate.MaxLength("account", "body", string(m.Account), 30); err != nil {
		return err
	}

	return nil
}

func (m *WritableProvider) validateAsn(formats strfmt.Registry) error {

	if swag.IsZero(m.Asn) { // not required
		return nil
	}

	if err := validate.MinimumInt("asn", "body", int64(m.Asn), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("asn", "body", int64(m.Asn), 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableProvider) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableProvider) validatePortalURL(formats strfmt.Registry) error {

	if swag.IsZero(m.PortalURL) { // not required
		return nil
	}

	if err := validate.MaxLength("portal_url", "body", string(m.PortalURL), 200); err != nil {
		return err
	}

	if err := validate.FormatOf("portal_url", "body", "uri", m.PortalURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableProvider) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", string(*m.Slug), 50); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", string(*m.Slug), `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableProvider) UnmarshalBinary(b []byte) error {
	var res WritableProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
