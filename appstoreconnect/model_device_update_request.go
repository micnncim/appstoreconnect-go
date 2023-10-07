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

// checks if the DeviceUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceUpdateRequest{}

// DeviceUpdateRequest struct for DeviceUpdateRequest
type DeviceUpdateRequest struct {
	Data DeviceUpdateRequestData `json:"data"`
}

// NewDeviceUpdateRequest instantiates a new DeviceUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUpdateRequest(data DeviceUpdateRequestData) *DeviceUpdateRequest {
	this := DeviceUpdateRequest{}
	this.Data = data
	return &this
}

// NewDeviceUpdateRequestWithDefaults instantiates a new DeviceUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUpdateRequestWithDefaults() *DeviceUpdateRequest {
	this := DeviceUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *DeviceUpdateRequest) GetData() DeviceUpdateRequestData {
	if o == nil {
		var ret DeviceUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdateRequest) GetDataOk() (*DeviceUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeviceUpdateRequest) SetData(v DeviceUpdateRequestData) {
	o.Data = v
}

func (o DeviceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableDeviceUpdateRequest struct {
	value *DeviceUpdateRequest
	isSet bool
}

func (v NullableDeviceUpdateRequest) Get() *DeviceUpdateRequest {
	return v.value
}

func (v *NullableDeviceUpdateRequest) Set(val *DeviceUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUpdateRequest(val *DeviceUpdateRequest) *NullableDeviceUpdateRequest {
	return &NullableDeviceUpdateRequest{value: val, isSet: true}
}

func (v NullableDeviceUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
