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

// checks if the BetaAppClipInvocationLocalizationInlineCreateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationLocalizationInlineCreateAttributes{}

// BetaAppClipInvocationLocalizationInlineCreateAttributes struct for BetaAppClipInvocationLocalizationInlineCreateAttributes
type BetaAppClipInvocationLocalizationInlineCreateAttributes struct {
	Title  string `json:"title"`
	Locale string `json:"locale"`
}

// NewBetaAppClipInvocationLocalizationInlineCreateAttributes instantiates a new BetaAppClipInvocationLocalizationInlineCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationLocalizationInlineCreateAttributes(title string, locale string) *BetaAppClipInvocationLocalizationInlineCreateAttributes {
	this := BetaAppClipInvocationLocalizationInlineCreateAttributes{}
	this.Title = title
	this.Locale = locale
	return &this
}

// NewBetaAppClipInvocationLocalizationInlineCreateAttributesWithDefaults instantiates a new BetaAppClipInvocationLocalizationInlineCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationLocalizationInlineCreateAttributesWithDefaults() *BetaAppClipInvocationLocalizationInlineCreateAttributes {
	this := BetaAppClipInvocationLocalizationInlineCreateAttributes{}
	return &this
}

// GetTitle returns the Title field value
func (o *BetaAppClipInvocationLocalizationInlineCreateAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreateAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BetaAppClipInvocationLocalizationInlineCreateAttributes) SetTitle(v string) {
	o.Title = v
}

// GetLocale returns the Locale field value
func (o *BetaAppClipInvocationLocalizationInlineCreateAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreateAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *BetaAppClipInvocationLocalizationInlineCreateAttributes) SetLocale(v string) {
	o.Locale = v
}

func (o BetaAppClipInvocationLocalizationInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationLocalizationInlineCreateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["locale"] = o.Locale
	return toSerialize, nil
}

type NullableBetaAppClipInvocationLocalizationInlineCreateAttributes struct {
	value *BetaAppClipInvocationLocalizationInlineCreateAttributes
	isSet bool
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreateAttributes) Get() *BetaAppClipInvocationLocalizationInlineCreateAttributes {
	return v.value
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreateAttributes) Set(val *BetaAppClipInvocationLocalizationInlineCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationLocalizationInlineCreateAttributes(val *BetaAppClipInvocationLocalizationInlineCreateAttributes) *NullableBetaAppClipInvocationLocalizationInlineCreateAttributes {
	return &NullableBetaAppClipInvocationLocalizationInlineCreateAttributes{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
