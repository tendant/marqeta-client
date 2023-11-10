# \PaymentSchedulesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentSchedule**](PaymentSchedulesAPI.md#CreatePaymentSchedule) | **Post** /credit/accounts/{account_token}/paymentschedules | Create payment schedule
[**CreatePaymentScheduleTransition**](PaymentSchedulesAPI.md#CreatePaymentScheduleTransition) | **Post** /credit/accounts/{account_token}/paymentschedules/{payment_schedule_token}/transitions | Create payment schedule transition
[**RetrievePaymentSchedule**](PaymentSchedulesAPI.md#RetrievePaymentSchedule) | **Get** /credit/accounts/{account_token}/paymentschedules/{payment_schedule_token} | Retrieve payment schedule
[**RetrievePaymentScheduleTransition**](PaymentSchedulesAPI.md#RetrievePaymentScheduleTransition) | **Get** /credit/accounts/{account_token}/paymentschedules/{payment_schedule_token}/transitions/{token} | Retrieve payment schedule transition
[**RetrievePaymentScheduleTransitions**](PaymentSchedulesAPI.md#RetrievePaymentScheduleTransitions) | **Get** /credit/accounts/{account_token}/paymentschedules/{payment_schedule_token}/transitions | Retrieve payment schedule transitions
[**RetrievePaymentSchedules**](PaymentSchedulesAPI.md#RetrievePaymentSchedules) | **Get** /credit/accounts/{account_token}/paymentschedules | List payment schedules



## CreatePaymentSchedule

> PaymentScheduleResponse CreatePaymentSchedule(ctx, accountToken).PaymentScheduleCreateReq(paymentScheduleCreateReq).Execute()

Create payment schedule



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to create a payment schedule.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentScheduleCreateReq := *openapiclient.NewPaymentScheduleCreateReq(openapiclient.PaymentScheduleAmountCategory("FIXED"), openapiclient.CurrencyCode("USD"), openapiclient.PaymentScheduleFrequency("ONCE"), "PaymentSourceToken_example") // PaymentScheduleCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesAPI.CreatePaymentSchedule(context.Background(), accountToken).PaymentScheduleCreateReq(paymentScheduleCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.CreatePaymentSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentSchedule`: PaymentScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.CreatePaymentSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to create a payment schedule.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentScheduleCreateReq** | [**PaymentScheduleCreateReq**](PaymentScheduleCreateReq.md) |  | 

### Return type

[**PaymentScheduleResponse**](PaymentScheduleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentScheduleTransition

> PaymentScheduleTransitionResponse CreatePaymentScheduleTransition(ctx, accountToken, paymentScheduleToken).PaymentScheduleTransitionCreateReq(paymentScheduleTransitionCreateReq).Execute()

Create payment schedule transition



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account on which to transition a payment schedule.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentScheduleToken := "paymentScheduleToken_example" // string | The unique identifier of the payment schedule whose status is to transition.  Send a `GET` request to `/credit/accounts/{account_token}/paymentschedules` to retrieve existing payment schedule tokens.
    paymentScheduleTransitionCreateReq := *openapiclient.NewPaymentScheduleTransitionCreateReq(openapiclient.PaymentScheduleStatus("ACTIVE")) // PaymentScheduleTransitionCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesAPI.CreatePaymentScheduleTransition(context.Background(), accountToken, paymentScheduleToken).PaymentScheduleTransitionCreateReq(paymentScheduleTransitionCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.CreatePaymentScheduleTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentScheduleTransition`: PaymentScheduleTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.CreatePaymentScheduleTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account on which to transition a payment schedule.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**paymentScheduleToken** | **string** | The unique identifier of the payment schedule whose status is to transition.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/paymentschedules&#x60; to retrieve existing payment schedule tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentScheduleTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **paymentScheduleTransitionCreateReq** | [**PaymentScheduleTransitionCreateReq**](PaymentScheduleTransitionCreateReq.md) |  | 

### Return type

[**PaymentScheduleTransitionResponse**](PaymentScheduleTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePaymentSchedule

> PaymentScheduleResponse RetrievePaymentSchedule(ctx, accountToken, paymentScheduleToken).Execute()

Retrieve payment schedule



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve a payment schedule.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentScheduleToken := "paymentScheduleToken_example" // string | The unique identifier of the payment schedule that you want to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/paymentschedules` to retrieve existing payment schedule tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesAPI.RetrievePaymentSchedule(context.Background(), accountToken, paymentScheduleToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.RetrievePaymentSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePaymentSchedule`: PaymentScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.RetrievePaymentSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve a payment schedule.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**paymentScheduleToken** | **string** | The unique identifier of the payment schedule that you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/paymentschedules&#x60; to retrieve existing payment schedule tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentScheduleResponse**](PaymentScheduleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePaymentScheduleTransition

> PaymentScheduleTransitionResponse RetrievePaymentScheduleTransition(ctx, accountToken, paymentScheduleToken, token).Execute()

Retrieve payment schedule transition



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve a payment schedule transition.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentScheduleToken := "paymentScheduleToken_example" // string | The unique identifier of the payment schedule you want to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/paymentschedules` to retrieve existing payment schedule tokens.
    token := "token_example" // string | The unique identifier of the payment schedule transition you want to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/paymentschedules/{payment_schedule_token}/transitions` to retrieve existing payment schedule transition tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesAPI.RetrievePaymentScheduleTransition(context.Background(), accountToken, paymentScheduleToken, token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.RetrievePaymentScheduleTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePaymentScheduleTransition`: PaymentScheduleTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.RetrievePaymentScheduleTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve a payment schedule transition.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**paymentScheduleToken** | **string** | The unique identifier of the payment schedule you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/paymentschedules&#x60; to retrieve existing payment schedule tokens. | 
**token** | **string** | The unique identifier of the payment schedule transition you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/paymentschedules/{payment_schedule_token}/transitions&#x60; to retrieve existing payment schedule transition tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentScheduleTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaymentScheduleTransitionResponse**](PaymentScheduleTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePaymentScheduleTransitions

> PaymentScheduleTransitionPage RetrievePaymentScheduleTransitions(ctx, accountToken, paymentScheduleToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

Retrieve payment schedule transitions



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve payment schedule transitions.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    paymentScheduleToken := "paymentScheduleToken_example" // string | The unique identifier of the payment schedule for which you want to retrieve transitions.  Send a `GET` request to `/credit/accounts/{account_token}/paymentschedules` to retrieve existing payment schedule tokens.
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `createdTime`, and not by the field names appearing in response bodies such as `created_time`. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesAPI.RetrievePaymentScheduleTransitions(context.Background(), accountToken, paymentScheduleToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.RetrievePaymentScheduleTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePaymentScheduleTransitions`: PaymentScheduleTransitionPage
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.RetrievePaymentScheduleTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve payment schedule transitions.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**paymentScheduleToken** | **string** | The unique identifier of the payment schedule for which you want to retrieve transitions.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/paymentschedules&#x60; to retrieve existing payment schedule tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentScheduleTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]

### Return type

[**PaymentScheduleTransitionPage**](PaymentScheduleTransitionPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePaymentSchedules

> PaymentSchedulePage RetrievePaymentSchedules(ctx, accountToken).Statuses(statuses).Frequency(frequency).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List payment schedules



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve payment schedules.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    statuses := []openapiclient.PaymentScheduleStatus{openapiclient.PaymentScheduleStatus("ACTIVE")} // []PaymentScheduleStatus | Status of the payment schedules to retrieve. (optional)
    frequency := []openapiclient.PaymentScheduleFrequency{openapiclient.PaymentScheduleFrequency("ONCE")} // []PaymentScheduleFrequency | Frequency of the payment schedules to retrieve. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesAPI.RetrievePaymentSchedules(context.Background(), accountToken).Statuses(statuses).Frequency(frequency).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesAPI.RetrievePaymentSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePaymentSchedules`: PaymentSchedulePage
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesAPI.RetrievePaymentSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve payment schedules.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statuses** | [**[]PaymentScheduleStatus**](PaymentScheduleStatus.md) | Status of the payment schedules to retrieve. | 
 **frequency** | [**[]PaymentScheduleFrequency**](PaymentScheduleFrequency.md) | Frequency of the payment schedules to retrieve. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**PaymentSchedulePage**](PaymentSchedulePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

