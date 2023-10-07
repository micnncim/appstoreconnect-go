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

// checks if the BetaAppClipInvocationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationRelationships{}

// BetaAppClipInvocationRelationships struct for BetaAppClipInvocationRelationships
type BetaAppClipInvocationRelationships struct {
	BetaAppClipInvocationLocalizations *BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizations `json:"betaAppClipInvocationLocalizations,omitempty"`
}

// NewBetaAppClipInvocationRelationships instantiates a new BetaAppClipInvocationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationRelationships() *BetaAppClipInvocationRelationships {
	this := BetaAppClipInvocationRelationships{}
	return &this
}

// NewBetaAppClipInvocationRelationshipsWithDefaults instantiates a new BetaAppClipInvocationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationRelationshipsWithDefaults() *BetaAppClipInvocationRelationships {
	this := BetaAppClipInvocationRelationships{}
	return &this
}

// GetBetaAppClipInvocationLocalizations returns the BetaAppClipInvocationLocalizations field value if set, zero value otherwise.
func (o *BetaAppClipInvocationRelationships) GetBetaAppClipInvocationLocalizations() BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizations {
	if o == nil || IsNil(o.BetaAppClipInvocationLocalizations) {
		var ret BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizations
		return ret
	}
	return *o.BetaAppClipInvocationLocalizations
}

// GetBetaAppClipInvocationLocalizationsOk returns a tuple with the BetaAppClipInvocationLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationRelationships) GetBetaAppClipInvocationLocalizationsOk() (*BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizations, bool) {
	if o == nil || IsNil(o.BetaAppClipInvocationLocalizations) {
		return nil, false
	}
	return o.BetaAppClipInvocationLocalizations, true
}

// HasBetaAppClipInvocationLocalizations returns a boolean if a field has been set.
func (o *BetaAppClipInvocationRelationships) HasBetaAppClipInvocationLocalizations() bool {
	if o != nil && !IsNil(o.BetaAppClipInvocationLocalizations) {
		return true
	}

	return false
}

// SetBetaAppClipInvocationLocalizations gets a reference to the given BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizations and assigns it to the BetaAppClipInvocationLocalizations field.
func (o *BetaAppClipInvocationRelationships) SetBetaAppClipInvocationLocalizations(v BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizations) {
	o.BetaAppClipInvocationLocalizations = &v
}

func (o BetaAppClipInvocationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BetaAppClipInvocationLocalizations) {
		toSerialize["betaAppClipInvocationLocalizations"] = o.BetaAppClipInvocationLocalizations
	}
	return toSerialize, nil
}

type NullableBetaAppClipInvocationRelationships struct {
	value *BetaAppClipInvocationRelationships
	isSet bool
}

func (v NullableBetaAppClipInvocationRelationships) Get() *BetaAppClipInvocationRelationships {
	return v.value
}

func (v *NullableBetaAppClipInvocationRelationships) Set(val *BetaAppClipInvocationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationRelationships(val *BetaAppClipInvocationRelationships) *NullableBetaAppClipInvocationRelationships {
	return &NullableBetaAppClipInvocationRelationships{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
