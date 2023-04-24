/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"encoding/json"
)

// AssociationInput struct for AssociationInput
type AssociationInput struct {
	To    AssociationInputTo     `json:"to"`
	Types []AssociationTypeInput `json:"types"`
}

// NewAssociationInput instantiates a new AssociationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationInput(to AssociationInputTo, types []AssociationTypeInput) *AssociationInput {
	this := AssociationInput{}
	this.To = to
	this.Types = types
	return &this
}

// NewAssociationInputWithDefaults instantiates a new AssociationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationInputWithDefaults() *AssociationInput {
	this := AssociationInput{}
	return &this
}

// GetTo returns the To field value
func (o *AssociationInput) GetTo() AssociationInputTo {
	if o == nil {
		var ret AssociationInputTo
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *AssociationInput) GetToOk() (*AssociationInputTo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *AssociationInput) SetTo(v AssociationInputTo) {
	o.To = v
}

// GetTypes returns the Types field value
func (o *AssociationInput) GetTypes() []AssociationTypeInput {
	if o == nil {
		var ret []AssociationTypeInput
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *AssociationInput) GetTypesOk() ([]AssociationTypeInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *AssociationInput) SetTypes(v []AssociationTypeInput) {
	o.Types = v
}

func (o AssociationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableAssociationInput struct {
	value *AssociationInput
	isSet bool
}

func (v NullableAssociationInput) Get() *AssociationInput {
	return v.value
}

func (v *NullableAssociationInput) Set(val *AssociationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationInput(val *AssociationInput) *NullableAssociationInput {
	return &NullableAssociationInput{value: val, isSet: true}
}

func (v NullableAssociationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
