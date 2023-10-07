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

// checks if the SubscriptionUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionUpdateRequestData{}

// SubscriptionUpdateRequestData struct for SubscriptionUpdateRequestData
type SubscriptionUpdateRequestData struct {
	Type          string                                      `json:"type"`
	Id            string                                      `json:"id"`
	Attributes    *SubscriptionUpdateRequestDataAttributes    `json:"attributes,omitempty"`
	Relationships *SubscriptionUpdateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewSubscriptionUpdateRequestData instantiates a new SubscriptionUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUpdateRequestData(type_ string, id string) *SubscriptionUpdateRequestData {
	this := SubscriptionUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionUpdateRequestDataWithDefaults instantiates a new SubscriptionUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUpdateRequestDataWithDefaults() *SubscriptionUpdateRequestData {
	this := SubscriptionUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionUpdateRequestData) GetAttributes() SubscriptionUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequestData) GetAttributesOk() (*SubscriptionUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *SubscriptionUpdateRequestData) SetAttributes(v SubscriptionUpdateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionUpdateRequestData) GetRelationships() SubscriptionUpdateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequestData) GetRelationshipsOk() (*SubscriptionUpdateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionUpdateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *SubscriptionUpdateRequestData) SetRelationships(v SubscriptionUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o SubscriptionUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableSubscriptionUpdateRequestData struct {
	value *SubscriptionUpdateRequestData
	isSet bool
}

func (v NullableSubscriptionUpdateRequestData) Get() *SubscriptionUpdateRequestData {
	return v.value
}

func (v *NullableSubscriptionUpdateRequestData) Set(val *SubscriptionUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUpdateRequestData(val *SubscriptionUpdateRequestData) *NullableSubscriptionUpdateRequestData {
	return &NullableSubscriptionUpdateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
