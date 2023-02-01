/*
Products

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package products

import (
	"encoding/json"
	"time"
)

// BatchResponseSimplePublicObject struct for BatchResponseSimplePublicObject
type BatchResponseSimplePublicObject struct {
	Status      string               `json:"status"`
	Results     []SimplePublicObject `json:"results"`
	RequestedAt *time.Time           `json:"requestedAt,omitempty"`
	StartedAt   time.Time            `json:"startedAt"`
	CompletedAt time.Time            `json:"completedAt"`
	Links       *map[string]string   `json:"links,omitempty"`
}

// NewBatchResponseSimplePublicObject instantiates a new BatchResponseSimplePublicObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseSimplePublicObject(status string, results []SimplePublicObject, startedAt time.Time, completedAt time.Time) *BatchResponseSimplePublicObject {
	this := BatchResponseSimplePublicObject{}
	this.Status = status
	this.Results = results
	this.StartedAt = startedAt
	this.CompletedAt = completedAt
	return &this
}

// NewBatchResponseSimplePublicObjectWithDefaults instantiates a new BatchResponseSimplePublicObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseSimplePublicObjectWithDefaults() *BatchResponseSimplePublicObject {
	this := BatchResponseSimplePublicObject{}
	return &this
}

// GetStatus returns the Status field value
func (o *BatchResponseSimplePublicObject) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObject) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseSimplePublicObject) SetStatus(v string) {
	o.Status = v
}

// GetResults returns the Results field value
func (o *BatchResponseSimplePublicObject) GetResults() []SimplePublicObject {
	if o == nil {
		var ret []SimplePublicObject
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObject) GetResultsOk() ([]SimplePublicObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseSimplePublicObject) SetResults(v []SimplePublicObject) {
	o.Results = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseSimplePublicObject) GetRequestedAt() time.Time {
	if o == nil || isNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObject) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseSimplePublicObject) HasRequestedAt() bool {
	if o != nil && !isNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseSimplePublicObject) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseSimplePublicObject) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObject) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseSimplePublicObject) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseSimplePublicObject) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObject) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseSimplePublicObject) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseSimplePublicObject) GetLinks() map[string]string {
	if o == nil || isNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSimplePublicObject) GetLinksOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseSimplePublicObject) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseSimplePublicObject) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o BatchResponseSimplePublicObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["results"] = o.Results
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

type NullableBatchResponseSimplePublicObject struct {
	value *BatchResponseSimplePublicObject
	isSet bool
}

func (v NullableBatchResponseSimplePublicObject) Get() *BatchResponseSimplePublicObject {
	return v.value
}

func (v *NullableBatchResponseSimplePublicObject) Set(val *BatchResponseSimplePublicObject) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseSimplePublicObject) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseSimplePublicObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseSimplePublicObject(val *BatchResponseSimplePublicObject) *NullableBatchResponseSimplePublicObject {
	return &NullableBatchResponseSimplePublicObject{value: val, isSet: true}
}

func (v NullableBatchResponseSimplePublicObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseSimplePublicObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
