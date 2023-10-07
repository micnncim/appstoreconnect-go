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

// checks if the GameCenterAppVersionRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionRelationships{}

// GameCenterAppVersionRelationships struct for GameCenterAppVersionRelationships
type GameCenterAppVersionRelationships struct {
	CompatibilityVersions *GameCenterAppVersionRelationshipsCompatibilityVersions          `json:"compatibilityVersions,omitempty"`
	AppStoreVersion       *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersion,omitempty"`
}

// NewGameCenterAppVersionRelationships instantiates a new GameCenterAppVersionRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionRelationships() *GameCenterAppVersionRelationships {
	this := GameCenterAppVersionRelationships{}
	return &this
}

// NewGameCenterAppVersionRelationshipsWithDefaults instantiates a new GameCenterAppVersionRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionRelationshipsWithDefaults() *GameCenterAppVersionRelationships {
	this := GameCenterAppVersionRelationships{}
	return &this
}

// GetCompatibilityVersions returns the CompatibilityVersions field value if set, zero value otherwise.
func (o *GameCenterAppVersionRelationships) GetCompatibilityVersions() GameCenterAppVersionRelationshipsCompatibilityVersions {
	if o == nil || IsNil(o.CompatibilityVersions) {
		var ret GameCenterAppVersionRelationshipsCompatibilityVersions
		return ret
	}
	return *o.CompatibilityVersions
}

// GetCompatibilityVersionsOk returns a tuple with the CompatibilityVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionRelationships) GetCompatibilityVersionsOk() (*GameCenterAppVersionRelationshipsCompatibilityVersions, bool) {
	if o == nil || IsNil(o.CompatibilityVersions) {
		return nil, false
	}
	return o.CompatibilityVersions, true
}

// HasCompatibilityVersions returns a boolean if a field has been set.
func (o *GameCenterAppVersionRelationships) HasCompatibilityVersions() bool {
	if o != nil && !IsNil(o.CompatibilityVersions) {
		return true
	}

	return false
}

// SetCompatibilityVersions gets a reference to the given GameCenterAppVersionRelationshipsCompatibilityVersions and assigns it to the CompatibilityVersions field.
func (o *GameCenterAppVersionRelationships) SetCompatibilityVersions(v GameCenterAppVersionRelationshipsCompatibilityVersions) {
	o.CompatibilityVersions = &v
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *GameCenterAppVersionRelationships) GetAppStoreVersion() AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersion) {
		var ret AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersion) {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *GameCenterAppVersionRelationships) HasAppStoreVersion() bool {
	if o != nil && !IsNil(o.AppStoreVersion) {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *GameCenterAppVersionRelationships) SetAppStoreVersion(v AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersion = &v
}

func (o GameCenterAppVersionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompatibilityVersions) {
		toSerialize["compatibilityVersions"] = o.CompatibilityVersions
	}
	if !IsNil(o.AppStoreVersion) {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}
	return toSerialize, nil
}

type NullableGameCenterAppVersionRelationships struct {
	value *GameCenterAppVersionRelationships
	isSet bool
}

func (v NullableGameCenterAppVersionRelationships) Get() *GameCenterAppVersionRelationships {
	return v.value
}

func (v *NullableGameCenterAppVersionRelationships) Set(val *GameCenterAppVersionRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionRelationships(val *GameCenterAppVersionRelationships) *NullableGameCenterAppVersionRelationships {
	return &NullableGameCenterAppVersionRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
