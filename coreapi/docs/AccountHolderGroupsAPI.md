# \AccountHolderGroupsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountholdergroups**](AccountHolderGroupsAPI.md#GetAccountholdergroups) | **Get** /accountholdergroups | List account holder groups
[**GetAccountholdergroupsToken**](AccountHolderGroupsAPI.md#GetAccountholdergroupsToken) | **Get** /accountholdergroups/{token} | Retrieve account holder group
[**PostAccountholdergroups**](AccountHolderGroupsAPI.md#PostAccountholdergroups) | **Post** /accountholdergroups | Create account holder group
[**PutAccountholdergroupsToken**](AccountHolderGroupsAPI.md#PutAccountholdergroupsToken) | **Put** /accountholdergroups/{token} | Update account holder group



## GetAccountholdergroups

> AccountHolderGroupListResponse GetAccountholdergroups(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List account holder groups



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
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderGroupsAPI.GetAccountholdergroups(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderGroupsAPI.GetAccountholdergroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountholdergroups`: AccountHolderGroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderGroupsAPI.GetAccountholdergroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountholdergroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of resources to retrieve. | [default to 10]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**AccountHolderGroupListResponse**](AccountHolderGroupListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountholdergroupsToken

> AccountHolderGroupResponse GetAccountholdergroupsToken(ctx, token).Execute()

Retrieve account holder group



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
    token := "token_example" // string | Unique identifier of the account holder group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderGroupsAPI.GetAccountholdergroupsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderGroupsAPI.GetAccountholdergroupsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountholdergroupsToken`: AccountHolderGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderGroupsAPI.GetAccountholdergroupsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the account holder group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountholdergroupsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountHolderGroupResponse**](AccountHolderGroupResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountholdergroups

> AccountHolderGroupResponse PostAccountholdergroups(ctx).AccountHolderGroupRequest(accountHolderGroupRequest).Execute()

Create account holder group



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
    accountHolderGroupRequest := *openapiclient.NewAccountHolderGroupRequest() // AccountHolderGroupRequest | Account holder group object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderGroupsAPI.PostAccountholdergroups(context.Background()).AccountHolderGroupRequest(accountHolderGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderGroupsAPI.PostAccountholdergroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAccountholdergroups`: AccountHolderGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderGroupsAPI.PostAccountholdergroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountholdergroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountHolderGroupRequest** | [**AccountHolderGroupRequest**](AccountHolderGroupRequest.md) | Account holder group object | 

### Return type

[**AccountHolderGroupResponse**](AccountHolderGroupResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccountholdergroupsToken

> AccountHolderGroupResponse PutAccountholdergroupsToken(ctx, token).AccountHolderGroupUpdateRequest(accountHolderGroupUpdateRequest).Execute()

Update account holder group



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
    token := "token_example" // string | Unique identifier of the account holder group.
    accountHolderGroupUpdateRequest := *openapiclient.NewAccountHolderGroupUpdateRequest() // AccountHolderGroupUpdateRequest | Account holder group update object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderGroupsAPI.PutAccountholdergroupsToken(context.Background(), token).AccountHolderGroupUpdateRequest(accountHolderGroupUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderGroupsAPI.PutAccountholdergroupsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAccountholdergroupsToken`: AccountHolderGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderGroupsAPI.PutAccountholdergroupsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the account holder group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccountholdergroupsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountHolderGroupUpdateRequest** | [**AccountHolderGroupUpdateRequest**](AccountHolderGroupUpdateRequest.md) | Account holder group update object | 

### Return type

[**AccountHolderGroupResponse**](AccountHolderGroupResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

