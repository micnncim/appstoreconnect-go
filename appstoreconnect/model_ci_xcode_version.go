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

// checks if the CiXcodeVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiXcodeVersion{}

// CiXcodeVersion struct for CiXcodeVersion
type CiXcodeVersion struct {
	Type          string                       `json:"type"`
	Id            string                       `json:"id"`
	Attributes    *CiXcodeVersionAttributes    `json:"attributes,omitempty"`
	Relationships *CiXcodeVersionRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks               `json:"links,omitempty"`
}

// NewCiXcodeVersion instantiates a new CiXcodeVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiXcodeVersion(type_ string, id string) *CiXcodeVersion {
	this := CiXcodeVersion{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiXcodeVersionWithDefaults instantiates a new CiXcodeVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiXcodeVersionWithDefaults() *CiXcodeVersion {
	this := CiXcodeVersion{}
	return &this
}

// GetType returns the Type field value
func (o *CiXcodeVersion) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiXcodeVersion) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiXcodeVersion) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiXcodeVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiXcodeVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiXcodeVersion) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CiXcodeVersion) GetAttributes() CiXcodeVersionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CiXcodeVersionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersion) GetAttributesOk() (*CiXcodeVersionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CiXcodeVersion) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CiXcodeVersionAttributes and assigns it to the Attributes field.
func (o *CiXcodeVersion) SetAttributes(v CiXcodeVersionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CiXcodeVersion) GetRelationships() CiXcodeVersionRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CiXcodeVersionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersion) GetRelationshipsOk() (*CiXcodeVersionRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CiXcodeVersion) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CiXcodeVersionRelationships and assigns it to the Relationships field.
func (o *CiXcodeVersion) SetRelationships(v CiXcodeVersionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CiXcodeVersion) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersion) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CiXcodeVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *CiXcodeVersion) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o CiXcodeVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiXcodeVersion) ToMap() (map[string]interface{}, error) {
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

type NullableCiXcodeVersion struct {
	value *CiXcodeVersion
	isSet bool
}

func (v NullableCiXcodeVersion) Get() *CiXcodeVersion {
	return v.value
}

func (v *NullableCiXcodeVersion) Set(val *CiXcodeVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCiXcodeVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCiXcodeVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiXcodeVersion(val *CiXcodeVersion) *NullableCiXcodeVersion {
	return &NullableCiXcodeVersion{value: val, isSet: true}
}

func (v NullableCiXcodeVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiXcodeVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
