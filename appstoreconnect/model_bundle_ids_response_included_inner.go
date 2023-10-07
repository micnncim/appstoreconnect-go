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

// BundleIdsResponseIncludedInner - struct for BundleIdsResponseIncludedInner
type BundleIdsResponseIncludedInner struct {
	App                *App
	BundleIdCapability *BundleIdCapability
	Profile            *Profile
}

// AppAsBundleIdsResponseIncludedInner is a convenience function that returns App wrapped in BundleIdsResponseIncludedInner
func AppAsBundleIdsResponseIncludedInner(v *App) BundleIdsResponseIncludedInner {
	return BundleIdsResponseIncludedInner{
		App: v,
	}
}

// BundleIdCapabilityAsBundleIdsResponseIncludedInner is a convenience function that returns BundleIdCapability wrapped in BundleIdsResponseIncludedInner
func BundleIdCapabilityAsBundleIdsResponseIncludedInner(v *BundleIdCapability) BundleIdsResponseIncludedInner {
	return BundleIdsResponseIncludedInner{
		BundleIdCapability: v,
	}
}

// ProfileAsBundleIdsResponseIncludedInner is a convenience function that returns Profile wrapped in BundleIdsResponseIncludedInner
func ProfileAsBundleIdsResponseIncludedInner(v *Profile) BundleIdsResponseIncludedInner {
	return BundleIdsResponseIncludedInner{
		Profile: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BundleIdsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into App
	err = newStrictDecoder(data).Decode(&dst.App)
	if err == nil {
		jsonApp, _ := json.Marshal(dst.App)
		if string(jsonApp) == "{}" { // empty struct
			dst.App = nil
		} else {
			match++
		}
	} else {
		dst.App = nil
	}

	// try to unmarshal data into BundleIdCapability
	err = newStrictDecoder(data).Decode(&dst.BundleIdCapability)
	if err == nil {
		jsonBundleIdCapability, _ := json.Marshal(dst.BundleIdCapability)
		if string(jsonBundleIdCapability) == "{}" { // empty struct
			dst.BundleIdCapability = nil
		} else {
			match++
		}
	} else {
		dst.BundleIdCapability = nil
	}

	// try to unmarshal data into Profile
	err = newStrictDecoder(data).Decode(&dst.Profile)
	if err == nil {
		jsonProfile, _ := json.Marshal(dst.Profile)
		if string(jsonProfile) == "{}" { // empty struct
			dst.Profile = nil
		} else {
			match++
		}
	} else {
		dst.Profile = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.App = nil
		dst.BundleIdCapability = nil
		dst.Profile = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BundleIdsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BundleIdsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BundleIdsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.BundleIdCapability != nil {
		return json.Marshal(&src.BundleIdCapability)
	}

	if src.Profile != nil {
		return json.Marshal(&src.Profile)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BundleIdsResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.App != nil {
		return obj.App
	}

	if obj.BundleIdCapability != nil {
		return obj.BundleIdCapability
	}

	if obj.Profile != nil {
		return obj.Profile
	}

	// all schemas are nil
	return nil
}

type NullableBundleIdsResponseIncludedInner struct {
	value *BundleIdsResponseIncludedInner
	isSet bool
}

func (v NullableBundleIdsResponseIncludedInner) Get() *BundleIdsResponseIncludedInner {
	return v.value
}

func (v *NullableBundleIdsResponseIncludedInner) Set(val *BundleIdsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdsResponseIncludedInner(val *BundleIdsResponseIncludedInner) *NullableBundleIdsResponseIncludedInner {
	return &NullableBundleIdsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableBundleIdsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
