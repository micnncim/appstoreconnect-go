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

// checks if the GameCenterGroupGameCenterLeaderboardsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupGameCenterLeaderboardsLinkagesRequest{}

// GameCenterGroupGameCenterLeaderboardsLinkagesRequest struct for GameCenterGroupGameCenterLeaderboardsLinkagesRequest
type GameCenterGroupGameCenterLeaderboardsLinkagesRequest struct {
	Data []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner `json:"data"`
}

// NewGameCenterGroupGameCenterLeaderboardsLinkagesRequest instantiates a new GameCenterGroupGameCenterLeaderboardsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupGameCenterLeaderboardsLinkagesRequest(data []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) *GameCenterGroupGameCenterLeaderboardsLinkagesRequest {
	this := GameCenterGroupGameCenterLeaderboardsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewGameCenterGroupGameCenterLeaderboardsLinkagesRequestWithDefaults instantiates a new GameCenterGroupGameCenterLeaderboardsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupGameCenterLeaderboardsLinkagesRequestWithDefaults() *GameCenterGroupGameCenterLeaderboardsLinkagesRequest {
	this := GameCenterGroupGameCenterLeaderboardsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesRequest) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	if o == nil {
		var ret []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesRequest) GetDataOk() ([]GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterGroupGameCenterLeaderboardsLinkagesRequest) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	o.Data = v
}

func (o GameCenterGroupGameCenterLeaderboardsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupGameCenterLeaderboardsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest struct {
	value *GameCenterGroupGameCenterLeaderboardsLinkagesRequest
	isSet bool
}

func (v NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest) Get() *GameCenterGroupGameCenterLeaderboardsLinkagesRequest {
	return v.value
}

func (v *NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest) Set(val *GameCenterGroupGameCenterLeaderboardsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest(val *GameCenterGroupGameCenterLeaderboardsLinkagesRequest) *NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest {
	return &NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest{value: val, isSet: true}
}

func (v NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupGameCenterLeaderboardsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
