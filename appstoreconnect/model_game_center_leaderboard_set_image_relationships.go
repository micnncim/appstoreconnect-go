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

// checks if the GameCenterLeaderboardSetImageRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetImageRelationships{}

// GameCenterLeaderboardSetImageRelationships struct for GameCenterLeaderboardSetImageRelationships
type GameCenterLeaderboardSetImageRelationships struct {
	GameCenterLeaderboardSetLocalization *GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalization `json:"gameCenterLeaderboardSetLocalization,omitempty"`
}

// NewGameCenterLeaderboardSetImageRelationships instantiates a new GameCenterLeaderboardSetImageRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetImageRelationships() *GameCenterLeaderboardSetImageRelationships {
	this := GameCenterLeaderboardSetImageRelationships{}
	return &this
}

// NewGameCenterLeaderboardSetImageRelationshipsWithDefaults instantiates a new GameCenterLeaderboardSetImageRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetImageRelationshipsWithDefaults() *GameCenterLeaderboardSetImageRelationships {
	this := GameCenterLeaderboardSetImageRelationships{}
	return &this
}

// GetGameCenterLeaderboardSetLocalization returns the GameCenterLeaderboardSetLocalization field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetImageRelationships) GetGameCenterLeaderboardSetLocalization() GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalization {
	if o == nil || IsNil(o.GameCenterLeaderboardSetLocalization) {
		var ret GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalization
		return ret
	}
	return *o.GameCenterLeaderboardSetLocalization
}

// GetGameCenterLeaderboardSetLocalizationOk returns a tuple with the GameCenterLeaderboardSetLocalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageRelationships) GetGameCenterLeaderboardSetLocalizationOk() (*GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalization, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboardSetLocalization) {
		return nil, false
	}
	return o.GameCenterLeaderboardSetLocalization, true
}

// HasGameCenterLeaderboardSetLocalization returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetImageRelationships) HasGameCenterLeaderboardSetLocalization() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboardSetLocalization) {
		return true
	}

	return false
}

// SetGameCenterLeaderboardSetLocalization gets a reference to the given GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalization and assigns it to the GameCenterLeaderboardSetLocalization field.
func (o *GameCenterLeaderboardSetImageRelationships) SetGameCenterLeaderboardSetLocalization(v GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalization) {
	o.GameCenterLeaderboardSetLocalization = &v
}

func (o GameCenterLeaderboardSetImageRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetImageRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterLeaderboardSetLocalization) {
		toSerialize["gameCenterLeaderboardSetLocalization"] = o.GameCenterLeaderboardSetLocalization
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetImageRelationships struct {
	value *GameCenterLeaderboardSetImageRelationships
	isSet bool
}

func (v NullableGameCenterLeaderboardSetImageRelationships) Get() *GameCenterLeaderboardSetImageRelationships {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetImageRelationships) Set(val *GameCenterLeaderboardSetImageRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetImageRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetImageRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetImageRelationships(val *GameCenterLeaderboardSetImageRelationships) *NullableGameCenterLeaderboardSetImageRelationships {
	return &NullableGameCenterLeaderboardSetImageRelationships{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetImageRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetImageRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
