# \AuthorizationControlsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthcontrols**](AuthorizationControlsAPI.md#GetAuthcontrols) | **Get** /authcontrols | List authorization controls
[**GetAuthcontrolsExemptmids**](AuthorizationControlsAPI.md#GetAuthcontrolsExemptmids) | **Get** /authcontrols/exemptmids | List merchant identifier (MID) exemptions
[**GetAuthcontrolsExemptmidsToken**](AuthorizationControlsAPI.md#GetAuthcontrolsExemptmidsToken) | **Get** /authcontrols/exemptmids/{token} | Retrieve a merchant identifier (MID) exemption
[**GetAuthcontrolsToken**](AuthorizationControlsAPI.md#GetAuthcontrolsToken) | **Get** /authcontrols/{token} | Retrieve authorization control
[**PostAuthcontrols**](AuthorizationControlsAPI.md#PostAuthcontrols) | **Post** /authcontrols | Create authorization control
[**PostAuthcontrolsExemptmids**](AuthorizationControlsAPI.md#PostAuthcontrolsExemptmids) | **Post** /authcontrols/exemptmids | Create a merchant identifier (MID) exemption
[**PutAuthcontrolsExemptmidsToken**](AuthorizationControlsAPI.md#PutAuthcontrolsExemptmidsToken) | **Put** /authcontrols/exemptmids/{token} | Update a merchant identifier (MID) exemption
[**PutAuthcontrolsToken**](AuthorizationControlsAPI.md#PutAuthcontrolsToken) | **Put** /authcontrols/{token} | Update authorization control



## GetAuthcontrols

> AuthControlListResponse GetAuthcontrols(ctx).CardProduct(cardProduct).User(user).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List authorization controls



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
    cardProduct := "cardProduct_example" // string | Unique identifier of the card product whose associated authorization controls you want to retrieve.  Enter the string \"null\" to list authorization controls that are not associated with a card product. (optional)
    user := "user_example" // string | Unique identifier of the user whose associated authorization controls you want to retrieve.  Enter the string \"null\" to list authorization controls that are not associated with a user. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationControlsAPI.GetAuthcontrols(context.Background()).CardProduct(cardProduct).User(user).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.GetAuthcontrols``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthcontrols`: AuthControlListResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationControlsAPI.GetAuthcontrols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthcontrolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **string** | Unique identifier of the card product whose associated authorization controls you want to retrieve.  Enter the string \&quot;null\&quot; to list authorization controls that are not associated with a card product. | 
 **user** | **string** | Unique identifier of the user whose associated authorization controls you want to retrieve.  Enter the string \&quot;null\&quot; to list authorization controls that are not associated with a user. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**AuthControlListResponse**](AuthControlListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthcontrolsExemptmids

> AuthControlExemptMidsListResponse GetAuthcontrolsExemptmids(ctx).CardProduct(cardProduct).User(user).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List merchant identifier (MID) exemptions



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
    cardProduct := "cardProduct_example" // string | Unique identifier of the card product whose associated MID exemptions you want to retrieve.  Enter the string \"null\" to list MID exemptions that are not associated with a card product. (optional)
    user := "user_example" // string | Unique identifier of the user whose associated MID exemptions you want to retrieve.  Enter the string \"null\" to list MID exemptions that are not associated with a user. (optional)
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationControlsAPI.GetAuthcontrolsExemptmids(context.Background()).CardProduct(cardProduct).User(user).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.GetAuthcontrolsExemptmids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthcontrolsExemptmids`: AuthControlExemptMidsListResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationControlsAPI.GetAuthcontrolsExemptmids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthcontrolsExemptmidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardProduct** | **string** | Unique identifier of the card product whose associated MID exemptions you want to retrieve.  Enter the string \&quot;null\&quot; to list MID exemptions that are not associated with a card product. | 
 **user** | **string** | Unique identifier of the user whose associated MID exemptions you want to retrieve.  Enter the string \&quot;null\&quot; to list MID exemptions that are not associated with a user. | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**AuthControlExemptMidsListResponse**](AuthControlExemptMidsListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthcontrolsExemptmidsToken

> AuthControlExemptMidsResponse GetAuthcontrolsExemptmidsToken(ctx, token).Execute()

Retrieve a merchant identifier (MID) exemption



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
    token := "token_example" // string | Unique identifier of the authorization control resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationControlsAPI.GetAuthcontrolsExemptmidsToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.GetAuthcontrolsExemptmidsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthcontrolsExemptmidsToken`: AuthControlExemptMidsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationControlsAPI.GetAuthcontrolsExemptmidsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the authorization control resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthcontrolsExemptmidsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthControlExemptMidsResponse**](AuthControlExemptMidsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthcontrolsToken

> AuthControlResponse GetAuthcontrolsToken(ctx, token).Fields(fields).Execute()

Retrieve authorization control



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
    token := "token_example" // string | Existing authorization control token.  Send a `GET` request to `/authcontrols` to retrieve authorization control tokens.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationControlsAPI.GetAuthcontrolsToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.GetAuthcontrolsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthcontrolsToken`: AuthControlResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationControlsAPI.GetAuthcontrolsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Existing authorization control token.  Send a &#x60;GET&#x60; request to &#x60;/authcontrols&#x60; to retrieve authorization control tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthcontrolsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**AuthControlResponse**](AuthControlResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthcontrols

> AuthControlResponse PostAuthcontrols(ctx).AuthControlRequest(authControlRequest).Execute()

Create authorization control



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
    authControlRequest := *openapiclient.NewAuthControlRequest("Name_example") // AuthControlRequest | Auth control object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationControlsAPI.PostAuthcontrols(context.Background()).AuthControlRequest(authControlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.PostAuthcontrols``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAuthcontrols`: AuthControlResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationControlsAPI.PostAuthcontrols`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthcontrolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authControlRequest** | [**AuthControlRequest**](AuthControlRequest.md) | Auth control object | 

### Return type

[**AuthControlResponse**](AuthControlResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthcontrolsExemptmids

> AuthControlExemptMidsResponse PostAuthcontrolsExemptmids(ctx).AuthControlExemptMidsRequest(authControlExemptMidsRequest).Execute()

Create a merchant identifier (MID) exemption



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
    authControlExemptMidsRequest := *openapiclient.NewAuthControlExemptMidsRequest("Name_example") // AuthControlExemptMidsRequest | Auth control exempt MID object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationControlsAPI.PostAuthcontrolsExemptmids(context.Background()).AuthControlExemptMidsRequest(authControlExemptMidsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.PostAuthcontrolsExemptmids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAuthcontrolsExemptmids`: AuthControlExemptMidsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationControlsAPI.PostAuthcontrolsExemptmids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthcontrolsExemptmidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authControlExemptMidsRequest** | [**AuthControlExemptMidsRequest**](AuthControlExemptMidsRequest.md) | Auth control exempt MID object | 

### Return type

[**AuthControlExemptMidsResponse**](AuthControlExemptMidsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAuthcontrolsExemptmidsToken

> PutAuthcontrolsExemptmidsToken(ctx, token).AuthControlExemptMidsUpdateRequest(authControlExemptMidsUpdateRequest).Execute()

Update a merchant identifier (MID) exemption



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
    token := "token_example" // string | Unique identifier of the authorization control resource.
    authControlExemptMidsUpdateRequest := *openapiclient.NewAuthControlExemptMidsUpdateRequest() // AuthControlExemptMidsUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationControlsAPI.PutAuthcontrolsExemptmidsToken(context.Background(), token).AuthControlExemptMidsUpdateRequest(authControlExemptMidsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.PutAuthcontrolsExemptmidsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the authorization control resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAuthcontrolsExemptmidsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authControlExemptMidsUpdateRequest** | [**AuthControlExemptMidsUpdateRequest**](AuthControlExemptMidsUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAuthcontrolsToken

> AuthControlResponse PutAuthcontrolsToken(ctx, token).AuthControlUpdateRequest(authControlUpdateRequest).Execute()

Update authorization control



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
    token := "token_example" // string | Existing authorization control token.  Send a `GET` request to `/authcontrols` to retrieve authorization control tokens.
    authControlUpdateRequest := *openapiclient.NewAuthControlUpdateRequest("Token_example") // AuthControlUpdateRequest | Auth control object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationControlsAPI.PutAuthcontrolsToken(context.Background(), token).AuthControlUpdateRequest(authControlUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationControlsAPI.PutAuthcontrolsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAuthcontrolsToken`: AuthControlResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationControlsAPI.PutAuthcontrolsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Existing authorization control token.  Send a &#x60;GET&#x60; request to &#x60;/authcontrols&#x60; to retrieve authorization control tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAuthcontrolsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authControlUpdateRequest** | [**AuthControlUpdateRequest**](AuthControlUpdateRequest.md) | Auth control object | 

### Return type

[**AuthControlResponse**](AuthControlResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

