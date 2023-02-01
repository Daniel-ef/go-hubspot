# \AssociationsBatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive**](AssociationsBatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/archive | Delete
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate**](AssociationsBatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/create | Create
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive**](AssociationsBatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/labels/archive | Delete Specific Labels
[**PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead**](AssociationsBatchApi.md#PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead) | **Post** /crm/v4/associations/{fromObjectType}/{toObjectType}/batch/read | Read



## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive

> PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive).Execute()

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
    batchInputPublicAssociationMultiArchive := *openapiclient.NewBatchInputPublicAssociationMultiArchive([]openapiclient.PublicAssociationMultiArchive{*openapiclient.NewPublicAssociationMultiArchive(*openapiclient.NewPublicObjectId("Id_example"), []openapiclient.PublicObjectId{*openapiclient.NewPublicObjectId("Id_example")})}) // BatchInputPublicAssociationMultiArchive | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiArchive** | [**BatchInputPublicAssociationMultiArchive**](BatchInputPublicAssociationMultiArchive.md) |  | 

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


## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate

> BatchResponseLabelsBetweenObjectPair PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

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
    batchInputPublicAssociationMultiPost := *openapiclient.NewBatchInputPublicAssociationMultiPost([]openapiclient.PublicAssociationMultiPost{*openapiclient.NewPublicAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}) // BatchInputPublicAssociationMultiPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate`: BatchResponseLabelsBetweenObjectPair
    fmt.Fprintf(os.Stdout, "Response from `AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiPost** | [**BatchInputPublicAssociationMultiPost**](BatchInputPublicAssociationMultiPost.md) |  | 

### Return type

[**BatchResponseLabelsBetweenObjectPair**](BatchResponseLabelsBetweenObjectPair.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive

> PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()

Delete Specific Labels



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
    batchInputPublicAssociationMultiPost := *openapiclient.NewBatchInputPublicAssociationMultiPost([]openapiclient.PublicAssociationMultiPost{*openapiclient.NewPublicAssociationMultiPost(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))})}) // BatchInputPublicAssociationMultiPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociationMultiPost** | [**BatchInputPublicAssociationMultiPost**](BatchInputPublicAssociationMultiPost.md) |  | 

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


## PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead

> BatchResponsePublicAssociationMultiWithLabel PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead(ctx, fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()

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
    batchInputPublicFetchAssociationsBatchRequest := *openapiclient.NewBatchInputPublicFetchAssociationsBatchRequest([]openapiclient.PublicFetchAssociationsBatchRequest{*openapiclient.NewPublicFetchAssociationsBatchRequest("Id_example")}) // BatchInputPublicFetchAssociationsBatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead(context.Background(), fromObjectType, toObjectType).BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead`: BatchResponsePublicAssociationMultiWithLabel
    fmt.Fprintf(os.Stdout, "Response from `AssociationsBatchApi.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicFetchAssociationsBatchRequest** | [**BatchInputPublicFetchAssociationsBatchRequest**](BatchInputPublicFetchAssociationsBatchRequest.md) |  | 

### Return type

[**BatchResponsePublicAssociationMultiWithLabel**](BatchResponsePublicAssociationMultiWithLabel.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

