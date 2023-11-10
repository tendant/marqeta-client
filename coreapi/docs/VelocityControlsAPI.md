# \VelocityControlsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVelocitycontrols**](VelocityControlsAPI.md#GetVelocitycontrols) | **Get** /velocitycontrols | List velocity controls
[**GetVelocitycontrolsToken**](VelocityControlsAPI.md#GetVelocitycontrolsToken) | **Get** /velocitycontrols/{token} | Returns a specific velocity control
[**GetVelocitycontrolsUserUsertokenAvailable**](VelocityControlsAPI.md#GetVelocitycontrolsUserUsertokenAvailable) | **Get** /velocitycontrols/user/{user_token}/available | List user velocity control balances
[**PostVelocitycontrols**](VelocityControlsAPI.md#PostVelocitycontrols) | **Post** /velocitycontrols | Create velocity control
[**PutVelocitycontrolsToken**](VelocityControlsAPI.md#PutVelocitycontrolsToken) | **Put** /velocitycontrols/{token} | Update velocity control



## GetVelocitycontrols

> VelocityControlListResponse GetVelocitycontrols(ctx).CardProduct(cardProduct).User(user).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List velocity controls



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
    cardProduct := "cardProduct_example" // string | Unique identifier of the card product. Enter the string `null` to retrieve velocity controls that are not associated with any card product. (optional)
    user := "user_example" // string | Unique identifier of the user. Enter the string `null` to retrieve velocity controls that are not associated with any user. (optional)
    count := int32(56) // int32 | Number of velocity control resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityControlsAPI.GetVelocitycontrols(context.Background()).CardProduct(cardProduct).User(user).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityControlsAPI.GetVelocitycontrols``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVelocitycontrols`: VelocityControlListResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityControlsAPI.GetVelocitycontrols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVelocitycontrolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **string** | Unique identifier of the card product. Enter the string &#x60;null&#x60; to retrieve velocity controls that are not associated with any card product. | 
 **user** | **string** | Unique identifier of the user. Enter the string &#x60;null&#x60; to retrieve velocity controls that are not associated with any user. | 
 **count** | **int32** | Number of velocity control resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**VelocityControlListResponse**](VelocityControlListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVelocitycontrolsToken

> VelocityControlResponse GetVelocitycontrolsToken(ctx, token).Fields(fields).Execute()

Returns a specific velocity control



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
    token := "token_example" // string | Unique identifier of the velocity control resource.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityControlsAPI.GetVelocitycontrolsToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityControlsAPI.GetVelocitycontrolsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVelocitycontrolsToken`: VelocityControlResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityControlsAPI.GetVelocitycontrolsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the velocity control resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVelocitycontrolsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**VelocityControlResponse**](VelocityControlResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVelocitycontrolsUserUsertokenAvailable

> VelocityControlBalanceListResponse GetVelocitycontrolsUserUsertokenAvailable(ctx, userToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List user velocity control balances



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
    userToken := "userToken_example" // string | Unique identifier of the cardholder.
    count := int32(56) // int32 | Number of available balance resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityControlsAPI.GetVelocitycontrolsUserUsertokenAvailable(context.Background(), userToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityControlsAPI.GetVelocitycontrolsUserUsertokenAvailable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVelocitycontrolsUserUsertokenAvailable`: VelocityControlBalanceListResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityControlsAPI.GetVelocitycontrolsUserUsertokenAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userToken** | **string** | Unique identifier of the cardholder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVelocitycontrolsUserUsertokenAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of available balance resources to retrieve. | [default to 5]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**VelocityControlBalanceListResponse**](VelocityControlBalanceListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVelocitycontrols

> VelocityControlResponse PostVelocitycontrols(ctx).VelocityControlRequest(velocityControlRequest).Execute()

Create velocity control



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
    velocityControlRequest := *openapiclient.NewVelocityControlRequest(float32(123), "CurrencyCode_example", "VelocityWindow_example") // VelocityControlRequest | Velocity control object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityControlsAPI.PostVelocitycontrols(context.Background()).VelocityControlRequest(velocityControlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityControlsAPI.PostVelocitycontrols``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostVelocitycontrols`: VelocityControlResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityControlsAPI.PostVelocitycontrols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVelocitycontrolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **velocityControlRequest** | [**VelocityControlRequest**](VelocityControlRequest.md) | Velocity control object | 

### Return type

[**VelocityControlResponse**](VelocityControlResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVelocitycontrolsToken

> VelocityControlResponse PutVelocitycontrolsToken(ctx, token).VelocityControlUpdateRequest(velocityControlUpdateRequest).Execute()

Update velocity control



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
    token := "token_example" // string | Unique identifier of the velocity control resource.
    velocityControlUpdateRequest := *openapiclient.NewVelocityControlUpdateRequest("Token_example") // VelocityControlUpdateRequest | Velocity control object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityControlsAPI.PutVelocitycontrolsToken(context.Background(), token).VelocityControlUpdateRequest(velocityControlUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityControlsAPI.PutVelocitycontrolsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutVelocitycontrolsToken`: VelocityControlResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityControlsAPI.PutVelocitycontrolsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the velocity control resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVelocitycontrolsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **velocityControlUpdateRequest** | [**VelocityControlUpdateRequest**](VelocityControlUpdateRequest.md) | Velocity control object | 

### Return type

[**VelocityControlResponse**](VelocityControlResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

