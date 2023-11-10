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


// VelocityControlsAPIService VelocityControlsAPI service
type VelocityControlsAPIService service

type ApiGetVelocitycontrolsRequest struct {
	ctx context.Context
	ApiService *VelocityControlsAPIService
	cardProduct *string
	user *string
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
}

// Unique identifier of the card product. Enter the string &#x60;null&#x60; to retrieve velocity controls that are not associated with any card product.
func (r ApiGetVelocitycontrolsRequest) CardProduct(cardProduct string) ApiGetVelocitycontrolsRequest {
	r.cardProduct = &cardProduct
	return r
}

// Unique identifier of the user. Enter the string &#x60;null&#x60; to retrieve velocity controls that are not associated with any user.
func (r ApiGetVelocitycontrolsRequest) User(user string) ApiGetVelocitycontrolsRequest {
	r.user = &user
	return r
}

// Number of velocity control resources to retrieve.
func (r ApiGetVelocitycontrolsRequest) Count(count int32) ApiGetVelocitycontrolsRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r ApiGetVelocitycontrolsRequest) StartIndex(startIndex int32) ApiGetVelocitycontrolsRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetVelocitycontrolsRequest) Fields(fields string) ApiGetVelocitycontrolsRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r ApiGetVelocitycontrolsRequest) SortBy(sortBy string) ApiGetVelocitycontrolsRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiGetVelocitycontrolsRequest) Execute() (*VelocityControlListResponse, *http.Response, error) {
	return r.ApiService.GetVelocitycontrolsExecute(r)
}

/*
GetVelocitycontrols List velocity controls

Retrieves a list of all the velocity controls associated with a specific user or card product, or lists all the velocity controls defined for your program.

Include either a `user` or a `card_product` query parameter to indicate the user or card product whose associated velocity controls you want to retrieve (do not include both).

To list all velocity controls for your program, omit the `user` and `card_product` query parameters from your request.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVelocitycontrolsRequest
*/
func (a *VelocityControlsAPIService) GetVelocitycontrols(ctx context.Context) ApiGetVelocitycontrolsRequest {
	return ApiGetVelocitycontrolsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VelocityControlListResponse
func (a *VelocityControlsAPIService) GetVelocitycontrolsExecute(r ApiGetVelocitycontrolsRequest) (*VelocityControlListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VelocityControlListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VelocityControlsAPIService.GetVelocitycontrols")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/velocitycontrols"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cardProduct != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "card_product", r.cardProduct, "")
	}
	if r.user != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user", r.user, "")
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
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
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

type ApiGetVelocitycontrolsTokenRequest struct {
	ctx context.Context
	ApiService *VelocityControlsAPIService
	token string
	fields *string
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetVelocitycontrolsTokenRequest) Fields(fields string) ApiGetVelocitycontrolsTokenRequest {
	r.fields = &fields
	return r
}

func (r ApiGetVelocitycontrolsTokenRequest) Execute() (*VelocityControlResponse, *http.Response, error) {
	return r.ApiService.GetVelocitycontrolsTokenExecute(r)
}

/*
GetVelocitycontrolsToken Returns a specific velocity control

Retrieves a specific velocity control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the velocity control resource.
 @return ApiGetVelocitycontrolsTokenRequest
*/
func (a *VelocityControlsAPIService) GetVelocitycontrolsToken(ctx context.Context, token string) ApiGetVelocitycontrolsTokenRequest {
	return ApiGetVelocitycontrolsTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return VelocityControlResponse
func (a *VelocityControlsAPIService) GetVelocitycontrolsTokenExecute(r ApiGetVelocitycontrolsTokenRequest) (*VelocityControlResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VelocityControlResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VelocityControlsAPIService.GetVelocitycontrolsToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/velocitycontrols/{token}"
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

type ApiGetVelocitycontrolsUserUsertokenAvailableRequest struct {
	ctx context.Context
	ApiService *VelocityControlsAPIService
	userToken string
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
}

// Number of available balance resources to retrieve.
func (r ApiGetVelocitycontrolsUserUsertokenAvailableRequest) Count(count int32) ApiGetVelocitycontrolsUserUsertokenAvailableRequest {
	r.count = &count
	return r
}

// The sort order index of the first resource in the returned array.
func (r ApiGetVelocitycontrolsUserUsertokenAvailableRequest) StartIndex(startIndex int32) ApiGetVelocitycontrolsUserUsertokenAvailableRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetVelocitycontrolsUserUsertokenAvailableRequest) Fields(fields string) ApiGetVelocitycontrolsUserUsertokenAvailableRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r ApiGetVelocitycontrolsUserUsertokenAvailableRequest) SortBy(sortBy string) ApiGetVelocitycontrolsUserUsertokenAvailableRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiGetVelocitycontrolsUserUsertokenAvailableRequest) Execute() (*VelocityControlBalanceListResponse, *http.Response, error) {
	return r.ApiService.GetVelocitycontrolsUserUsertokenAvailableExecute(r)
}

/*
GetVelocitycontrolsUserUsertokenAvailable List user velocity control balances

Retrieves a list of the available balances of the velocity controls associated with a user.
This operation is unavailable for velocity controls with a window of `TRANSACTION`, because available balances do not apply to single-transaction velocity windows.

Depending on the control, the balance can include an available monetary amount, the number of transactions remaining, and the number of days remaining in the time window.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userToken Unique identifier of the cardholder.
 @return ApiGetVelocitycontrolsUserUsertokenAvailableRequest
*/
func (a *VelocityControlsAPIService) GetVelocitycontrolsUserUsertokenAvailable(ctx context.Context, userToken string) ApiGetVelocitycontrolsUserUsertokenAvailableRequest {
	return ApiGetVelocitycontrolsUserUsertokenAvailableRequest{
		ApiService: a,
		ctx: ctx,
		userToken: userToken,
	}
}

// Execute executes the request
//  @return VelocityControlBalanceListResponse
func (a *VelocityControlsAPIService) GetVelocitycontrolsUserUsertokenAvailableExecute(r ApiGetVelocitycontrolsUserUsertokenAvailableRequest) (*VelocityControlBalanceListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VelocityControlBalanceListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VelocityControlsAPIService.GetVelocitycontrolsUserUsertokenAvailable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/velocitycontrols/user/{user_token}/available"
	localVarPath = strings.Replace(localVarPath, "{"+"user_token"+"}", url.PathEscape(parameterValueToString(r.userToken, "userToken")), -1)

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

type ApiPostVelocitycontrolsRequest struct {
	ctx context.Context
	ApiService *VelocityControlsAPIService
	velocityControlRequest *VelocityControlRequest
}

// Velocity control object
func (r ApiPostVelocitycontrolsRequest) VelocityControlRequest(velocityControlRequest VelocityControlRequest) ApiPostVelocitycontrolsRequest {
	r.velocityControlRequest = &velocityControlRequest
	return r
}

func (r ApiPostVelocitycontrolsRequest) Execute() (*VelocityControlResponse, *http.Response, error) {
	return r.ApiService.PostVelocitycontrolsExecute(r)
}

/*
PostVelocitycontrols Create velocity control

Limits how much and how frequently a user can spend funds.
If multiple velocity controls apply to the same user, the user cannot exceed any of the defined spending limits.

[TIP]
You can create program-level controls in the sandbox environment.
However, you must work with your Marqeta representative to create program-level controls in a production environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostVelocitycontrolsRequest
*/
func (a *VelocityControlsAPIService) PostVelocitycontrols(ctx context.Context) ApiPostVelocitycontrolsRequest {
	return ApiPostVelocitycontrolsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VelocityControlResponse
func (a *VelocityControlsAPIService) PostVelocitycontrolsExecute(r ApiPostVelocitycontrolsRequest) (*VelocityControlResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VelocityControlResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VelocityControlsAPIService.PostVelocitycontrols")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/velocitycontrols"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.velocityControlRequest == nil {
		return localVarReturnValue, nil, reportError("velocityControlRequest is required and must be specified")
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
	localVarPostBody = r.velocityControlRequest
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

type ApiPutVelocitycontrolsTokenRequest struct {
	ctx context.Context
	ApiService *VelocityControlsAPIService
	token string
	velocityControlUpdateRequest *VelocityControlUpdateRequest
}

// Velocity control object
func (r ApiPutVelocitycontrolsTokenRequest) VelocityControlUpdateRequest(velocityControlUpdateRequest VelocityControlUpdateRequest) ApiPutVelocitycontrolsTokenRequest {
	r.velocityControlUpdateRequest = &velocityControlUpdateRequest
	return r
}

func (r ApiPutVelocitycontrolsTokenRequest) Execute() (*VelocityControlResponse, *http.Response, error) {
	return r.ApiService.PutVelocitycontrolsTokenExecute(r)
}

/*
PutVelocitycontrolsToken Update velocity control

Updates a specific velocity control.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the velocity control resource.
 @return ApiPutVelocitycontrolsTokenRequest
*/
func (a *VelocityControlsAPIService) PutVelocitycontrolsToken(ctx context.Context, token string) ApiPutVelocitycontrolsTokenRequest {
	return ApiPutVelocitycontrolsTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return VelocityControlResponse
func (a *VelocityControlsAPIService) PutVelocitycontrolsTokenExecute(r ApiPutVelocitycontrolsTokenRequest) (*VelocityControlResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VelocityControlResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VelocityControlsAPIService.PutVelocitycontrolsToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/velocitycontrols/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.velocityControlUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("velocityControlUpdateRequest is required and must be specified")
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
	localVarPostBody = r.velocityControlUpdateRequest
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
