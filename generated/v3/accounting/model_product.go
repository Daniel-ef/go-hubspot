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

// Product Representation of a product in the external accounting system.
type Product struct {
	UnitPrice UnitPrice `json:"unitPrice"`
	// Identifies if the product is tax exempt or not.
	TaxExempt    bool     `json:"taxExempt"`
	SalesTaxType *TaxType `json:"salesTaxType,omitempty"`
	// The display name of the product.
	Name string `json:"name"`
	// The description of the product.
	Description *string `json:"description,omitempty"`
	// The ID of the product in the external accounting system.
	Id string `json:"id"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(unitPrice UnitPrice, taxExempt bool, name string, id string) *Product {
	this := Product{}
	this.UnitPrice = unitPrice
	this.TaxExempt = taxExempt
	this.Name = name
	this.Id = id
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetUnitPrice returns the UnitPrice field value
func (o *Product) GetUnitPrice() UnitPrice {
	if o == nil {
		var ret UnitPrice
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *Product) GetUnitPriceOk() (*UnitPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *Product) SetUnitPrice(v UnitPrice) {
	o.UnitPrice = v
}

// GetTaxExempt returns the TaxExempt field value
func (o *Product) GetTaxExempt() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TaxExempt
}

// GetTaxExemptOk returns a tuple with the TaxExempt field value
// and a boolean to check if the value has been set.
func (o *Product) GetTaxExemptOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxExempt, true
}

// SetTaxExempt sets field value
func (o *Product) SetTaxExempt(v bool) {
	o.TaxExempt = v
}

// GetSalesTaxType returns the SalesTaxType field value if set, zero value otherwise.
func (o *Product) GetSalesTaxType() TaxType {
	if o == nil || isNil(o.SalesTaxType) {
		var ret TaxType
		return ret
	}
	return *o.SalesTaxType
}

// GetSalesTaxTypeOk returns a tuple with the SalesTaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetSalesTaxTypeOk() (*TaxType, bool) {
	if o == nil || isNil(o.SalesTaxType) {
		return nil, false
	}
	return o.SalesTaxType, true
}

// HasSalesTaxType returns a boolean if a field has been set.
func (o *Product) HasSalesTaxType() bool {
	if o != nil && !isNil(o.SalesTaxType) {
		return true
	}

	return false
}

// SetSalesTaxType gets a reference to the given TaxType and assigns it to the SalesTaxType field.
func (o *Product) SetSalesTaxType(v TaxType) {
	o.SalesTaxType = &v
}

// GetName returns the Name field value
func (o *Product) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Product) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Product) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Product) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Product) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *Product) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Product) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Product) SetId(v string) {
	o.Id = v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if true {
		toSerialize["taxExempt"] = o.TaxExempt
	}
	if !isNil(o.SalesTaxType) {
		toSerialize["salesTaxType"] = o.SalesTaxType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
