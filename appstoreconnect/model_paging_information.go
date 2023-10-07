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

// checks if the PagingInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingInformation{}

// PagingInformation struct for PagingInformation
type PagingInformation struct {
	Paging PagingInformationPaging `json:"paging"`
}

// NewPagingInformation instantiates a new PagingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingInformation(paging PagingInformationPaging) *PagingInformation {
	this := PagingInformation{}
	this.Paging = paging
	return &this
}

// NewPagingInformationWithDefaults instantiates a new PagingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingInformationWithDefaults() *PagingInformation {
	this := PagingInformation{}
	return &this
}

// GetPaging returns the Paging field value
func (o *PagingInformation) GetPaging() PagingInformationPaging {
	if o == nil {
		var ret PagingInformationPaging
		return ret
	}

	return o.Paging
}

// GetPagingOk returns a tuple with the Paging field value
// and a boolean to check if the value has been set.
func (o *PagingInformation) GetPagingOk() (*PagingInformationPaging, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Paging, true
}

// SetPaging sets field value
func (o *PagingInformation) SetPaging(v PagingInformationPaging) {
	o.Paging = v
}

func (o PagingInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paging"] = o.Paging
	return toSerialize, nil
}

type NullablePagingInformation struct {
	value *PagingInformation
	isSet bool
}

func (v NullablePagingInformation) Get() *PagingInformation {
	return v.value
}

func (v *NullablePagingInformation) Set(val *PagingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingInformation(val *PagingInformation) *NullablePagingInformation {
	return &NullablePagingInformation{value: val, isSet: true}
}

func (v NullablePagingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
