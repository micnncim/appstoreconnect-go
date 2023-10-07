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

// checks if the AppPreOrderCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreOrderCreateRequestDataAttributes{}

// AppPreOrderCreateRequestDataAttributes struct for AppPreOrderCreateRequestDataAttributes
type AppPreOrderCreateRequestDataAttributes struct {
	AppReleaseDate *string `json:"appReleaseDate,omitempty"`
}

// NewAppPreOrderCreateRequestDataAttributes instantiates a new AppPreOrderCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreOrderCreateRequestDataAttributes() *AppPreOrderCreateRequestDataAttributes {
	this := AppPreOrderCreateRequestDataAttributes{}
	return &this
}

// NewAppPreOrderCreateRequestDataAttributesWithDefaults instantiates a new AppPreOrderCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreOrderCreateRequestDataAttributesWithDefaults() *AppPreOrderCreateRequestDataAttributes {
	this := AppPreOrderCreateRequestDataAttributes{}
	return &this
}

// GetAppReleaseDate returns the AppReleaseDate field value if set, zero value otherwise.
func (o *AppPreOrderCreateRequestDataAttributes) GetAppReleaseDate() string {
	if o == nil || IsNil(o.AppReleaseDate) {
		var ret string
		return ret
	}
	return *o.AppReleaseDate
}

// GetAppReleaseDateOk returns a tuple with the AppReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreOrderCreateRequestDataAttributes) GetAppReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.AppReleaseDate) {
		return nil, false
	}
	return o.AppReleaseDate, true
}

// HasAppReleaseDate returns a boolean if a field has been set.
func (o *AppPreOrderCreateRequestDataAttributes) HasAppReleaseDate() bool {
	if o != nil && !IsNil(o.AppReleaseDate) {
		return true
	}

	return false
}

// SetAppReleaseDate gets a reference to the given string and assigns it to the AppReleaseDate field.
func (o *AppPreOrderCreateRequestDataAttributes) SetAppReleaseDate(v string) {
	o.AppReleaseDate = &v
}

func (o AppPreOrderCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreOrderCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppReleaseDate) {
		toSerialize["appReleaseDate"] = o.AppReleaseDate
	}
	return toSerialize, nil
}

type NullableAppPreOrderCreateRequestDataAttributes struct {
	value *AppPreOrderCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppPreOrderCreateRequestDataAttributes) Get() *AppPreOrderCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppPreOrderCreateRequestDataAttributes) Set(val *AppPreOrderCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreOrderCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreOrderCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreOrderCreateRequestDataAttributes(val *AppPreOrderCreateRequestDataAttributes) *NullableAppPreOrderCreateRequestDataAttributes {
	return &NullableAppPreOrderCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppPreOrderCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreOrderCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
