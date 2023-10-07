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

// checks if the BuildBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBundle{}

// BuildBundle struct for BuildBundle
type BuildBundle struct {
	Type          string                    `json:"type"`
	Id            string                    `json:"id"`
	Attributes    *BuildBundleAttributes    `json:"attributes,omitempty"`
	Relationships *BuildBundleRelationships `json:"relationships,omitempty"`
}

// NewBuildBundle instantiates a new BuildBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBundle(type_ string, id string) *BuildBundle {
	this := BuildBundle{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBuildBundleWithDefaults instantiates a new BuildBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBundleWithDefaults() *BuildBundle {
	this := BuildBundle{}
	return &this
}

// GetType returns the Type field value
func (o *BuildBundle) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BuildBundle) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuildBundle) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BuildBundle) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BuildBundle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BuildBundle) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BuildBundle) GetAttributes() BuildBundleAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BuildBundleAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundle) GetAttributesOk() (*BuildBundleAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BuildBundle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BuildBundleAttributes and assigns it to the Attributes field.
func (o *BuildBundle) SetAttributes(v BuildBundleAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BuildBundle) GetRelationships() BuildBundleRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BuildBundleRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundle) GetRelationshipsOk() (*BuildBundleRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BuildBundle) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BuildBundleRelationships and assigns it to the Relationships field.
func (o *BuildBundle) SetRelationships(v BuildBundleRelationships) {
	o.Relationships = &v
}

func (o BuildBundle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableBuildBundle struct {
	value *BuildBundle
	isSet bool
}

func (v NullableBuildBundle) Get() *BuildBundle {
	return v.value
}

func (v *NullableBuildBundle) Set(val *BuildBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBundle(val *BuildBundle) *NullableBuildBundle {
	return &NullableBuildBundle{value: val, isSet: true}
}

func (v NullableBuildBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
