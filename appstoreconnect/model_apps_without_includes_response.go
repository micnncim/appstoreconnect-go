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

// checks if the AppsWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsWithoutIncludesResponse{}

// AppsWithoutIncludesResponse struct for AppsWithoutIncludesResponse
type AppsWithoutIncludesResponse struct {
	Data  []User             `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// NewAppsWithoutIncludesResponse instantiates a new AppsWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsWithoutIncludesResponse(data []User, links PagedDocumentLinks) *AppsWithoutIncludesResponse {
	this := AppsWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppsWithoutIncludesResponseWithDefaults instantiates a new AppsWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsWithoutIncludesResponseWithDefaults() *AppsWithoutIncludesResponse {
	this := AppsWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppsWithoutIncludesResponse) GetData() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppsWithoutIncludesResponse) GetDataOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppsWithoutIncludesResponse) SetData(v []User) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppsWithoutIncludesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppsWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppsWithoutIncludesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppsWithoutIncludesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppsWithoutIncludesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppsWithoutIncludesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppsWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableAppsWithoutIncludesResponse struct {
	value *AppsWithoutIncludesResponse
	isSet bool
}

func (v NullableAppsWithoutIncludesResponse) Get() *AppsWithoutIncludesResponse {
	return v.value
}

func (v *NullableAppsWithoutIncludesResponse) Set(val *AppsWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsWithoutIncludesResponse(val *AppsWithoutIncludesResponse) *NullableAppsWithoutIncludesResponse {
	return &NullableAppsWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableAppsWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
