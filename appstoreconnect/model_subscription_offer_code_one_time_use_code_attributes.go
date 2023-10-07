/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
	"time"
)

// checks if the SubscriptionOfferCodeOneTimeUseCodeAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeOneTimeUseCodeAttributes{}

// SubscriptionOfferCodeOneTimeUseCodeAttributes struct for SubscriptionOfferCodeOneTimeUseCodeAttributes
type SubscriptionOfferCodeOneTimeUseCodeAttributes struct {
	NumberOfCodes  *int32     `json:"numberOfCodes,omitempty"`
	CreatedDate    *time.Time `json:"createdDate,omitempty"`
	ExpirationDate *string    `json:"expirationDate,omitempty"`
	Active         *bool      `json:"active,omitempty"`
}

// NewSubscriptionOfferCodeOneTimeUseCodeAttributes instantiates a new SubscriptionOfferCodeOneTimeUseCodeAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeOneTimeUseCodeAttributes() *SubscriptionOfferCodeOneTimeUseCodeAttributes {
	this := SubscriptionOfferCodeOneTimeUseCodeAttributes{}
	return &this
}

// NewSubscriptionOfferCodeOneTimeUseCodeAttributesWithDefaults instantiates a new SubscriptionOfferCodeOneTimeUseCodeAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeOneTimeUseCodeAttributesWithDefaults() *SubscriptionOfferCodeOneTimeUseCodeAttributes {
	this := SubscriptionOfferCodeOneTimeUseCodeAttributes{}
	return &this
}

// GetNumberOfCodes returns the NumberOfCodes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetNumberOfCodes() int32 {
	if o == nil || IsNil(o.NumberOfCodes) {
		var ret int32
		return ret
	}
	return *o.NumberOfCodes
}

// GetNumberOfCodesOk returns a tuple with the NumberOfCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetNumberOfCodesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfCodes) {
		return nil, false
	}
	return o.NumberOfCodes, true
}

// HasNumberOfCodes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) HasNumberOfCodes() bool {
	if o != nil && !IsNil(o.NumberOfCodes) {
		return true
	}

	return false
}

// SetNumberOfCodes gets a reference to the given int32 and assigns it to the NumberOfCodes field.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) SetNumberOfCodes(v int32) {
	o.NumberOfCodes = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SubscriptionOfferCodeOneTimeUseCodeAttributes) SetActive(v bool) {
	o.Active = &v
}

func (o SubscriptionOfferCodeOneTimeUseCodeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeOneTimeUseCodeAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfCodes) {
		toSerialize["numberOfCodes"] = o.NumberOfCodes
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeOneTimeUseCodeAttributes struct {
	value *SubscriptionOfferCodeOneTimeUseCodeAttributes
	isSet bool
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeAttributes) Get() *SubscriptionOfferCodeOneTimeUseCodeAttributes {
	return v.value
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeAttributes) Set(val *SubscriptionOfferCodeOneTimeUseCodeAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeOneTimeUseCodeAttributes(val *SubscriptionOfferCodeOneTimeUseCodeAttributes) *NullableSubscriptionOfferCodeOneTimeUseCodeAttributes {
	return &NullableSubscriptionOfferCodeOneTimeUseCodeAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
