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

// checks if the AppStoreReviewAttachmentUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewAttachmentUpdateRequestData{}

// AppStoreReviewAttachmentUpdateRequestData struct for AppStoreReviewAttachmentUpdateRequestData
type AppStoreReviewAttachmentUpdateRequestData struct {
	Type       string                                                     `json:"type"`
	Id         string                                                     `json:"id"`
	Attributes *AppClipAdvancedExperienceImageUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewAppStoreReviewAttachmentUpdateRequestData instantiates a new AppStoreReviewAttachmentUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentUpdateRequestData(type_ string, id string) *AppStoreReviewAttachmentUpdateRequestData {
	this := AppStoreReviewAttachmentUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreReviewAttachmentUpdateRequestDataWithDefaults instantiates a new AppStoreReviewAttachmentUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentUpdateRequestDataWithDefaults() *AppStoreReviewAttachmentUpdateRequestData {
	this := AppStoreReviewAttachmentUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreReviewAttachmentUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreReviewAttachmentUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreReviewAttachmentUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreReviewAttachmentUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetAttributes() AppClipAdvancedExperienceImageUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceImageUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceImageUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppStoreReviewAttachmentUpdateRequestData) SetAttributes(v AppClipAdvancedExperienceImageUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppStoreReviewAttachmentUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewAttachmentUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAppStoreReviewAttachmentUpdateRequestData struct {
	value *AppStoreReviewAttachmentUpdateRequestData
	isSet bool
}

func (v NullableAppStoreReviewAttachmentUpdateRequestData) Get() *AppStoreReviewAttachmentUpdateRequestData {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentUpdateRequestData) Set(val *AppStoreReviewAttachmentUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentUpdateRequestData(val *AppStoreReviewAttachmentUpdateRequestData) *NullableAppStoreReviewAttachmentUpdateRequestData {
	return &NullableAppStoreReviewAttachmentUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
