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

// checks if the AppStoreReviewAttachmentCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewAttachmentCreateRequestData{}

// AppStoreReviewAttachmentCreateRequestData struct for AppStoreReviewAttachmentCreateRequestData
type AppStoreReviewAttachmentCreateRequestData struct {
	Type          string                                                    `json:"type"`
	Attributes    AppClipAdvancedExperienceImageCreateRequestDataAttributes `json:"attributes"`
	Relationships AppStoreReviewAttachmentCreateRequestDataRelationships    `json:"relationships"`
}

// NewAppStoreReviewAttachmentCreateRequestData instantiates a new AppStoreReviewAttachmentCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentCreateRequestData(type_ string, attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes, relationships AppStoreReviewAttachmentCreateRequestDataRelationships) *AppStoreReviewAttachmentCreateRequestData {
	this := AppStoreReviewAttachmentCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppStoreReviewAttachmentCreateRequestDataWithDefaults instantiates a new AppStoreReviewAttachmentCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentCreateRequestDataWithDefaults() *AppStoreReviewAttachmentCreateRequestData {
	this := AppStoreReviewAttachmentCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreReviewAttachmentCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreReviewAttachmentCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppStoreReviewAttachmentCreateRequestData) GetAttributes() AppClipAdvancedExperienceImageCreateRequestDataAttributes {
	if o == nil {
		var ret AppClipAdvancedExperienceImageCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppStoreReviewAttachmentCreateRequestData) SetAttributes(v AppClipAdvancedExperienceImageCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreReviewAttachmentCreateRequestData) GetRelationships() AppStoreReviewAttachmentCreateRequestDataRelationships {
	if o == nil {
		var ret AppStoreReviewAttachmentCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentCreateRequestData) GetRelationshipsOk() (*AppStoreReviewAttachmentCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreReviewAttachmentCreateRequestData) SetRelationships(v AppStoreReviewAttachmentCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreReviewAttachmentCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewAttachmentCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppStoreReviewAttachmentCreateRequestData struct {
	value *AppStoreReviewAttachmentCreateRequestData
	isSet bool
}

func (v NullableAppStoreReviewAttachmentCreateRequestData) Get() *AppStoreReviewAttachmentCreateRequestData {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentCreateRequestData) Set(val *AppStoreReviewAttachmentCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentCreateRequestData(val *AppStoreReviewAttachmentCreateRequestData) *NullableAppStoreReviewAttachmentCreateRequestData {
	return &NullableAppStoreReviewAttachmentCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
