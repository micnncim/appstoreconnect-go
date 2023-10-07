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

// checks if the GameCenterLeaderboardCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardCreateRequestDataAttributes{}

// GameCenterLeaderboardCreateRequestDataAttributes struct for GameCenterLeaderboardCreateRequestDataAttributes
type GameCenterLeaderboardCreateRequestDataAttributes struct {
	DefaultFormatter    GameCenterLeaderboardFormatter `json:"defaultFormatter"`
	ReferenceName       string                         `json:"referenceName"`
	VendorIdentifier    string                         `json:"vendorIdentifier"`
	SubmissionType      string                         `json:"submissionType"`
	ScoreSortType       string                         `json:"scoreSortType"`
	ScoreRangeStart     *float64                       `json:"scoreRangeStart,omitempty"`
	ScoreRangeEnd       *float64                       `json:"scoreRangeEnd,omitempty"`
	RecurrenceStartDate *time.Time                     `json:"recurrenceStartDate,omitempty"`
	RecurrenceDuration  *string                        `json:"recurrenceDuration,omitempty"`
	RecurrenceRule      *string                        `json:"recurrenceRule,omitempty"`
}

// NewGameCenterLeaderboardCreateRequestDataAttributes instantiates a new GameCenterLeaderboardCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardCreateRequestDataAttributes(defaultFormatter GameCenterLeaderboardFormatter, referenceName string, vendorIdentifier string, submissionType string, scoreSortType string) *GameCenterLeaderboardCreateRequestDataAttributes {
	this := GameCenterLeaderboardCreateRequestDataAttributes{}
	this.DefaultFormatter = defaultFormatter
	this.ReferenceName = referenceName
	this.VendorIdentifier = vendorIdentifier
	this.SubmissionType = submissionType
	this.ScoreSortType = scoreSortType
	return &this
}

// NewGameCenterLeaderboardCreateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardCreateRequestDataAttributesWithDefaults() *GameCenterLeaderboardCreateRequestDataAttributes {
	this := GameCenterLeaderboardCreateRequestDataAttributes{}
	return &this
}

// GetDefaultFormatter returns the DefaultFormatter field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetDefaultFormatter() GameCenterLeaderboardFormatter {
	if o == nil {
		var ret GameCenterLeaderboardFormatter
		return ret
	}

	return o.DefaultFormatter
}

// GetDefaultFormatterOk returns a tuple with the DefaultFormatter field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetDefaultFormatterOk() (*GameCenterLeaderboardFormatter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultFormatter, true
}

// SetDefaultFormatter sets field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetDefaultFormatter(v GameCenterLeaderboardFormatter) {
	o.DefaultFormatter = v
}

// GetReferenceName returns the ReferenceName field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetVendorIdentifier returns the VendorIdentifier field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetVendorIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorIdentifier
}

// GetVendorIdentifierOk returns a tuple with the VendorIdentifier field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetVendorIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorIdentifier, true
}

// SetVendorIdentifier sets field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetVendorIdentifier(v string) {
	o.VendorIdentifier = v
}

// GetSubmissionType returns the SubmissionType field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetSubmissionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubmissionType
}

// GetSubmissionTypeOk returns a tuple with the SubmissionType field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetSubmissionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmissionType, true
}

// SetSubmissionType sets field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetSubmissionType(v string) {
	o.SubmissionType = v
}

// GetScoreSortType returns the ScoreSortType field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreSortType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScoreSortType
}

// GetScoreSortTypeOk returns a tuple with the ScoreSortType field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreSortTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScoreSortType, true
}

// SetScoreSortType sets field value
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetScoreSortType(v string) {
	o.ScoreSortType = v
}

// GetScoreRangeStart returns the ScoreRangeStart field value if set, zero value otherwise.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeStart() float64 {
	if o == nil || IsNil(o.ScoreRangeStart) {
		var ret float64
		return ret
	}
	return *o.ScoreRangeStart
}

// GetScoreRangeStartOk returns a tuple with the ScoreRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeStartOk() (*float64, bool) {
	if o == nil || IsNil(o.ScoreRangeStart) {
		return nil, false
	}
	return o.ScoreRangeStart, true
}

// HasScoreRangeStart returns a boolean if a field has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasScoreRangeStart() bool {
	if o != nil && !IsNil(o.ScoreRangeStart) {
		return true
	}

	return false
}

// SetScoreRangeStart gets a reference to the given float64 and assigns it to the ScoreRangeStart field.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetScoreRangeStart(v float64) {
	o.ScoreRangeStart = &v
}

// GetScoreRangeEnd returns the ScoreRangeEnd field value if set, zero value otherwise.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeEnd() float64 {
	if o == nil || IsNil(o.ScoreRangeEnd) {
		var ret float64
		return ret
	}
	return *o.ScoreRangeEnd
}

// GetScoreRangeEndOk returns a tuple with the ScoreRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeEndOk() (*float64, bool) {
	if o == nil || IsNil(o.ScoreRangeEnd) {
		return nil, false
	}
	return o.ScoreRangeEnd, true
}

// HasScoreRangeEnd returns a boolean if a field has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasScoreRangeEnd() bool {
	if o != nil && !IsNil(o.ScoreRangeEnd) {
		return true
	}

	return false
}

// SetScoreRangeEnd gets a reference to the given float64 and assigns it to the ScoreRangeEnd field.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetScoreRangeEnd(v float64) {
	o.ScoreRangeEnd = &v
}

// GetRecurrenceStartDate returns the RecurrenceStartDate field value if set, zero value otherwise.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceStartDate() time.Time {
	if o == nil || IsNil(o.RecurrenceStartDate) {
		var ret time.Time
		return ret
	}
	return *o.RecurrenceStartDate
}

// GetRecurrenceStartDateOk returns a tuple with the RecurrenceStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RecurrenceStartDate) {
		return nil, false
	}
	return o.RecurrenceStartDate, true
}

// HasRecurrenceStartDate returns a boolean if a field has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasRecurrenceStartDate() bool {
	if o != nil && !IsNil(o.RecurrenceStartDate) {
		return true
	}

	return false
}

// SetRecurrenceStartDate gets a reference to the given time.Time and assigns it to the RecurrenceStartDate field.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetRecurrenceStartDate(v time.Time) {
	o.RecurrenceStartDate = &v
}

// GetRecurrenceDuration returns the RecurrenceDuration field value if set, zero value otherwise.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceDuration() string {
	if o == nil || IsNil(o.RecurrenceDuration) {
		var ret string
		return ret
	}
	return *o.RecurrenceDuration
}

// GetRecurrenceDurationOk returns a tuple with the RecurrenceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceDurationOk() (*string, bool) {
	if o == nil || IsNil(o.RecurrenceDuration) {
		return nil, false
	}
	return o.RecurrenceDuration, true
}

// HasRecurrenceDuration returns a boolean if a field has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasRecurrenceDuration() bool {
	if o != nil && !IsNil(o.RecurrenceDuration) {
		return true
	}

	return false
}

// SetRecurrenceDuration gets a reference to the given string and assigns it to the RecurrenceDuration field.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetRecurrenceDuration(v string) {
	o.RecurrenceDuration = &v
}

// GetRecurrenceRule returns the RecurrenceRule field value if set, zero value otherwise.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceRule() string {
	if o == nil || IsNil(o.RecurrenceRule) {
		var ret string
		return ret
	}
	return *o.RecurrenceRule
}

// GetRecurrenceRuleOk returns a tuple with the RecurrenceRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceRuleOk() (*string, bool) {
	if o == nil || IsNil(o.RecurrenceRule) {
		return nil, false
	}
	return o.RecurrenceRule, true
}

// HasRecurrenceRule returns a boolean if a field has been set.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasRecurrenceRule() bool {
	if o != nil && !IsNil(o.RecurrenceRule) {
		return true
	}

	return false
}

// SetRecurrenceRule gets a reference to the given string and assigns it to the RecurrenceRule field.
func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetRecurrenceRule(v string) {
	o.RecurrenceRule = &v
}

func (o GameCenterLeaderboardCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultFormatter"] = o.DefaultFormatter
	toSerialize["referenceName"] = o.ReferenceName
	toSerialize["vendorIdentifier"] = o.VendorIdentifier
	toSerialize["submissionType"] = o.SubmissionType
	toSerialize["scoreSortType"] = o.ScoreSortType
	if !IsNil(o.ScoreRangeStart) {
		toSerialize["scoreRangeStart"] = o.ScoreRangeStart
	}
	if !IsNil(o.ScoreRangeEnd) {
		toSerialize["scoreRangeEnd"] = o.ScoreRangeEnd
	}
	if !IsNil(o.RecurrenceStartDate) {
		toSerialize["recurrenceStartDate"] = o.RecurrenceStartDate
	}
	if !IsNil(o.RecurrenceDuration) {
		toSerialize["recurrenceDuration"] = o.RecurrenceDuration
	}
	if !IsNil(o.RecurrenceRule) {
		toSerialize["recurrenceRule"] = o.RecurrenceRule
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardCreateRequestDataAttributes struct {
	value *GameCenterLeaderboardCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterLeaderboardCreateRequestDataAttributes) Get() *GameCenterLeaderboardCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterLeaderboardCreateRequestDataAttributes) Set(val *GameCenterLeaderboardCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardCreateRequestDataAttributes(val *GameCenterLeaderboardCreateRequestDataAttributes) *NullableGameCenterLeaderboardCreateRequestDataAttributes {
	return &NullableGameCenterLeaderboardCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
