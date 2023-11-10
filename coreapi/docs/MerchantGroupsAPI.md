# \MerchantGroupsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMerchantGroup**](MerchantGroupsAPI.md#GetMerchantGroup) | **Get** /merchantgroups/{token} | Retrieve merchant group
[**GetMerchantGroups**](MerchantGroupsAPI.md#GetMerchantGroups) | **Get** /merchantgroups | List merchant groups
[**PostMerchantGroup**](MerchantGroupsAPI.md#PostMerchantGroup) | **Post** /merchantgroups | Create merchant group
[**PutMerchantGroupsToken**](MerchantGroupsAPI.md#PutMerchantGroupsToken) | **Put** /merchantgroups/{token} | Update merchant group



## GetMerchantGroup

> MerchantGroupResponse GetMerchantGroup(ctx, token).Execute()

Retrieve merchant group



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
    token := "token_example" // string | Unique identifier of the merchant group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantGroupsAPI.GetMerchantGroup(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantGroupsAPI.GetMerchantGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantGroup`: MerchantGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `MerchantGroupsAPI.GetMerchantGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the merchant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MerchantGroupResponse**](MerchantGroupResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantGroups

> MerchantGroupListResponse GetMerchantGroups(ctx).Mid(mid).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List merchant groups



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
    mid := "mid_example" // string | Returns all merchant groups that contain the specified merchant identifier. (optional)
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantGroupsAPI.GetMerchantGroups(context.Background()).Mid(mid).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantGroupsAPI.GetMerchantGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantGroups`: MerchantGroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `MerchantGroupsAPI.GetMerchantGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mid** | **string** | Returns all merchant groups that contain the specified merchant identifier. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 10]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**MerchantGroupListResponse**](MerchantGroupListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMerchantGroup

> MerchantGroupResponse PostMerchantGroup(ctx).MerchantGroupRequest(merchantGroupRequest).Execute()

Create merchant group



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
    merchantGroupRequest := *openapiclient.NewMerchantGroupRequest("Name_example") // MerchantGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantGroupsAPI.PostMerchantGroup(context.Background()).MerchantGroupRequest(merchantGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantGroupsAPI.PostMerchantGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMerchantGroup`: MerchantGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `MerchantGroupsAPI.PostMerchantGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMerchantGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantGroupRequest** | [**MerchantGroupRequest**](MerchantGroupRequest.md) |  | 

### Return type

[**MerchantGroupResponse**](MerchantGroupResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMerchantGroupsToken

> MerchantGroupResponse PutMerchantGroupsToken(ctx, token).MerchantGroupUpdateRequest(merchantGroupUpdateRequest).Execute()

Update merchant group



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
    token := "token_example" // string | Unique identifier of the merchant group.
    merchantGroupUpdateRequest := *openapiclient.NewMerchantGroupUpdateRequest() // MerchantGroupUpdateRequest | Merchant Group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantGroupsAPI.PutMerchantGroupsToken(context.Background(), token).MerchantGroupUpdateRequest(merchantGroupUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantGroupsAPI.PutMerchantGroupsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMerchantGroupsToken`: MerchantGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `MerchantGroupsAPI.PutMerchantGroupsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the merchant group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMerchantGroupsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantGroupUpdateRequest** | [**MerchantGroupUpdateRequest**](MerchantGroupUpdateRequest.md) | Merchant Group | 

### Return type

[**MerchantGroupResponse**](MerchantGroupResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

