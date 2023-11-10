# \AdjustmentsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdjustmentForAccount**](AdjustmentsAPI.md#CreateAdjustmentForAccount) | **Post** /credit/accounts/{account_token}/adjustments | Create account adjustment
[**GetAdjustmentsByAccount**](AdjustmentsAPI.md#GetAdjustmentsByAccount) | **Get** /credit/accounts/{account_token}/adjustments | List account adjustments
[**RetrieveAdjustment**](AdjustmentsAPI.md#RetrieveAdjustment) | **Get** /credit/accounts/{account_token}/adjustments/{adjustment_token} | Retrieve account adjustment



## CreateAdjustmentForAccount

> AccountAdjustmentResponse CreateAdjustmentForAccount(ctx, accountToken).AccountAdjustmentReq(accountAdjustmentReq).Execute()

Create account adjustment



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to create an adjustment.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    accountAdjustmentReq := *openapiclient.NewAccountAdjustmentReq(float32(123), openapiclient.CurrencyCode("USD"), "Description_example") // AccountAdjustmentReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsAPI.CreateAdjustmentForAccount(context.Background(), accountToken).AccountAdjustmentReq(accountAdjustmentReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsAPI.CreateAdjustmentForAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdjustmentForAccount`: AccountAdjustmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsAPI.CreateAdjustmentForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to create an adjustment.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdjustmentForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountAdjustmentReq** | [**AccountAdjustmentReq**](AccountAdjustmentReq.md) |  | 

### Return type

[**AccountAdjustmentResponse**](AccountAdjustmentResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdjustmentsByAccount

> AccountAdjustmentPage GetAdjustmentsByAccount(ctx, accountToken).Count(count).StartIndex(startIndex).Execute()

List account adjustments



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve adjustments.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    count := int32(56) // int32 | The number of resources to retrieve (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsAPI.GetAdjustmentsByAccount(context.Background(), accountToken).Count(count).StartIndex(startIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsAPI.GetAdjustmentsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdjustmentsByAccount`: AccountAdjustmentPage
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsAPI.GetAdjustmentsByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve adjustments.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdjustmentsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]

### Return type

[**AccountAdjustmentPage**](AccountAdjustmentPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAdjustment

> AccountAdjustmentResponse RetrieveAdjustment(ctx, accountToken, adjustmentToken).Execute()

Retrieve account adjustment



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve an adjustment.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    adjustmentToken := "adjustmentToken_example" // string | The unique identifier of the adjustment to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/adjustments` to retrieve existing account adjustment tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdjustmentsAPI.RetrieveAdjustment(context.Background(), accountToken, adjustmentToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentsAPI.RetrieveAdjustment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAdjustment`: AccountAdjustmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AdjustmentsAPI.RetrieveAdjustment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve an adjustment.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**adjustmentToken** | **string** | The unique identifier of the adjustment to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/adjustments&#x60; to retrieve existing account adjustment tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAdjustmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountAdjustmentResponse**](AccountAdjustmentResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

