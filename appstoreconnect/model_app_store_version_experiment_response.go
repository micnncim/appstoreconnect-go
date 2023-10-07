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

// checks if the AppStoreVersionExperimentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentResponse{}

// AppStoreVersionExperimentResponse struct for AppStoreVersionExperimentResponse
type AppStoreVersionExperimentResponse struct {
	// Deprecated
	Data     AppStoreVersionExperiment                         `json:"data"`
	Included []AppStoreVersionExperimentsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                                     `json:"links"`
}

// NewAppStoreVersionExperimentResponse instantiates a new AppStoreVersionExperimentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentResponse(data AppStoreVersionExperiment, links DocumentLinks) *AppStoreVersionExperimentResponse {
	this := AppStoreVersionExperimentResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionExperimentResponseWithDefaults instantiates a new AppStoreVersionExperimentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentResponseWithDefaults() *AppStoreVersionExperimentResponse {
	this := AppStoreVersionExperimentResponse{}
	return &this
}

// GetData returns the Data field value
// Deprecated
func (o *AppStoreVersionExperimentResponse) GetData() AppStoreVersionExperiment {
	if o == nil {
		var ret AppStoreVersionExperiment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppStoreVersionExperimentResponse) GetDataOk() (*AppStoreVersionExperiment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
// Deprecated
func (o *AppStoreVersionExperimentResponse) SetData(v AppStoreVersionExperiment) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentResponse) GetIncluded() []AppStoreVersionExperimentsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppStoreVersionExperimentsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentResponse) GetIncludedOk() ([]AppStoreVersionExperimentsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppStoreVersionExperimentsResponseIncludedInner and assigns it to the Included field.
func (o *AppStoreVersionExperimentResponse) SetIncluded(v []AppStoreVersionExperimentsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionExperimentResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionExperimentResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionExperimentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentResponse struct {
	value *AppStoreVersionExperimentResponse
	isSet bool
}

func (v NullableAppStoreVersionExperimentResponse) Get() *AppStoreVersionExperimentResponse {
	return v.value
}

func (v *NullableAppStoreVersionExperimentResponse) Set(val *AppStoreVersionExperimentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentResponse(val *AppStoreVersionExperimentResponse) *NullableAppStoreVersionExperimentResponse {
	return &NullableAppStoreVersionExperimentResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
