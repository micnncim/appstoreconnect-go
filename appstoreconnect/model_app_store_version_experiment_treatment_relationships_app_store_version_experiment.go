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

// checks if the AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment{}

// AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment struct for AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
type AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks                   `json:"links,omitempty"`
	Data  *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData `json:"data,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment instantiates a new AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment() *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	this := AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment{}
	return &this
}

// NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentWithDefaults instantiates a new AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentWithDefaults() *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	this := AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) GetData() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData {
	if o == nil || IsNil(o.Data) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) GetDataOk() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData and assigns it to the Data field.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) SetData(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) {
	o.Data = &v
}

func (o AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment struct {
	value *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) Get() *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) Set(val *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment(val *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	return &NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
