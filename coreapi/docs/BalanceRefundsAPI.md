# \BalanceRefundsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBalanceRefund**](BalanceRefundsAPI.md#CreateBalanceRefund) | **Post** /credit/accounts/{account_token}/creditbalancerefunds | Create balance refund



## CreateBalanceRefund

> AccountCreditBalanceRefundResponse CreateBalanceRefund(ctx, accountToken).AccountCreditBalanceRefundReq(accountCreditBalanceRefundReq).Execute()

Create balance refund



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to create a balance refund.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    accountCreditBalanceRefundReq := *openapiclient.NewAccountCreditBalanceRefundReq(float32(123), openapiclient.CurrencyCode("USD"), "Description_example", openapiclient.RefundMethod("CHECK")) // AccountCreditBalanceRefundReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalanceRefundsAPI.CreateBalanceRefund(context.Background(), accountToken).AccountCreditBalanceRefundReq(accountCreditBalanceRefundReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalanceRefundsAPI.CreateBalanceRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBalanceRefund`: AccountCreditBalanceRefundResponse
    fmt.Fprintf(os.Stdout, "Response from `BalanceRefundsAPI.CreateBalanceRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to create a balance refund.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBalanceRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountCreditBalanceRefundReq** | [**AccountCreditBalanceRefundReq**](AccountCreditBalanceRefundReq.md) |  | 

### Return type

[**AccountCreditBalanceRefundResponse**](AccountCreditBalanceRefundResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

