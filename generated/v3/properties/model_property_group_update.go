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

// PropertyGroupUpdate struct for PropertyGroupUpdate
type PropertyGroupUpdate struct {
	// A human-readable label that will be shown in HubSpot.
	Label *string `json:"label,omitempty"`
	// Property groups are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property group to be displayed after any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
}

// NewPropertyGroupUpdate instantiates a new PropertyGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyGroupUpdate() *PropertyGroupUpdate {
	this := PropertyGroupUpdate{}
	return &this
}

// NewPropertyGroupUpdateWithDefaults instantiates a new PropertyGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyGroupUpdateWithDefaults() *PropertyGroupUpdate {
	this := PropertyGroupUpdate{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PropertyGroupUpdate) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyGroupUpdate) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PropertyGroupUpdate) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PropertyGroupUpdate) SetLabel(v string) {
	o.Label = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *PropertyGroupUpdate) GetDisplayOrder() int32 {
	if o == nil || isNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyGroupUpdate) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || isNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *PropertyGroupUpdate) HasDisplayOrder() bool {
	if o != nil && !isNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *PropertyGroupUpdate) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

func (o PropertyGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyGroupUpdate struct {
	value *PropertyGroupUpdate
	isSet bool
}

func (v NullablePropertyGroupUpdate) Get() *PropertyGroupUpdate {
	return v.value
}

func (v *NullablePropertyGroupUpdate) Set(val *PropertyGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyGroupUpdate(val *PropertyGroupUpdate) *NullablePropertyGroupUpdate {
	return &NullablePropertyGroupUpdate{value: val, isSet: true}
}

func (v NullablePropertyGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
