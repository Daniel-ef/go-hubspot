/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/Daniel-ef/go-hubspot"
	"net/url"
	"strings"
)

// AssociationsBatchApiService AssociationsBatchApi service
type AssociationsBatchApiService service

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest struct {
	ctx                                     context.Context
	ApiService                              *AssociationsBatchApiService
	fromObjectType                          string
	toObjectType                            string
	batchInputPublicAssociationMultiArchive *BatchInputPublicAssociationMultiArchive
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest) BatchInputPublicAssociationMultiArchive(batchInputPublicAssociationMultiArchive BatchInputPublicAssociationMultiArchive) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest {
	r.batchInputPublicAssociationMultiArchive = &batchInputPublicAssociationMultiArchive
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive Delete

Batch delete associations for objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest
*/
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsBatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicAssociationMultiArchive == nil {
		return nil, reportError("batchInputPublicAssociationMultiArchive is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputPublicAssociationMultiArchive
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest struct {
	ctx                                  context.Context
	ApiService                           *AssociationsBatchApiService
	fromObjectType                       string
	toObjectType                         string
	batchInputPublicAssociationMultiPost *BatchInputPublicAssociationMultiPost
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest) BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost BatchInputPublicAssociationMultiPost) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest {
	r.batchInputPublicAssociationMultiPost = &batchInputPublicAssociationMultiPost
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest) Execute() (*BatchResponseLabelsBetweenObjectPair, *http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate Create

Batch create associations for objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest
*/
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//
//	@return BatchResponseLabelsBetweenObjectPair
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreateRequest) (*BatchResponseLabelsBetweenObjectPair, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseLabelsBetweenObjectPair
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsBatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/create"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicAssociationMultiPost == nil {
		return localVarReturnValue, nil, reportError("batchInputPublicAssociationMultiPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputPublicAssociationMultiPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest struct {
	ctx                                  context.Context
	ApiService                           *AssociationsBatchApiService
	fromObjectType                       string
	toObjectType                         string
	batchInputPublicAssociationMultiPost *BatchInputPublicAssociationMultiPost
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest) BatchInputPublicAssociationMultiPost(batchInputPublicAssociationMultiPost BatchInputPublicAssociationMultiPost) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest {
	r.batchInputPublicAssociationMultiPost = &batchInputPublicAssociationMultiPost
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive Delete Specific Labels

Batch delete specific association labels for objects. Deleting an unlabeled association will also delete all labeled associations between those two objects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest
*/
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsBatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchLabelsArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/labels/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicAssociationMultiPost == nil {
		return nil, reportError("batchInputPublicAssociationMultiPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputPublicAssociationMultiPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest struct {
	ctx                                           context.Context
	ApiService                                    *AssociationsBatchApiService
	fromObjectType                                string
	toObjectType                                  string
	batchInputPublicFetchAssociationsBatchRequest *BatchInputPublicFetchAssociationsBatchRequest
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest) BatchInputPublicFetchAssociationsBatchRequest(batchInputPublicFetchAssociationsBatchRequest BatchInputPublicFetchAssociationsBatchRequest) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest {
	r.batchInputPublicFetchAssociationsBatchRequest = &batchInputPublicFetchAssociationsBatchRequest
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest) Execute() (*BatchResponsePublicAssociationMultiWithLabel, *http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead Read

Batch read associations for objects to specific object type. The 'after' field in a returned paging object  can be added alongside the 'id' to retrieve the next page of associations from that objectId. The 'link' field is deprecated and should be ignored.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest
*/
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//
//	@return BatchResponsePublicAssociationMultiWithLabel
func (a *AssociationsBatchApiService) PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeBatchReadRequest) (*BatchResponsePublicAssociationMultiWithLabel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponsePublicAssociationMultiWithLabel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsBatchApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeBatchRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/batch/read"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputPublicFetchAssociationsBatchRequest == nil {
		return localVarReturnValue, nil, reportError("batchInputPublicFetchAssociationsBatchRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputPublicFetchAssociationsBatchRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
