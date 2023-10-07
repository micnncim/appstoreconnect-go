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

// checks if the InAppPurchasePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePrice{}

// InAppPurchasePrice struct for InAppPurchasePrice
type InAppPurchasePrice struct {
	Type          string                           `json:"type"`
	Id            string                           `json:"id"`
	Attributes    *InAppPurchasePriceAttributes    `json:"attributes,omitempty"`
	Relationships *InAppPurchasePriceRelationships `json:"relationships,omitempty"`
}

// NewInAppPurchasePrice instantiates a new InAppPurchasePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePrice(type_ string, id string) *InAppPurchasePrice {
	this := InAppPurchasePrice{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInAppPurchasePriceWithDefaults instantiates a new InAppPurchasePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceWithDefaults() *InAppPurchasePrice {
	this := InAppPurchasePrice{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchasePrice) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePrice) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchasePrice) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InAppPurchasePrice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePrice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InAppPurchasePrice) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InAppPurchasePrice) GetAttributes() InAppPurchasePriceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret InAppPurchasePriceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePrice) GetAttributesOk() (*InAppPurchasePriceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InAppPurchasePrice) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given InAppPurchasePriceAttributes and assigns it to the Attributes field.
func (o *InAppPurchasePrice) SetAttributes(v InAppPurchasePriceAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InAppPurchasePrice) GetRelationships() InAppPurchasePriceRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InAppPurchasePriceRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePrice) GetRelationshipsOk() (*InAppPurchasePriceRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InAppPurchasePrice) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InAppPurchasePriceRelationships and assigns it to the Relationships field.
func (o *InAppPurchasePrice) SetRelationships(v InAppPurchasePriceRelationships) {
	o.Relationships = &v
}

func (o InAppPurchasePrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePrice) ToMap() (map[string]interface{}, error) {
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

type NullableInAppPurchasePrice struct {
	value *InAppPurchasePrice
	isSet bool
}

func (v NullableInAppPurchasePrice) Get() *InAppPurchasePrice {
	return v.value
}

func (v *NullableInAppPurchasePrice) Set(val *InAppPurchasePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePrice(val *InAppPurchasePrice) *NullableInAppPurchasePrice {
	return &NullableInAppPurchasePrice{value: val, isSet: true}
}

func (v NullableInAppPurchasePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
