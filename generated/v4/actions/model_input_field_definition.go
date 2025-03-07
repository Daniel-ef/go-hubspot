/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
)

// InputFieldDefinition Configuration for an input field on the custom action
type InputFieldDefinition struct {
	TypeDefinition FieldTypeDefinition `json:"typeDefinition"`
	// Controls what kind of input a customer can use to specify the field value. Must contain exactly one of `STATIC_VALUE` or `OBJECT_PROPERTY`. If `STATIC_VALUE`, the customer will be able to choose a value when configuring the custom action; if `OBJECT_PROPERTY`, the customer will be able to choose a property from the enrolled workflow object that the field value will be copied from. In the future we may support more than one input control type here.
	SupportedValueTypes []string `json:"supportedValueTypes,omitempty"`
	// Whether the field is required for the custom action to be valid
	IsRequired bool `json:"isRequired"`
}

// NewInputFieldDefinition instantiates a new InputFieldDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputFieldDefinition(typeDefinition FieldTypeDefinition, isRequired bool) *InputFieldDefinition {
	this := InputFieldDefinition{}
	this.TypeDefinition = typeDefinition
	this.IsRequired = isRequired
	return &this
}

// NewInputFieldDefinitionWithDefaults instantiates a new InputFieldDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputFieldDefinitionWithDefaults() *InputFieldDefinition {
	this := InputFieldDefinition{}
	return &this
}

// GetTypeDefinition returns the TypeDefinition field value
func (o *InputFieldDefinition) GetTypeDefinition() FieldTypeDefinition {
	if o == nil {
		var ret FieldTypeDefinition
		return ret
	}

	return o.TypeDefinition
}

// GetTypeDefinitionOk returns a tuple with the TypeDefinition field value
// and a boolean to check if the value has been set.
func (o *InputFieldDefinition) GetTypeDefinitionOk() (*FieldTypeDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeDefinition, true
}

// SetTypeDefinition sets field value
func (o *InputFieldDefinition) SetTypeDefinition(v FieldTypeDefinition) {
	o.TypeDefinition = v
}

// GetSupportedValueTypes returns the SupportedValueTypes field value if set, zero value otherwise.
func (o *InputFieldDefinition) GetSupportedValueTypes() []string {
	if o == nil || isNil(o.SupportedValueTypes) {
		var ret []string
		return ret
	}
	return o.SupportedValueTypes
}

// GetSupportedValueTypesOk returns a tuple with the SupportedValueTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputFieldDefinition) GetSupportedValueTypesOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedValueTypes) {
		return nil, false
	}
	return o.SupportedValueTypes, true
}

// HasSupportedValueTypes returns a boolean if a field has been set.
func (o *InputFieldDefinition) HasSupportedValueTypes() bool {
	if o != nil && !isNil(o.SupportedValueTypes) {
		return true
	}

	return false
}

// SetSupportedValueTypes gets a reference to the given []string and assigns it to the SupportedValueTypes field.
func (o *InputFieldDefinition) SetSupportedValueTypes(v []string) {
	o.SupportedValueTypes = v
}

// GetIsRequired returns the IsRequired field value
func (o *InputFieldDefinition) GetIsRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *InputFieldDefinition) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *InputFieldDefinition) SetIsRequired(v bool) {
	o.IsRequired = v
}

func (o InputFieldDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["typeDefinition"] = o.TypeDefinition
	}
	if !isNil(o.SupportedValueTypes) {
		toSerialize["supportedValueTypes"] = o.SupportedValueTypes
	}
	if true {
		toSerialize["isRequired"] = o.IsRequired
	}
	return json.Marshal(toSerialize)
}

type NullableInputFieldDefinition struct {
	value *InputFieldDefinition
	isSet bool
}

func (v NullableInputFieldDefinition) Get() *InputFieldDefinition {
	return v.value
}

func (v *NullableInputFieldDefinition) Set(val *InputFieldDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableInputFieldDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableInputFieldDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputFieldDefinition(val *InputFieldDefinition) *NullableInputFieldDefinition {
	return &NullableInputFieldDefinition{value: val, isSet: true}
}

func (v NullableInputFieldDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputFieldDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
