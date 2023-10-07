/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
	"fmt"
)

// AppClipAdvancedExperienceLanguage the model 'AppClipAdvancedExperienceLanguage'
type AppClipAdvancedExperienceLanguage string

// List of AppClipAdvancedExperienceLanguage
const (
	APPCLIPADVANCEDEXPERIENCELANGUAGE_AR AppClipAdvancedExperienceLanguage = "AR"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_CA AppClipAdvancedExperienceLanguage = "CA"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_CS AppClipAdvancedExperienceLanguage = "CS"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_DA AppClipAdvancedExperienceLanguage = "DA"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_DE AppClipAdvancedExperienceLanguage = "DE"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_EL AppClipAdvancedExperienceLanguage = "EL"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_EN AppClipAdvancedExperienceLanguage = "EN"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_ES AppClipAdvancedExperienceLanguage = "ES"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_FI AppClipAdvancedExperienceLanguage = "FI"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_FR AppClipAdvancedExperienceLanguage = "FR"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_HE AppClipAdvancedExperienceLanguage = "HE"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_HI AppClipAdvancedExperienceLanguage = "HI"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_HR AppClipAdvancedExperienceLanguage = "HR"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_HU AppClipAdvancedExperienceLanguage = "HU"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_ID AppClipAdvancedExperienceLanguage = "ID"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_IT AppClipAdvancedExperienceLanguage = "IT"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_JA AppClipAdvancedExperienceLanguage = "JA"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_KO AppClipAdvancedExperienceLanguage = "KO"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_MS AppClipAdvancedExperienceLanguage = "MS"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_NL AppClipAdvancedExperienceLanguage = "NL"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_NO AppClipAdvancedExperienceLanguage = "NO"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_PL AppClipAdvancedExperienceLanguage = "PL"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_PT AppClipAdvancedExperienceLanguage = "PT"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_RO AppClipAdvancedExperienceLanguage = "RO"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_RU AppClipAdvancedExperienceLanguage = "RU"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_SK AppClipAdvancedExperienceLanguage = "SK"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_SV AppClipAdvancedExperienceLanguage = "SV"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_TH AppClipAdvancedExperienceLanguage = "TH"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_TR AppClipAdvancedExperienceLanguage = "TR"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_UK AppClipAdvancedExperienceLanguage = "UK"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_VI AppClipAdvancedExperienceLanguage = "VI"
	APPCLIPADVANCEDEXPERIENCELANGUAGE_ZH AppClipAdvancedExperienceLanguage = "ZH"
)

// All allowed values of AppClipAdvancedExperienceLanguage enum
var AllowedAppClipAdvancedExperienceLanguageEnumValues = []AppClipAdvancedExperienceLanguage{
	"AR",
	"CA",
	"CS",
	"DA",
	"DE",
	"EL",
	"EN",
	"ES",
	"FI",
	"FR",
	"HE",
	"HI",
	"HR",
	"HU",
	"ID",
	"IT",
	"JA",
	"KO",
	"MS",
	"NL",
	"NO",
	"PL",
	"PT",
	"RO",
	"RU",
	"SK",
	"SV",
	"TH",
	"TR",
	"UK",
	"VI",
	"ZH",
}

func (v *AppClipAdvancedExperienceLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppClipAdvancedExperienceLanguage(value)
	for _, existing := range AllowedAppClipAdvancedExperienceLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppClipAdvancedExperienceLanguage", value)
}

// NewAppClipAdvancedExperienceLanguageFromValue returns a pointer to a valid AppClipAdvancedExperienceLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppClipAdvancedExperienceLanguageFromValue(v string) (*AppClipAdvancedExperienceLanguage, error) {
	ev := AppClipAdvancedExperienceLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppClipAdvancedExperienceLanguage: valid values are %v", v, AllowedAppClipAdvancedExperienceLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppClipAdvancedExperienceLanguage) IsValid() bool {
	for _, existing := range AllowedAppClipAdvancedExperienceLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppClipAdvancedExperienceLanguage value
func (v AppClipAdvancedExperienceLanguage) Ptr() *AppClipAdvancedExperienceLanguage {
	return &v
}

type NullableAppClipAdvancedExperienceLanguage struct {
	value *AppClipAdvancedExperienceLanguage
	isSet bool
}

func (v NullableAppClipAdvancedExperienceLanguage) Get() *AppClipAdvancedExperienceLanguage {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceLanguage) Set(val *AppClipAdvancedExperienceLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceLanguage(val *AppClipAdvancedExperienceLanguage) *NullableAppClipAdvancedExperienceLanguage {
	return &NullableAppClipAdvancedExperienceLanguage{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
