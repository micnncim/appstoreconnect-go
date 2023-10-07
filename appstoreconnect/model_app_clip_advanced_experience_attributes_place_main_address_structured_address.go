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

// checks if the AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress{}

// AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress struct for AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress
type AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress struct {
	StreetAddress []string `json:"streetAddress,omitempty"`
	Floor         *string  `json:"floor,omitempty"`
	Neighborhood  *string  `json:"neighborhood,omitempty"`
	Locality      *string  `json:"locality,omitempty"`
	StateProvince *string  `json:"stateProvince,omitempty"`
	PostalCode    *string  `json:"postalCode,omitempty"`
	CountryCode   *string  `json:"countryCode,omitempty"`
}

// NewAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress instantiates a new AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress() *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress {
	this := AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress{}
	return &this
}

// NewAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddressWithDefaults instantiates a new AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddressWithDefaults() *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress {
	this := AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress{}
	return &this
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetStreetAddress() []string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret []string
		return ret
	}
	return o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetStreetAddressOk() ([]string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given []string and assigns it to the StreetAddress field.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) SetStreetAddress(v []string) {
	o.StreetAddress = v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetFloor() string {
	if o == nil || IsNil(o.Floor) {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetFloorOk() (*string, bool) {
	if o == nil || IsNil(o.Floor) {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) HasFloor() bool {
	if o != nil && !IsNil(o.Floor) {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) SetFloor(v string) {
	o.Floor = &v
}

// GetNeighborhood returns the Neighborhood field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetNeighborhood() string {
	if o == nil || IsNil(o.Neighborhood) {
		var ret string
		return ret
	}
	return *o.Neighborhood
}

// GetNeighborhoodOk returns a tuple with the Neighborhood field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetNeighborhoodOk() (*string, bool) {
	if o == nil || IsNil(o.Neighborhood) {
		return nil, false
	}
	return o.Neighborhood, true
}

// HasNeighborhood returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) HasNeighborhood() bool {
	if o != nil && !IsNil(o.Neighborhood) {
		return true
	}

	return false
}

// SetNeighborhood gets a reference to the given string and assigns it to the Neighborhood field.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) SetNeighborhood(v string) {
	o.Neighborhood = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetLocality() string {
	if o == nil || IsNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetLocalityOk() (*string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) SetLocality(v string) {
	o.Locality = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetStateProvince() string {
	if o == nil || IsNil(o.StateProvince) {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetStateProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.StateProvince) {
		return nil, false
	}
	return o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) HasStateProvince() bool {
	if o != nil && !IsNil(o.StateProvince) {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) SetStateProvince(v string) {
	o.StateProvince = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StreetAddress) {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if !IsNil(o.Floor) {
		toSerialize["floor"] = o.Floor
	}
	if !IsNil(o.Neighborhood) {
		toSerialize["neighborhood"] = o.Neighborhood
	}
	if !IsNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !IsNil(o.StateProvince) {
		toSerialize["stateProvince"] = o.StateProvince
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress struct {
	value *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress
	isSet bool
}

func (v NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) Get() *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) Set(val *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress(val *AppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) *NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress {
	return &NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceAttributesPlaceMainAddressStructuredAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
