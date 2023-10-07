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

// checks if the SubscriptionOfferCodeOneTimeUseCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeOneTimeUseCode{}

// SubscriptionOfferCodeOneTimeUseCode struct for SubscriptionOfferCodeOneTimeUseCode
type SubscriptionOfferCodeOneTimeUseCode struct {
	Type          string                                         `json:"type"`
	Id            string                                         `json:"id"`
	Attributes    *SubscriptionOfferCodeOneTimeUseCodeAttributes `json:"attributes,omitempty"`
	Relationships *SubscriptionOfferCodeCustomCodeRelationships  `json:"relationships,omitempty"`
	Links         *ResourceLinks                                 `json:"links,omitempty"`
}

// NewSubscriptionOfferCodeOneTimeUseCode instantiates a new SubscriptionOfferCodeOneTimeUseCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeOneTimeUseCode(type_ string, id string) *SubscriptionOfferCodeOneTimeUseCode {
	this := SubscriptionOfferCodeOneTimeUseCode{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionOfferCodeOneTimeUseCodeWithDefaults instantiates a new SubscriptionOfferCodeOneTimeUseCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeOneTimeUseCodeWithDefaults() *SubscriptionOfferCodeOneTimeUseCode {
	this := SubscriptionOfferCodeOneTimeUseCode{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeOneTimeUseCode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeOneTimeUseCode) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionOfferCodeOneTimeUseCode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionOfferCodeOneTimeUseCode) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetAttributes() SubscriptionOfferCodeOneTimeUseCodeAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionOfferCodeOneTimeUseCodeAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetAttributesOk() (*SubscriptionOfferCodeOneTimeUseCodeAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionOfferCodeOneTimeUseCodeAttributes and assigns it to the Attributes field.
func (o *SubscriptionOfferCodeOneTimeUseCode) SetAttributes(v SubscriptionOfferCodeOneTimeUseCodeAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetRelationships() SubscriptionOfferCodeCustomCodeRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionOfferCodeCustomCodeRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetRelationshipsOk() (*SubscriptionOfferCodeCustomCodeRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionOfferCodeCustomCodeRelationships and assigns it to the Relationships field.
func (o *SubscriptionOfferCodeOneTimeUseCode) SetRelationships(v SubscriptionOfferCodeCustomCodeRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCode) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *SubscriptionOfferCodeOneTimeUseCode) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o SubscriptionOfferCodeOneTimeUseCode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeOneTimeUseCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeOneTimeUseCode struct {
	value *SubscriptionOfferCodeOneTimeUseCode
	isSet bool
}

func (v NullableSubscriptionOfferCodeOneTimeUseCode) Get() *SubscriptionOfferCodeOneTimeUseCode {
	return v.value
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCode) Set(val *SubscriptionOfferCodeOneTimeUseCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeOneTimeUseCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeOneTimeUseCode(val *SubscriptionOfferCodeOneTimeUseCode) *NullableSubscriptionOfferCodeOneTimeUseCode {
	return &NullableSubscriptionOfferCodeOneTimeUseCode{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeOneTimeUseCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
