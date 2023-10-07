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

// checks if the BundleIdCapabilityCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdCapabilityCreateRequestData{}

// BundleIdCapabilityCreateRequestData struct for BundleIdCapabilityCreateRequestData
type BundleIdCapabilityCreateRequestData struct {
	Type          string                                           `json:"type"`
	Attributes    BundleIdCapabilityCreateRequestDataAttributes    `json:"attributes"`
	Relationships BundleIdCapabilityCreateRequestDataRelationships `json:"relationships"`
}

// NewBundleIdCapabilityCreateRequestData instantiates a new BundleIdCapabilityCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityCreateRequestData(type_ string, attributes BundleIdCapabilityCreateRequestDataAttributes, relationships BundleIdCapabilityCreateRequestDataRelationships) *BundleIdCapabilityCreateRequestData {
	this := BundleIdCapabilityCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewBundleIdCapabilityCreateRequestDataWithDefaults instantiates a new BundleIdCapabilityCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityCreateRequestDataWithDefaults() *BundleIdCapabilityCreateRequestData {
	this := BundleIdCapabilityCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BundleIdCapabilityCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleIdCapabilityCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BundleIdCapabilityCreateRequestData) GetAttributes() BundleIdCapabilityCreateRequestDataAttributes {
	if o == nil {
		var ret BundleIdCapabilityCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestData) GetAttributesOk() (*BundleIdCapabilityCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BundleIdCapabilityCreateRequestData) SetAttributes(v BundleIdCapabilityCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *BundleIdCapabilityCreateRequestData) GetRelationships() BundleIdCapabilityCreateRequestDataRelationships {
	if o == nil {
		var ret BundleIdCapabilityCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestData) GetRelationshipsOk() (*BundleIdCapabilityCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *BundleIdCapabilityCreateRequestData) SetRelationships(v BundleIdCapabilityCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o BundleIdCapabilityCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdCapabilityCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableBundleIdCapabilityCreateRequestData struct {
	value *BundleIdCapabilityCreateRequestData
	isSet bool
}

func (v NullableBundleIdCapabilityCreateRequestData) Get() *BundleIdCapabilityCreateRequestData {
	return v.value
}

func (v *NullableBundleIdCapabilityCreateRequestData) Set(val *BundleIdCapabilityCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityCreateRequestData(val *BundleIdCapabilityCreateRequestData) *NullableBundleIdCapabilityCreateRequestData {
	return &NullableBundleIdCapabilityCreateRequestData{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
