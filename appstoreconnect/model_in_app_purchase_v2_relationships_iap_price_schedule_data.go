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

// checks if the InAppPurchaseV2RelationshipsIapPriceScheduleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2RelationshipsIapPriceScheduleData{}

// InAppPurchaseV2RelationshipsIapPriceScheduleData struct for InAppPurchaseV2RelationshipsIapPriceScheduleData
type InAppPurchaseV2RelationshipsIapPriceScheduleData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewInAppPurchaseV2RelationshipsIapPriceScheduleData instantiates a new InAppPurchaseV2RelationshipsIapPriceScheduleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2RelationshipsIapPriceScheduleData(type_ string, id string) *InAppPurchaseV2RelationshipsIapPriceScheduleData {
	this := InAppPurchaseV2RelationshipsIapPriceScheduleData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInAppPurchaseV2RelationshipsIapPriceScheduleDataWithDefaults instantiates a new InAppPurchaseV2RelationshipsIapPriceScheduleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsIapPriceScheduleDataWithDefaults() *InAppPurchaseV2RelationshipsIapPriceScheduleData {
	this := InAppPurchaseV2RelationshipsIapPriceScheduleData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseV2RelationshipsIapPriceScheduleData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsIapPriceScheduleData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseV2RelationshipsIapPriceScheduleData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InAppPurchaseV2RelationshipsIapPriceScheduleData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsIapPriceScheduleData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InAppPurchaseV2RelationshipsIapPriceScheduleData) SetId(v string) {
	o.Id = v
}

func (o InAppPurchaseV2RelationshipsIapPriceScheduleData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2RelationshipsIapPriceScheduleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableInAppPurchaseV2RelationshipsIapPriceScheduleData struct {
	value *InAppPurchaseV2RelationshipsIapPriceScheduleData
	isSet bool
}

func (v NullableInAppPurchaseV2RelationshipsIapPriceScheduleData) Get() *InAppPurchaseV2RelationshipsIapPriceScheduleData {
	return v.value
}

func (v *NullableInAppPurchaseV2RelationshipsIapPriceScheduleData) Set(val *InAppPurchaseV2RelationshipsIapPriceScheduleData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2RelationshipsIapPriceScheduleData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2RelationshipsIapPriceScheduleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2RelationshipsIapPriceScheduleData(val *InAppPurchaseV2RelationshipsIapPriceScheduleData) *NullableInAppPurchaseV2RelationshipsIapPriceScheduleData {
	return &NullableInAppPurchaseV2RelationshipsIapPriceScheduleData{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2RelationshipsIapPriceScheduleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2RelationshipsIapPriceScheduleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
