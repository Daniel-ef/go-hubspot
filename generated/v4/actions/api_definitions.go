/*
Custom Workflow Actions

Create custom workflow actions

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// DefinitionsApiService DefinitionsApi service
type DefinitionsApiService service

type ApiArchiveRequest struct {
	ctx          context.Context
	ApiService   *DefinitionsApiService
	definitionId string
	appId        int32
}

func (r ApiArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.ArchiveExecute(r)
}

/*
Archive Archive a custom action

Archives a single custom workflow action with the specified ID. Workflows that currently use this custom action will stop attempting to execute the action, and all future executions will be marked as a failure.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId The ID of the custom workflow action.
	@param appId
	@return ApiArchiveRequest
*/
func (a *DefinitionsApiService) Archive(ctx context.Context, definitionId string, appId int32) ApiArchiveRequest {
	return ApiArchiveRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
func (a *DefinitionsApiService) ArchiveExecute(r ApiArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsApiService.Archive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiCreateRequest struct {
	ctx                            context.Context
	ApiService                     *DefinitionsApiService
	appId                          int32
	extensionActionDefinitionInput *ExtensionActionDefinitionInput
}

// The custom workflow action to create.
func (r ApiCreateRequest) ExtensionActionDefinitionInput(extensionActionDefinitionInput ExtensionActionDefinitionInput) ApiCreateRequest {
	r.extensionActionDefinitionInput = &extensionActionDefinitionInput
	return r
}

func (r ApiCreateRequest) Execute() (*ExtensionActionDefinition, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create new custom action

Creates a new custom workflow action.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiCreateRequest
*/
func (a *DefinitionsApiService) Create(ctx context.Context, appId int32) ApiCreateRequest {
	return ApiCreateRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return ExtensionActionDefinition
func (a *DefinitionsApiService) CreateExecute(r ApiCreateRequest) (*ExtensionActionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExtensionActionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsApiService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.extensionActionDefinitionInput == nil {
		return localVarReturnValue, nil, reportError("extensionActionDefinitionInput is required and must be specified")
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
	localVarPostBody = r.extensionActionDefinitionInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiGetByIDRequest struct {
	ctx          context.Context
	ApiService   *DefinitionsApiService
	definitionId string
	appId        int32
	archived     *bool
}

// Whether to include archived custom actions.
func (r ApiGetByIDRequest) Archived(archived bool) ApiGetByIDRequest {
	r.archived = &archived
	return r
}

func (r ApiGetByIDRequest) Execute() (*ExtensionActionDefinition, *http.Response, error) {
	return r.ApiService.GetByIDExecute(r)
}

/*
GetByID Get a custom action

Returns a single custom workflow action with the specified ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId The ID of the custom workflow action.
	@param appId
	@return ApiGetByIDRequest
*/
func (a *DefinitionsApiService) GetByID(ctx context.Context, definitionId string, appId int32) ApiGetByIDRequest {
	return ApiGetByIDRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return ExtensionActionDefinition
func (a *DefinitionsApiService) GetByIDExecute(r ApiGetByIDRequest) (*ExtensionActionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExtensionActionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsApiService.GetByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiGetPageRequest struct {
	ctx        context.Context
	ApiService *DefinitionsApiService
	appId      int32
	limit      *int32
	after      *string
	archived   *bool
}

// Maximum number of results per page.
func (r ApiGetPageRequest) Limit(limit int32) ApiGetPageRequest {
	r.limit = &limit
	return r
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetPageRequest) After(after string) ApiGetPageRequest {
	r.after = &after
	return r
}

// Whether to include archived custom actions.
func (r ApiGetPageRequest) Archived(archived bool) ApiGetPageRequest {
	r.archived = &archived
	return r
}

func (r ApiGetPageRequest) Execute() (*CollectionResponseExtensionActionDefinitionForwardPaging, *http.Response, error) {
	return r.ApiService.GetPageExecute(r)
}

/*
GetPage Get all custom actions

Returns a list of all custom workflow actions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiGetPageRequest
*/
func (a *DefinitionsApiService) GetPage(ctx context.Context, appId int32) ApiGetPageRequest {
	return ApiGetPageRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return CollectionResponseExtensionActionDefinitionForwardPaging
func (a *DefinitionsApiService) GetPageExecute(r ApiGetPageRequest) (*CollectionResponseExtensionActionDefinitionForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseExtensionActionDefinitionForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsApiService.GetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
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
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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

type ApiUpdateRequest struct {
	ctx                            context.Context
	ApiService                     *DefinitionsApiService
	definitionId                   string
	appId                          int32
	extensionActionDefinitionPatch *ExtensionActionDefinitionPatch
}

// The custom workflow action fields to be updated.
func (r ApiUpdateRequest) ExtensionActionDefinitionPatch(extensionActionDefinitionPatch ExtensionActionDefinitionPatch) ApiUpdateRequest {
	r.extensionActionDefinitionPatch = &extensionActionDefinitionPatch
	return r
}

func (r ApiUpdateRequest) Execute() (*ExtensionActionDefinition, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a custom action

Updates a custom workflow action with new values for the specified fields.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId The ID of the custom workflow action.
	@param appId
	@return ApiUpdateRequest
*/
func (a *DefinitionsApiService) Update(ctx context.Context, definitionId string, appId int32) ApiUpdateRequest {
	return ApiUpdateRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return ExtensionActionDefinition
func (a *DefinitionsApiService) UpdateExecute(r ApiUpdateRequest) (*ExtensionActionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExtensionActionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsApiService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterToString(r.definitionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.extensionActionDefinitionPatch == nil {
		return localVarReturnValue, nil, reportError("extensionActionDefinitionPatch is required and must be specified")
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
	localVarPostBody = r.extensionActionDefinitionPatch
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
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
