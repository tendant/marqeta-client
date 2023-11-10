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
	"time"
	"reflect"
)


// RewardRedemptionsBetaAPIService RewardRedemptionsBetaAPI service
type RewardRedemptionsBetaAPIService service

type ApiGetRedemptionRequest struct {
	ctx context.Context
	ApiService *RewardRedemptionsBetaAPIService
	token string
	redemptionToken string
}

func (r ApiGetRedemptionRequest) Execute() (*RedemptionsResponse, *http.Response, error) {
	return r.ApiService.GetRedemptionExecute(r)
}

/*
GetRedemption Retrieve reward redemption

Retrieve a specific redemption for a specific reward program.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the reward program.
 @param redemptionToken Unique identifier of the reward redemption.
 @return ApiGetRedemptionRequest
*/
func (a *RewardRedemptionsBetaAPIService) GetRedemption(ctx context.Context, token string, redemptionToken string) ApiGetRedemptionRequest {
	return ApiGetRedemptionRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
		redemptionToken: redemptionToken,
	}
}

// Execute executes the request
//  @return RedemptionsResponse
func (a *RewardRedemptionsBetaAPIService) GetRedemptionExecute(r ApiGetRedemptionRequest) (*RedemptionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RedemptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RewardRedemptionsBetaAPIService.GetRedemption")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/rewardprograms/{token}/redemptions/{redemption_token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"redemption_token"+"}", url.PathEscape(parameterValueToString(r.redemptionToken, "redemptionToken")), -1)

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

type ApiPostRedemptionTransitionRequest struct {
	ctx context.Context
	ApiService *RewardRedemptionsBetaAPIService
	token string
	redemptionToken string
	createRedemptionTransitionsRequest *CreateRedemptionTransitionsRequest
}

func (r ApiPostRedemptionTransitionRequest) CreateRedemptionTransitionsRequest(createRedemptionTransitionsRequest CreateRedemptionTransitionsRequest) ApiPostRedemptionTransitionRequest {
	r.createRedemptionTransitionsRequest = &createRedemptionTransitionsRequest
	return r
}

func (r ApiPostRedemptionTransitionRequest) Execute() (*RedemptionTransitionsResponse, *http.Response, error) {
	return r.ApiService.PostRedemptionTransitionExecute(r)
}

/*
PostRedemptionTransition Transition reward redemption status

Transition the current status of a reward redemption to a new status.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the reward program.
 @param redemptionToken Unique identifier of the reward redemption.
 @return ApiPostRedemptionTransitionRequest
*/
func (a *RewardRedemptionsBetaAPIService) PostRedemptionTransition(ctx context.Context, token string, redemptionToken string) ApiPostRedemptionTransitionRequest {
	return ApiPostRedemptionTransitionRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
		redemptionToken: redemptionToken,
	}
}

// Execute executes the request
//  @return RedemptionTransitionsResponse
func (a *RewardRedemptionsBetaAPIService) PostRedemptionTransitionExecute(r ApiPostRedemptionTransitionRequest) (*RedemptionTransitionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RedemptionTransitionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RewardRedemptionsBetaAPIService.PostRedemptionTransition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/rewardprograms/{token}/redemptions/{redemption_token}/transitions"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"redemption_token"+"}", url.PathEscape(parameterValueToString(r.redemptionToken, "redemptionToken")), -1)

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
	localVarPostBody = r.createRedemptionTransitionsRequest
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

type ApiPostRewardProgramRedemptionRequest struct {
	ctx context.Context
	ApiService *RewardRedemptionsBetaAPIService
	token string
	createRedemptionsRequest *CreateRedemptionsRequest
}

func (r ApiPostRewardProgramRedemptionRequest) CreateRedemptionsRequest(createRedemptionsRequest CreateRedemptionsRequest) ApiPostRewardProgramRedemptionRequest {
	r.createRedemptionsRequest = &createRedemptionsRequest
	return r
}

func (r ApiPostRewardProgramRedemptionRequest) Execute() (*RedemptionsResponse, *http.Response, error) {
	return r.ApiService.PostRewardProgramRedemptionExecute(r)
}

/*
PostRewardProgramRedemption Create reward redemption

Create a redemption to redeem rewards on a specific reward program.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the reward program.
 @return ApiPostRewardProgramRedemptionRequest
*/
func (a *RewardRedemptionsBetaAPIService) PostRewardProgramRedemption(ctx context.Context, token string) ApiPostRewardProgramRedemptionRequest {
	return ApiPostRewardProgramRedemptionRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return RedemptionsResponse
func (a *RewardRedemptionsBetaAPIService) PostRewardProgramRedemptionExecute(r ApiPostRewardProgramRedemptionRequest) (*RedemptionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RedemptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RewardRedemptionsBetaAPIService.PostRewardProgramRedemption")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/rewardprograms/{token}/redemptions"
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
	localVarPostBody = r.createRedemptionsRequest
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

type ApiRetrieveRedemptionsRequest struct {
	ctx context.Context
	ApiService *RewardRedemptionsBetaAPIService
	token string
	startDate *time.Time
	endDate *time.Time
	count *int32
	startIndex *int64
	sortBy *string
	type_ *RedemptionType
}

// The starting date (or date-time) of a date range from which to return resources, in UTC.
func (r ApiRetrieveRedemptionsRequest) StartDate(startDate time.Time) ApiRetrieveRedemptionsRequest {
	r.startDate = &startDate
	return r
}

// The ending date (or date-time) of a date range from which to return resources, in UTC.
func (r ApiRetrieveRedemptionsRequest) EndDate(endDate time.Time) ApiRetrieveRedemptionsRequest {
	r.endDate = &endDate
	return r
}

// Number of resources to retrieve.
func (r ApiRetrieveRedemptionsRequest) Count(count int32) ApiRetrieveRedemptionsRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ApiRetrieveRedemptionsRequest) StartIndex(startIndex int64) ApiRetrieveRedemptionsRequest {
	r.startIndex = &startIndex
	return r
}

// Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;.
func (r ApiRetrieveRedemptionsRequest) SortBy(sortBy string) ApiRetrieveRedemptionsRequest {
	r.sortBy = &sortBy
	return r
}

// Type of reward redemptions in the returned array.
func (r ApiRetrieveRedemptionsRequest) Type_(type_ RedemptionType) ApiRetrieveRedemptionsRequest {
	r.type_ = &type_
	return r
}

func (r ApiRetrieveRedemptionsRequest) Execute() (*RedemptionsPage, *http.Response, error) {
	return r.ApiService.RetrieveRedemptionsExecute(r)
}

/*
RetrieveRedemptions List reward redemptions

Retrieve an array of reward redemptions for a specific reward program.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the reward program.
 @return ApiRetrieveRedemptionsRequest
*/
func (a *RewardRedemptionsBetaAPIService) RetrieveRedemptions(ctx context.Context, token string) ApiRetrieveRedemptionsRequest {
	return ApiRetrieveRedemptionsRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return RedemptionsPage
func (a *RewardRedemptionsBetaAPIService) RetrieveRedemptionsExecute(r ApiRetrieveRedemptionsRequest) (*RedemptionsPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RedemptionsPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RewardRedemptionsBetaAPIService.RetrieveRedemptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/rewardprograms/{token}/redemptions"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 5
		r.count = &defaultValue
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_index", r.startIndex, "")
	} else {
		var defaultValue int64 = 0
		r.startIndex = &defaultValue
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-createdTime"
		r.sortBy = &defaultValue
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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

type ApiRetrieveRedemptionsBalanceRequest struct {
	ctx context.Context
	ApiService *RewardRedemptionsBetaAPIService
	token string
	startDate *time.Time
	endDate *time.Time
	type_ *[]RedemptionType
}

// The starting date (or date-time) of a date range from which to return resources, in UTC.
func (r ApiRetrieveRedemptionsBalanceRequest) StartDate(startDate time.Time) ApiRetrieveRedemptionsBalanceRequest {
	r.startDate = &startDate
	return r
}

// The ending date (or date-time) of a date range from which to return resources, in UTC.
func (r ApiRetrieveRedemptionsBalanceRequest) EndDate(endDate time.Time) ApiRetrieveRedemptionsBalanceRequest {
	r.endDate = &endDate
	return r
}

// Type of reward redemptions in the returned array.
func (r ApiRetrieveRedemptionsBalanceRequest) Type_(type_ []RedemptionType) ApiRetrieveRedemptionsBalanceRequest {
	r.type_ = &type_
	return r
}

func (r ApiRetrieveRedemptionsBalanceRequest) Execute() (*RedemptionsBalanceResponse, *http.Response, error) {
	return r.ApiService.RetrieveRedemptionsBalanceExecute(r)
}

/*
RetrieveRedemptionsBalance Retrieve reward redemption balance

Retrieve the balance for reward redemptions within a specified date range.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the reward program.
 @return ApiRetrieveRedemptionsBalanceRequest
*/
func (a *RewardRedemptionsBetaAPIService) RetrieveRedemptionsBalance(ctx context.Context, token string) ApiRetrieveRedemptionsBalanceRequest {
	return ApiRetrieveRedemptionsBalanceRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return RedemptionsBalanceResponse
func (a *RewardRedemptionsBetaAPIService) RetrieveRedemptionsBalanceExecute(r ApiRetrieveRedemptionsBalanceRequest) (*RedemptionsBalanceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RedemptionsBalanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RewardRedemptionsBetaAPIService.RetrieveRedemptionsBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit/rewardprograms/{token}/redemptions/balance"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	if r.type_ != nil {
		t := *r.type_
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "type", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "type", t, "multi")
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