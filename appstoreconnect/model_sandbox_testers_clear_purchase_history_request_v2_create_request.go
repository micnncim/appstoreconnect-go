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

// checks if the SandboxTestersClearPurchaseHistoryRequestV2CreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTestersClearPurchaseHistoryRequestV2CreateRequest{}

// SandboxTestersClearPurchaseHistoryRequestV2CreateRequest struct for SandboxTestersClearPurchaseHistoryRequestV2CreateRequest
type SandboxTestersClearPurchaseHistoryRequestV2CreateRequest struct {
	Data SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData `json:"data"`
}

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequest instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequest(data SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequest{}
	this.Data = data
	return &this
}

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestWithDefaults instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestWithDefaults() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) GetData() SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData {
	if o == nil {
		var ret SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) GetDataOk() (*SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) SetData(v SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) {
	o.Data = v
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest struct {
	value *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest
	isSet bool
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest) Get() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest {
	return v.value
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest) Set(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequest) *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest {
	return &NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest{value: val, isSet: true}
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
