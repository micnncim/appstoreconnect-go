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

// checks if the SubscriptionLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionLocalizationUpdateRequest{}

// SubscriptionLocalizationUpdateRequest struct for SubscriptionLocalizationUpdateRequest
type SubscriptionLocalizationUpdateRequest struct {
	Data SubscriptionLocalizationUpdateRequestData `json:"data"`
}

// NewSubscriptionLocalizationUpdateRequest instantiates a new SubscriptionLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionLocalizationUpdateRequest(data SubscriptionLocalizationUpdateRequestData) *SubscriptionLocalizationUpdateRequest {
	this := SubscriptionLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionLocalizationUpdateRequestWithDefaults instantiates a new SubscriptionLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionLocalizationUpdateRequestWithDefaults() *SubscriptionLocalizationUpdateRequest {
	this := SubscriptionLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionLocalizationUpdateRequest) GetData() SubscriptionLocalizationUpdateRequestData {
	if o == nil {
		var ret SubscriptionLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionLocalizationUpdateRequest) GetDataOk() (*SubscriptionLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionLocalizationUpdateRequest) SetData(v SubscriptionLocalizationUpdateRequestData) {
	o.Data = v
}

func (o SubscriptionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionLocalizationUpdateRequest struct {
	value *SubscriptionLocalizationUpdateRequest
	isSet bool
}

func (v NullableSubscriptionLocalizationUpdateRequest) Get() *SubscriptionLocalizationUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionLocalizationUpdateRequest) Set(val *SubscriptionLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionLocalizationUpdateRequest(val *SubscriptionLocalizationUpdateRequest) *NullableSubscriptionLocalizationUpdateRequest {
	return &NullableSubscriptionLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
