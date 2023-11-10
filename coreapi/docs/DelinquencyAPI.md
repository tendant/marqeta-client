# \DelinquencyAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResendWebhookEvent**](DelinquencyAPI.md#ResendWebhookEvent) | **Post** /credit/webhooks/{event_type}/{resource_token} | Resend credit event notification
[**RetrieveDelinquencyState**](DelinquencyAPI.md#RetrieveDelinquencyState) | **Get** /credit/accounts/{account_token}/delinquencystate | Retrieve delinquency state
[**RetrieveDelinquencyTransition**](DelinquencyAPI.md#RetrieveDelinquencyTransition) | **Get** /credit/accounts/{account_token}/delinquencystate/transitions/{delinquency_transition_token} | Retrieve delinquency state transition
[**RetrieveDelinquencyTransitions**](DelinquencyAPI.md#RetrieveDelinquencyTransitions) | **Get** /credit/accounts/{account_token}/delinquencystate/transitions | List delinquency state transitions



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
    resp, r, err := apiClient.DelinquencyAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelinquencyAPI.ResendWebhookEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendWebhookEvent`: WebhookEventResendContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `DelinquencyAPI.ResendWebhookEvent`: %v\n", resp)
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


## RetrieveDelinquencyState

> DelinquencyStateResponse RetrieveDelinquencyState(ctx, accountToken).Execute()

Retrieve delinquency state



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
    accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve delinquency state details.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelinquencyAPI.RetrieveDelinquencyState(context.Background(), accountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelinquencyAPI.RetrieveDelinquencyState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDelinquencyState`: DelinquencyStateResponse
    fmt.Fprintf(os.Stdout, "Response from `DelinquencyAPI.RetrieveDelinquencyState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve delinquency state details.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDelinquencyStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DelinquencyStateResponse**](DelinquencyStateResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDelinquencyTransition

> DelinquencyTransitionResponse RetrieveDelinquencyTransition(ctx, accountToken, delinquencyTransitionToken).Execute()

Retrieve delinquency state transition



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
    accountToken := "accountToken_example" // string | Unique identifier of the credit account whose delinquency state transition you want to retrieve.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    delinquencyTransitionToken := "delinquencyTransitionToken_example" // string | Unique identifier of the delinquency state transition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelinquencyAPI.RetrieveDelinquencyTransition(context.Background(), accountToken, delinquencyTransitionToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelinquencyAPI.RetrieveDelinquencyTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDelinquencyTransition`: DelinquencyTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DelinquencyAPI.RetrieveDelinquencyTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account whose delinquency state transition you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**delinquencyTransitionToken** | **string** | Unique identifier of the delinquency state transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDelinquencyTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DelinquencyTransitionResponse**](DelinquencyTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDelinquencyTransitions

> DelinquencyTransitionsResponsePage RetrieveDelinquencyTransitions(ctx, accountToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List delinquency state transitions



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
    accountToken := "accountToken_example" // string | Unique identifier of the credit account whose delinquency state transitions you want to retrieve.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `impactTime`, and not by the field names appearing in response bodies such as `impact_time`. (optional) (default to "-impactTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelinquencyAPI.RetrieveDelinquencyTransitions(context.Background(), accountToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelinquencyAPI.RetrieveDelinquencyTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDelinquencyTransitions`: DelinquencyTransitionsResponsePage
    fmt.Fprintf(os.Stdout, "Response from `DelinquencyAPI.RetrieveDelinquencyTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account whose delinquency state transitions you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDelinquencyTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;impactTime&#x60;, and not by the field names appearing in response bodies such as &#x60;impact_time&#x60;. | [default to &quot;-impactTime&quot;]

### Return type

[**DelinquencyTransitionsResponsePage**](DelinquencyTransitionsResponsePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

