# \PeerTransfersAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeertransfersToken**](PeerTransfersAPI.md#GetPeertransfersToken) | **Get** /peertransfers/{token} | Retrieve peer transfer
[**GetPeertransfersUserUserorbusinesstoken**](PeerTransfersAPI.md#GetPeertransfersUserUserorbusinesstoken) | **Get** /peertransfers/user/{user_or_business_token} | List peer transfers by account holder
[**GetPeertransfersUserUserorbusinesstokenRecipient**](PeerTransfersAPI.md#GetPeertransfersUserUserorbusinesstokenRecipient) | **Get** /peertransfers/user/{user_or_business_token}/recipient | List received peer transfers
[**GetPeertransfersUserUserorbusinesstokenSender**](PeerTransfersAPI.md#GetPeertransfersUserUserorbusinesstokenSender) | **Get** /peertransfers/user/{user_or_business_token}/sender | List sent peer transfers
[**PostPeertransfers**](PeerTransfersAPI.md#PostPeertransfers) | **Post** /peertransfers | Create peer transfer



## GetPeertransfersToken

> PeerTransferResponse GetPeertransfersToken(ctx, token).Execute()

Retrieve peer transfer



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
    token := "token_example" // string | Unique identifier of the peer transfer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeerTransfersAPI.GetPeertransfersToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeerTransfersAPI.GetPeertransfersToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeertransfersToken`: PeerTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `PeerTransfersAPI.GetPeertransfersToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the peer transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeertransfersTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PeerTransferResponse**](PeerTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeertransfersUserUserorbusinesstoken

> PeerTransferResponse GetPeertransfersUserUserorbusinesstoken(ctx, userOrBusinessToken).Count(count).StartIndex(startIndex).Fields(fields).Execute()

List peer transfers by account holder



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
    userOrBusinessToken := "userOrBusinessToken_example" // string | Existing user or business token.  Send a `GET` request to `/users` to retrieve user tokens or to `/businesses` to retrieve business tokens.
    count := int32(56) // int32 | Number of peer transfer resources to retrieve. (optional) (default to 25)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeerTransfersAPI.GetPeertransfersUserUserorbusinesstoken(context.Background(), userOrBusinessToken).Count(count).StartIndex(startIndex).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeerTransfersAPI.GetPeertransfersUserUserorbusinesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeertransfersUserUserorbusinesstoken`: PeerTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `PeerTransfersAPI.GetPeertransfersUserUserorbusinesstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userOrBusinessToken** | **string** | Existing user or business token.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens or to &#x60;/businesses&#x60; to retrieve business tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeertransfersUserUserorbusinesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of peer transfer resources to retrieve. | [default to 25]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**PeerTransferResponse**](PeerTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeertransfersUserUserorbusinesstokenRecipient

> PeerTransferResponse GetPeertransfersUserUserorbusinesstokenRecipient(ctx, userOrBusinessToken).Count(count).StartIndex(startIndex).Fields(fields).Execute()

List received peer transfers



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
    userOrBusinessToken := "userOrBusinessToken_example" // string | Existing user or business token.  Send a `GET` request to `/users` to retrieve user tokens or to `/businesses` to retrieve business tokens.
    count := int32(56) // int32 | Number of peer transfer resources to retrieve. (optional) (default to 25)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeerTransfersAPI.GetPeertransfersUserUserorbusinesstokenRecipient(context.Background(), userOrBusinessToken).Count(count).StartIndex(startIndex).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeerTransfersAPI.GetPeertransfersUserUserorbusinesstokenRecipient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeertransfersUserUserorbusinesstokenRecipient`: PeerTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `PeerTransfersAPI.GetPeertransfersUserUserorbusinesstokenRecipient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userOrBusinessToken** | **string** | Existing user or business token.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens or to &#x60;/businesses&#x60; to retrieve business tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeertransfersUserUserorbusinesstokenRecipientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of peer transfer resources to retrieve. | [default to 25]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**PeerTransferResponse**](PeerTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeertransfersUserUserorbusinesstokenSender

> PeerTransferResponse GetPeertransfersUserUserorbusinesstokenSender(ctx, userOrBusinessToken).Count(count).StartIndex(startIndex).Fields(fields).Execute()

List sent peer transfers



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
    userOrBusinessToken := "userOrBusinessToken_example" // string | Existing user or business token.  Send a `GET` request to `/users` to retrieve user tokens or to `/businesses` to retrieve business tokens.
    count := int32(56) // int32 | Number of peer transfer resources to retrieve. (optional) (default to 25)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeerTransfersAPI.GetPeertransfersUserUserorbusinesstokenSender(context.Background(), userOrBusinessToken).Count(count).StartIndex(startIndex).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeerTransfersAPI.GetPeertransfersUserUserorbusinesstokenSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeertransfersUserUserorbusinesstokenSender`: PeerTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `PeerTransfersAPI.GetPeertransfersUserUserorbusinesstokenSender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userOrBusinessToken** | **string** | Existing user or business token.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens or to &#x60;/businesses&#x60; to retrieve business tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeertransfersUserUserorbusinesstokenSenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of peer transfer resources to retrieve. | [default to 25]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**PeerTransferResponse**](PeerTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPeertransfers

> PeerTransferResponse PostPeertransfers(ctx).PeerTransferRequest(peerTransferRequest).Execute()

Create peer transfer



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
    peerTransferRequest := *openapiclient.NewPeerTransferRequest(float32(123), "CurrencyCode_example") // PeerTransferRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeerTransfersAPI.PostPeertransfers(context.Background()).PeerTransferRequest(peerTransferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeerTransfersAPI.PostPeertransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPeertransfers`: PeerTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `PeerTransfersAPI.PostPeertransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPeertransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **peerTransferRequest** | [**PeerTransferRequest**](PeerTransferRequest.md) |  | 

### Return type

[**PeerTransferResponse**](PeerTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

