# \AccountCardsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCardForAccount**](AccountCardsAPI.md#CreateCardForAccount) | **Post** /credit/accounts/{account_token}/cards | Create account card
[**GetCardByAccount**](AccountCardsAPI.md#GetCardByAccount) | **Get** /credit/accounts/{account_token}/cards/{card_token} | Retrieve account card
[**GetCardsByAccount**](AccountCardsAPI.md#GetCardsByAccount) | **Get** /credit/accounts/{account_token}/cards | List account cards



## CreateCardForAccount

> CardResponse CreateCardForAccount(ctx, accountToken).CardCreateReq(cardCreateReq).Execute()

Create account card



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to create a credit card.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    cardCreateReq := *openapiclient.NewCardCreateReq("CardProductToken_example", "UserToken_example") // CardCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountCardsAPI.CreateCardForAccount(context.Background(), accountToken).CardCreateReq(cardCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountCardsAPI.CreateCardForAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCardForAccount`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountCardsAPI.CreateCardForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to create a credit card.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCardForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardCreateReq** | [**CardCreateReq**](CardCreateReq.md) |  | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardByAccount

> CardResponse GetCardByAccount(ctx, accountToken, cardToken).Execute()

Retrieve account card



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to retrieve a credit card.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    cardToken := "cardToken_example" // string | The unique identifier of the credit card to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/cards` to retrieve existing credit card tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountCardsAPI.GetCardByAccount(context.Background(), accountToken, cardToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountCardsAPI.GetCardByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardByAccount`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountCardsAPI.GetCardByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to retrieve a credit card.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**cardToken** | **string** | The unique identifier of the credit card to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/cards&#x60; to retrieve existing credit card tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardsByAccount

> AccountCardsPage GetCardsByAccount(ctx, accountToken).Status(status).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List account cards



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to retrieve credit cards.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    status := "status_example" // string | The status of the credit cards to return. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountCardsAPI.GetCardsByAccount(context.Background(), accountToken).Status(status).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountCardsAPI.GetCardsByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardsByAccount`: AccountCardsPage
    fmt.Fprintf(os.Stdout, "Response from `AccountCardsAPI.GetCardsByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to retrieve credit cards.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | The status of the credit cards to return. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**AccountCardsPage**](AccountCardsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

