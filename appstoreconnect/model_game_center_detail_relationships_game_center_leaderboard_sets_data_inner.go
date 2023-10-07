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

// checks if the GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner{}

// GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner struct for GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner
type GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner instantiates a new GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner(type_ string, id string) *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	this := GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInnerWithDefaults instantiates a new GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInnerWithDefaults() *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	this := GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) SetId(v string) {
	o.Id = v
}

func (o GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner struct {
	value *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner
	isSet bool
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) Get() *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	return v.value
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) Set(val *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner(val *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	return &NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner{value: val, isSet: true}
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
