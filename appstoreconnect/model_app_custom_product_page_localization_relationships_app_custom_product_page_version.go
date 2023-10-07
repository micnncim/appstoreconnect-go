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

// checks if the AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion{}

// AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion struct for AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion
type AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks                   `json:"links,omitempty"`
	Data  *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData `json:"data,omitempty"`
}

// NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion instantiates a new AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion() *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion {
	this := AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion{}
	return &this
}

// NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionWithDefaults instantiates a new AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionWithDefaults() *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion {
	this := AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) GetData() AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData {
	if o == nil || IsNil(o.Data) {
		var ret AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) GetDataOk() (*AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData and assigns it to the Data field.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) SetData(v AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) {
	o.Data = &v
}

func (o AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion struct {
	value *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) Get() *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) Set(val *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion(val *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion {
	return &NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
