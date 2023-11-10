# \CardProductsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCardproducts**](CardProductsAPI.md#GetCardproducts) | **Get** /cardproducts | List card products
[**GetCardproductsToken**](CardProductsAPI.md#GetCardproductsToken) | **Get** /cardproducts/{token} | Retrieve card product
[**PostCardproducts**](CardProductsAPI.md#PostCardproducts) | **Post** /cardproducts | Create card product
[**PutCardproductsToken**](CardProductsAPI.md#PutCardproductsToken) | **Put** /cardproducts/{token} | Update card product



## GetCardproducts

> CardProductListResponse GetCardproducts(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List card products



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
    count := int32(56) // int32 | Number of resources to retrieve. Count can be between 1 - 10 items. (optional) (default to 5)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardProductsAPI.GetCardproducts(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardProductsAPI.GetCardproducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardproducts`: CardProductListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardProductsAPI.GetCardproducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCardproductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of resources to retrieve. Count can be between 1 - 10 items. | [default to 5]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**CardProductListResponse**](CardProductListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardproductsToken

> CardProductResponse GetCardproductsToken(ctx, token).Execute()

Retrieve card product



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
    token := "token_example" // string | Unique identifier of the card product to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardProductsAPI.GetCardproductsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardProductsAPI.GetCardproductsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardproductsToken`: CardProductResponse
    fmt.Fprintf(os.Stdout, "Response from `CardProductsAPI.GetCardproductsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card product to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardproductsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardProductResponse**](CardProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCardproducts

> CardProductResponse PostCardproducts(ctx).CardProductRequest(cardProductRequest).Execute()

Create card product



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
    cardProductRequest := *openapiclient.NewCardProductRequest("Name_example", time.Now()) // CardProductRequest | Card product object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardProductsAPI.PostCardproducts(context.Background()).CardProductRequest(cardProductRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardProductsAPI.PostCardproducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCardproducts`: CardProductResponse
    fmt.Fprintf(os.Stdout, "Response from `CardProductsAPI.PostCardproducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCardproductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProductRequest** | [**CardProductRequest**](CardProductRequest.md) | Card product object | 

### Return type

[**CardProductResponse**](CardProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCardproductsToken

> CardProductResponse PutCardproductsToken(ctx, token).CardProductUpdateModel(cardProductUpdateModel).Execute()

Update card product



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
    token := "token_example" // string | Unique identifier of the card product to update.
    cardProductUpdateModel := *openapiclient.NewCardProductUpdateModel() // CardProductUpdateModel | Card product object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardProductsAPI.PutCardproductsToken(context.Background(), token).CardProductUpdateModel(cardProductUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardProductsAPI.PutCardproductsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCardproductsToken`: CardProductResponse
    fmt.Fprintf(os.Stdout, "Response from `CardProductsAPI.PutCardproductsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card product to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCardproductsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardProductUpdateModel** | [**CardProductUpdateModel**](CardProductUpdateModel.md) | Card product object | 

### Return type

[**CardProductResponse**](CardProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

