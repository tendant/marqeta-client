# \GPAOrdersAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGpaordersToken**](GPAOrdersAPI.md#GetGpaordersToken) | **Get** /gpaorders/{token} | Retrieve GPA order
[**GetGpaordersUnloads**](GPAOrdersAPI.md#GetGpaordersUnloads) | **Get** /gpaorders/unloads | List GPA unloads
[**GetGpaordersUnloadsUnloadtoken**](GPAOrdersAPI.md#GetGpaordersUnloadsUnloadtoken) | **Get** /gpaorders/unloads/{unload_token} | Retrieve GPA unload
[**PostGpaorders**](GPAOrdersAPI.md#PostGpaorders) | **Post** /gpaorders | Create GPA order
[**PostGpaordersUnloads**](GPAOrdersAPI.md#PostGpaordersUnloads) | **Post** /gpaorders/unloads | Create GPA unload



## GetGpaordersToken

> GpaResponse GetGpaordersToken(ctx, token).Execute()

Retrieve GPA order



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
    token := "token_example" // string | Unique identifier of the GPA order.  Send a `GET` request to `/transactions?type=gpa.credit` to retrieve GPA order tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPAOrdersAPI.GetGpaordersToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPAOrdersAPI.GetGpaordersToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGpaordersToken`: GpaResponse
    fmt.Fprintf(os.Stdout, "Response from `GPAOrdersAPI.GetGpaordersToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the GPA order.  Send a &#x60;GET&#x60; request to &#x60;/transactions?type&#x3D;gpa.credit&#x60; to retrieve GPA order tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGpaordersTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GpaResponse**](GpaResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGpaordersUnloads

> GPAUnloadListResponse GetGpaordersUnloads(ctx).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).UserToken(userToken).BusinessToken(businessToken).OriginalOrderToken(originalOrderToken).Execute()

List GPA unloads



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
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")
    userToken := "userToken_example" // string | Unique identifier of the user resource.  Send a `GET` request to `/users` to retrieve user tokens. (optional)
    businessToken := "businessToken_example" // string | Unique identifier of the business resource.  Send a `GET` request to `/businesses` to retrieve business tokens. (optional)
    originalOrderToken := "originalOrderToken_example" // string | Unique identifier of the original GPA order.  Send a `GET` request to `/transactions?type=gpa.credit` to retrieve GPA order tokens. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPAOrdersAPI.GetGpaordersUnloads(context.Background()).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).UserToken(userToken).BusinessToken(businessToken).OriginalOrderToken(originalOrderToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPAOrdersAPI.GetGpaordersUnloads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGpaordersUnloads`: GPAUnloadListResponse
    fmt.Fprintf(os.Stdout, "Response from `GPAOrdersAPI.GetGpaordersUnloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGpaordersUnloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]
 **userToken** | **string** | Unique identifier of the user resource.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | 
 **businessToken** | **string** | Unique identifier of the business resource.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | 
 **originalOrderToken** | **string** | Unique identifier of the original GPA order.  Send a &#x60;GET&#x60; request to &#x60;/transactions?type&#x3D;gpa.credit&#x60; to retrieve GPA order tokens. | 

### Return type

[**GPAUnloadListResponse**](GPAUnloadListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGpaordersUnloadsUnloadtoken

> GpaReturns GetGpaordersUnloadsUnloadtoken(ctx, unloadToken).Execute()

Retrieve GPA unload



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
    unloadToken := "unloadToken_example" // string | Unique identifier of the GPA unload.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPAOrdersAPI.GetGpaordersUnloadsUnloadtoken(context.Background(), unloadToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPAOrdersAPI.GetGpaordersUnloadsUnloadtoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGpaordersUnloadsUnloadtoken`: GpaReturns
    fmt.Fprintf(os.Stdout, "Response from `GPAOrdersAPI.GetGpaordersUnloadsUnloadtoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unloadToken** | **string** | Unique identifier of the GPA unload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGpaordersUnloadsUnloadtokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GpaReturns**](GpaReturns.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGpaorders

> GpaResponse PostGpaorders(ctx).GpaRequest(gpaRequest).Execute()

Create GPA order



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
    gpaRequest := *openapiclient.NewGpaRequest(float32(123), "CurrencyCode_example", "FundingSourceToken_example") // GpaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPAOrdersAPI.PostGpaorders(context.Background()).GpaRequest(gpaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPAOrdersAPI.PostGpaorders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGpaorders`: GpaResponse
    fmt.Fprintf(os.Stdout, "Response from `GPAOrdersAPI.PostGpaorders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGpaordersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpaRequest** | [**GpaRequest**](GpaRequest.md) |  | 

### Return type

[**GpaResponse**](GpaResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGpaordersUnloads

> GpaReturns PostGpaordersUnloads(ctx).UnloadRequestModel(unloadRequestModel).Execute()

Create GPA unload



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
    unloadRequestModel := *openapiclient.NewUnloadRequestModel(float32(123), "OriginalOrderToken_example") // UnloadRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPAOrdersAPI.PostGpaordersUnloads(context.Background()).UnloadRequestModel(unloadRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPAOrdersAPI.PostGpaordersUnloads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGpaordersUnloads`: GpaReturns
    fmt.Fprintf(os.Stdout, "Response from `GPAOrdersAPI.PostGpaordersUnloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGpaordersUnloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unloadRequestModel** | [**UnloadRequestModel**](UnloadRequestModel.md) |  | 

### Return type

[**GpaReturns**](GpaReturns.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

