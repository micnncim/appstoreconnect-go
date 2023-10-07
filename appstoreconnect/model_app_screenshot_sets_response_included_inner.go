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

// AppScreenshotSetsResponseIncludedInner - struct for AppScreenshotSetsResponseIncludedInner
type AppScreenshotSetsResponseIncludedInner struct {
	AppCustomProductPageLocalization               *AppCustomProductPageLocalization
	AppScreenshot                                  *AppScreenshot
	AppStoreVersionExperimentTreatmentLocalization *AppStoreVersionExperimentTreatmentLocalization
	AppStoreVersionLocalization                    *AppStoreVersionLocalization
}

// AppCustomProductPageLocalizationAsAppScreenshotSetsResponseIncludedInner is a convenience function that returns AppCustomProductPageLocalization wrapped in AppScreenshotSetsResponseIncludedInner
func AppCustomProductPageLocalizationAsAppScreenshotSetsResponseIncludedInner(v *AppCustomProductPageLocalization) AppScreenshotSetsResponseIncludedInner {
	return AppScreenshotSetsResponseIncludedInner{
		AppCustomProductPageLocalization: v,
	}
}

// AppScreenshotAsAppScreenshotSetsResponseIncludedInner is a convenience function that returns AppScreenshot wrapped in AppScreenshotSetsResponseIncludedInner
func AppScreenshotAsAppScreenshotSetsResponseIncludedInner(v *AppScreenshot) AppScreenshotSetsResponseIncludedInner {
	return AppScreenshotSetsResponseIncludedInner{
		AppScreenshot: v,
	}
}

// AppStoreVersionExperimentTreatmentLocalizationAsAppScreenshotSetsResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentTreatmentLocalization wrapped in AppScreenshotSetsResponseIncludedInner
func AppStoreVersionExperimentTreatmentLocalizationAsAppScreenshotSetsResponseIncludedInner(v *AppStoreVersionExperimentTreatmentLocalization) AppScreenshotSetsResponseIncludedInner {
	return AppScreenshotSetsResponseIncludedInner{
		AppStoreVersionExperimentTreatmentLocalization: v,
	}
}

// AppStoreVersionLocalizationAsAppScreenshotSetsResponseIncludedInner is a convenience function that returns AppStoreVersionLocalization wrapped in AppScreenshotSetsResponseIncludedInner
func AppStoreVersionLocalizationAsAppScreenshotSetsResponseIncludedInner(v *AppStoreVersionLocalization) AppScreenshotSetsResponseIncludedInner {
	return AppScreenshotSetsResponseIncludedInner{
		AppStoreVersionLocalization: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppScreenshotSetsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppCustomProductPageLocalization
	err = newStrictDecoder(data).Decode(&dst.AppCustomProductPageLocalization)
	if err == nil {
		jsonAppCustomProductPageLocalization, _ := json.Marshal(dst.AppCustomProductPageLocalization)
		if string(jsonAppCustomProductPageLocalization) == "{}" { // empty struct
			dst.AppCustomProductPageLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppCustomProductPageLocalization = nil
	}

	// try to unmarshal data into AppScreenshot
	err = newStrictDecoder(data).Decode(&dst.AppScreenshot)
	if err == nil {
		jsonAppScreenshot, _ := json.Marshal(dst.AppScreenshot)
		if string(jsonAppScreenshot) == "{}" { // empty struct
			dst.AppScreenshot = nil
		} else {
			match++
		}
	} else {
		dst.AppScreenshot = nil
	}

	// try to unmarshal data into AppStoreVersionExperimentTreatmentLocalization
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperimentTreatmentLocalization)
	if err == nil {
		jsonAppStoreVersionExperimentTreatmentLocalization, _ := json.Marshal(dst.AppStoreVersionExperimentTreatmentLocalization)
		if string(jsonAppStoreVersionExperimentTreatmentLocalization) == "{}" { // empty struct
			dst.AppStoreVersionExperimentTreatmentLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperimentTreatmentLocalization = nil
	}

	// try to unmarshal data into AppStoreVersionLocalization
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionLocalization)
	if err == nil {
		jsonAppStoreVersionLocalization, _ := json.Marshal(dst.AppStoreVersionLocalization)
		if string(jsonAppStoreVersionLocalization) == "{}" { // empty struct
			dst.AppStoreVersionLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionLocalization = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppCustomProductPageLocalization = nil
		dst.AppScreenshot = nil
		dst.AppStoreVersionExperimentTreatmentLocalization = nil
		dst.AppStoreVersionLocalization = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppScreenshotSetsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppScreenshotSetsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppScreenshotSetsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppCustomProductPageLocalization != nil {
		return json.Marshal(&src.AppCustomProductPageLocalization)
	}

	if src.AppScreenshot != nil {
		return json.Marshal(&src.AppScreenshot)
	}

	if src.AppStoreVersionExperimentTreatmentLocalization != nil {
		return json.Marshal(&src.AppStoreVersionExperimentTreatmentLocalization)
	}

	if src.AppStoreVersionLocalization != nil {
		return json.Marshal(&src.AppStoreVersionLocalization)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppScreenshotSetsResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AppCustomProductPageLocalization != nil {
		return obj.AppCustomProductPageLocalization
	}

	if obj.AppScreenshot != nil {
		return obj.AppScreenshot
	}

	if obj.AppStoreVersionExperimentTreatmentLocalization != nil {
		return obj.AppStoreVersionExperimentTreatmentLocalization
	}

	if obj.AppStoreVersionLocalization != nil {
		return obj.AppStoreVersionLocalization
	}

	// all schemas are nil
	return nil
}

type NullableAppScreenshotSetsResponseIncludedInner struct {
	value *AppScreenshotSetsResponseIncludedInner
	isSet bool
}

func (v NullableAppScreenshotSetsResponseIncludedInner) Get() *AppScreenshotSetsResponseIncludedInner {
	return v.value
}

func (v *NullableAppScreenshotSetsResponseIncludedInner) Set(val *AppScreenshotSetsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetsResponseIncludedInner(val *AppScreenshotSetsResponseIncludedInner) *NullableAppScreenshotSetsResponseIncludedInner {
	return &NullableAppScreenshotSetsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppScreenshotSetsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
