/*
Feedback Submissions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feedback_submissions

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/Daniel-ef/go-hubspot"
	"net/url"
	"strings"
)

// AssociationsApiService AssociationsApi service
type AssociationsApiService service

type ApiAssociationsGetPageRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId int64
	toObjectType         string
	after                *string
	limit                *int32
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiAssociationsGetPageRequest) After(after string) ApiAssociationsGetPageRequest {
	r.after = &after
	return r
}

// The maximum number of results to display per page.
func (r ApiAssociationsGetPageRequest) Limit(limit int32) ApiAssociationsGetPageRequest {
	r.limit = &limit
	return r
}

func (r ApiAssociationsGetPageRequest) Execute() (*CollectionResponseMultiAssociatedObjectWithLabelForwardPaging, *http.Response, error) {
	return r.ApiService.AssociationsGetPageExecute(r)
}

/*
AssociationsGetPage List

List all associations of a feedback submission by object type. Limit 1000 per call.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param feedbackSubmissionId
	@param toObjectType
	@return ApiAssociationsGetPageRequest
*/
func (a *AssociationsApiService) AssociationsGetPage(ctx context.Context, feedbackSubmissionId int64, toObjectType string) ApiAssociationsGetPageRequest {
	return ApiAssociationsGetPageRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
	}
}

// Execute executes the request
//
//	@return CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
func (a *AssociationsApiService) AssociationsGetPageExecute(r ApiAssociationsGetPageRequest) (*CollectionResponseMultiAssociatedObjectWithLabelForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.AssociationsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiDeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId int64
	toObjectType         string
	toObjectId           int64
}

func (r ApiDeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveExecute(r)
}

/*
DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive Delete

deletes all associations between two records.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param feedbackSubmissionId
	@param toObjectType
	@param toObjectId
	@return ApiDeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveRequest
*/
func (a *AssociationsApiService) DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive(ctx context.Context, feedbackSubmissionId int64, toObjectType string, toObjectId int64) ApiDeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveRequest {
	return ApiDeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
		toObjectId:           toObjectId,
	}
}

// Execute executes the request
func (a *AssociationsApiService) DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveExecute(r ApiDeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.DeleteCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", url.PathEscape(parameterToString(r.toObjectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest struct {
	ctx                  context.Context
	ApiService           *AssociationsApiService
	feedbackSubmissionId int64
	toObjectType         string
	toObjectId           int64
	associationSpec      *[]AssociationSpec
}

func (r ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest) AssociationSpec(associationSpec []AssociationSpec) ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest {
	r.associationSpec = &associationSpec
	return r
}

func (r ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest) Execute() (*LabelsBetweenObjectPair, *http.Response, error) {
	return r.ApiService.PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateExecute(r)
}

/*
PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate Create

Set association labels between two records.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param feedbackSubmissionId
	@param toObjectType
	@param toObjectId
	@return ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest
*/
func (a *AssociationsApiService) PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate(ctx context.Context, feedbackSubmissionId int64, toObjectType string, toObjectId int64) ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest {
	return ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest{
		ApiService:           a,
		ctx:                  ctx,
		feedbackSubmissionId: feedbackSubmissionId,
		toObjectType:         toObjectType,
		toObjectId:           toObjectId,
	}
}

// Execute executes the request
//
//	@return LabelsBetweenObjectPair
func (a *AssociationsApiService) PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateExecute(r ApiPutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreateRequest) (*LabelsBetweenObjectPair, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LabelsBetweenObjectPair
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssociationsApiService.PutCrmV4ObjectsFeedbackSubmissionsFeedbackSubmissionIdAssociationsToObjectTypeToObjectIdCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/objects/feedback_submissions/{feedbackSubmissionId}/associations/{toObjectType}/{toObjectId}"
	localVarPath = strings.Replace(localVarPath, "{"+"feedbackSubmissionId"+"}", url.PathEscape(parameterToString(r.feedbackSubmissionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterToString(r.toObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectId"+"}", url.PathEscape(parameterToString(r.toObjectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.associationSpec == nil {
		return localVarReturnValue, nil, reportError("associationSpec is required and must be specified")
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
	localVarPostBody = r.associationSpec
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
