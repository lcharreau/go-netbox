/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.1 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// RackWidthValue * `10` - 10 inches * `19` - 19 inches * `21` - 21 inches * `23` - 23 inches
type RackWidthValue int32

// List of Rack_width_value
const (
	RACKWIDTHVALUE__10 RackWidthValue = 10
	RACKWIDTHVALUE__19 RackWidthValue = 19
	RACKWIDTHVALUE__21 RackWidthValue = 21
	RACKWIDTHVALUE__23 RackWidthValue = 23
)

// All allowed values of RackWidthValue enum
var AllowedRackWidthValueEnumValues = []RackWidthValue{
	10,
	19,
	21,
	23,
}

func (v *RackWidthValue) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackWidthValue(value)
	for _, existing := range AllowedRackWidthValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackWidthValue", value)
}

// NewRackWidthValueFromValue returns a pointer to a valid RackWidthValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackWidthValueFromValue(v int32) (*RackWidthValue, error) {
	ev := RackWidthValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackWidthValue: valid values are %v", v, AllowedRackWidthValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackWidthValue) IsValid() bool {
	for _, existing := range AllowedRackWidthValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_width_value value
func (v RackWidthValue) Ptr() *RackWidthValue {
	return &v
}

type NullableRackWidthValue struct {
	value *RackWidthValue
	isSet bool
}

func (v NullableRackWidthValue) Get() *RackWidthValue {
	return v.value
}

func (v *NullableRackWidthValue) Set(val *RackWidthValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRackWidthValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRackWidthValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackWidthValue(val *RackWidthValue) *NullableRackWidthValue {
	return &NullableRackWidthValue{value: val, isSet: true}
}

func (v NullableRackWidthValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackWidthValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
