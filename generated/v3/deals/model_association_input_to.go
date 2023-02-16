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

// AssociationInputTo struct for AssociationInputTo
type AssociationInputTo struct {
	Id string `json:"id"`
}

// NewAssociationInputTo instantiates a new AssociationInputTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationInputTo(id string) *AssociationInputTo {
	this := AssociationInputTo{}
	this.Id = id
	return &this
}

// NewAssociationInputToWithDefaults instantiates a new AssociationInputTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationInputToWithDefaults() *AssociationInputTo {
	this := AssociationInputTo{}
	return &this
}

// GetId returns the Id field value
func (o *AssociationInputTo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssociationInputTo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssociationInputTo) SetId(v string) {
	o.Id = v
}

func (o AssociationInputTo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAssociationInputTo struct {
	value *AssociationInputTo
	isSet bool
}

func (v NullableAssociationInputTo) Get() *AssociationInputTo {
	return v.value
}

func (v *NullableAssociationInputTo) Set(val *AssociationInputTo) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationInputTo) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationInputTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationInputTo(val *AssociationInputTo) *NullableAssociationInputTo {
	return &NullableAssociationInputTo{value: val, isSet: true}
}

func (v NullableAssociationInputTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationInputTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
