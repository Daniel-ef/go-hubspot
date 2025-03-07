/*
CRM Imports

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imports

import (
	"encoding/json"
	"time"
)

// PublicImportResponse A current summary of the import, whether complete or not.
type PublicImportResponse struct {
	// The status of the import.
	State             string                 `json:"state"`
	ImportRequestJson map[string]interface{} `json:"importRequestJson,omitempty"`
	CreatedAt         time.Time              `json:"createdAt"`
	Metadata          PublicImportMetadata   `json:"metadata"`
	ImportName        *string                `json:"importName,omitempty"`
	UpdatedAt         time.Time              `json:"updatedAt"`
	// Whether or not the import is a list of people disqualified from receiving emails.
	OptOutImport bool   `json:"optOutImport"`
	Id           string `json:"id"`
}

// NewPublicImportResponse instantiates a new PublicImportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicImportResponse(state string, createdAt time.Time, metadata PublicImportMetadata, updatedAt time.Time, optOutImport bool, id string) *PublicImportResponse {
	this := PublicImportResponse{}
	this.State = state
	this.CreatedAt = createdAt
	this.Metadata = metadata
	this.UpdatedAt = updatedAt
	this.OptOutImport = optOutImport
	this.Id = id
	return &this
}

// NewPublicImportResponseWithDefaults instantiates a new PublicImportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicImportResponseWithDefaults() *PublicImportResponse {
	this := PublicImportResponse{}
	return &this
}

// GetState returns the State field value
func (o *PublicImportResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PublicImportResponse) SetState(v string) {
	o.State = v
}

// GetImportRequestJson returns the ImportRequestJson field value if set, zero value otherwise.
func (o *PublicImportResponse) GetImportRequestJson() map[string]interface{} {
	if o == nil || isNil(o.ImportRequestJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.ImportRequestJson
}

// GetImportRequestJsonOk returns a tuple with the ImportRequestJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetImportRequestJsonOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ImportRequestJson) {
		return map[string]interface{}{}, false
	}
	return o.ImportRequestJson, true
}

// HasImportRequestJson returns a boolean if a field has been set.
func (o *PublicImportResponse) HasImportRequestJson() bool {
	if o != nil && !isNil(o.ImportRequestJson) {
		return true
	}

	return false
}

// SetImportRequestJson gets a reference to the given map[string]interface{} and assigns it to the ImportRequestJson field.
func (o *PublicImportResponse) SetImportRequestJson(v map[string]interface{}) {
	o.ImportRequestJson = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PublicImportResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PublicImportResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMetadata returns the Metadata field value
func (o *PublicImportResponse) GetMetadata() PublicImportMetadata {
	if o == nil {
		var ret PublicImportMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetMetadataOk() (*PublicImportMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *PublicImportResponse) SetMetadata(v PublicImportMetadata) {
	o.Metadata = v
}

// GetImportName returns the ImportName field value if set, zero value otherwise.
func (o *PublicImportResponse) GetImportName() string {
	if o == nil || isNil(o.ImportName) {
		var ret string
		return ret
	}
	return *o.ImportName
}

// GetImportNameOk returns a tuple with the ImportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetImportNameOk() (*string, bool) {
	if o == nil || isNil(o.ImportName) {
		return nil, false
	}
	return o.ImportName, true
}

// HasImportName returns a boolean if a field has been set.
func (o *PublicImportResponse) HasImportName() bool {
	if o != nil && !isNil(o.ImportName) {
		return true
	}

	return false
}

// SetImportName gets a reference to the given string and assigns it to the ImportName field.
func (o *PublicImportResponse) SetImportName(v string) {
	o.ImportName = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PublicImportResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PublicImportResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetOptOutImport returns the OptOutImport field value
func (o *PublicImportResponse) GetOptOutImport() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OptOutImport
}

// GetOptOutImportOk returns a tuple with the OptOutImport field value
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetOptOutImportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptOutImport, true
}

// SetOptOutImport sets field value
func (o *PublicImportResponse) SetOptOutImport(v bool) {
	o.OptOutImport = v
}

// GetId returns the Id field value
func (o *PublicImportResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicImportResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicImportResponse) SetId(v string) {
	o.Id = v
}

func (o PublicImportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if !isNil(o.ImportRequestJson) {
		toSerialize["importRequestJson"] = o.ImportRequestJson
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.ImportName) {
		toSerialize["importName"] = o.ImportName
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["optOutImport"] = o.OptOutImport
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePublicImportResponse struct {
	value *PublicImportResponse
	isSet bool
}

func (v NullablePublicImportResponse) Get() *PublicImportResponse {
	return v.value
}

func (v *NullablePublicImportResponse) Set(val *PublicImportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicImportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicImportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicImportResponse(val *PublicImportResponse) *NullablePublicImportResponse {
	return &NullablePublicImportResponse{value: val, isSet: true}
}

func (v NullablePublicImportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicImportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
