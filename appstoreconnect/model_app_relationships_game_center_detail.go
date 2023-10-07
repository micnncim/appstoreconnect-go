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

// checks if the AppRelationshipsGameCenterDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsGameCenterDetail{}

// AppRelationshipsGameCenterDetail struct for AppRelationshipsGameCenterDetail
type AppRelationshipsGameCenterDetail struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data  *AppRelationshipsGameCenterDetailData                       `json:"data,omitempty"`
}

// NewAppRelationshipsGameCenterDetail instantiates a new AppRelationshipsGameCenterDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsGameCenterDetail() *AppRelationshipsGameCenterDetail {
	this := AppRelationshipsGameCenterDetail{}
	return &this
}

// NewAppRelationshipsGameCenterDetailWithDefaults instantiates a new AppRelationshipsGameCenterDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsGameCenterDetailWithDefaults() *AppRelationshipsGameCenterDetail {
	this := AppRelationshipsGameCenterDetail{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsGameCenterDetail) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsGameCenterDetail) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsGameCenterDetail) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppRelationshipsGameCenterDetail) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsGameCenterDetail) GetData() AppRelationshipsGameCenterDetailData {
	if o == nil || IsNil(o.Data) {
		var ret AppRelationshipsGameCenterDetailData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsGameCenterDetail) GetDataOk() (*AppRelationshipsGameCenterDetailData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsGameCenterDetail) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppRelationshipsGameCenterDetailData and assigns it to the Data field.
func (o *AppRelationshipsGameCenterDetail) SetData(v AppRelationshipsGameCenterDetailData) {
	o.Data = &v
}

func (o AppRelationshipsGameCenterDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsGameCenterDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsGameCenterDetail struct {
	value *AppRelationshipsGameCenterDetail
	isSet bool
}

func (v NullableAppRelationshipsGameCenterDetail) Get() *AppRelationshipsGameCenterDetail {
	return v.value
}

func (v *NullableAppRelationshipsGameCenterDetail) Set(val *AppRelationshipsGameCenterDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsGameCenterDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsGameCenterDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsGameCenterDetail(val *AppRelationshipsGameCenterDetail) *NullableAppRelationshipsGameCenterDetail {
	return &NullableAppRelationshipsGameCenterDetail{value: val, isSet: true}
}

func (v NullableAppRelationshipsGameCenterDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsGameCenterDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
