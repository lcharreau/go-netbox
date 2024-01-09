/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.0 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the TunnelTerminationRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelTerminationRole{}

// TunnelTerminationRole struct for TunnelTerminationRole
type TunnelTerminationRole struct {
	Value                *PatchedWritableTunnelTerminationRequestRole `json:"value,omitempty"`
	Label                *TunnelTerminationRoleLabel                  `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelTerminationRole TunnelTerminationRole

// NewTunnelTerminationRole instantiates a new TunnelTerminationRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelTerminationRole() *TunnelTerminationRole {
	this := TunnelTerminationRole{}
	return &this
}

// NewTunnelTerminationRoleWithDefaults instantiates a new TunnelTerminationRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelTerminationRoleWithDefaults() *TunnelTerminationRole {
	this := TunnelTerminationRole{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TunnelTerminationRole) GetValue() PatchedWritableTunnelTerminationRequestRole {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritableTunnelTerminationRequestRole
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelTerminationRole) GetValueOk() (*PatchedWritableTunnelTerminationRequestRole, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TunnelTerminationRole) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritableTunnelTerminationRequestRole and assigns it to the Value field.
func (o *TunnelTerminationRole) SetValue(v PatchedWritableTunnelTerminationRequestRole) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *TunnelTerminationRole) GetLabel() TunnelTerminationRoleLabel {
	if o == nil || IsNil(o.Label) {
		var ret TunnelTerminationRoleLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelTerminationRole) GetLabelOk() (*TunnelTerminationRoleLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *TunnelTerminationRole) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given TunnelTerminationRoleLabel and assigns it to the Label field.
func (o *TunnelTerminationRole) SetLabel(v TunnelTerminationRoleLabel) {
	o.Label = &v
}

func (o TunnelTerminationRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelTerminationRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelTerminationRole) UnmarshalJSON(data []byte) (err error) {
	varTunnelTerminationRole := _TunnelTerminationRole{}

	err = json.Unmarshal(data, &varTunnelTerminationRole)

	if err != nil {
		return err
	}

	*o = TunnelTerminationRole(varTunnelTerminationRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelTerminationRole struct {
	value *TunnelTerminationRole
	isSet bool
}

func (v NullableTunnelTerminationRole) Get() *TunnelTerminationRole {
	return v.value
}

func (v *NullableTunnelTerminationRole) Set(val *TunnelTerminationRole) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelTerminationRole) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelTerminationRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelTerminationRole(val *TunnelTerminationRole) *NullableTunnelTerminationRole {
	return &NullableTunnelTerminationRole{value: val, isSet: true}
}

func (v NullableTunnelTerminationRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelTerminationRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}