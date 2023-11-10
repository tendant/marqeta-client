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


// CreditDisputesAPIService CreditDisputesAPI service
type CreditDisputesAPIService service

type ApiCreateDisputeForAccountRequest struct {
	ctx context.Context
	ApiService *CreditDisputesAPIService
	accountToken string
	disputeCreateReq *DisputeCreateReq
}

func (r ApiCreateDisputeForAccountRequest) DisputeCreateReq(disputeCreateReq DisputeCreateReq) ApiCreateDisputeForAccountRequest {
	r.disputeCreateReq = &disputeCreateReq
	return r
}

func (r ApiCreateDisputeForAccountRequest) Execute() (*DisputeResponse, *http.Response, error) {
	return r.ApiService.CreateDisputeForAccountExecute(r)
}

/*
CreateDisputeForAccount Create account dispute

Create a dispute of a journal entry on a credit account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken The unique identifier of the credit account for which to create a dispute.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @return ApiCreateDisputeForAccountRequest
*/
func (a *CreditDisputesAPIService) CreateDisputeForAccount(ctx context.Context, accountToken string) ApiCreateDisputeForAccountRequest {
	return ApiCreateDisputeForAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
	}
}

// Execute executes the request
//  @return DisputeResponse
func (a *CreditDisputesAPIService) CreateDisputeForAccountExecute(r ApiCreateDisputeForAccountRequest) (*DisputeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DisputeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditDisputesAPIService.CreateDisputeForAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/accounts/{account_token}/disputes"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.disputeCreateReq == nil {
		return localVarReturnValue, nil, reportError("disputeCreateReq is required and must be specified")
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
	localVarPostBody = r.disputeCreateReq
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

type ApiGetDisputesByAccountRequest struct {
	ctx context.Context
	ApiService *CreditDisputesAPIService
	accountToken string
	count *int32
	startIndex *int32
	sortBy *string
}

// The number of resources to retrieve.
func (r ApiGetDisputesByAccountRequest) Count(count int32) ApiGetDisputesByAccountRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ApiGetDisputesByAccountRequest) StartIndex(startIndex int32) ApiGetDisputesByAccountRequest {
	r.startIndex = &startIndex
	return r
}

// Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;.
func (r ApiGetDisputesByAccountRequest) SortBy(sortBy string) ApiGetDisputesByAccountRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiGetDisputesByAccountRequest) Execute() (*DisputeResponsePage, *http.Response, error) {
	return r.ApiService.GetDisputesByAccountExecute(r)
}

/*
GetDisputesByAccount List account disputes

Retrieve an array of disputes on a credit account.

This endpoint supports <</core-api/sorting-and-pagination, sorting and pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken The unique identifier of the credit account for which to retrieve the disputes.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @return ApiGetDisputesByAccountRequest
*/
func (a *CreditDisputesAPIService) GetDisputesByAccount(ctx context.Context, accountToken string) ApiGetDisputesByAccountRequest {
	return ApiGetDisputesByAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
	}
}

// Execute executes the request
//  @return DisputeResponsePage
func (a *CreditDisputesAPIService) GetDisputesByAccountExecute(r ApiGetDisputesByAccountRequest) (*DisputeResponsePage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DisputeResponsePage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditDisputesAPIService.GetDisputesByAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/accounts/{account_token}/disputes"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)

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

type ApiRetrieveDisputeRequest struct {
	ctx context.Context
	ApiService *CreditDisputesAPIService
	accountToken string
	disputeToken string
}

func (r ApiRetrieveDisputeRequest) Execute() (*DisputeResponse, *http.Response, error) {
	return r.ApiService.RetrieveDisputeExecute(r)
}

/*
RetrieveDispute Retrieve account dispute

Retrieve a dispute from a credit account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken The unique identifier of the credit account from which to retrieve a dispute.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @param disputeToken The unique identifier of the dispute to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/disputes` to retrieve existing dispute tokens.
 @return ApiRetrieveDisputeRequest
*/
func (a *CreditDisputesAPIService) RetrieveDispute(ctx context.Context, accountToken string, disputeToken string) ApiRetrieveDisputeRequest {
	return ApiRetrieveDisputeRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
		disputeToken: disputeToken,
	}
}

// Execute executes the request
//  @return DisputeResponse
func (a *CreditDisputesAPIService) RetrieveDisputeExecute(r ApiRetrieveDisputeRequest) (*DisputeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DisputeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditDisputesAPIService.RetrieveDispute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/accounts/{account_token}/disputes/{dispute_token}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dispute_token"+"}", url.PathEscape(parameterValueToString(r.disputeToken, "disputeToken")), -1)

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

type ApiTransitionDisputeRequest struct {
	ctx context.Context
	ApiService *CreditDisputesAPIService
	accountToken string
	disputeToken string
	disputeTransitionReq *DisputeTransitionReq
}

func (r ApiTransitionDisputeRequest) DisputeTransitionReq(disputeTransitionReq DisputeTransitionReq) ApiTransitionDisputeRequest {
	r.disputeTransitionReq = &disputeTransitionReq
	return r
}

func (r ApiTransitionDisputeRequest) Execute() (*DisputeTransitionResponse, *http.Response, error) {
	return r.ApiService.TransitionDisputeExecute(r)
}

/*
TransitionDispute Update account dispute

Update the amount or status of a dispute on a credit account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken The unique identifier of the credit account from which to update a dispute.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @param disputeToken The unique identifier of the dispute to update.  Send a `GET` request to `/credit/accounts/{account_token}/disputes` to retrieve existing credit account tokens.
 @return ApiTransitionDisputeRequest
*/
func (a *CreditDisputesAPIService) TransitionDispute(ctx context.Context, accountToken string, disputeToken string) ApiTransitionDisputeRequest {
	return ApiTransitionDisputeRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
		disputeToken: disputeToken,
	}
}

// Execute executes the request
//  @return DisputeTransitionResponse
func (a *CreditDisputesAPIService) TransitionDisputeExecute(r ApiTransitionDisputeRequest) (*DisputeTransitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DisputeTransitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditDisputesAPIService.TransitionDispute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/accounts/{account_token}/disputes/{dispute_token}/transitions"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dispute_token"+"}", url.PathEscape(parameterValueToString(r.disputeToken, "disputeToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.disputeTransitionReq == nil {
		return localVarReturnValue, nil, reportError("disputeTransitionReq is required and must be specified")
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
	localVarPostBody = r.disputeTransitionReq
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
