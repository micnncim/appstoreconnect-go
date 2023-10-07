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

// checks if the AppEncryptionDeclarationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationAttributes{}

// AppEncryptionDeclarationAttributes struct for AppEncryptionDeclarationAttributes
type AppEncryptionDeclarationAttributes struct {
	AppDescription *string    `json:"appDescription,omitempty"`
	CreatedDate    *time.Time `json:"createdDate,omitempty"`
	// Deprecated
	UsesEncryption                  *bool     `json:"usesEncryption,omitempty"`
	Exempt                          *bool     `json:"exempt,omitempty"`
	ContainsProprietaryCryptography *bool     `json:"containsProprietaryCryptography,omitempty"`
	ContainsThirdPartyCryptography  *bool     `json:"containsThirdPartyCryptography,omitempty"`
	AvailableOnFrenchStore          *bool     `json:"availableOnFrenchStore,omitempty"`
	Platform                        *Platform `json:"platform,omitempty"`
	// Deprecated
	UploadedDate *time.Time `json:"uploadedDate,omitempty"`
	// Deprecated
	DocumentUrl *string `json:"documentUrl,omitempty"`
	// Deprecated
	DocumentName *string `json:"documentName,omitempty"`
	// Deprecated
	DocumentType                  *string                        `json:"documentType,omitempty"`
	AppEncryptionDeclarationState *AppEncryptionDeclarationState `json:"appEncryptionDeclarationState,omitempty"`
	CodeValue                     *string                        `json:"codeValue,omitempty"`
}

// NewAppEncryptionDeclarationAttributes instantiates a new AppEncryptionDeclarationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationAttributes() *AppEncryptionDeclarationAttributes {
	this := AppEncryptionDeclarationAttributes{}
	return &this
}

// NewAppEncryptionDeclarationAttributesWithDefaults instantiates a new AppEncryptionDeclarationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationAttributesWithDefaults() *AppEncryptionDeclarationAttributes {
	this := AppEncryptionDeclarationAttributes{}
	return &this
}

// GetAppDescription returns the AppDescription field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetAppDescription() string {
	if o == nil || IsNil(o.AppDescription) {
		var ret string
		return ret
	}
	return *o.AppDescription
}

// GetAppDescriptionOk returns a tuple with the AppDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetAppDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AppDescription) {
		return nil, false
	}
	return o.AppDescription, true
}

// HasAppDescription returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasAppDescription() bool {
	if o != nil && !IsNil(o.AppDescription) {
		return true
	}

	return false
}

// SetAppDescription gets a reference to the given string and assigns it to the AppDescription field.
func (o *AppEncryptionDeclarationAttributes) SetAppDescription(v string) {
	o.AppDescription = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *AppEncryptionDeclarationAttributes) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetUsesEncryption returns the UsesEncryption field value if set, zero value otherwise.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetUsesEncryption() bool {
	if o == nil || IsNil(o.UsesEncryption) {
		var ret bool
		return ret
	}
	return *o.UsesEncryption
}

// GetUsesEncryptionOk returns a tuple with the UsesEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetUsesEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesEncryption) {
		return nil, false
	}
	return o.UsesEncryption, true
}

// HasUsesEncryption returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasUsesEncryption() bool {
	if o != nil && !IsNil(o.UsesEncryption) {
		return true
	}

	return false
}

// SetUsesEncryption gets a reference to the given bool and assigns it to the UsesEncryption field.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) SetUsesEncryption(v bool) {
	o.UsesEncryption = &v
}

// GetExempt returns the Exempt field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetExempt() bool {
	if o == nil || IsNil(o.Exempt) {
		var ret bool
		return ret
	}
	return *o.Exempt
}

// GetExemptOk returns a tuple with the Exempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.Exempt) {
		return nil, false
	}
	return o.Exempt, true
}

// HasExempt returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasExempt() bool {
	if o != nil && !IsNil(o.Exempt) {
		return true
	}

	return false
}

// SetExempt gets a reference to the given bool and assigns it to the Exempt field.
func (o *AppEncryptionDeclarationAttributes) SetExempt(v bool) {
	o.Exempt = &v
}

// GetContainsProprietaryCryptography returns the ContainsProprietaryCryptography field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetContainsProprietaryCryptography() bool {
	if o == nil || IsNil(o.ContainsProprietaryCryptography) {
		var ret bool
		return ret
	}
	return *o.ContainsProprietaryCryptography
}

// GetContainsProprietaryCryptographyOk returns a tuple with the ContainsProprietaryCryptography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetContainsProprietaryCryptographyOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainsProprietaryCryptography) {
		return nil, false
	}
	return o.ContainsProprietaryCryptography, true
}

// HasContainsProprietaryCryptography returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasContainsProprietaryCryptography() bool {
	if o != nil && !IsNil(o.ContainsProprietaryCryptography) {
		return true
	}

	return false
}

// SetContainsProprietaryCryptography gets a reference to the given bool and assigns it to the ContainsProprietaryCryptography field.
func (o *AppEncryptionDeclarationAttributes) SetContainsProprietaryCryptography(v bool) {
	o.ContainsProprietaryCryptography = &v
}

// GetContainsThirdPartyCryptography returns the ContainsThirdPartyCryptography field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetContainsThirdPartyCryptography() bool {
	if o == nil || IsNil(o.ContainsThirdPartyCryptography) {
		var ret bool
		return ret
	}
	return *o.ContainsThirdPartyCryptography
}

// GetContainsThirdPartyCryptographyOk returns a tuple with the ContainsThirdPartyCryptography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetContainsThirdPartyCryptographyOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainsThirdPartyCryptography) {
		return nil, false
	}
	return o.ContainsThirdPartyCryptography, true
}

// HasContainsThirdPartyCryptography returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasContainsThirdPartyCryptography() bool {
	if o != nil && !IsNil(o.ContainsThirdPartyCryptography) {
		return true
	}

	return false
}

// SetContainsThirdPartyCryptography gets a reference to the given bool and assigns it to the ContainsThirdPartyCryptography field.
func (o *AppEncryptionDeclarationAttributes) SetContainsThirdPartyCryptography(v bool) {
	o.ContainsThirdPartyCryptography = &v
}

// GetAvailableOnFrenchStore returns the AvailableOnFrenchStore field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetAvailableOnFrenchStore() bool {
	if o == nil || IsNil(o.AvailableOnFrenchStore) {
		var ret bool
		return ret
	}
	return *o.AvailableOnFrenchStore
}

// GetAvailableOnFrenchStoreOk returns a tuple with the AvailableOnFrenchStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetAvailableOnFrenchStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailableOnFrenchStore) {
		return nil, false
	}
	return o.AvailableOnFrenchStore, true
}

// HasAvailableOnFrenchStore returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasAvailableOnFrenchStore() bool {
	if o != nil && !IsNil(o.AvailableOnFrenchStore) {
		return true
	}

	return false
}

// SetAvailableOnFrenchStore gets a reference to the given bool and assigns it to the AvailableOnFrenchStore field.
func (o *AppEncryptionDeclarationAttributes) SetAvailableOnFrenchStore(v bool) {
	o.AvailableOnFrenchStore = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *AppEncryptionDeclarationAttributes) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetUploadedDate returns the UploadedDate field value if set, zero value otherwise.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetUploadedDate() time.Time {
	if o == nil || IsNil(o.UploadedDate) {
		var ret time.Time
		return ret
	}
	return *o.UploadedDate
}

// GetUploadedDateOk returns a tuple with the UploadedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetUploadedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UploadedDate) {
		return nil, false
	}
	return o.UploadedDate, true
}

// HasUploadedDate returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasUploadedDate() bool {
	if o != nil && !IsNil(o.UploadedDate) {
		return true
	}

	return false
}

// SetUploadedDate gets a reference to the given time.Time and assigns it to the UploadedDate field.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) SetUploadedDate(v time.Time) {
	o.UploadedDate = &v
}

// GetDocumentUrl returns the DocumentUrl field value if set, zero value otherwise.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetDocumentUrl() string {
	if o == nil || IsNil(o.DocumentUrl) {
		var ret string
		return ret
	}
	return *o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetDocumentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentUrl) {
		return nil, false
	}
	return o.DocumentUrl, true
}

// HasDocumentUrl returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasDocumentUrl() bool {
	if o != nil && !IsNil(o.DocumentUrl) {
		return true
	}

	return false
}

// SetDocumentUrl gets a reference to the given string and assigns it to the DocumentUrl field.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) SetDocumentUrl(v string) {
	o.DocumentUrl = &v
}

// GetDocumentName returns the DocumentName field value if set, zero value otherwise.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetDocumentName() string {
	if o == nil || IsNil(o.DocumentName) {
		var ret string
		return ret
	}
	return *o.DocumentName
}

// GetDocumentNameOk returns a tuple with the DocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetDocumentNameOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentName) {
		return nil, false
	}
	return o.DocumentName, true
}

// HasDocumentName returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasDocumentName() bool {
	if o != nil && !IsNil(o.DocumentName) {
		return true
	}

	return false
}

// SetDocumentName gets a reference to the given string and assigns it to the DocumentName field.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) SetDocumentName(v string) {
	o.DocumentName = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) GetDocumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
// Deprecated
func (o *AppEncryptionDeclarationAttributes) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetAppEncryptionDeclarationState returns the AppEncryptionDeclarationState field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetAppEncryptionDeclarationState() AppEncryptionDeclarationState {
	if o == nil || IsNil(o.AppEncryptionDeclarationState) {
		var ret AppEncryptionDeclarationState
		return ret
	}
	return *o.AppEncryptionDeclarationState
}

// GetAppEncryptionDeclarationStateOk returns a tuple with the AppEncryptionDeclarationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetAppEncryptionDeclarationStateOk() (*AppEncryptionDeclarationState, bool) {
	if o == nil || IsNil(o.AppEncryptionDeclarationState) {
		return nil, false
	}
	return o.AppEncryptionDeclarationState, true
}

// HasAppEncryptionDeclarationState returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasAppEncryptionDeclarationState() bool {
	if o != nil && !IsNil(o.AppEncryptionDeclarationState) {
		return true
	}

	return false
}

// SetAppEncryptionDeclarationState gets a reference to the given AppEncryptionDeclarationState and assigns it to the AppEncryptionDeclarationState field.
func (o *AppEncryptionDeclarationAttributes) SetAppEncryptionDeclarationState(v AppEncryptionDeclarationState) {
	o.AppEncryptionDeclarationState = &v
}

// GetCodeValue returns the CodeValue field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationAttributes) GetCodeValue() string {
	if o == nil || IsNil(o.CodeValue) {
		var ret string
		return ret
	}
	return *o.CodeValue
}

// GetCodeValueOk returns a tuple with the CodeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationAttributes) GetCodeValueOk() (*string, bool) {
	if o == nil || IsNil(o.CodeValue) {
		return nil, false
	}
	return o.CodeValue, true
}

// HasCodeValue returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationAttributes) HasCodeValue() bool {
	if o != nil && !IsNil(o.CodeValue) {
		return true
	}

	return false
}

// SetCodeValue gets a reference to the given string and assigns it to the CodeValue field.
func (o *AppEncryptionDeclarationAttributes) SetCodeValue(v string) {
	o.CodeValue = &v
}

func (o AppEncryptionDeclarationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppDescription) {
		toSerialize["appDescription"] = o.AppDescription
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.UsesEncryption) {
		toSerialize["usesEncryption"] = o.UsesEncryption
	}
	if !IsNil(o.Exempt) {
		toSerialize["exempt"] = o.Exempt
	}
	if !IsNil(o.ContainsProprietaryCryptography) {
		toSerialize["containsProprietaryCryptography"] = o.ContainsProprietaryCryptography
	}
	if !IsNil(o.ContainsThirdPartyCryptography) {
		toSerialize["containsThirdPartyCryptography"] = o.ContainsThirdPartyCryptography
	}
	if !IsNil(o.AvailableOnFrenchStore) {
		toSerialize["availableOnFrenchStore"] = o.AvailableOnFrenchStore
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.UploadedDate) {
		toSerialize["uploadedDate"] = o.UploadedDate
	}
	if !IsNil(o.DocumentUrl) {
		toSerialize["documentUrl"] = o.DocumentUrl
	}
	if !IsNil(o.DocumentName) {
		toSerialize["documentName"] = o.DocumentName
	}
	if !IsNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !IsNil(o.AppEncryptionDeclarationState) {
		toSerialize["appEncryptionDeclarationState"] = o.AppEncryptionDeclarationState
	}
	if !IsNil(o.CodeValue) {
		toSerialize["codeValue"] = o.CodeValue
	}
	return toSerialize, nil
}

type NullableAppEncryptionDeclarationAttributes struct {
	value *AppEncryptionDeclarationAttributes
	isSet bool
}

func (v NullableAppEncryptionDeclarationAttributes) Get() *AppEncryptionDeclarationAttributes {
	return v.value
}

func (v *NullableAppEncryptionDeclarationAttributes) Set(val *AppEncryptionDeclarationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationAttributes(val *AppEncryptionDeclarationAttributes) *NullableAppEncryptionDeclarationAttributes {
	return &NullableAppEncryptionDeclarationAttributes{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
