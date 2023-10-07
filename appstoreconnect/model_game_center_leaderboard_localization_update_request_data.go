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

// checks if the GameCenterLeaderboardLocalizationUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationUpdateRequestData{}

// GameCenterLeaderboardLocalizationUpdateRequestData struct for GameCenterLeaderboardLocalizationUpdateRequestData
type GameCenterLeaderboardLocalizationUpdateRequestData struct {
	Type       string                                                        `json:"type"`
	Id         string                                                        `json:"id"`
	Attributes *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewGameCenterLeaderboardLocalizationUpdateRequestData instantiates a new GameCenterLeaderboardLocalizationUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationUpdateRequestData(type_ string, id string) *GameCenterLeaderboardLocalizationUpdateRequestData {
	this := GameCenterLeaderboardLocalizationUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardLocalizationUpdateRequestDataWithDefaults instantiates a new GameCenterLeaderboardLocalizationUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationUpdateRequestDataWithDefaults() *GameCenterLeaderboardLocalizationUpdateRequestData {
	this := GameCenterLeaderboardLocalizationUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) GetAttributes() GameCenterLeaderboardLocalizationUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterLeaderboardLocalizationUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) GetAttributesOk() (*GameCenterLeaderboardLocalizationUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterLeaderboardLocalizationUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardLocalizationUpdateRequestData) SetAttributes(v GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o GameCenterLeaderboardLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalizationUpdateRequestData struct {
	value *GameCenterLeaderboardLocalizationUpdateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequestData) Get() *GameCenterLeaderboardLocalizationUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequestData) Set(val *GameCenterLeaderboardLocalizationUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationUpdateRequestData(val *GameCenterLeaderboardLocalizationUpdateRequestData) *NullableGameCenterLeaderboardLocalizationUpdateRequestData {
	return &NullableGameCenterLeaderboardLocalizationUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
