/*
Line Items

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package line_items

import (
	"encoding/json"
)

// BatchReadInputSimplePublicObjectId struct for BatchReadInputSimplePublicObjectId
type BatchReadInputSimplePublicObjectId struct {
	Properties            []string               `json:"properties"`
	PropertiesWithHistory []string               `json:"propertiesWithHistory"`
	IdProperty            *string                `json:"idProperty,omitempty"`
	Inputs                []SimplePublicObjectId `json:"inputs"`
}

// NewBatchReadInputSimplePublicObjectId instantiates a new BatchReadInputSimplePublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchReadInputSimplePublicObjectId(properties []string, propertiesWithHistory []string, inputs []SimplePublicObjectId) *BatchReadInputSimplePublicObjectId {
	this := BatchReadInputSimplePublicObjectId{}
	this.Properties = properties
	this.PropertiesWithHistory = propertiesWithHistory
	this.Inputs = inputs
	return &this
}

// NewBatchReadInputSimplePublicObjectIdWithDefaults instantiates a new BatchReadInputSimplePublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchReadInputSimplePublicObjectIdWithDefaults() *BatchReadInputSimplePublicObjectId {
	this := BatchReadInputSimplePublicObjectId{}
	return &this
}

// GetProperties returns the Properties field value
func (o *BatchReadInputSimplePublicObjectId) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *BatchReadInputSimplePublicObjectId) SetProperties(v []string) {
	o.Properties = v
}

// GetPropertiesWithHistory returns the PropertiesWithHistory field value
func (o *BatchReadInputSimplePublicObjectId) GetPropertiesWithHistory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertiesWithHistory
}

// GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetPropertiesWithHistoryOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertiesWithHistory, true
}

// SetPropertiesWithHistory sets field value
func (o *BatchReadInputSimplePublicObjectId) SetPropertiesWithHistory(v []string) {
	o.PropertiesWithHistory = v
}

// GetIdProperty returns the IdProperty field value if set, zero value otherwise.
func (o *BatchReadInputSimplePublicObjectId) GetIdProperty() string {
	if o == nil || isNil(o.IdProperty) {
		var ret string
		return ret
	}
	return *o.IdProperty
}

// GetIdPropertyOk returns a tuple with the IdProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetIdPropertyOk() (*string, bool) {
	if o == nil || isNil(o.IdProperty) {
		return nil, false
	}
	return o.IdProperty, true
}

// HasIdProperty returns a boolean if a field has been set.
func (o *BatchReadInputSimplePublicObjectId) HasIdProperty() bool {
	if o != nil && !isNil(o.IdProperty) {
		return true
	}

	return false
}

// SetIdProperty gets a reference to the given string and assigns it to the IdProperty field.
func (o *BatchReadInputSimplePublicObjectId) SetIdProperty(v string) {
	o.IdProperty = &v
}

// GetInputs returns the Inputs field value
func (o *BatchReadInputSimplePublicObjectId) GetInputs() []SimplePublicObjectId {
	if o == nil {
		var ret []SimplePublicObjectId
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetInputsOk() ([]SimplePublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchReadInputSimplePublicObjectId) SetInputs(v []SimplePublicObjectId) {
	o.Inputs = v
}

func (o BatchReadInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["propertiesWithHistory"] = o.PropertiesWithHistory
	}
	if !isNil(o.IdProperty) {
		toSerialize["idProperty"] = o.IdProperty
	}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchReadInputSimplePublicObjectId struct {
	value *BatchReadInputSimplePublicObjectId
	isSet bool
}

func (v NullableBatchReadInputSimplePublicObjectId) Get() *BatchReadInputSimplePublicObjectId {
	return v.value
}

func (v *NullableBatchReadInputSimplePublicObjectId) Set(val *BatchReadInputSimplePublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchReadInputSimplePublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchReadInputSimplePublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchReadInputSimplePublicObjectId(val *BatchReadInputSimplePublicObjectId) *NullableBatchReadInputSimplePublicObjectId {
	return &NullableBatchReadInputSimplePublicObjectId{value: val, isSet: true}
}

func (v NullableBatchReadInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchReadInputSimplePublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
