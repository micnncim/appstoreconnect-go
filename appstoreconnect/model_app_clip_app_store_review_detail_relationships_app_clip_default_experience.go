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

// checks if the AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience{}

// AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience struct for AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience
type AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks           `json:"links,omitempty"`
	Data  *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData `json:"data,omitempty"`
}

// NewAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience instantiates a new AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience() *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience {
	this := AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience{}
	return &this
}

// NewAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceWithDefaults instantiates a new AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceWithDefaults() *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience {
	this := AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) GetData() AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData {
	if o == nil || IsNil(o.Data) {
		var ret AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) GetDataOk() (*AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData and assigns it to the Data field.
func (o *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) SetData(v AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData) {
	o.Data = &v
}

func (o AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience struct {
	value *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) Get() *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) Set(val *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience(val *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) *NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience {
	return &NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
