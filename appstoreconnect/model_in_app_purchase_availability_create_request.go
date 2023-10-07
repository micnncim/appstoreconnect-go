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

// checks if the InAppPurchaseAvailabilityCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAvailabilityCreateRequest{}

// InAppPurchaseAvailabilityCreateRequest struct for InAppPurchaseAvailabilityCreateRequest
type InAppPurchaseAvailabilityCreateRequest struct {
	Data InAppPurchaseAvailabilityCreateRequestData `json:"data"`
}

// NewInAppPurchaseAvailabilityCreateRequest instantiates a new InAppPurchaseAvailabilityCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAvailabilityCreateRequest(data InAppPurchaseAvailabilityCreateRequestData) *InAppPurchaseAvailabilityCreateRequest {
	this := InAppPurchaseAvailabilityCreateRequest{}
	this.Data = data
	return &this
}

// NewInAppPurchaseAvailabilityCreateRequestWithDefaults instantiates a new InAppPurchaseAvailabilityCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAvailabilityCreateRequestWithDefaults() *InAppPurchaseAvailabilityCreateRequest {
	this := InAppPurchaseAvailabilityCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseAvailabilityCreateRequest) GetData() InAppPurchaseAvailabilityCreateRequestData {
	if o == nil {
		var ret InAppPurchaseAvailabilityCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailabilityCreateRequest) GetDataOk() (*InAppPurchaseAvailabilityCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseAvailabilityCreateRequest) SetData(v InAppPurchaseAvailabilityCreateRequestData) {
	o.Data = v
}

func (o InAppPurchaseAvailabilityCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAvailabilityCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInAppPurchaseAvailabilityCreateRequest struct {
	value *InAppPurchaseAvailabilityCreateRequest
	isSet bool
}

func (v NullableInAppPurchaseAvailabilityCreateRequest) Get() *InAppPurchaseAvailabilityCreateRequest {
	return v.value
}

func (v *NullableInAppPurchaseAvailabilityCreateRequest) Set(val *InAppPurchaseAvailabilityCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAvailabilityCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAvailabilityCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAvailabilityCreateRequest(val *InAppPurchaseAvailabilityCreateRequest) *NullableInAppPurchaseAvailabilityCreateRequest {
	return &NullableInAppPurchaseAvailabilityCreateRequest{value: val, isSet: true}
}

func (v NullableInAppPurchaseAvailabilityCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAvailabilityCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
