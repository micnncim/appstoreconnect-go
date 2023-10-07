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

// checks if the GameCenterGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroup{}

// GameCenterGroup struct for GameCenterGroup
type GameCenterGroup struct {
	Type          string                        `json:"type"`
	Id            string                        `json:"id"`
	Attributes    *GameCenterGroupAttributes    `json:"attributes,omitempty"`
	Relationships *GameCenterGroupRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                `json:"links,omitempty"`
}

// NewGameCenterGroup instantiates a new GameCenterGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroup(type_ string, id string) *GameCenterGroup {
	this := GameCenterGroup{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterGroupWithDefaults instantiates a new GameCenterGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupWithDefaults() *GameCenterGroup {
	this := GameCenterGroup{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterGroup) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroup) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterGroup) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterGroup) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterGroup) GetAttributes() GameCenterGroupAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterGroupAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroup) GetAttributesOk() (*GameCenterGroupAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterGroup) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterGroupAttributes and assigns it to the Attributes field.
func (o *GameCenterGroup) SetAttributes(v GameCenterGroupAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterGroup) GetRelationships() GameCenterGroupRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterGroupRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroup) GetRelationshipsOk() (*GameCenterGroupRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterGroup) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterGroupRelationships and assigns it to the Relationships field.
func (o *GameCenterGroup) SetRelationships(v GameCenterGroupRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterGroup) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroup) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterGroup) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroup) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterGroup struct {
	value *GameCenterGroup
	isSet bool
}

func (v NullableGameCenterGroup) Get() *GameCenterGroup {
	return v.value
}

func (v *NullableGameCenterGroup) Set(val *GameCenterGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroup(val *GameCenterGroup) *NullableGameCenterGroup {
	return &NullableGameCenterGroup{value: val, isSet: true}
}

func (v NullableGameCenterGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
