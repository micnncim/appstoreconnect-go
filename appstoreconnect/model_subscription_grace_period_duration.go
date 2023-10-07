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

// SubscriptionGracePeriodDuration the model 'SubscriptionGracePeriodDuration'
type SubscriptionGracePeriodDuration string

// List of SubscriptionGracePeriodDuration
const (
	SUBSCRIPTIONGRACEPERIODDURATION_THREE_DAYS        SubscriptionGracePeriodDuration = "THREE_DAYS"
	SUBSCRIPTIONGRACEPERIODDURATION_SIXTEEN_DAYS      SubscriptionGracePeriodDuration = "SIXTEEN_DAYS"
	SUBSCRIPTIONGRACEPERIODDURATION_TWENTY_EIGHT_DAYS SubscriptionGracePeriodDuration = "TWENTY_EIGHT_DAYS"
)

// All allowed values of SubscriptionGracePeriodDuration enum
var AllowedSubscriptionGracePeriodDurationEnumValues = []SubscriptionGracePeriodDuration{
	"THREE_DAYS",
	"SIXTEEN_DAYS",
	"TWENTY_EIGHT_DAYS",
}

func (v *SubscriptionGracePeriodDuration) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionGracePeriodDuration(value)
	for _, existing := range AllowedSubscriptionGracePeriodDurationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionGracePeriodDuration", value)
}

// NewSubscriptionGracePeriodDurationFromValue returns a pointer to a valid SubscriptionGracePeriodDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionGracePeriodDurationFromValue(v string) (*SubscriptionGracePeriodDuration, error) {
	ev := SubscriptionGracePeriodDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionGracePeriodDuration: valid values are %v", v, AllowedSubscriptionGracePeriodDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionGracePeriodDuration) IsValid() bool {
	for _, existing := range AllowedSubscriptionGracePeriodDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionGracePeriodDuration value
func (v SubscriptionGracePeriodDuration) Ptr() *SubscriptionGracePeriodDuration {
	return &v
}

type NullableSubscriptionGracePeriodDuration struct {
	value *SubscriptionGracePeriodDuration
	isSet bool
}

func (v NullableSubscriptionGracePeriodDuration) Get() *SubscriptionGracePeriodDuration {
	return v.value
}

func (v *NullableSubscriptionGracePeriodDuration) Set(val *SubscriptionGracePeriodDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGracePeriodDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGracePeriodDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGracePeriodDuration(val *SubscriptionGracePeriodDuration) *NullableSubscriptionGracePeriodDuration {
	return &NullableSubscriptionGracePeriodDuration{value: val, isSet: true}
}

func (v NullableSubscriptionGracePeriodDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGracePeriodDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
