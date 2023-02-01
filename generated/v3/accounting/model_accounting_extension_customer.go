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

// AccountingExtensionCustomer Representation of a customer in the external accounting system.
type AccountingExtensionCustomer struct {
	// The customer's email address
	EmailAddress *string `json:"emailAddress,omitempty"`
	// The customer's full name
	Name string `json:"name"`
	// The ID of the customer in the external accounting system.
	Id             string   `json:"id"`
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// The ISO 4217 currency code that represents the currency the customer should be billed in.
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewAccountingExtensionCustomer instantiates a new AccountingExtensionCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountingExtensionCustomer(name string, id string) *AccountingExtensionCustomer {
	this := AccountingExtensionCustomer{}
	this.Name = name
	this.Id = id
	return &this
}

// NewAccountingExtensionCustomerWithDefaults instantiates a new AccountingExtensionCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountingExtensionCustomerWithDefaults() *AccountingExtensionCustomer {
	this := AccountingExtensionCustomer{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *AccountingExtensionCustomer) GetEmailAddress() string {
	if o == nil || isNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountingExtensionCustomer) GetEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *AccountingExtensionCustomer) HasEmailAddress() bool {
	if o != nil && !isNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *AccountingExtensionCustomer) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetName returns the Name field value
func (o *AccountingExtensionCustomer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionCustomer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountingExtensionCustomer) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *AccountingExtensionCustomer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionCustomer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountingExtensionCustomer) SetId(v string) {
	o.Id = v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *AccountingExtensionCustomer) GetBillingAddress() Address {
	if o == nil || isNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountingExtensionCustomer) GetBillingAddressOk() (*Address, bool) {
	if o == nil || isNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *AccountingExtensionCustomer) HasBillingAddress() bool {
	if o != nil && !isNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *AccountingExtensionCustomer) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AccountingExtensionCustomer) GetCurrencyCode() string {
	if o == nil || isNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountingExtensionCustomer) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || isNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AccountingExtensionCustomer) HasCurrencyCode() bool {
	if o != nil && !isNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AccountingExtensionCustomer) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o AccountingExtensionCustomer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !isNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableAccountingExtensionCustomer struct {
	value *AccountingExtensionCustomer
	isSet bool
}

func (v NullableAccountingExtensionCustomer) Get() *AccountingExtensionCustomer {
	return v.value
}

func (v *NullableAccountingExtensionCustomer) Set(val *AccountingExtensionCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountingExtensionCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountingExtensionCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountingExtensionCustomer(val *AccountingExtensionCustomer) *NullableAccountingExtensionCustomer {
	return &NullableAccountingExtensionCustomer{value: val, isSet: true}
}

func (v NullableAccountingExtensionCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountingExtensionCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
