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

// checks if the CiMacOsVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiMacOsVersion{}

// CiMacOsVersion struct for CiMacOsVersion
type CiMacOsVersion struct {
	Type          string                       `json:"type"`
	Id            string                       `json:"id"`
	Attributes    *CiMacOsVersionAttributes    `json:"attributes,omitempty"`
	Relationships *CiMacOsVersionRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks               `json:"links,omitempty"`
}

// NewCiMacOsVersion instantiates a new CiMacOsVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiMacOsVersion(type_ string, id string) *CiMacOsVersion {
	this := CiMacOsVersion{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiMacOsVersionWithDefaults instantiates a new CiMacOsVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiMacOsVersionWithDefaults() *CiMacOsVersion {
	this := CiMacOsVersion{}
	return &this
}

// GetType returns the Type field value
func (o *CiMacOsVersion) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersion) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiMacOsVersion) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiMacOsVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiMacOsVersion) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CiMacOsVersion) GetAttributes() CiMacOsVersionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CiMacOsVersionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersion) GetAttributesOk() (*CiMacOsVersionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CiMacOsVersion) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CiMacOsVersionAttributes and assigns it to the Attributes field.
func (o *CiMacOsVersion) SetAttributes(v CiMacOsVersionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CiMacOsVersion) GetRelationships() CiMacOsVersionRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CiMacOsVersionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersion) GetRelationshipsOk() (*CiMacOsVersionRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CiMacOsVersion) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CiMacOsVersionRelationships and assigns it to the Relationships field.
func (o *CiMacOsVersion) SetRelationships(v CiMacOsVersionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CiMacOsVersion) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersion) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CiMacOsVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *CiMacOsVersion) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o CiMacOsVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiMacOsVersion) ToMap() (map[string]interface{}, error) {
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

type NullableCiMacOsVersion struct {
	value *CiMacOsVersion
	isSet bool
}

func (v NullableCiMacOsVersion) Get() *CiMacOsVersion {
	return v.value
}

func (v *NullableCiMacOsVersion) Set(val *CiMacOsVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCiMacOsVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCiMacOsVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiMacOsVersion(val *CiMacOsVersion) *NullableCiMacOsVersion {
	return &NullableCiMacOsVersion{value: val, isSet: true}
}

func (v NullableCiMacOsVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiMacOsVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
