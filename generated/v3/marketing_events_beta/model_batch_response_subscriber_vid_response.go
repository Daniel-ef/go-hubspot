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

// BatchResponseSubscriberVidResponse struct for BatchResponseSubscriberVidResponse
type BatchResponseSubscriberVidResponse struct {
	Status      string                  `json:"status"`
	Results     []SubscriberVidResponse `json:"results"`
	NumErrors   *int32                  `json:"numErrors,omitempty"`
	Errors      []StandardError         `json:"errors,omitempty"`
	RequestedAt *time.Time              `json:"requestedAt,omitempty"`
	StartedAt   time.Time               `json:"startedAt"`
	CompletedAt time.Time               `json:"completedAt"`
	Links       *map[string]string      `json:"links,omitempty"`
}

// NewBatchResponseSubscriberVidResponse instantiates a new BatchResponseSubscriberVidResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseSubscriberVidResponse(status string, results []SubscriberVidResponse, startedAt time.Time, completedAt time.Time) *BatchResponseSubscriberVidResponse {
	this := BatchResponseSubscriberVidResponse{}
	this.Status = status
	this.Results = results
	this.StartedAt = startedAt
	this.CompletedAt = completedAt
	return &this
}

// NewBatchResponseSubscriberVidResponseWithDefaults instantiates a new BatchResponseSubscriberVidResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseSubscriberVidResponseWithDefaults() *BatchResponseSubscriberVidResponse {
	this := BatchResponseSubscriberVidResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *BatchResponseSubscriberVidResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseSubscriberVidResponse) SetStatus(v string) {
	o.Status = v
}

// GetResults returns the Results field value
func (o *BatchResponseSubscriberVidResponse) GetResults() []SubscriberVidResponse {
	if o == nil {
		var ret []SubscriberVidResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetResultsOk() ([]SubscriberVidResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseSubscriberVidResponse) SetResults(v []SubscriberVidResponse) {
	o.Results = v
}

// GetNumErrors returns the NumErrors field value if set, zero value otherwise.
func (o *BatchResponseSubscriberVidResponse) GetNumErrors() int32 {
	if o == nil || isNil(o.NumErrors) {
		var ret int32
		return ret
	}
	return *o.NumErrors
}

// GetNumErrorsOk returns a tuple with the NumErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetNumErrorsOk() (*int32, bool) {
	if o == nil || isNil(o.NumErrors) {
		return nil, false
	}
	return o.NumErrors, true
}

// HasNumErrors returns a boolean if a field has been set.
func (o *BatchResponseSubscriberVidResponse) HasNumErrors() bool {
	if o != nil && !isNil(o.NumErrors) {
		return true
	}

	return false
}

// SetNumErrors gets a reference to the given int32 and assigns it to the NumErrors field.
func (o *BatchResponseSubscriberVidResponse) SetNumErrors(v int32) {
	o.NumErrors = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BatchResponseSubscriberVidResponse) GetErrors() []StandardError {
	if o == nil || isNil(o.Errors) {
		var ret []StandardError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetErrorsOk() ([]StandardError, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BatchResponseSubscriberVidResponse) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []StandardError and assigns it to the Errors field.
func (o *BatchResponseSubscriberVidResponse) SetErrors(v []StandardError) {
	o.Errors = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseSubscriberVidResponse) GetRequestedAt() time.Time {
	if o == nil || isNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseSubscriberVidResponse) HasRequestedAt() bool {
	if o != nil && !isNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseSubscriberVidResponse) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseSubscriberVidResponse) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseSubscriberVidResponse) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseSubscriberVidResponse) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseSubscriberVidResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseSubscriberVidResponse) GetLinks() map[string]string {
	if o == nil || isNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriberVidResponse) GetLinksOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseSubscriberVidResponse) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseSubscriberVidResponse) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o BatchResponseSubscriberVidResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if !isNil(o.NumErrors) {
		toSerialize["numErrors"] = o.NumErrors
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.RequestedAt) {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if true {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableBatchResponseSubscriberVidResponse struct {
	value *BatchResponseSubscriberVidResponse
	isSet bool
}

func (v NullableBatchResponseSubscriberVidResponse) Get() *BatchResponseSubscriberVidResponse {
	return v.value
}

func (v *NullableBatchResponseSubscriberVidResponse) Set(val *BatchResponseSubscriberVidResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseSubscriberVidResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseSubscriberVidResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseSubscriberVidResponse(val *BatchResponseSubscriberVidResponse) *NullableBatchResponseSubscriberVidResponse {
	return &NullableBatchResponseSubscriberVidResponse{value: val, isSet: true}
}

func (v NullableBatchResponseSubscriberVidResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseSubscriberVidResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
