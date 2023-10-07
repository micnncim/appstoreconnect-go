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

// checks if the AppClipDefaultExperience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperience{}

// AppClipDefaultExperience struct for AppClipDefaultExperience
type AppClipDefaultExperience struct {
	Type          string                                 `json:"type"`
	Id            string                                 `json:"id"`
	Attributes    *AppClipDefaultExperienceAttributes    `json:"attributes,omitempty"`
	Relationships *AppClipDefaultExperienceRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                         `json:"links,omitempty"`
}

// NewAppClipDefaultExperience instantiates a new AppClipDefaultExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperience(type_ string, id string) *AppClipDefaultExperience {
	this := AppClipDefaultExperience{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipDefaultExperienceWithDefaults instantiates a new AppClipDefaultExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceWithDefaults() *AppClipDefaultExperience {
	this := AppClipDefaultExperience{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDefaultExperience) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperience) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDefaultExperience) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipDefaultExperience) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperience) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipDefaultExperience) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipDefaultExperience) GetAttributes() AppClipDefaultExperienceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipDefaultExperienceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperience) GetAttributesOk() (*AppClipDefaultExperienceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipDefaultExperience) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipDefaultExperienceAttributes and assigns it to the Attributes field.
func (o *AppClipDefaultExperience) SetAttributes(v AppClipDefaultExperienceAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppClipDefaultExperience) GetRelationships() AppClipDefaultExperienceRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppClipDefaultExperienceRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperience) GetRelationshipsOk() (*AppClipDefaultExperienceRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppClipDefaultExperience) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppClipDefaultExperienceRelationships and assigns it to the Relationships field.
func (o *AppClipDefaultExperience) SetRelationships(v AppClipDefaultExperienceRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppClipDefaultExperience) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperience) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppClipDefaultExperience) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppClipDefaultExperience) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppClipDefaultExperience) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperience) ToMap() (map[string]interface{}, error) {
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

type NullableAppClipDefaultExperience struct {
	value *AppClipDefaultExperience
	isSet bool
}

func (v NullableAppClipDefaultExperience) Get() *AppClipDefaultExperience {
	return v.value
}

func (v *NullableAppClipDefaultExperience) Set(val *AppClipDefaultExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperience(val *AppClipDefaultExperience) *NullableAppClipDefaultExperience {
	return &NullableAppClipDefaultExperience{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
