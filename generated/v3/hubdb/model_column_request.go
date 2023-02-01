/*
HubDB endpoints

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"encoding/json"
)

// ColumnRequest struct for ColumnRequest
type ColumnRequest struct {
	// Column Id
	Id int32 `json:"id"`
	// Name of the column
	Name string `json:"name"`
	// Label of the column
	Label string `json:"label"`
	// Type of the column
	Type string `json:"type"`
	// Options to choose for select and multi-select columns
	Options []Option `json:"options"`
	// The id of another table to which the column refers/points to.
	ForeignTableId *int64 `json:"foreignTableId,omitempty"`
	// The id of the column from another table to which the column refers/points to.
	ForeignColumnId *int32 `json:"foreignColumnId,omitempty"`
}

// NewColumnRequest instantiates a new ColumnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumnRequest(id int32, name string, label string, type_ string, options []Option) *ColumnRequest {
	this := ColumnRequest{}
	this.Id = id
	this.Name = name
	this.Label = label
	this.Type = type_
	this.Options = options
	return &this
}

// NewColumnRequestWithDefaults instantiates a new ColumnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnRequestWithDefaults() *ColumnRequest {
	this := ColumnRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ColumnRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ColumnRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ColumnRequest) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ColumnRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ColumnRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ColumnRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *ColumnRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ColumnRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ColumnRequest) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *ColumnRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ColumnRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ColumnRequest) SetType(v string) {
	o.Type = v
}

// GetOptions returns the Options field value
func (o *ColumnRequest) GetOptions() []Option {
	if o == nil {
		var ret []Option
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ColumnRequest) GetOptionsOk() ([]Option, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *ColumnRequest) SetOptions(v []Option) {
	o.Options = v
}

// GetForeignTableId returns the ForeignTableId field value if set, zero value otherwise.
func (o *ColumnRequest) GetForeignTableId() int64 {
	if o == nil || isNil(o.ForeignTableId) {
		var ret int64
		return ret
	}
	return *o.ForeignTableId
}

// GetForeignTableIdOk returns a tuple with the ForeignTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnRequest) GetForeignTableIdOk() (*int64, bool) {
	if o == nil || isNil(o.ForeignTableId) {
		return nil, false
	}
	return o.ForeignTableId, true
}

// HasForeignTableId returns a boolean if a field has been set.
func (o *ColumnRequest) HasForeignTableId() bool {
	if o != nil && !isNil(o.ForeignTableId) {
		return true
	}

	return false
}

// SetForeignTableId gets a reference to the given int64 and assigns it to the ForeignTableId field.
func (o *ColumnRequest) SetForeignTableId(v int64) {
	o.ForeignTableId = &v
}

// GetForeignColumnId returns the ForeignColumnId field value if set, zero value otherwise.
func (o *ColumnRequest) GetForeignColumnId() int32 {
	if o == nil || isNil(o.ForeignColumnId) {
		var ret int32
		return ret
	}
	return *o.ForeignColumnId
}

// GetForeignColumnIdOk returns a tuple with the ForeignColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnRequest) GetForeignColumnIdOk() (*int32, bool) {
	if o == nil || isNil(o.ForeignColumnId) {
		return nil, false
	}
	return o.ForeignColumnId, true
}

// HasForeignColumnId returns a boolean if a field has been set.
func (o *ColumnRequest) HasForeignColumnId() bool {
	if o != nil && !isNil(o.ForeignColumnId) {
		return true
	}

	return false
}

// SetForeignColumnId gets a reference to the given int32 and assigns it to the ForeignColumnId field.
func (o *ColumnRequest) SetForeignColumnId(v int32) {
	o.ForeignColumnId = &v
}

func (o ColumnRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["options"] = o.Options
	}
	if !isNil(o.ForeignTableId) {
		toSerialize["foreignTableId"] = o.ForeignTableId
	}
	if !isNil(o.ForeignColumnId) {
		toSerialize["foreignColumnId"] = o.ForeignColumnId
	}
	return json.Marshal(toSerialize)
}

type NullableColumnRequest struct {
	value *ColumnRequest
	isSet bool
}

func (v NullableColumnRequest) Get() *ColumnRequest {
	return v.value
}

func (v *NullableColumnRequest) Set(val *ColumnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableColumnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableColumnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumnRequest(val *ColumnRequest) *NullableColumnRequest {
	return &NullableColumnRequest{value: val, isSet: true}
}

func (v NullableColumnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
