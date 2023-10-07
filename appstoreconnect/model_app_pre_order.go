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

// checks if the AppPreOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreOrder{}

// AppPreOrder struct for AppPreOrder
type AppPreOrder struct {
	Type          string                    `json:"type"`
	Id            string                    `json:"id"`
	Attributes    *AppPreOrderAttributes    `json:"attributes,omitempty"`
	Relationships *AppPreOrderRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks            `json:"links,omitempty"`
}

// NewAppPreOrder instantiates a new AppPreOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreOrder(type_ string, id string) *AppPreOrder {
	this := AppPreOrder{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppPreOrderWithDefaults instantiates a new AppPreOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreOrderWithDefaults() *AppPreOrder {
	this := AppPreOrder{}
	return &this
}

// GetType returns the Type field value
func (o *AppPreOrder) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPreOrder) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPreOrder) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPreOrder) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPreOrder) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPreOrder) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppPreOrder) GetAttributes() AppPreOrderAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppPreOrderAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreOrder) GetAttributesOk() (*AppPreOrderAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppPreOrder) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppPreOrderAttributes and assigns it to the Attributes field.
func (o *AppPreOrder) SetAttributes(v AppPreOrderAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppPreOrder) GetRelationships() AppPreOrderRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppPreOrderRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreOrder) GetRelationshipsOk() (*AppPreOrderRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppPreOrder) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppPreOrderRelationships and assigns it to the Relationships field.
func (o *AppPreOrder) SetRelationships(v AppPreOrderRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPreOrder) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreOrder) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPreOrder) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppPreOrder) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppPreOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreOrder) ToMap() (map[string]interface{}, error) {
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

type NullableAppPreOrder struct {
	value *AppPreOrder
	isSet bool
}

func (v NullableAppPreOrder) Get() *AppPreOrder {
	return v.value
}

func (v *NullableAppPreOrder) Set(val *AppPreOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreOrder(val *AppPreOrder) *NullableAppPreOrder {
	return &NullableAppPreOrder{value: val, isSet: true}
}

func (v NullableAppPreOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
