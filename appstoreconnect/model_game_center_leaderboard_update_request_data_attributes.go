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

// checks if the GameCenterLeaderboardUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardUpdateRequestDataAttributes{}

// GameCenterLeaderboardUpdateRequestDataAttributes struct for GameCenterLeaderboardUpdateRequestDataAttributes
type GameCenterLeaderboardUpdateRequestDataAttributes struct {
	DefaultFormatter    *GameCenterLeaderboardFormatter `json:"defaultFormatter,omitempty"`
	ReferenceName       *string                         `json:"referenceName,omitempty"`
	SubmissionType      *string                         `json:"submissionType,omitempty"`
	ScoreSortType       *string                         `json:"scoreSortType,omitempty"`
	ScoreRangeStart     *float64                        `json:"scoreRangeStart,omitempty"`
	ScoreRangeEnd       *float64                        `json:"scoreRangeEnd,omitempty"`
	RecurrenceStartDate *time.Time                      `json:"recurrenceStartDate,omitempty"`
	RecurrenceDuration  *string                         `json:"recurrenceDuration,omitempty"`
	RecurrenceRule      *string                         `json:"recurrenceRule,omitempty"`
	Archived            *bool                           `json:"archived,omitempty"`
}

// NewGameCenterLeaderboardUpdateRequestDataAttributes instantiates a new GameCenterLeaderboardUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardUpdateRequestDataAttributes() *GameCenterLeaderboardUpdateRequestDataAttributes {
	this := GameCenterLeaderboardUpdateRequestDataAttributes{}
	return &this
}

// NewGameCenterLeaderboardUpdateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardUpdateRequestDataAttributesWithDefaults() *GameCenterLeaderboardUpdateRequestDataAttributes {
	this := GameCenterLeaderboardUpdateRequestDataAttributes{}
	return &this
}

// GetDefaultFormatter returns the DefaultFormatter field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetDefaultFormatter() GameCenterLeaderboardFormatter {
	if o == nil || IsNil(o.DefaultFormatter) {
		var ret GameCenterLeaderboardFormatter
		return ret
	}
	return *o.DefaultFormatter
}

// GetDefaultFormatterOk returns a tuple with the DefaultFormatter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetDefaultFormatterOk() (*GameCenterLeaderboardFormatter, bool) {
	if o == nil || IsNil(o.DefaultFormatter) {
		return nil, false
	}
	return o.DefaultFormatter, true
}

// HasDefaultFormatter returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasDefaultFormatter() bool {
	if o != nil && !IsNil(o.DefaultFormatter) {
		return true
	}

	return false
}

// SetDefaultFormatter gets a reference to the given GameCenterLeaderboardFormatter and assigns it to the DefaultFormatter field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetDefaultFormatter(v GameCenterLeaderboardFormatter) {
	o.DefaultFormatter = &v
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetSubmissionType returns the SubmissionType field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetSubmissionType() string {
	if o == nil || IsNil(o.SubmissionType) {
		var ret string
		return ret
	}
	return *o.SubmissionType
}

// GetSubmissionTypeOk returns a tuple with the SubmissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetSubmissionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubmissionType) {
		return nil, false
	}
	return o.SubmissionType, true
}

// HasSubmissionType returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasSubmissionType() bool {
	if o != nil && !IsNil(o.SubmissionType) {
		return true
	}

	return false
}

// SetSubmissionType gets a reference to the given string and assigns it to the SubmissionType field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetSubmissionType(v string) {
	o.SubmissionType = &v
}

// GetScoreSortType returns the ScoreSortType field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreSortType() string {
	if o == nil || IsNil(o.ScoreSortType) {
		var ret string
		return ret
	}
	return *o.ScoreSortType
}

// GetScoreSortTypeOk returns a tuple with the ScoreSortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScoreSortType) {
		return nil, false
	}
	return o.ScoreSortType, true
}

// HasScoreSortType returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasScoreSortType() bool {
	if o != nil && !IsNil(o.ScoreSortType) {
		return true
	}

	return false
}

// SetScoreSortType gets a reference to the given string and assigns it to the ScoreSortType field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetScoreSortType(v string) {
	o.ScoreSortType = &v
}

// GetScoreRangeStart returns the ScoreRangeStart field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeStart() float64 {
	if o == nil || IsNil(o.ScoreRangeStart) {
		var ret float64
		return ret
	}
	return *o.ScoreRangeStart
}

// GetScoreRangeStartOk returns a tuple with the ScoreRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeStartOk() (*float64, bool) {
	if o == nil || IsNil(o.ScoreRangeStart) {
		return nil, false
	}
	return o.ScoreRangeStart, true
}

// HasScoreRangeStart returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasScoreRangeStart() bool {
	if o != nil && !IsNil(o.ScoreRangeStart) {
		return true
	}

	return false
}

// SetScoreRangeStart gets a reference to the given float64 and assigns it to the ScoreRangeStart field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetScoreRangeStart(v float64) {
	o.ScoreRangeStart = &v
}

// GetScoreRangeEnd returns the ScoreRangeEnd field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeEnd() float64 {
	if o == nil || IsNil(o.ScoreRangeEnd) {
		var ret float64
		return ret
	}
	return *o.ScoreRangeEnd
}

// GetScoreRangeEndOk returns a tuple with the ScoreRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeEndOk() (*float64, bool) {
	if o == nil || IsNil(o.ScoreRangeEnd) {
		return nil, false
	}
	return o.ScoreRangeEnd, true
}

// HasScoreRangeEnd returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasScoreRangeEnd() bool {
	if o != nil && !IsNil(o.ScoreRangeEnd) {
		return true
	}

	return false
}

// SetScoreRangeEnd gets a reference to the given float64 and assigns it to the ScoreRangeEnd field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetScoreRangeEnd(v float64) {
	o.ScoreRangeEnd = &v
}

// GetRecurrenceStartDate returns the RecurrenceStartDate field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceStartDate() time.Time {
	if o == nil || IsNil(o.RecurrenceStartDate) {
		var ret time.Time
		return ret
	}
	return *o.RecurrenceStartDate
}

// GetRecurrenceStartDateOk returns a tuple with the RecurrenceStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RecurrenceStartDate) {
		return nil, false
	}
	return o.RecurrenceStartDate, true
}

// HasRecurrenceStartDate returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasRecurrenceStartDate() bool {
	if o != nil && !IsNil(o.RecurrenceStartDate) {
		return true
	}

	return false
}

// SetRecurrenceStartDate gets a reference to the given time.Time and assigns it to the RecurrenceStartDate field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetRecurrenceStartDate(v time.Time) {
	o.RecurrenceStartDate = &v
}

// GetRecurrenceDuration returns the RecurrenceDuration field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceDuration() string {
	if o == nil || IsNil(o.RecurrenceDuration) {
		var ret string
		return ret
	}
	return *o.RecurrenceDuration
}

// GetRecurrenceDurationOk returns a tuple with the RecurrenceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceDurationOk() (*string, bool) {
	if o == nil || IsNil(o.RecurrenceDuration) {
		return nil, false
	}
	return o.RecurrenceDuration, true
}

// HasRecurrenceDuration returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasRecurrenceDuration() bool {
	if o != nil && !IsNil(o.RecurrenceDuration) {
		return true
	}

	return false
}

// SetRecurrenceDuration gets a reference to the given string and assigns it to the RecurrenceDuration field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetRecurrenceDuration(v string) {
	o.RecurrenceDuration = &v
}

// GetRecurrenceRule returns the RecurrenceRule field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceRule() string {
	if o == nil || IsNil(o.RecurrenceRule) {
		var ret string
		return ret
	}
	return *o.RecurrenceRule
}

// GetRecurrenceRuleOk returns a tuple with the RecurrenceRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceRuleOk() (*string, bool) {
	if o == nil || IsNil(o.RecurrenceRule) {
		return nil, false
	}
	return o.RecurrenceRule, true
}

// HasRecurrenceRule returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasRecurrenceRule() bool {
	if o != nil && !IsNil(o.RecurrenceRule) {
		return true
	}

	return false
}

// SetRecurrenceRule gets a reference to the given string and assigns it to the RecurrenceRule field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetRecurrenceRule(v string) {
	o.RecurrenceRule = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetArchived(v bool) {
	o.Archived = &v
}

func (o GameCenterLeaderboardUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultFormatter) {
		toSerialize["defaultFormatter"] = o.DefaultFormatter
	}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if !IsNil(o.SubmissionType) {
		toSerialize["submissionType"] = o.SubmissionType
	}
	if !IsNil(o.ScoreSortType) {
		toSerialize["scoreSortType"] = o.ScoreSortType
	}
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
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardUpdateRequestDataAttributes struct {
	value *GameCenterLeaderboardUpdateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterLeaderboardUpdateRequestDataAttributes) Get() *GameCenterLeaderboardUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterLeaderboardUpdateRequestDataAttributes) Set(val *GameCenterLeaderboardUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardUpdateRequestDataAttributes(val *GameCenterLeaderboardUpdateRequestDataAttributes) *NullableGameCenterLeaderboardUpdateRequestDataAttributes {
	return &NullableGameCenterLeaderboardUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
