# \CommandoModeAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommandomodes**](CommandoModeAPI.md#GetCommandomodes) | **Get** /commandomodes | List Commando Mode control sets
[**GetCommandomodesCommandomodetokenTransitions**](CommandoModeAPI.md#GetCommandomodesCommandomodetokenTransitions) | **Get** /commandomodes/{commandomode_token}/transitions | List Commando Mode transitions set
[**GetCommandomodesToken**](CommandoModeAPI.md#GetCommandomodesToken) | **Get** /commandomodes/{token} | Retrieve Commando Mode control set
[**GetCommandomodesTransitionsToken**](CommandoModeAPI.md#GetCommandomodesTransitionsToken) | **Get** /commandomodes/transitions/{token} | Retrieve Commando Mode transition



## GetCommandomodes

> CommandoModeListResponse GetCommandomodes(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List Commando Mode control sets



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
    count := int32(56) // int32 | Number of Commando Mode control sets to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommandoModeAPI.GetCommandomodes(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommandoModeAPI.GetCommandomodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandomodes`: CommandoModeListResponse
    fmt.Fprintf(os.Stdout, "Response from `CommandoModeAPI.GetCommandomodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandomodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of Commando Mode control sets to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**CommandoModeListResponse**](CommandoModeListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandomodesCommandomodetokenTransitions

> CommandoModeTransitionListResponse GetCommandomodesCommandomodetokenTransitions(ctx, commandomodeToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List Commando Mode transitions set



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
    commandomodeToken := "commandomodeToken_example" // string | Unique identifier of the Commando Mode control set.
    count := int32(56) // int32 | Number of Commando Mode control set transitions to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommandoModeAPI.GetCommandomodesCommandomodetokenTransitions(context.Background(), commandomodeToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommandoModeAPI.GetCommandomodesCommandomodetokenTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandomodesCommandomodetokenTransitions`: CommandoModeTransitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `CommandoModeAPI.GetCommandomodesCommandomodetokenTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandomodeToken** | **string** | Unique identifier of the Commando Mode control set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandomodesCommandomodetokenTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of Commando Mode control set transitions to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**CommandoModeTransitionListResponse**](CommandoModeTransitionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandomodesToken

> CommandoModeResponse GetCommandomodesToken(ctx, token).Execute()

Retrieve Commando Mode control set



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
    token := "token_example" // string | Unique identifier of the Commando Mode control set you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommandoModeAPI.GetCommandomodesToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommandoModeAPI.GetCommandomodesToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandomodesToken`: CommandoModeResponse
    fmt.Fprintf(os.Stdout, "Response from `CommandoModeAPI.GetCommandomodesToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the Commando Mode control set you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandomodesTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommandoModeResponse**](CommandoModeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandomodesTransitionsToken

> CommandoModeTransitionResponse GetCommandomodesTransitionsToken(ctx, token).Execute()

Retrieve Commando Mode transition



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
    token := "token_example" // string | Unique identifier of the Commando Mode control set transition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommandoModeAPI.GetCommandomodesTransitionsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommandoModeAPI.GetCommandomodesTransitionsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandomodesTransitionsToken`: CommandoModeTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CommandoModeAPI.GetCommandomodesTransitionsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the Commando Mode control set transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandomodesTransitionsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommandoModeTransitionResponse**](CommandoModeTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

