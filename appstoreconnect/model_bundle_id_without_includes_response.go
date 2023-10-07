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

// checks if the BundleIdWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdWithoutIncludesResponse{}

// BundleIdWithoutIncludesResponse struct for BundleIdWithoutIncludesResponse
type BundleIdWithoutIncludesResponse struct {
	Data  Profile       `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewBundleIdWithoutIncludesResponse instantiates a new BundleIdWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdWithoutIncludesResponse(data Profile, links DocumentLinks) *BundleIdWithoutIncludesResponse {
	this := BundleIdWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBundleIdWithoutIncludesResponseWithDefaults instantiates a new BundleIdWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdWithoutIncludesResponseWithDefaults() *BundleIdWithoutIncludesResponse {
	this := BundleIdWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdWithoutIncludesResponse) GetData() Profile {
	if o == nil {
		var ret Profile
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdWithoutIncludesResponse) GetDataOk() (*Profile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdWithoutIncludesResponse) SetData(v Profile) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BundleIdWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BundleIdWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BundleIdWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BundleIdWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBundleIdWithoutIncludesResponse struct {
	value *BundleIdWithoutIncludesResponse
	isSet bool
}

func (v NullableBundleIdWithoutIncludesResponse) Get() *BundleIdWithoutIncludesResponse {
	return v.value
}

func (v *NullableBundleIdWithoutIncludesResponse) Set(val *BundleIdWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdWithoutIncludesResponse(val *BundleIdWithoutIncludesResponse) *NullableBundleIdWithoutIncludesResponse {
	return &NullableBundleIdWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableBundleIdWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
