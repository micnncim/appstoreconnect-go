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

// checks if the BetaAppClipInvocationLocalizationUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationLocalizationUpdateRequestDataAttributes{}

// BetaAppClipInvocationLocalizationUpdateRequestDataAttributes struct for BetaAppClipInvocationLocalizationUpdateRequestDataAttributes
type BetaAppClipInvocationLocalizationUpdateRequestDataAttributes struct {
	Title *string `json:"title,omitempty"`
}

// NewBetaAppClipInvocationLocalizationUpdateRequestDataAttributes instantiates a new BetaAppClipInvocationLocalizationUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationLocalizationUpdateRequestDataAttributes() *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes {
	this := BetaAppClipInvocationLocalizationUpdateRequestDataAttributes{}
	return &this
}

// NewBetaAppClipInvocationLocalizationUpdateRequestDataAttributesWithDefaults instantiates a new BetaAppClipInvocationLocalizationUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationLocalizationUpdateRequestDataAttributesWithDefaults() *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes {
	this := BetaAppClipInvocationLocalizationUpdateRequestDataAttributes{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) SetTitle(v string) {
	o.Title = &v
}

func (o BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes struct {
	value *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes
	isSet bool
}

func (v NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes) Get() *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes) Set(val *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes(val *BetaAppClipInvocationLocalizationUpdateRequestDataAttributes) *NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes {
	return &NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationLocalizationUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
