# \AutoReloadAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoreloads**](AutoReloadAPI.md#GetAutoreloads) | **Get** /autoreloads | List auto reloads
[**GetAutoreloadsToken**](AutoReloadAPI.md#GetAutoreloadsToken) | **Get** /autoreloads/{token} | Retrieve auto reload
[**PostAutoreloads**](AutoReloadAPI.md#PostAutoreloads) | **Post** /autoreloads | Create auto reload
[**PutAutoreloadsToken**](AutoReloadAPI.md#PutAutoreloadsToken) | **Put** /autoreloads/{token} | Update auto reload



## GetAutoreloads

> AutoReloadListResponse GetAutoreloads(ctx).CardProduct(cardProduct).UserToken(userToken).BusinessToken(businessToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List auto reloads



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
    cardProduct := "cardProduct_example" // string | Unique identifier of the card product whose auto reloads you want to retrieve. (optional)
    userToken := "userToken_example" // string | Unique identifier of the user whose auto reloads you want to retrieve. (optional)
    businessToken := "businessToken_example" // string | Unique identifier of the business whose auto reloads you want to retrieve. (optional)
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoReloadAPI.GetAutoreloads(context.Background()).CardProduct(cardProduct).UserToken(userToken).BusinessToken(businessToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoReloadAPI.GetAutoreloads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoreloads`: AutoReloadListResponse
    fmt.Fprintf(os.Stdout, "Response from `AutoReloadAPI.GetAutoreloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoreloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **string** | Unique identifier of the card product whose auto reloads you want to retrieve. | 
 **userToken** | **string** | Unique identifier of the user whose auto reloads you want to retrieve. | 
 **businessToken** | **string** | Unique identifier of the business whose auto reloads you want to retrieve. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 10]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**AutoReloadListResponse**](AutoReloadListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoreloadsToken

> AutoReloadResponseModel GetAutoreloadsToken(ctx, token).Fields(fields).Execute()

Retrieve auto reload



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
    token := "token_example" // string | Unique identifier of the auto reload.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoReloadAPI.GetAutoreloadsToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoReloadAPI.GetAutoreloadsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoreloadsToken`: AutoReloadResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AutoReloadAPI.GetAutoreloadsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the auto reload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoreloadsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**AutoReloadResponseModel**](AutoReloadResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutoreloads

> AutoReloadResponseModel PostAutoreloads(ctx).AutoReloadModel(autoReloadModel).Execute()

Create auto reload



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
    autoReloadModel := *openapiclient.NewAutoReloadModel("CurrencyCode_example", *openapiclient.NewOrderScope()) // AutoReloadModel | Auto reload object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoReloadAPI.PostAutoreloads(context.Background()).AutoReloadModel(autoReloadModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoReloadAPI.PostAutoreloads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAutoreloads`: AutoReloadResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AutoReloadAPI.PostAutoreloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAutoreloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoReloadModel** | [**AutoReloadModel**](AutoReloadModel.md) | Auto reload object | 

### Return type

[**AutoReloadResponseModel**](AutoReloadResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutoreloadsToken

> AutoReloadResponseModel PutAutoreloadsToken(ctx, token).AutoReloadUpdateModel(autoReloadUpdateModel).Execute()

Update auto reload



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
    token := "token_example" // string | Unique identifier of the auto reload.
    autoReloadUpdateModel := *openapiclient.NewAutoReloadUpdateModel() // AutoReloadUpdateModel | Auto reload object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoReloadAPI.PutAutoreloadsToken(context.Background(), token).AutoReloadUpdateModel(autoReloadUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoReloadAPI.PutAutoreloadsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAutoreloadsToken`: AutoReloadResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AutoReloadAPI.PutAutoreloadsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the auto reload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutoreloadsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoReloadUpdateModel** | [**AutoReloadUpdateModel**](AutoReloadUpdateModel.md) | Auto reload object | 

### Return type

[**AutoReloadResponseModel**](AutoReloadResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

