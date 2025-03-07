/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"encoding/json"
)

// InvoiceUpdateRequest The invoice data to update in HubSpot
type InvoiceUpdateRequest struct {
	ExternalInvoiceNumber *string `json:"externalInvoiceNumber,omitempty"`
	// The ISO 4217 currency code that represents the currency used in the invoice to bill the recipient
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The ISO-8601 due date of the invoice.
	DueDate *string `json:"dueDate,omitempty"`
	// The ID of the invoice recipient. This is the recipient ID from the external accounting system.
	ExternalRecipientId     *string `json:"externalRecipientId,omitempty"`
	ReceivedByRecipientDate *int64  `json:"receivedByRecipientDate,omitempty"`
	// States if the invoice is voided or not.
	IsVoided *bool `json:"isVoided,omitempty"`
	// The ISO-8601 datetime of when the customer received the invoice.
	ReceivedByCustomerDate *string `json:"receivedByCustomerDate,omitempty"`
	// The number / name of the invoice.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
}

// NewInvoiceUpdateRequest instantiates a new InvoiceUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceUpdateRequest() *InvoiceUpdateRequest {
	this := InvoiceUpdateRequest{}
	return &this
}

// NewInvoiceUpdateRequestWithDefaults instantiates a new InvoiceUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceUpdateRequestWithDefaults() *InvoiceUpdateRequest {
	this := InvoiceUpdateRequest{}
	return &this
}

// GetExternalInvoiceNumber returns the ExternalInvoiceNumber field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetExternalInvoiceNumber() string {
	if o == nil || isNil(o.ExternalInvoiceNumber) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceNumber
}

// GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetExternalInvoiceNumberOk() (*string, bool) {
	if o == nil || isNil(o.ExternalInvoiceNumber) {
		return nil, false
	}
	return o.ExternalInvoiceNumber, true
}

// HasExternalInvoiceNumber returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasExternalInvoiceNumber() bool {
	if o != nil && !isNil(o.ExternalInvoiceNumber) {
		return true
	}

	return false
}

// SetExternalInvoiceNumber gets a reference to the given string and assigns it to the ExternalInvoiceNumber field.
func (o *InvoiceUpdateRequest) SetExternalInvoiceNumber(v string) {
	o.ExternalInvoiceNumber = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetCurrencyCode() string {
	if o == nil || isNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || isNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasCurrencyCode() bool {
	if o != nil && !isNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *InvoiceUpdateRequest) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetDueDate() string {
	if o == nil || isNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetDueDateOk() (*string, bool) {
	if o == nil || isNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasDueDate() bool {
	if o != nil && !isNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *InvoiceUpdateRequest) SetDueDate(v string) {
	o.DueDate = &v
}

// GetExternalRecipientId returns the ExternalRecipientId field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetExternalRecipientId() string {
	if o == nil || isNil(o.ExternalRecipientId) {
		var ret string
		return ret
	}
	return *o.ExternalRecipientId
}

// GetExternalRecipientIdOk returns a tuple with the ExternalRecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetExternalRecipientIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalRecipientId) {
		return nil, false
	}
	return o.ExternalRecipientId, true
}

// HasExternalRecipientId returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasExternalRecipientId() bool {
	if o != nil && !isNil(o.ExternalRecipientId) {
		return true
	}

	return false
}

// SetExternalRecipientId gets a reference to the given string and assigns it to the ExternalRecipientId field.
func (o *InvoiceUpdateRequest) SetExternalRecipientId(v string) {
	o.ExternalRecipientId = &v
}

// GetReceivedByRecipientDate returns the ReceivedByRecipientDate field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetReceivedByRecipientDate() int64 {
	if o == nil || isNil(o.ReceivedByRecipientDate) {
		var ret int64
		return ret
	}
	return *o.ReceivedByRecipientDate
}

// GetReceivedByRecipientDateOk returns a tuple with the ReceivedByRecipientDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetReceivedByRecipientDateOk() (*int64, bool) {
	if o == nil || isNil(o.ReceivedByRecipientDate) {
		return nil, false
	}
	return o.ReceivedByRecipientDate, true
}

// HasReceivedByRecipientDate returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasReceivedByRecipientDate() bool {
	if o != nil && !isNil(o.ReceivedByRecipientDate) {
		return true
	}

	return false
}

// SetReceivedByRecipientDate gets a reference to the given int64 and assigns it to the ReceivedByRecipientDate field.
func (o *InvoiceUpdateRequest) SetReceivedByRecipientDate(v int64) {
	o.ReceivedByRecipientDate = &v
}

// GetIsVoided returns the IsVoided field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetIsVoided() bool {
	if o == nil || isNil(o.IsVoided) {
		var ret bool
		return ret
	}
	return *o.IsVoided
}

// GetIsVoidedOk returns a tuple with the IsVoided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetIsVoidedOk() (*bool, bool) {
	if o == nil || isNil(o.IsVoided) {
		return nil, false
	}
	return o.IsVoided, true
}

// HasIsVoided returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasIsVoided() bool {
	if o != nil && !isNil(o.IsVoided) {
		return true
	}

	return false
}

// SetIsVoided gets a reference to the given bool and assigns it to the IsVoided field.
func (o *InvoiceUpdateRequest) SetIsVoided(v bool) {
	o.IsVoided = &v
}

// GetReceivedByCustomerDate returns the ReceivedByCustomerDate field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetReceivedByCustomerDate() string {
	if o == nil || isNil(o.ReceivedByCustomerDate) {
		var ret string
		return ret
	}
	return *o.ReceivedByCustomerDate
}

// GetReceivedByCustomerDateOk returns a tuple with the ReceivedByCustomerDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetReceivedByCustomerDateOk() (*string, bool) {
	if o == nil || isNil(o.ReceivedByCustomerDate) {
		return nil, false
	}
	return o.ReceivedByCustomerDate, true
}

// HasReceivedByCustomerDate returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasReceivedByCustomerDate() bool {
	if o != nil && !isNil(o.ReceivedByCustomerDate) {
		return true
	}

	return false
}

// SetReceivedByCustomerDate gets a reference to the given string and assigns it to the ReceivedByCustomerDate field.
func (o *InvoiceUpdateRequest) SetReceivedByCustomerDate(v string) {
	o.ReceivedByCustomerDate = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *InvoiceUpdateRequest) GetInvoiceNumber() string {
	if o == nil || isNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUpdateRequest) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || isNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *InvoiceUpdateRequest) HasInvoiceNumber() bool {
	if o != nil && !isNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *InvoiceUpdateRequest) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

func (o InvoiceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalInvoiceNumber) {
		toSerialize["externalInvoiceNumber"] = o.ExternalInvoiceNumber
	}
	if !isNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !isNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !isNil(o.ExternalRecipientId) {
		toSerialize["externalRecipientId"] = o.ExternalRecipientId
	}
	if !isNil(o.ReceivedByRecipientDate) {
		toSerialize["receivedByRecipientDate"] = o.ReceivedByRecipientDate
	}
	if !isNil(o.IsVoided) {
		toSerialize["isVoided"] = o.IsVoided
	}
	if !isNil(o.ReceivedByCustomerDate) {
		toSerialize["receivedByCustomerDate"] = o.ReceivedByCustomerDate
	}
	if !isNil(o.InvoiceNumber) {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	return json.Marshal(toSerialize)
}

type NullableInvoiceUpdateRequest struct {
	value *InvoiceUpdateRequest
	isSet bool
}

func (v NullableInvoiceUpdateRequest) Get() *InvoiceUpdateRequest {
	return v.value
}

func (v *NullableInvoiceUpdateRequest) Set(val *InvoiceUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceUpdateRequest(val *InvoiceUpdateRequest) *NullableInvoiceUpdateRequest {
	return &NullableInvoiceUpdateRequest{value: val, isSet: true}
}

func (v NullableInvoiceUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
