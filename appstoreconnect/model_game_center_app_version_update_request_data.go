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

// checks if the GameCenterAppVersionUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionUpdateRequestData{}

// GameCenterAppVersionUpdateRequestData struct for GameCenterAppVersionUpdateRequestData
type GameCenterAppVersionUpdateRequestData struct {
	Type       string                          `json:"type"`
	Id         string                          `json:"id"`
	Attributes *GameCenterAppVersionAttributes `json:"attributes,omitempty"`
}

// NewGameCenterAppVersionUpdateRequestData instantiates a new GameCenterAppVersionUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionUpdateRequestData(type_ string, id string) *GameCenterAppVersionUpdateRequestData {
	this := GameCenterAppVersionUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterAppVersionUpdateRequestDataWithDefaults instantiates a new GameCenterAppVersionUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionUpdateRequestDataWithDefaults() *GameCenterAppVersionUpdateRequestData {
	this := GameCenterAppVersionUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAppVersionUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAppVersionUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterAppVersionUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterAppVersionUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterAppVersionUpdateRequestData) GetAttributes() GameCenterAppVersionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterAppVersionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionUpdateRequestData) GetAttributesOk() (*GameCenterAppVersionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterAppVersionUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterAppVersionAttributes and assigns it to the Attributes field.
func (o *GameCenterAppVersionUpdateRequestData) SetAttributes(v GameCenterAppVersionAttributes) {
	o.Attributes = &v
}

func (o GameCenterAppVersionUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGameCenterAppVersionUpdateRequestData struct {
	value *GameCenterAppVersionUpdateRequestData
	isSet bool
}

func (v NullableGameCenterAppVersionUpdateRequestData) Get() *GameCenterAppVersionUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterAppVersionUpdateRequestData) Set(val *GameCenterAppVersionUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionUpdateRequestData(val *GameCenterAppVersionUpdateRequestData) *NullableGameCenterAppVersionUpdateRequestData {
	return &NullableGameCenterAppVersionUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
