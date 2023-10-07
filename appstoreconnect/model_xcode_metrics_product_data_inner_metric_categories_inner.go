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

// checks if the XcodeMetricsProductDataInnerMetricCategoriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XcodeMetricsProductDataInnerMetricCategoriesInner{}

// XcodeMetricsProductDataInnerMetricCategoriesInner struct for XcodeMetricsProductDataInnerMetricCategoriesInner
type XcodeMetricsProductDataInnerMetricCategoriesInner struct {
	Identifier *MetricCategory                                                 `json:"identifier,omitempty"`
	Metrics    []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner `json:"metrics,omitempty"`
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInner instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXcodeMetricsProductDataInnerMetricCategoriesInner() *XcodeMetricsProductDataInnerMetricCategoriesInner {
	this := XcodeMetricsProductDataInnerMetricCategoriesInner{}
	return &this
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInnerWithDefaults instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXcodeMetricsProductDataInnerMetricCategoriesInnerWithDefaults() *XcodeMetricsProductDataInnerMetricCategoriesInner {
	this := XcodeMetricsProductDataInnerMetricCategoriesInner{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) GetIdentifier() MetricCategory {
	if o == nil || IsNil(o.Identifier) {
		var ret MetricCategory
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) GetIdentifierOk() (*MetricCategory, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given MetricCategory and assigns it to the Identifier field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) SetIdentifier(v MetricCategory) {
	o.Identifier = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) GetMetrics() []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner {
	if o == nil || IsNil(o.Metrics) {
		var ret []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) GetMetricsOk() ([]XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner and assigns it to the Metrics field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInner) SetMetrics(v []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) {
	o.Metrics = v
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableXcodeMetricsProductDataInnerMetricCategoriesInner struct {
	value *XcodeMetricsProductDataInnerMetricCategoriesInner
	isSet bool
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInner) Get() *XcodeMetricsProductDataInnerMetricCategoriesInner {
	return v.value
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInner) Set(val *XcodeMetricsProductDataInnerMetricCategoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXcodeMetricsProductDataInnerMetricCategoriesInner(val *XcodeMetricsProductDataInnerMetricCategoriesInner) *NullableXcodeMetricsProductDataInnerMetricCategoriesInner {
	return &NullableXcodeMetricsProductDataInnerMetricCategoriesInner{value: val, isSet: true}
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
