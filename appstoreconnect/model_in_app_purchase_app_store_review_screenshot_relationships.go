/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
)

// checks if the InAppPurchaseAppStoreReviewScreenshotRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAppStoreReviewScreenshotRelationships{}

// InAppPurchaseAppStoreReviewScreenshotRelationships struct for InAppPurchaseAppStoreReviewScreenshotRelationships
type InAppPurchaseAppStoreReviewScreenshotRelationships struct {
	InAppPurchaseV2 *InAppPurchaseAppStoreReviewScreenshotRelationshipsInAppPurchaseV2 `json:"inAppPurchaseV2,omitempty"`
}

// NewInAppPurchaseAppStoreReviewScreenshotRelationships instantiates a new InAppPurchaseAppStoreReviewScreenshotRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAppStoreReviewScreenshotRelationships() *InAppPurchaseAppStoreReviewScreenshotRelationships {
	this := InAppPurchaseAppStoreReviewScreenshotRelationships{}
	return &this
}

// NewInAppPurchaseAppStoreReviewScreenshotRelationshipsWithDefaults instantiates a new InAppPurchaseAppStoreReviewScreenshotRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAppStoreReviewScreenshotRelationshipsWithDefaults() *InAppPurchaseAppStoreReviewScreenshotRelationships {
	this := InAppPurchaseAppStoreReviewScreenshotRelationships{}
	return &this
}

// GetInAppPurchaseV2 returns the InAppPurchaseV2 field value if set, zero value otherwise.
func (o *InAppPurchaseAppStoreReviewScreenshotRelationships) GetInAppPurchaseV2() InAppPurchaseAppStoreReviewScreenshotRelationshipsInAppPurchaseV2 {
	if o == nil || IsNil(o.InAppPurchaseV2) {
		var ret InAppPurchaseAppStoreReviewScreenshotRelationshipsInAppPurchaseV2
		return ret
	}
	return *o.InAppPurchaseV2
}

// GetInAppPurchaseV2Ok returns a tuple with the InAppPurchaseV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotRelationships) GetInAppPurchaseV2Ok() (*InAppPurchaseAppStoreReviewScreenshotRelationshipsInAppPurchaseV2, bool) {
	if o == nil || IsNil(o.InAppPurchaseV2) {
		return nil, false
	}
	return o.InAppPurchaseV2, true
}

// HasInAppPurchaseV2 returns a boolean if a field has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotRelationships) HasInAppPurchaseV2() bool {
	if o != nil && !IsNil(o.InAppPurchaseV2) {
		return true
	}

	return false
}

// SetInAppPurchaseV2 gets a reference to the given InAppPurchaseAppStoreReviewScreenshotRelationshipsInAppPurchaseV2 and assigns it to the InAppPurchaseV2 field.
func (o *InAppPurchaseAppStoreReviewScreenshotRelationships) SetInAppPurchaseV2(v InAppPurchaseAppStoreReviewScreenshotRelationshipsInAppPurchaseV2) {
	o.InAppPurchaseV2 = &v
}

func (o InAppPurchaseAppStoreReviewScreenshotRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAppStoreReviewScreenshotRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InAppPurchaseV2) {
		toSerialize["inAppPurchaseV2"] = o.InAppPurchaseV2
	}
	return toSerialize, nil
}

type NullableInAppPurchaseAppStoreReviewScreenshotRelationships struct {
	value *InAppPurchaseAppStoreReviewScreenshotRelationships
	isSet bool
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotRelationships) Get() *InAppPurchaseAppStoreReviewScreenshotRelationships {
	return v.value
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotRelationships) Set(val *InAppPurchaseAppStoreReviewScreenshotRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAppStoreReviewScreenshotRelationships(val *InAppPurchaseAppStoreReviewScreenshotRelationships) *NullableInAppPurchaseAppStoreReviewScreenshotRelationships {
	return &NullableInAppPurchaseAppStoreReviewScreenshotRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
