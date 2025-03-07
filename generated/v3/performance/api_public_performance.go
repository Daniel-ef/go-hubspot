/*
CMS Performance API

Use these endpoints to get a time series view of your website's performance.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package performance

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/Daniel-ef/go-hubspot"
	"net/url"
)

// PublicPerformanceApiService PublicPerformanceApi service
type PublicPerformanceApiService service

type ApiGetPageRequest struct {
	ctx        context.Context
	ApiService *PublicPerformanceApiService
	domain     *string
	path       *string
	pad        *bool
	sum        *bool
	period     *string
	interval   *string
	start      *int64
	end        *int64
}

// The domain to search return data for.
func (r ApiGetPageRequest) Domain(domain string) ApiGetPageRequest {
	r.domain = &domain
	return r
}

// The url path of the domain to return data for.
func (r ApiGetPageRequest) Path(path string) ApiGetPageRequest {
	r.path = &path
	return r
}

// Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set.
func (r ApiGetPageRequest) Pad(pad bool) ApiGetPageRequest {
	r.pad = &pad
	return r
}

// Specifies whether the time series data should be summated for the given period. Defaults to false.
func (r ApiGetPageRequest) Sum(sum bool) ApiGetPageRequest {
	r.sum = &sum
	return r
}

// A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w]
func (r ApiGetPageRequest) Period(period string) ApiGetPageRequest {
	r.period = &period
	return r
}

// The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w]
func (r ApiGetPageRequest) Interval(interval string) ApiGetPageRequest {
	r.interval = &interval
	return r
}

// A timestamp in milliseconds that indicates the start of the time period.
func (r ApiGetPageRequest) Start(start int64) ApiGetPageRequest {
	r.start = &start
	return r
}

// A timestamp in milliseconds that indicates the end of the time period.
func (r ApiGetPageRequest) End(end int64) ApiGetPageRequest {
	r.end = &end
	return r
}

func (r ApiGetPageRequest) Execute() (*PublicPerformanceResponse, *http.Response, error) {
	return r.ApiService.GetPageExecute(r)
}

/*
GetPage View your website's performance.

Returns time series data website performance data for the given domain and/or path.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPageRequest
*/
func (a *PublicPerformanceApiService) GetPage(ctx context.Context) ApiGetPageRequest {
	return ApiGetPageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PublicPerformanceResponse
func (a *PublicPerformanceApiService) GetPageExecute(r ApiGetPageRequest) (*PublicPerformanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicPerformanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicPerformanceApiService.GetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/performance/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.domain != nil {
		localVarQueryParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.path != nil {
		localVarQueryParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.pad != nil {
		localVarQueryParams.Add("pad", parameterToString(*r.pad, ""))
	}
	if r.sum != nil {
		localVarQueryParams.Add("sum", parameterToString(*r.sum, ""))
	}
	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
	}
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("end", parameterToString(*r.end, ""))
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

type ApiGetUptimeRequest struct {
	ctx        context.Context
	ApiService *PublicPerformanceApiService
	domain     *string
	path       *string
	pad        *bool
	sum        *bool
	period     *string
	interval   *string
	start      *int64
	end        *int64
}

// The domain to search return data for.
func (r ApiGetUptimeRequest) Domain(domain string) ApiGetUptimeRequest {
	r.domain = &domain
	return r
}

func (r ApiGetUptimeRequest) Path(path string) ApiGetUptimeRequest {
	r.path = &path
	return r
}

// Specifies whether the time series data should have empty intervals if performance data is not present to create a continuous set.
func (r ApiGetUptimeRequest) Pad(pad bool) ApiGetUptimeRequest {
	r.pad = &pad
	return r
}

// Specifies whether the time series data should be summated for the given period. Defaults to false.
func (r ApiGetUptimeRequest) Sum(sum bool) ApiGetUptimeRequest {
	r.sum = &sum
	return r
}

// A relative period to return time series data for. This value is ignored if start and/or end are provided. Valid periods: [15m, 30m, 1h, 4h, 12h, 1d, 1w]
func (r ApiGetUptimeRequest) Period(period string) ApiGetUptimeRequest {
	r.period = &period
	return r
}

// The time series interval to group data by. Valid intervals: [1m, 5m, 15m, 30m, 1h, 4h, 12h, 1d, 1w]
func (r ApiGetUptimeRequest) Interval(interval string) ApiGetUptimeRequest {
	r.interval = &interval
	return r
}

// A timestamp in milliseconds that indicates the start of the time period.
func (r ApiGetUptimeRequest) Start(start int64) ApiGetUptimeRequest {
	r.start = &start
	return r
}

// A timestamp in milliseconds that indicates the end of the time period.
func (r ApiGetUptimeRequest) End(end int64) ApiGetUptimeRequest {
	r.end = &end
	return r
}

func (r ApiGetUptimeRequest) Execute() (*PublicPerformanceResponse, *http.Response, error) {
	return r.ApiService.GetUptimeExecute(r)
}

/*
GetUptime View your website's uptime.

Returns uptime time series website performance data for the given domain.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetUptimeRequest
*/
func (a *PublicPerformanceApiService) GetUptime(ctx context.Context) ApiGetUptimeRequest {
	return ApiGetUptimeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PublicPerformanceResponse
func (a *PublicPerformanceApiService) GetUptimeExecute(r ApiGetUptimeRequest) (*PublicPerformanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicPerformanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicPerformanceApiService.GetUptime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cms/v3/performance/uptime"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.domain != nil {
		localVarQueryParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.path != nil {
		localVarQueryParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.pad != nil {
		localVarQueryParams.Add("pad", parameterToString(*r.pad, ""))
	}
	if r.sum != nil {
		localVarQueryParams.Add("sum", parameterToString(*r.sum, ""))
	}
	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
	}
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("end", parameterToString(*r.end, ""))
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
