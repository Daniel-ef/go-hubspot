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

// TimelineEventResponse The current state of the timeline event.
type TimelineEventResponse struct {
	// Identifier for the event. This should be unique to the app and event template. If you use the same ID for different CRM objects, the last to be processed will win and the first will not have a record. You can also use `{{uuid}}` anywhere in the ID to generate a unique string, guaranteeing uniqueness.
	Id string `json:"id"`
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
	// The ObjectType associated with the EventTemplate.
	ObjectType string     `json:"objectType"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
}

// NewTimelineEventResponse instantiates a new TimelineEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineEventResponse(id string, eventTemplateId string, tokens map[string]string, objectType string) *TimelineEventResponse {
	this := TimelineEventResponse{}
	this.Id = id
	this.EventTemplateId = eventTemplateId
	this.Tokens = tokens
	this.ObjectType = objectType
	return &this
}

// NewTimelineEventResponseWithDefaults instantiates a new TimelineEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineEventResponseWithDefaults() *TimelineEventResponse {
	this := TimelineEventResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TimelineEventResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TimelineEventResponse) SetId(v string) {
	o.Id = v
}

// GetEventTemplateId returns the EventTemplateId field value
func (o *TimelineEventResponse) GetEventTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventTemplateId
}

// GetEventTemplateIdOk returns a tuple with the EventTemplateId field value
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetEventTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTemplateId, true
}

// SetEventTemplateId sets field value
func (o *TimelineEventResponse) SetEventTemplateId(v string) {
	o.EventTemplateId = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *TimelineEventResponse) SetEmail(v string) {
	o.Email = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetObjectId() string {
	if o == nil || isNil(o.ObjectId) {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetObjectIdOk() (*string, bool) {
	if o == nil || isNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasObjectId() bool {
	if o != nil && !isNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *TimelineEventResponse) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetUtk returns the Utk field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetUtk() string {
	if o == nil || isNil(o.Utk) {
		var ret string
		return ret
	}
	return *o.Utk
}

// GetUtkOk returns a tuple with the Utk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetUtkOk() (*string, bool) {
	if o == nil || isNil(o.Utk) {
		return nil, false
	}
	return o.Utk, true
}

// HasUtk returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasUtk() bool {
	if o != nil && !isNil(o.Utk) {
		return true
	}

	return false
}

// SetUtk gets a reference to the given string and assigns it to the Utk field.
func (o *TimelineEventResponse) SetUtk(v string) {
	o.Utk = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetDomain() string {
	if o == nil || isNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetDomainOk() (*string, bool) {
	if o == nil || isNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasDomain() bool {
	if o != nil && !isNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *TimelineEventResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TimelineEventResponse) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTokens returns the Tokens field value
func (o *TimelineEventResponse) GetTokens() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetTokensOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokens, true
}

// SetTokens sets field value
func (o *TimelineEventResponse) SetTokens(v map[string]string) {
	o.Tokens = v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetExtraData() map[string]interface{} {
	if o == nil || isNil(o.ExtraData) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetExtraDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ExtraData) {
		return map[string]interface{}{}, false
	}
	return o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasExtraData() bool {
	if o != nil && !isNil(o.ExtraData) {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given map[string]interface{} and assigns it to the ExtraData field.
func (o *TimelineEventResponse) SetExtraData(v map[string]interface{}) {
	o.ExtraData = v
}

// GetTimelineIFrame returns the TimelineIFrame field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetTimelineIFrame() TimelineEventIFrame {
	if o == nil || isNil(o.TimelineIFrame) {
		var ret TimelineEventIFrame
		return ret
	}
	return *o.TimelineIFrame
}

// GetTimelineIFrameOk returns a tuple with the TimelineIFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetTimelineIFrameOk() (*TimelineEventIFrame, bool) {
	if o == nil || isNil(o.TimelineIFrame) {
		return nil, false
	}
	return o.TimelineIFrame, true
}

// HasTimelineIFrame returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasTimelineIFrame() bool {
	if o != nil && !isNil(o.TimelineIFrame) {
		return true
	}

	return false
}

// SetTimelineIFrame gets a reference to the given TimelineEventIFrame and assigns it to the TimelineIFrame field.
func (o *TimelineEventResponse) SetTimelineIFrame(v TimelineEventIFrame) {
	o.TimelineIFrame = &v
}

// GetObjectType returns the ObjectType field value
func (o *TimelineEventResponse) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TimelineEventResponse) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TimelineEventResponse) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TimelineEventResponse) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TimelineEventResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o TimelineEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
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
	if true {
		toSerialize["objectType"] = o.ObjectType
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineEventResponse struct {
	value *TimelineEventResponse
	isSet bool
}

func (v NullableTimelineEventResponse) Get() *TimelineEventResponse {
	return v.value
}

func (v *NullableTimelineEventResponse) Set(val *TimelineEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineEventResponse(val *TimelineEventResponse) *NullableTimelineEventResponse {
	return &NullableTimelineEventResponse{value: val, isSet: true}
}

func (v NullableTimelineEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
