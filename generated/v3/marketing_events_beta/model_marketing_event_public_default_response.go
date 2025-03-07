/*
Marketing Events Extension

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
	"time"
)

// MarketingEventPublicDefaultResponse struct for MarketingEventPublicDefaultResponse
type MarketingEventPublicDefaultResponse struct {
	// The name of the marketing event.
	EventName string `json:"eventName"`
	// The type of the marketing event.
	EventType *string `json:"eventType,omitempty"`
	// The start date and time of the marketing event.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The end date and time of the marketing event.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer"`
	// The description of the marketing event.
	EventDescription *string `json:"eventDescription,omitempty"`
	// A URL in the external event application where the marketing event can be managed.
	EventUrl *string `json:"eventUrl,omitempty"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled *bool `json:"eventCancelled,omitempty"`
	// A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account's definition of the MarketingEvent Object. If they don't they will be filtered out and not set. In order to do this you'll need to create a new PropertyGroup on the HubSpot account's MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
	Id               string          `json:"id"`
	CreatedAt        time.Time       `json:"createdAt"`
	UpdatedAt        time.Time       `json:"updatedAt"`
}

// NewMarketingEventPublicDefaultResponse instantiates a new MarketingEventPublicDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventPublicDefaultResponse(eventName string, eventOrganizer string, id string, createdAt time.Time, updatedAt time.Time) *MarketingEventPublicDefaultResponse {
	this := MarketingEventPublicDefaultResponse{}
	this.EventName = eventName
	this.EventOrganizer = eventOrganizer
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewMarketingEventPublicDefaultResponseWithDefaults instantiates a new MarketingEventPublicDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventPublicDefaultResponseWithDefaults() *MarketingEventPublicDefaultResponse {
	this := MarketingEventPublicDefaultResponse{}
	return &this
}

// GetEventName returns the EventName field value
func (o *MarketingEventPublicDefaultResponse) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *MarketingEventPublicDefaultResponse) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MarketingEventPublicDefaultResponse) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MarketingEventPublicDefaultResponse) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MarketingEventPublicDefaultResponse) SetEventType(v string) {
	o.EventType = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *MarketingEventPublicDefaultResponse) GetStartDateTime() time.Time {
	if o == nil || isNil(o.StartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MarketingEventPublicDefaultResponse) HasStartDateTime() bool {
	if o != nil && !isNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MarketingEventPublicDefaultResponse) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *MarketingEventPublicDefaultResponse) GetEndDateTime() time.Time {
	if o == nil || isNil(o.EndDateTime) {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndDateTime) {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MarketingEventPublicDefaultResponse) HasEndDateTime() bool {
	if o != nil && !isNil(o.EndDateTime) {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *MarketingEventPublicDefaultResponse) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

// GetEventOrganizer returns the EventOrganizer field value
func (o *MarketingEventPublicDefaultResponse) GetEventOrganizer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventOrganizer
}

// GetEventOrganizerOk returns a tuple with the EventOrganizer field value
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetEventOrganizerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventOrganizer, true
}

// SetEventOrganizer sets field value
func (o *MarketingEventPublicDefaultResponse) SetEventOrganizer(v string) {
	o.EventOrganizer = v
}

// GetEventDescription returns the EventDescription field value if set, zero value otherwise.
func (o *MarketingEventPublicDefaultResponse) GetEventDescription() string {
	if o == nil || isNil(o.EventDescription) {
		var ret string
		return ret
	}
	return *o.EventDescription
}

// GetEventDescriptionOk returns a tuple with the EventDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetEventDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.EventDescription) {
		return nil, false
	}
	return o.EventDescription, true
}

// HasEventDescription returns a boolean if a field has been set.
func (o *MarketingEventPublicDefaultResponse) HasEventDescription() bool {
	if o != nil && !isNil(o.EventDescription) {
		return true
	}

	return false
}

// SetEventDescription gets a reference to the given string and assigns it to the EventDescription field.
func (o *MarketingEventPublicDefaultResponse) SetEventDescription(v string) {
	o.EventDescription = &v
}

// GetEventUrl returns the EventUrl field value if set, zero value otherwise.
func (o *MarketingEventPublicDefaultResponse) GetEventUrl() string {
	if o == nil || isNil(o.EventUrl) {
		var ret string
		return ret
	}
	return *o.EventUrl
}

// GetEventUrlOk returns a tuple with the EventUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetEventUrlOk() (*string, bool) {
	if o == nil || isNil(o.EventUrl) {
		return nil, false
	}
	return o.EventUrl, true
}

// HasEventUrl returns a boolean if a field has been set.
func (o *MarketingEventPublicDefaultResponse) HasEventUrl() bool {
	if o != nil && !isNil(o.EventUrl) {
		return true
	}

	return false
}

// SetEventUrl gets a reference to the given string and assigns it to the EventUrl field.
func (o *MarketingEventPublicDefaultResponse) SetEventUrl(v string) {
	o.EventUrl = &v
}

// GetEventCancelled returns the EventCancelled field value if set, zero value otherwise.
func (o *MarketingEventPublicDefaultResponse) GetEventCancelled() bool {
	if o == nil || isNil(o.EventCancelled) {
		var ret bool
		return ret
	}
	return *o.EventCancelled
}

// GetEventCancelledOk returns a tuple with the EventCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetEventCancelledOk() (*bool, bool) {
	if o == nil || isNil(o.EventCancelled) {
		return nil, false
	}
	return o.EventCancelled, true
}

// HasEventCancelled returns a boolean if a field has been set.
func (o *MarketingEventPublicDefaultResponse) HasEventCancelled() bool {
	if o != nil && !isNil(o.EventCancelled) {
		return true
	}

	return false
}

// SetEventCancelled gets a reference to the given bool and assigns it to the EventCancelled field.
func (o *MarketingEventPublicDefaultResponse) SetEventCancelled(v bool) {
	o.EventCancelled = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *MarketingEventPublicDefaultResponse) GetCustomProperties() []PropertyValue {
	if o == nil || isNil(o.CustomProperties) {
		var ret []PropertyValue
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetCustomPropertiesOk() ([]PropertyValue, bool) {
	if o == nil || isNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *MarketingEventPublicDefaultResponse) HasCustomProperties() bool {
	if o != nil && !isNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []PropertyValue and assigns it to the CustomProperties field.
func (o *MarketingEventPublicDefaultResponse) SetCustomProperties(v []PropertyValue) {
	o.CustomProperties = v
}

// GetId returns the Id field value
func (o *MarketingEventPublicDefaultResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MarketingEventPublicDefaultResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MarketingEventPublicDefaultResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MarketingEventPublicDefaultResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *MarketingEventPublicDefaultResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *MarketingEventPublicDefaultResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *MarketingEventPublicDefaultResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o MarketingEventPublicDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !isNil(o.EndDateTime) {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if true {
		toSerialize["eventOrganizer"] = o.EventOrganizer
	}
	if !isNil(o.EventDescription) {
		toSerialize["eventDescription"] = o.EventDescription
	}
	if !isNil(o.EventUrl) {
		toSerialize["eventUrl"] = o.EventUrl
	}
	if !isNil(o.EventCancelled) {
		toSerialize["eventCancelled"] = o.EventCancelled
	}
	if !isNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingEventPublicDefaultResponse struct {
	value *MarketingEventPublicDefaultResponse
	isSet bool
}

func (v NullableMarketingEventPublicDefaultResponse) Get() *MarketingEventPublicDefaultResponse {
	return v.value
}

func (v *NullableMarketingEventPublicDefaultResponse) Set(val *MarketingEventPublicDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventPublicDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventPublicDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventPublicDefaultResponse(val *MarketingEventPublicDefaultResponse) *NullableMarketingEventPublicDefaultResponse {
	return &NullableMarketingEventPublicDefaultResponse{value: val, isSet: true}
}

func (v NullableMarketingEventPublicDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventPublicDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
