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

// checks if the SubscriptionSubmissionCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionSubmissionCreateRequestData{}

// SubscriptionSubmissionCreateRequestData struct for SubscriptionSubmissionCreateRequestData
type SubscriptionSubmissionCreateRequestData struct {
	Type          string                                                             `json:"type"`
	Relationships SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships `json:"relationships"`
}

// NewSubscriptionSubmissionCreateRequestData instantiates a new SubscriptionSubmissionCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionSubmissionCreateRequestData(type_ string, relationships SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships) *SubscriptionSubmissionCreateRequestData {
	this := SubscriptionSubmissionCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewSubscriptionSubmissionCreateRequestDataWithDefaults instantiates a new SubscriptionSubmissionCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionSubmissionCreateRequestDataWithDefaults() *SubscriptionSubmissionCreateRequestData {
	this := SubscriptionSubmissionCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionSubmissionCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSubmissionCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionSubmissionCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionSubmissionCreateRequestData) GetRelationships() SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSubmissionCreateRequestData) GetRelationshipsOk() (*SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionSubmissionCreateRequestData) SetRelationships(v SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionSubmissionCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableSubscriptionSubmissionCreateRequestData struct {
	value *SubscriptionSubmissionCreateRequestData
	isSet bool
}

func (v NullableSubscriptionSubmissionCreateRequestData) Get() *SubscriptionSubmissionCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionSubmissionCreateRequestData) Set(val *SubscriptionSubmissionCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionSubmissionCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionSubmissionCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionSubmissionCreateRequestData(val *SubscriptionSubmissionCreateRequestData) *NullableSubscriptionSubmissionCreateRequestData {
	return &NullableSubscriptionSubmissionCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionSubmissionCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
