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


// BusinessTransitionsAPIService BusinessTransitionsAPI service
type BusinessTransitionsAPIService service

type ApiGetBusinesstransitionsBusinessBusinesstokenRequest struct {
	ctx context.Context
	ApiService *BusinessTransitionsAPIService
	businessToken string
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
}

// Number of business transitions to retrieve.
func (r ApiGetBusinesstransitionsBusinessBusinesstokenRequest) Count(count int32) ApiGetBusinesstransitionsBusinessBusinesstokenRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ApiGetBusinesstransitionsBusinessBusinesstokenRequest) StartIndex(startIndex int32) ApiGetBusinesstransitionsBusinessBusinesstokenRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetBusinesstransitionsBusinessBusinesstokenRequest) Fields(fields string) ApiGetBusinesstransitionsBusinessBusinesstokenRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r ApiGetBusinesstransitionsBusinessBusinesstokenRequest) SortBy(sortBy string) ApiGetBusinesstransitionsBusinessBusinesstokenRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiGetBusinesstransitionsBusinessBusinesstokenRequest) Execute() (*BusinessTransitionListResponse, *http.Response, error) {
	return r.ApiService.GetBusinesstransitionsBusinessBusinesstokenExecute(r)
}

/*
GetBusinesstransitionsBusinessBusinesstoken List business transitions

List all transitions for a given business.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param businessToken Unique identifier of the business resource associated with the transitions to retrieve.
 @return ApiGetBusinesstransitionsBusinessBusinesstokenRequest
*/
func (a *BusinessTransitionsAPIService) GetBusinesstransitionsBusinessBusinesstoken(ctx context.Context, businessToken string) ApiGetBusinesstransitionsBusinessBusinesstokenRequest {
	return ApiGetBusinesstransitionsBusinessBusinesstokenRequest{
		ApiService: a,
		ctx: ctx,
		businessToken: businessToken,
	}
}

// Execute executes the request
//  @return BusinessTransitionListResponse
func (a *BusinessTransitionsAPIService) GetBusinesstransitionsBusinessBusinesstokenExecute(r ApiGetBusinesstransitionsBusinessBusinesstokenRequest) (*BusinessTransitionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessTransitionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessTransitionsAPIService.GetBusinesstransitionsBusinessBusinesstoken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesstransitions/business/{business_token}"
	localVarPath = strings.Replace(localVarPath, "{"+"business_token"+"}", url.PathEscape(parameterValueToString(r.businessToken, "businessToken")), -1)

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
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-id"
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

type ApiGetBusinesstransitionsTokenRequest struct {
	ctx context.Context
	ApiService *BusinessTransitionsAPIService
	token string
	fields *string
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r ApiGetBusinesstransitionsTokenRequest) Fields(fields string) ApiGetBusinesstransitionsTokenRequest {
	r.fields = &fields
	return r
}

func (r ApiGetBusinesstransitionsTokenRequest) Execute() (*BusinessTransitionResponse, *http.Response, error) {
	return r.ApiService.GetBusinesstransitionsTokenExecute(r)
}

/*
GetBusinesstransitionsToken Retrieve business transition

Use this endpoint to retrieve a business transition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token The unique identifier of the business transition you want to retrieve.
 @return ApiGetBusinesstransitionsTokenRequest
*/
func (a *BusinessTransitionsAPIService) GetBusinesstransitionsToken(ctx context.Context, token string) ApiGetBusinesstransitionsTokenRequest {
	return ApiGetBusinesstransitionsTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BusinessTransitionResponse
func (a *BusinessTransitionsAPIService) GetBusinesstransitionsTokenExecute(r ApiGetBusinesstransitionsTokenRequest) (*BusinessTransitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessTransitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessTransitionsAPIService.GetBusinesstransitionsToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesstransitions/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
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

type ApiPostBusinesstransitionsRequest struct {
	ctx context.Context
	ApiService *BusinessTransitionsAPIService
	businessTransitionRequest *BusinessTransitionRequest
}

func (r ApiPostBusinesstransitionsRequest) BusinessTransitionRequest(businessTransitionRequest BusinessTransitionRequest) ApiPostBusinesstransitionsRequest {
	r.businessTransitionRequest = &businessTransitionRequest
	return r
}

func (r ApiPostBusinesstransitionsRequest) Execute() (*BusinessTransitionResponse, *http.Response, error) {
	return r.ApiService.PostBusinesstransitionsExecute(r)
}

/*
PostBusinesstransitions Create business transition

This endpoint enables you to change a business' status, depending on your role and the previous status change.
By changing a business' status, you can control the business' capabilities and the setting of the `business.active` field.
The `business.active` field is `true` if your business is in the `LIMITED` or `ACTIVE` state, and `false` if your business is in the `UNVERIFIED` state.
You cannot control the value of the `business.active` field directly.

[cols=",2a,2a"]
|===
| The business.status field | Description | Business limitations

| Unverified
| Initial status of a newly created business belonging to an `accountholdergroup` where KYC is always required.
| Cannot load funds.

*The business.active field:*   +
`false`

*Allowable transitions:*   +
`ACTIVE`, `SUSPENDED`, `CLOSED`

| Limited
| Initial status of a newly created business belonging to an `accountholdergroup` where KYC is conditionally required.
| Restricted by rules in `accountholdergroups.pre_kyc_controls`.

*The business.active field:*   +
`true`

*Allowable transitions:*   +
`ACTIVE`, `SUSPENDED`, `CLOSED`

| Active
| Status of a business that has passed KYC; initial status of a newly created business belonging to an `accountholdergroup` where KYC is never required.
| None.

*The business.active field:*   +
`true`

*Allowable transitions:*   +
`SUSPENDED`, `CLOSED`

| Suspended
| The business is temporarily inactive.

*NOTE:* Transitioning a suspended business to the `ACTIVE` status is restricted, based on your role and the details of the previous status change.
| Cannot load funds or activate cards.

*The business.active field:*   +
`false`

*Allowable transitions:*   +
`ACTIVE`, `LIMITED`, `UNVERIFIED`, `CLOSED`

| Closed
| The business is permanently inactive.

*NOTE:* `CLOSED` is a terminal status.
In exceptional cases, you can transition a business from `CLOSED` to another status, depending on your role and the details of the previous status change.
Contact your Marqeta representative for more information.
| Cannot load funds.

*The business.active field:*   +
`false`

*Allowable transitions:*   +
`ACTIVE`, `LIMITED`, `UNVERIFIED`, `SUSPENDED`
|===

[NOTE]
The Marqeta platform transitions a business' status in response to certain events.
For example, a business with an `UNVERIFIED` status transitions to `ACTIVE` when the business passes KYC.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostBusinesstransitionsRequest
*/
func (a *BusinessTransitionsAPIService) PostBusinesstransitions(ctx context.Context) ApiPostBusinesstransitionsRequest {
	return ApiPostBusinesstransitionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BusinessTransitionResponse
func (a *BusinessTransitionsAPIService) PostBusinesstransitionsExecute(r ApiPostBusinesstransitionsRequest) (*BusinessTransitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessTransitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessTransitionsAPIService.PostBusinesstransitions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesstransitions"

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
	localVarPostBody = r.businessTransitionRequest
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