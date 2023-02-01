/*
Video Conference Extension

These APIs allow you to specify URLs that can be used to interact with a video conferencing application, to allow HubSpot to add video conference links to meeting requests with contacts.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package videoconferencing

import (
	"encoding/json"
)

// ExternalSettings The URLs of the various actions provided by the video conferencing application. All URLs must use the `https` protocol.
type ExternalSettings struct {
	// The URL that HubSpot will send requests to create a new video conference.
	CreateMeetingUrl string `json:"createMeetingUrl"`
	// The URL that HubSpot will send updates to existing meetings. Typically called when the user changes the topic or times of a meeting.
	UpdateMeetingUrl *string `json:"updateMeetingUrl,omitempty"`
	// The URL that HubSpot will send notifications of meetings that have been deleted in HubSpot.
	DeleteMeetingUrl *string `json:"deleteMeetingUrl,omitempty"`
	// The URL that HubSpot will use to verify that a user exists in the video conference application.
	UserVerifyUrl *string `json:"userVerifyUrl,omitempty"`
}

// NewExternalSettings instantiates a new ExternalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalSettings(createMeetingUrl string) *ExternalSettings {
	this := ExternalSettings{}
	this.CreateMeetingUrl = createMeetingUrl
	return &this
}

// NewExternalSettingsWithDefaults instantiates a new ExternalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalSettingsWithDefaults() *ExternalSettings {
	this := ExternalSettings{}
	return &this
}

// GetCreateMeetingUrl returns the CreateMeetingUrl field value
func (o *ExternalSettings) GetCreateMeetingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreateMeetingUrl
}

// GetCreateMeetingUrlOk returns a tuple with the CreateMeetingUrl field value
// and a boolean to check if the value has been set.
func (o *ExternalSettings) GetCreateMeetingUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateMeetingUrl, true
}

// SetCreateMeetingUrl sets field value
func (o *ExternalSettings) SetCreateMeetingUrl(v string) {
	o.CreateMeetingUrl = v
}

// GetUpdateMeetingUrl returns the UpdateMeetingUrl field value if set, zero value otherwise.
func (o *ExternalSettings) GetUpdateMeetingUrl() string {
	if o == nil || isNil(o.UpdateMeetingUrl) {
		var ret string
		return ret
	}
	return *o.UpdateMeetingUrl
}

// GetUpdateMeetingUrlOk returns a tuple with the UpdateMeetingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSettings) GetUpdateMeetingUrlOk() (*string, bool) {
	if o == nil || isNil(o.UpdateMeetingUrl) {
		return nil, false
	}
	return o.UpdateMeetingUrl, true
}

// HasUpdateMeetingUrl returns a boolean if a field has been set.
func (o *ExternalSettings) HasUpdateMeetingUrl() bool {
	if o != nil && !isNil(o.UpdateMeetingUrl) {
		return true
	}

	return false
}

// SetUpdateMeetingUrl gets a reference to the given string and assigns it to the UpdateMeetingUrl field.
func (o *ExternalSettings) SetUpdateMeetingUrl(v string) {
	o.UpdateMeetingUrl = &v
}

// GetDeleteMeetingUrl returns the DeleteMeetingUrl field value if set, zero value otherwise.
func (o *ExternalSettings) GetDeleteMeetingUrl() string {
	if o == nil || isNil(o.DeleteMeetingUrl) {
		var ret string
		return ret
	}
	return *o.DeleteMeetingUrl
}

// GetDeleteMeetingUrlOk returns a tuple with the DeleteMeetingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSettings) GetDeleteMeetingUrlOk() (*string, bool) {
	if o == nil || isNil(o.DeleteMeetingUrl) {
		return nil, false
	}
	return o.DeleteMeetingUrl, true
}

// HasDeleteMeetingUrl returns a boolean if a field has been set.
func (o *ExternalSettings) HasDeleteMeetingUrl() bool {
	if o != nil && !isNil(o.DeleteMeetingUrl) {
		return true
	}

	return false
}

// SetDeleteMeetingUrl gets a reference to the given string and assigns it to the DeleteMeetingUrl field.
func (o *ExternalSettings) SetDeleteMeetingUrl(v string) {
	o.DeleteMeetingUrl = &v
}

// GetUserVerifyUrl returns the UserVerifyUrl field value if set, zero value otherwise.
func (o *ExternalSettings) GetUserVerifyUrl() string {
	if o == nil || isNil(o.UserVerifyUrl) {
		var ret string
		return ret
	}
	return *o.UserVerifyUrl
}

// GetUserVerifyUrlOk returns a tuple with the UserVerifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSettings) GetUserVerifyUrlOk() (*string, bool) {
	if o == nil || isNil(o.UserVerifyUrl) {
		return nil, false
	}
	return o.UserVerifyUrl, true
}

// HasUserVerifyUrl returns a boolean if a field has been set.
func (o *ExternalSettings) HasUserVerifyUrl() bool {
	if o != nil && !isNil(o.UserVerifyUrl) {
		return true
	}

	return false
}

// SetUserVerifyUrl gets a reference to the given string and assigns it to the UserVerifyUrl field.
func (o *ExternalSettings) SetUserVerifyUrl(v string) {
	o.UserVerifyUrl = &v
}

func (o ExternalSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createMeetingUrl"] = o.CreateMeetingUrl
	}
	if !isNil(o.UpdateMeetingUrl) {
		toSerialize["updateMeetingUrl"] = o.UpdateMeetingUrl
	}
	if !isNil(o.DeleteMeetingUrl) {
		toSerialize["deleteMeetingUrl"] = o.DeleteMeetingUrl
	}
	if !isNil(o.UserVerifyUrl) {
		toSerialize["userVerifyUrl"] = o.UserVerifyUrl
	}
	return json.Marshal(toSerialize)
}

type NullableExternalSettings struct {
	value *ExternalSettings
	isSet bool
}

func (v NullableExternalSettings) Get() *ExternalSettings {
	return v.value
}

func (v *NullableExternalSettings) Set(val *ExternalSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalSettings(val *ExternalSettings) *NullableExternalSettings {
	return &NullableExternalSettings{value: val, isSet: true}
}

func (v NullableExternalSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
