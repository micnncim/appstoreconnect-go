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

// checks if the AppEventScreenshotUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshotUpdateRequestData{}

// AppEventScreenshotUpdateRequestData struct for AppEventScreenshotUpdateRequestData
type AppEventScreenshotUpdateRequestData struct {
	Type       string                                         `json:"type"`
	Id         string                                         `json:"id"`
	Attributes *AppEventScreenshotUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewAppEventScreenshotUpdateRequestData instantiates a new AppEventScreenshotUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshotUpdateRequestData(type_ string, id string) *AppEventScreenshotUpdateRequestData {
	this := AppEventScreenshotUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppEventScreenshotUpdateRequestDataWithDefaults instantiates a new AppEventScreenshotUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotUpdateRequestDataWithDefaults() *AppEventScreenshotUpdateRequestData {
	this := AppEventScreenshotUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventScreenshotUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventScreenshotUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEventScreenshotUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventScreenshotUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppEventScreenshotUpdateRequestData) GetAttributes() AppEventScreenshotUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppEventScreenshotUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotUpdateRequestData) GetAttributesOk() (*AppEventScreenshotUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppEventScreenshotUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppEventScreenshotUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppEventScreenshotUpdateRequestData) SetAttributes(v AppEventScreenshotUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppEventScreenshotUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshotUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAppEventScreenshotUpdateRequestData struct {
	value *AppEventScreenshotUpdateRequestData
	isSet bool
}

func (v NullableAppEventScreenshotUpdateRequestData) Get() *AppEventScreenshotUpdateRequestData {
	return v.value
}

func (v *NullableAppEventScreenshotUpdateRequestData) Set(val *AppEventScreenshotUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshotUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshotUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshotUpdateRequestData(val *AppEventScreenshotUpdateRequestData) *NullableAppEventScreenshotUpdateRequestData {
	return &NullableAppEventScreenshotUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppEventScreenshotUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshotUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
