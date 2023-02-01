/*
CMS Site Search

Use these endpoints for searching content on your HubSpot hosted CMS website(s).

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package site_search

import (
	"encoding/json"
)

// IndexedField struct for IndexedField
type IndexedField struct {
	Name          string                   `json:"name"`
	Value         map[string]interface{}   `json:"value"`
	Values        []map[string]interface{} `json:"values"`
	MetadataField bool                     `json:"metadataField"`
}

// NewIndexedField instantiates a new IndexedField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexedField(name string, value map[string]interface{}, values []map[string]interface{}, metadataField bool) *IndexedField {
	this := IndexedField{}
	this.Name = name
	this.Value = value
	this.Values = values
	this.MetadataField = metadataField
	return &this
}

// NewIndexedFieldWithDefaults instantiates a new IndexedField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexedFieldWithDefaults() *IndexedField {
	this := IndexedField{}
	return &this
}

// GetName returns the Name field value
func (o *IndexedField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IndexedField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IndexedField) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *IndexedField) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *IndexedField) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *IndexedField) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetValues returns the Values field value
func (o *IndexedField) GetValues() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *IndexedField) GetValuesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *IndexedField) SetValues(v []map[string]interface{}) {
	o.Values = v
}

// GetMetadataField returns the MetadataField field value
func (o *IndexedField) GetMetadataField() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MetadataField
}

// GetMetadataFieldOk returns a tuple with the MetadataField field value
// and a boolean to check if the value has been set.
func (o *IndexedField) GetMetadataFieldOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataField, true
}

// SetMetadataField sets field value
func (o *IndexedField) SetMetadataField(v bool) {
	o.MetadataField = v
}

func (o IndexedField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["values"] = o.Values
	}
	if true {
		toSerialize["metadataField"] = o.MetadataField
	}
	return json.Marshal(toSerialize)
}

type NullableIndexedField struct {
	value *IndexedField
	isSet bool
}

func (v NullableIndexedField) Get() *IndexedField {
	return v.value
}

func (v *NullableIndexedField) Set(val *IndexedField) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexedField) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexedField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexedField(val *IndexedField) *NullableIndexedField {
	return &NullableIndexedField{value: val, isSet: true}
}

func (v NullableIndexedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexedField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
