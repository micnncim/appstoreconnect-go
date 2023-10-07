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

// checks if the ReviewSubmissionItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionItemsResponse{}

// ReviewSubmissionItemsResponse struct for ReviewSubmissionItemsResponse
type ReviewSubmissionItemsResponse struct {
	Data     []ReviewSubmissionItem                       `json:"data"`
	Included []ReviewSubmissionItemsResponseIncludedInner `json:"included,omitempty"`
	Links    PagedDocumentLinks                           `json:"links"`
	Meta     *PagingInformation                           `json:"meta,omitempty"`
}

// NewReviewSubmissionItemsResponse instantiates a new ReviewSubmissionItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionItemsResponse(data []ReviewSubmissionItem, links PagedDocumentLinks) *ReviewSubmissionItemsResponse {
	this := ReviewSubmissionItemsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewReviewSubmissionItemsResponseWithDefaults instantiates a new ReviewSubmissionItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionItemsResponseWithDefaults() *ReviewSubmissionItemsResponse {
	this := ReviewSubmissionItemsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ReviewSubmissionItemsResponse) GetData() []ReviewSubmissionItem {
	if o == nil {
		var ret []ReviewSubmissionItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemsResponse) GetDataOk() ([]ReviewSubmissionItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ReviewSubmissionItemsResponse) SetData(v []ReviewSubmissionItem) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ReviewSubmissionItemsResponse) GetIncluded() []ReviewSubmissionItemsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []ReviewSubmissionItemsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemsResponse) GetIncludedOk() ([]ReviewSubmissionItemsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ReviewSubmissionItemsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []ReviewSubmissionItemsResponseIncludedInner and assigns it to the Included field.
func (o *ReviewSubmissionItemsResponse) SetIncluded(v []ReviewSubmissionItemsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *ReviewSubmissionItemsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReviewSubmissionItemsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReviewSubmissionItemsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReviewSubmissionItemsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *ReviewSubmissionItemsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o ReviewSubmissionItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableReviewSubmissionItemsResponse struct {
	value *ReviewSubmissionItemsResponse
	isSet bool
}

func (v NullableReviewSubmissionItemsResponse) Get() *ReviewSubmissionItemsResponse {
	return v.value
}

func (v *NullableReviewSubmissionItemsResponse) Set(val *ReviewSubmissionItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionItemsResponse(val *ReviewSubmissionItemsResponse) *NullableReviewSubmissionItemsResponse {
	return &NullableReviewSubmissionItemsResponse{value: val, isSet: true}
}

func (v NullableReviewSubmissionItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
