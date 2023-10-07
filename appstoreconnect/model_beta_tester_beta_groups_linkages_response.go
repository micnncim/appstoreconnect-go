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

// checks if the BetaTesterBetaGroupsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterBetaGroupsLinkagesResponse{}

// BetaTesterBetaGroupsLinkagesResponse struct for BetaTesterBetaGroupsLinkagesResponse
type BetaTesterBetaGroupsLinkagesResponse struct {
	Data  []AppRelationshipsBetaGroupsDataInner `json:"data"`
	Links PagedDocumentLinks                    `json:"links"`
	Meta  *PagingInformation                    `json:"meta,omitempty"`
}

// NewBetaTesterBetaGroupsLinkagesResponse instantiates a new BetaTesterBetaGroupsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterBetaGroupsLinkagesResponse(data []AppRelationshipsBetaGroupsDataInner, links PagedDocumentLinks) *BetaTesterBetaGroupsLinkagesResponse {
	this := BetaTesterBetaGroupsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaTesterBetaGroupsLinkagesResponseWithDefaults instantiates a new BetaTesterBetaGroupsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterBetaGroupsLinkagesResponseWithDefaults() *BetaTesterBetaGroupsLinkagesResponse {
	this := BetaTesterBetaGroupsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterBetaGroupsLinkagesResponse) GetData() []AppRelationshipsBetaGroupsDataInner {
	if o == nil {
		var ret []AppRelationshipsBetaGroupsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterBetaGroupsLinkagesResponse) GetDataOk() ([]AppRelationshipsBetaGroupsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaTesterBetaGroupsLinkagesResponse) SetData(v []AppRelationshipsBetaGroupsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaTesterBetaGroupsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaTesterBetaGroupsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaTesterBetaGroupsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaTesterBetaGroupsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterBetaGroupsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaTesterBetaGroupsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaTesterBetaGroupsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaTesterBetaGroupsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterBetaGroupsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBetaTesterBetaGroupsLinkagesResponse struct {
	value *BetaTesterBetaGroupsLinkagesResponse
	isSet bool
}

func (v NullableBetaTesterBetaGroupsLinkagesResponse) Get() *BetaTesterBetaGroupsLinkagesResponse {
	return v.value
}

func (v *NullableBetaTesterBetaGroupsLinkagesResponse) Set(val *BetaTesterBetaGroupsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterBetaGroupsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterBetaGroupsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterBetaGroupsLinkagesResponse(val *BetaTesterBetaGroupsLinkagesResponse) *NullableBetaTesterBetaGroupsLinkagesResponse {
	return &NullableBetaTesterBetaGroupsLinkagesResponse{value: val, isSet: true}
}

func (v NullableBetaTesterBetaGroupsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterBetaGroupsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
