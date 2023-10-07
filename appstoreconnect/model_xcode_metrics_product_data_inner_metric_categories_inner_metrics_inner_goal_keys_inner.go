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

// checks if the XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner{}

// XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner struct for XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner
type XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner struct {
	GoalKey    *string `json:"goalKey,omitempty"`
	LowerBound *int32  `json:"lowerBound,omitempty"`
	UpperBound *int32  `json:"upperBound,omitempty"`
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner {
	this := XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner{}
	return &this
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInnerWithDefaults instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInnerWithDefaults() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner {
	this := XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner{}
	return &this
}

// GetGoalKey returns the GoalKey field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) GetGoalKey() string {
	if o == nil || IsNil(o.GoalKey) {
		var ret string
		return ret
	}
	return *o.GoalKey
}

// GetGoalKeyOk returns a tuple with the GoalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) GetGoalKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GoalKey) {
		return nil, false
	}
	return o.GoalKey, true
}

// HasGoalKey returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) HasGoalKey() bool {
	if o != nil && !IsNil(o.GoalKey) {
		return true
	}

	return false
}

// SetGoalKey gets a reference to the given string and assigns it to the GoalKey field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) SetGoalKey(v string) {
	o.GoalKey = &v
}

// GetLowerBound returns the LowerBound field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) GetLowerBound() int32 {
	if o == nil || IsNil(o.LowerBound) {
		var ret int32
		return ret
	}
	return *o.LowerBound
}

// GetLowerBoundOk returns a tuple with the LowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) GetLowerBoundOk() (*int32, bool) {
	if o == nil || IsNil(o.LowerBound) {
		return nil, false
	}
	return o.LowerBound, true
}

// HasLowerBound returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) HasLowerBound() bool {
	if o != nil && !IsNil(o.LowerBound) {
		return true
	}

	return false
}

// SetLowerBound gets a reference to the given int32 and assigns it to the LowerBound field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) SetLowerBound(v int32) {
	o.LowerBound = &v
}

// GetUpperBound returns the UpperBound field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) GetUpperBound() int32 {
	if o == nil || IsNil(o.UpperBound) {
		var ret int32
		return ret
	}
	return *o.UpperBound
}

// GetUpperBoundOk returns a tuple with the UpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) GetUpperBoundOk() (*int32, bool) {
	if o == nil || IsNil(o.UpperBound) {
		return nil, false
	}
	return o.UpperBound, true
}

// HasUpperBound returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) HasUpperBound() bool {
	if o != nil && !IsNil(o.UpperBound) {
		return true
	}

	return false
}

// SetUpperBound gets a reference to the given int32 and assigns it to the UpperBound field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) SetUpperBound(v int32) {
	o.UpperBound = &v
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GoalKey) {
		toSerialize["goalKey"] = o.GoalKey
	}
	if !IsNil(o.LowerBound) {
		toSerialize["lowerBound"] = o.LowerBound
	}
	if !IsNil(o.UpperBound) {
		toSerialize["upperBound"] = o.UpperBound
	}
	return toSerialize, nil
}

type NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner struct {
	value *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner
	isSet bool
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) Get() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner {
	return v.value
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) Set(val *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner(val *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner {
	return &NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner{value: val, isSet: true}
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
