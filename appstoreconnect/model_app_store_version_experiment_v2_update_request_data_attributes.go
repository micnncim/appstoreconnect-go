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

// checks if the AppStoreVersionExperimentV2UpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentV2UpdateRequestDataAttributes{}

// AppStoreVersionExperimentV2UpdateRequestDataAttributes struct for AppStoreVersionExperimentV2UpdateRequestDataAttributes
type AppStoreVersionExperimentV2UpdateRequestDataAttributes struct {
	Name              *string `json:"name,omitempty"`
	TrafficProportion *int32  `json:"trafficProportion,omitempty"`
	Started           *bool   `json:"started,omitempty"`
}

// NewAppStoreVersionExperimentV2UpdateRequestDataAttributes instantiates a new AppStoreVersionExperimentV2UpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentV2UpdateRequestDataAttributes() *AppStoreVersionExperimentV2UpdateRequestDataAttributes {
	this := AppStoreVersionExperimentV2UpdateRequestDataAttributes{}
	return &this
}

// NewAppStoreVersionExperimentV2UpdateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionExperimentV2UpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentV2UpdateRequestDataAttributesWithDefaults() *AppStoreVersionExperimentV2UpdateRequestDataAttributes {
	this := AppStoreVersionExperimentV2UpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetTrafficProportion returns the TrafficProportion field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) GetTrafficProportion() int32 {
	if o == nil || IsNil(o.TrafficProportion) {
		var ret int32
		return ret
	}
	return *o.TrafficProportion
}

// GetTrafficProportionOk returns a tuple with the TrafficProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) GetTrafficProportionOk() (*int32, bool) {
	if o == nil || IsNil(o.TrafficProportion) {
		return nil, false
	}
	return o.TrafficProportion, true
}

// HasTrafficProportion returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) HasTrafficProportion() bool {
	if o != nil && !IsNil(o.TrafficProportion) {
		return true
	}

	return false
}

// SetTrafficProportion gets a reference to the given int32 and assigns it to the TrafficProportion field.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) SetTrafficProportion(v int32) {
	o.TrafficProportion = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) GetStarted() bool {
	if o == nil || IsNil(o.Started) {
		var ret bool
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) GetStartedOk() (*bool, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given bool and assigns it to the Started field.
func (o *AppStoreVersionExperimentV2UpdateRequestDataAttributes) SetStarted(v bool) {
	o.Started = &v
}

func (o AppStoreVersionExperimentV2UpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentV2UpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TrafficProportion) {
		toSerialize["trafficProportion"] = o.TrafficProportion
	}
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes struct {
	value *AppStoreVersionExperimentV2UpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes) Get() *AppStoreVersionExperimentV2UpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes) Set(val *AppStoreVersionExperimentV2UpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentV2UpdateRequestDataAttributes(val *AppStoreVersionExperimentV2UpdateRequestDataAttributes) *NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes {
	return &NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentV2UpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
