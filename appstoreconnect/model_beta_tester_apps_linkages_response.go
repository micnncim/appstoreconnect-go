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

// checks if the BetaTesterAppsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterAppsLinkagesResponse{}

// BetaTesterAppsLinkagesResponse struct for BetaTesterAppsLinkagesResponse
type BetaTesterAppsLinkagesResponse struct {
	Data  []AppAvailabilityV2CreateRequestDataRelationshipsAppData `json:"data"`
	Links PagedDocumentLinks                                       `json:"links"`
	Meta  *PagingInformation                                       `json:"meta,omitempty"`
}

// NewBetaTesterAppsLinkagesResponse instantiates a new BetaTesterAppsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterAppsLinkagesResponse(data []AppAvailabilityV2CreateRequestDataRelationshipsAppData, links PagedDocumentLinks) *BetaTesterAppsLinkagesResponse {
	this := BetaTesterAppsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaTesterAppsLinkagesResponseWithDefaults instantiates a new BetaTesterAppsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterAppsLinkagesResponseWithDefaults() *BetaTesterAppsLinkagesResponse {
	this := BetaTesterAppsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterAppsLinkagesResponse) GetData() []AppAvailabilityV2CreateRequestDataRelationshipsAppData {
	if o == nil {
		var ret []AppAvailabilityV2CreateRequestDataRelationshipsAppData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterAppsLinkagesResponse) GetDataOk() ([]AppAvailabilityV2CreateRequestDataRelationshipsAppData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaTesterAppsLinkagesResponse) SetData(v []AppAvailabilityV2CreateRequestDataRelationshipsAppData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaTesterAppsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaTesterAppsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaTesterAppsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaTesterAppsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterAppsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaTesterAppsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaTesterAppsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaTesterAppsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterAppsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBetaTesterAppsLinkagesResponse struct {
	value *BetaTesterAppsLinkagesResponse
	isSet bool
}

func (v NullableBetaTesterAppsLinkagesResponse) Get() *BetaTesterAppsLinkagesResponse {
	return v.value
}

func (v *NullableBetaTesterAppsLinkagesResponse) Set(val *BetaTesterAppsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterAppsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterAppsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterAppsLinkagesResponse(val *BetaTesterAppsLinkagesResponse) *NullableBetaTesterAppsLinkagesResponse {
	return &NullableBetaTesterAppsLinkagesResponse{value: val, isSet: true}
}

func (v NullableBetaTesterAppsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterAppsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
