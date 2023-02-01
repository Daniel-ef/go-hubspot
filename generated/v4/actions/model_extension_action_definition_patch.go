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

// ExtensionActionDefinitionPatch Fields on custom workflow action to be updated.
type ExtensionActionDefinitionPatch struct {
	// The URL that will accept an HTTPS request each time workflows executes the custom action.
	ActionUrl *string `json:"actionUrl,omitempty"`
	// Whether this custom action is published to customers.
	Published *bool `json:"published,omitempty"`
	// The list of input fields to display in this custom action.
	InputFields          []InputFieldDefinition `json:"inputFields,omitempty"`
	ObjectRequestOptions *ObjectRequestOptions  `json:"objectRequestOptions,omitempty"`
	// A list of dependencies between the input fields. These configure when the input fields should be visible.
	InputFieldDependencies []ExtensionActionDefinitionPatchInputFieldDependenciesInner `json:"inputFieldDependencies,omitempty"`
	// The user-facing labels for the custom action.
	Labels *map[string]ActionLabels `json:"labels,omitempty"`
	// The object types that this custom action supports.
	ObjectTypes []string `json:"objectTypes,omitempty"`
}

// NewExtensionActionDefinitionPatch instantiates a new ExtensionActionDefinitionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionActionDefinitionPatch() *ExtensionActionDefinitionPatch {
	this := ExtensionActionDefinitionPatch{}
	return &this
}

// NewExtensionActionDefinitionPatchWithDefaults instantiates a new ExtensionActionDefinitionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionActionDefinitionPatchWithDefaults() *ExtensionActionDefinitionPatch {
	this := ExtensionActionDefinitionPatch{}
	return &this
}

// GetActionUrl returns the ActionUrl field value if set, zero value otherwise.
func (o *ExtensionActionDefinitionPatch) GetActionUrl() string {
	if o == nil || isNil(o.ActionUrl) {
		var ret string
		return ret
	}
	return *o.ActionUrl
}

// GetActionUrlOk returns a tuple with the ActionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionActionDefinitionPatch) GetActionUrlOk() (*string, bool) {
	if o == nil || isNil(o.ActionUrl) {
		return nil, false
	}
	return o.ActionUrl, true
}

// HasActionUrl returns a boolean if a field has been set.
func (o *ExtensionActionDefinitionPatch) HasActionUrl() bool {
	if o != nil && !isNil(o.ActionUrl) {
		return true
	}

	return false
}

// SetActionUrl gets a reference to the given string and assigns it to the ActionUrl field.
func (o *ExtensionActionDefinitionPatch) SetActionUrl(v string) {
	o.ActionUrl = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *ExtensionActionDefinitionPatch) GetPublished() bool {
	if o == nil || isNil(o.Published) {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionActionDefinitionPatch) GetPublishedOk() (*bool, bool) {
	if o == nil || isNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *ExtensionActionDefinitionPatch) HasPublished() bool {
	if o != nil && !isNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *ExtensionActionDefinitionPatch) SetPublished(v bool) {
	o.Published = &v
}

// GetInputFields returns the InputFields field value if set, zero value otherwise.
func (o *ExtensionActionDefinitionPatch) GetInputFields() []InputFieldDefinition {
	if o == nil || isNil(o.InputFields) {
		var ret []InputFieldDefinition
		return ret
	}
	return o.InputFields
}

// GetInputFieldsOk returns a tuple with the InputFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionActionDefinitionPatch) GetInputFieldsOk() ([]InputFieldDefinition, bool) {
	if o == nil || isNil(o.InputFields) {
		return nil, false
	}
	return o.InputFields, true
}

// HasInputFields returns a boolean if a field has been set.
func (o *ExtensionActionDefinitionPatch) HasInputFields() bool {
	if o != nil && !isNil(o.InputFields) {
		return true
	}

	return false
}

// SetInputFields gets a reference to the given []InputFieldDefinition and assigns it to the InputFields field.
func (o *ExtensionActionDefinitionPatch) SetInputFields(v []InputFieldDefinition) {
	o.InputFields = v
}

// GetObjectRequestOptions returns the ObjectRequestOptions field value if set, zero value otherwise.
func (o *ExtensionActionDefinitionPatch) GetObjectRequestOptions() ObjectRequestOptions {
	if o == nil || isNil(o.ObjectRequestOptions) {
		var ret ObjectRequestOptions
		return ret
	}
	return *o.ObjectRequestOptions
}

// GetObjectRequestOptionsOk returns a tuple with the ObjectRequestOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionActionDefinitionPatch) GetObjectRequestOptionsOk() (*ObjectRequestOptions, bool) {
	if o == nil || isNil(o.ObjectRequestOptions) {
		return nil, false
	}
	return o.ObjectRequestOptions, true
}

// HasObjectRequestOptions returns a boolean if a field has been set.
func (o *ExtensionActionDefinitionPatch) HasObjectRequestOptions() bool {
	if o != nil && !isNil(o.ObjectRequestOptions) {
		return true
	}

	return false
}

// SetObjectRequestOptions gets a reference to the given ObjectRequestOptions and assigns it to the ObjectRequestOptions field.
func (o *ExtensionActionDefinitionPatch) SetObjectRequestOptions(v ObjectRequestOptions) {
	o.ObjectRequestOptions = &v
}

// GetInputFieldDependencies returns the InputFieldDependencies field value if set, zero value otherwise.
func (o *ExtensionActionDefinitionPatch) GetInputFieldDependencies() []ExtensionActionDefinitionPatchInputFieldDependenciesInner {
	if o == nil || isNil(o.InputFieldDependencies) {
		var ret []ExtensionActionDefinitionPatchInputFieldDependenciesInner
		return ret
	}
	return o.InputFieldDependencies
}

// GetInputFieldDependenciesOk returns a tuple with the InputFieldDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionActionDefinitionPatch) GetInputFieldDependenciesOk() ([]ExtensionActionDefinitionPatchInputFieldDependenciesInner, bool) {
	if o == nil || isNil(o.InputFieldDependencies) {
		return nil, false
	}
	return o.InputFieldDependencies, true
}

// HasInputFieldDependencies returns a boolean if a field has been set.
func (o *ExtensionActionDefinitionPatch) HasInputFieldDependencies() bool {
	if o != nil && !isNil(o.InputFieldDependencies) {
		return true
	}

	return false
}

// SetInputFieldDependencies gets a reference to the given []ExtensionActionDefinitionPatchInputFieldDependenciesInner and assigns it to the InputFieldDependencies field.
func (o *ExtensionActionDefinitionPatch) SetInputFieldDependencies(v []ExtensionActionDefinitionPatchInputFieldDependenciesInner) {
	o.InputFieldDependencies = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ExtensionActionDefinitionPatch) GetLabels() map[string]ActionLabels {
	if o == nil || isNil(o.Labels) {
		var ret map[string]ActionLabels
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionActionDefinitionPatch) GetLabelsOk() (*map[string]ActionLabels, bool) {
	if o == nil || isNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ExtensionActionDefinitionPatch) HasLabels() bool {
	if o != nil && !isNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]ActionLabels and assigns it to the Labels field.
func (o *ExtensionActionDefinitionPatch) SetLabels(v map[string]ActionLabels) {
	o.Labels = &v
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *ExtensionActionDefinitionPatch) GetObjectTypes() []string {
	if o == nil || isNil(o.ObjectTypes) {
		var ret []string
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionActionDefinitionPatch) GetObjectTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ObjectTypes) {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *ExtensionActionDefinitionPatch) HasObjectTypes() bool {
	if o != nil && !isNil(o.ObjectTypes) {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []string and assigns it to the ObjectTypes field.
func (o *ExtensionActionDefinitionPatch) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

func (o ExtensionActionDefinitionPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ActionUrl) {
		toSerialize["actionUrl"] = o.ActionUrl
	}
	if !isNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	if !isNil(o.InputFields) {
		toSerialize["inputFields"] = o.InputFields
	}
	if !isNil(o.ObjectRequestOptions) {
		toSerialize["objectRequestOptions"] = o.ObjectRequestOptions
	}
	if !isNil(o.InputFieldDependencies) {
		toSerialize["inputFieldDependencies"] = o.InputFieldDependencies
	}
	if !isNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !isNil(o.ObjectTypes) {
		toSerialize["objectTypes"] = o.ObjectTypes
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionActionDefinitionPatch struct {
	value *ExtensionActionDefinitionPatch
	isSet bool
}

func (v NullableExtensionActionDefinitionPatch) Get() *ExtensionActionDefinitionPatch {
	return v.value
}

func (v *NullableExtensionActionDefinitionPatch) Set(val *ExtensionActionDefinitionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionActionDefinitionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionActionDefinitionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionActionDefinitionPatch(val *ExtensionActionDefinitionPatch) *NullableExtensionActionDefinitionPatch {
	return &NullableExtensionActionDefinitionPatch{value: val, isSet: true}
}

func (v NullableExtensionActionDefinitionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionActionDefinitionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
