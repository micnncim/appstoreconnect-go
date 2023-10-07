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

// checks if the SubscriptionOfferCodePriceInlineCreateRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodePriceInlineCreateRelationships{}

// SubscriptionOfferCodePriceInlineCreateRelationships struct for SubscriptionOfferCodePriceInlineCreateRelationships
type SubscriptionOfferCodePriceInlineCreateRelationships struct {
	Territory              *InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory        `json:"territory,omitempty"`
	SubscriptionPricePoint *SubscriptionIntroductoryOfferInlineCreateRelationshipsSubscriptionPricePoint `json:"subscriptionPricePoint,omitempty"`
}

// NewSubscriptionOfferCodePriceInlineCreateRelationships instantiates a new SubscriptionOfferCodePriceInlineCreateRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodePriceInlineCreateRelationships() *SubscriptionOfferCodePriceInlineCreateRelationships {
	this := SubscriptionOfferCodePriceInlineCreateRelationships{}
	return &this
}

// NewSubscriptionOfferCodePriceInlineCreateRelationshipsWithDefaults instantiates a new SubscriptionOfferCodePriceInlineCreateRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodePriceInlineCreateRelationshipsWithDefaults() *SubscriptionOfferCodePriceInlineCreateRelationships {
	this := SubscriptionOfferCodePriceInlineCreateRelationships{}
	return &this
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) GetTerritory() InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory {
	if o == nil || IsNil(o.Territory) {
		var ret InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) GetTerritoryOk() (*InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory and assigns it to the Territory field.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) SetTerritory(v InAppPurchasePriceScheduleCreateRequestDataRelationshipsBaseTerritory) {
	o.Territory = &v
}

// GetSubscriptionPricePoint returns the SubscriptionPricePoint field value if set, zero value otherwise.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) GetSubscriptionPricePoint() SubscriptionIntroductoryOfferInlineCreateRelationshipsSubscriptionPricePoint {
	if o == nil || IsNil(o.SubscriptionPricePoint) {
		var ret SubscriptionIntroductoryOfferInlineCreateRelationshipsSubscriptionPricePoint
		return ret
	}
	return *o.SubscriptionPricePoint
}

// GetSubscriptionPricePointOk returns a tuple with the SubscriptionPricePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) GetSubscriptionPricePointOk() (*SubscriptionIntroductoryOfferInlineCreateRelationshipsSubscriptionPricePoint, bool) {
	if o == nil || IsNil(o.SubscriptionPricePoint) {
		return nil, false
	}
	return o.SubscriptionPricePoint, true
}

// HasSubscriptionPricePoint returns a boolean if a field has been set.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) HasSubscriptionPricePoint() bool {
	if o != nil && !IsNil(o.SubscriptionPricePoint) {
		return true
	}

	return false
}

// SetSubscriptionPricePoint gets a reference to the given SubscriptionIntroductoryOfferInlineCreateRelationshipsSubscriptionPricePoint and assigns it to the SubscriptionPricePoint field.
func (o *SubscriptionOfferCodePriceInlineCreateRelationships) SetSubscriptionPricePoint(v SubscriptionIntroductoryOfferInlineCreateRelationshipsSubscriptionPricePoint) {
	o.SubscriptionPricePoint = &v
}

func (o SubscriptionOfferCodePriceInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodePriceInlineCreateRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	if !IsNil(o.SubscriptionPricePoint) {
		toSerialize["subscriptionPricePoint"] = o.SubscriptionPricePoint
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodePriceInlineCreateRelationships struct {
	value *SubscriptionOfferCodePriceInlineCreateRelationships
	isSet bool
}

func (v NullableSubscriptionOfferCodePriceInlineCreateRelationships) Get() *SubscriptionOfferCodePriceInlineCreateRelationships {
	return v.value
}

func (v *NullableSubscriptionOfferCodePriceInlineCreateRelationships) Set(val *SubscriptionOfferCodePriceInlineCreateRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodePriceInlineCreateRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodePriceInlineCreateRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodePriceInlineCreateRelationships(val *SubscriptionOfferCodePriceInlineCreateRelationships) *NullableSubscriptionOfferCodePriceInlineCreateRelationships {
	return &NullableSubscriptionOfferCodePriceInlineCreateRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodePriceInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodePriceInlineCreateRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
