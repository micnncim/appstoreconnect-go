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

// checks if the AppRelationshipsBetaAppReviewDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsBetaAppReviewDetail{}

// AppRelationshipsBetaAppReviewDetail struct for AppRelationshipsBetaAppReviewDetail
type AppRelationshipsBetaAppReviewDetail struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data  *AppRelationshipsBetaAppReviewDetailData                    `json:"data,omitempty"`
}

// NewAppRelationshipsBetaAppReviewDetail instantiates a new AppRelationshipsBetaAppReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsBetaAppReviewDetail() *AppRelationshipsBetaAppReviewDetail {
	this := AppRelationshipsBetaAppReviewDetail{}
	return &this
}

// NewAppRelationshipsBetaAppReviewDetailWithDefaults instantiates a new AppRelationshipsBetaAppReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsBetaAppReviewDetailWithDefaults() *AppRelationshipsBetaAppReviewDetail {
	this := AppRelationshipsBetaAppReviewDetail{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsBetaAppReviewDetail) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppReviewDetail) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsBetaAppReviewDetail) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppRelationshipsBetaAppReviewDetail) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsBetaAppReviewDetail) GetData() AppRelationshipsBetaAppReviewDetailData {
	if o == nil || IsNil(o.Data) {
		var ret AppRelationshipsBetaAppReviewDetailData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppReviewDetail) GetDataOk() (*AppRelationshipsBetaAppReviewDetailData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsBetaAppReviewDetail) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppRelationshipsBetaAppReviewDetailData and assigns it to the Data field.
func (o *AppRelationshipsBetaAppReviewDetail) SetData(v AppRelationshipsBetaAppReviewDetailData) {
	o.Data = &v
}

func (o AppRelationshipsBetaAppReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsBetaAppReviewDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsBetaAppReviewDetail struct {
	value *AppRelationshipsBetaAppReviewDetail
	isSet bool
}

func (v NullableAppRelationshipsBetaAppReviewDetail) Get() *AppRelationshipsBetaAppReviewDetail {
	return v.value
}

func (v *NullableAppRelationshipsBetaAppReviewDetail) Set(val *AppRelationshipsBetaAppReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsBetaAppReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsBetaAppReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsBetaAppReviewDetail(val *AppRelationshipsBetaAppReviewDetail) *NullableAppRelationshipsBetaAppReviewDetail {
	return &NullableAppRelationshipsBetaAppReviewDetail{value: val, isSet: true}
}

func (v NullableAppRelationshipsBetaAppReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsBetaAppReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
