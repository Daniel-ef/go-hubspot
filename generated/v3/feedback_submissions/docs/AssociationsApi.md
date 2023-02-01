# \AssociationsApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociationsGetPage**](AssociationsApi.md#AssociationsGetPage) | **Get** /crm/v4/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType} | List
[**DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive**](AssociationsApi.md#DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive) | **Delete** /crm/v4/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId} | Delete
[**PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate**](AssociationsApi.md#PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate) | **Put** /crm/v4/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId} | Create



## AssociationsGetPage

> CollectionResponseMultiAssociatedObjectWithLabelForwardPaging AssociationsGetPage(ctx, feedbackSubmissionId, toObjectType).After(after).Limit(limit).Execute()

List



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
    feedbackSubmissionId := int64(789) // int64 | 
    toObjectType := "toObjectType_example" // string | 
    after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
    limit := int32(56) // int32 | The maximum number of results to display per page. (optional) (default to 500)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.AssociationsGetPage(context.Background(), feedbackSubmissionId, toObjectType).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.AssociationsGetPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssociationsGetPage`: CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.AssociationsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackSubmissionId** | **int64** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociationsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to display per page. | [default to 500]

### Return type

[**CollectionResponseMultiAssociatedObjectWithLabelForwardPaging**](CollectionResponseMultiAssociatedObjectWithLabelForwardPaging.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive

> DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive(ctx, feedbackSubmissionId, toObjectType, toObjectId).Execute()

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
    feedbackSubmissionId := int64(789) // int64 | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive(context.Background(), feedbackSubmissionId, toObjectType, toObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackSubmissionId** | **int64** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveRequest struct via the builder pattern


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


## PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate

> LabelsBetweenObjectPair PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate(ctx, feedbackSubmissionId, toObjectType, toObjectId).AssociationSpec(associationSpec).Execute()

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
    feedbackSubmissionId := int64(789) // int64 | 
    toObjectType := "toObjectType_example" // string | 
    toObjectId := int64(789) // int64 | 
    associationSpec := []openapiclient.AssociationSpec{*openapiclient.NewAssociationSpec("AssociationCategory_example", int32(123))} // []AssociationSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssociationsApi.PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate(context.Background(), feedbackSubmissionId, toObjectType, toObjectId).AssociationSpec(associationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssociationsApi.PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate`: LabelsBetweenObjectPair
    fmt.Fprintf(os.Stdout, "Response from `AssociationsApi.PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedbackSubmissionId** | **int64** |  | 
**toObjectType** | **string** |  | 
**toObjectId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **associationSpec** | [**[]AssociationSpec**](AssociationSpec.md) |  | 

### Return type

[**LabelsBetweenObjectPair**](LabelsBetweenObjectPair.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

