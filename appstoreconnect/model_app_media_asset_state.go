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

// checks if the AppMediaAssetState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppMediaAssetState{}

// AppMediaAssetState struct for AppMediaAssetState
type AppMediaAssetState struct {
	Errors   []AppMediaStateError `json:"errors,omitempty"`
	Warnings []AppMediaStateError `json:"warnings,omitempty"`
	State    *string              `json:"state,omitempty"`
}

// NewAppMediaAssetState instantiates a new AppMediaAssetState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppMediaAssetState() *AppMediaAssetState {
	this := AppMediaAssetState{}
	return &this
}

// NewAppMediaAssetStateWithDefaults instantiates a new AppMediaAssetState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppMediaAssetStateWithDefaults() *AppMediaAssetState {
	this := AppMediaAssetState{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AppMediaAssetState) GetErrors() []AppMediaStateError {
	if o == nil || IsNil(o.Errors) {
		var ret []AppMediaStateError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMediaAssetState) GetErrorsOk() ([]AppMediaStateError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AppMediaAssetState) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []AppMediaStateError and assigns it to the Errors field.
func (o *AppMediaAssetState) SetErrors(v []AppMediaStateError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AppMediaAssetState) GetWarnings() []AppMediaStateError {
	if o == nil || IsNil(o.Warnings) {
		var ret []AppMediaStateError
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMediaAssetState) GetWarningsOk() ([]AppMediaStateError, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AppMediaAssetState) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AppMediaStateError and assigns it to the Warnings field.
func (o *AppMediaAssetState) SetWarnings(v []AppMediaStateError) {
	o.Warnings = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AppMediaAssetState) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMediaAssetState) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AppMediaAssetState) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AppMediaAssetState) SetState(v string) {
	o.State = &v
}

func (o AppMediaAssetState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppMediaAssetState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableAppMediaAssetState struct {
	value *AppMediaAssetState
	isSet bool
}

func (v NullableAppMediaAssetState) Get() *AppMediaAssetState {
	return v.value
}

func (v *NullableAppMediaAssetState) Set(val *AppMediaAssetState) {
	v.value = val
	v.isSet = true
}

func (v NullableAppMediaAssetState) IsSet() bool {
	return v.isSet
}

func (v *NullableAppMediaAssetState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppMediaAssetState(val *AppMediaAssetState) *NullableAppMediaAssetState {
	return &NullableAppMediaAssetState{value: val, isSet: true}
}

func (v NullableAppMediaAssetState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppMediaAssetState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
