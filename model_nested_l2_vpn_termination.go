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

// checks if the NestedL2VPNTermination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedL2VPNTermination{}

// NestedL2VPNTermination Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedL2VPNTermination struct {
	Id                   int32       `json:"id"`
	Url                  string      `json:"url"`
	Display              string      `json:"display"`
	L2vpn                NestedL2VPN `json:"l2vpn"`
	AdditionalProperties map[string]interface{}
}

type _NestedL2VPNTermination NestedL2VPNTermination

// NewNestedL2VPNTermination instantiates a new NestedL2VPNTermination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedL2VPNTermination(id int32, url string, display string, l2vpn NestedL2VPN) *NestedL2VPNTermination {
	this := NestedL2VPNTermination{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.L2vpn = l2vpn
	return &this
}

// NewNestedL2VPNTerminationWithDefaults instantiates a new NestedL2VPNTermination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedL2VPNTerminationWithDefaults() *NestedL2VPNTermination {
	this := NestedL2VPNTermination{}
	return &this
}

// GetId returns the Id field value
func (o *NestedL2VPNTermination) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPNTermination) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedL2VPNTermination) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedL2VPNTermination) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPNTermination) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedL2VPNTermination) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedL2VPNTermination) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPNTermination) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedL2VPNTermination) SetDisplay(v string) {
	o.Display = v
}

// GetL2vpn returns the L2vpn field value
func (o *NestedL2VPNTermination) GetL2vpn() NestedL2VPN {
	if o == nil {
		var ret NestedL2VPN
		return ret
	}

	return o.L2vpn
}

// GetL2vpnOk returns a tuple with the L2vpn field value
// and a boolean to check if the value has been set.
func (o *NestedL2VPNTermination) GetL2vpnOk() (*NestedL2VPN, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2vpn, true
}

// SetL2vpn sets field value
func (o *NestedL2VPNTermination) SetL2vpn(v NestedL2VPN) {
	o.L2vpn = v
}

func (o NestedL2VPNTermination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedL2VPNTermination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["l2vpn"] = o.L2vpn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedL2VPNTermination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"l2vpn",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNestedL2VPNTermination := _NestedL2VPNTermination{}

	err = json.Unmarshal(data, &varNestedL2VPNTermination)

	if err != nil {
		return err
	}

	*o = NestedL2VPNTermination(varNestedL2VPNTermination)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "l2vpn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedL2VPNTermination struct {
	value *NestedL2VPNTermination
	isSet bool
}

func (v NullableNestedL2VPNTermination) Get() *NestedL2VPNTermination {
	return v.value
}

func (v *NullableNestedL2VPNTermination) Set(val *NestedL2VPNTermination) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedL2VPNTermination) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedL2VPNTermination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedL2VPNTermination(val *NestedL2VPNTermination) *NullableNestedL2VPNTermination {
	return &NullableNestedL2VPNTermination{value: val, isSet: true}
}

func (v NullableNestedL2VPNTermination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedL2VPNTermination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
