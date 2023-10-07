/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
	"time"
)

// checks if the CertificateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAttributes{}

// CertificateAttributes struct for CertificateAttributes
type CertificateAttributes struct {
	Name               *string           `json:"name,omitempty"`
	CertificateType    *CertificateType  `json:"certificateType,omitempty"`
	DisplayName        *string           `json:"displayName,omitempty"`
	SerialNumber       *string           `json:"serialNumber,omitempty"`
	Platform           *BundleIdPlatform `json:"platform,omitempty"`
	ExpirationDate     *time.Time        `json:"expirationDate,omitempty"`
	CertificateContent *string           `json:"certificateContent,omitempty"`
}

// NewCertificateAttributes instantiates a new CertificateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAttributes() *CertificateAttributes {
	this := CertificateAttributes{}
	return &this
}

// NewCertificateAttributesWithDefaults instantiates a new CertificateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAttributesWithDefaults() *CertificateAttributes {
	this := CertificateAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CertificateAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CertificateAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CertificateAttributes) SetName(v string) {
	o.Name = &v
}

// GetCertificateType returns the CertificateType field value if set, zero value otherwise.
func (o *CertificateAttributes) GetCertificateType() CertificateType {
	if o == nil || IsNil(o.CertificateType) {
		var ret CertificateType
		return ret
	}
	return *o.CertificateType
}

// GetCertificateTypeOk returns a tuple with the CertificateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAttributes) GetCertificateTypeOk() (*CertificateType, bool) {
	if o == nil || IsNil(o.CertificateType) {
		return nil, false
	}
	return o.CertificateType, true
}

// HasCertificateType returns a boolean if a field has been set.
func (o *CertificateAttributes) HasCertificateType() bool {
	if o != nil && !IsNil(o.CertificateType) {
		return true
	}

	return false
}

// SetCertificateType gets a reference to the given CertificateType and assigns it to the CertificateType field.
func (o *CertificateAttributes) SetCertificateType(v CertificateType) {
	o.CertificateType = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CertificateAttributes) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CertificateAttributes) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CertificateAttributes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *CertificateAttributes) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAttributes) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *CertificateAttributes) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *CertificateAttributes) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CertificateAttributes) GetPlatform() BundleIdPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret BundleIdPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAttributes) GetPlatformOk() (*BundleIdPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CertificateAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given BundleIdPlatform and assigns it to the Platform field.
func (o *CertificateAttributes) SetPlatform(v BundleIdPlatform) {
	o.Platform = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CertificateAttributes) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAttributes) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CertificateAttributes) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *CertificateAttributes) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetCertificateContent returns the CertificateContent field value if set, zero value otherwise.
func (o *CertificateAttributes) GetCertificateContent() string {
	if o == nil || IsNil(o.CertificateContent) {
		var ret string
		return ret
	}
	return *o.CertificateContent
}

// GetCertificateContentOk returns a tuple with the CertificateContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAttributes) GetCertificateContentOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateContent) {
		return nil, false
	}
	return o.CertificateContent, true
}

// HasCertificateContent returns a boolean if a field has been set.
func (o *CertificateAttributes) HasCertificateContent() bool {
	if o != nil && !IsNil(o.CertificateContent) {
		return true
	}

	return false
}

// SetCertificateContent gets a reference to the given string and assigns it to the CertificateContent field.
func (o *CertificateAttributes) SetCertificateContent(v string) {
	o.CertificateContent = &v
}

func (o CertificateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CertificateType) {
		toSerialize["certificateType"] = o.CertificateType
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.CertificateContent) {
		toSerialize["certificateContent"] = o.CertificateContent
	}
	return toSerialize, nil
}

type NullableCertificateAttributes struct {
	value *CertificateAttributes
	isSet bool
}

func (v NullableCertificateAttributes) Get() *CertificateAttributes {
	return v.value
}

func (v *NullableCertificateAttributes) Set(val *CertificateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAttributes(val *CertificateAttributes) *NullableCertificateAttributes {
	return &NullableCertificateAttributes{value: val, isSet: true}
}

func (v NullableCertificateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
