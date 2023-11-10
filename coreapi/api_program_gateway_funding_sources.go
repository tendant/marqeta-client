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
)


// ProgramGatewayFundingSourcesAPIService ProgramGatewayFundingSourcesAPI service
type ProgramGatewayFundingSourcesAPIService service

type ApiGetFundingsourcesProgramgatewayTokenRequest struct {
	ctx context.Context
	ApiService *ProgramGatewayFundingSourcesAPIService
	token string
}

func (r ApiGetFundingsourcesProgramgatewayTokenRequest) Execute() (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	return r.ApiService.GetFundingsourcesProgramgatewayTokenExecute(r)
}

/*
GetFundingsourcesProgramgatewayToken Retrieve program gateway source

Retrieves a specific program gateway funding source.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the program gateway funding source.
 @return ApiGetFundingsourcesProgramgatewayTokenRequest
*/
func (a *ProgramGatewayFundingSourcesAPIService) GetFundingsourcesProgramgatewayToken(ctx context.Context, token string) ApiGetFundingsourcesProgramgatewayTokenRequest {
	return ApiGetFundingsourcesProgramgatewayTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return GatewayProgramFundingSourceResponse
func (a *ProgramGatewayFundingSourcesAPIService) GetFundingsourcesProgramgatewayTokenExecute(r ApiGetFundingsourcesProgramgatewayTokenRequest) (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GatewayProgramFundingSourceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgramGatewayFundingSourcesAPIService.GetFundingsourcesProgramgatewayToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fundingsources/programgateway/{token}"
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

type ApiPostFundingsourcesProgramgatewayRequest struct {
	ctx context.Context
	ApiService *ProgramGatewayFundingSourcesAPIService
	gatewayProgramFundingSourceRequest *GatewayProgramFundingSourceRequest
}

func (r ApiPostFundingsourcesProgramgatewayRequest) GatewayProgramFundingSourceRequest(gatewayProgramFundingSourceRequest GatewayProgramFundingSourceRequest) ApiPostFundingsourcesProgramgatewayRequest {
	r.gatewayProgramFundingSourceRequest = &gatewayProgramFundingSourceRequest
	return r
}

func (r ApiPostFundingsourcesProgramgatewayRequest) Execute() (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	return r.ApiService.PostFundingsourcesProgramgatewayExecute(r)
}

/*
PostFundingsourcesProgramgateway Create program gateway source

Creates a program gateway funding source.
A program gateway funding source is a transaction relay that allows you to approve or decline transactions in real time.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostFundingsourcesProgramgatewayRequest
*/
func (a *ProgramGatewayFundingSourcesAPIService) PostFundingsourcesProgramgateway(ctx context.Context) ApiPostFundingsourcesProgramgatewayRequest {
	return ApiPostFundingsourcesProgramgatewayRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GatewayProgramFundingSourceResponse
func (a *ProgramGatewayFundingSourcesAPIService) PostFundingsourcesProgramgatewayExecute(r ApiPostFundingsourcesProgramgatewayRequest) (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GatewayProgramFundingSourceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgramGatewayFundingSourcesAPIService.PostFundingsourcesProgramgateway")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fundingsources/programgateway"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.gatewayProgramFundingSourceRequest
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

type ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest struct {
	ctx context.Context
	ApiService *ProgramGatewayFundingSourcesAPIService
	token string
	gatewayProgramCustomHeaderUpdateRequest *GatewayProgramCustomHeaderUpdateRequest
}

func (r ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest) GatewayProgramCustomHeaderUpdateRequest(gatewayProgramCustomHeaderUpdateRequest GatewayProgramCustomHeaderUpdateRequest) ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest {
	r.gatewayProgramCustomHeaderUpdateRequest = &gatewayProgramCustomHeaderUpdateRequest
	return r
}

func (r ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest) Execute() (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	return r.ApiService.PutFundingsourcesProgramgatewayCustomHeaderTokenExecute(r)
}

/*
PutFundingsourcesProgramgatewayCustomHeaderToken Update program gateway source custom headers

Adds or updates custom HTTP headers for a specific program gateway funding source.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the program gateway funding source.
 @return ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest
*/
func (a *ProgramGatewayFundingSourcesAPIService) PutFundingsourcesProgramgatewayCustomHeaderToken(ctx context.Context, token string) ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest {
	return ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return GatewayProgramFundingSourceResponse
func (a *ProgramGatewayFundingSourcesAPIService) PutFundingsourcesProgramgatewayCustomHeaderTokenExecute(r ApiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest) (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GatewayProgramFundingSourceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgramGatewayFundingSourcesAPIService.PutFundingsourcesProgramgatewayCustomHeaderToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fundingsources/programgateway/customheaders/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.gatewayProgramCustomHeaderUpdateRequest
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

type ApiPutFundingsourcesProgramgatewayTokenRequest struct {
	ctx context.Context
	ApiService *ProgramGatewayFundingSourcesAPIService
	token string
	gatewayProgramFundingSourceUpdateRequest *GatewayProgramFundingSourceUpdateRequest
}

func (r ApiPutFundingsourcesProgramgatewayTokenRequest) GatewayProgramFundingSourceUpdateRequest(gatewayProgramFundingSourceUpdateRequest GatewayProgramFundingSourceUpdateRequest) ApiPutFundingsourcesProgramgatewayTokenRequest {
	r.gatewayProgramFundingSourceUpdateRequest = &gatewayProgramFundingSourceUpdateRequest
	return r
}

func (r ApiPutFundingsourcesProgramgatewayTokenRequest) Execute() (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	return r.ApiService.PutFundingsourcesProgramgatewayTokenExecute(r)
}

/*
PutFundingsourcesProgramgatewayToken Update program gateway source

Update a program gateway funding source.
Only the values of parameters specified in the request are modified; all others are left unchanged.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the program gateway funding source.
 @return ApiPutFundingsourcesProgramgatewayTokenRequest
*/
func (a *ProgramGatewayFundingSourcesAPIService) PutFundingsourcesProgramgatewayToken(ctx context.Context, token string) ApiPutFundingsourcesProgramgatewayTokenRequest {
	return ApiPutFundingsourcesProgramgatewayTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return GatewayProgramFundingSourceResponse
func (a *ProgramGatewayFundingSourcesAPIService) PutFundingsourcesProgramgatewayTokenExecute(r ApiPutFundingsourcesProgramgatewayTokenRequest) (*GatewayProgramFundingSourceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GatewayProgramFundingSourceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProgramGatewayFundingSourcesAPIService.PutFundingsourcesProgramgatewayToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fundingsources/programgateway/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.gatewayProgramFundingSourceUpdateRequest
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
