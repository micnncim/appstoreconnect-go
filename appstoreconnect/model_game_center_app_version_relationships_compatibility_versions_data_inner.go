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

// checks if the GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner{}

// GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner struct for GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner
type GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner instantiates a new GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner(type_ string, id string) *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner {
	this := GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterAppVersionRelationshipsCompatibilityVersionsDataInnerWithDefaults instantiates a new GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionRelationshipsCompatibilityVersionsDataInnerWithDefaults() *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner {
	this := GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) SetId(v string) {
	o.Id = v
}

func (o GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner struct {
	value *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner
	isSet bool
}

func (v NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) Get() *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner {
	return v.value
}

func (v *NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) Set(val *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner(val *GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) *NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner {
	return &NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
