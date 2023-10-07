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

// checks if the AppMediaStateError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppMediaStateError{}

// AppMediaStateError struct for AppMediaStateError
type AppMediaStateError struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewAppMediaStateError instantiates a new AppMediaStateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppMediaStateError() *AppMediaStateError {
	this := AppMediaStateError{}
	return &this
}

// NewAppMediaStateErrorWithDefaults instantiates a new AppMediaStateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppMediaStateErrorWithDefaults() *AppMediaStateError {
	this := AppMediaStateError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AppMediaStateError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMediaStateError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AppMediaStateError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AppMediaStateError) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppMediaStateError) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMediaStateError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppMediaStateError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppMediaStateError) SetDescription(v string) {
	o.Description = &v
}

func (o AppMediaStateError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppMediaStateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAppMediaStateError struct {
	value *AppMediaStateError
	isSet bool
}

func (v NullableAppMediaStateError) Get() *AppMediaStateError {
	return v.value
}

func (v *NullableAppMediaStateError) Set(val *AppMediaStateError) {
	v.value = val
	v.isSet = true
}

func (v NullableAppMediaStateError) IsSet() bool {
	return v.isSet
}

func (v *NullableAppMediaStateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppMediaStateError(val *AppMediaStateError) *NullableAppMediaStateError {
	return &NullableAppMediaStateError{value: val, isSet: true}
}

func (v NullableAppMediaStateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppMediaStateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
