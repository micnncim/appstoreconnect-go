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

// checks if the GameCenterLeaderboardLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationCreateRequestData{}

// GameCenterLeaderboardLocalizationCreateRequestData struct for GameCenterLeaderboardLocalizationCreateRequestData
type GameCenterLeaderboardLocalizationCreateRequestData struct {
	Type          string                                                          `json:"type"`
	Attributes    GameCenterLeaderboardLocalizationCreateRequestDataAttributes    `json:"attributes"`
	Relationships GameCenterLeaderboardLocalizationCreateRequestDataRelationships `json:"relationships"`
}

// NewGameCenterLeaderboardLocalizationCreateRequestData instantiates a new GameCenterLeaderboardLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationCreateRequestData(type_ string, attributes GameCenterLeaderboardLocalizationCreateRequestDataAttributes, relationships GameCenterLeaderboardLocalizationCreateRequestDataRelationships) *GameCenterLeaderboardLocalizationCreateRequestData {
	this := GameCenterLeaderboardLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewGameCenterLeaderboardLocalizationCreateRequestDataWithDefaults instantiates a new GameCenterLeaderboardLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationCreateRequestDataWithDefaults() *GameCenterLeaderboardLocalizationCreateRequestData {
	this := GameCenterLeaderboardLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterLeaderboardLocalizationCreateRequestData) GetAttributes() GameCenterLeaderboardLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterLeaderboardLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationCreateRequestData) GetAttributesOk() (*GameCenterLeaderboardLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterLeaderboardLocalizationCreateRequestData) SetAttributes(v GameCenterLeaderboardLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *GameCenterLeaderboardLocalizationCreateRequestData) GetRelationships() GameCenterLeaderboardLocalizationCreateRequestDataRelationships {
	if o == nil {
		var ret GameCenterLeaderboardLocalizationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationCreateRequestData) GetRelationshipsOk() (*GameCenterLeaderboardLocalizationCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GameCenterLeaderboardLocalizationCreateRequestData) SetRelationships(v GameCenterLeaderboardLocalizationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o GameCenterLeaderboardLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalizationCreateRequestData struct {
	value *GameCenterLeaderboardLocalizationCreateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestData) Get() *GameCenterLeaderboardLocalizationCreateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestData) Set(val *GameCenterLeaderboardLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationCreateRequestData(val *GameCenterLeaderboardLocalizationCreateRequestData) *NullableGameCenterLeaderboardLocalizationCreateRequestData {
	return &NullableGameCenterLeaderboardLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
