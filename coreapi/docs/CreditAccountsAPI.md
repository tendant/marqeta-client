# \CreditAccountsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreditAccount**](CreditAccountsAPI.md#CreateCreditAccount) | **Post** /credit/accounts | Create account
[**ListAccounts**](CreditAccountsAPI.md#ListAccounts) | **Get** /credit/accounts | List accounts
[**RetrieveAccount**](CreditAccountsAPI.md#RetrieveAccount) | **Get** /credit/accounts/{account_token} | Retrieve account
[**UpdateAccount**](CreditAccountsAPI.md#UpdateAccount) | **Put** /credit/accounts/{account_token} | Update account



## CreateCreditAccount

> AccountResponse CreateCreditAccount(ctx).AccountCreateReq(accountCreateReq).Execute()

Create account



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
    accountCreateReq := *openapiclient.NewAccountCreateReq(float32(123), []openapiclient.AccountUsageCreateReq{*openapiclient.NewAccountUsageCreateReq([]openapiclient.AprScheduleCreateReq{*openapiclient.NewAprScheduleCreateReq([]openapiclient.AprScheduleEntryCreateReq{*openapiclient.NewAprScheduleEntryCreateReq(float32(123))}, openapiclient.AccountAprType("GO_TO"))}, openapiclient.BalanceType("PURCHASE"))}, "UserToken_example") // AccountCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditAccountsAPI.CreateCreditAccount(context.Background()).AccountCreateReq(accountCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditAccountsAPI.CreateCreditAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreditAccount`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditAccountsAPI.CreateCreditAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountCreateReq** | [**AccountCreateReq**](AccountCreateReq.md) |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountsPage ListAccounts(ctx).CardToken(cardToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List accounts



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
    cardToken := "cardToken_example" // string | The unique identifier of the credit card associated with the account. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditAccountsAPI.ListAccounts(context.Background()).CardToken(cardToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditAccountsAPI.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: AccountsPage
    fmt.Fprintf(os.Stdout, "Response from `CreditAccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardToken** | **string** | The unique identifier of the credit card associated with the account. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**AccountsPage**](AccountsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAccount

> AccountResponse RetrieveAccount(ctx, accountToken).Execute()

Retrieve account



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account to retrieve.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditAccountsAPI.RetrieveAccount(context.Background(), accountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditAccountsAPI.RetrieveAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAccount`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditAccountsAPI.RetrieveAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> AccountResponse UpdateAccount(ctx, accountToken).AccountUpdateReq(accountUpdateReq).Execute()

Update account



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account to update.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    accountUpdateReq := *openapiclient.NewAccountUpdateReq() // AccountUpdateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditAccountsAPI.UpdateAccount(context.Background(), accountToken).AccountUpdateReq(accountUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditAccountsAPI.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditAccountsAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account to update.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountUpdateReq** | [**AccountUpdateReq**](AccountUpdateReq.md) |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

