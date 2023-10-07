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

// AppStoreVersionExperimentTreatmentsResponseIncludedInner - struct for AppStoreVersionExperimentTreatmentsResponseIncludedInner
type AppStoreVersionExperimentTreatmentsResponseIncludedInner struct {
	AppStoreVersionExperiment                      *AppStoreVersionExperiment
	AppStoreVersionExperimentTreatmentLocalization *AppStoreVersionExperimentTreatmentLocalization
	AppStoreVersionExperimentV2                    *AppStoreVersionExperimentV2
}

// AppStoreVersionExperimentAsAppStoreVersionExperimentTreatmentsResponseIncludedInner is a convenience function that returns AppStoreVersionExperiment wrapped in AppStoreVersionExperimentTreatmentsResponseIncludedInner
func AppStoreVersionExperimentAsAppStoreVersionExperimentTreatmentsResponseIncludedInner(v *AppStoreVersionExperiment) AppStoreVersionExperimentTreatmentsResponseIncludedInner {
	return AppStoreVersionExperimentTreatmentsResponseIncludedInner{
		AppStoreVersionExperiment: v,
	}
}

// AppStoreVersionExperimentTreatmentLocalizationAsAppStoreVersionExperimentTreatmentsResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentTreatmentLocalization wrapped in AppStoreVersionExperimentTreatmentsResponseIncludedInner
func AppStoreVersionExperimentTreatmentLocalizationAsAppStoreVersionExperimentTreatmentsResponseIncludedInner(v *AppStoreVersionExperimentTreatmentLocalization) AppStoreVersionExperimentTreatmentsResponseIncludedInner {
	return AppStoreVersionExperimentTreatmentsResponseIncludedInner{
		AppStoreVersionExperimentTreatmentLocalization: v,
	}
}

// AppStoreVersionExperimentV2AsAppStoreVersionExperimentTreatmentsResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentV2 wrapped in AppStoreVersionExperimentTreatmentsResponseIncludedInner
func AppStoreVersionExperimentV2AsAppStoreVersionExperimentTreatmentsResponseIncludedInner(v *AppStoreVersionExperimentV2) AppStoreVersionExperimentTreatmentsResponseIncludedInner {
	return AppStoreVersionExperimentTreatmentsResponseIncludedInner{
		AppStoreVersionExperimentV2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppStoreVersionExperimentTreatmentsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppStoreVersionExperiment
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperiment)
	if err == nil {
		jsonAppStoreVersionExperiment, _ := json.Marshal(dst.AppStoreVersionExperiment)
		if string(jsonAppStoreVersionExperiment) == "{}" { // empty struct
			dst.AppStoreVersionExperiment = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperiment = nil
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

	// try to unmarshal data into AppStoreVersionExperimentV2
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperimentV2)
	if err == nil {
		jsonAppStoreVersionExperimentV2, _ := json.Marshal(dst.AppStoreVersionExperimentV2)
		if string(jsonAppStoreVersionExperimentV2) == "{}" { // empty struct
			dst.AppStoreVersionExperimentV2 = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperimentV2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppStoreVersionExperiment = nil
		dst.AppStoreVersionExperimentTreatmentLocalization = nil
		dst.AppStoreVersionExperimentV2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppStoreVersionExperimentTreatmentsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppStoreVersionExperimentTreatmentsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppStoreVersionExperimentTreatmentsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppStoreVersionExperiment != nil {
		return json.Marshal(&src.AppStoreVersionExperiment)
	}

	if src.AppStoreVersionExperimentTreatmentLocalization != nil {
		return json.Marshal(&src.AppStoreVersionExperimentTreatmentLocalization)
	}

	if src.AppStoreVersionExperimentV2 != nil {
		return json.Marshal(&src.AppStoreVersionExperimentV2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppStoreVersionExperimentTreatmentsResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AppStoreVersionExperiment != nil {
		return obj.AppStoreVersionExperiment
	}

	if obj.AppStoreVersionExperimentTreatmentLocalization != nil {
		return obj.AppStoreVersionExperimentTreatmentLocalization
	}

	if obj.AppStoreVersionExperimentV2 != nil {
		return obj.AppStoreVersionExperimentV2
	}

	// all schemas are nil
	return nil
}

type NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner struct {
	value *AppStoreVersionExperimentTreatmentsResponseIncludedInner
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner) Get() *AppStoreVersionExperimentTreatmentsResponseIncludedInner {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner) Set(val *AppStoreVersionExperimentTreatmentsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentsResponseIncludedInner(val *AppStoreVersionExperimentTreatmentsResponseIncludedInner) *NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner {
	return &NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
