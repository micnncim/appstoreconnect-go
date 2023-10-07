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

// checks if the BetaTesterInvitationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterInvitationCreateRequestDataRelationships{}

// BetaTesterInvitationCreateRequestDataRelationships struct for BetaTesterInvitationCreateRequestDataRelationships
type BetaTesterInvitationCreateRequestDataRelationships struct {
	BetaTester BetaTesterInvitationCreateRequestDataRelationshipsBetaTester `json:"betaTester"`
	App        AppAvailabilityV2CreateRequestDataRelationshipsApp           `json:"app"`
}

// NewBetaTesterInvitationCreateRequestDataRelationships instantiates a new BetaTesterInvitationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterInvitationCreateRequestDataRelationships(betaTester BetaTesterInvitationCreateRequestDataRelationshipsBetaTester, app AppAvailabilityV2CreateRequestDataRelationshipsApp) *BetaTesterInvitationCreateRequestDataRelationships {
	this := BetaTesterInvitationCreateRequestDataRelationships{}
	this.BetaTester = betaTester
	this.App = app
	return &this
}

// NewBetaTesterInvitationCreateRequestDataRelationshipsWithDefaults instantiates a new BetaTesterInvitationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterInvitationCreateRequestDataRelationshipsWithDefaults() *BetaTesterInvitationCreateRequestDataRelationships {
	this := BetaTesterInvitationCreateRequestDataRelationships{}
	return &this
}

// GetBetaTester returns the BetaTester field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetBetaTester() BetaTesterInvitationCreateRequestDataRelationshipsBetaTester {
	if o == nil {
		var ret BetaTesterInvitationCreateRequestDataRelationshipsBetaTester
		return ret
	}

	return o.BetaTester
}

// GetBetaTesterOk returns a tuple with the BetaTester field value
// and a boolean to check if the value has been set.
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetBetaTesterOk() (*BetaTesterInvitationCreateRequestDataRelationshipsBetaTester, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BetaTester, true
}

// SetBetaTester sets field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) SetBetaTester(v BetaTesterInvitationCreateRequestDataRelationshipsBetaTester) {
	o.BetaTester = v
}

// GetApp returns the App field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetApp() AppAvailabilityV2CreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetAppOk() (*AppAvailabilityV2CreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) SetApp(v AppAvailabilityV2CreateRequestDataRelationshipsApp) {
	o.App = v
}

func (o BetaTesterInvitationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterInvitationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["betaTester"] = o.BetaTester
	toSerialize["app"] = o.App
	return toSerialize, nil
}

type NullableBetaTesterInvitationCreateRequestDataRelationships struct {
	value *BetaTesterInvitationCreateRequestDataRelationships
	isSet bool
}

func (v NullableBetaTesterInvitationCreateRequestDataRelationships) Get() *BetaTesterInvitationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBetaTesterInvitationCreateRequestDataRelationships) Set(val *BetaTesterInvitationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterInvitationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterInvitationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterInvitationCreateRequestDataRelationships(val *BetaTesterInvitationCreateRequestDataRelationships) *NullableBetaTesterInvitationCreateRequestDataRelationships {
	return &NullableBetaTesterInvitationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBetaTesterInvitationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterInvitationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
