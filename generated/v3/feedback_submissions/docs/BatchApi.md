# \BatchApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchArchive**](BatchApi.md#BatchArchive) | **Post** /crm/v3/objects/feedback_submissions/batch/archive | Archive a batch of feedback submissions by ID
[**BatchCreate**](BatchApi.md#BatchCreate) | **Post** /crm/v3/objects/feedback_submissions/batch/create | Create a batch of feedback submissions
[**BatchRead**](BatchApi.md#BatchRead) | **Post** /crm/v3/objects/feedback_submissions/batch/read | Read a batch of feedback submissions by internal ID, or unique property values
[**BatchUpdate**](BatchApi.md#BatchUpdate) | **Post** /crm/v3/objects/feedback_submissions/batch/update | Update a batch of feedback submissions



## BatchArchive

> BatchArchive(ctx).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()

Archive a batch of feedback submissions by ID

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
    batchInputSimplePublicObjectId := *openapiclient.NewBatchInputSimplePublicObjectId([]openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}) // BatchInputSimplePublicObjectId | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchArchive(context.Background()).BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectId** | [**BatchInputSimplePublicObjectId**](BatchInputSimplePublicObjectId.md) |  | 

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


## BatchCreate

> BatchResponseSimplePublicObject BatchCreate(ctx).BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput).Execute()

Create a batch of feedback submissions

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
    batchInputSimplePublicObjectInput := *openapiclient.NewBatchInputSimplePublicObjectInput([]openapiclient.SimplePublicObjectInput{*openapiclient.NewSimplePublicObjectInput(map[string]string{"key": "Inner_example"})}) // BatchInputSimplePublicObjectInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchCreate(context.Background()).BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCreate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectInput** | [**BatchInputSimplePublicObjectInput**](BatchInputSimplePublicObjectInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRead

> BatchResponseSimplePublicObject BatchRead(ctx).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()

Read a batch of feedback submissions by internal ID, or unique property values

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
    batchReadInputSimplePublicObjectId := *openapiclient.NewBatchReadInputSimplePublicObjectId([]string{"Properties_example"}, []string{"PropertiesWithHistory_example"}, []openapiclient.SimplePublicObjectId{*openapiclient.NewSimplePublicObjectId("Id_example")}) // BatchReadInputSimplePublicObjectId | 
    archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchRead(context.Background()).BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId).Archived(archived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchRead`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchReadInputSimplePublicObjectId** | [**BatchReadInputSimplePublicObjectId**](BatchReadInputSimplePublicObjectId.md) |  | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdate

> BatchResponseSimplePublicObject BatchUpdate(ctx).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()

Update a batch of feedback submissions

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
    batchInputSimplePublicObjectBatchInput := *openapiclient.NewBatchInputSimplePublicObjectBatchInput([]openapiclient.SimplePublicObjectBatchInput{*openapiclient.NewSimplePublicObjectBatchInput(map[string]string{"key": "Inner_example"}, "Id_example")}) // BatchInputSimplePublicObjectBatchInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchUpdate(context.Background()).BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUpdate`: BatchResponseSimplePublicObject
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputSimplePublicObjectBatchInput** | [**BatchInputSimplePublicObjectBatchInput**](BatchInputSimplePublicObjectBatchInput.md) |  | 

### Return type

[**BatchResponseSimplePublicObject**](BatchResponseSimplePublicObject.md)

### Authorization

[hapikey](../README.md#hapikey), [oauth2](../README.md#oauth2), [oauth2_legacy](../README.md#oauth2_legacy), [private_apps](../README.md#private_apps), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

