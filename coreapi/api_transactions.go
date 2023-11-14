/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TransactionsAPIService TransactionsAPI service
type TransactionsAPIService service

type ApiGetTransactionsRequest struct {
	ctx context.Context
	ApiService *TransactionsAPIService
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
	startDate *string
	endDate *string
	type_ *string
	userToken *string
	businessToken *string
	actingUserToken *string
	cardToken *string
	state *string
	version *string
	verbose *bool
	startIdentifier *int64
}

// The number of transactions to retrieve.
func (r ApiGetTransactionsRequest) Count(count int32) ApiGetTransactionsRequest {
	r.count = &count
	return r
}

// The sort order index of the first resource in the returned array.
func (r ApiGetTransactionsRequest) StartIndex(startIndex int32) ApiGetTransactionsRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetTransactionsRequest) Fields(fields string) ApiGetTransactionsRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r ApiGetTransactionsRequest) SortBy(sortBy string) ApiGetTransactionsRequest {
	r.sortBy = &sortBy
	return r
}

// The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;start_date&#x60; and &#x60;end_date&#x60; fields.
func (r ApiGetTransactionsRequest) StartDate(startDate string) ApiGetTransactionsRequest {
	r.startDate = &startDate
	return r
}

// The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;end_date&#x60; and &#x60;start_date&#x60; fields.
func (r ApiGetTransactionsRequest) EndDate(endDate string) ApiGetTransactionsRequest {
	r.endDate = &endDate
	return r
}

// Comma-delimited list of transaction types to include.
func (r ApiGetTransactionsRequest) Type_(type_ string) ApiGetTransactionsRequest {
	r.type_ = &type_
	return r
}

// The unique identifier of the user account holder.
func (r ApiGetTransactionsRequest) UserToken(userToken string) ApiGetTransactionsRequest {
	r.userToken = &userToken
	return r
}

// The unique identifier of the business account holder.
func (r ApiGetTransactionsRequest) BusinessToken(businessToken string) ApiGetTransactionsRequest {
	r.businessToken = &businessToken
	return r
}

// The unique identifier of the acting user.
func (r ApiGetTransactionsRequest) ActingUserToken(actingUserToken string) ApiGetTransactionsRequest {
	r.actingUserToken = &actingUserToken
	return r
}

// The unique identifier of the card.
func (r ApiGetTransactionsRequest) CardToken(cardToken string) ApiGetTransactionsRequest {
	r.cardToken = &cardToken
	return r
}

// Comma-delimited list of transaction states to display.
func (r ApiGetTransactionsRequest) State(state string) ApiGetTransactionsRequest {
	r.state = &state
	return r
}

// Specifies the API version for the request.
func (r ApiGetTransactionsRequest) Version(version string) ApiGetTransactionsRequest {
	r.version = &version
	return r
}

// If &#x60;true&#x60;, the query returns additional information for diagnostic purposes.
func (r ApiGetTransactionsRequest) Verbose(verbose bool) ApiGetTransactionsRequest {
	r.verbose = &verbose
	return r
}

// Start identifier
func (r ApiGetTransactionsRequest) StartIdentifier(startIdentifier int64) ApiGetTransactionsRequest {
	r.startIdentifier = &startIdentifier
	return r
}

func (r ApiGetTransactionsRequest) Execute() (*TransactionModelListResponse, *http.Response, error) {
	return r.ApiService.GetTransactionsExecute(r)
}

/*
GetTransactions List transactions

List all transactions.

By default, this endpoint returns transactions conducted within the last 30 days.
To return transactions older than 30 days, you must include the `start_date` and `end_date` query parameters in your request.

By default, `GET /transactions` returns transactions having either `PENDING` or `COMPLETION` states.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTransactionsRequest
*/
func (a *TransactionsAPIService) GetTransactions(ctx context.Context) ApiGetTransactionsRequest {
	return ApiGetTransactionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TransactionModelListResponse
func (a *TransactionsAPIService) GetTransactionsExecute(r ApiGetTransactionsRequest) (*TransactionModelListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModelListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 10
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
		var defaultValue string = "-user_transaction_time"
		r.sortBy = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.userToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_token", r.userToken, "")
	}
	if r.businessToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "business_token", r.businessToken, "")
	}
	if r.actingUserToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "acting_user_token", r.actingUserToken, "")
	}
	if r.cardToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "card_token", r.cardToken, "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
	} else {
		var defaultValue string = "PENDING,COMPLETION"
		r.state = &defaultValue
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.verbose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "verbose", r.verbose, "")
	}
	if r.startIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_identifier", r.startIdentifier, "")
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

type ApiGetTransactionsFundingsourceFundingsourcetokenRequest struct {
	ctx context.Context
	ApiService *TransactionsAPIService
	fundingSourceToken string
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
	startDate *string
	endDate *string
	type_ *string
	polarity *string
	version *string
	verbose *bool
}

// The number of transactions to retrieve.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) Count(count int32) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.count = &count
	return r
}

// The sort order index of the first resource in the returned array.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) StartIndex(startIndex int32) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) Fields(fields string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) SortBy(sortBy string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.sortBy = &sortBy
	return r
}

// The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;start_date&#x60; and &#x60;end_date&#x60; fields.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) StartDate(startDate string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.startDate = &startDate
	return r
}

// The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;end_date&#x60; and &#x60;start_date&#x60; fields.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) EndDate(endDate string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.endDate = &endDate
	return r
}

// Comma-delimited list of transaction types to include.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) Type_(type_ string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.type_ = &type_
	return r
}

// Specifies whether to return credit or debit transactions.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) Polarity(polarity string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.polarity = &polarity
	return r
}

// Specifies the API version for the request.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) Version(version string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.version = &version
	return r
}

// If &#x60;true&#x60;, the query returns additional information for diagnostic purposes.
func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) Verbose(verbose bool) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	r.verbose = &verbose
	return r
}

func (r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) Execute() (*TransactionModelListResponse, *http.Response, error) {
	return r.ApiService.GetTransactionsFundingsourceFundingsourcetokenExecute(r)
}

/*
GetTransactionsFundingsourceFundingsourcetoken List transactions for a funding account

List transactions for a specific funding source.

By default, this endpoint returns transactions conducted within the last 30 days.
To return transactions older than 30 days, you must include the `start_date` and `end_date` query parameters in your request.

By default, `GET /transactions/fundingsource/{funding_source_token}` returns transactions having either `PENDING` or `COMPLETION` states.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fundingSourceToken The unique identifier of the funding account.
 @return ApiGetTransactionsFundingsourceFundingsourcetokenRequest
*/
func (a *TransactionsAPIService) GetTransactionsFundingsourceFundingsourcetoken(ctx context.Context, fundingSourceToken string) ApiGetTransactionsFundingsourceFundingsourcetokenRequest {
	return ApiGetTransactionsFundingsourceFundingsourcetokenRequest{
		ApiService: a,
		ctx: ctx,
		fundingSourceToken: fundingSourceToken,
	}
}

// Execute executes the request
//  @return TransactionModelListResponse
func (a *TransactionsAPIService) GetTransactionsFundingsourceFundingsourcetokenExecute(r ApiGetTransactionsFundingsourceFundingsourcetokenRequest) (*TransactionModelListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModelListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionsFundingsourceFundingsourcetoken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/fundingsource/{funding_source_token}"
	localVarPath = strings.Replace(localVarPath, "{"+"funding_source_token"+"}", url.PathEscape(parameterValueToString(r.fundingSourceToken, "fundingSourceToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 10
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
		var defaultValue string = "-user_transaction_time"
		r.sortBy = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.polarity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "polarity", r.polarity, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.verbose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "verbose", r.verbose, "")
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

type ApiGetTransactionsTokenRequest struct {
	ctx context.Context
	ApiService *TransactionsAPIService
	token string
	fields *string
	version *string
	verbose *bool
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetTransactionsTokenRequest) Fields(fields string) ApiGetTransactionsTokenRequest {
	r.fields = &fields
	return r
}

// Specifies the API version for the request.
func (r ApiGetTransactionsTokenRequest) Version(version string) ApiGetTransactionsTokenRequest {
	r.version = &version
	return r
}

// If &#x60;true&#x60;, the query returns additional information for diagnostic purposes.
func (r ApiGetTransactionsTokenRequest) Verbose(verbose bool) ApiGetTransactionsTokenRequest {
	r.verbose = &verbose
	return r
}

func (r ApiGetTransactionsTokenRequest) Execute() (*TransactionModel, *http.Response, error) {
	return r.ApiService.GetTransactionsTokenExecute(r)
}

/*
GetTransactionsToken Retrieve transaction

Retrieve a specific transaction.
Include the `token` path parameter to identify the transaction.

[NOTE]
Transactions are not available in real time via this endpoint, and typically appear after 30 seconds.
On occasion, a transaction may require up to four hours to appear.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token The unique identifier of the transaction.
 @return ApiGetTransactionsTokenRequest
*/
func (a *TransactionsAPIService) GetTransactionsToken(ctx context.Context, token string) ApiGetTransactionsTokenRequest {
	return ApiGetTransactionsTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return TransactionModel
func (a *TransactionsAPIService) GetTransactionsTokenExecute(r ApiGetTransactionsTokenRequest) (*TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionsToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.verbose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "verbose", r.verbose, "")
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

type ApiGetTransactionsTokenRelatedRequest struct {
	ctx context.Context
	ApiService *TransactionsAPIService
	token string
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
	startDate *string
	endDate *string
	type_ *string
	state *string
	version *string
	verbose *bool
}

// The number of transactions to retrieve.
func (r ApiGetTransactionsTokenRelatedRequest) Count(count int32) ApiGetTransactionsTokenRelatedRequest {
	r.count = &count
	return r
}

// The sort order index of the first resource in the returned array.
func (r ApiGetTransactionsTokenRelatedRequest) StartIndex(startIndex int32) ApiGetTransactionsTokenRelatedRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r ApiGetTransactionsTokenRelatedRequest) Fields(fields string) ApiGetTransactionsTokenRelatedRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r ApiGetTransactionsTokenRelatedRequest) SortBy(sortBy string) ApiGetTransactionsTokenRelatedRequest {
	r.sortBy = &sortBy
	return r
}

// The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;start_date&#x60; and &#x60;end_date&#x60; fields.
func (r ApiGetTransactionsTokenRelatedRequest) StartDate(startDate string) ApiGetTransactionsTokenRelatedRequest {
	r.startDate = &startDate
	return r
}

// The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;end_date&#x60; and &#x60;start_date&#x60; fields.
func (r ApiGetTransactionsTokenRelatedRequest) EndDate(endDate string) ApiGetTransactionsTokenRelatedRequest {
	r.endDate = &endDate
	return r
}

// Comma-delimited list of transaction types to include.
func (r ApiGetTransactionsTokenRelatedRequest) Type_(type_ string) ApiGetTransactionsTokenRelatedRequest {
	r.type_ = &type_
	return r
}

// Comma-delimited list of transaction states to display.
func (r ApiGetTransactionsTokenRelatedRequest) State(state string) ApiGetTransactionsTokenRelatedRequest {
	r.state = &state
	return r
}

// Specifies the API version for the request.
func (r ApiGetTransactionsTokenRelatedRequest) Version(version string) ApiGetTransactionsTokenRelatedRequest {
	r.version = &version
	return r
}

// If &#x60;true&#x60;, the query returns additional information for diagnostic purposes.
func (r ApiGetTransactionsTokenRelatedRequest) Verbose(verbose bool) ApiGetTransactionsTokenRelatedRequest {
	r.verbose = &verbose
	return r
}

func (r ApiGetTransactionsTokenRelatedRequest) Execute() (*TransactionModelListResponse, *http.Response, error) {
	return r.ApiService.GetTransactionsTokenRelatedExecute(r)
}

/*
GetTransactionsTokenRelated List related transactions

List all transactions related to the specified transaction.

By default, this endpoint returns transactions conducted within the last 30 days.
To return transactions older than 30 days, you must include the `start_date` and `end_date` query parameters in your request.

By default, this endpoint returns transactions of any state.
To return transactions in specific states, you must include the `state` query parameter in your request.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token The unique identifier of the transaction.
 @return ApiGetTransactionsTokenRelatedRequest
*/
func (a *TransactionsAPIService) GetTransactionsTokenRelated(ctx context.Context, token string) ApiGetTransactionsTokenRelatedRequest {
	return ApiGetTransactionsTokenRelatedRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return TransactionModelListResponse
func (a *TransactionsAPIService) GetTransactionsTokenRelatedExecute(r ApiGetTransactionsTokenRelatedRequest) (*TransactionModelListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModelListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsAPIService.GetTransactionsTokenRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{token}/related"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 10
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
		var defaultValue string = "-user_transaction_time"
		r.sortBy = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
	} else {
		var defaultValue string = "ALL"
		r.state = &defaultValue
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.verbose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "verbose", r.verbose, "")
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
