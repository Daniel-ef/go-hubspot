/*
Timeline events

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"encoding/json"
)

// TimelineEventTemplateTokenUpdateRequest State of the token definition for update requests.
type TimelineEventTemplateTokenUpdateRequest struct {
	// Used for list segmentation and reporting.
	Label string `json:"label"`
	// The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API.
	ObjectPropertyName *string `json:"objectPropertyName,omitempty"`
	// If type is `enumeration`, we should have a list of options to choose from.
	Options []TimelineEventTemplateTokenOption `json:"options"`
}

// NewTimelineEventTemplateTokenUpdateRequest instantiates a new TimelineEventTemplateTokenUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineEventTemplateTokenUpdateRequest(label string, options []TimelineEventTemplateTokenOption) *TimelineEventTemplateTokenUpdateRequest {
	this := TimelineEventTemplateTokenUpdateRequest{}
	this.Label = label
	this.Options = options
	return &this
}

// NewTimelineEventTemplateTokenUpdateRequestWithDefaults instantiates a new TimelineEventTemplateTokenUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineEventTemplateTokenUpdateRequestWithDefaults() *TimelineEventTemplateTokenUpdateRequest {
	this := TimelineEventTemplateTokenUpdateRequest{}
	return &this
}

// GetLabel returns the Label field value
func (o *TimelineEventTemplateTokenUpdateRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateTokenUpdateRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TimelineEventTemplateTokenUpdateRequest) SetLabel(v string) {
	o.Label = v
}

// GetObjectPropertyName returns the ObjectPropertyName field value if set, zero value otherwise.
func (o *TimelineEventTemplateTokenUpdateRequest) GetObjectPropertyName() string {
	if o == nil || isNil(o.ObjectPropertyName) {
		var ret string
		return ret
	}
	return *o.ObjectPropertyName
}

// GetObjectPropertyNameOk returns a tuple with the ObjectPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateTokenUpdateRequest) GetObjectPropertyNameOk() (*string, bool) {
	if o == nil || isNil(o.ObjectPropertyName) {
		return nil, false
	}
	return o.ObjectPropertyName, true
}

// HasObjectPropertyName returns a boolean if a field has been set.
func (o *TimelineEventTemplateTokenUpdateRequest) HasObjectPropertyName() bool {
	if o != nil && !isNil(o.ObjectPropertyName) {
		return true
	}

	return false
}

// SetObjectPropertyName gets a reference to the given string and assigns it to the ObjectPropertyName field.
func (o *TimelineEventTemplateTokenUpdateRequest) SetObjectPropertyName(v string) {
	o.ObjectPropertyName = &v
}

// GetOptions returns the Options field value
func (o *TimelineEventTemplateTokenUpdateRequest) GetOptions() []TimelineEventTemplateTokenOption {
	if o == nil {
		var ret []TimelineEventTemplateTokenOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateTokenUpdateRequest) GetOptionsOk() ([]TimelineEventTemplateTokenOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *TimelineEventTemplateTokenUpdateRequest) SetOptions(v []TimelineEventTemplateTokenOption) {
	o.Options = v
}

func (o TimelineEventTemplateTokenUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.ObjectPropertyName) {
		toSerialize["objectPropertyName"] = o.ObjectPropertyName
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableTimelineEventTemplateTokenUpdateRequest struct {
	value *TimelineEventTemplateTokenUpdateRequest
	isSet bool
}

func (v NullableTimelineEventTemplateTokenUpdateRequest) Get() *TimelineEventTemplateTokenUpdateRequest {
	return v.value
}

func (v *NullableTimelineEventTemplateTokenUpdateRequest) Set(val *TimelineEventTemplateTokenUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineEventTemplateTokenUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineEventTemplateTokenUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineEventTemplateTokenUpdateRequest(val *TimelineEventTemplateTokenUpdateRequest) *NullableTimelineEventTemplateTokenUpdateRequest {
	return &NullableTimelineEventTemplateTokenUpdateRequest{value: val, isSet: true}
}

func (v NullableTimelineEventTemplateTokenUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineEventTemplateTokenUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
