# \BusinessTransitionsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBusinesstransitionsBusinessBusinesstoken**](BusinessTransitionsAPI.md#GetBusinesstransitionsBusinessBusinesstoken) | **Get** /businesstransitions/business/{business_token} | List business transitions
[**GetBusinesstransitionsToken**](BusinessTransitionsAPI.md#GetBusinesstransitionsToken) | **Get** /businesstransitions/{token} | Retrieve business transition
[**PostBusinesstransitions**](BusinessTransitionsAPI.md#PostBusinesstransitions) | **Post** /businesstransitions | Create business transition



## GetBusinesstransitionsBusinessBusinesstoken

> BusinessTransitionListResponse GetBusinesstransitionsBusinessBusinesstoken(ctx, businessToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List business transitions



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
    businessToken := "businessToken_example" // string | Unique identifier of the business resource associated with the transitions to retrieve.
    count := int32(56) // int32 | Number of business transitions to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-id")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessTransitionsAPI.GetBusinesstransitionsBusinessBusinesstoken(context.Background(), businessToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessTransitionsAPI.GetBusinesstransitionsBusinessBusinesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinesstransitionsBusinessBusinesstoken`: BusinessTransitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessTransitionsAPI.GetBusinesstransitionsBusinessBusinesstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessToken** | **string** | Unique identifier of the business resource associated with the transitions to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinesstransitionsBusinessBusinesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of business transitions to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-id&quot;]

### Return type

[**BusinessTransitionListResponse**](BusinessTransitionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinesstransitionsToken

> BusinessTransitionResponse GetBusinesstransitionsToken(ctx, token).Fields(fields).Execute()

Retrieve business transition



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
    token := "token_example" // string | The unique identifier of the business transition you want to retrieve.
    fields := "fields_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessTransitionsAPI.GetBusinesstransitionsToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessTransitionsAPI.GetBusinesstransitionsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinesstransitionsToken`: BusinessTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessTransitionsAPI.GetBusinesstransitionsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the business transition you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinesstransitionsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | 

### Return type

[**BusinessTransitionResponse**](BusinessTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBusinesstransitions

> BusinessTransitionResponse PostBusinesstransitions(ctx).BusinessTransitionRequest(businessTransitionRequest).Execute()

Create business transition



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
    businessTransitionRequest := *openapiclient.NewBusinessTransitionRequest("BusinessToken_example", "Channel_example", "ReasonCode_example", "Status_example") // BusinessTransitionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessTransitionsAPI.PostBusinesstransitions(context.Background()).BusinessTransitionRequest(businessTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessTransitionsAPI.PostBusinesstransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBusinesstransitions`: BusinessTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessTransitionsAPI.PostBusinesstransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBusinesstransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessTransitionRequest** | [**BusinessTransitionRequest**](BusinessTransitionRequest.md) |  | 

### Return type

[**BusinessTransitionResponse**](BusinessTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

