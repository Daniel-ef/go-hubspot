# \AssociationDefinitionsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeId**](AssociationDefinitionsApi.md#DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeId) | **Delete** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels/{associationTypeId} | Delete
[**GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels**](AssociationDefinitionsApi.md#GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels) | **Get** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels | Read
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels**](AssociationDefinitionsApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels | Create
[**PutCrmV4AssociationsFromObjectTypeToObjectTypeLabels**](AssociationDefinitionsApi.md#PutCrmV4AssociationsFromObjectTypeToObjectTypeLabels) | **Put** /crm/v4/associations/{fromObjectType}/{toObjectType}/labels | Update



## DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeId

> DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeId(ctx, fromObjectType, toObjectType, associationTypeId).Execute()

Delete



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fromObjectType := "fromObjectType_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    associationTypeId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationDefinitionsApi.DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeId(context.Background(), fromObjectType, toObjectType, associationTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationDefinitionsApi.DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 
**associationTypeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels

> CollectionResponseAssociationSpecWithLabelNoPaging GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels(ctx, fromObjectType, toObjectType).Execute()

Read



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fromObjectType := "fromObjectType_example" // string | 
    toObjectType := "toObjectType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationDefinitionsApi.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels(context.Background(), fromObjectType, toObjectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationDefinitionsApi.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels`: CollectionResponseAssociationSpecWithLabelNoPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationDefinitionsApi.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponseAssociationSpecWithLabelNoPaging**](CollectionResponseAssociationSpecWithLabelNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels

> CollectionResponseAssociationSpecWithLabelNoPaging PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()

Create



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fromObjectType := "fromObjectType_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    publicAssociationDefinitionCreateRequest := *openapiclient.NewPublicAssociationDefinitionCreateRequest("Label_example", "Name_example") // PublicAssociationDefinitionCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationDefinitionsApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationDefinitionsApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels`: CollectionResponseAssociationSpecWithLabelNoPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationDefinitionsApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicAssociationDefinitionCreateRequest** | [**PublicAssociationDefinitionCreateRequest**](PublicAssociationDefinitionCreateRequest.md) |  | 

### Return type

[**CollectionResponseAssociationSpecWithLabelNoPaging**](CollectionResponseAssociationSpecWithLabelNoPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV4AssociationsFromObjectTypeToObjectTypeLabels

> PutCrmV4AssociationsFromObjectTypeToObjectTypeLabels(ctx, fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()

Update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fromObjectType := "fromObjectType_example" // string | 
    toObjectType := "toObjectType_example" // string | 
    publicAssociationDefinitionUpdateRequest := *openapiclient.NewPublicAssociationDefinitionUpdateRequest("Label_example", int32(123)) // PublicAssociationDefinitionUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationDefinitionsApi.PutCrmV4AssociationsFromObjectTypeToObjectTypeLabels(context.Background(), fromObjectType, toObjectType).PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationDefinitionsApi.PutCrmV4AssociationsFromObjectTypeToObjectTypeLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicAssociationDefinitionUpdateRequest** | [**PublicAssociationDefinitionUpdateRequest**](PublicAssociationDefinitionUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

