/*
CRM Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"encoding/json"
	"time"
)

// Pipeline A pipeline definition.
type Pipeline struct {
	// A unique label used to organize pipelines in HubSpot's UI
	Label string `json:"label"`
	// The order for displaying this pipeline. If two pipelines have a matching `displayOrder`, they will be sorted alphabetically by label.
	DisplayOrder int32 `json:"displayOrder"`
	// A unique identifier generated by HubSpot that can be used to retrieve and update the pipeline.
	Id string `json:"id"`
	// The stages associated with the pipeline. They can be retrieved and updated via the pipeline stages endpoints.
	Stages []PipelineStage `json:"stages"`
	// The date the pipeline was created. The default pipelines will have createdAt = 0.
	CreatedAt time.Time `json:"createdAt"`
	// The date the pipeline was archived. `archivedAt` will only be present if the pipeline is archived.
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
	// The date the pipeline was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// Whether the pipeline is archived.
	Archived bool `json:"archived"`
}

// NewPipeline instantiates a new Pipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipeline(label string, displayOrder int32, id string, stages []PipelineStage, createdAt time.Time, updatedAt time.Time, archived bool) *Pipeline {
	this := Pipeline{}
	this.Label = label
	this.DisplayOrder = displayOrder
	this.Id = id
	this.Stages = stages
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Archived = archived
	return &this
}

// NewPipelineWithDefaults instantiates a new Pipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineWithDefaults() *Pipeline {
	this := Pipeline{}
	return &this
}

// GetLabel returns the Label field value
func (o *Pipeline) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Pipeline) SetLabel(v string) {
	o.Label = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *Pipeline) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *Pipeline) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetId returns the Id field value
func (o *Pipeline) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Pipeline) SetId(v string) {
	o.Id = v
}

// GetStages returns the Stages field value
func (o *Pipeline) GetStages() []PipelineStage {
	if o == nil {
		var ret []PipelineStage
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetStagesOk() ([]PipelineStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *Pipeline) SetStages(v []PipelineStage) {
	o.Stages = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Pipeline) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Pipeline) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Pipeline) GetArchivedAt() time.Time {
	if o == nil || isNil(o.ArchivedAt) {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Pipeline) HasArchivedAt() bool {
	if o != nil && !isNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *Pipeline) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Pipeline) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Pipeline) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetArchived returns the Archived field value
func (o *Pipeline) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *Pipeline) SetArchived(v bool) {
	o.Archived = v
}

func (o Pipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["stages"] = o.Stages
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["archived"] = o.Archived
	}
	return json.Marshal(toSerialize)
}

type NullablePipeline struct {
	value *Pipeline
	isSet bool
}

func (v NullablePipeline) Get() *Pipeline {
	return v.value
}

func (v *NullablePipeline) Set(val *Pipeline) {
	v.value = val
	v.isSet = true
}

func (v NullablePipeline) IsSet() bool {
	return v.isSet
}

func (v *NullablePipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipeline(val *Pipeline) *NullablePipeline {
	return &NullablePipeline{value: val, isSet: true}
}

func (v NullablePipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
