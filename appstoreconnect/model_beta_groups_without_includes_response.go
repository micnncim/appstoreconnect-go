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

// checks if the BetaGroupsWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupsWithoutIncludesResponse{}

// BetaGroupsWithoutIncludesResponse struct for BetaGroupsWithoutIncludesResponse
type BetaGroupsWithoutIncludesResponse struct {
	Data  []BetaTester       `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// NewBetaGroupsWithoutIncludesResponse instantiates a new BetaGroupsWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupsWithoutIncludesResponse(data []BetaTester, links PagedDocumentLinks) *BetaGroupsWithoutIncludesResponse {
	this := BetaGroupsWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaGroupsWithoutIncludesResponseWithDefaults instantiates a new BetaGroupsWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupsWithoutIncludesResponseWithDefaults() *BetaGroupsWithoutIncludesResponse {
	this := BetaGroupsWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupsWithoutIncludesResponse) GetData() []BetaTester {
	if o == nil {
		var ret []BetaTester
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupsWithoutIncludesResponse) GetDataOk() ([]BetaTester, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaGroupsWithoutIncludesResponse) SetData(v []BetaTester) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaGroupsWithoutIncludesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaGroupsWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaGroupsWithoutIncludesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaGroupsWithoutIncludesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupsWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaGroupsWithoutIncludesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaGroupsWithoutIncludesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaGroupsWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupsWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBetaGroupsWithoutIncludesResponse struct {
	value *BetaGroupsWithoutIncludesResponse
	isSet bool
}

func (v NullableBetaGroupsWithoutIncludesResponse) Get() *BetaGroupsWithoutIncludesResponse {
	return v.value
}

func (v *NullableBetaGroupsWithoutIncludesResponse) Set(val *BetaGroupsWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupsWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupsWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupsWithoutIncludesResponse(val *BetaGroupsWithoutIncludesResponse) *NullableBetaGroupsWithoutIncludesResponse {
	return &NullableBetaGroupsWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableBetaGroupsWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupsWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
