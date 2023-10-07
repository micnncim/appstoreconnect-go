/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
	"fmt"
)

// BetaInviteType the model 'BetaInviteType'
type BetaInviteType string

// List of BetaInviteType
const (
	BETAINVITETYPE_EMAIL       BetaInviteType = "EMAIL"
	BETAINVITETYPE_PUBLIC_LINK BetaInviteType = "PUBLIC_LINK"
)

// All allowed values of BetaInviteType enum
var AllowedBetaInviteTypeEnumValues = []BetaInviteType{
	"EMAIL",
	"PUBLIC_LINK",
}

func (v *BetaInviteType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BetaInviteType(value)
	for _, existing := range AllowedBetaInviteTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BetaInviteType", value)
}

// NewBetaInviteTypeFromValue returns a pointer to a valid BetaInviteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBetaInviteTypeFromValue(v string) (*BetaInviteType, error) {
	ev := BetaInviteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BetaInviteType: valid values are %v", v, AllowedBetaInviteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BetaInviteType) IsValid() bool {
	for _, existing := range AllowedBetaInviteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BetaInviteType value
func (v BetaInviteType) Ptr() *BetaInviteType {
	return &v
}

type NullableBetaInviteType struct {
	value *BetaInviteType
	isSet bool
}

func (v NullableBetaInviteType) Get() *BetaInviteType {
	return v.value
}

func (v *NullableBetaInviteType) Set(val *BetaInviteType) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaInviteType) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaInviteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaInviteType(val *BetaInviteType) *NullableBetaInviteType {
	return &NullableBetaInviteType{value: val, isSet: true}
}

func (v NullableBetaInviteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaInviteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
