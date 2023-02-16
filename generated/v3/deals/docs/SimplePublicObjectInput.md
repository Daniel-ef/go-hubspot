# SimplePublicObjectInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | **map[string]string** |  | 
**Associations** | Pointer to [**[]AssociationInput**](AssociationInput.md) |  | [optional] 

## Methods

### NewSimplePublicObjectInput

`func NewSimplePublicObjectInput(properties map[string]string, ) *SimplePublicObjectInput`

NewSimplePublicObjectInput instantiates a new SimplePublicObjectInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplePublicObjectInputWithDefaults

`func NewSimplePublicObjectInputWithDefaults() *SimplePublicObjectInput`

NewSimplePublicObjectInputWithDefaults instantiates a new SimplePublicObjectInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SimplePublicObjectInput) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SimplePublicObjectInput) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SimplePublicObjectInput) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetAssociations

`func (o *SimplePublicObjectInput) GetAssociations() []AssociationInput`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *SimplePublicObjectInput) GetAssociationsOk() (*[]AssociationInput, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *SimplePublicObjectInput) SetAssociations(v []AssociationInput)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *SimplePublicObjectInput) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


