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

// checks if the SubscriptionRelationshipsIntroductoryOffers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionRelationshipsIntroductoryOffers{}

// SubscriptionRelationshipsIntroductoryOffers struct for SubscriptionRelationshipsIntroductoryOffers
type SubscriptionRelationshipsIntroductoryOffers struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Meta  *PagingInformation                                          `json:"meta,omitempty"`
	Data  []SubscriptionRelationshipsIntroductoryOffersDataInner      `json:"data,omitempty"`
}

// NewSubscriptionRelationshipsIntroductoryOffers instantiates a new SubscriptionRelationshipsIntroductoryOffers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionRelationshipsIntroductoryOffers() *SubscriptionRelationshipsIntroductoryOffers {
	this := SubscriptionRelationshipsIntroductoryOffers{}
	return &this
}

// NewSubscriptionRelationshipsIntroductoryOffersWithDefaults instantiates a new SubscriptionRelationshipsIntroductoryOffers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionRelationshipsIntroductoryOffersWithDefaults() *SubscriptionRelationshipsIntroductoryOffers {
	this := SubscriptionRelationshipsIntroductoryOffers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsIntroductoryOffers) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsIntroductoryOffers) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsIntroductoryOffers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *SubscriptionRelationshipsIntroductoryOffers) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsIntroductoryOffers) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsIntroductoryOffers) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsIntroductoryOffers) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionRelationshipsIntroductoryOffers) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsIntroductoryOffers) GetData() []SubscriptionRelationshipsIntroductoryOffersDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []SubscriptionRelationshipsIntroductoryOffersDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsIntroductoryOffers) GetDataOk() ([]SubscriptionRelationshipsIntroductoryOffersDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsIntroductoryOffers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SubscriptionRelationshipsIntroductoryOffersDataInner and assigns it to the Data field.
func (o *SubscriptionRelationshipsIntroductoryOffers) SetData(v []SubscriptionRelationshipsIntroductoryOffersDataInner) {
	o.Data = v
}

func (o SubscriptionRelationshipsIntroductoryOffers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionRelationshipsIntroductoryOffers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSubscriptionRelationshipsIntroductoryOffers struct {
	value *SubscriptionRelationshipsIntroductoryOffers
	isSet bool
}

func (v NullableSubscriptionRelationshipsIntroductoryOffers) Get() *SubscriptionRelationshipsIntroductoryOffers {
	return v.value
}

func (v *NullableSubscriptionRelationshipsIntroductoryOffers) Set(val *SubscriptionRelationshipsIntroductoryOffers) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionRelationshipsIntroductoryOffers) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionRelationshipsIntroductoryOffers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionRelationshipsIntroductoryOffers(val *SubscriptionRelationshipsIntroductoryOffers) *NullableSubscriptionRelationshipsIntroductoryOffers {
	return &NullableSubscriptionRelationshipsIntroductoryOffers{value: val, isSet: true}
}

func (v NullableSubscriptionRelationshipsIntroductoryOffers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionRelationshipsIntroductoryOffers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
