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

// checks if the ScmRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmRepository{}

// ScmRepository struct for ScmRepository
type ScmRepository struct {
	Type          string                      `json:"type"`
	Id            string                      `json:"id"`
	Attributes    *ScmRepositoryAttributes    `json:"attributes,omitempty"`
	Relationships *ScmRepositoryRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks              `json:"links,omitempty"`
}

// NewScmRepository instantiates a new ScmRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmRepository(type_ string, id string) *ScmRepository {
	this := ScmRepository{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewScmRepositoryWithDefaults instantiates a new ScmRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmRepositoryWithDefaults() *ScmRepository {
	this := ScmRepository{}
	return &this
}

// GetType returns the Type field value
func (o *ScmRepository) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScmRepository) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScmRepository) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ScmRepository) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScmRepository) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScmRepository) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ScmRepository) GetAttributes() ScmRepositoryAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ScmRepositoryAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepository) GetAttributesOk() (*ScmRepositoryAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ScmRepository) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ScmRepositoryAttributes and assigns it to the Attributes field.
func (o *ScmRepository) SetAttributes(v ScmRepositoryAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScmRepository) GetRelationships() ScmRepositoryRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ScmRepositoryRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepository) GetRelationshipsOk() (*ScmRepositoryRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScmRepository) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ScmRepositoryRelationships and assigns it to the Relationships field.
func (o *ScmRepository) SetRelationships(v ScmRepositoryRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ScmRepository) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepository) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ScmRepository) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *ScmRepository) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o ScmRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmRepository) ToMap() (map[string]interface{}, error) {
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

type NullableScmRepository struct {
	value *ScmRepository
	isSet bool
}

func (v NullableScmRepository) Get() *ScmRepository {
	return v.value
}

func (v *NullableScmRepository) Set(val *ScmRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableScmRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableScmRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmRepository(val *ScmRepository) *NullableScmRepository {
	return &NullableScmRepository{value: val, isSet: true}
}

func (v NullableScmRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
