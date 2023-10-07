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

// checks if the DeviceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAttributes{}

// DeviceAttributes struct for DeviceAttributes
type DeviceAttributes struct {
	Name        *string           `json:"name,omitempty"`
	Platform    *BundleIdPlatform `json:"platform,omitempty"`
	Udid        *string           `json:"udid,omitempty"`
	DeviceClass *string           `json:"deviceClass,omitempty"`
	Status      *string           `json:"status,omitempty"`
	Model       *string           `json:"model,omitempty"`
	AddedDate   *time.Time        `json:"addedDate,omitempty"`
}

// NewDeviceAttributes instantiates a new DeviceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAttributes() *DeviceAttributes {
	this := DeviceAttributes{}
	return &this
}

// NewDeviceAttributesWithDefaults instantiates a new DeviceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAttributesWithDefaults() *DeviceAttributes {
	this := DeviceAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceAttributes) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DeviceAttributes) GetPlatform() BundleIdPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret BundleIdPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetPlatformOk() (*BundleIdPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DeviceAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given BundleIdPlatform and assigns it to the Platform field.
func (o *DeviceAttributes) SetPlatform(v BundleIdPlatform) {
	o.Platform = &v
}

// GetUdid returns the Udid field value if set, zero value otherwise.
func (o *DeviceAttributes) GetUdid() string {
	if o == nil || IsNil(o.Udid) {
		var ret string
		return ret
	}
	return *o.Udid
}

// GetUdidOk returns a tuple with the Udid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetUdidOk() (*string, bool) {
	if o == nil || IsNil(o.Udid) {
		return nil, false
	}
	return o.Udid, true
}

// HasUdid returns a boolean if a field has been set.
func (o *DeviceAttributes) HasUdid() bool {
	if o != nil && !IsNil(o.Udid) {
		return true
	}

	return false
}

// SetUdid gets a reference to the given string and assigns it to the Udid field.
func (o *DeviceAttributes) SetUdid(v string) {
	o.Udid = &v
}

// GetDeviceClass returns the DeviceClass field value if set, zero value otherwise.
func (o *DeviceAttributes) GetDeviceClass() string {
	if o == nil || IsNil(o.DeviceClass) {
		var ret string
		return ret
	}
	return *o.DeviceClass
}

// GetDeviceClassOk returns a tuple with the DeviceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetDeviceClassOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceClass) {
		return nil, false
	}
	return o.DeviceClass, true
}

// HasDeviceClass returns a boolean if a field has been set.
func (o *DeviceAttributes) HasDeviceClass() bool {
	if o != nil && !IsNil(o.DeviceClass) {
		return true
	}

	return false
}

// SetDeviceClass gets a reference to the given string and assigns it to the DeviceClass field.
func (o *DeviceAttributes) SetDeviceClass(v string) {
	o.DeviceClass = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceAttributes) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceAttributes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DeviceAttributes) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DeviceAttributes) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DeviceAttributes) SetModel(v string) {
	o.Model = &v
}

// GetAddedDate returns the AddedDate field value if set, zero value otherwise.
func (o *DeviceAttributes) GetAddedDate() time.Time {
	if o == nil || IsNil(o.AddedDate) {
		var ret time.Time
		return ret
	}
	return *o.AddedDate
}

// GetAddedDateOk returns a tuple with the AddedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetAddedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddedDate) {
		return nil, false
	}
	return o.AddedDate, true
}

// HasAddedDate returns a boolean if a field has been set.
func (o *DeviceAttributes) HasAddedDate() bool {
	if o != nil && !IsNil(o.AddedDate) {
		return true
	}

	return false
}

// SetAddedDate gets a reference to the given time.Time and assigns it to the AddedDate field.
func (o *DeviceAttributes) SetAddedDate(v time.Time) {
	o.AddedDate = &v
}

func (o DeviceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Udid) {
		toSerialize["udid"] = o.Udid
	}
	if !IsNil(o.DeviceClass) {
		toSerialize["deviceClass"] = o.DeviceClass
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.AddedDate) {
		toSerialize["addedDate"] = o.AddedDate
	}
	return toSerialize, nil
}

type NullableDeviceAttributes struct {
	value *DeviceAttributes
	isSet bool
}

func (v NullableDeviceAttributes) Get() *DeviceAttributes {
	return v.value
}

func (v *NullableDeviceAttributes) Set(val *DeviceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAttributes(val *DeviceAttributes) *NullableDeviceAttributes {
	return &NullableDeviceAttributes{value: val, isSet: true}
}

func (v NullableDeviceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
