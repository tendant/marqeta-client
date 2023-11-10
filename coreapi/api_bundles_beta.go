/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marqeta_coreapi_client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// BundlesBetaAPIService BundlesBetaAPI service
type BundlesBetaAPIService service

type ApiCloneBundleRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	token string
}

func (r ApiCloneBundleRequest) Execute() (*BundleResponse, *http.Response, error) {
	return r.ApiService.CloneBundleExecute(r)
}

/*
CloneBundle Clone bundle

Create a new bundle based on an existing bundle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token The unique identifier of the bundle to clone.  Send a `GET` request to `/bundles` to retrieve existing bundle tokens.
 @return ApiCloneBundleRequest
*/
func (a *BundlesBetaAPIService) CloneBundle(ctx context.Context, token string) ApiCloneBundleRequest {
	return ApiCloneBundleRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BundleResponse
func (a *BundlesBetaAPIService) CloneBundleExecute(r ApiCloneBundleRequest) (*BundleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.CloneBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles/{token}/clone"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiCreateBundleRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	bundleCreateReq *BundleCreateReq
}

func (r ApiCreateBundleRequest) BundleCreateReq(bundleCreateReq BundleCreateReq) ApiCreateBundleRequest {
	r.bundleCreateReq = &bundleCreateReq
	return r
}

func (r ApiCreateBundleRequest) Execute() (*BundleResponse, *http.Response, error) {
	return r.ApiService.CreateBundleExecute(r)
}

/*
CreateBundle Create bundle

Create a bundle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateBundleRequest
*/
func (a *BundlesBetaAPIService) CreateBundle(ctx context.Context) ApiCreateBundleRequest {
	return ApiCreateBundleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BundleResponse
func (a *BundlesBetaAPIService) CreateBundleExecute(r ApiCreateBundleRequest) (*BundleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.CreateBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bundleCreateReq == nil {
		return localVarReturnValue, nil, reportError("bundleCreateReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bundleCreateReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListBundlesRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	count *int32
	startIndex *int32
	sortBy *string
	status *[]BundleResourceStatus
}

// The number of resources to retrieve.
func (r ApiListBundlesRequest) Count(count int32) ApiListBundlesRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ApiListBundlesRequest) StartIndex(startIndex int32) ApiListBundlesRequest {
	r.startIndex = &startIndex
	return r
}

// Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;.
func (r ApiListBundlesRequest) SortBy(sortBy string) ApiListBundlesRequest {
	r.sortBy = &sortBy
	return r
}

// An array of statuses by which to filter bundles.
func (r ApiListBundlesRequest) Status(status []BundleResourceStatus) ApiListBundlesRequest {
	r.status = &status
	return r
}

func (r ApiListBundlesRequest) Execute() (*BundleResponsePage, *http.Response, error) {
	return r.ApiService.ListBundlesExecute(r)
}

/*
ListBundles List bundles

Retrieve an array of bundles.

This endpoint supports <</core-api/sorting-and-pagination, sorting and pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListBundlesRequest
*/
func (a *BundlesBetaAPIService) ListBundles(ctx context.Context) ApiListBundlesRequest {
	return ApiListBundlesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BundleResponsePage
func (a *BundlesBetaAPIService) ListBundlesExecute(r ApiListBundlesRequest) (*BundleResponsePage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleResponsePage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.ListBundles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 5
		r.count = &defaultValue
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_index", r.startIndex, "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-lastModifiedTime"
		r.sortBy = &defaultValue
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "status", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "status", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListRelatedBundlesRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	token string
	count *int32
	startIndex *int32
	sortBy *string
	status *[]BundleResourceStatus
}

// The number of resources to retrieve.
func (r ApiListRelatedBundlesRequest) Count(count int32) ApiListRelatedBundlesRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ApiListRelatedBundlesRequest) StartIndex(startIndex int32) ApiListRelatedBundlesRequest {
	r.startIndex = &startIndex
	return r
}

// Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;.
func (r ApiListRelatedBundlesRequest) SortBy(sortBy string) ApiListRelatedBundlesRequest {
	r.sortBy = &sortBy
	return r
}

// An array of statuses by which to filter bundles.
func (r ApiListRelatedBundlesRequest) Status(status []BundleResourceStatus) ApiListRelatedBundlesRequest {
	r.status = &status
	return r
}

func (r ApiListRelatedBundlesRequest) Execute() (*BundleResponsePage, *http.Response, error) {
	return r.ApiService.ListRelatedBundlesExecute(r)
}

/*
ListRelatedBundles List related bundles

Retrieve an array of related product bundles.

This endpoint supports <</core-api/sorting-and-pagination, sorting and pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token The unique identifier of the parent product bundle.  Send a `GET` request to `/bundles` to retrieve existing product bundles tokens.
 @return ApiListRelatedBundlesRequest
*/
func (a *BundlesBetaAPIService) ListRelatedBundles(ctx context.Context, token string) ApiListRelatedBundlesRequest {
	return ApiListRelatedBundlesRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BundleResponsePage
func (a *BundlesBetaAPIService) ListRelatedBundlesExecute(r ApiListRelatedBundlesRequest) (*BundleResponsePage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleResponsePage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.ListRelatedBundles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles/{token}/lineage"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 5
		r.count = &defaultValue
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_index", r.startIndex, "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-lastModifiedTime"
		r.sortBy = &defaultValue
	}
	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "status", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "status", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPromoteBundleRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	token string
}

func (r ApiPromoteBundleRequest) Execute() (*BundleResponse, *http.Response, error) {
	return r.ApiService.PromoteBundleExecute(r)
}

/*
PromoteBundle Promote bundle

Promote a specific bundle, which replaces the current active bundle and activates the promoted bundle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token The unique identifier of the bundle to promote.  Send a `GET` request to `/bundles` to retrieve existing bundle tokens.
 @return ApiPromoteBundleRequest
*/
func (a *BundlesBetaAPIService) PromoteBundle(ctx context.Context, token string) ApiPromoteBundleRequest {
	return ApiPromoteBundleRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BundleResponse
func (a *BundlesBetaAPIService) PromoteBundleExecute(r ApiPromoteBundleRequest) (*BundleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.PromoteBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles/{token}/promote"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiRetrieveBundleRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	token string
	expandObjects *[]PolicyType
}

// Embeds the associated object of the specified type into the response. For more, see &lt;&lt;/core-api/object-expansion, object expansion&gt;&gt;.
func (r ApiRetrieveBundleRequest) ExpandObjects(expandObjects []PolicyType) ApiRetrieveBundleRequest {
	r.expandObjects = &expandObjects
	return r
}

func (r ApiRetrieveBundleRequest) Execute() (*BundleResponse, *http.Response, error) {
	return r.ApiService.RetrieveBundleExecute(r)
}

/*
RetrieveBundle Retrieve bundle

Retrieve a specific bundle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token The unique identifier of the bundle to retrieve.  Send a `GET` request to `/credit/bundles` to retrieve existing  tokens.
 @return ApiRetrieveBundleRequest
*/
func (a *BundlesBetaAPIService) RetrieveBundle(ctx context.Context, token string) ApiRetrieveBundleRequest {
	return ApiRetrieveBundleRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BundleResponse
func (a *BundlesBetaAPIService) RetrieveBundleExecute(r ApiRetrieveBundleRequest) (*BundleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.RetrieveBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expandObjects != nil {
		t := *r.expandObjects
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "expand_objects", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "expand_objects", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiTransitionBundleRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	token string
	bundleTransitionReq *BundleTransitionReq
}

func (r ApiTransitionBundleRequest) BundleTransitionReq(bundleTransitionReq BundleTransitionReq) ApiTransitionBundleRequest {
	r.bundleTransitionReq = &bundleTransitionReq
	return r
}

func (r ApiTransitionBundleRequest) Execute() (*BundleTransitionResponse, *http.Response, error) {
	return r.ApiService.TransitionBundleExecute(r)
}

/*
TransitionBundle Transition a bundle

Transition the status of a specific bundle.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Token of the bundle whose status you want to transition.  Send a `GET` request to `/credit/bundles` to retrieve existing  tokens.
 @return ApiTransitionBundleRequest
*/
func (a *BundlesBetaAPIService) TransitionBundle(ctx context.Context, token string) ApiTransitionBundleRequest {
	return ApiTransitionBundleRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BundleTransitionResponse
func (a *BundlesBetaAPIService) TransitionBundleExecute(r ApiTransitionBundleRequest) (*BundleTransitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleTransitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.TransitionBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles/{token}/transitions"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bundleTransitionReq == nil {
		return localVarReturnValue, nil, reportError("bundleTransitionReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bundleTransitionReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateBundleRequest struct {
	ctx context.Context
	ApiService *BundlesBetaAPIService
	token string
	bundleUpdateReq *BundleUpdateReq
}

func (r ApiUpdateBundleRequest) BundleUpdateReq(bundleUpdateReq BundleUpdateReq) ApiUpdateBundleRequest {
	r.bundleUpdateReq = &bundleUpdateReq
	return r
}

func (r ApiUpdateBundleRequest) Execute() (*BundleResponse, *http.Response, error) {
	return r.ApiService.UpdateBundleExecute(r)
}

/*
UpdateBundle Update bundle

Update a specific bundle that is not `ACTIVE` or `ARCHIVED`. Bundles are created in a `DRAFT` state, and are still modifiable at this point. Using the transitions endpoint a bundle can be transitioned from `DRAFT`, to `ACTIVE`. Once a bundle is active, it is immutable, and cannot be modified.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Token of the bundle to update.  Send a `GET` request to `/credit/bundles` to retrieve existing  tokens.
 @return ApiUpdateBundleRequest
*/
func (a *BundlesBetaAPIService) UpdateBundle(ctx context.Context, token string) ApiUpdateBundleRequest {
	return ApiUpdateBundleRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BundleResponse
func (a *BundlesBetaAPIService) UpdateBundleExecute(r ApiUpdateBundleRequest) (*BundleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BundleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesBetaAPIService.UpdateBundle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/bundles/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bundleUpdateReq == nil {
		return localVarReturnValue, nil, reportError("bundleUpdateReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bundleUpdateReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
