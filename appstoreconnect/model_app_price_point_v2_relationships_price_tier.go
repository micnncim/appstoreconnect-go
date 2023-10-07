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

// checks if the AppPricePointV2RelationshipsPriceTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPricePointV2RelationshipsPriceTier{}

// AppPricePointV2RelationshipsPriceTier struct for AppPricePointV2RelationshipsPriceTier
type AppPricePointV2RelationshipsPriceTier struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data  *AppPricePointV2RelationshipsPriceTierData                  `json:"data,omitempty"`
}

// NewAppPricePointV2RelationshipsPriceTier instantiates a new AppPricePointV2RelationshipsPriceTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointV2RelationshipsPriceTier() *AppPricePointV2RelationshipsPriceTier {
	this := AppPricePointV2RelationshipsPriceTier{}
	return &this
}

// NewAppPricePointV2RelationshipsPriceTierWithDefaults instantiates a new AppPricePointV2RelationshipsPriceTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointV2RelationshipsPriceTierWithDefaults() *AppPricePointV2RelationshipsPriceTier {
	this := AppPricePointV2RelationshipsPriceTier{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPricePointV2RelationshipsPriceTier) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointV2RelationshipsPriceTier) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPricePointV2RelationshipsPriceTier) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppPricePointV2RelationshipsPriceTier) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPricePointV2RelationshipsPriceTier) GetData() AppPricePointV2RelationshipsPriceTierData {
	if o == nil || IsNil(o.Data) {
		var ret AppPricePointV2RelationshipsPriceTierData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointV2RelationshipsPriceTier) GetDataOk() (*AppPricePointV2RelationshipsPriceTierData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPricePointV2RelationshipsPriceTier) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppPricePointV2RelationshipsPriceTierData and assigns it to the Data field.
func (o *AppPricePointV2RelationshipsPriceTier) SetData(v AppPricePointV2RelationshipsPriceTierData) {
	o.Data = &v
}

func (o AppPricePointV2RelationshipsPriceTier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPricePointV2RelationshipsPriceTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppPricePointV2RelationshipsPriceTier struct {
	value *AppPricePointV2RelationshipsPriceTier
	isSet bool
}

func (v NullableAppPricePointV2RelationshipsPriceTier) Get() *AppPricePointV2RelationshipsPriceTier {
	return v.value
}

func (v *NullableAppPricePointV2RelationshipsPriceTier) Set(val *AppPricePointV2RelationshipsPriceTier) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointV2RelationshipsPriceTier) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointV2RelationshipsPriceTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointV2RelationshipsPriceTier(val *AppPricePointV2RelationshipsPriceTier) *NullableAppPricePointV2RelationshipsPriceTier {
	return &NullableAppPricePointV2RelationshipsPriceTier{value: val, isSet: true}
}

func (v NullableAppPricePointV2RelationshipsPriceTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointV2RelationshipsPriceTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
