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

// checks if the SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData{}

// SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData struct for SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData
type SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData instantiates a new SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData(type_ string, id string) *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData {
	this := SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeDataWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeDataWithDefaults() *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData {
	this := SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) SetId(v string) {
	o.Id = v
}

func (o SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData struct {
	value *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) Get() *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) Set(val *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData(val *SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) *NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData {
	return &NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
