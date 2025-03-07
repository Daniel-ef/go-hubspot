/*
Deals

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deals

import (
	"encoding/json"
)

// CollectionResponseSimplePublicObjectWithAssociationsForwardPaging struct for CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
type CollectionResponseSimplePublicObjectWithAssociationsForwardPaging struct {
	Results []SimplePublicObjectWithAssociations `json:"results"`
	Paging  *ForwardPaging                       `json:"paging,omitempty"`
}

// NewCollectionResponseSimplePublicObjectWithAssociationsForwardPaging instantiates a new CollectionResponseSimplePublicObjectWithAssociationsForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseSimplePublicObjectWithAssociationsForwardPaging(results []SimplePublicObjectWithAssociations) *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging {
	this := CollectionResponseSimplePublicObjectWithAssociationsForwardPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponseSimplePublicObjectWithAssociationsForwardPagingWithDefaults instantiates a new CollectionResponseSimplePublicObjectWithAssociationsForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseSimplePublicObjectWithAssociationsForwardPagingWithDefaults() *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging {
	this := CollectionResponseSimplePublicObjectWithAssociationsForwardPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) GetResults() []SimplePublicObjectWithAssociations {
	if o == nil {
		var ret []SimplePublicObjectWithAssociations
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) GetResultsOk() ([]SimplePublicObjectWithAssociations, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) SetResults(v []SimplePublicObjectWithAssociations) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) GetPaging() ForwardPaging {
	if o == nil || isNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || isNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) HasPaging() bool {
	if o != nil && !isNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

func (o CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	if !isNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging struct {
	value *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
	isSet bool
}

func (v NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging) Get() *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging) Set(val *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging(val *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) *NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging {
	return &NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseSimplePublicObjectWithAssociationsForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
