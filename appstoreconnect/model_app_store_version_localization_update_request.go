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

// checks if the AppStoreVersionLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationUpdateRequest{}

// AppStoreVersionLocalizationUpdateRequest struct for AppStoreVersionLocalizationUpdateRequest
type AppStoreVersionLocalizationUpdateRequest struct {
	Data AppStoreVersionLocalizationUpdateRequestData `json:"data"`
}

// NewAppStoreVersionLocalizationUpdateRequest instantiates a new AppStoreVersionLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationUpdateRequest(data AppStoreVersionLocalizationUpdateRequestData) *AppStoreVersionLocalizationUpdateRequest {
	this := AppStoreVersionLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionLocalizationUpdateRequestWithDefaults instantiates a new AppStoreVersionLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationUpdateRequestWithDefaults() *AppStoreVersionLocalizationUpdateRequest {
	this := AppStoreVersionLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionLocalizationUpdateRequest) GetData() AppStoreVersionLocalizationUpdateRequestData {
	if o == nil {
		var ret AppStoreVersionLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationUpdateRequest) GetDataOk() (*AppStoreVersionLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionLocalizationUpdateRequest) SetData(v AppStoreVersionLocalizationUpdateRequestData) {
	o.Data = v
}

func (o AppStoreVersionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionLocalizationUpdateRequest struct {
	value *AppStoreVersionLocalizationUpdateRequest
	isSet bool
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) Get() *AppStoreVersionLocalizationUpdateRequest {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) Set(val *AppStoreVersionLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationUpdateRequest(val *AppStoreVersionLocalizationUpdateRequest) *NullableAppStoreVersionLocalizationUpdateRequest {
	return &NullableAppStoreVersionLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
