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

// checks if the AppScreenshotSetCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotSetCreateRequestData{}

// AppScreenshotSetCreateRequestData struct for AppScreenshotSetCreateRequestData
type AppScreenshotSetCreateRequestData struct {
	Type          string                                       `json:"type"`
	Attributes    AppScreenshotSetCreateRequestDataAttributes  `json:"attributes"`
	Relationships *AppPreviewSetCreateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewAppScreenshotSetCreateRequestData instantiates a new AppScreenshotSetCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetCreateRequestData(type_ string, attributes AppScreenshotSetCreateRequestDataAttributes) *AppScreenshotSetCreateRequestData {
	this := AppScreenshotSetCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAppScreenshotSetCreateRequestDataWithDefaults instantiates a new AppScreenshotSetCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetCreateRequestDataWithDefaults() *AppScreenshotSetCreateRequestData {
	this := AppScreenshotSetCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppScreenshotSetCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppScreenshotSetCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppScreenshotSetCreateRequestData) GetAttributes() AppScreenshotSetCreateRequestDataAttributes {
	if o == nil {
		var ret AppScreenshotSetCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequestData) GetAttributesOk() (*AppScreenshotSetCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppScreenshotSetCreateRequestData) SetAttributes(v AppScreenshotSetCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppScreenshotSetCreateRequestData) GetRelationships() AppPreviewSetCreateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppPreviewSetCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequestData) GetRelationshipsOk() (*AppPreviewSetCreateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppScreenshotSetCreateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppPreviewSetCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *AppScreenshotSetCreateRequestData) SetRelationships(v AppPreviewSetCreateRequestDataRelationships) {
	o.Relationships = &v
}

func (o AppScreenshotSetCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotSetCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAppScreenshotSetCreateRequestData struct {
	value *AppScreenshotSetCreateRequestData
	isSet bool
}

func (v NullableAppScreenshotSetCreateRequestData) Get() *AppScreenshotSetCreateRequestData {
	return v.value
}

func (v *NullableAppScreenshotSetCreateRequestData) Set(val *AppScreenshotSetCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetCreateRequestData(val *AppScreenshotSetCreateRequestData) *NullableAppScreenshotSetCreateRequestData {
	return &NullableAppScreenshotSetCreateRequestData{value: val, isSet: true}
}

func (v NullableAppScreenshotSetCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
