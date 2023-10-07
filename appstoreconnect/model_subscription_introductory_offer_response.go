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

// checks if the SubscriptionIntroductoryOfferResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferResponse{}

// SubscriptionIntroductoryOfferResponse struct for SubscriptionIntroductoryOfferResponse
type SubscriptionIntroductoryOfferResponse struct {
	Data     SubscriptionIntroductoryOffer                         `json:"data"`
	Included []SubscriptionIntroductoryOffersResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                                         `json:"links"`
}

// NewSubscriptionIntroductoryOfferResponse instantiates a new SubscriptionIntroductoryOfferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferResponse(data SubscriptionIntroductoryOffer, links DocumentLinks) *SubscriptionIntroductoryOfferResponse {
	this := SubscriptionIntroductoryOfferResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionIntroductoryOfferResponseWithDefaults instantiates a new SubscriptionIntroductoryOfferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferResponseWithDefaults() *SubscriptionIntroductoryOfferResponse {
	this := SubscriptionIntroductoryOfferResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionIntroductoryOfferResponse) GetData() SubscriptionIntroductoryOffer {
	if o == nil {
		var ret SubscriptionIntroductoryOffer
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferResponse) GetDataOk() (*SubscriptionIntroductoryOffer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionIntroductoryOfferResponse) SetData(v SubscriptionIntroductoryOffer) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferResponse) GetIncluded() []SubscriptionIntroductoryOffersResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionIntroductoryOffersResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferResponse) GetIncludedOk() ([]SubscriptionIntroductoryOffersResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionIntroductoryOffersResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionIntroductoryOfferResponse) SetIncluded(v []SubscriptionIntroductoryOffersResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionIntroductoryOfferResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionIntroductoryOfferResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionIntroductoryOfferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSubscriptionIntroductoryOfferResponse struct {
	value *SubscriptionIntroductoryOfferResponse
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferResponse) Get() *SubscriptionIntroductoryOfferResponse {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferResponse) Set(val *SubscriptionIntroductoryOfferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferResponse(val *SubscriptionIntroductoryOfferResponse) *NullableSubscriptionIntroductoryOfferResponse {
	return &NullableSubscriptionIntroductoryOfferResponse{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
