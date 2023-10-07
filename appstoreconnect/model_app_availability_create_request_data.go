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

// checks if the AppAvailabilityCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityCreateRequestData{}

// AppAvailabilityCreateRequestData struct for AppAvailabilityCreateRequestData
type AppAvailabilityCreateRequestData struct {
	Type          string                                        `json:"type"`
	Attributes    AppAvailabilityV2CreateRequestDataAttributes  `json:"attributes"`
	Relationships AppAvailabilityCreateRequestDataRelationships `json:"relationships"`
}

// NewAppAvailabilityCreateRequestData instantiates a new AppAvailabilityCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityCreateRequestData(type_ string, attributes AppAvailabilityV2CreateRequestDataAttributes, relationships AppAvailabilityCreateRequestDataRelationships) *AppAvailabilityCreateRequestData {
	this := AppAvailabilityCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppAvailabilityCreateRequestDataWithDefaults instantiates a new AppAvailabilityCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityCreateRequestDataWithDefaults() *AppAvailabilityCreateRequestData {
	this := AppAvailabilityCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppAvailabilityCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppAvailabilityCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppAvailabilityCreateRequestData) GetAttributes() AppAvailabilityV2CreateRequestDataAttributes {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityCreateRequestData) GetAttributesOk() (*AppAvailabilityV2CreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppAvailabilityCreateRequestData) SetAttributes(v AppAvailabilityV2CreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppAvailabilityCreateRequestData) GetRelationships() AppAvailabilityCreateRequestDataRelationships {
	if o == nil {
		var ret AppAvailabilityCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityCreateRequestData) GetRelationshipsOk() (*AppAvailabilityCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppAvailabilityCreateRequestData) SetRelationships(v AppAvailabilityCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppAvailabilityCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppAvailabilityCreateRequestData struct {
	value *AppAvailabilityCreateRequestData
	isSet bool
}

func (v NullableAppAvailabilityCreateRequestData) Get() *AppAvailabilityCreateRequestData {
	return v.value
}

func (v *NullableAppAvailabilityCreateRequestData) Set(val *AppAvailabilityCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityCreateRequestData(val *AppAvailabilityCreateRequestData) *NullableAppAvailabilityCreateRequestData {
	return &NullableAppAvailabilityCreateRequestData{value: val, isSet: true}
}

func (v NullableAppAvailabilityCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
