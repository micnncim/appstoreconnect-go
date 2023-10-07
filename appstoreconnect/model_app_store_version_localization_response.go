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

// checks if the AppStoreVersionLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationResponse{}

// AppStoreVersionLocalizationResponse struct for AppStoreVersionLocalizationResponse
type AppStoreVersionLocalizationResponse struct {
	Data     AppStoreVersionLocalization                         `json:"data"`
	Included []AppStoreVersionLocalizationsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                                       `json:"links"`
}

// NewAppStoreVersionLocalizationResponse instantiates a new AppStoreVersionLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationResponse(data AppStoreVersionLocalization, links DocumentLinks) *AppStoreVersionLocalizationResponse {
	this := AppStoreVersionLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionLocalizationResponseWithDefaults instantiates a new AppStoreVersionLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationResponseWithDefaults() *AppStoreVersionLocalizationResponse {
	this := AppStoreVersionLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionLocalizationResponse) GetData() AppStoreVersionLocalization {
	if o == nil {
		var ret AppStoreVersionLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationResponse) GetDataOk() (*AppStoreVersionLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionLocalizationResponse) SetData(v AppStoreVersionLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationResponse) GetIncluded() []AppStoreVersionLocalizationsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppStoreVersionLocalizationsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationResponse) GetIncludedOk() ([]AppStoreVersionLocalizationsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppStoreVersionLocalizationsResponseIncludedInner and assigns it to the Included field.
func (o *AppStoreVersionLocalizationResponse) SetIncluded(v []AppStoreVersionLocalizationsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppStoreVersionLocalizationResponse struct {
	value *AppStoreVersionLocalizationResponse
	isSet bool
}

func (v NullableAppStoreVersionLocalizationResponse) Get() *AppStoreVersionLocalizationResponse {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationResponse) Set(val *AppStoreVersionLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationResponse(val *AppStoreVersionLocalizationResponse) *NullableAppStoreVersionLocalizationResponse {
	return &NullableAppStoreVersionLocalizationResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
