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

// checks if the GameCenterGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupResponse{}

// GameCenterGroupResponse struct for GameCenterGroupResponse
type GameCenterGroupResponse struct {
	Data     GameCenterGroup                         `json:"data"`
	Included []GameCenterGroupsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                           `json:"links"`
}

// NewGameCenterGroupResponse instantiates a new GameCenterGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupResponse(data GameCenterGroup, links DocumentLinks) *GameCenterGroupResponse {
	this := GameCenterGroupResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterGroupResponseWithDefaults instantiates a new GameCenterGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupResponseWithDefaults() *GameCenterGroupResponse {
	this := GameCenterGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterGroupResponse) GetData() GameCenterGroup {
	if o == nil {
		var ret GameCenterGroup
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupResponse) GetDataOk() (*GameCenterGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterGroupResponse) SetData(v GameCenterGroup) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterGroupResponse) GetIncluded() []GameCenterGroupsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterGroupsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroupResponse) GetIncludedOk() ([]GameCenterGroupsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterGroupResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterGroupsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterGroupResponse) SetIncluded(v []GameCenterGroupsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterGroupResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterGroupResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableGameCenterGroupResponse struct {
	value *GameCenterGroupResponse
	isSet bool
}

func (v NullableGameCenterGroupResponse) Get() *GameCenterGroupResponse {
	return v.value
}

func (v *NullableGameCenterGroupResponse) Set(val *GameCenterGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupResponse(val *GameCenterGroupResponse) *NullableGameCenterGroupResponse {
	return &NullableGameCenterGroupResponse{value: val, isSet: true}
}

func (v NullableGameCenterGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
