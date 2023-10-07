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

// checks if the GameCenterLeaderboardCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardCreateRequestData{}

// GameCenterLeaderboardCreateRequestData struct for GameCenterLeaderboardCreateRequestData
type GameCenterLeaderboardCreateRequestData struct {
	Type          string                                               `json:"type"`
	Attributes    GameCenterLeaderboardCreateRequestDataAttributes     `json:"attributes"`
	Relationships *GameCenterLeaderboardCreateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewGameCenterLeaderboardCreateRequestData instantiates a new GameCenterLeaderboardCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardCreateRequestData(type_ string, attributes GameCenterLeaderboardCreateRequestDataAttributes) *GameCenterLeaderboardCreateRequestData {
	this := GameCenterLeaderboardCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGameCenterLeaderboardCreateRequestDataWithDefaults instantiates a new GameCenterLeaderboardCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardCreateRequestDataWithDefaults() *GameCenterLeaderboardCreateRequestData {
	this := GameCenterLeaderboardCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterLeaderboardCreateRequestData) GetAttributes() GameCenterLeaderboardCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterLeaderboardCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestData) GetAttributesOk() (*GameCenterLeaderboardCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterLeaderboardCreateRequestData) SetAttributes(v GameCenterLeaderboardCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterLeaderboardCreateRequestData) GetRelationships() GameCenterLeaderboardCreateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterLeaderboardCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestData) GetRelationshipsOk() (*GameCenterLeaderboardCreateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterLeaderboardCreateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterLeaderboardCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *GameCenterLeaderboardCreateRequestData) SetRelationships(v GameCenterLeaderboardCreateRequestDataRelationships) {
	o.Relationships = &v
}

func (o GameCenterLeaderboardCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardCreateRequestData struct {
	value *GameCenterLeaderboardCreateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardCreateRequestData) Get() *GameCenterLeaderboardCreateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardCreateRequestData) Set(val *GameCenterLeaderboardCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardCreateRequestData(val *GameCenterLeaderboardCreateRequestData) *NullableGameCenterLeaderboardCreateRequestData {
	return &NullableGameCenterLeaderboardCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
