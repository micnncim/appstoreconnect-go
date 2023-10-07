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

// checks if the AppAvailabilityV2CreateRequestDataRelationshipsApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityV2CreateRequestDataRelationshipsApp{}

// AppAvailabilityV2CreateRequestDataRelationshipsApp struct for AppAvailabilityV2CreateRequestDataRelationshipsApp
type AppAvailabilityV2CreateRequestDataRelationshipsApp struct {
	Data AppAvailabilityV2CreateRequestDataRelationshipsAppData `json:"data"`
}

// NewAppAvailabilityV2CreateRequestDataRelationshipsApp instantiates a new AppAvailabilityV2CreateRequestDataRelationshipsApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityV2CreateRequestDataRelationshipsApp(data AppAvailabilityV2CreateRequestDataRelationshipsAppData) *AppAvailabilityV2CreateRequestDataRelationshipsApp {
	this := AppAvailabilityV2CreateRequestDataRelationshipsApp{}
	this.Data = data
	return &this
}

// NewAppAvailabilityV2CreateRequestDataRelationshipsAppWithDefaults instantiates a new AppAvailabilityV2CreateRequestDataRelationshipsApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityV2CreateRequestDataRelationshipsAppWithDefaults() *AppAvailabilityV2CreateRequestDataRelationshipsApp {
	this := AppAvailabilityV2CreateRequestDataRelationshipsApp{}
	return &this
}

// GetData returns the Data field value
func (o *AppAvailabilityV2CreateRequestDataRelationshipsApp) GetData() AppAvailabilityV2CreateRequestDataRelationshipsAppData {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsAppData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2CreateRequestDataRelationshipsApp) GetDataOk() (*AppAvailabilityV2CreateRequestDataRelationshipsAppData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppAvailabilityV2CreateRequestDataRelationshipsApp) SetData(v AppAvailabilityV2CreateRequestDataRelationshipsAppData) {
	o.Data = v
}

func (o AppAvailabilityV2CreateRequestDataRelationshipsApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityV2CreateRequestDataRelationshipsApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppAvailabilityV2CreateRequestDataRelationshipsApp struct {
	value *AppAvailabilityV2CreateRequestDataRelationshipsApp
	isSet bool
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationshipsApp) Get() *AppAvailabilityV2CreateRequestDataRelationshipsApp {
	return v.value
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationshipsApp) Set(val *AppAvailabilityV2CreateRequestDataRelationshipsApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationshipsApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationshipsApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityV2CreateRequestDataRelationshipsApp(val *AppAvailabilityV2CreateRequestDataRelationshipsApp) *NullableAppAvailabilityV2CreateRequestDataRelationshipsApp {
	return &NullableAppAvailabilityV2CreateRequestDataRelationshipsApp{value: val, isSet: true}
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationshipsApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationshipsApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
