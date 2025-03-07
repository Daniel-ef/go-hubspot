/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
)

// TagCloneRequestVNext Request body object for cloning blog tags.
type TagCloneRequestVNext struct {
	// ID of the object to be cloned.
	Id string `json:"id"`
	// Target language of new variant.
	Language *string `json:"language,omitempty"`
	// Language of primary blog tag to clone.
	PrimaryLanguage *string `json:"primaryLanguage,omitempty"`
	// Name of newly cloned blog tag.
	Name string `json:"name"`
}

// NewTagCloneRequestVNext instantiates a new TagCloneRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCloneRequestVNext(id string, name string) *TagCloneRequestVNext {
	this := TagCloneRequestVNext{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTagCloneRequestVNextWithDefaults instantiates a new TagCloneRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCloneRequestVNextWithDefaults() *TagCloneRequestVNext {
	this := TagCloneRequestVNext{}
	return &this
}

// GetId returns the Id field value
func (o *TagCloneRequestVNext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagCloneRequestVNext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagCloneRequestVNext) SetId(v string) {
	o.Id = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TagCloneRequestVNext) GetLanguage() string {
	if o == nil || isNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCloneRequestVNext) GetLanguageOk() (*string, bool) {
	if o == nil || isNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TagCloneRequestVNext) HasLanguage() bool {
	if o != nil && !isNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *TagCloneRequestVNext) SetLanguage(v string) {
	o.Language = &v
}

// GetPrimaryLanguage returns the PrimaryLanguage field value if set, zero value otherwise.
func (o *TagCloneRequestVNext) GetPrimaryLanguage() string {
	if o == nil || isNil(o.PrimaryLanguage) {
		var ret string
		return ret
	}
	return *o.PrimaryLanguage
}

// GetPrimaryLanguageOk returns a tuple with the PrimaryLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCloneRequestVNext) GetPrimaryLanguageOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryLanguage) {
		return nil, false
	}
	return o.PrimaryLanguage, true
}

// HasPrimaryLanguage returns a boolean if a field has been set.
func (o *TagCloneRequestVNext) HasPrimaryLanguage() bool {
	if o != nil && !isNil(o.PrimaryLanguage) {
		return true
	}

	return false
}

// SetPrimaryLanguage gets a reference to the given string and assigns it to the PrimaryLanguage field.
func (o *TagCloneRequestVNext) SetPrimaryLanguage(v string) {
	o.PrimaryLanguage = &v
}

// GetName returns the Name field value
func (o *TagCloneRequestVNext) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagCloneRequestVNext) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagCloneRequestVNext) SetName(v string) {
	o.Name = v
}

func (o TagCloneRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !isNil(o.PrimaryLanguage) {
		toSerialize["primaryLanguage"] = o.PrimaryLanguage
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTagCloneRequestVNext struct {
	value *TagCloneRequestVNext
	isSet bool
}

func (v NullableTagCloneRequestVNext) Get() *TagCloneRequestVNext {
	return v.value
}

func (v *NullableTagCloneRequestVNext) Set(val *TagCloneRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCloneRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCloneRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCloneRequestVNext(val *TagCloneRequestVNext) *NullableTagCloneRequestVNext {
	return &NullableTagCloneRequestVNext{value: val, isSet: true}
}

func (v NullableTagCloneRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCloneRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
