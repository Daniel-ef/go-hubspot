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

// InvoicePdfResponse A response that contains the PDF of an invoice
type InvoicePdfResponse struct {
	// Designates if the response is a success ('OK') or failure ('ERR').
	Result *string `json:"@result,omitempty"`
	// The bytes of the invoice PDF.
	Invoice string `json:"invoice"`
}

// NewInvoicePdfResponse instantiates a new InvoicePdfResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicePdfResponse(invoice string) *InvoicePdfResponse {
	this := InvoicePdfResponse{}
	this.Invoice = invoice
	return &this
}

// NewInvoicePdfResponseWithDefaults instantiates a new InvoicePdfResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicePdfResponseWithDefaults() *InvoicePdfResponse {
	this := InvoicePdfResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InvoicePdfResponse) GetResult() string {
	if o == nil || isNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePdfResponse) GetResultOk() (*string, bool) {
	if o == nil || isNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InvoicePdfResponse) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *InvoicePdfResponse) SetResult(v string) {
	o.Result = &v
}

// GetInvoice returns the Invoice field value
func (o *InvoicePdfResponse) GetInvoice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value
// and a boolean to check if the value has been set.
func (o *InvoicePdfResponse) GetInvoiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoice, true
}

// SetInvoice sets field value
func (o *InvoicePdfResponse) SetInvoice(v string) {
	o.Invoice = v
}

func (o InvoicePdfResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Result) {
		toSerialize["@result"] = o.Result
	}
	if true {
		toSerialize["invoice"] = o.Invoice
	}
	return json.Marshal(toSerialize)
}

type NullableInvoicePdfResponse struct {
	value *InvoicePdfResponse
	isSet bool
}

func (v NullableInvoicePdfResponse) Get() *InvoicePdfResponse {
	return v.value
}

func (v *NullableInvoicePdfResponse) Set(val *InvoicePdfResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicePdfResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicePdfResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicePdfResponse(val *InvoicePdfResponse) *NullableInvoicePdfResponse {
	return &NullableInvoicePdfResponse{value: val, isSet: true}
}

func (v NullableInvoicePdfResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicePdfResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
