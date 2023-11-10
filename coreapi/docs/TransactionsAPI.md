# \TransactionsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactions**](TransactionsAPI.md#GetTransactions) | **Get** /transactions | List transactions
[**GetTransactionsFundingsourceFundingsourcetoken**](TransactionsAPI.md#GetTransactionsFundingsourceFundingsourcetoken) | **Get** /transactions/fundingsource/{funding_source_token} | List transactions for a funding account
[**GetTransactionsToken**](TransactionsAPI.md#GetTransactionsToken) | **Get** /transactions/{token} | Retrieve transaction
[**GetTransactionsTokenRelated**](TransactionsAPI.md#GetTransactionsTokenRelated) | **Get** /transactions/{token}/related | List related transactions



## GetTransactions

> TransactionModelListResponse GetTransactions(ctx).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Type_(type_).UserToken(userToken).BusinessToken(businessToken).ActingUserToken(actingUserToken).CardToken(cardToken).State(state).Version(version).Verbose(verbose).StartIdentifier(startIdentifier).Execute()

List transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    count := int32(56) // int32 | The number of transactions to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-user_transaction_time")
    startDate := "startDate_example" // string | The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the `start_date` and `end_date` fields. (optional)
    endDate := "endDate_example" // string | The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the `end_date` and `start_date` fields. (optional)
    type_ := "type__example" // string | Comma-delimited list of transaction types to include. (optional)
    userToken := "userToken_example" // string | The unique identifier of the user account holder. (optional)
    businessToken := "businessToken_example" // string | The unique identifier of the business account holder. (optional)
    actingUserToken := "actingUserToken_example" // string | The unique identifier of the acting user. (optional)
    cardToken := "cardToken_example" // string | The unique identifier of the card. (optional)
    state := "state_example" // string | Comma-delimited list of transaction states to display. (optional) (default to "PENDING,COMPLETION")
    version := "version_example" // string | Specifies the API version for the request. (optional)
    verbose := true // bool | If `true`, the query returns additional information for diagnostic purposes. (optional)
    startIdentifier := int64(789) // int64 | Start identifier (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetTransactions(context.Background()).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Type_(type_).UserToken(userToken).BusinessToken(businessToken).ActingUserToken(actingUserToken).CardToken(cardToken).State(state).Version(version).Verbose(verbose).StartIdentifier(startIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactions`: TransactionModelListResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of transactions to retrieve. | [default to 10]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-user_transaction_time&quot;]
 **startDate** | **string** | The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;start_date&#x60; and &#x60;end_date&#x60; fields. | 
 **endDate** | **string** | The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;end_date&#x60; and &#x60;start_date&#x60; fields. | 
 **type_** | **string** | Comma-delimited list of transaction types to include. | 
 **userToken** | **string** | The unique identifier of the user account holder. | 
 **businessToken** | **string** | The unique identifier of the business account holder. | 
 **actingUserToken** | **string** | The unique identifier of the acting user. | 
 **cardToken** | **string** | The unique identifier of the card. | 
 **state** | **string** | Comma-delimited list of transaction states to display. | [default to &quot;PENDING,COMPLETION&quot;]
 **version** | **string** | Specifies the API version for the request. | 
 **verbose** | **bool** | If &#x60;true&#x60;, the query returns additional information for diagnostic purposes. | 
 **startIdentifier** | **int64** | Start identifier | 

### Return type

[**TransactionModelListResponse**](TransactionModelListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsFundingsourceFundingsourcetoken

> TransactionModelListResponse GetTransactionsFundingsourceFundingsourcetoken(ctx, fundingSourceToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Type_(type_).Polarity(polarity).Version(version).Verbose(verbose).Execute()

List transactions for a funding account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    fundingSourceToken := "fundingSourceToken_example" // string | The unique identifier of the funding account.
    count := int32(56) // int32 | The number of transactions to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-user_transaction_time")
    startDate := "startDate_example" // string | The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the `start_date` and `end_date` fields. (optional)
    endDate := "endDate_example" // string | The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the `end_date` and `start_date` fields. (optional)
    type_ := "type__example" // string | Comma-delimited list of transaction types to include. (optional)
    polarity := "polarity_example" // string | Specifies whether to return credit or debit transactions. (optional)
    version := "version_example" // string | Specifies the API version for the request. (optional)
    verbose := true // bool | If `true`, the query returns additional information for diagnostic purposes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetTransactionsFundingsourceFundingsourcetoken(context.Background(), fundingSourceToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Type_(type_).Polarity(polarity).Version(version).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsFundingsourceFundingsourcetoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsFundingsourceFundingsourcetoken`: TransactionModelListResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsFundingsourceFundingsourcetoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceToken** | **string** | The unique identifier of the funding account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsFundingsourceFundingsourcetokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of transactions to retrieve. | [default to 10]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-user_transaction_time&quot;]
 **startDate** | **string** | The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;start_date&#x60; and &#x60;end_date&#x60; fields. | 
 **endDate** | **string** | The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;end_date&#x60; and &#x60;start_date&#x60; fields. | 
 **type_** | **string** | Comma-delimited list of transaction types to include. | 
 **polarity** | **string** | Specifies whether to return credit or debit transactions. | 
 **version** | **string** | Specifies the API version for the request. | 
 **verbose** | **bool** | If &#x60;true&#x60;, the query returns additional information for diagnostic purposes. | 

### Return type

[**TransactionModelListResponse**](TransactionModelListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsToken

> TransactionModel GetTransactionsToken(ctx, token).Fields(fields).Version(version).Verbose(verbose).Execute()

Retrieve transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    token := "token_example" // string | The unique identifier of the transaction.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    version := "version_example" // string | Specifies the API version for the request. (optional)
    verbose := true // bool | If `true`, the query returns additional information for diagnostic purposes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetTransactionsToken(context.Background(), token).Fields(fields).Version(version).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsToken`: TransactionModel
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **version** | **string** | Specifies the API version for the request. | 
 **verbose** | **bool** | If &#x60;true&#x60;, the query returns additional information for diagnostic purposes. | 

### Return type

[**TransactionModel**](TransactionModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsTokenRelated

> TransactionModelListResponse GetTransactionsTokenRelated(ctx, token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Type_(type_).State(state).Version(version).Verbose(verbose).Execute()

List related transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    token := "token_example" // string | The unique identifier of the transaction.
    count := int32(56) // int32 | The number of transactions to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-user_transaction_time")
    startDate := "startDate_example" // string | The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the `start_date` and `end_date` fields. (optional)
    endDate := "endDate_example" // string | The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the `end_date` and `start_date` fields. (optional)
    type_ := "type__example" // string | Comma-delimited list of transaction types to include. (optional)
    state := "state_example" // string | Comma-delimited list of transaction states to display. (optional) (default to "ALL")
    version := "version_example" // string | Specifies the API version for the request. (optional)
    verbose := true // bool | If `true`, the query returns additional information for diagnostic purposes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetTransactionsTokenRelated(context.Background(), token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Type_(type_).State(state).Version(version).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionsTokenRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsTokenRelated`: TransactionModelListResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionsTokenRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsTokenRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of transactions to retrieve. | [default to 10]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-user_transaction_time&quot;]
 **startDate** | **string** | The starting date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;start_date&#x60; and &#x60;end_date&#x60; fields. | 
 **endDate** | **string** | The ending date (or date-time) of a date range from which to return transactions. To return transactions for a single day, enter the same date in both the &#x60;end_date&#x60; and &#x60;start_date&#x60; fields. | 
 **type_** | **string** | Comma-delimited list of transaction types to include. | 
 **state** | **string** | Comma-delimited list of transaction states to display. | [default to &quot;ALL&quot;]
 **version** | **string** | Specifies the API version for the request. | 
 **verbose** | **bool** | If &#x60;true&#x60;, the query returns additional information for diagnostic purposes. | 

### Return type

[**TransactionModelListResponse**](TransactionModelListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

