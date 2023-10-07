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

// checks if the UserInvitationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitationCreateRequest{}

// UserInvitationCreateRequest struct for UserInvitationCreateRequest
type UserInvitationCreateRequest struct {
	Data UserInvitationCreateRequestData `json:"data"`
}

// NewUserInvitationCreateRequest instantiates a new UserInvitationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationCreateRequest(data UserInvitationCreateRequestData) *UserInvitationCreateRequest {
	this := UserInvitationCreateRequest{}
	this.Data = data
	return &this
}

// NewUserInvitationCreateRequestWithDefaults instantiates a new UserInvitationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationCreateRequestWithDefaults() *UserInvitationCreateRequest {
	this := UserInvitationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *UserInvitationCreateRequest) GetData() UserInvitationCreateRequestData {
	if o == nil {
		var ret UserInvitationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequest) GetDataOk() (*UserInvitationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserInvitationCreateRequest) SetData(v UserInvitationCreateRequestData) {
	o.Data = v
}

func (o UserInvitationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableUserInvitationCreateRequest struct {
	value *UserInvitationCreateRequest
	isSet bool
}

func (v NullableUserInvitationCreateRequest) Get() *UserInvitationCreateRequest {
	return v.value
}

func (v *NullableUserInvitationCreateRequest) Set(val *UserInvitationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationCreateRequest(val *UserInvitationCreateRequest) *NullableUserInvitationCreateRequest {
	return &NullableUserInvitationCreateRequest{value: val, isSet: true}
}

func (v NullableUserInvitationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
