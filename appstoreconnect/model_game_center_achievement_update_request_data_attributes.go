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

// checks if the GameCenterAchievementUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementUpdateRequestDataAttributes{}

// GameCenterAchievementUpdateRequestDataAttributes struct for GameCenterAchievementUpdateRequestDataAttributes
type GameCenterAchievementUpdateRequestDataAttributes struct {
	ReferenceName    *string `json:"referenceName,omitempty"`
	Points           *int32  `json:"points,omitempty"`
	ShowBeforeEarned *bool   `json:"showBeforeEarned,omitempty"`
	Repeatable       *bool   `json:"repeatable,omitempty"`
	Archived         *bool   `json:"archived,omitempty"`
}

// NewGameCenterAchievementUpdateRequestDataAttributes instantiates a new GameCenterAchievementUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementUpdateRequestDataAttributes() *GameCenterAchievementUpdateRequestDataAttributes {
	this := GameCenterAchievementUpdateRequestDataAttributes{}
	return &this
}

// NewGameCenterAchievementUpdateRequestDataAttributesWithDefaults instantiates a new GameCenterAchievementUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementUpdateRequestDataAttributesWithDefaults() *GameCenterAchievementUpdateRequestDataAttributes {
	this := GameCenterAchievementUpdateRequestDataAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *GameCenterAchievementUpdateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetPoints() int32 {
	if o == nil || IsNil(o.Points) {
		var ret int32
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given int32 and assigns it to the Points field.
func (o *GameCenterAchievementUpdateRequestDataAttributes) SetPoints(v int32) {
	o.Points = &v
}

// GetShowBeforeEarned returns the ShowBeforeEarned field value if set, zero value otherwise.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetShowBeforeEarned() bool {
	if o == nil || IsNil(o.ShowBeforeEarned) {
		var ret bool
		return ret
	}
	return *o.ShowBeforeEarned
}

// GetShowBeforeEarnedOk returns a tuple with the ShowBeforeEarned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetShowBeforeEarnedOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowBeforeEarned) {
		return nil, false
	}
	return o.ShowBeforeEarned, true
}

// HasShowBeforeEarned returns a boolean if a field has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) HasShowBeforeEarned() bool {
	if o != nil && !IsNil(o.ShowBeforeEarned) {
		return true
	}

	return false
}

// SetShowBeforeEarned gets a reference to the given bool and assigns it to the ShowBeforeEarned field.
func (o *GameCenterAchievementUpdateRequestDataAttributes) SetShowBeforeEarned(v bool) {
	o.ShowBeforeEarned = &v
}

// GetRepeatable returns the Repeatable field value if set, zero value otherwise.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetRepeatable() bool {
	if o == nil || IsNil(o.Repeatable) {
		var ret bool
		return ret
	}
	return *o.Repeatable
}

// GetRepeatableOk returns a tuple with the Repeatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetRepeatableOk() (*bool, bool) {
	if o == nil || IsNil(o.Repeatable) {
		return nil, false
	}
	return o.Repeatable, true
}

// HasRepeatable returns a boolean if a field has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) HasRepeatable() bool {
	if o != nil && !IsNil(o.Repeatable) {
		return true
	}

	return false
}

// SetRepeatable gets a reference to the given bool and assigns it to the Repeatable field.
func (o *GameCenterAchievementUpdateRequestDataAttributes) SetRepeatable(v bool) {
	o.Repeatable = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *GameCenterAchievementUpdateRequestDataAttributes) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *GameCenterAchievementUpdateRequestDataAttributes) SetArchived(v bool) {
	o.Archived = &v
}

func (o GameCenterAchievementUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	if !IsNil(o.ShowBeforeEarned) {
		toSerialize["showBeforeEarned"] = o.ShowBeforeEarned
	}
	if !IsNil(o.Repeatable) {
		toSerialize["repeatable"] = o.Repeatable
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementUpdateRequestDataAttributes struct {
	value *GameCenterAchievementUpdateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterAchievementUpdateRequestDataAttributes) Get() *GameCenterAchievementUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterAchievementUpdateRequestDataAttributes) Set(val *GameCenterAchievementUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementUpdateRequestDataAttributes(val *GameCenterAchievementUpdateRequestDataAttributes) *NullableGameCenterAchievementUpdateRequestDataAttributes {
	return &NullableGameCenterAchievementUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterAchievementUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
