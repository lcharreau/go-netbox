/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.0 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// TunnelEncapsulationLabel the model 'TunnelEncapsulationLabel'
type TunnelEncapsulationLabel string

// List of Tunnel_encapsulation_label
const (
	TUNNELENCAPSULATIONLABEL_I_PSEC___TRANSPORT TunnelEncapsulationLabel = "IPsec - Transport"
	TUNNELENCAPSULATIONLABEL_I_PSEC___TUNNEL    TunnelEncapsulationLabel = "IPsec - Tunnel"
	TUNNELENCAPSULATIONLABEL_IP_IN_IP           TunnelEncapsulationLabel = "IP-in-IP"
	TUNNELENCAPSULATIONLABEL_GRE                TunnelEncapsulationLabel = "GRE"
)

// All allowed values of TunnelEncapsulationLabel enum
var AllowedTunnelEncapsulationLabelEnumValues = []TunnelEncapsulationLabel{
	"IPsec - Transport",
	"IPsec - Tunnel",
	"IP-in-IP",
	"GRE",
}

func (v *TunnelEncapsulationLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TunnelEncapsulationLabel(value)
	for _, existing := range AllowedTunnelEncapsulationLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TunnelEncapsulationLabel", value)
}

// NewTunnelEncapsulationLabelFromValue returns a pointer to a valid TunnelEncapsulationLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTunnelEncapsulationLabelFromValue(v string) (*TunnelEncapsulationLabel, error) {
	ev := TunnelEncapsulationLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TunnelEncapsulationLabel: valid values are %v", v, AllowedTunnelEncapsulationLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TunnelEncapsulationLabel) IsValid() bool {
	for _, existing := range AllowedTunnelEncapsulationLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Tunnel_encapsulation_label value
func (v TunnelEncapsulationLabel) Ptr() *TunnelEncapsulationLabel {
	return &v
}

type NullableTunnelEncapsulationLabel struct {
	value *TunnelEncapsulationLabel
	isSet bool
}

func (v NullableTunnelEncapsulationLabel) Get() *TunnelEncapsulationLabel {
	return v.value
}

func (v *NullableTunnelEncapsulationLabel) Set(val *TunnelEncapsulationLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelEncapsulationLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelEncapsulationLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelEncapsulationLabel(val *TunnelEncapsulationLabel) *NullableTunnelEncapsulationLabel {
	return &NullableTunnelEncapsulationLabel{value: val, isSet: true}
}

func (v NullableTunnelEncapsulationLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelEncapsulationLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}