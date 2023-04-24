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

// AssociationInputTo struct for AssociationInputTo
type AssociationInputTo struct {
	Id string `json:"id"`
}

// NewAssociationInputTo instantiates a new AssociationInputTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationInputTo(id string) *AssociationInputTo {
	this := AssociationInputTo{}
	this.Id = id
	return &this
}

// NewAssociationInputToWithDefaults instantiates a new AssociationInputTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationInputToWithDefaults() *AssociationInputTo {
	this := AssociationInputTo{}
	return &this
}

// GetId returns the Id field value
func (o *AssociationInputTo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssociationInputTo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssociationInputTo) SetId(v string) {
	o.Id = v
}

func (o AssociationInputTo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAssociationInputTo struct {
	value *AssociationInputTo
	isSet bool
}

func (v NullableAssociationInputTo) Get() *AssociationInputTo {
	return v.value
}

func (v *NullableAssociationInputTo) Set(val *AssociationInputTo) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationInputTo) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationInputTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationInputTo(val *AssociationInputTo) *NullableAssociationInputTo {
	return &NullableAssociationInputTo{value: val, isSet: true}
}

func (v NullableAssociationInputTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationInputTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
