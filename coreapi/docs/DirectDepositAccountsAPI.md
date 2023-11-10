# \DirectDepositAccountsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](DirectDepositAccountsAPI.md#CreateAccount) | **Post** /depositaccounts | Creates new direct deposit account for cardholder.
[**CreateTransition**](DirectDepositAccountsAPI.md#CreateTransition) | **Post** /depositaccounts/transitions | Creates new transition for a direct deposit account.
[**GetCDDInfo**](DirectDepositAccountsAPI.md#GetCDDInfo) | **Get** /depositaccounts/{token}/cdd | Get direct deposit account transition list for card holder.
[**GetDirectDepositAccount**](DirectDepositAccountsAPI.md#GetDirectDepositAccount) | **Get** /depositaccounts/{token} | Get direct deposit account.
[**GetDirectDepositAccountTransition**](DirectDepositAccountsAPI.md#GetDirectDepositAccountTransition) | **Get** /depositaccounts/transitions/{token} | Get direct deposit account transition.
[**GetTransitionList**](DirectDepositAccountsAPI.md#GetTransitionList) | **Get** /depositaccounts/{user_token}/transitions | Get direct deposit account transition list for card holder.
[**GetUserDirectDepositAccounts**](DirectDepositAccountsAPI.md#GetUserDirectDepositAccounts) | **Get** /depositaccounts/user/{token} | List all specific direct deposit accounts.
[**GetUserForDirectDepositAccount**](DirectDepositAccountsAPI.md#GetUserForDirectDepositAccount) | **Get** /depositaccounts/account/{account_number}/user | Get User for Plain Text Account Number
[**Update**](DirectDepositAccountsAPI.md#Update) | **Put** /depositaccounts/{token} | Update direct deposit account.
[**UpdateCDDInfo**](DirectDepositAccountsAPI.md#UpdateCDDInfo) | **Put** /depositaccounts/{token}/cdd/{cddtoken} | Update CDD answers for Direct Deposit Account



## CreateAccount

> DirectDepositAccountResponse CreateAccount(ctx).DirectDepositAccountRequest(directDepositAccountRequest).Execute()

Creates new direct deposit account for cardholder.

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
    directDepositAccountRequest := *openapiclient.NewDirectDepositAccountRequest() // DirectDepositAccountRequest | Create direct deposit account for cardholder

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.CreateAccount(context.Background()).DirectDepositAccountRequest(directDepositAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: DirectDepositAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directDepositAccountRequest** | [**DirectDepositAccountRequest**](DirectDepositAccountRequest.md) | Create direct deposit account for cardholder | 

### Return type

[**DirectDepositAccountResponse**](DirectDepositAccountResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransition

> DirectDepositAccountTransitionResponse CreateTransition(ctx).DirectDepositAccountTransitionRequest(directDepositAccountTransitionRequest).Execute()

Creates new transition for a direct deposit account.

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
    directDepositAccountTransitionRequest := *openapiclient.NewDirectDepositAccountTransitionRequest("AccountToken_example", "Channel_example") // DirectDepositAccountTransitionRequest | Create transition for direct deposit account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.CreateTransition(context.Background()).DirectDepositAccountTransitionRequest(directDepositAccountTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.CreateTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransition`: DirectDepositAccountTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.CreateTransition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directDepositAccountTransitionRequest** | [**DirectDepositAccountTransitionRequest**](DirectDepositAccountTransitionRequest.md) | Create transition for direct deposit account | 

### Return type

[**DirectDepositAccountTransitionResponse**](DirectDepositAccountTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCDDInfo

> CustomerDueDiligenceResponse GetCDDInfo(ctx, token).Execute()

Get direct deposit account transition list for card holder.

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
    token := "token_example" // string | Get CDD info for a specific DDA token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.GetCDDInfo(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.GetCDDInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCDDInfo`: CustomerDueDiligenceResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.GetCDDInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Get CDD info for a specific DDA token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCDDInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDueDiligenceResponse**](CustomerDueDiligenceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectDepositAccount

> DirectDepositAccountResponse GetDirectDepositAccount(ctx, token).Execute()

Get direct deposit account.

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
    token := "token_example" // string | Get specific direct deposit account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.GetDirectDepositAccount(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.GetDirectDepositAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectDepositAccount`: DirectDepositAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.GetDirectDepositAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Get specific direct deposit account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectDepositAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectDepositAccountResponse**](DirectDepositAccountResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectDepositAccountTransition

> DirectDepositAccountTransitionResponse GetDirectDepositAccountTransition(ctx, token).Execute()

Get direct deposit account transition.

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
    token := "token_example" // string | Get specific direct deposit account transition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.GetDirectDepositAccountTransition(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.GetDirectDepositAccountTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDirectDepositAccountTransition`: DirectDepositAccountTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.GetDirectDepositAccountTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Get specific direct deposit account transition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectDepositAccountTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectDepositAccountTransitionResponse**](DirectDepositAccountTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitionList

> DirectDepositAccountTransitionResponse GetTransitionList(ctx, userToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

Get direct deposit account transition list for card holder.

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
    userToken := "userToken_example" // string | Get direct deposit account transition list for user
    count := int32(56) // int32 | Number of users to retrieve (optional) (default to 5)
    startIndex := int32(56) // int32 | Start index (optional) (default to 0)
    sortBy := "sortBy_example" // string | Sort order (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.GetTransitionList(context.Background(), userToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.GetTransitionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitionList`: DirectDepositAccountTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.GetTransitionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userToken** | **string** | Get direct deposit account transition list for user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of users to retrieve | [default to 5]
 **startIndex** | **int32** | Start index | [default to 0]
 **sortBy** | **string** | Sort order | [default to &quot;-createdTime&quot;]

### Return type

[**DirectDepositAccountTransitionResponse**](DirectDepositAccountTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDirectDepositAccounts

> DirectDepositAccountListResponse GetUserDirectDepositAccounts(ctx, token).Count(count).StartIndex(startIndex).SortBy(sortBy).State(state).Execute()

List all specific direct deposit accounts.

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
    token := "token_example" // string | Get specific direct deposit account
    count := int32(56) // int32 | Number of users to retrieve (optional) (default to 5)
    startIndex := int32(56) // int32 | Start index (optional) (default to 0)
    sortBy := "sortBy_example" // string | Sort order (optional) (default to "-lastModifiedTime")
    state := "state_example" // string | Direct deposit account status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.GetUserDirectDepositAccounts(context.Background(), token).Count(count).StartIndex(startIndex).SortBy(sortBy).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.GetUserDirectDepositAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDirectDepositAccounts`: DirectDepositAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.GetUserDirectDepositAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Get specific direct deposit account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDirectDepositAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of users to retrieve | [default to 5]
 **startIndex** | **int32** | Start index | [default to 0]
 **sortBy** | **string** | Sort order | [default to &quot;-lastModifiedTime&quot;]
 **state** | **string** | Direct deposit account status | 

### Return type

[**DirectDepositAccountListResponse**](DirectDepositAccountListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserForDirectDepositAccount

> DirectDepositAccountResponse GetUserForDirectDepositAccount(ctx, accountNumber).Execute()

Get User for Plain Text Account Number

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
    accountNumber := "accountNumber_example" // string | Get user associated with direct deposit account number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.GetUserForDirectDepositAccount(context.Background(), accountNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.GetUserForDirectDepositAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserForDirectDepositAccount`: DirectDepositAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.GetUserForDirectDepositAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | Get user associated with direct deposit account number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserForDirectDepositAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DirectDepositAccountResponse**](DirectDepositAccountResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> DirectDepositAccountResponse Update(ctx, token).DepositAccountUpdateRequest(depositAccountUpdateRequest).Execute()

Update direct deposit account.

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
    depositAccountUpdateRequest := *openapiclient.NewDepositAccountUpdateRequest() // DepositAccountUpdateRequest | Update direct deposit account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.Update(context.Background(), token).DepositAccountUpdateRequest(depositAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: DirectDepositAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depositAccountUpdateRequest** | [**DepositAccountUpdateRequest**](DepositAccountUpdateRequest.md) | Update direct deposit account | 

### Return type

[**DirectDepositAccountResponse**](DirectDepositAccountResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCDDInfo

> CustomerDueDiligenceResponse UpdateCDDInfo(ctx, token, cddtoken).CustomerDueDiligenceUpdateResponse(customerDueDiligenceUpdateResponse).Execute()

Update CDD answers for Direct Deposit Account

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
    cddtoken := "cddtoken_example" // string | 
    customerDueDiligenceUpdateResponse := *openapiclient.NewCustomerDueDiligenceUpdateResponse() // CustomerDueDiligenceUpdateResponse | Update CDD answers

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectDepositAccountsAPI.UpdateCDDInfo(context.Background(), token, cddtoken).CustomerDueDiligenceUpdateResponse(customerDueDiligenceUpdateResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectDepositAccountsAPI.UpdateCDDInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCDDInfo`: CustomerDueDiligenceResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectDepositAccountsAPI.UpdateCDDInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 
**cddtoken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCDDInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customerDueDiligenceUpdateResponse** | [**CustomerDueDiligenceUpdateResponse**](CustomerDueDiligenceUpdateResponse.md) | Update CDD answers | 

### Return type

[**CustomerDueDiligenceResponse**](CustomerDueDiligenceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

