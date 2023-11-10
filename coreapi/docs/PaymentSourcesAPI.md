# \PaymentSourcesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentSource**](PaymentSourcesAPI.md#CreatePaymentSource) | **Post** /credit/paymentsources | Create payment source
[**ListPaymentSources**](PaymentSourcesAPI.md#ListPaymentSources) | **Get** /credit/paymentsources | List payment sources
[**RetrievePaymentSource**](PaymentSourcesAPI.md#RetrievePaymentSource) | **Get** /credit/paymentsources/{token} | Retrieve payment source
[**UpdatePaymentSource**](PaymentSourcesAPI.md#UpdatePaymentSource) | **Put** /credit/paymentsources/{token} | Update payment source



## CreatePaymentSource

> PaymentSourceResponse CreatePaymentSource(ctx).PaymentSourceCreateReq(paymentSourceCreateReq).Execute()

Create payment source



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
    paymentSourceCreateReq := *openapiclient.NewPaymentSourceCreateReq("AccountNumber_example", "AccountToken_example", "Name_example", "RoutingNumber_example", "SourceType_example", "UserToken_example", false) // PaymentSourceCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSourcesAPI.CreatePaymentSource(context.Background()).PaymentSourceCreateReq(paymentSourceCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSourcesAPI.CreatePaymentSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentSource`: PaymentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentSourcesAPI.CreatePaymentSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentSourceCreateReq** | [**PaymentSourceCreateReq**](PaymentSourceCreateReq.md) |  | 

### Return type

[**PaymentSourceResponse**](PaymentSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentSources

> PaymentSourcePage ListPaymentSources(ctx).AccountToken(accountToken).UserToken(userToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List payment sources



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account associated with the payment source. (optional)
    userToken := "userToken_example" // string | The unique identifier of the user associated with the payment source. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSourcesAPI.ListPaymentSources(context.Background()).AccountToken(accountToken).UserToken(userToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSourcesAPI.ListPaymentSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentSources`: PaymentSourcePage
    fmt.Fprintf(os.Stdout, "Response from `PaymentSourcesAPI.ListPaymentSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountToken** | **string** | The unique identifier of the credit account associated with the payment source. | 
 **userToken** | **string** | The unique identifier of the user associated with the payment source. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**PaymentSourcePage**](PaymentSourcePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePaymentSource

> PaymentSourceResponse RetrievePaymentSource(ctx, token).Execute()

Retrieve payment source



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
    token := "token_example" // string | The unique identifier of the payment source to retrieve.  Send a `GET` request to `/credit/paymentsources` to retrieve existing payment source tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSourcesAPI.RetrievePaymentSource(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSourcesAPI.RetrievePaymentSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePaymentSource`: PaymentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentSourcesAPI.RetrievePaymentSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the payment source to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/paymentsources&#x60; to retrieve existing payment source tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentSourceResponse**](PaymentSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentSource

> PaymentSourceResponse UpdatePaymentSource(ctx, token).PaymentSourceUpdateReq(paymentSourceUpdateReq).Execute()

Update payment source



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
    token := "token_example" // string | The unique identifier of the payment source to retrieve.  Send a `GET` request to `/credit/paymentsources` to retrieve existing payment source tokens.
    paymentSourceUpdateReq := *openapiclient.NewPaymentSourceUpdateReq(openapiclient.PaymentSourceStatusEnum("ACTIVE")) // PaymentSourceUpdateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSourcesAPI.UpdatePaymentSource(context.Background(), token).PaymentSourceUpdateReq(paymentSourceUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSourcesAPI.UpdatePaymentSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentSource`: PaymentSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentSourcesAPI.UpdatePaymentSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the payment source to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/paymentsources&#x60; to retrieve existing payment source tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentSourceUpdateReq** | [**PaymentSourceUpdateReq**](PaymentSourceUpdateReq.md) |  | 

### Return type

[**PaymentSourceResponse**](PaymentSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

