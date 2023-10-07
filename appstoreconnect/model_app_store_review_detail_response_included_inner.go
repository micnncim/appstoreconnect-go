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

// AppStoreReviewDetailResponseIncludedInner - struct for AppStoreReviewDetailResponseIncludedInner
type AppStoreReviewDetailResponseIncludedInner struct {
	AppStoreReviewAttachment *AppStoreReviewAttachment
	AppStoreVersion          *AppStoreVersion
}

// AppStoreReviewAttachmentAsAppStoreReviewDetailResponseIncludedInner is a convenience function that returns AppStoreReviewAttachment wrapped in AppStoreReviewDetailResponseIncludedInner
func AppStoreReviewAttachmentAsAppStoreReviewDetailResponseIncludedInner(v *AppStoreReviewAttachment) AppStoreReviewDetailResponseIncludedInner {
	return AppStoreReviewDetailResponseIncludedInner{
		AppStoreReviewAttachment: v,
	}
}

// AppStoreVersionAsAppStoreReviewDetailResponseIncludedInner is a convenience function that returns AppStoreVersion wrapped in AppStoreReviewDetailResponseIncludedInner
func AppStoreVersionAsAppStoreReviewDetailResponseIncludedInner(v *AppStoreVersion) AppStoreReviewDetailResponseIncludedInner {
	return AppStoreReviewDetailResponseIncludedInner{
		AppStoreVersion: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppStoreReviewDetailResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppStoreReviewAttachment
	err = newStrictDecoder(data).Decode(&dst.AppStoreReviewAttachment)
	if err == nil {
		jsonAppStoreReviewAttachment, _ := json.Marshal(dst.AppStoreReviewAttachment)
		if string(jsonAppStoreReviewAttachment) == "{}" { // empty struct
			dst.AppStoreReviewAttachment = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreReviewAttachment = nil
	}

	// try to unmarshal data into AppStoreVersion
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersion)
	if err == nil {
		jsonAppStoreVersion, _ := json.Marshal(dst.AppStoreVersion)
		if string(jsonAppStoreVersion) == "{}" { // empty struct
			dst.AppStoreVersion = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersion = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppStoreReviewAttachment = nil
		dst.AppStoreVersion = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppStoreReviewDetailResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppStoreReviewDetailResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppStoreReviewDetailResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppStoreReviewAttachment != nil {
		return json.Marshal(&src.AppStoreReviewAttachment)
	}

	if src.AppStoreVersion != nil {
		return json.Marshal(&src.AppStoreVersion)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppStoreReviewDetailResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AppStoreReviewAttachment != nil {
		return obj.AppStoreReviewAttachment
	}

	if obj.AppStoreVersion != nil {
		return obj.AppStoreVersion
	}

	// all schemas are nil
	return nil
}

type NullableAppStoreReviewDetailResponseIncludedInner struct {
	value *AppStoreReviewDetailResponseIncludedInner
	isSet bool
}

func (v NullableAppStoreReviewDetailResponseIncludedInner) Get() *AppStoreReviewDetailResponseIncludedInner {
	return v.value
}

func (v *NullableAppStoreReviewDetailResponseIncludedInner) Set(val *AppStoreReviewDetailResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailResponseIncludedInner(val *AppStoreReviewDetailResponseIncludedInner) *NullableAppStoreReviewDetailResponseIncludedInner {
	return &NullableAppStoreReviewDetailResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
