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

// checks if the VLANGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VLANGroupRequest{}

// VLANGroupRequest Adds support for custom fields and tags.
type VLANGroupRequest struct {
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	ScopeType NullableString `json:"scope_type,omitempty"`
	ScopeId   NullableInt32  `json:"scope_id,omitempty"`
	// Lowest permissible ID of a child VLAN
	MinVid *int32 `json:"min_vid,omitempty"`
	// Highest permissible ID of a child VLAN
	MaxVid               *int32                 `json:"max_vid,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VLANGroupRequest VLANGroupRequest

// NewVLANGroupRequest instantiates a new VLANGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVLANGroupRequest(name string, slug string) *VLANGroupRequest {
	this := VLANGroupRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewVLANGroupRequestWithDefaults instantiates a new VLANGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVLANGroupRequestWithDefaults() *VLANGroupRequest {
	this := VLANGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *VLANGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VLANGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VLANGroupRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *VLANGroupRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *VLANGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *VLANGroupRequest) SetSlug(v string) {
	o.Slug = v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLANGroupRequest) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType.Get()) {
		var ret string
		return ret
	}
	return *o.ScopeType.Get()
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroupRequest) GetScopeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeType.Get(), o.ScopeType.IsSet()
}

// HasScopeType returns a boolean if a field has been set.
func (o *VLANGroupRequest) HasScopeType() bool {
	if o != nil && o.ScopeType.IsSet() {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given NullableString and assigns it to the ScopeType field.
func (o *VLANGroupRequest) SetScopeType(v string) {
	o.ScopeType.Set(&v)
}

// SetScopeTypeNil sets the value for ScopeType to be an explicit nil
func (o *VLANGroupRequest) SetScopeTypeNil() {
	o.ScopeType.Set(nil)
}

// UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
func (o *VLANGroupRequest) UnsetScopeType() {
	o.ScopeType.Unset()
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLANGroupRequest) GetScopeId() int32 {
	if o == nil || IsNil(o.ScopeId.Get()) {
		var ret int32
		return ret
	}
	return *o.ScopeId.Get()
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroupRequest) GetScopeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeId.Get(), o.ScopeId.IsSet()
}

// HasScopeId returns a boolean if a field has been set.
func (o *VLANGroupRequest) HasScopeId() bool {
	if o != nil && o.ScopeId.IsSet() {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given NullableInt32 and assigns it to the ScopeId field.
func (o *VLANGroupRequest) SetScopeId(v int32) {
	o.ScopeId.Set(&v)
}

// SetScopeIdNil sets the value for ScopeId to be an explicit nil
func (o *VLANGroupRequest) SetScopeIdNil() {
	o.ScopeId.Set(nil)
}

// UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
func (o *VLANGroupRequest) UnsetScopeId() {
	o.ScopeId.Unset()
}

// GetMinVid returns the MinVid field value if set, zero value otherwise.
func (o *VLANGroupRequest) GetMinVid() int32 {
	if o == nil || IsNil(o.MinVid) {
		var ret int32
		return ret
	}
	return *o.MinVid
}

// GetMinVidOk returns a tuple with the MinVid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroupRequest) GetMinVidOk() (*int32, bool) {
	if o == nil || IsNil(o.MinVid) {
		return nil, false
	}
	return o.MinVid, true
}

// HasMinVid returns a boolean if a field has been set.
func (o *VLANGroupRequest) HasMinVid() bool {
	if o != nil && !IsNil(o.MinVid) {
		return true
	}

	return false
}

// SetMinVid gets a reference to the given int32 and assigns it to the MinVid field.
func (o *VLANGroupRequest) SetMinVid(v int32) {
	o.MinVid = &v
}

// GetMaxVid returns the MaxVid field value if set, zero value otherwise.
func (o *VLANGroupRequest) GetMaxVid() int32 {
	if o == nil || IsNil(o.MaxVid) {
		var ret int32
		return ret
	}
	return *o.MaxVid
}

// GetMaxVidOk returns a tuple with the MaxVid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroupRequest) GetMaxVidOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxVid) {
		return nil, false
	}
	return o.MaxVid, true
}

// HasMaxVid returns a boolean if a field has been set.
func (o *VLANGroupRequest) HasMaxVid() bool {
	if o != nil && !IsNil(o.MaxVid) {
		return true
	}

	return false
}

// SetMaxVid gets a reference to the given int32 and assigns it to the MaxVid field.
func (o *VLANGroupRequest) SetMaxVid(v int32) {
	o.MaxVid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VLANGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VLANGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VLANGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VLANGroupRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroupRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VLANGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *VLANGroupRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VLANGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VLANGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VLANGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o VLANGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VLANGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.ScopeType.IsSet() {
		toSerialize["scope_type"] = o.ScopeType.Get()
	}
	if o.ScopeId.IsSet() {
		toSerialize["scope_id"] = o.ScopeId.Get()
	}
	if !IsNil(o.MinVid) {
		toSerialize["min_vid"] = o.MinVid
	}
	if !IsNil(o.MaxVid) {
		toSerialize["max_vid"] = o.MaxVid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VLANGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varVLANGroupRequest := _VLANGroupRequest{}

	err = json.Unmarshal(data, &varVLANGroupRequest)

	if err != nil {
		return err
	}

	*o = VLANGroupRequest(varVLANGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "scope_type")
		delete(additionalProperties, "scope_id")
		delete(additionalProperties, "min_vid")
		delete(additionalProperties, "max_vid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVLANGroupRequest struct {
	value *VLANGroupRequest
	isSet bool
}

func (v NullableVLANGroupRequest) Get() *VLANGroupRequest {
	return v.value
}

func (v *NullableVLANGroupRequest) Set(val *VLANGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVLANGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVLANGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVLANGroupRequest(val *VLANGroupRequest) *NullableVLANGroupRequest {
	return &NullableVLANGroupRequest{value: val, isSet: true}
}

func (v NullableVLANGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVLANGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}