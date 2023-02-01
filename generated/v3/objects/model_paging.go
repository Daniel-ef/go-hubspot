/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"encoding/json"
)

// Paging struct for Paging
type Paging struct {
	Next *NextPage     `json:"next,omitempty"`
	Prev *PreviousPage `json:"prev,omitempty"`
}

// NewPaging instantiates a new Paging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaging() *Paging {
	this := Paging{}
	return &this
}

// NewPagingWithDefaults instantiates a new Paging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingWithDefaults() *Paging {
	this := Paging{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Paging) GetNext() NextPage {
	if o == nil || isNil(o.Next) {
		var ret NextPage
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paging) GetNextOk() (*NextPage, bool) {
	if o == nil || isNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Paging) HasNext() bool {
	if o != nil && !isNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextPage and assigns it to the Next field.
func (o *Paging) SetNext(v NextPage) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *Paging) GetPrev() PreviousPage {
	if o == nil || isNil(o.Prev) {
		var ret PreviousPage
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paging) GetPrevOk() (*PreviousPage, bool) {
	if o == nil || isNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *Paging) HasPrev() bool {
	if o != nil && !isNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given PreviousPage and assigns it to the Prev field.
func (o *Paging) SetPrev(v PreviousPage) {
	o.Prev = &v
}

func (o Paging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !isNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	return json.Marshal(toSerialize)
}

type NullablePaging struct {
	value *Paging
	isSet bool
}

func (v NullablePaging) Get() *Paging {
	return v.value
}

func (v *NullablePaging) Set(val *Paging) {
	v.value = val
	v.isSet = true
}

func (v NullablePaging) IsSet() bool {
	return v.isSet
}

func (v *NullablePaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaging(val *Paging) *NullablePaging {
	return &NullablePaging{value: val, isSet: true}
}

func (v NullablePaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
