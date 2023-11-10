# \PushToCardAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPushtocardsDisburse**](PushToCardAPI.md#GetPushtocardsDisburse) | **Get** /pushtocards/disburse | Lists all push-to-card disbursements
[**GetPushtocardsDisburseToken**](PushToCardAPI.md#GetPushtocardsDisburseToken) | **Get** /pushtocards/disburse/{token} | Returns a specific push-to-card disbursement
[**GetPushtocardsPaymentcard**](PushToCardAPI.md#GetPushtocardsPaymentcard) | **Get** /pushtocards/paymentcard | Returns all push-to-card payment card details
[**GetPushtocardsPaymentcardToken**](PushToCardAPI.md#GetPushtocardsPaymentcardToken) | **Get** /pushtocards/paymentcard/{token} | Returns a specific paymentcard object
[**PostPushtocardsDisburse**](PushToCardAPI.md#PostPushtocardsDisburse) | **Post** /pushtocards/disburse | Initiates a push-to-card money disbursement
[**PostPushtocardsPaymentcard**](PushToCardAPI.md#PostPushtocardsPaymentcard) | **Post** /pushtocards/paymentcard | Adds an external card to which funds will be pushed



## GetPushtocardsDisburse

> PushToCardDisburseListResponse GetPushtocardsDisburse(ctx).Count(count).Fields(fields).StartIndex(startIndex).SortBy(sortBy).Execute()

Lists all push-to-card disbursements

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
    count := int32(56) // int32 | Number of push-to-card disbursements to retrieve (optional) (default to 10)
    fields := "fields_example" // string | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. (optional)
    startIndex := int32(56) // int32 | Start index (optional) (default to 0)
    sortBy := "sortBy_example" // string | Sort order (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushToCardAPI.GetPushtocardsDisburse(context.Background()).Count(count).Fields(fields).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushToCardAPI.GetPushtocardsDisburse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushtocardsDisburse`: PushToCardDisburseListResponse
    fmt.Fprintf(os.Stdout, "Response from `PushToCardAPI.GetPushtocardsDisburse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPushtocardsDisburseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of push-to-card disbursements to retrieve | [default to 10]
 **fields** | **string** | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **startIndex** | **int32** | Start index | [default to 0]
 **sortBy** | **string** | Sort order | [default to &quot;-createdTime&quot;]

### Return type

[**PushToCardDisburseListResponse**](PushToCardDisburseListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushtocardsDisburseToken

> PushToCardDisbursementResponse GetPushtocardsDisburseToken(ctx, token).Fields(fields).Execute()

Returns a specific push-to-card disbursement

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
    token := "token_example" // string | Push-to-card disbursement token
    fields := "fields_example" // string | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushToCardAPI.GetPushtocardsDisburseToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushToCardAPI.GetPushtocardsDisburseToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushtocardsDisburseToken`: PushToCardDisbursementResponse
    fmt.Fprintf(os.Stdout, "Response from `PushToCardAPI.GetPushtocardsDisburseToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Push-to-card disbursement token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushtocardsDisburseTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**PushToCardDisbursementResponse**](PushToCardDisbursementResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushtocardsPaymentcard

> PushToCardListResponse GetPushtocardsPaymentcard(ctx).UserToken(userToken).Count(count).Fields(fields).StartIndex(startIndex).SortBy(sortBy).Execute()

Returns all push-to-card payment card details

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
    userToken := "userToken_example" // string | User token
    count := int32(56) // int32 | Number of push-to-card payment cards to retrieve (optional) (default to 10)
    fields := "fields_example" // string | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. (optional)
    startIndex := int32(56) // int32 | Start index (optional) (default to 0)
    sortBy := "sortBy_example" // string | Sort order (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushToCardAPI.GetPushtocardsPaymentcard(context.Background()).UserToken(userToken).Count(count).Fields(fields).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushToCardAPI.GetPushtocardsPaymentcard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushtocardsPaymentcard`: PushToCardListResponse
    fmt.Fprintf(os.Stdout, "Response from `PushToCardAPI.GetPushtocardsPaymentcard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPushtocardsPaymentcardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userToken** | **string** | User token | 
 **count** | **int32** | Number of push-to-card payment cards to retrieve | [default to 10]
 **fields** | **string** | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 
 **startIndex** | **int32** | Start index | [default to 0]
 **sortBy** | **string** | Sort order | [default to &quot;-createdTime&quot;]

### Return type

[**PushToCardListResponse**](PushToCardListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPushtocardsPaymentcardToken

> PushToCardResponse GetPushtocardsPaymentcardToken(ctx, token).Fields(fields).Execute()

Returns a specific paymentcard object

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
    token := "token_example" // string | Push-to-card token
    fields := "fields_example" // string | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushToCardAPI.GetPushtocardsPaymentcardToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushToCardAPI.GetPushtocardsPaymentcardToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushtocardsPaymentcardToken`: PushToCardResponse
    fmt.Fprintf(os.Stdout, "Response from `PushToCardAPI.GetPushtocardsPaymentcardToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Push-to-card token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushtocardsPaymentcardTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (e.g. field_1,field_2,..). Leave blank to return all fields. | 

### Return type

[**PushToCardResponse**](PushToCardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPushtocardsDisburse

> PushToCardDisbursementResponse PostPushtocardsDisburse(ctx).PushToCardDisburseRequest(pushToCardDisburseRequest).Execute()

Initiates a push-to-card money disbursement

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
    pushToCardDisburseRequest := *openapiclient.NewPushToCardDisburseRequest(float32(123), "CurrencyCode_example", "PaymentInstrumentToken_example") // PushToCardDisburseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushToCardAPI.PostPushtocardsDisburse(context.Background()).PushToCardDisburseRequest(pushToCardDisburseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushToCardAPI.PostPushtocardsDisburse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPushtocardsDisburse`: PushToCardDisbursementResponse
    fmt.Fprintf(os.Stdout, "Response from `PushToCardAPI.PostPushtocardsDisburse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPushtocardsDisburseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushToCardDisburseRequest** | [**PushToCardDisburseRequest**](PushToCardDisburseRequest.md) |  | 

### Return type

[**PushToCardDisbursementResponse**](PushToCardDisbursementResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPushtocardsPaymentcard

> PushToCardResponse PostPushtocardsPaymentcard(ctx).PushToCardRequest(pushToCardRequest).Execute()

Adds an external card to which funds will be pushed

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
    pushToCardRequest := *openapiclient.NewPushToCardRequest("Address1_example", "City_example", "Country_example", "Cvv_example", "ExpDate_example", "NameOnCard_example", "Pan_example", "PostalCode_example", "State_example", "UserToken_example") // PushToCardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushToCardAPI.PostPushtocardsPaymentcard(context.Background()).PushToCardRequest(pushToCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushToCardAPI.PostPushtocardsPaymentcard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPushtocardsPaymentcard`: PushToCardResponse
    fmt.Fprintf(os.Stdout, "Response from `PushToCardAPI.PostPushtocardsPaymentcard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPushtocardsPaymentcardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushToCardRequest** | [**PushToCardRequest**](PushToCardRequest.md) |  | 

### Return type

[**PushToCardResponse**](PushToCardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

