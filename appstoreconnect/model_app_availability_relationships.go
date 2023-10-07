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

// checks if the AppAvailabilityRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityRelationships{}

// AppAvailabilityRelationships struct for AppAvailabilityRelationships
type AppAvailabilityRelationships struct {
	App                  *AppAvailabilityRelationshipsApp                  `json:"app,omitempty"`
	AvailableTerritories *AppAvailabilityRelationshipsAvailableTerritories `json:"availableTerritories,omitempty"`
}

// NewAppAvailabilityRelationships instantiates a new AppAvailabilityRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityRelationships() *AppAvailabilityRelationships {
	this := AppAvailabilityRelationships{}
	return &this
}

// NewAppAvailabilityRelationshipsWithDefaults instantiates a new AppAvailabilityRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityRelationshipsWithDefaults() *AppAvailabilityRelationships {
	this := AppAvailabilityRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AppAvailabilityRelationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AppAvailabilityRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *AppAvailabilityRelationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

// GetAvailableTerritories returns the AvailableTerritories field value if set, zero value otherwise.
func (o *AppAvailabilityRelationships) GetAvailableTerritories() AppAvailabilityRelationshipsAvailableTerritories {
	if o == nil || IsNil(o.AvailableTerritories) {
		var ret AppAvailabilityRelationshipsAvailableTerritories
		return ret
	}
	return *o.AvailableTerritories
}

// GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityRelationships) GetAvailableTerritoriesOk() (*AppAvailabilityRelationshipsAvailableTerritories, bool) {
	if o == nil || IsNil(o.AvailableTerritories) {
		return nil, false
	}
	return o.AvailableTerritories, true
}

// HasAvailableTerritories returns a boolean if a field has been set.
func (o *AppAvailabilityRelationships) HasAvailableTerritories() bool {
	if o != nil && !IsNil(o.AvailableTerritories) {
		return true
	}

	return false
}

// SetAvailableTerritories gets a reference to the given AppAvailabilityRelationshipsAvailableTerritories and assigns it to the AvailableTerritories field.
func (o *AppAvailabilityRelationships) SetAvailableTerritories(v AppAvailabilityRelationshipsAvailableTerritories) {
	o.AvailableTerritories = &v
}

func (o AppAvailabilityRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.AvailableTerritories) {
		toSerialize["availableTerritories"] = o.AvailableTerritories
	}
	return toSerialize, nil
}

type NullableAppAvailabilityRelationships struct {
	value *AppAvailabilityRelationships
	isSet bool
}

func (v NullableAppAvailabilityRelationships) Get() *AppAvailabilityRelationships {
	return v.value
}

func (v *NullableAppAvailabilityRelationships) Set(val *AppAvailabilityRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityRelationships(val *AppAvailabilityRelationships) *NullableAppAvailabilityRelationships {
	return &NullableAppAvailabilityRelationships{value: val, isSet: true}
}

func (v NullableAppAvailabilityRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
