# \MCCGroupsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMccgroups**](MCCGroupsAPI.md#GetMccgroups) | **Get** /mccgroups | List MCC groups
[**GetMccgroupsToken**](MCCGroupsAPI.md#GetMccgroupsToken) | **Get** /mccgroups/{token} | Retrieve MCC group
[**PostMccgroups**](MCCGroupsAPI.md#PostMccgroups) | **Post** /mccgroups | Create MCC group
[**PutMccgroupsToken**](MCCGroupsAPI.md#PutMccgroupsToken) | **Put** /mccgroups/{token} | Update MCC group



## GetMccgroups

> MCCGroupListResponse GetMccgroups(ctx).Mcc(mcc).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List MCC groups



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
    mcc := "mcc_example" // string | Returns all MCC groups that contain the specified merchant category code. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCCGroupsAPI.GetMccgroups(context.Background()).Mcc(mcc).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCCGroupsAPI.GetMccgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMccgroups`: MCCGroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `MCCGroupsAPI.GetMccgroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMccgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcc** | **string** | Returns all MCC groups that contain the specified merchant category code. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 10]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**MCCGroupListResponse**](MCCGroupListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMccgroupsToken

> MccGroupModel GetMccgroupsToken(ctx, token).Execute()

Retrieve MCC group



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
    token := "token_example" // string | Unique identifier of the MCC group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCCGroupsAPI.GetMccgroupsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCCGroupsAPI.GetMccgroupsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMccgroupsToken`: MccGroupModel
    fmt.Fprintf(os.Stdout, "Response from `MCCGroupsAPI.GetMccgroupsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the MCC group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMccgroupsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MccGroupModel**](MccGroupModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMccgroups

> MccGroupModel PostMccgroups(ctx).MccGroupModel(mccGroupModel).Execute()

Create MCC group



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
    mccGroupModel := *openapiclient.NewMccGroupModel([]map[string]interface{}{map[string]interface{}(123)}, "Name_example") // MccGroupModel | MCC group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCCGroupsAPI.PostMccgroups(context.Background()).MccGroupModel(mccGroupModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCCGroupsAPI.PostMccgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMccgroups`: MccGroupModel
    fmt.Fprintf(os.Stdout, "Response from `MCCGroupsAPI.PostMccgroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMccgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mccGroupModel** | [**MccGroupModel**](MccGroupModel.md) | MCC group | 

### Return type

[**MccGroupModel**](MccGroupModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMccgroupsToken

> MccGroupUpdateModel PutMccgroupsToken(ctx, token).MccGroupUpdateModel(mccGroupUpdateModel).Execute()

Update MCC group



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
    token := "token_example" // string | The unique identifier of the MCC group. Send a `GET` request to `/mccgroups` to retrieve MCC group tokens.
    mccGroupUpdateModel := *openapiclient.NewMccGroupUpdateModel() // MccGroupUpdateModel | MCC group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCCGroupsAPI.PutMccgroupsToken(context.Background(), token).MccGroupUpdateModel(mccGroupUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCCGroupsAPI.PutMccgroupsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMccgroupsToken`: MccGroupUpdateModel
    fmt.Fprintf(os.Stdout, "Response from `MCCGroupsAPI.PutMccgroupsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the MCC group. Send a &#x60;GET&#x60; request to &#x60;/mccgroups&#x60; to retrieve MCC group tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMccgroupsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mccGroupUpdateModel** | [**MccGroupUpdateModel**](MccGroupUpdateModel.md) | MCC group | 

### Return type

[**MccGroupUpdateModel**](MccGroupUpdateModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

