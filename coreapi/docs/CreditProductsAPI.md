# \CreditProductsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](CreditProductsAPI.md#CreateProduct) | **Post** /credit/products | Create credit product
[**LineageProducts**](CreditProductsAPI.md#LineageProducts) | **Get** /credit/products/{token}/lineage | Retrieve credit product lineage
[**ListProducts**](CreditProductsAPI.md#ListProducts) | **Get** /credit/products | List credit products
[**RetrieveProduct**](CreditProductsAPI.md#RetrieveProduct) | **Get** /credit/products/{token} | Retrieve credit product



## CreateProduct

> ProductResponse CreateProduct(ctx).ProductCreateReq(productCreateReq).Execute()

Create credit product



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
    productCreateReq := *openapiclient.NewProductCreateReq([]string{"CardProductTokens_example"}, openapiclient.ProductClassification("CONSUMER"), *openapiclient.NewProductConfig(int32(123), int32(123)), *openapiclient.NewProductCreditLine(float32(123), float32(123)), openapiclient.CurrencyCode("USD"), *openapiclient.NewInterestCalculation(*openapiclient.NewApplicationOfCredits(openapiclient.CycleType("BEGINNING_REVOLVING"), int32(123)), "DayCount_example", "GraceDaysApplication_example", []string{"InterestApplication_example"}, openapiclient.InterestOnGraceReactivationEnum("ACCRUE_FULL_CYCLE"), "Method_example", float32(123)), float32(123), float32(123), "Name_example", []openapiclient.PaymentAllocationOrderEnum{openapiclient.PaymentAllocationOrderEnum("INTEREST")}, openapiclient.ProductSubType("CREDIT_CARD"), openapiclient.ProductType("REVOLVING"), []openapiclient.BalanceType{openapiclient.BalanceType("PURCHASE")}) // ProductCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditProductsAPI.CreateProduct(context.Background()).ProductCreateReq(productCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditProductsAPI.CreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProduct`: ProductResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditProductsAPI.CreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCreateReq** | [**ProductCreateReq**](ProductCreateReq.md) |  | 

### Return type

[**ProductResponse**](ProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LineageProducts

> ProductsPage LineageProducts(ctx, token).Status(status).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

Retrieve credit product lineage



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
    token := "token_example" // string | The unique identifier of the credit product whose lineage you want to retrieve.  Send a `GET` request to `/credit/products` to retrieve existing credit product tokens.
    status := []openapiclient.ResourceStatus{openapiclient.ResourceStatus("DRAFT")} // []ResourceStatus | An array of statuses by which to filter credit products. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditProductsAPI.LineageProducts(context.Background(), token).Status(status).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditProductsAPI.LineageProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LineageProducts`: ProductsPage
    fmt.Fprintf(os.Stdout, "Response from `CreditProductsAPI.LineageProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the credit product whose lineage you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/products&#x60; to retrieve existing credit product tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLineageProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**[]ResourceStatus**](ResourceStatus.md) | An array of statuses by which to filter credit products. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**ProductsPage**](ProductsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ProductsPage ListProducts(ctx).Status(status).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List credit products



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
    status := []openapiclient.ResourceStatus{openapiclient.ResourceStatus("DRAFT")} // []ResourceStatus | An array of statuses by which to filter credit products. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditProductsAPI.ListProducts(context.Background()).Status(status).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditProductsAPI.ListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProducts`: ProductsPage
    fmt.Fprintf(os.Stdout, "Response from `CreditProductsAPI.ListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**[]ResourceStatus**](ResourceStatus.md) | An array of statuses by which to filter credit products. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**ProductsPage**](ProductsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveProduct

> ProductResponse RetrieveProduct(ctx, token).Execute()

Retrieve credit product



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
    token := "token_example" // string | The unique identifier of the credit product to retrieve.  Send a `GET` request to `/credit/products` to retrieve existing credit product tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditProductsAPI.RetrieveProduct(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditProductsAPI.RetrieveProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveProduct`: ProductResponse
    fmt.Fprintf(os.Stdout, "Response from `CreditProductsAPI.RetrieveProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the credit product to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/products&#x60; to retrieve existing credit product tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductResponse**](ProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

