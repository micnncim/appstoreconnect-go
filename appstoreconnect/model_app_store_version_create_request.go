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

// checks if the AppStoreVersionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionCreateRequest{}

// AppStoreVersionCreateRequest struct for AppStoreVersionCreateRequest
type AppStoreVersionCreateRequest struct {
	Data AppStoreVersionCreateRequestData `json:"data"`
}

// NewAppStoreVersionCreateRequest instantiates a new AppStoreVersionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionCreateRequest(data AppStoreVersionCreateRequestData) *AppStoreVersionCreateRequest {
	this := AppStoreVersionCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionCreateRequestWithDefaults instantiates a new AppStoreVersionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionCreateRequestWithDefaults() *AppStoreVersionCreateRequest {
	this := AppStoreVersionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionCreateRequest) GetData() AppStoreVersionCreateRequestData {
	if o == nil {
		var ret AppStoreVersionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequest) GetDataOk() (*AppStoreVersionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionCreateRequest) SetData(v AppStoreVersionCreateRequestData) {
	o.Data = v
}

func (o AppStoreVersionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionCreateRequest struct {
	value *AppStoreVersionCreateRequest
	isSet bool
}

func (v NullableAppStoreVersionCreateRequest) Get() *AppStoreVersionCreateRequest {
	return v.value
}

func (v *NullableAppStoreVersionCreateRequest) Set(val *AppStoreVersionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionCreateRequest(val *AppStoreVersionCreateRequest) *NullableAppStoreVersionCreateRequest {
	return &NullableAppStoreVersionCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
