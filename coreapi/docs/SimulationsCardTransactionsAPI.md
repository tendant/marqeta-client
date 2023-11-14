# \SimulationsCardTransactionsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountFundingAuthPlusCaptureEvent**](SimulationsCardTransactionsAPI.md#AccountFundingAuthPlusCaptureEvent) | **Post** /simulations/cardtransactions/account.funding.auth_plus_capture | Simulate AFT authorization and capture
[**AccountFundingAuthPlusCaptureReversalEvent**](SimulationsCardTransactionsAPI.md#AccountFundingAuthPlusCaptureReversalEvent) | **Post** /simulations/cardtransactions/account.funding.auth_plus_capture.reversal | Simulate AFT authorization and capture reversal
[**AccountFundingAuthorizationClearingEvent**](SimulationsCardTransactionsAPI.md#AccountFundingAuthorizationClearingEvent) | **Post** /simulations/cardtransactions/account.funding.authorization.clearing | Simulate AFT authorization clearing
[**AccountFundingAuthorizationEvent**](SimulationsCardTransactionsAPI.md#AccountFundingAuthorizationEvent) | **Post** /simulations/cardtransactions/account.funding.authorization | Simulate AFT authorization
[**AccountFundingAuthorizationReversalEvent**](SimulationsCardTransactionsAPI.md#AccountFundingAuthorizationReversalEvent) | **Post** /simulations/cardtransactions/account.funding.authorization.reversal | Simulate AFT authorization reversal
[**AuthorizationAdviceEvent**](SimulationsCardTransactionsAPI.md#AuthorizationAdviceEvent) | **Post** /simulations/cardtransactions/authorization.advice | Simulate authorization advice
[**AuthorizationAtmWithdrawalEvent**](SimulationsCardTransactionsAPI.md#AuthorizationAtmWithdrawalEvent) | **Post** /simulations/cardtransactions/authorization.atm.withdrawal | Simulate ATM withdrawal authorization
[**AuthorizationCashbackEvent**](SimulationsCardTransactionsAPI.md#AuthorizationCashbackEvent) | **Post** /simulations/cardtransactions/authorization.cashback | Simulate authorization cash back
[**AuthorizationClearingAtmWithdrawalEvent**](SimulationsCardTransactionsAPI.md#AuthorizationClearingAtmWithdrawalEvent) | **Post** /simulations/cardtransactions/authorization.clearing.atm.withdrawal | Simulate ATM withdrawal authorization clearing
[**AuthorizationClearingEvent**](SimulationsCardTransactionsAPI.md#AuthorizationClearingEvent) | **Post** /simulations/cardtransactions/authorization.clearing | Simulate authorization clearing
[**AuthorizationClearingQuasiCashEvent**](SimulationsCardTransactionsAPI.md#AuthorizationClearingQuasiCashEvent) | **Post** /simulations/cardtransactions/authorization.clearing.quasi.cash | Simulate quasi-cash authorization clearing
[**AuthorizationEvent**](SimulationsCardTransactionsAPI.md#AuthorizationEvent) | **Post** /simulations/cardtransactions/authorization | Simulate authorization
[**AuthorizationIncrementalEvent**](SimulationsCardTransactionsAPI.md#AuthorizationIncrementalEvent) | **Post** /simulations/cardtransactions/authorization.incremental | Simulate incremental authorization
[**AuthorizationQuasiCashEvent**](SimulationsCardTransactionsAPI.md#AuthorizationQuasiCashEvent) | **Post** /simulations/cardtransactions/authorization.quasi.cash | Simulate quasi-cash authorization
[**AuthorizationReversalEvent**](SimulationsCardTransactionsAPI.md#AuthorizationReversalEvent) | **Post** /simulations/cardtransactions/authorization.reversal | Simulate authorization reversal
[**RefundAuthorizationClearingEvent**](SimulationsCardTransactionsAPI.md#RefundAuthorizationClearingEvent) | **Post** /simulations/cardtransactions/refund.authorization.clearing | Simulate refund authorization clearing
[**RefundAuthorizationEvent**](SimulationsCardTransactionsAPI.md#RefundAuthorizationEvent) | **Post** /simulations/cardtransactions/refund.authorization | Simulate refund authorization
[**RefundAuthorizationReversalEvent**](SimulationsCardTransactionsAPI.md#RefundAuthorizationReversalEvent) | **Post** /simulations/cardtransactions/refund.authorization.reversal | Simulate refund authorization reversal
[**RefundEvent**](SimulationsCardTransactionsAPI.md#RefundEvent) | **Post** /simulations/cardtransactions/refund | Simulate refund



## AccountFundingAuthPlusCaptureEvent

> CardTransactionResponse AccountFundingAuthPlusCaptureEvent(ctx).AccountFundingAuthPlusCaptureEvent(accountFundingAuthPlusCaptureEvent).Execute()

Simulate AFT authorization and capture



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
    accountFundingAuthPlusCaptureEvent := *openapiclient.NewAccountFundingAuthPlusCaptureEvent("CardToken_example", float32(123)) // AccountFundingAuthPlusCaptureEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AccountFundingAuthPlusCaptureEvent(context.Background()).AccountFundingAuthPlusCaptureEvent(accountFundingAuthPlusCaptureEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AccountFundingAuthPlusCaptureEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountFundingAuthPlusCaptureEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AccountFundingAuthPlusCaptureEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountFundingAuthPlusCaptureEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountFundingAuthPlusCaptureEvent** | [**AccountFundingAuthPlusCaptureEvent**](AccountFundingAuthPlusCaptureEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountFundingAuthPlusCaptureReversalEvent

> CardTransactionResponse AccountFundingAuthPlusCaptureReversalEvent(ctx).AccountFundingAuthPlusCaptureReversalEvent(accountFundingAuthPlusCaptureReversalEvent).Execute()

Simulate AFT authorization and capture reversal



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
    accountFundingAuthPlusCaptureReversalEvent := *openapiclient.NewAccountFundingAuthPlusCaptureReversalEvent("PrecedingRelatedTransactionToken_example", float32(123)) // AccountFundingAuthPlusCaptureReversalEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AccountFundingAuthPlusCaptureReversalEvent(context.Background()).AccountFundingAuthPlusCaptureReversalEvent(accountFundingAuthPlusCaptureReversalEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AccountFundingAuthPlusCaptureReversalEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountFundingAuthPlusCaptureReversalEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AccountFundingAuthPlusCaptureReversalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountFundingAuthPlusCaptureReversalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountFundingAuthPlusCaptureReversalEvent** | [**AccountFundingAuthPlusCaptureReversalEvent**](AccountFundingAuthPlusCaptureReversalEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountFundingAuthorizationClearingEvent

> CardTransactionResponse AccountFundingAuthorizationClearingEvent(ctx).AccountFundingAuthorizationClearingEvent(accountFundingAuthorizationClearingEvent).Execute()

Simulate AFT authorization clearing



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
    accountFundingAuthorizationClearingEvent := *openapiclient.NewAccountFundingAuthorizationClearingEvent(float32(123)) // AccountFundingAuthorizationClearingEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AccountFundingAuthorizationClearingEvent(context.Background()).AccountFundingAuthorizationClearingEvent(accountFundingAuthorizationClearingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AccountFundingAuthorizationClearingEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountFundingAuthorizationClearingEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AccountFundingAuthorizationClearingEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountFundingAuthorizationClearingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountFundingAuthorizationClearingEvent** | [**AccountFundingAuthorizationClearingEvent**](AccountFundingAuthorizationClearingEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountFundingAuthorizationEvent

> CardTransactionResponse AccountFundingAuthorizationEvent(ctx).AccountFundingAuthorizationEvent(accountFundingAuthorizationEvent).Execute()

Simulate AFT authorization



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
    accountFundingAuthorizationEvent := *openapiclient.NewAccountFundingAuthorizationEvent("CardToken_example", float32(123)) // AccountFundingAuthorizationEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AccountFundingAuthorizationEvent(context.Background()).AccountFundingAuthorizationEvent(accountFundingAuthorizationEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AccountFundingAuthorizationEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountFundingAuthorizationEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AccountFundingAuthorizationEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountFundingAuthorizationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountFundingAuthorizationEvent** | [**AccountFundingAuthorizationEvent**](AccountFundingAuthorizationEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountFundingAuthorizationReversalEvent

> CardTransactionResponse AccountFundingAuthorizationReversalEvent(ctx).AccountFundingAuthorizationReversalEvent(accountFundingAuthorizationReversalEvent).Execute()

Simulate AFT authorization reversal



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
    accountFundingAuthorizationReversalEvent := *openapiclient.NewAccountFundingAuthorizationReversalEvent("PrecedingRelatedTransactionToken_example", float32(123)) // AccountFundingAuthorizationReversalEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AccountFundingAuthorizationReversalEvent(context.Background()).AccountFundingAuthorizationReversalEvent(accountFundingAuthorizationReversalEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AccountFundingAuthorizationReversalEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountFundingAuthorizationReversalEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AccountFundingAuthorizationReversalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountFundingAuthorizationReversalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountFundingAuthorizationReversalEvent** | [**AccountFundingAuthorizationReversalEvent**](AccountFundingAuthorizationReversalEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationAdviceEvent

> CardTransactionResponse AuthorizationAdviceEvent(ctx).AuthorizationAdviceEvent(authorizationAdviceEvent).Execute()

Simulate authorization advice



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
    authorizationAdviceEvent := *openapiclient.NewAuthorizationAdviceEvent("PrecedingRelatedTransactionToken_example", float32(123)) // AuthorizationAdviceEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationAdviceEvent(context.Background()).AuthorizationAdviceEvent(authorizationAdviceEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationAdviceEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationAdviceEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationAdviceEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationAdviceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationAdviceEvent** | [**AuthorizationAdviceEvent**](AuthorizationAdviceEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationAtmWithdrawalEvent

> CardTransactionResponse AuthorizationAtmWithdrawalEvent(ctx).AuthorizationAtmWithdrawalEvent(authorizationAtmWithdrawalEvent).Execute()

Simulate ATM withdrawal authorization



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
    authorizationAtmWithdrawalEvent := *openapiclient.NewAuthorizationAtmWithdrawalEvent("CardToken_example", float32(123)) // AuthorizationAtmWithdrawalEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationAtmWithdrawalEvent(context.Background()).AuthorizationAtmWithdrawalEvent(authorizationAtmWithdrawalEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationAtmWithdrawalEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationAtmWithdrawalEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationAtmWithdrawalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationAtmWithdrawalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationAtmWithdrawalEvent** | [**AuthorizationAtmWithdrawalEvent**](AuthorizationAtmWithdrawalEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationCashbackEvent

> CardTransactionResponse AuthorizationCashbackEvent(ctx).AuthorizationCashbackEvent(authorizationCashbackEvent).Execute()

Simulate authorization cash back



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
    authorizationCashbackEvent := *openapiclient.NewAuthorizationCashbackEvent("CardToken_example", float32(123), float32(123)) // AuthorizationCashbackEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationCashbackEvent(context.Background()).AuthorizationCashbackEvent(authorizationCashbackEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationCashbackEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationCashbackEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationCashbackEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationCashbackEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationCashbackEvent** | [**AuthorizationCashbackEvent**](AuthorizationCashbackEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationClearingAtmWithdrawalEvent

> CardTransactionResponse AuthorizationClearingAtmWithdrawalEvent(ctx).AuthorizationClearingAtmWithdrawalEvent(authorizationClearingAtmWithdrawalEvent).Execute()

Simulate ATM withdrawal authorization clearing



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
    authorizationClearingAtmWithdrawalEvent := *openapiclient.NewAuthorizationClearingAtmWithdrawalEvent(float32(123)) // AuthorizationClearingAtmWithdrawalEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationClearingAtmWithdrawalEvent(context.Background()).AuthorizationClearingAtmWithdrawalEvent(authorizationClearingAtmWithdrawalEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationClearingAtmWithdrawalEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationClearingAtmWithdrawalEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationClearingAtmWithdrawalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationClearingAtmWithdrawalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationClearingAtmWithdrawalEvent** | [**AuthorizationClearingAtmWithdrawalEvent**](AuthorizationClearingAtmWithdrawalEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationClearingEvent

> CardTransactionResponse AuthorizationClearingEvent(ctx).AuthorizationClearingEvent(authorizationClearingEvent).Execute()

Simulate authorization clearing



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
    authorizationClearingEvent := *openapiclient.NewAuthorizationClearingEvent(float32(123)) // AuthorizationClearingEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationClearingEvent(context.Background()).AuthorizationClearingEvent(authorizationClearingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationClearingEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationClearingEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationClearingEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationClearingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationClearingEvent** | [**AuthorizationClearingEvent**](AuthorizationClearingEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationClearingQuasiCashEvent

> CardTransactionResponse AuthorizationClearingQuasiCashEvent(ctx).AuthorizationClearingQuasiCashEvent(authorizationClearingQuasiCashEvent).Execute()

Simulate quasi-cash authorization clearing



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
    authorizationClearingQuasiCashEvent := *openapiclient.NewAuthorizationClearingQuasiCashEvent(float32(123)) // AuthorizationClearingQuasiCashEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationClearingQuasiCashEvent(context.Background()).AuthorizationClearingQuasiCashEvent(authorizationClearingQuasiCashEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationClearingQuasiCashEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationClearingQuasiCashEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationClearingQuasiCashEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationClearingQuasiCashEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationClearingQuasiCashEvent** | [**AuthorizationClearingQuasiCashEvent**](AuthorizationClearingQuasiCashEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationEvent

> CardTransactionResponse AuthorizationEvent(ctx).AuthorizationEvent(authorizationEvent).Execute()

Simulate authorization



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
    authorizationEvent := *openapiclient.NewAuthorizationEvent("CardToken_example", float32(123)) // AuthorizationEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationEvent(context.Background()).AuthorizationEvent(authorizationEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationEvent** | [**AuthorizationEvent**](AuthorizationEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationIncrementalEvent

> CardTransactionResponse AuthorizationIncrementalEvent(ctx).AuthorizationIncrementalEvent(authorizationIncrementalEvent).Execute()

Simulate incremental authorization



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
    authorizationIncrementalEvent := *openapiclient.NewAuthorizationIncrementalEvent("PrecedingRelatedTransactionToken_example", float32(123)) // AuthorizationIncrementalEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationIncrementalEvent(context.Background()).AuthorizationIncrementalEvent(authorizationIncrementalEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationIncrementalEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationIncrementalEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationIncrementalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationIncrementalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationIncrementalEvent** | [**AuthorizationIncrementalEvent**](AuthorizationIncrementalEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationQuasiCashEvent

> CardTransactionResponse AuthorizationQuasiCashEvent(ctx).AuthorizationQuasiCashEvent(authorizationQuasiCashEvent).Execute()

Simulate quasi-cash authorization



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
    authorizationQuasiCashEvent := *openapiclient.NewAuthorizationQuasiCashEvent("CardToken_example", float32(123)) // AuthorizationQuasiCashEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationQuasiCashEvent(context.Background()).AuthorizationQuasiCashEvent(authorizationQuasiCashEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationQuasiCashEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationQuasiCashEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationQuasiCashEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationQuasiCashEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationQuasiCashEvent** | [**AuthorizationQuasiCashEvent**](AuthorizationQuasiCashEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizationReversalEvent

> CardTransactionResponse AuthorizationReversalEvent(ctx).AuthorizationReversalEvent(authorizationReversalEvent).Execute()

Simulate authorization reversal



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
    authorizationReversalEvent := *openapiclient.NewAuthorizationReversalEvent("PrecedingRelatedTransactionToken_example", float32(123)) // AuthorizationReversalEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.AuthorizationReversalEvent(context.Background()).AuthorizationReversalEvent(authorizationReversalEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.AuthorizationReversalEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizationReversalEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.AuthorizationReversalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizationReversalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationReversalEvent** | [**AuthorizationReversalEvent**](AuthorizationReversalEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundAuthorizationClearingEvent

> CardTransactionResponse RefundAuthorizationClearingEvent(ctx).RefundAuthorizationClearingEvent(refundAuthorizationClearingEvent).Execute()

Simulate refund authorization clearing



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
    refundAuthorizationClearingEvent := *openapiclient.NewRefundAuthorizationClearingEvent(float32(123)) // RefundAuthorizationClearingEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.RefundAuthorizationClearingEvent(context.Background()).RefundAuthorizationClearingEvent(refundAuthorizationClearingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.RefundAuthorizationClearingEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundAuthorizationClearingEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.RefundAuthorizationClearingEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundAuthorizationClearingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundAuthorizationClearingEvent** | [**RefundAuthorizationClearingEvent**](RefundAuthorizationClearingEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundAuthorizationEvent

> CardTransactionResponse RefundAuthorizationEvent(ctx).RefundAuthorizationEvent(refundAuthorizationEvent).Execute()

Simulate refund authorization



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
    refundAuthorizationEvent := *openapiclient.NewRefundAuthorizationEvent("CardToken_example", float32(123)) // RefundAuthorizationEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.RefundAuthorizationEvent(context.Background()).RefundAuthorizationEvent(refundAuthorizationEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.RefundAuthorizationEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundAuthorizationEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.RefundAuthorizationEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundAuthorizationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundAuthorizationEvent** | [**RefundAuthorizationEvent**](RefundAuthorizationEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundAuthorizationReversalEvent

> CardTransactionResponse RefundAuthorizationReversalEvent(ctx).RefundAuthorizationReversalEvent(refundAuthorizationReversalEvent).Execute()

Simulate refund authorization reversal



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
    refundAuthorizationReversalEvent := *openapiclient.NewRefundAuthorizationReversalEvent("PrecedingRelatedTransactionToken_example", float32(123)) // RefundAuthorizationReversalEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.RefundAuthorizationReversalEvent(context.Background()).RefundAuthorizationReversalEvent(refundAuthorizationReversalEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.RefundAuthorizationReversalEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundAuthorizationReversalEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.RefundAuthorizationReversalEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundAuthorizationReversalEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundAuthorizationReversalEvent** | [**RefundAuthorizationReversalEvent**](RefundAuthorizationReversalEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundEvent

> CardTransactionResponse RefundEvent(ctx).RefundEvent(refundEvent).Execute()

Simulate refund



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
    refundEvent := *openapiclient.NewRefundEvent("CardToken_example", float32(123)) // RefundEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsCardTransactionsAPI.RefundEvent(context.Background()).RefundEvent(refundEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsCardTransactionsAPI.RefundEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundEvent`: CardTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsCardTransactionsAPI.RefundEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundEvent** | [**RefundEvent**](RefundEvent.md) |  | 

### Return type

[**CardTransactionResponse**](CardTransactionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

