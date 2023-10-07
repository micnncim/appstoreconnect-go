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

// checks if the AppInfoLocalizationUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoLocalizationUpdateRequestData{}

// AppInfoLocalizationUpdateRequestData struct for AppInfoLocalizationUpdateRequestData
type AppInfoLocalizationUpdateRequestData struct {
	Type       string                                          `json:"type"`
	Id         string                                          `json:"id"`
	Attributes *AppInfoLocalizationUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewAppInfoLocalizationUpdateRequestData instantiates a new AppInfoLocalizationUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationUpdateRequestData(type_ string, id string) *AppInfoLocalizationUpdateRequestData {
	this := AppInfoLocalizationUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppInfoLocalizationUpdateRequestDataWithDefaults instantiates a new AppInfoLocalizationUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationUpdateRequestDataWithDefaults() *AppInfoLocalizationUpdateRequestData {
	this := AppInfoLocalizationUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppInfoLocalizationUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppInfoLocalizationUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppInfoLocalizationUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppInfoLocalizationUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestData) GetAttributes() AppInfoLocalizationUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppInfoLocalizationUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestData) GetAttributesOk() (*AppInfoLocalizationUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppInfoLocalizationUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppInfoLocalizationUpdateRequestData) SetAttributes(v AppInfoLocalizationUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppInfoLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoLocalizationUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAppInfoLocalizationUpdateRequestData struct {
	value *AppInfoLocalizationUpdateRequestData
	isSet bool
}

func (v NullableAppInfoLocalizationUpdateRequestData) Get() *AppInfoLocalizationUpdateRequestData {
	return v.value
}

func (v *NullableAppInfoLocalizationUpdateRequestData) Set(val *AppInfoLocalizationUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationUpdateRequestData(val *AppInfoLocalizationUpdateRequestData) *NullableAppInfoLocalizationUpdateRequestData {
	return &NullableAppInfoLocalizationUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
