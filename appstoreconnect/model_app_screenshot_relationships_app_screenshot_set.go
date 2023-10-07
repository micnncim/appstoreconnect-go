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

// checks if the AppScreenshotRelationshipsAppScreenshotSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotRelationshipsAppScreenshotSet{}

// AppScreenshotRelationshipsAppScreenshotSet struct for AppScreenshotRelationshipsAppScreenshotSet
type AppScreenshotRelationshipsAppScreenshotSet struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks              `json:"links,omitempty"`
	Data  *AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner `json:"data,omitempty"`
}

// NewAppScreenshotRelationshipsAppScreenshotSet instantiates a new AppScreenshotRelationshipsAppScreenshotSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotRelationshipsAppScreenshotSet() *AppScreenshotRelationshipsAppScreenshotSet {
	this := AppScreenshotRelationshipsAppScreenshotSet{}
	return &this
}

// NewAppScreenshotRelationshipsAppScreenshotSetWithDefaults instantiates a new AppScreenshotRelationshipsAppScreenshotSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotRelationshipsAppScreenshotSetWithDefaults() *AppScreenshotRelationshipsAppScreenshotSet {
	this := AppScreenshotRelationshipsAppScreenshotSet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppScreenshotRelationshipsAppScreenshotSet) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotRelationshipsAppScreenshotSet) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppScreenshotRelationshipsAppScreenshotSet) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppScreenshotRelationshipsAppScreenshotSet) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppScreenshotRelationshipsAppScreenshotSet) GetData() AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotRelationshipsAppScreenshotSet) GetDataOk() (*AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppScreenshotRelationshipsAppScreenshotSet) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner and assigns it to the Data field.
func (o *AppScreenshotRelationshipsAppScreenshotSet) SetData(v AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) {
	o.Data = &v
}

func (o AppScreenshotRelationshipsAppScreenshotSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotRelationshipsAppScreenshotSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppScreenshotRelationshipsAppScreenshotSet struct {
	value *AppScreenshotRelationshipsAppScreenshotSet
	isSet bool
}

func (v NullableAppScreenshotRelationshipsAppScreenshotSet) Get() *AppScreenshotRelationshipsAppScreenshotSet {
	return v.value
}

func (v *NullableAppScreenshotRelationshipsAppScreenshotSet) Set(val *AppScreenshotRelationshipsAppScreenshotSet) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotRelationshipsAppScreenshotSet) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotRelationshipsAppScreenshotSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotRelationshipsAppScreenshotSet(val *AppScreenshotRelationshipsAppScreenshotSet) *NullableAppScreenshotRelationshipsAppScreenshotSet {
	return &NullableAppScreenshotRelationshipsAppScreenshotSet{value: val, isSet: true}
}

func (v NullableAppScreenshotRelationshipsAppScreenshotSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotRelationshipsAppScreenshotSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
