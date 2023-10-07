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

// checks if the AppAvailabilityV2CreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityV2CreateRequest{}

// AppAvailabilityV2CreateRequest struct for AppAvailabilityV2CreateRequest
type AppAvailabilityV2CreateRequest struct {
	Data     AppAvailabilityV2CreateRequestData  `json:"data"`
	Included []TerritoryAvailabilityInlineCreate `json:"included,omitempty"`
}

// NewAppAvailabilityV2CreateRequest instantiates a new AppAvailabilityV2CreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityV2CreateRequest(data AppAvailabilityV2CreateRequestData) *AppAvailabilityV2CreateRequest {
	this := AppAvailabilityV2CreateRequest{}
	this.Data = data
	return &this
}

// NewAppAvailabilityV2CreateRequestWithDefaults instantiates a new AppAvailabilityV2CreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityV2CreateRequestWithDefaults() *AppAvailabilityV2CreateRequest {
	this := AppAvailabilityV2CreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppAvailabilityV2CreateRequest) GetData() AppAvailabilityV2CreateRequestData {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2CreateRequest) GetDataOk() (*AppAvailabilityV2CreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppAvailabilityV2CreateRequest) SetData(v AppAvailabilityV2CreateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppAvailabilityV2CreateRequest) GetIncluded() []TerritoryAvailabilityInlineCreate {
	if o == nil || IsNil(o.Included) {
		var ret []TerritoryAvailabilityInlineCreate
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2CreateRequest) GetIncludedOk() ([]TerritoryAvailabilityInlineCreate, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppAvailabilityV2CreateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []TerritoryAvailabilityInlineCreate and assigns it to the Included field.
func (o *AppAvailabilityV2CreateRequest) SetIncluded(v []TerritoryAvailabilityInlineCreate) {
	o.Included = v
}

func (o AppAvailabilityV2CreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityV2CreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableAppAvailabilityV2CreateRequest struct {
	value *AppAvailabilityV2CreateRequest
	isSet bool
}

func (v NullableAppAvailabilityV2CreateRequest) Get() *AppAvailabilityV2CreateRequest {
	return v.value
}

func (v *NullableAppAvailabilityV2CreateRequest) Set(val *AppAvailabilityV2CreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityV2CreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityV2CreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityV2CreateRequest(val *AppAvailabilityV2CreateRequest) *NullableAppAvailabilityV2CreateRequest {
	return &NullableAppAvailabilityV2CreateRequest{value: val, isSet: true}
}

func (v NullableAppAvailabilityV2CreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityV2CreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
