# \RealTimeFeeGroupsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRealtimefeegroups**](RealTimeFeeGroupsAPI.md#GetRealtimefeegroups) | **Get** /realtimefeegroups | List real-time fee groups
[**GetRealtimefeegroupsToken**](RealTimeFeeGroupsAPI.md#GetRealtimefeegroupsToken) | **Get** /realtimefeegroups/{token} | Retrieve real-time fee group
[**PostRealtimefeegroups**](RealTimeFeeGroupsAPI.md#PostRealtimefeegroups) | **Post** /realtimefeegroups | Create real-time fee group
[**PutRealtimefeegroupsToken**](RealTimeFeeGroupsAPI.md#PutRealtimefeegroupsToken) | **Put** /realtimefeegroups/{token} | Update real-time fee group



## GetRealtimefeegroups

> RealTimeFeeGroupListResponse GetRealtimefeegroups(ctx).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List real-time fee groups



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
    count := int32(56) // int32 | Number of real-time fee groups to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on).  Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeFeeGroupsAPI.GetRealtimefeegroups(context.Background()).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeFeeGroupsAPI.GetRealtimefeegroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRealtimefeegroups`: RealTimeFeeGroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `RealTimeFeeGroupsAPI.GetRealtimefeegroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRealtimefeegroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of real-time fee groups to retrieve. | [default to 5]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on).  Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**RealTimeFeeGroupListResponse**](RealTimeFeeGroupListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealtimefeegroupsToken

> RealTimeFeeGroup GetRealtimefeegroupsToken(ctx, token).Execute()

Retrieve real-time fee group



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
    token := "token_example" // string | Unique identifier of the real-time fee group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeFeeGroupsAPI.GetRealtimefeegroupsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeFeeGroupsAPI.GetRealtimefeegroupsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRealtimefeegroupsToken`: RealTimeFeeGroup
    fmt.Fprintf(os.Stdout, "Response from `RealTimeFeeGroupsAPI.GetRealtimefeegroupsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the real-time fee group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealtimefeegroupsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RealTimeFeeGroup**](RealTimeFeeGroup.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRealtimefeegroups

> RealTimeFeeGroup PostRealtimefeegroups(ctx).RealTimeFeeGroupCreateRequest(realTimeFeeGroupCreateRequest).Execute()

Create real-time fee group



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
    realTimeFeeGroupCreateRequest := *openapiclient.NewRealTimeFeeGroupCreateRequest("Name_example") // RealTimeFeeGroupCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeFeeGroupsAPI.PostRealtimefeegroups(context.Background()).RealTimeFeeGroupCreateRequest(realTimeFeeGroupCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeFeeGroupsAPI.PostRealtimefeegroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRealtimefeegroups`: RealTimeFeeGroup
    fmt.Fprintf(os.Stdout, "Response from `RealTimeFeeGroupsAPI.PostRealtimefeegroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRealtimefeegroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **realTimeFeeGroupCreateRequest** | [**RealTimeFeeGroupCreateRequest**](RealTimeFeeGroupCreateRequest.md) |  | 

### Return type

[**RealTimeFeeGroup**](RealTimeFeeGroup.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRealtimefeegroupsToken

> RealTimeFeeGroup PutRealtimefeegroupsToken(ctx, token).RealTimeFeeGroupRequest(realTimeFeeGroupRequest).Execute()

Update real-time fee group



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
    token := "token_example" // string | Unique identifier of the real-time fee group.
    realTimeFeeGroupRequest := *openapiclient.NewRealTimeFeeGroupRequest() // RealTimeFeeGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeFeeGroupsAPI.PutRealtimefeegroupsToken(context.Background(), token).RealTimeFeeGroupRequest(realTimeFeeGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeFeeGroupsAPI.PutRealtimefeegroupsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutRealtimefeegroupsToken`: RealTimeFeeGroup
    fmt.Fprintf(os.Stdout, "Response from `RealTimeFeeGroupsAPI.PutRealtimefeegroupsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the real-time fee group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRealtimefeegroupsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realTimeFeeGroupRequest** | [**RealTimeFeeGroupRequest**](RealTimeFeeGroupRequest.md) |  | 

### Return type

[**RealTimeFeeGroup**](RealTimeFeeGroup.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

