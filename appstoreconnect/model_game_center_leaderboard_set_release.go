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

// checks if the GameCenterLeaderboardSetRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetRelease{}

// GameCenterLeaderboardSetRelease struct for GameCenterLeaderboardSetRelease
type GameCenterLeaderboardSetRelease struct {
	Type          string                                        `json:"type"`
	Id            string                                        `json:"id"`
	Attributes    *GameCenterAchievementReleaseAttributes       `json:"attributes,omitempty"`
	Relationships *GameCenterLeaderboardSetReleaseRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                                `json:"links,omitempty"`
}

// NewGameCenterLeaderboardSetRelease instantiates a new GameCenterLeaderboardSetRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetRelease(type_ string, id string) *GameCenterLeaderboardSetRelease {
	this := GameCenterLeaderboardSetRelease{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardSetReleaseWithDefaults instantiates a new GameCenterLeaderboardSetRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetReleaseWithDefaults() *GameCenterLeaderboardSetRelease {
	this := GameCenterLeaderboardSetRelease{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardSetRelease) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetRelease) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardSetRelease) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardSetRelease) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetRelease) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardSetRelease) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetRelease) GetAttributes() GameCenterAchievementReleaseAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterAchievementReleaseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetRelease) GetAttributesOk() (*GameCenterAchievementReleaseAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetRelease) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterAchievementReleaseAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardSetRelease) SetAttributes(v GameCenterAchievementReleaseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetRelease) GetRelationships() GameCenterLeaderboardSetReleaseRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterLeaderboardSetReleaseRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetRelease) GetRelationshipsOk() (*GameCenterLeaderboardSetReleaseRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetRelease) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterLeaderboardSetReleaseRelationships and assigns it to the Relationships field.
func (o *GameCenterLeaderboardSetRelease) SetRelationships(v GameCenterLeaderboardSetReleaseRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetRelease) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetRelease) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetRelease) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardSetRelease) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterLeaderboardSetRelease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetRelease struct {
	value *GameCenterLeaderboardSetRelease
	isSet bool
}

func (v NullableGameCenterLeaderboardSetRelease) Get() *GameCenterLeaderboardSetRelease {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetRelease) Set(val *GameCenterLeaderboardSetRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetRelease(val *GameCenterLeaderboardSetRelease) *NullableGameCenterLeaderboardSetRelease {
	return &NullableGameCenterLeaderboardSetRelease{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
