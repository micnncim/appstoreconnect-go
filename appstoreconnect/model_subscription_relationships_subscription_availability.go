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

// checks if the SubscriptionRelationshipsSubscriptionAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionRelationshipsSubscriptionAvailability{}

// SubscriptionRelationshipsSubscriptionAvailability struct for SubscriptionRelationshipsSubscriptionAvailability
type SubscriptionRelationshipsSubscriptionAvailability struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data  *SubscriptionRelationshipsSubscriptionAvailabilityData      `json:"data,omitempty"`
}

// NewSubscriptionRelationshipsSubscriptionAvailability instantiates a new SubscriptionRelationshipsSubscriptionAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionRelationshipsSubscriptionAvailability() *SubscriptionRelationshipsSubscriptionAvailability {
	this := SubscriptionRelationshipsSubscriptionAvailability{}
	return &this
}

// NewSubscriptionRelationshipsSubscriptionAvailabilityWithDefaults instantiates a new SubscriptionRelationshipsSubscriptionAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionRelationshipsSubscriptionAvailabilityWithDefaults() *SubscriptionRelationshipsSubscriptionAvailability {
	this := SubscriptionRelationshipsSubscriptionAvailability{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsSubscriptionAvailability) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsSubscriptionAvailability) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsSubscriptionAvailability) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *SubscriptionRelationshipsSubscriptionAvailability) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsSubscriptionAvailability) GetData() SubscriptionRelationshipsSubscriptionAvailabilityData {
	if o == nil || IsNil(o.Data) {
		var ret SubscriptionRelationshipsSubscriptionAvailabilityData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsSubscriptionAvailability) GetDataOk() (*SubscriptionRelationshipsSubscriptionAvailabilityData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsSubscriptionAvailability) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SubscriptionRelationshipsSubscriptionAvailabilityData and assigns it to the Data field.
func (o *SubscriptionRelationshipsSubscriptionAvailability) SetData(v SubscriptionRelationshipsSubscriptionAvailabilityData) {
	o.Data = &v
}

func (o SubscriptionRelationshipsSubscriptionAvailability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionRelationshipsSubscriptionAvailability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSubscriptionRelationshipsSubscriptionAvailability struct {
	value *SubscriptionRelationshipsSubscriptionAvailability
	isSet bool
}

func (v NullableSubscriptionRelationshipsSubscriptionAvailability) Get() *SubscriptionRelationshipsSubscriptionAvailability {
	return v.value
}

func (v *NullableSubscriptionRelationshipsSubscriptionAvailability) Set(val *SubscriptionRelationshipsSubscriptionAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionRelationshipsSubscriptionAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionRelationshipsSubscriptionAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionRelationshipsSubscriptionAvailability(val *SubscriptionRelationshipsSubscriptionAvailability) *NullableSubscriptionRelationshipsSubscriptionAvailability {
	return &NullableSubscriptionRelationshipsSubscriptionAvailability{value: val, isSet: true}
}

func (v NullableSubscriptionRelationshipsSubscriptionAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionRelationshipsSubscriptionAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
