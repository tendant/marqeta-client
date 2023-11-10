# \BulkCardOrdersAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBulkissuances**](BulkCardOrdersAPI.md#GetBulkissuances) | **Get** /bulkissuances | List bulk card orders
[**GetBulkissuancesToken**](BulkCardOrdersAPI.md#GetBulkissuancesToken) | **Get** /bulkissuances/{token} | Retrieve bulk card order
[**PostBulkissuances**](BulkCardOrdersAPI.md#PostBulkissuances) | **Post** /bulkissuances | Create bulk card order



## GetBulkissuances

> BulkCardOrderListResponse GetBulkissuances(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List bulk card orders



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
    count := int32(56) // int32 | Number of bulk card orders to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkCardOrdersAPI.GetBulkissuances(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkCardOrdersAPI.GetBulkissuances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBulkissuances`: BulkCardOrderListResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkCardOrdersAPI.GetBulkissuances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkissuancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of bulk card orders to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**BulkCardOrderListResponse**](BulkCardOrderListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkissuancesToken

> BulkIssuanceResponse GetBulkissuancesToken(ctx, token).Execute()

Retrieve bulk card order



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
    token := "token_example" // string | The unique identifier of the bulk card order to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkCardOrdersAPI.GetBulkissuancesToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkCardOrdersAPI.GetBulkissuancesToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBulkissuancesToken`: BulkIssuanceResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkCardOrdersAPI.GetBulkissuancesToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the bulk card order to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkissuancesTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkIssuanceResponse**](BulkIssuanceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBulkissuances

> BulkIssuanceResponse PostBulkissuances(ctx).BulkIssuanceRequest(bulkIssuanceRequest).Execute()

Create bulk card order



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
    bulkIssuanceRequest := *openapiclient.NewBulkIssuanceRequest(int32(123), "CardProductToken_example", *openapiclient.NewFulfillmentRequest(*openapiclient.NewCardPersonalization(*openapiclient.NewText(*openapiclient.NewTextValue()))), "Token_example") // BulkIssuanceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkCardOrdersAPI.PostBulkissuances(context.Background()).BulkIssuanceRequest(bulkIssuanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkCardOrdersAPI.PostBulkissuances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBulkissuances`: BulkIssuanceResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkCardOrdersAPI.PostBulkissuances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBulkissuancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkIssuanceRequest** | [**BulkIssuanceRequest**](BulkIssuanceRequest.md) |  | 

### Return type

[**BulkIssuanceResponse**](BulkIssuanceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

