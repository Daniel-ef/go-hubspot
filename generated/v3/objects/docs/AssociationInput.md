# AssociationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**AssociationInputTo**](AssociationInputTo.md) |  | 
**Types** | [**[]AssociationTypeInput**](AssociationTypeInput.md) |  | 

## Methods

### NewAssociationInput

`func NewAssociationInput(to AssociationInputTo, types []AssociationTypeInput, ) *AssociationInput`

NewAssociationInput instantiates a new AssociationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationInputWithDefaults

`func NewAssociationInputWithDefaults() *AssociationInput`

NewAssociationInputWithDefaults instantiates a new AssociationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *AssociationInput) GetTo() AssociationInputTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AssociationInput) GetToOk() (*AssociationInputTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AssociationInput) SetTo(v AssociationInputTo)`

SetTo sets To field to given value.


### GetTypes

`func (o *AssociationInput) GetTypes() []AssociationTypeInput`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *AssociationInput) GetTypesOk() (*[]AssociationTypeInput, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *AssociationInput) SetTypes(v []AssociationTypeInput)`

SetTypes sets Types field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


