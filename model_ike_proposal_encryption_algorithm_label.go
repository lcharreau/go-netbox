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

// IKEProposalEncryptionAlgorithmLabel the model 'IKEProposalEncryptionAlgorithmLabel'
type IKEProposalEncryptionAlgorithmLabel string

// List of IKEProposal_encryption_algorithm_label
const (
	IKEPROPOSALENCRYPTIONALGORITHMLABEL__128_BIT_AES__CBC IKEProposalEncryptionAlgorithmLabel = "128-bit AES (CBC)"
	IKEPROPOSALENCRYPTIONALGORITHMLABEL__128_BIT_AES__GCM IKEProposalEncryptionAlgorithmLabel = "128-bit AES (GCM)"
	IKEPROPOSALENCRYPTIONALGORITHMLABEL__192_BIT_AES__CBC IKEProposalEncryptionAlgorithmLabel = "192-bit AES (CBC)"
	IKEPROPOSALENCRYPTIONALGORITHMLABEL__192_BIT_AES__GCM IKEProposalEncryptionAlgorithmLabel = "192-bit AES (GCM)"
	IKEPROPOSALENCRYPTIONALGORITHMLABEL__256_BIT_AES__CBC IKEProposalEncryptionAlgorithmLabel = "256-bit AES (CBC)"
	IKEPROPOSALENCRYPTIONALGORITHMLABEL__256_BIT_AES__GCM IKEProposalEncryptionAlgorithmLabel = "256-bit AES (GCM)"
	IKEPROPOSALENCRYPTIONALGORITHMLABEL__3_DES            IKEProposalEncryptionAlgorithmLabel = "3DES"
	IKEPROPOSALENCRYPTIONALGORITHMLABEL_DES               IKEProposalEncryptionAlgorithmLabel = "DES"
)

// All allowed values of IKEProposalEncryptionAlgorithmLabel enum
var AllowedIKEProposalEncryptionAlgorithmLabelEnumValues = []IKEProposalEncryptionAlgorithmLabel{
	"128-bit AES (CBC)",
	"128-bit AES (GCM)",
	"192-bit AES (CBC)",
	"192-bit AES (GCM)",
	"256-bit AES (CBC)",
	"256-bit AES (GCM)",
	"3DES",
	"DES",
}

func (v *IKEProposalEncryptionAlgorithmLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEProposalEncryptionAlgorithmLabel(value)
	for _, existing := range AllowedIKEProposalEncryptionAlgorithmLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEProposalEncryptionAlgorithmLabel", value)
}

// NewIKEProposalEncryptionAlgorithmLabelFromValue returns a pointer to a valid IKEProposalEncryptionAlgorithmLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEProposalEncryptionAlgorithmLabelFromValue(v string) (*IKEProposalEncryptionAlgorithmLabel, error) {
	ev := IKEProposalEncryptionAlgorithmLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEProposalEncryptionAlgorithmLabel: valid values are %v", v, AllowedIKEProposalEncryptionAlgorithmLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEProposalEncryptionAlgorithmLabel) IsValid() bool {
	for _, existing := range AllowedIKEProposalEncryptionAlgorithmLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEProposal_encryption_algorithm_label value
func (v IKEProposalEncryptionAlgorithmLabel) Ptr() *IKEProposalEncryptionAlgorithmLabel {
	return &v
}

type NullableIKEProposalEncryptionAlgorithmLabel struct {
	value *IKEProposalEncryptionAlgorithmLabel
	isSet bool
}

func (v NullableIKEProposalEncryptionAlgorithmLabel) Get() *IKEProposalEncryptionAlgorithmLabel {
	return v.value
}

func (v *NullableIKEProposalEncryptionAlgorithmLabel) Set(val *IKEProposalEncryptionAlgorithmLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalEncryptionAlgorithmLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalEncryptionAlgorithmLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalEncryptionAlgorithmLabel(val *IKEProposalEncryptionAlgorithmLabel) *NullableIKEProposalEncryptionAlgorithmLabel {
	return &NullableIKEProposalEncryptionAlgorithmLabel{value: val, isSet: true}
}

func (v NullableIKEProposalEncryptionAlgorithmLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalEncryptionAlgorithmLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
