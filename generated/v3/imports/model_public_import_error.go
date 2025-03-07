/*
CRM Imports

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imports

import (
	"encoding/json"
)

// PublicImportError struct for PublicImportError
type PublicImportError struct {
	ErrorType         string        `json:"errorType"`
	SourceData        ImportRowCore `json:"sourceData"`
	ObjectType        *string       `json:"objectType,omitempty"`
	InvalidValue      *string       `json:"invalidValue,omitempty"`
	ExtraContext      *string       `json:"extraContext,omitempty"`
	ObjectTypeId      *string       `json:"objectTypeId,omitempty"`
	KnownColumnNumber int32         `json:"knownColumnNumber"`
	CreatedAt         int32         `json:"createdAt"`
	Id                string        `json:"id"`
}

// NewPublicImportError instantiates a new PublicImportError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicImportError(errorType string, sourceData ImportRowCore, knownColumnNumber int32, createdAt int32, id string) *PublicImportError {
	this := PublicImportError{}
	this.ErrorType = errorType
	this.SourceData = sourceData
	this.KnownColumnNumber = knownColumnNumber
	this.CreatedAt = createdAt
	this.Id = id
	return &this
}

// NewPublicImportErrorWithDefaults instantiates a new PublicImportError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicImportErrorWithDefaults() *PublicImportError {
	this := PublicImportError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *PublicImportError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *PublicImportError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetSourceData returns the SourceData field value
func (o *PublicImportError) GetSourceData() ImportRowCore {
	if o == nil {
		var ret ImportRowCore
		return ret
	}

	return o.SourceData
}

// GetSourceDataOk returns a tuple with the SourceData field value
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetSourceDataOk() (*ImportRowCore, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceData, true
}

// SetSourceData sets field value
func (o *PublicImportError) SetSourceData(v ImportRowCore) {
	o.SourceData = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *PublicImportError) GetObjectType() string {
	if o == nil || isNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *PublicImportError) HasObjectType() bool {
	if o != nil && !isNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *PublicImportError) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetInvalidValue returns the InvalidValue field value if set, zero value otherwise.
func (o *PublicImportError) GetInvalidValue() string {
	if o == nil || isNil(o.InvalidValue) {
		var ret string
		return ret
	}
	return *o.InvalidValue
}

// GetInvalidValueOk returns a tuple with the InvalidValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetInvalidValueOk() (*string, bool) {
	if o == nil || isNil(o.InvalidValue) {
		return nil, false
	}
	return o.InvalidValue, true
}

// HasInvalidValue returns a boolean if a field has been set.
func (o *PublicImportError) HasInvalidValue() bool {
	if o != nil && !isNil(o.InvalidValue) {
		return true
	}

	return false
}

// SetInvalidValue gets a reference to the given string and assigns it to the InvalidValue field.
func (o *PublicImportError) SetInvalidValue(v string) {
	o.InvalidValue = &v
}

// GetExtraContext returns the ExtraContext field value if set, zero value otherwise.
func (o *PublicImportError) GetExtraContext() string {
	if o == nil || isNil(o.ExtraContext) {
		var ret string
		return ret
	}
	return *o.ExtraContext
}

// GetExtraContextOk returns a tuple with the ExtraContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetExtraContextOk() (*string, bool) {
	if o == nil || isNil(o.ExtraContext) {
		return nil, false
	}
	return o.ExtraContext, true
}

// HasExtraContext returns a boolean if a field has been set.
func (o *PublicImportError) HasExtraContext() bool {
	if o != nil && !isNil(o.ExtraContext) {
		return true
	}

	return false
}

// SetExtraContext gets a reference to the given string and assigns it to the ExtraContext field.
func (o *PublicImportError) SetExtraContext(v string) {
	o.ExtraContext = &v
}

// GetObjectTypeId returns the ObjectTypeId field value if set, zero value otherwise.
func (o *PublicImportError) GetObjectTypeId() string {
	if o == nil || isNil(o.ObjectTypeId) {
		var ret string
		return ret
	}
	return *o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetObjectTypeIdOk() (*string, bool) {
	if o == nil || isNil(o.ObjectTypeId) {
		return nil, false
	}
	return o.ObjectTypeId, true
}

// HasObjectTypeId returns a boolean if a field has been set.
func (o *PublicImportError) HasObjectTypeId() bool {
	if o != nil && !isNil(o.ObjectTypeId) {
		return true
	}

	return false
}

// SetObjectTypeId gets a reference to the given string and assigns it to the ObjectTypeId field.
func (o *PublicImportError) SetObjectTypeId(v string) {
	o.ObjectTypeId = &v
}

// GetKnownColumnNumber returns the KnownColumnNumber field value
func (o *PublicImportError) GetKnownColumnNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KnownColumnNumber
}

// GetKnownColumnNumberOk returns a tuple with the KnownColumnNumber field value
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetKnownColumnNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KnownColumnNumber, true
}

// SetKnownColumnNumber sets field value
func (o *PublicImportError) SetKnownColumnNumber(v int32) {
	o.KnownColumnNumber = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PublicImportError) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PublicImportError) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *PublicImportError) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicImportError) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicImportError) SetId(v string) {
	o.Id = v
}

func (o PublicImportError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errorType"] = o.ErrorType
	}
	if true {
		toSerialize["sourceData"] = o.SourceData
	}
	if !isNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !isNil(o.InvalidValue) {
		toSerialize["invalidValue"] = o.InvalidValue
	}
	if !isNil(o.ExtraContext) {
		toSerialize["extraContext"] = o.ExtraContext
	}
	if !isNil(o.ObjectTypeId) {
		toSerialize["objectTypeId"] = o.ObjectTypeId
	}
	if true {
		toSerialize["knownColumnNumber"] = o.KnownColumnNumber
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePublicImportError struct {
	value *PublicImportError
	isSet bool
}

func (v NullablePublicImportError) Get() *PublicImportError {
	return v.value
}

func (v *NullablePublicImportError) Set(val *PublicImportError) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicImportError) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicImportError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicImportError(val *PublicImportError) *NullablePublicImportError {
	return &NullablePublicImportError{value: val, isSet: true}
}

func (v NullablePublicImportError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicImportError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
