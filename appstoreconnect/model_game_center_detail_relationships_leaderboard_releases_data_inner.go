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

// checks if the GameCenterDetailRelationshipsLeaderboardReleasesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailRelationshipsLeaderboardReleasesDataInner{}

// GameCenterDetailRelationshipsLeaderboardReleasesDataInner struct for GameCenterDetailRelationshipsLeaderboardReleasesDataInner
type GameCenterDetailRelationshipsLeaderboardReleasesDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewGameCenterDetailRelationshipsLeaderboardReleasesDataInner instantiates a new GameCenterDetailRelationshipsLeaderboardReleasesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailRelationshipsLeaderboardReleasesDataInner(type_ string, id string) *GameCenterDetailRelationshipsLeaderboardReleasesDataInner {
	this := GameCenterDetailRelationshipsLeaderboardReleasesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterDetailRelationshipsLeaderboardReleasesDataInnerWithDefaults instantiates a new GameCenterDetailRelationshipsLeaderboardReleasesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailRelationshipsLeaderboardReleasesDataInnerWithDefaults() *GameCenterDetailRelationshipsLeaderboardReleasesDataInner {
	this := GameCenterDetailRelationshipsLeaderboardReleasesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) SetId(v string) {
	o.Id = v
}

func (o GameCenterDetailRelationshipsLeaderboardReleasesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailRelationshipsLeaderboardReleasesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner struct {
	value *GameCenterDetailRelationshipsLeaderboardReleasesDataInner
	isSet bool
}

func (v NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner) Get() *GameCenterDetailRelationshipsLeaderboardReleasesDataInner {
	return v.value
}

func (v *NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner) Set(val *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner(val *GameCenterDetailRelationshipsLeaderboardReleasesDataInner) *NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner {
	return &NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner{value: val, isSet: true}
}

func (v NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailRelationshipsLeaderboardReleasesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
