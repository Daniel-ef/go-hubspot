/*
Webhooks API

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"encoding/json"
)

// SubscriptionPatchRequest Updated details for the subscription.
type SubscriptionPatchRequest struct {
	// Determines if the subscription is active or paused.
	Active *bool `json:"active,omitempty"`
}

// NewSubscriptionPatchRequest instantiates a new SubscriptionPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPatchRequest() *SubscriptionPatchRequest {
	this := SubscriptionPatchRequest{}
	return &this
}

// NewSubscriptionPatchRequestWithDefaults instantiates a new SubscriptionPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPatchRequestWithDefaults() *SubscriptionPatchRequest {
	this := SubscriptionPatchRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SubscriptionPatchRequest) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPatchRequest) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SubscriptionPatchRequest) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SubscriptionPatchRequest) SetActive(v bool) {
	o.Active = &v
}

func (o SubscriptionPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionPatchRequest struct {
	value *SubscriptionPatchRequest
	isSet bool
}

func (v NullableSubscriptionPatchRequest) Get() *SubscriptionPatchRequest {
	return v.value
}

func (v *NullableSubscriptionPatchRequest) Set(val *SubscriptionPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPatchRequest(val *SubscriptionPatchRequest) *NullableSubscriptionPatchRequest {
	return &NullableSubscriptionPatchRequest{value: val, isSet: true}
}

func (v NullableSubscriptionPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
