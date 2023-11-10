# \CreditDisputesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDisputeForAccount**](CreditDisputesAPI.md#CreateDisputeForAccount) | **Post** /credit/accounts/{account_token}/disputes | Create account dispute
[**GetDisputesByAccount**](CreditDisputesAPI.md#GetDisputesByAccount) | **Get** /credit/accounts/{account_token}/disputes | List account disputes
[**RetrieveDispute**](CreditDisputesAPI.md#RetrieveDispute) | **Get** /credit/accounts/{account_token}/disputes/{dispute_token} | Retrieve account dispute
[**TransitionDispute**](CreditDisputesAPI.md#TransitionDispute) | **Post** /credit/accounts/{account_token}/disputes/{dispute_token}/transitions | Update account dispute



## CreateDisputeForAccount

> DisputeResponse CreateDisputeForAccount(ctx, accountToken).DisputeCreateReq(disputeCreateReq).Execute()

Create account dispute



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to create a dispute.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    disputeCreateReq := *openapiclient.NewDisputeCreateReq(float32(123), openapiclient.DisputeCategory("FRAUD"), "LedgerEntryToken_example") // DisputeCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditDisputesAPI.CreateDisputeForAccount(context.Background(), accountToken).DisputeCreateReq(disputeCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditDisputesAPI.CreateDisputeForAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDisputeForAccount`: DisputeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditDisputesAPI.CreateDisputeForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to create a dispute.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDisputeForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disputeCreateReq** | [**DisputeCreateReq**](DisputeCreateReq.md) |  | 

### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisputesByAccount

> DisputeResponsePage GetDisputesByAccount(ctx, accountToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List account disputes



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to retrieve the disputes.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditDisputesAPI.GetDisputesByAccount(context.Background(), accountToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditDisputesAPI.GetDisputesByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDisputesByAccount`: DisputeResponsePage
    fmt.Fprintf(os.Stdout, "Response from `CreditDisputesAPI.GetDisputesByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to retrieve the disputes.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisputesByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**DisputeResponsePage**](DisputeResponsePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDispute

> DisputeResponse RetrieveDispute(ctx, accountToken, disputeToken).Execute()

Retrieve account dispute



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account from which to retrieve a dispute.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    disputeToken := "disputeToken_example" // string | The unique identifier of the dispute to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/disputes` to retrieve existing dispute tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditDisputesAPI.RetrieveDispute(context.Background(), accountToken, disputeToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditDisputesAPI.RetrieveDispute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDispute`: DisputeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditDisputesAPI.RetrieveDispute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account from which to retrieve a dispute.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**disputeToken** | **string** | The unique identifier of the dispute to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/disputes&#x60; to retrieve existing dispute tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DisputeResponse**](DisputeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionDispute

> DisputeTransitionResponse TransitionDispute(ctx, accountToken, disputeToken).DisputeTransitionReq(disputeTransitionReq).Execute()

Update account dispute



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account from which to update a dispute.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    disputeToken := "disputeToken_example" // string | The unique identifier of the dispute to update.  Send a `GET` request to `/credit/accounts/{account_token}/disputes` to retrieve existing credit account tokens.
    disputeTransitionReq := *openapiclient.NewDisputeTransitionReq(float32(123), openapiclient.DisputeStatus("ACTIVE")) // DisputeTransitionReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditDisputesAPI.TransitionDispute(context.Background(), accountToken, disputeToken).DisputeTransitionReq(disputeTransitionReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditDisputesAPI.TransitionDispute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransitionDispute`: DisputeTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditDisputesAPI.TransitionDispute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account from which to update a dispute.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**disputeToken** | **string** | The unique identifier of the dispute to update.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/disputes&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disputeTransitionReq** | [**DisputeTransitionReq**](DisputeTransitionReq.md) |  | 

### Return type

[**DisputeTransitionResponse**](DisputeTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

