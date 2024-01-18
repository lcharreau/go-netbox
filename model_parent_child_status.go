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

// ParentChildStatus Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.  * `parent` - Parent * `child` - Child
type ParentChildStatus string

// List of Parent_child_status
const (
	PARENTCHILDSTATUS_PARENT ParentChildStatus = "parent"
	PARENTCHILDSTATUS_CHILD  ParentChildStatus = "child"
	PARENTCHILDSTATUS_EMPTY  ParentChildStatus = ""
)

// All allowed values of ParentChildStatus enum
var AllowedParentChildStatusEnumValues = []ParentChildStatus{
	"parent",
	"child",
	"",
}

func (v *ParentChildStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ParentChildStatus(value)
	for _, existing := range AllowedParentChildStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ParentChildStatus", value)
}

// NewParentChildStatusFromValue returns a pointer to a valid ParentChildStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewParentChildStatusFromValue(v string) (*ParentChildStatus, error) {
	ev := ParentChildStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ParentChildStatus: valid values are %v", v, AllowedParentChildStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ParentChildStatus) IsValid() bool {
	for _, existing := range AllowedParentChildStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Parent_child_status value
func (v ParentChildStatus) Ptr() *ParentChildStatus {
	return &v
}

type NullableParentChildStatus struct {
	value *ParentChildStatus
	isSet bool
}

func (v NullableParentChildStatus) Get() *ParentChildStatus {
	return v.value
}

func (v *NullableParentChildStatus) Set(val *ParentChildStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableParentChildStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableParentChildStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentChildStatus(val *ParentChildStatus) *NullableParentChildStatus {
	return &NullableParentChildStatus{value: val, isSet: true}
}

func (v NullableParentChildStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentChildStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
