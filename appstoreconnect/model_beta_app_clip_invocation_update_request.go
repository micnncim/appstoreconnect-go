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

// checks if the BetaAppClipInvocationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationUpdateRequest{}

// BetaAppClipInvocationUpdateRequest struct for BetaAppClipInvocationUpdateRequest
type BetaAppClipInvocationUpdateRequest struct {
	Data BetaAppClipInvocationUpdateRequestData `json:"data"`
}

// NewBetaAppClipInvocationUpdateRequest instantiates a new BetaAppClipInvocationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationUpdateRequest(data BetaAppClipInvocationUpdateRequestData) *BetaAppClipInvocationUpdateRequest {
	this := BetaAppClipInvocationUpdateRequest{}
	this.Data = data
	return &this
}

// NewBetaAppClipInvocationUpdateRequestWithDefaults instantiates a new BetaAppClipInvocationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationUpdateRequestWithDefaults() *BetaAppClipInvocationUpdateRequest {
	this := BetaAppClipInvocationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationUpdateRequest) GetData() BetaAppClipInvocationUpdateRequestData {
	if o == nil {
		var ret BetaAppClipInvocationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationUpdateRequest) GetDataOk() (*BetaAppClipInvocationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationUpdateRequest) SetData(v BetaAppClipInvocationUpdateRequestData) {
	o.Data = v
}

func (o BetaAppClipInvocationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaAppClipInvocationUpdateRequest struct {
	value *BetaAppClipInvocationUpdateRequest
	isSet bool
}

func (v NullableBetaAppClipInvocationUpdateRequest) Get() *BetaAppClipInvocationUpdateRequest {
	return v.value
}

func (v *NullableBetaAppClipInvocationUpdateRequest) Set(val *BetaAppClipInvocationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationUpdateRequest(val *BetaAppClipInvocationUpdateRequest) *NullableBetaAppClipInvocationUpdateRequest {
	return &NullableBetaAppClipInvocationUpdateRequest{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
