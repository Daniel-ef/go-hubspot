/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"encoding/json"
)

// CollectionResponseMultiAssociatedObjectWithLabelForwardPaging struct for CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
type CollectionResponseMultiAssociatedObjectWithLabelForwardPaging struct {
	Results []MultiAssociatedObjectWithLabel `json:"results"`
	Paging  *ForwardPaging                   `json:"paging,omitempty"`
}

// NewCollectionResponseMultiAssociatedObjectWithLabelForwardPaging instantiates a new CollectionResponseMultiAssociatedObjectWithLabelForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseMultiAssociatedObjectWithLabelForwardPaging(results []MultiAssociatedObjectWithLabel) *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging {
	this := CollectionResponseMultiAssociatedObjectWithLabelForwardPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseMultiAssociatedObjectWithLabelForwardPagingWithDefaults instantiates a new CollectionResponseMultiAssociatedObjectWithLabelForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseMultiAssociatedObjectWithLabelForwardPagingWithDefaults() *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging {
	this := CollectionResponseMultiAssociatedObjectWithLabelForwardPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) GetResults() []MultiAssociatedObjectWithLabel {
	if o == nil {
		var ret []MultiAssociatedObjectWithLabel
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) GetResultsOk() ([]MultiAssociatedObjectWithLabel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) SetResults(v []MultiAssociatedObjectWithLabel) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) GetPaging() ForwardPaging {
	if o == nil || isNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || isNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) HasPaging() bool {
	if o != nil && !isNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

func (o CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	if !isNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging struct {
	value *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	isSet bool
}

func (v NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging) Get() *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging) Set(val *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging(val *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) *NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging {
	return &NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseMultiAssociatedObjectWithLabelForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
