/*
CMS Performance API

Use these endpoints to get a time series view of your website's performance.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package performance

import (
	"encoding/json"
)

// PublicPerformanceResponse struct for PublicPerformanceResponse
type PublicPerformanceResponse struct {
	Data          []PerformanceView `json:"data"`
	Domain        *string           `json:"domain,omitempty"`
	Path          *string           `json:"path,omitempty"`
	StartInterval int64             `json:"startInterval"`
	EndInterval   int64             `json:"endInterval"`
	Interval      string            `json:"interval"`
	Period        *string           `json:"period,omitempty"`
}

// NewPublicPerformanceResponse instantiates a new PublicPerformanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPerformanceResponse(data []PerformanceView, startInterval int64, endInterval int64, interval string) *PublicPerformanceResponse {
	this := PublicPerformanceResponse{}
	this.Data = data
	this.StartInterval = startInterval
	this.EndInterval = endInterval
	this.Interval = interval
	return &this
}

// NewPublicPerformanceResponseWithDefaults instantiates a new PublicPerformanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPerformanceResponseWithDefaults() *PublicPerformanceResponse {
	this := PublicPerformanceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PublicPerformanceResponse) GetData() []PerformanceView {
	if o == nil {
		var ret []PerformanceView
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PublicPerformanceResponse) GetDataOk() ([]PerformanceView, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PublicPerformanceResponse) SetData(v []PerformanceView) {
	o.Data = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PublicPerformanceResponse) GetDomain() string {
	if o == nil || isNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPerformanceResponse) GetDomainOk() (*string, bool) {
	if o == nil || isNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PublicPerformanceResponse) HasDomain() bool {
	if o != nil && !isNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *PublicPerformanceResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PublicPerformanceResponse) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPerformanceResponse) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PublicPerformanceResponse) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PublicPerformanceResponse) SetPath(v string) {
	o.Path = &v
}

// GetStartInterval returns the StartInterval field value
func (o *PublicPerformanceResponse) GetStartInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartInterval
}

// GetStartIntervalOk returns a tuple with the StartInterval field value
// and a boolean to check if the value has been set.
func (o *PublicPerformanceResponse) GetStartIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartInterval, true
}

// SetStartInterval sets field value
func (o *PublicPerformanceResponse) SetStartInterval(v int64) {
	o.StartInterval = v
}

// GetEndInterval returns the EndInterval field value
func (o *PublicPerformanceResponse) GetEndInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndInterval
}

// GetEndIntervalOk returns a tuple with the EndInterval field value
// and a boolean to check if the value has been set.
func (o *PublicPerformanceResponse) GetEndIntervalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndInterval, true
}

// SetEndInterval sets field value
func (o *PublicPerformanceResponse) SetEndInterval(v int64) {
	o.EndInterval = v
}

// GetInterval returns the Interval field value
func (o *PublicPerformanceResponse) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *PublicPerformanceResponse) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *PublicPerformanceResponse) SetInterval(v string) {
	o.Interval = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *PublicPerformanceResponse) GetPeriod() string {
	if o == nil || isNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPerformanceResponse) GetPeriodOk() (*string, bool) {
	if o == nil || isNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *PublicPerformanceResponse) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *PublicPerformanceResponse) SetPeriod(v string) {
	o.Period = &v
}

func (o PublicPerformanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["startInterval"] = o.StartInterval
	}
	if true {
		toSerialize["endInterval"] = o.EndInterval
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if !isNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullablePublicPerformanceResponse struct {
	value *PublicPerformanceResponse
	isSet bool
}

func (v NullablePublicPerformanceResponse) Get() *PublicPerformanceResponse {
	return v.value
}

func (v *NullablePublicPerformanceResponse) Set(val *PublicPerformanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPerformanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPerformanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPerformanceResponse(val *PublicPerformanceResponse) *NullablePublicPerformanceResponse {
	return &NullablePublicPerformanceResponse{value: val, isSet: true}
}

func (v NullablePublicPerformanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPerformanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
