# \CardTransitionsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCardtransitionsCardToken**](CardTransitionsAPI.md#GetCardtransitionsCardToken) | **Get** /cardtransitions/card/{token} | List transitions for card
[**GetCardtransitionsToken**](CardTransitionsAPI.md#GetCardtransitionsToken) | **Get** /cardtransitions/{token} | Retrieve card transition
[**PostCardtransitions**](CardTransitionsAPI.md#PostCardtransitions) | **Post** /cardtransitions | Create card transition



## GetCardtransitionsCardToken

> CardTransitionListResponse GetCardtransitionsCardToken(ctx, token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List transitions for card



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
    token := "token_example" // string | Unique identifier of the card.
    count := int32(56) // int32 | Number of card transitions to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransitionsAPI.GetCardtransitionsCardToken(context.Background(), token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransitionsAPI.GetCardtransitionsCardToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardtransitionsCardToken`: CardTransitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardTransitionsAPI.GetCardtransitionsCardToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardtransitionsCardTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of card transitions to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**CardTransitionListResponse**](CardTransitionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardtransitionsToken

> CardTransitionResponse GetCardtransitionsToken(ctx, token).Fields(fields).Execute()

Retrieve card transition



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
    token := "token_example" // string | Unique identifier of the card transition. Send a `GET` request to `/cardtransitions/card/{token}` to retrieve card transition tokens for a specific card.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransitionsAPI.GetCardtransitionsToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransitionsAPI.GetCardtransitionsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardtransitionsToken`: CardTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CardTransitionsAPI.GetCardtransitionsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card transition. Send a &#x60;GET&#x60; request to &#x60;/cardtransitions/card/{token}&#x60; to retrieve card transition tokens for a specific card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardtransitionsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**CardTransitionResponse**](CardTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCardtransitions

> CardTransitionResponse PostCardtransitions(ctx).CardTransitionRequest(cardTransitionRequest).Execute()

Create card transition



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
    cardTransitionRequest := *openapiclient.NewCardTransitionRequest("CardToken_example", "Channel_example", "State_example") // CardTransitionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransitionsAPI.PostCardtransitions(context.Background()).CardTransitionRequest(cardTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransitionsAPI.PostCardtransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCardtransitions`: CardTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CardTransitionsAPI.PostCardtransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCardtransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardTransitionRequest** | [**CardTransitionRequest**](CardTransitionRequest.md) |  | 

### Return type

[**CardTransitionResponse**](CardTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

