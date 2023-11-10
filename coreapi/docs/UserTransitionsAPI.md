# \UserTransitionsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsertransitionsToken**](UserTransitionsAPI.md#GetUsertransitionsToken) | **Get** /usertransitions/{token} | Retrieve user transition
[**GetUsertransitionsUserUsertoken**](UserTransitionsAPI.md#GetUsertransitionsUserUsertoken) | **Get** /usertransitions/user/{user_token} | List transitions for user
[**PostUsertransitions**](UserTransitionsAPI.md#PostUsertransitions) | **Post** /usertransitions | Create user transition



## GetUsertransitionsToken

> UserTransitionResponse GetUsertransitionsToken(ctx, token).Fields(fields).Execute()

Retrieve user transition



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
    token := "token_example" // string | Unique identifier of the user transition you want to retrieve.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTransitionsAPI.GetUsertransitionsToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTransitionsAPI.GetUsertransitionsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsertransitionsToken`: UserTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `UserTransitionsAPI.GetUsertransitionsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the user transition you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsertransitionsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**UserTransitionResponse**](UserTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsertransitionsUserUsertoken

> UserTransitionListResponse GetUsertransitionsUserUsertoken(ctx, userToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List transitions for user



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
    userToken := "userToken_example" // string | Unique identifier of the user resource.
    count := int32(56) // int32 | Number of user transitions to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-id")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTransitionsAPI.GetUsertransitionsUserUsertoken(context.Background(), userToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTransitionsAPI.GetUsertransitionsUserUsertoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsertransitionsUserUsertoken`: UserTransitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `UserTransitionsAPI.GetUsertransitionsUserUsertoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userToken** | **string** | Unique identifier of the user resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsertransitionsUserUsertokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of user transitions to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-id&quot;]

### Return type

[**UserTransitionListResponse**](UserTransitionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsertransitions

> UserTransitionResponse PostUsertransitions(ctx).UserTransitionRequest(userTransitionRequest).Execute()

Create user transition



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
    userTransitionRequest := *openapiclient.NewUserTransitionRequest("Channel_example", "ReasonCode_example", "Status_example", "UserToken_example") // UserTransitionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserTransitionsAPI.PostUsertransitions(context.Background()).UserTransitionRequest(userTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTransitionsAPI.PostUsertransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsertransitions`: UserTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `UserTransitionsAPI.PostUsertransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsertransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userTransitionRequest** | [**UserTransitionRequest**](UserTransitionRequest.md) |  | 

### Return type

[**UserTransitionResponse**](UserTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

