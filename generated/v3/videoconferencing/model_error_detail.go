/*
Video Conference Extension

These APIs allow you to specify URLs that can be used to interact with a video conferencing application, to allow HubSpot to add video conference links to meeting requests with contacts.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package videoconferencing

import (
	"encoding/json"
)

// ErrorDetail struct for ErrorDetail
type ErrorDetail struct {
	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`
	// The name of the field or parameter in which the error was found.
	In *string `json:"in,omitempty"`
	// The status code associated with the error detail
	Code *string `json:"code,omitempty"`
	// A specific category that contains more specific detail about the error
	SubCategory *string `json:"subCategory,omitempty"`
	// Context about the error condition
	Context *map[string][]string `json:"context,omitempty"`
}

// NewErrorDetail instantiates a new ErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetail(message string) *ErrorDetail {
	this := ErrorDetail{}
	this.Message = message
	return &this
}

// NewErrorDetailWithDefaults instantiates a new ErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailWithDefaults() *ErrorDetail {
	this := ErrorDetail{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorDetail) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorDetail) SetMessage(v string) {
	o.Message = v
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *ErrorDetail) GetIn() string {
	if o == nil || isNil(o.In) {
		var ret string
		return ret
	}
	return *o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetInOk() (*string, bool) {
	if o == nil || isNil(o.In) {
		return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *ErrorDetail) HasIn() bool {
	if o != nil && !isNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given string and assigns it to the In field.
func (o *ErrorDetail) SetIn(v string) {
	o.In = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorDetail) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorDetail) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorDetail) SetCode(v string) {
	o.Code = &v
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise.
func (o *ErrorDetail) GetSubCategory() string {
	if o == nil || isNil(o.SubCategory) {
		var ret string
		return ret
	}
	return *o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetSubCategoryOk() (*string, bool) {
	if o == nil || isNil(o.SubCategory) {
		return nil, false
	}
	return o.SubCategory, true
}

// HasSubCategory returns a boolean if a field has been set.
func (o *ErrorDetail) HasSubCategory() bool {
	if o != nil && !isNil(o.SubCategory) {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given string and assigns it to the SubCategory field.
func (o *ErrorDetail) SetSubCategory(v string) {
	o.SubCategory = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ErrorDetail) GetContext() map[string][]string {
	if o == nil || isNil(o.Context) {
		var ret map[string][]string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetail) GetContextOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ErrorDetail) HasContext() bool {
	if o != nil && !isNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string][]string and assigns it to the Context field.
func (o *ErrorDetail) SetContext(v map[string][]string) {
	o.Context = &v
}

func (o ErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.In) {
		toSerialize["in"] = o.In
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.SubCategory) {
		toSerialize["subCategory"] = o.SubCategory
	}
	if !isNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableErrorDetail struct {
	value *ErrorDetail
	isSet bool
}

func (v NullableErrorDetail) Get() *ErrorDetail {
	return v.value
}

func (v *NullableErrorDetail) Set(val *ErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetail(val *ErrorDetail) *NullableErrorDetail {
	return &NullableErrorDetail{value: val, isSet: true}
}

func (v NullableErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
