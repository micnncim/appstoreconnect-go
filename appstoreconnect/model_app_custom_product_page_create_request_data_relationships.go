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

// checks if the AppCustomProductPageCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageCreateRequestDataRelationships{}

// AppCustomProductPageCreateRequestDataRelationships struct for AppCustomProductPageCreateRequestDataRelationships
type AppCustomProductPageCreateRequestDataRelationships struct {
	App                          AppAvailabilityV2CreateRequestDataRelationshipsApp                                `json:"app"`
	AppCustomProductPageVersions *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions   `json:"appCustomProductPageVersions,omitempty"`
	AppStoreVersionTemplate      *AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersionTemplate,omitempty"`
	CustomProductPageTemplate    *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage         `json:"customProductPageTemplate,omitempty"`
}

// NewAppCustomProductPageCreateRequestDataRelationships instantiates a new AppCustomProductPageCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageCreateRequestDataRelationships(app AppAvailabilityV2CreateRequestDataRelationshipsApp) *AppCustomProductPageCreateRequestDataRelationships {
	this := AppCustomProductPageCreateRequestDataRelationships{}
	this.App = app
	return &this
}

// NewAppCustomProductPageCreateRequestDataRelationshipsWithDefaults instantiates a new AppCustomProductPageCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageCreateRequestDataRelationshipsWithDefaults() *AppCustomProductPageCreateRequestDataRelationships {
	this := AppCustomProductPageCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *AppCustomProductPageCreateRequestDataRelationships) GetApp() AppAvailabilityV2CreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AppAvailabilityV2CreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestDataRelationships) GetAppOk() (*AppAvailabilityV2CreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *AppCustomProductPageCreateRequestDataRelationships) SetApp(v AppAvailabilityV2CreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetAppCustomProductPageVersions returns the AppCustomProductPageVersions field value if set, zero value otherwise.
func (o *AppCustomProductPageCreateRequestDataRelationships) GetAppCustomProductPageVersions() AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions {
	if o == nil || IsNil(o.AppCustomProductPageVersions) {
		var ret AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions
		return ret
	}
	return *o.AppCustomProductPageVersions
}

// GetAppCustomProductPageVersionsOk returns a tuple with the AppCustomProductPageVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestDataRelationships) GetAppCustomProductPageVersionsOk() (*AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions, bool) {
	if o == nil || IsNil(o.AppCustomProductPageVersions) {
		return nil, false
	}
	return o.AppCustomProductPageVersions, true
}

// HasAppCustomProductPageVersions returns a boolean if a field has been set.
func (o *AppCustomProductPageCreateRequestDataRelationships) HasAppCustomProductPageVersions() bool {
	if o != nil && !IsNil(o.AppCustomProductPageVersions) {
		return true
	}

	return false
}

// SetAppCustomProductPageVersions gets a reference to the given AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions and assigns it to the AppCustomProductPageVersions field.
func (o *AppCustomProductPageCreateRequestDataRelationships) SetAppCustomProductPageVersions(v AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) {
	o.AppCustomProductPageVersions = &v
}

// GetAppStoreVersionTemplate returns the AppStoreVersionTemplate field value if set, zero value otherwise.
func (o *AppCustomProductPageCreateRequestDataRelationships) GetAppStoreVersionTemplate() AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersionTemplate) {
		var ret AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersionTemplate
}

// GetAppStoreVersionTemplateOk returns a tuple with the AppStoreVersionTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestDataRelationships) GetAppStoreVersionTemplateOk() (*AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersionTemplate) {
		return nil, false
	}
	return o.AppStoreVersionTemplate, true
}

// HasAppStoreVersionTemplate returns a boolean if a field has been set.
func (o *AppCustomProductPageCreateRequestDataRelationships) HasAppStoreVersionTemplate() bool {
	if o != nil && !IsNil(o.AppStoreVersionTemplate) {
		return true
	}

	return false
}

// SetAppStoreVersionTemplate gets a reference to the given AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersionTemplate field.
func (o *AppCustomProductPageCreateRequestDataRelationships) SetAppStoreVersionTemplate(v AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersionTemplate = &v
}

// GetCustomProductPageTemplate returns the CustomProductPageTemplate field value if set, zero value otherwise.
func (o *AppCustomProductPageCreateRequestDataRelationships) GetCustomProductPageTemplate() AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage {
	if o == nil || IsNil(o.CustomProductPageTemplate) {
		var ret AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage
		return ret
	}
	return *o.CustomProductPageTemplate
}

// GetCustomProductPageTemplateOk returns a tuple with the CustomProductPageTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestDataRelationships) GetCustomProductPageTemplateOk() (*AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage, bool) {
	if o == nil || IsNil(o.CustomProductPageTemplate) {
		return nil, false
	}
	return o.CustomProductPageTemplate, true
}

// HasCustomProductPageTemplate returns a boolean if a field has been set.
func (o *AppCustomProductPageCreateRequestDataRelationships) HasCustomProductPageTemplate() bool {
	if o != nil && !IsNil(o.CustomProductPageTemplate) {
		return true
	}

	return false
}

// SetCustomProductPageTemplate gets a reference to the given AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage and assigns it to the CustomProductPageTemplate field.
func (o *AppCustomProductPageCreateRequestDataRelationships) SetCustomProductPageTemplate(v AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) {
	o.CustomProductPageTemplate = &v
}

func (o AppCustomProductPageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	if !IsNil(o.AppCustomProductPageVersions) {
		toSerialize["appCustomProductPageVersions"] = o.AppCustomProductPageVersions
	}
	if !IsNil(o.AppStoreVersionTemplate) {
		toSerialize["appStoreVersionTemplate"] = o.AppStoreVersionTemplate
	}
	if !IsNil(o.CustomProductPageTemplate) {
		toSerialize["customProductPageTemplate"] = o.CustomProductPageTemplate
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageCreateRequestDataRelationships struct {
	value *AppCustomProductPageCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppCustomProductPageCreateRequestDataRelationships) Get() *AppCustomProductPageCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppCustomProductPageCreateRequestDataRelationships) Set(val *AppCustomProductPageCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageCreateRequestDataRelationships(val *AppCustomProductPageCreateRequestDataRelationships) *NullableAppCustomProductPageCreateRequestDataRelationships {
	return &NullableAppCustomProductPageCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppCustomProductPageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
