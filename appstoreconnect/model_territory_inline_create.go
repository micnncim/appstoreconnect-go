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

// checks if the TerritoryInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerritoryInlineCreate{}

// TerritoryInlineCreate struct for TerritoryInlineCreate
type TerritoryInlineCreate struct {
	Type string  `json:"type"`
	Id   *string `json:"id,omitempty"`
}

// NewTerritoryInlineCreate instantiates a new TerritoryInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerritoryInlineCreate(type_ string) *TerritoryInlineCreate {
	this := TerritoryInlineCreate{}
	this.Type = type_
	return &this
}

// NewTerritoryInlineCreateWithDefaults instantiates a new TerritoryInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerritoryInlineCreateWithDefaults() *TerritoryInlineCreate {
	this := TerritoryInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *TerritoryInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerritoryInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TerritoryInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TerritoryInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TerritoryInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TerritoryInlineCreate) SetId(v string) {
	o.Id = &v
}

func (o TerritoryInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerritoryInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableTerritoryInlineCreate struct {
	value *TerritoryInlineCreate
	isSet bool
}

func (v NullableTerritoryInlineCreate) Get() *TerritoryInlineCreate {
	return v.value
}

func (v *NullableTerritoryInlineCreate) Set(val *TerritoryInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTerritoryInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTerritoryInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerritoryInlineCreate(val *TerritoryInlineCreate) *NullableTerritoryInlineCreate {
	return &NullableTerritoryInlineCreate{value: val, isSet: true}
}

func (v NullableTerritoryInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerritoryInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
