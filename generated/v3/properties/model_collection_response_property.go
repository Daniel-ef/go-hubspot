/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"encoding/json"
)

// CollectionResponseProperty struct for CollectionResponseProperty
type CollectionResponseProperty struct {
	Results []Property `json:"results"`
	Paging  *Paging    `json:"paging,omitempty"`
}

// NewCollectionResponseProperty instantiates a new CollectionResponseProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseProperty(results []Property) *CollectionResponseProperty {
	this := CollectionResponseProperty{}
	this.Results = results
	return &this
}

// NewCollectionResponsePropertyWithDefaults instantiates a new CollectionResponseProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponsePropertyWithDefaults() *CollectionResponseProperty {
	this := CollectionResponseProperty{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponseProperty) GetResults() []Property {
	if o == nil {
		var ret []Property
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseProperty) GetResultsOk() ([]Property, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseProperty) SetResults(v []Property) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseProperty) GetPaging() Paging {
	if o == nil || isNil(o.Paging) {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseProperty) GetPagingOk() (*Paging, bool) {
	if o == nil || isNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseProperty) HasPaging() bool {
	if o != nil && !isNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *CollectionResponseProperty) SetPaging(v Paging) {
	o.Paging = &v
}

func (o CollectionResponseProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	if !isNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseProperty struct {
	value *CollectionResponseProperty
	isSet bool
}

func (v NullableCollectionResponseProperty) Get() *CollectionResponseProperty {
	return v.value
}

func (v *NullableCollectionResponseProperty) Set(val *CollectionResponseProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseProperty(val *CollectionResponseProperty) *NullableCollectionResponseProperty {
	return &NullableCollectionResponseProperty{value: val, isSet: true}
}

func (v NullableCollectionResponseProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
