/*
Timeline events

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// TokensApiService TokensApi service
type TokensApiService service

type ApiTemplatesTokensArchiveRequest struct {
	ctx             context.Context
	ApiService      *TokensApiService
	eventTemplateId string
	tokenName       string
	appId           int32
}

func (r ApiTemplatesTokensArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.TemplatesTokensArchiveExecute(r)
}

/*
TemplatesTokensArchive Removes a token from the event template

This will remove the token from an existing template. Existing events and CRM objects will still retain the token and its mapped object properties, but new ones will not.

The timeline will still display this property for older CRM objects if it's still referenced in the template `Markdown`. New events will not.

Any lists or reports referencing deleted tokens will no longer return new contacts, but old ones will still exist in the lists.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventTemplateId The event template ID.
	@param tokenName The token name.
	@param appId The ID of the target app.
	@return ApiTemplatesTokensArchiveRequest
*/
func (a *TokensApiService) TemplatesTokensArchive(ctx context.Context, eventTemplateId string, tokenName string, appId int32) ApiTemplatesTokensArchiveRequest {
	return ApiTemplatesTokensArchiveRequest{
		ApiService:      a,
		ctx:             ctx,
		eventTemplateId: eventTemplateId,
		tokenName:       tokenName,
		appId:           appId,
	}
}

// Execute executes the request
func (a *TokensApiService) TemplatesTokensArchiveExecute(r ApiTemplatesTokensArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.TemplatesTokensArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName}"
	localVarPath = strings.Replace(localVarPath, "{"+"eventTemplateId"+"}", url.PathEscape(parameterToString(r.eventTemplateId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenName"+"}", url.PathEscape(parameterToString(r.tokenName, "")), -1)
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

type ApiTemplatesTokensCreateRequest struct {
	ctx                        context.Context
	ApiService                 *TokensApiService
	eventTemplateId            string
	appId                      int32
	timelineEventTemplateToken *TimelineEventTemplateToken
}

// The new token definition.
func (r ApiTemplatesTokensCreateRequest) TimelineEventTemplateToken(timelineEventTemplateToken TimelineEventTemplateToken) ApiTemplatesTokensCreateRequest {
	r.timelineEventTemplateToken = &timelineEventTemplateToken
	return r
}

func (r ApiTemplatesTokensCreateRequest) Execute() (*TimelineEventTemplateToken, *http.Response, error) {
	return r.ApiService.TemplatesTokensCreateExecute(r)
}

/*
TemplatesTokensCreate Adds a token to an existing event template

Once you've defined an event template, it's likely that you'll want to define tokens for it as well. You can do this on the event template itself or update individual tokens here.

Event type tokens allow you to attach custom data to events displayed in a timeline or used for list segmentation.

You can also use `objectPropertyName` to associate any CRM object properties. This will allow you to fully build out CRM objects.

Token names should be unique across the template.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventTemplateId The event template ID.
	@param appId The ID of the target app.
	@return ApiTemplatesTokensCreateRequest
*/
func (a *TokensApiService) TemplatesTokensCreate(ctx context.Context, eventTemplateId string, appId int32) ApiTemplatesTokensCreateRequest {
	return ApiTemplatesTokensCreateRequest{
		ApiService:      a,
		ctx:             ctx,
		eventTemplateId: eventTemplateId,
		appId:           appId,
	}
}

// Execute executes the request
//
//	@return TimelineEventTemplateToken
func (a *TokensApiService) TemplatesTokensCreateExecute(r ApiTemplatesTokensCreateRequest) (*TimelineEventTemplateToken, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TimelineEventTemplateToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.TemplatesTokensCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"eventTemplateId"+"}", url.PathEscape(parameterToString(r.eventTemplateId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timelineEventTemplateToken == nil {
		return localVarReturnValue, nil, reportError("timelineEventTemplateToken is required and must be specified")
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
	localVarPostBody = r.timelineEventTemplateToken
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

type ApiTemplatesTokensUpdateRequest struct {
	ctx                                     context.Context
	ApiService                              *TokensApiService
	eventTemplateId                         string
	tokenName                               string
	appId                                   int32
	timelineEventTemplateTokenUpdateRequest *TimelineEventTemplateTokenUpdateRequest
}

// The updated token definition.
func (r ApiTemplatesTokensUpdateRequest) TimelineEventTemplateTokenUpdateRequest(timelineEventTemplateTokenUpdateRequest TimelineEventTemplateTokenUpdateRequest) ApiTemplatesTokensUpdateRequest {
	r.timelineEventTemplateTokenUpdateRequest = &timelineEventTemplateTokenUpdateRequest
	return r
}

func (r ApiTemplatesTokensUpdateRequest) Execute() (*TimelineEventTemplateToken, *http.Response, error) {
	return r.ApiService.TemplatesTokensUpdateExecute(r)
}

/*
TemplatesTokensUpdate Updates an existing token on an event template

This will update the existing token on an event template. Name and type can't be changed on existing tokens.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventTemplateId The event template ID.
	@param tokenName The token name.
	@param appId The ID of the target app.
	@return ApiTemplatesTokensUpdateRequest
*/
func (a *TokensApiService) TemplatesTokensUpdate(ctx context.Context, eventTemplateId string, tokenName string, appId int32) ApiTemplatesTokensUpdateRequest {
	return ApiTemplatesTokensUpdateRequest{
		ApiService:      a,
		ctx:             ctx,
		eventTemplateId: eventTemplateId,
		tokenName:       tokenName,
		appId:           appId,
	}
}

// Execute executes the request
//
//	@return TimelineEventTemplateToken
func (a *TokensApiService) TemplatesTokensUpdateExecute(r ApiTemplatesTokensUpdateRequest) (*TimelineEventTemplateToken, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TimelineEventTemplateToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokensApiService.TemplatesTokensUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName}"
	localVarPath = strings.Replace(localVarPath, "{"+"eventTemplateId"+"}", url.PathEscape(parameterToString(r.eventTemplateId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenName"+"}", url.PathEscape(parameterToString(r.tokenName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timelineEventTemplateTokenUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("timelineEventTemplateTokenUpdateRequest is required and must be specified")
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
	localVarPostBody = r.timelineEventTemplateTokenUpdateRequest
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
