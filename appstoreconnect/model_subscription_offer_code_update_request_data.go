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

// checks if the SubscriptionOfferCodeUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeUpdateRequestData{}

// SubscriptionOfferCodeUpdateRequestData struct for SubscriptionOfferCodeUpdateRequestData
type SubscriptionOfferCodeUpdateRequestData struct {
	Type       string                                                      `json:"type"`
	Id         string                                                      `json:"id"`
	Attributes *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewSubscriptionOfferCodeUpdateRequestData instantiates a new SubscriptionOfferCodeUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeUpdateRequestData(type_ string, id string) *SubscriptionOfferCodeUpdateRequestData {
	this := SubscriptionOfferCodeUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionOfferCodeUpdateRequestDataWithDefaults instantiates a new SubscriptionOfferCodeUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeUpdateRequestDataWithDefaults() *SubscriptionOfferCodeUpdateRequestData {
	this := SubscriptionOfferCodeUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionOfferCodeUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionOfferCodeUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeUpdateRequestData) GetAttributes() SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeUpdateRequestData) GetAttributesOk() (*SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *SubscriptionOfferCodeUpdateRequestData) SetAttributes(v SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o SubscriptionOfferCodeUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeUpdateRequestData struct {
	value *SubscriptionOfferCodeUpdateRequestData
	isSet bool
}

func (v NullableSubscriptionOfferCodeUpdateRequestData) Get() *SubscriptionOfferCodeUpdateRequestData {
	return v.value
}

func (v *NullableSubscriptionOfferCodeUpdateRequestData) Set(val *SubscriptionOfferCodeUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeUpdateRequestData(val *SubscriptionOfferCodeUpdateRequestData) *NullableSubscriptionOfferCodeUpdateRequestData {
	return &NullableSubscriptionOfferCodeUpdateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
