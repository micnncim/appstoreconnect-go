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

// checks if the AppClipAppStoreReviewDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailResponse{}

// AppClipAppStoreReviewDetailResponse struct for AppClipAppStoreReviewDetailResponse
type AppClipAppStoreReviewDetailResponse struct {
	Data     AppClipAppStoreReviewDetail `json:"data"`
	Included []AppClipDefaultExperience  `json:"included,omitempty"`
	Links    DocumentLinks               `json:"links"`
}

// NewAppClipAppStoreReviewDetailResponse instantiates a new AppClipAppStoreReviewDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailResponse(data AppClipAppStoreReviewDetail, links DocumentLinks) *AppClipAppStoreReviewDetailResponse {
	this := AppClipAppStoreReviewDetailResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipAppStoreReviewDetailResponseWithDefaults instantiates a new AppClipAppStoreReviewDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailResponseWithDefaults() *AppClipAppStoreReviewDetailResponse {
	this := AppClipAppStoreReviewDetailResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAppStoreReviewDetailResponse) GetData() AppClipAppStoreReviewDetail {
	if o == nil {
		var ret AppClipAppStoreReviewDetail
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailResponse) GetDataOk() (*AppClipAppStoreReviewDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAppStoreReviewDetailResponse) SetData(v AppClipAppStoreReviewDetail) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppClipAppStoreReviewDetailResponse) GetIncluded() []AppClipDefaultExperience {
	if o == nil || IsNil(o.Included) {
		var ret []AppClipDefaultExperience
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailResponse) GetIncludedOk() ([]AppClipDefaultExperience, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppClipAppStoreReviewDetailResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppClipDefaultExperience and assigns it to the Included field.
func (o *AppClipAppStoreReviewDetailResponse) SetIncluded(v []AppClipDefaultExperience) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppClipAppStoreReviewDetailResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipAppStoreReviewDetailResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppClipAppStoreReviewDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailResponse struct {
	value *AppClipAppStoreReviewDetailResponse
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailResponse) Get() *AppClipAppStoreReviewDetailResponse {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailResponse) Set(val *AppClipAppStoreReviewDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailResponse(val *AppClipAppStoreReviewDetailResponse) *NullableAppClipAppStoreReviewDetailResponse {
	return &NullableAppClipAppStoreReviewDetailResponse{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
