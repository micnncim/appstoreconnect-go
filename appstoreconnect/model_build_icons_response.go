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

// checks if the BuildIconsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildIconsResponse{}

// BuildIconsResponse struct for BuildIconsResponse
type BuildIconsResponse struct {
	Data  []BuildIcon        `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// NewBuildIconsResponse instantiates a new BuildIconsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildIconsResponse(data []BuildIcon, links PagedDocumentLinks) *BuildIconsResponse {
	this := BuildIconsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildIconsResponseWithDefaults instantiates a new BuildIconsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildIconsResponseWithDefaults() *BuildIconsResponse {
	this := BuildIconsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildIconsResponse) GetData() []BuildIcon {
	if o == nil {
		var ret []BuildIcon
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildIconsResponse) GetDataOk() ([]BuildIcon, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BuildIconsResponse) SetData(v []BuildIcon) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BuildIconsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildIconsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildIconsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BuildIconsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildIconsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BuildIconsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BuildIconsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BuildIconsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildIconsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBuildIconsResponse struct {
	value *BuildIconsResponse
	isSet bool
}

func (v NullableBuildIconsResponse) Get() *BuildIconsResponse {
	return v.value
}

func (v *NullableBuildIconsResponse) Set(val *BuildIconsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildIconsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildIconsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildIconsResponse(val *BuildIconsResponse) *NullableBuildIconsResponse {
	return &NullableBuildIconsResponse{value: val, isSet: true}
}

func (v NullableBuildIconsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildIconsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
