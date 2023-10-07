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

// checks if the SubscriptionPricePointsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPricePointsResponse{}

// SubscriptionPricePointsResponse struct for SubscriptionPricePointsResponse
type SubscriptionPricePointsResponse struct {
	Data     []SubscriptionPricePoint `json:"data"`
	Included []Territory              `json:"included,omitempty"`
	Links    PagedDocumentLinks       `json:"links"`
	Meta     *PagingInformation       `json:"meta,omitempty"`
}

// NewSubscriptionPricePointsResponse instantiates a new SubscriptionPricePointsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPricePointsResponse(data []SubscriptionPricePoint, links PagedDocumentLinks) *SubscriptionPricePointsResponse {
	this := SubscriptionPricePointsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionPricePointsResponseWithDefaults instantiates a new SubscriptionPricePointsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPricePointsResponseWithDefaults() *SubscriptionPricePointsResponse {
	this := SubscriptionPricePointsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPricePointsResponse) GetData() []SubscriptionPricePoint {
	if o == nil {
		var ret []SubscriptionPricePoint
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPricePointsResponse) GetDataOk() ([]SubscriptionPricePoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionPricePointsResponse) SetData(v []SubscriptionPricePoint) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionPricePointsResponse) GetIncluded() []Territory {
	if o == nil || IsNil(o.Included) {
		var ret []Territory
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricePointsResponse) GetIncludedOk() ([]Territory, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionPricePointsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []Territory and assigns it to the Included field.
func (o *SubscriptionPricePointsResponse) SetIncluded(v []Territory) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionPricePointsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPricePointsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionPricePointsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionPricePointsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricePointsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionPricePointsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionPricePointsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionPricePointsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPricePointsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionPricePointsResponse struct {
	value *SubscriptionPricePointsResponse
	isSet bool
}

func (v NullableSubscriptionPricePointsResponse) Get() *SubscriptionPricePointsResponse {
	return v.value
}

func (v *NullableSubscriptionPricePointsResponse) Set(val *SubscriptionPricePointsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPricePointsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPricePointsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPricePointsResponse(val *SubscriptionPricePointsResponse) *NullableSubscriptionPricePointsResponse {
	return &NullableSubscriptionPricePointsResponse{value: val, isSet: true}
}

func (v NullableSubscriptionPricePointsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPricePointsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
