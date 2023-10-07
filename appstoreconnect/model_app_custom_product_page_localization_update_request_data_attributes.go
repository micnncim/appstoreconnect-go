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

// checks if the AppCustomProductPageLocalizationUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationUpdateRequestDataAttributes{}

// AppCustomProductPageLocalizationUpdateRequestDataAttributes struct for AppCustomProductPageLocalizationUpdateRequestDataAttributes
type AppCustomProductPageLocalizationUpdateRequestDataAttributes struct {
	PromotionalText *string `json:"promotionalText,omitempty"`
}

// NewAppCustomProductPageLocalizationUpdateRequestDataAttributes instantiates a new AppCustomProductPageLocalizationUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationUpdateRequestDataAttributes() *AppCustomProductPageLocalizationUpdateRequestDataAttributes {
	this := AppCustomProductPageLocalizationUpdateRequestDataAttributes{}
	return &this
}

// NewAppCustomProductPageLocalizationUpdateRequestDataAttributesWithDefaults instantiates a new AppCustomProductPageLocalizationUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationUpdateRequestDataAttributesWithDefaults() *AppCustomProductPageLocalizationUpdateRequestDataAttributes {
	this := AppCustomProductPageLocalizationUpdateRequestDataAttributes{}
	return &this
}

// GetPromotionalText returns the PromotionalText field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationUpdateRequestDataAttributes) GetPromotionalText() string {
	if o == nil || IsNil(o.PromotionalText) {
		var ret string
		return ret
	}
	return *o.PromotionalText
}

// GetPromotionalTextOk returns a tuple with the PromotionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationUpdateRequestDataAttributes) GetPromotionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionalText) {
		return nil, false
	}
	return o.PromotionalText, true
}

// HasPromotionalText returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationUpdateRequestDataAttributes) HasPromotionalText() bool {
	if o != nil && !IsNil(o.PromotionalText) {
		return true
	}

	return false
}

// SetPromotionalText gets a reference to the given string and assigns it to the PromotionalText field.
func (o *AppCustomProductPageLocalizationUpdateRequestDataAttributes) SetPromotionalText(v string) {
	o.PromotionalText = &v
}

func (o AppCustomProductPageLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromotionalText) {
		toSerialize["promotionalText"] = o.PromotionalText
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes struct {
	value *AppCustomProductPageLocalizationUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes) Get() *AppCustomProductPageLocalizationUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes) Set(val *AppCustomProductPageLocalizationUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationUpdateRequestDataAttributes(val *AppCustomProductPageLocalizationUpdateRequestDataAttributes) *NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes {
	return &NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
