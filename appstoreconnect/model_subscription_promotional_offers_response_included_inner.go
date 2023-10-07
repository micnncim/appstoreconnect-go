/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
	"fmt"
)

// SubscriptionPromotionalOffersResponseIncludedInner - struct for SubscriptionPromotionalOffersResponseIncludedInner
type SubscriptionPromotionalOffersResponseIncludedInner struct {
	Subscription                      *Subscription
	SubscriptionPromotionalOfferPrice *SubscriptionPromotionalOfferPrice
}

// SubscriptionAsSubscriptionPromotionalOffersResponseIncludedInner is a convenience function that returns Subscription wrapped in SubscriptionPromotionalOffersResponseIncludedInner
func SubscriptionAsSubscriptionPromotionalOffersResponseIncludedInner(v *Subscription) SubscriptionPromotionalOffersResponseIncludedInner {
	return SubscriptionPromotionalOffersResponseIncludedInner{
		Subscription: v,
	}
}

// SubscriptionPromotionalOfferPriceAsSubscriptionPromotionalOffersResponseIncludedInner is a convenience function that returns SubscriptionPromotionalOfferPrice wrapped in SubscriptionPromotionalOffersResponseIncludedInner
func SubscriptionPromotionalOfferPriceAsSubscriptionPromotionalOffersResponseIncludedInner(v *SubscriptionPromotionalOfferPrice) SubscriptionPromotionalOffersResponseIncludedInner {
	return SubscriptionPromotionalOffersResponseIncludedInner{
		SubscriptionPromotionalOfferPrice: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionPromotionalOffersResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Subscription
	err = newStrictDecoder(data).Decode(&dst.Subscription)
	if err == nil {
		jsonSubscription, _ := json.Marshal(dst.Subscription)
		if string(jsonSubscription) == "{}" { // empty struct
			dst.Subscription = nil
		} else {
			match++
		}
	} else {
		dst.Subscription = nil
	}

	// try to unmarshal data into SubscriptionPromotionalOfferPrice
	err = newStrictDecoder(data).Decode(&dst.SubscriptionPromotionalOfferPrice)
	if err == nil {
		jsonSubscriptionPromotionalOfferPrice, _ := json.Marshal(dst.SubscriptionPromotionalOfferPrice)
		if string(jsonSubscriptionPromotionalOfferPrice) == "{}" { // empty struct
			dst.SubscriptionPromotionalOfferPrice = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionPromotionalOfferPrice = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Subscription = nil
		dst.SubscriptionPromotionalOfferPrice = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscriptionPromotionalOffersResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscriptionPromotionalOffersResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionPromotionalOffersResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.Subscription != nil {
		return json.Marshal(&src.Subscription)
	}

	if src.SubscriptionPromotionalOfferPrice != nil {
		return json.Marshal(&src.SubscriptionPromotionalOfferPrice)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscriptionPromotionalOffersResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Subscription != nil {
		return obj.Subscription
	}

	if obj.SubscriptionPromotionalOfferPrice != nil {
		return obj.SubscriptionPromotionalOfferPrice
	}

	// all schemas are nil
	return nil
}

type NullableSubscriptionPromotionalOffersResponseIncludedInner struct {
	value *SubscriptionPromotionalOffersResponseIncludedInner
	isSet bool
}

func (v NullableSubscriptionPromotionalOffersResponseIncludedInner) Get() *SubscriptionPromotionalOffersResponseIncludedInner {
	return v.value
}

func (v *NullableSubscriptionPromotionalOffersResponseIncludedInner) Set(val *SubscriptionPromotionalOffersResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOffersResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOffersResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOffersResponseIncludedInner(val *SubscriptionPromotionalOffersResponseIncludedInner) *NullableSubscriptionPromotionalOffersResponseIncludedInner {
	return &NullableSubscriptionPromotionalOffersResponseIncludedInner{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOffersResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOffersResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
