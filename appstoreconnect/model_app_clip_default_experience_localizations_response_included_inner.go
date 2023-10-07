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

// AppClipDefaultExperienceLocalizationsResponseIncludedInner - struct for AppClipDefaultExperienceLocalizationsResponseIncludedInner
type AppClipDefaultExperienceLocalizationsResponseIncludedInner struct {
	AppClipDefaultExperience *AppClipDefaultExperience
	AppClipHeaderImage       *AppClipHeaderImage
}

// AppClipDefaultExperienceAsAppClipDefaultExperienceLocalizationsResponseIncludedInner is a convenience function that returns AppClipDefaultExperience wrapped in AppClipDefaultExperienceLocalizationsResponseIncludedInner
func AppClipDefaultExperienceAsAppClipDefaultExperienceLocalizationsResponseIncludedInner(v *AppClipDefaultExperience) AppClipDefaultExperienceLocalizationsResponseIncludedInner {
	return AppClipDefaultExperienceLocalizationsResponseIncludedInner{
		AppClipDefaultExperience: v,
	}
}

// AppClipHeaderImageAsAppClipDefaultExperienceLocalizationsResponseIncludedInner is a convenience function that returns AppClipHeaderImage wrapped in AppClipDefaultExperienceLocalizationsResponseIncludedInner
func AppClipHeaderImageAsAppClipDefaultExperienceLocalizationsResponseIncludedInner(v *AppClipHeaderImage) AppClipDefaultExperienceLocalizationsResponseIncludedInner {
	return AppClipDefaultExperienceLocalizationsResponseIncludedInner{
		AppClipHeaderImage: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppClipDefaultExperienceLocalizationsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppClipDefaultExperience
	err = newStrictDecoder(data).Decode(&dst.AppClipDefaultExperience)
	if err == nil {
		jsonAppClipDefaultExperience, _ := json.Marshal(dst.AppClipDefaultExperience)
		if string(jsonAppClipDefaultExperience) == "{}" { // empty struct
			dst.AppClipDefaultExperience = nil
		} else {
			match++
		}
	} else {
		dst.AppClipDefaultExperience = nil
	}

	// try to unmarshal data into AppClipHeaderImage
	err = newStrictDecoder(data).Decode(&dst.AppClipHeaderImage)
	if err == nil {
		jsonAppClipHeaderImage, _ := json.Marshal(dst.AppClipHeaderImage)
		if string(jsonAppClipHeaderImage) == "{}" { // empty struct
			dst.AppClipHeaderImage = nil
		} else {
			match++
		}
	} else {
		dst.AppClipHeaderImage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppClipDefaultExperience = nil
		dst.AppClipHeaderImage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppClipDefaultExperienceLocalizationsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppClipDefaultExperienceLocalizationsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppClipDefaultExperienceLocalizationsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppClipDefaultExperience != nil {
		return json.Marshal(&src.AppClipDefaultExperience)
	}

	if src.AppClipHeaderImage != nil {
		return json.Marshal(&src.AppClipHeaderImage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppClipDefaultExperienceLocalizationsResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AppClipDefaultExperience != nil {
		return obj.AppClipDefaultExperience
	}

	if obj.AppClipHeaderImage != nil {
		return obj.AppClipHeaderImage
	}

	// all schemas are nil
	return nil
}

type NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner struct {
	value *AppClipDefaultExperienceLocalizationsResponseIncludedInner
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner) Get() *AppClipDefaultExperienceLocalizationsResponseIncludedInner {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner) Set(val *AppClipDefaultExperienceLocalizationsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalizationsResponseIncludedInner(val *AppClipDefaultExperienceLocalizationsResponseIncludedInner) *NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner {
	return &NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalizationsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
