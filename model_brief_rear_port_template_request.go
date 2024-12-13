/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the BriefRearPortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefRearPortTemplateRequest{}

// BriefRearPortTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BriefRearPortTemplateRequest struct {
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefRearPortTemplateRequest BriefRearPortTemplateRequest

// NewBriefRearPortTemplateRequest instantiates a new BriefRearPortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefRearPortTemplateRequest(name string) *BriefRearPortTemplateRequest {
	this := BriefRearPortTemplateRequest{}
	this.Name = name
	return &this
}

// NewBriefRearPortTemplateRequestWithDefaults instantiates a new BriefRearPortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefRearPortTemplateRequestWithDefaults() *BriefRearPortTemplateRequest {
	this := BriefRearPortTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefRearPortTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefRearPortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefRearPortTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefRearPortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefRearPortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefRearPortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefRearPortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefRearPortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefRearPortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefRearPortTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varBriefRearPortTemplateRequest := _BriefRearPortTemplateRequest{}

	err = json.Unmarshal(data, &varBriefRearPortTemplateRequest)

	if err != nil {
		return err
	}

	*o = BriefRearPortTemplateRequest(varBriefRearPortTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefRearPortTemplateRequest struct {
	value *BriefRearPortTemplateRequest
	isSet bool
}

func (v NullableBriefRearPortTemplateRequest) Get() *BriefRearPortTemplateRequest {
	return v.value
}

func (v *NullableBriefRearPortTemplateRequest) Set(val *BriefRearPortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefRearPortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefRearPortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefRearPortTemplateRequest(val *BriefRearPortTemplateRequest) *NullableBriefRearPortTemplateRequest {
	return &NullableBriefRearPortTemplateRequest{value: val, isSet: true}
}

func (v NullableBriefRearPortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefRearPortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
