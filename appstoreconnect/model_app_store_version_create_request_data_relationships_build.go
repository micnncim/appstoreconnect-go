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

// checks if the AppStoreVersionCreateRequestDataRelationshipsBuild type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionCreateRequestDataRelationshipsBuild{}

// AppStoreVersionCreateRequestDataRelationshipsBuild struct for AppStoreVersionCreateRequestDataRelationshipsBuild
type AppStoreVersionCreateRequestDataRelationshipsBuild struct {
	Data *AppEncryptionDeclarationRelationshipsBuildsDataInner `json:"data,omitempty"`
}

// NewAppStoreVersionCreateRequestDataRelationshipsBuild instantiates a new AppStoreVersionCreateRequestDataRelationshipsBuild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionCreateRequestDataRelationshipsBuild() *AppStoreVersionCreateRequestDataRelationshipsBuild {
	this := AppStoreVersionCreateRequestDataRelationshipsBuild{}
	return &this
}

// NewAppStoreVersionCreateRequestDataRelationshipsBuildWithDefaults instantiates a new AppStoreVersionCreateRequestDataRelationshipsBuild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionCreateRequestDataRelationshipsBuildWithDefaults() *AppStoreVersionCreateRequestDataRelationshipsBuild {
	this := AppStoreVersionCreateRequestDataRelationshipsBuild{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionCreateRequestDataRelationshipsBuild) GetData() AppEncryptionDeclarationRelationshipsBuildsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppEncryptionDeclarationRelationshipsBuildsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataRelationshipsBuild) GetDataOk() (*AppEncryptionDeclarationRelationshipsBuildsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionCreateRequestDataRelationshipsBuild) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppEncryptionDeclarationRelationshipsBuildsDataInner and assigns it to the Data field.
func (o *AppStoreVersionCreateRequestDataRelationshipsBuild) SetData(v AppEncryptionDeclarationRelationshipsBuildsDataInner) {
	o.Data = &v
}

func (o AppStoreVersionCreateRequestDataRelationshipsBuild) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionCreateRequestDataRelationshipsBuild) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreVersionCreateRequestDataRelationshipsBuild struct {
	value *AppStoreVersionCreateRequestDataRelationshipsBuild
	isSet bool
}

func (v NullableAppStoreVersionCreateRequestDataRelationshipsBuild) Get() *AppStoreVersionCreateRequestDataRelationshipsBuild {
	return v.value
}

func (v *NullableAppStoreVersionCreateRequestDataRelationshipsBuild) Set(val *AppStoreVersionCreateRequestDataRelationshipsBuild) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionCreateRequestDataRelationshipsBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionCreateRequestDataRelationshipsBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionCreateRequestDataRelationshipsBuild(val *AppStoreVersionCreateRequestDataRelationshipsBuild) *NullableAppStoreVersionCreateRequestDataRelationshipsBuild {
	return &NullableAppStoreVersionCreateRequestDataRelationshipsBuild{value: val, isSet: true}
}

func (v NullableAppStoreVersionCreateRequestDataRelationshipsBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionCreateRequestDataRelationshipsBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
