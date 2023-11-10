# \PaymentsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayment**](PaymentsAPI.md#CreatePayment) | **Post** /credit/accounts/{account_token}/payments | Create account payment
[**ListPayments**](PaymentsAPI.md#ListPayments) | **Get** /credit/accounts/{account_token}/payments | List account payments
[**ReleasePaymentHold**](PaymentsAPI.md#ReleasePaymentHold) | **Post** /credit/accounts/{account_token}/payments/{payment_token}/releasehold | Release payment hold
[**ResendWebhookEvent**](PaymentsAPI.md#ResendWebhookEvent) | **Post** /credit/webhooks/{event_type}/{resource_token} | Resend credit event notification
[**RetrievePayment**](PaymentsAPI.md#RetrievePayment) | **Get** /credit/accounts/{account_token}/payments/{payment_token} | Retrieve account payment
[**TransitionPayment**](PaymentsAPI.md#TransitionPayment) | **Post** /credit/accounts/{account_token}/payments/{payment_token}/transitions | Transition account payment status



## CreatePayment

> PaymentDetailResponse CreatePayment(ctx, accountToken).PaymentCreateReq(paymentCreateReq).Execute()

Create account payment



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to create a payment.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentCreateReq := *openapiclient.NewPaymentCreateReq(float32(123), openapiclient.CurrencyCode("USD"), "Description_example", "Method_example") // PaymentCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.CreatePayment(context.Background(), accountToken).PaymentCreateReq(paymentCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.CreatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayment`: PaymentDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.CreatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to create a payment.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentCreateReq** | [**PaymentCreateReq**](PaymentCreateReq.md) |  | 

### Return type

[**PaymentDetailResponse**](PaymentDetailResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayments

> PaymentsPage ListPayments(ctx, accountToken).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Statuses(statuses).Execute()

List account payments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to retrieve payments.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    startDate := time.Now() // string | Beginning of the date range of the payments to return. (optional)
    endDate := time.Now() // string | End of the date range of the payments to return. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")
    statuses := []openapiclient.PaymentStatus{openapiclient.PaymentStatus("INITIATED")} // []PaymentStatus | An array of statuses by which to filter payments. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.ListPayments(context.Background(), accountToken).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ListPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayments`: PaymentsPage
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ListPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to retrieve payments.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Beginning of the date range of the payments to return. | 
 **endDate** | **string** | End of the date range of the payments to return. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]
 **statuses** | [**[]PaymentStatus**](PaymentStatus.md) | An array of statuses by which to filter payments. | 

### Return type

[**PaymentsPage**](PaymentsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleasePaymentHold

> PaymentDetailResponse ReleasePaymentHold(ctx, accountToken, paymentToken).Execute()

Release payment hold



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account on which a payment hold is being released.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentToken := "paymentToken_example" // string | The unique identifier of the payment currently on hold.  Send a `GET` request to `/credit/accounts/{account_token}/payments` to retrieve existing payment tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.ReleasePaymentHold(context.Background(), accountToken, paymentToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ReleasePaymentHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleasePaymentHold`: PaymentDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ReleasePaymentHold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account on which a payment hold is being released.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**paymentToken** | **string** | The unique identifier of the payment currently on hold.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/payments&#x60; to retrieve existing payment tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleasePaymentHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentDetailResponse**](PaymentDetailResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendWebhookEvent

> WebhookEventResendContainerResponse ResendWebhookEvent(ctx, eventType, resourceToken).Execute()

Resend credit event notification



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
    eventType := "eventType_example" // string | Specifies the type of event you want to resend.
    resourceToken := "resourceToken_example" // string | The unique identifier of the resource for which you want to resend a notification.  Send a `GET` request to `/credit/accounts/{account_token}/journalentries` to retrieve existing journal entry tokens.  Send a `GET` request to `/credit/accounts/{account_token}/ledgerentries` to retrieve existing ledger entry tokens.  Send a `GET` request to `/accounts/{account_token}/accounttransitions` to retrieve existing account transition tokens.  Send a `GET` request to `/credit/accounts/{account_token}/payments/{payment_token}` to retrieve existing payment transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements` to retrieve existing statement summary tokens.  Send a `GET` request to `/accounts/{account_token}/delinquencystate/transitions` to retrieve existing delinquency state transition tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ResendWebhookEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendWebhookEvent`: WebhookEventResendContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ResendWebhookEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventType** | **string** | Specifies the type of event you want to resend. | 
**resourceToken** | **string** | The unique identifier of the resource for which you want to resend a notification.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/journalentries&#x60; to retrieve existing journal entry tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/ledgerentries&#x60; to retrieve existing ledger entry tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/accounttransitions&#x60; to retrieve existing account transition tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/payments/{payment_token}&#x60; to retrieve existing payment transition tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/statements&#x60; to retrieve existing statement summary tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/delinquencystate/transitions&#x60; to retrieve existing delinquency state transition tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebhookEventResendContainerResponse**](WebhookEventResendContainerResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePayment

> PaymentDetailResponse RetrievePayment(ctx, accountToken, paymentToken).Execute()

Retrieve account payment



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to retrieve a payment.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentToken := "paymentToken_example" // string | The unique identifier of the payment to retrieve.  Send a `GET` request to `/credit/accounts/{token}/payments` to retrieve existing payment tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.RetrievePayment(context.Background(), accountToken, paymentToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.RetrievePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePayment`: PaymentDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.RetrievePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to retrieve a payment.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**paymentToken** | **string** | The unique identifier of the payment to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/payments&#x60; to retrieve existing payment tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentDetailResponse**](PaymentDetailResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionPayment

> PaymentTransitionResponse TransitionPayment(ctx, accountToken, paymentToken).PaymentTransitionReq(paymentTransitionReq).Execute()

Transition account payment status



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to transition a payment status.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentToken := "paymentToken_example" // string | The unique identifier of the payment whose status you want to transition.  Send a `GET` request to `/credit/accounts/{account_token}/payments` endpoint to retrieve existing payment tokens for a given account.
    paymentTransitionReq := *openapiclient.NewPaymentTransitionReq(openapiclient.PaymentStatus("INITIATED")) // PaymentTransitionReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsAPI.TransitionPayment(context.Background(), accountToken, paymentToken).PaymentTransitionReq(paymentTransitionReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.TransitionPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransitionPayment`: PaymentTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.TransitionPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to transition a payment status.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**paymentToken** | **string** | The unique identifier of the payment whose status you want to transition.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/payments&#x60; endpoint to retrieve existing payment tokens for a given account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **paymentTransitionReq** | [**PaymentTransitionReq**](PaymentTransitionReq.md) |  | 

### Return type

[**PaymentTransitionResponse**](PaymentTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

