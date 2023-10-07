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

// checks if the AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner{}

// AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner struct for AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner
type AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner instantiates a new AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner(type_ string, id string) *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner {
	this := AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInnerWithDefaults instantiates a new AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInnerWithDefaults() *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner {
	this := AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner struct {
	value *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner
	isSet bool
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) Get() *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner {
	return v.value
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) Set(val *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner(val *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner {
	return &NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
