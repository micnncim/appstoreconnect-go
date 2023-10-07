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

// checks if the GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData{}

// GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData struct for GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData
type GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData instantiates a new GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData(type_ string, id string) *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData {
	this := GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageDataWithDefaults instantiates a new GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageDataWithDefaults() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData {
	this := GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) SetId(v string) {
	o.Id = v
}

func (o GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData struct {
	value *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) Get() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) Set(val *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData(val *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData {
	return &NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
