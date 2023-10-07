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

// checks if the AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization{}

// AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization struct for AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization
type AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks                          `json:"links,omitempty"`
	Data  *AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner `json:"data,omitempty"`
}

// NewAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization instantiates a new AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization() *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization {
	this := AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization{}
	return &this
}

// NewAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalizationWithDefaults instantiates a new AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalizationWithDefaults() *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization {
	this := AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) GetData() AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) GetDataOk() (*AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner and assigns it to the Data field.
func (o *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) SetData(v AppClipDefaultExperienceRelationshipsAppClipDefaultExperienceLocalizationsDataInner) {
	o.Data = &v
}

func (o AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization struct {
	value *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization
	isSet bool
}

func (v NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) Get() *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization {
	return v.value
}

func (v *NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) Set(val *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization(val *AppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) *NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization {
	return &NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization{value: val, isSet: true}
}

func (v NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipHeaderImageRelationshipsAppClipDefaultExperienceLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
