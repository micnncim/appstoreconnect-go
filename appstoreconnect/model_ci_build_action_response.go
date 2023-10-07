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

// checks if the CiBuildActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildActionResponse{}

// CiBuildActionResponse struct for CiBuildActionResponse
type CiBuildActionResponse struct {
	Data     CiBuildAction `json:"data"`
	Included []CiBuildRun  `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}

// NewCiBuildActionResponse instantiates a new CiBuildActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildActionResponse(data CiBuildAction, links DocumentLinks) *CiBuildActionResponse {
	this := CiBuildActionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiBuildActionResponseWithDefaults instantiates a new CiBuildActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildActionResponseWithDefaults() *CiBuildActionResponse {
	this := CiBuildActionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiBuildActionResponse) GetData() CiBuildAction {
	if o == nil {
		var ret CiBuildAction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiBuildActionResponse) GetDataOk() (*CiBuildAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiBuildActionResponse) SetData(v CiBuildAction) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CiBuildActionResponse) GetIncluded() []CiBuildRun {
	if o == nil || IsNil(o.Included) {
		var ret []CiBuildRun
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionResponse) GetIncludedOk() ([]CiBuildRun, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CiBuildActionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CiBuildRun and assigns it to the Included field.
func (o *CiBuildActionResponse) SetIncluded(v []CiBuildRun) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *CiBuildActionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiBuildActionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiBuildActionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CiBuildActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableCiBuildActionResponse struct {
	value *CiBuildActionResponse
	isSet bool
}

func (v NullableCiBuildActionResponse) Get() *CiBuildActionResponse {
	return v.value
}

func (v *NullableCiBuildActionResponse) Set(val *CiBuildActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildActionResponse(val *CiBuildActionResponse) *NullableCiBuildActionResponse {
	return &NullableCiBuildActionResponse{value: val, isSet: true}
}

func (v NullableCiBuildActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
