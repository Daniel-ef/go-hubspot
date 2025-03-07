/*
Timeline events

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"encoding/json"
	"time"
)

// TimelineEvent The state of the timeline event.
type TimelineEvent struct {
	// The event template ID.
	EventTemplateId string `json:"eventTemplateId"`
	// The email address used for contact-specific events. This can be used to identify existing contacts, create new ones, or change the email for an existing contact (if paired with the `objectId`).
	Email *string `json:"email,omitempty"`
	// The CRM object identifier. This is required for every event other than contacts (where utk or email can be used).
	ObjectId *string `json:"objectId,omitempty"`
	// Use the `utk` parameter to associate an event with a contact by `usertoken`. This is recommended if you don't know a user's email, but have an identifying user token in your cookie.
	Utk *string `json:"utk,omitempty"`
	// The event domain (often paired with utk).
	Domain *string `json:"domain,omitempty"`
	// The time the event occurred. If not passed in, the curren time will be assumed. This is used to determine where an event is shown on a CRM object's timeline.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// A collection of token keys and values associated with the template tokens.
	Tokens map[string]string `json:"tokens"`
	// Additional event-specific data that can be interpreted by the template's markdown.
	ExtraData      map[string]interface{} `json:"extraData,omitempty"`
	TimelineIFrame *TimelineEventIFrame   `json:"timelineIFrame,omitempty"`
	// Identifier for the event. This is optional, and we recommend you do not pass this in. We will create one for you if you omit this. You can also use `{{uuid}}` anywhere in the ID to generate a unique string, guaranteeing uniqueness.
	Id *string `json:"id,omitempty"`
}

// NewTimelineEvent instantiates a new TimelineEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineEvent(eventTemplateId string, tokens map[string]string) *TimelineEvent {
	this := TimelineEvent{}
	this.EventTemplateId = eventTemplateId
	this.Tokens = tokens
	return &this
}

// NewTimelineEventWithDefaults instantiates a new TimelineEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineEventWithDefaults() *TimelineEvent {
	this := TimelineEvent{}
	return &this
}

// GetEventTemplateId returns the EventTemplateId field value
func (o *TimelineEvent) GetEventTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventTemplateId
}

// GetEventTemplateIdOk returns a tuple with the EventTemplateId field value
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetEventTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTemplateId, true
}

// SetEventTemplateId sets field value
func (o *TimelineEvent) SetEventTemplateId(v string) {
	o.EventTemplateId = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TimelineEvent) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TimelineEvent) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *TimelineEvent) SetEmail(v string) {
	o.Email = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *TimelineEvent) GetObjectId() string {
	if o == nil || isNil(o.ObjectId) {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetObjectIdOk() (*string, bool) {
	if o == nil || isNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *TimelineEvent) HasObjectId() bool {
	if o != nil && !isNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *TimelineEvent) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetUtk returns the Utk field value if set, zero value otherwise.
func (o *TimelineEvent) GetUtk() string {
	if o == nil || isNil(o.Utk) {
		var ret string
		return ret
	}
	return *o.Utk
}

// GetUtkOk returns a tuple with the Utk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetUtkOk() (*string, bool) {
	if o == nil || isNil(o.Utk) {
		return nil, false
	}
	return o.Utk, true
}

// HasUtk returns a boolean if a field has been set.
func (o *TimelineEvent) HasUtk() bool {
	if o != nil && !isNil(o.Utk) {
		return true
	}

	return false
}

// SetUtk gets a reference to the given string and assigns it to the Utk field.
func (o *TimelineEvent) SetUtk(v string) {
	o.Utk = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *TimelineEvent) GetDomain() string {
	if o == nil || isNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetDomainOk() (*string, bool) {
	if o == nil || isNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *TimelineEvent) HasDomain() bool {
	if o != nil && !isNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *TimelineEvent) SetDomain(v string) {
	o.Domain = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TimelineEvent) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TimelineEvent) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TimelineEvent) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTokens returns the Tokens field value
func (o *TimelineEvent) GetTokens() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetTokensOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokens, true
}

// SetTokens sets field value
func (o *TimelineEvent) SetTokens(v map[string]string) {
	o.Tokens = v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise.
func (o *TimelineEvent) GetExtraData() map[string]interface{} {
	if o == nil || isNil(o.ExtraData) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ExtraData) {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *TimelineEvent) HasExtraData() bool {
	if o != nil && !isNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]interface{} and assigns it to the ExtraData field.
func (o *TimelineEvent) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

// GetTimelineIFrame returns the TimelineIFrame field value if set, zero value otherwise.
func (o *TimelineEvent) GetTimelineIFrame() TimelineEventIFrame {
	if o == nil || isNil(o.TimelineIFrame) {
		var ret TimelineEventIFrame
		return ret
	}
	return *o.TimelineIFrame
}

// GetTimelineIFrameOk returns a tuple with the TimelineIFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetTimelineIFrameOk() (*TimelineEventIFrame, bool) {
	if o == nil || isNil(o.TimelineIFrame) {
		return nil, false
	}
	return o.TimelineIFrame, true
}

// HasTimelineIFrame returns a boolean if a field has been set.
func (o *TimelineEvent) HasTimelineIFrame() bool {
	if o != nil && !isNil(o.TimelineIFrame) {
		return true
	}

	return false
}

// SetTimelineIFrame gets a reference to the given TimelineEventIFrame and assigns it to the TimelineIFrame field.
func (o *TimelineEvent) SetTimelineIFrame(v TimelineEventIFrame) {
	o.TimelineIFrame = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimelineEvent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEvent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimelineEvent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TimelineEvent) SetId(v string) {
	o.Id = &v
}

func (o TimelineEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventTemplateId"] = o.EventTemplateId
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.ObjectId) {
		toSerialize["objectId"] = o.ObjectId
	}
	if !isNil(o.Utk) {
		toSerialize["utk"] = o.Utk
	}
	if !isNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["tokens"] = o.Tokens
	}
	if !isNil(o.ExtraData) {
		toSerialize["extraData"] = o.ExtraData
	}
	if !isNil(o.TimelineIFrame) {
		toSerialize["timelineIFrame"] = o.TimelineIFrame
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineEvent struct {
	value *TimelineEvent
	isSet bool
}

func (v NullableTimelineEvent) Get() *TimelineEvent {
	return v.value
}

func (v *NullableTimelineEvent) Set(val *TimelineEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineEvent(val *TimelineEvent) *NullableTimelineEvent {
	return &NullableTimelineEvent{value: val, isSet: true}
}

func (v NullableTimelineEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
