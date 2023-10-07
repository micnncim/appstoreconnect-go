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

// checks if the BetaAppClipInvocationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationCreateRequest{}

// BetaAppClipInvocationCreateRequest struct for BetaAppClipInvocationCreateRequest
type BetaAppClipInvocationCreateRequest struct {
	Data     BetaAppClipInvocationCreateRequestData          `json:"data"`
	Included []BetaAppClipInvocationLocalizationInlineCreate `json:"included,omitempty"`
}

// NewBetaAppClipInvocationCreateRequest instantiates a new BetaAppClipInvocationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationCreateRequest(data BetaAppClipInvocationCreateRequestData) *BetaAppClipInvocationCreateRequest {
	this := BetaAppClipInvocationCreateRequest{}
	this.Data = data
	return &this
}

// NewBetaAppClipInvocationCreateRequestWithDefaults instantiates a new BetaAppClipInvocationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationCreateRequestWithDefaults() *BetaAppClipInvocationCreateRequest {
	this := BetaAppClipInvocationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationCreateRequest) GetData() BetaAppClipInvocationCreateRequestData {
	if o == nil {
		var ret BetaAppClipInvocationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationCreateRequest) GetDataOk() (*BetaAppClipInvocationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationCreateRequest) SetData(v BetaAppClipInvocationCreateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaAppClipInvocationCreateRequest) GetIncluded() []BetaAppClipInvocationLocalizationInlineCreate {
	if o == nil || IsNil(o.Included) {
		var ret []BetaAppClipInvocationLocalizationInlineCreate
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationCreateRequest) GetIncludedOk() ([]BetaAppClipInvocationLocalizationInlineCreate, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaAppClipInvocationCreateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BetaAppClipInvocationLocalizationInlineCreate and assigns it to the Included field.
func (o *BetaAppClipInvocationCreateRequest) SetIncluded(v []BetaAppClipInvocationLocalizationInlineCreate) {
	o.Included = v
}

func (o BetaAppClipInvocationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableBetaAppClipInvocationCreateRequest struct {
	value *BetaAppClipInvocationCreateRequest
	isSet bool
}

func (v NullableBetaAppClipInvocationCreateRequest) Get() *BetaAppClipInvocationCreateRequest {
	return v.value
}

func (v *NullableBetaAppClipInvocationCreateRequest) Set(val *BetaAppClipInvocationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationCreateRequest(val *BetaAppClipInvocationCreateRequest) *NullableBetaAppClipInvocationCreateRequest {
	return &NullableBetaAppClipInvocationCreateRequest{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
