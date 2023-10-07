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

// checks if the AppPreviewSetRelationshipsAppStoreVersionLocalizationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetRelationshipsAppStoreVersionLocalizationData{}

// AppPreviewSetRelationshipsAppStoreVersionLocalizationData struct for AppPreviewSetRelationshipsAppStoreVersionLocalizationData
type AppPreviewSetRelationshipsAppStoreVersionLocalizationData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewAppPreviewSetRelationshipsAppStoreVersionLocalizationData instantiates a new AppPreviewSetRelationshipsAppStoreVersionLocalizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetRelationshipsAppStoreVersionLocalizationData(type_ string, id string) *AppPreviewSetRelationshipsAppStoreVersionLocalizationData {
	this := AppPreviewSetRelationshipsAppStoreVersionLocalizationData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppPreviewSetRelationshipsAppStoreVersionLocalizationDataWithDefaults instantiates a new AppPreviewSetRelationshipsAppStoreVersionLocalizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetRelationshipsAppStoreVersionLocalizationDataWithDefaults() *AppPreviewSetRelationshipsAppStoreVersionLocalizationData {
	this := AppPreviewSetRelationshipsAppStoreVersionLocalizationData{}
	return &this
}

// GetType returns the Type field value
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) SetId(v string) {
	o.Id = v
}

func (o AppPreviewSetRelationshipsAppStoreVersionLocalizationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetRelationshipsAppStoreVersionLocalizationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData struct {
	value *AppPreviewSetRelationshipsAppStoreVersionLocalizationData
	isSet bool
}

func (v NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData) Get() *AppPreviewSetRelationshipsAppStoreVersionLocalizationData {
	return v.value
}

func (v *NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData) Set(val *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData(val *AppPreviewSetRelationshipsAppStoreVersionLocalizationData) *NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData {
	return &NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData{value: val, isSet: true}
}

func (v NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetRelationshipsAppStoreVersionLocalizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
