# \DirectDepositsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDirectdeposits**](DirectDepositsAPI.md#GetDirectdeposits) | **Get** /directdeposits | Retrieves a list of all direct deposit records for your program.
[**GetDirectdepositsAccountsUserorbusinesstoken**](DirectDepositsAPI.md#GetDirectdepositsAccountsUserorbusinesstoken) | **Get** /directdeposits/accounts/{user_or_business_token} | Returns an account and routing number which can be used for direct deposit
[**GetDirectdepositsToken**](DirectDepositsAPI.md#GetDirectdepositsToken) | **Get** /directdeposits/{token} | Returns a direct deposit entry
[**GetDirectdepositsTransitions**](DirectDepositsAPI.md#GetDirectdepositsTransitions) | **Get** /directdeposits/transitions | Returns a list of direct deposit transitions
[**GetDirectdepositsTransitionsToken**](DirectDepositsAPI.md#GetDirectdepositsTransitionsToken) | **Get** /directdeposits/transitions/{token} | Returns a direct deposit transition
[**PostDirectdepositsTransitions**](DirectDepositsAPI.md#PostDirectdepositsTransitions) | **Post** /directdeposits/transitions | Creates a direct deposit transition
[**PutDirectdepositsAccountsUserorbusinesstoken**](DirectDepositsAPI.md#PutDirectdepositsAccountsUserorbusinesstoken) | **Put** /directdeposits/accounts/{user_or_business_token} | Updates a specific direct deposit account



## GetDirectdeposits

> DirectDepositListResponse GetDirectdeposits(ctx).Count(count).StartIndex(startIndex).ReversedAfterGracePeriod(reversedAfterGracePeriod).UserToken(userToken).BusinessToken(businessToken).DirectDepositState(directDepositState).StartSettlementDate(startSettlementDate).EndSettlementDate(endSettlementDate).SortBy(sortBy).Execute()

Retrieves a list of all direct deposit records for your program.

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
    count := int32(56) // int32 | The number of direct deposit records to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Start index (optional) (default to 0)
    reversedAfterGracePeriod := true // bool | Reversed after grace period (optional)
    userToken := "userToken_example" // string | User token (optional)
    businessToken := "businessToken_example" // string | Business token (optional)
    directDepositState := "directDepositState_example" // string | Comma-delimited list of direct deposit states to display e.g. PENDING | REVERSED | APPLIED | REJECTED  (optional)
    startSettlementDate := "startSettlementDate_example" // string | Start Settlement Date (optional)
    endSettlementDate := "endSettlementDate_example" // string | End Settlement Date (optional)
    sortBy := "sortBy_example" // string | Sort order (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositsAPI.GetDirectdeposits(context.Background()).Count(count).StartIndex(startIndex).ReversedAfterGracePeriod(reversedAfterGracePeriod).UserToken(userToken).BusinessToken(businessToken).DirectDepositState(directDepositState).StartSettlementDate(startSettlementDate).EndSettlementDate(endSettlementDate).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositsAPI.GetDirectdeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectdeposits`: DirectDepositListResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositsAPI.GetDirectdeposits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectdepositsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of direct deposit records to retrieve. | [default to 5]
 **startIndex** | **int32** | Start index | [default to 0]
 **reversedAfterGracePeriod** | **bool** | Reversed after grace period | 
 **userToken** | **string** | User token | 
 **businessToken** | **string** | Business token | 
 **directDepositState** | **string** | Comma-delimited list of direct deposit states to display e.g. PENDING | REVERSED | APPLIED | REJECTED  | 
 **startSettlementDate** | **string** | Start Settlement Date | 
 **endSettlementDate** | **string** | End Settlement Date | 
 **sortBy** | **string** | Sort order | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**DirectDepositListResponse**](DirectDepositListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectdepositsAccountsUserorbusinesstoken

> DepositAccount GetDirectdepositsAccountsUserorbusinesstoken(ctx, userOrBusinessToken).Execute()

Returns an account and routing number which can be used for direct deposit

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
    userOrBusinessToken := "userOrBusinessToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositsAPI.GetDirectdepositsAccountsUserorbusinesstoken(context.Background(), userOrBusinessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositsAPI.GetDirectdepositsAccountsUserorbusinesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectdepositsAccountsUserorbusinesstoken`: DepositAccount
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositsAPI.GetDirectdepositsAccountsUserorbusinesstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userOrBusinessToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectdepositsAccountsUserorbusinesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DepositAccount**](DepositAccount.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectdepositsToken

> DepositDepositResponse GetDirectdepositsToken(ctx, token).Execute()

Returns a direct deposit entry

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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositsAPI.GetDirectdepositsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositsAPI.GetDirectdepositsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectdepositsToken`: DepositDepositResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositsAPI.GetDirectdepositsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectdepositsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DepositDepositResponse**](DepositDepositResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectdepositsTransitions

> DirectDepositTransitionListResponse GetDirectdepositsTransitions(ctx).Count(count).UserToken(userToken).BusinessToken(businessToken).DirectDepositToken(directDepositToken).StartIndex(startIndex).SortBy(sortBy).States(states).Execute()

Returns a list of direct deposit transitions

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
    count := int32(56) // int32 | Number of direct deposit transitions to retrieve (optional) (default to 5)
    userToken := "userToken_example" // string | User token (optional)
    businessToken := "businessToken_example" // string | Business token (optional)
    directDepositToken := "directDepositToken_example" // string | Direct deposit token (optional)
    startIndex := int32(56) // int32 | Start index (optional) (default to 0)
    sortBy := "sortBy_example" // string | Sort order (optional) (default to "-createdTime")
    states := "states_example" // string | Comma-delimited list of direct deposit states to display e.g. PENDING | REVERSED | APPLIED | REJECTED  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositsAPI.GetDirectdepositsTransitions(context.Background()).Count(count).UserToken(userToken).BusinessToken(businessToken).DirectDepositToken(directDepositToken).StartIndex(startIndex).SortBy(sortBy).States(states).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositsAPI.GetDirectdepositsTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectdepositsTransitions`: DirectDepositTransitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositsAPI.GetDirectdepositsTransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectdepositsTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of direct deposit transitions to retrieve | [default to 5]
 **userToken** | **string** | User token | 
 **businessToken** | **string** | Business token | 
 **directDepositToken** | **string** | Direct deposit token | 
 **startIndex** | **int32** | Start index | [default to 0]
 **sortBy** | **string** | Sort order | [default to &quot;-createdTime&quot;]
 **states** | **string** | Comma-delimited list of direct deposit states to display e.g. PENDING | REVERSED | APPLIED | REJECTED  | 

### Return type

[**DirectDepositTransitionListResponse**](DirectDepositTransitionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectdepositsTransitionsToken

> DirectDepositTransitionResponse GetDirectdepositsTransitionsToken(ctx, token).Execute()

Returns a direct deposit transition

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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositsAPI.GetDirectdepositsTransitionsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositsAPI.GetDirectdepositsTransitionsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectdepositsTransitionsToken`: DirectDepositTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositsAPI.GetDirectdepositsTransitionsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectdepositsTransitionsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectDepositTransitionResponse**](DirectDepositTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDirectdepositsTransitions

> DirectDepositTransitionResponse PostDirectdepositsTransitions(ctx).DirectDepositTransitionRequest(directDepositTransitionRequest).Execute()

Creates a direct deposit transition

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
    directDepositTransitionRequest := *openapiclient.NewDirectDepositTransitionRequest("Channel_example", "DirectDepositToken_example", "Reason_example", "State_example") // DirectDepositTransitionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositsAPI.PostDirectdepositsTransitions(context.Background()).DirectDepositTransitionRequest(directDepositTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositsAPI.PostDirectdepositsTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDirectdepositsTransitions`: DirectDepositTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositsAPI.PostDirectdepositsTransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDirectdepositsTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directDepositTransitionRequest** | [**DirectDepositTransitionRequest**](DirectDepositTransitionRequest.md) |  | 

### Return type

[**DirectDepositTransitionResponse**](DirectDepositTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDirectdepositsAccountsUserorbusinesstoken

> DepositAccount PutDirectdepositsAccountsUserorbusinesstoken(ctx, userOrBusinessToken).DepositAccountUpdateRequest(depositAccountUpdateRequest).Execute()

Updates a specific direct deposit account

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
    userOrBusinessToken := "userOrBusinessToken_example" // string | User or business token
    depositAccountUpdateRequest := *openapiclient.NewDepositAccountUpdateRequest() // DepositAccountUpdateRequest | Deposit account update request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositsAPI.PutDirectdepositsAccountsUserorbusinesstoken(context.Background(), userOrBusinessToken).DepositAccountUpdateRequest(depositAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositsAPI.PutDirectdepositsAccountsUserorbusinesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDirectdepositsAccountsUserorbusinesstoken`: DepositAccount
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositsAPI.PutDirectdepositsAccountsUserorbusinesstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userOrBusinessToken** | **string** | User or business token | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDirectdepositsAccountsUserorbusinesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depositAccountUpdateRequest** | [**DepositAccountUpdateRequest**](DepositAccountUpdateRequest.md) | Deposit account update request | 

### Return type

[**DepositAccount**](DepositAccount.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

