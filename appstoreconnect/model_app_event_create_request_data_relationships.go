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

// checks if the AppEventCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventCreateRequestDataRelationships{}

// AppEventCreateRequestDataRelationships struct for AppEventCreateRequestDataRelationships
type AppEventCreateRequestDataRelationships struct {
	App AppAvailabilityV2CreateRequestDataRelationshipsApp `json:"app"`
}

// NewAppEventCreateRequestDataRelationships instantiates a new AppEventCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventCreateRequestDataRelationships(app AppAvailabilityV2CreateRequestDataRelationshipsApp) *AppEventCreateRequestDataRelationships {
	this := AppEventCreateRequestDataRelationships{}
	this.App = app
	return &this
}

// NewAppEventCreateRequestDataRelationshipsWithDefaults instantiates a new AppEventCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventCreateRequestDataRelationshipsWithDefaults() *AppEventCreateRequestDataRelationships {
	this := AppEventCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *AppEventCreateRequestDataRelationships) GetApp() AppAvailabilityV2CreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataRelationships) GetAppOk() (*AppAvailabilityV2CreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *AppEventCreateRequestDataRelationships) SetApp(v AppAvailabilityV2CreateRequestDataRelationshipsApp) {
	o.App = v
}

func (o AppEventCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	return toSerialize, nil
}

type NullableAppEventCreateRequestDataRelationships struct {
	value *AppEventCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppEventCreateRequestDataRelationships) Get() *AppEventCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppEventCreateRequestDataRelationships) Set(val *AppEventCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventCreateRequestDataRelationships(val *AppEventCreateRequestDataRelationships) *NullableAppEventCreateRequestDataRelationships {
	return &NullableAppEventCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppEventCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
