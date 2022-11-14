// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CableTermination cable termination
//
// swagger:model CableTermination
type CableTermination struct {

	// Cable
	// Required: true
	Cable *int64 `json:"cable"`

	// End
	// Required: true
	// Enum: [A B]
	CableEnd *string `json:"cable_end"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Termination
	// Read Only: true
	Termination interface{} `json:"termination,omitempty"`

	// Termination id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationID *int64 `json:"termination_id"`

	// Termination type
	// Required: true
	TerminationType *string `json:"termination_type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this cable termination
func (m *CableTermination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCableEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CableTermination) validateCable(formats strfmt.Registry) error {

	if err := validate.Required("cable", "body", m.Cable); err != nil {
		return err
	}

	return nil
}

var cableTerminationTypeCableEndPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableTerminationTypeCableEndPropEnum = append(cableTerminationTypeCableEndPropEnum, v)
	}
}

const (

	// CableTerminationCableEndA captures enum value "A"
	CableTerminationCableEndA string = "A"

	// CableTerminationCableEndB captures enum value "B"
	CableTerminationCableEndB string = "B"
)

// prop value enum
func (m *CableTermination) validateCableEndEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cableTerminationTypeCableEndPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CableTermination) validateCableEnd(formats strfmt.Registry) error {

	if err := validate.Required("cable_end", "body", m.CableEnd); err != nil {
		return err
	}

	// value enum
	if err := m.validateCableEndEnum("cable_end", "body", *m.CableEnd); err != nil {
		return err
	}

	return nil
}

func (m *CableTermination) validateTerminationID(formats strfmt.Registry) error {

	if err := validate.Required("termination_id", "body", m.TerminationID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_id", "body", *m.TerminationID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_id", "body", *m.TerminationID, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *CableTermination) validateTerminationType(formats strfmt.Registry) error {

	if err := validate.Required("termination_type", "body", m.TerminationType); err != nil {
		return err
	}

	return nil
}

func (m *CableTermination) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cable termination based on the context it is used
func (m *CableTermination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CableTermination) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *CableTermination) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CableTermination) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CableTermination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableTermination) UnmarshalBinary(b []byte) error {
	var res CableTermination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
