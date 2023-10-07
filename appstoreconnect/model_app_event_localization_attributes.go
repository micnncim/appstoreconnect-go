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

// checks if the AppEventLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationAttributes{}

// AppEventLocalizationAttributes struct for AppEventLocalizationAttributes
type AppEventLocalizationAttributes struct {
	Locale           *string `json:"locale,omitempty"`
	Name             *string `json:"name,omitempty"`
	ShortDescription *string `json:"shortDescription,omitempty"`
	LongDescription  *string `json:"longDescription,omitempty"`
}

// NewAppEventLocalizationAttributes instantiates a new AppEventLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationAttributes() *AppEventLocalizationAttributes {
	this := AppEventLocalizationAttributes{}
	return &this
}

// NewAppEventLocalizationAttributesWithDefaults instantiates a new AppEventLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationAttributesWithDefaults() *AppEventLocalizationAttributes {
	this := AppEventLocalizationAttributes{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AppEventLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AppEventLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AppEventLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppEventLocalizationAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppEventLocalizationAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppEventLocalizationAttributes) SetName(v string) {
	o.Name = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *AppEventLocalizationAttributes) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *AppEventLocalizationAttributes) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *AppEventLocalizationAttributes) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetLongDescription returns the LongDescription field value if set, zero value otherwise.
func (o *AppEventLocalizationAttributes) GetLongDescription() string {
	if o == nil || IsNil(o.LongDescription) {
		var ret string
		return ret
	}
	return *o.LongDescription
}

// GetLongDescriptionOk returns a tuple with the LongDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationAttributes) GetLongDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.LongDescription) {
		return nil, false
	}
	return o.LongDescription, true
}

// HasLongDescription returns a boolean if a field has been set.
func (o *AppEventLocalizationAttributes) HasLongDescription() bool {
	if o != nil && !IsNil(o.LongDescription) {
		return true
	}

	return false
}

// SetLongDescription gets a reference to the given string and assigns it to the LongDescription field.
func (o *AppEventLocalizationAttributes) SetLongDescription(v string) {
	o.LongDescription = &v
}

func (o AppEventLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.LongDescription) {
		toSerialize["longDescription"] = o.LongDescription
	}
	return toSerialize, nil
}

type NullableAppEventLocalizationAttributes struct {
	value *AppEventLocalizationAttributes
	isSet bool
}

func (v NullableAppEventLocalizationAttributes) Get() *AppEventLocalizationAttributes {
	return v.value
}

func (v *NullableAppEventLocalizationAttributes) Set(val *AppEventLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationAttributes(val *AppEventLocalizationAttributes) *NullableAppEventLocalizationAttributes {
	return &NullableAppEventLocalizationAttributes{value: val, isSet: true}
}

func (v NullableAppEventLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
