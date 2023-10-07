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

// checks if the InAppPurchaseV2RelationshipsContentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2RelationshipsContentData{}

// InAppPurchaseV2RelationshipsContentData struct for InAppPurchaseV2RelationshipsContentData
type InAppPurchaseV2RelationshipsContentData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewInAppPurchaseV2RelationshipsContentData instantiates a new InAppPurchaseV2RelationshipsContentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2RelationshipsContentData(type_ string, id string) *InAppPurchaseV2RelationshipsContentData {
	this := InAppPurchaseV2RelationshipsContentData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInAppPurchaseV2RelationshipsContentDataWithDefaults instantiates a new InAppPurchaseV2RelationshipsContentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsContentDataWithDefaults() *InAppPurchaseV2RelationshipsContentData {
	this := InAppPurchaseV2RelationshipsContentData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseV2RelationshipsContentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsContentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseV2RelationshipsContentData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InAppPurchaseV2RelationshipsContentData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsContentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InAppPurchaseV2RelationshipsContentData) SetId(v string) {
	o.Id = v
}

func (o InAppPurchaseV2RelationshipsContentData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2RelationshipsContentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableInAppPurchaseV2RelationshipsContentData struct {
	value *InAppPurchaseV2RelationshipsContentData
	isSet bool
}

func (v NullableInAppPurchaseV2RelationshipsContentData) Get() *InAppPurchaseV2RelationshipsContentData {
	return v.value
}

func (v *NullableInAppPurchaseV2RelationshipsContentData) Set(val *InAppPurchaseV2RelationshipsContentData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2RelationshipsContentData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2RelationshipsContentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2RelationshipsContentData(val *InAppPurchaseV2RelationshipsContentData) *NullableInAppPurchaseV2RelationshipsContentData {
	return &NullableInAppPurchaseV2RelationshipsContentData{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2RelationshipsContentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2RelationshipsContentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
