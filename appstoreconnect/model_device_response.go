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

// checks if the DeviceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceResponse{}

// DeviceResponse struct for DeviceResponse
type DeviceResponse struct {
	Data  Device        `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewDeviceResponse instantiates a new DeviceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceResponse(data Device, links DocumentLinks) *DeviceResponse {
	this := DeviceResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewDeviceResponseWithDefaults instantiates a new DeviceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceResponseWithDefaults() *DeviceResponse {
	this := DeviceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeviceResponse) GetData() Device {
	if o == nil {
		var ret Device
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetDataOk() (*Device, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeviceResponse) SetData(v Device) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *DeviceResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DeviceResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o DeviceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableDeviceResponse struct {
	value *DeviceResponse
	isSet bool
}

func (v NullableDeviceResponse) Get() *DeviceResponse {
	return v.value
}

func (v *NullableDeviceResponse) Set(val *DeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceResponse(val *DeviceResponse) *NullableDeviceResponse {
	return &NullableDeviceResponse{value: val, isSet: true}
}

func (v NullableDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
