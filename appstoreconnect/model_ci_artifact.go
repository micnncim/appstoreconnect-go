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

// checks if the CiArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiArtifact{}

// CiArtifact struct for CiArtifact
type CiArtifact struct {
	Type       string                `json:"type"`
	Id         string                `json:"id"`
	Attributes *CiArtifactAttributes `json:"attributes,omitempty"`
	Links      *ResourceLinks        `json:"links,omitempty"`
}

// NewCiArtifact instantiates a new CiArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiArtifact(type_ string, id string) *CiArtifact {
	this := CiArtifact{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiArtifactWithDefaults instantiates a new CiArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiArtifactWithDefaults() *CiArtifact {
	this := CiArtifact{}
	return &this
}

// GetType returns the Type field value
func (o *CiArtifact) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiArtifact) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiArtifact) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiArtifact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiArtifact) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiArtifact) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CiArtifact) GetAttributes() CiArtifactAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CiArtifactAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiArtifact) GetAttributesOk() (*CiArtifactAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CiArtifact) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CiArtifactAttributes and assigns it to the Attributes field.
func (o *CiArtifact) SetAttributes(v CiArtifactAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CiArtifact) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiArtifact) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CiArtifact) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *CiArtifact) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o CiArtifact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCiArtifact struct {
	value *CiArtifact
	isSet bool
}

func (v NullableCiArtifact) Get() *CiArtifact {
	return v.value
}

func (v *NullableCiArtifact) Set(val *CiArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableCiArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableCiArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiArtifact(val *CiArtifact) *NullableCiArtifact {
	return &NullableCiArtifact{value: val, isSet: true}
}

func (v NullableCiArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
