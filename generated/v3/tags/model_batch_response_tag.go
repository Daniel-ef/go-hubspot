/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
	"time"
)

// BatchResponseTag Response object for batch operations on blog tags.
type BatchResponseTag struct {
	// Status of batch operation.
	Status string `json:"status"`
	// Results of batch operation.
	Results []Tag `json:"results"`
	// Time of batch operation request.
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	// Time of batch operation start.
	StartedAt time.Time `json:"startedAt"`
	// Time of batch operation completion.
	CompletedAt time.Time `json:"completedAt"`
	// Links associated with batch operation.
	Links *map[string]string `json:"links,omitempty"`
}

// NewBatchResponseTag instantiates a new BatchResponseTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseTag(status string, results []Tag, startedAt time.Time, completedAt time.Time) *BatchResponseTag {
	this := BatchResponseTag{}
	this.Status = status
	this.Results = results
	this.StartedAt = startedAt
	this.CompletedAt = completedAt
	return &this
}

// NewBatchResponseTagWithDefaults instantiates a new BatchResponseTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseTagWithDefaults() *BatchResponseTag {
	this := BatchResponseTag{}
	return &this
}

// GetStatus returns the Status field value
func (o *BatchResponseTag) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTag) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseTag) SetStatus(v string) {
	o.Status = v
}

// GetResults returns the Results field value
func (o *BatchResponseTag) GetResults() []Tag {
	if o == nil {
		var ret []Tag
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTag) GetResultsOk() ([]Tag, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseTag) SetResults(v []Tag) {
	o.Results = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseTag) GetRequestedAt() time.Time {
	if o == nil || isNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseTag) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseTag) HasRequestedAt() bool {
	if o != nil && !isNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseTag) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseTag) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTag) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseTag) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseTag) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseTag) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseTag) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseTag) GetLinks() map[string]string {
	if o == nil || isNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseTag) GetLinksOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseTag) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseTag) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o BatchResponseTag) MarshalJSON() ([]byte, error) {
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

type NullableBatchResponseTag struct {
	value *BatchResponseTag
	isSet bool
}

func (v NullableBatchResponseTag) Get() *BatchResponseTag {
	return v.value
}

func (v *NullableBatchResponseTag) Set(val *BatchResponseTag) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseTag) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseTag(val *BatchResponseTag) *NullableBatchResponseTag {
	return &NullableBatchResponseTag{value: val, isSet: true}
}

func (v NullableBatchResponseTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
