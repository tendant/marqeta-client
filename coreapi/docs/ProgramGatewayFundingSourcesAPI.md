# \ProgramGatewayFundingSourcesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFundingsourcesProgramgatewayToken**](ProgramGatewayFundingSourcesAPI.md#GetFundingsourcesProgramgatewayToken) | **Get** /fundingsources/programgateway/{token} | Retrieve program gateway source
[**PostFundingsourcesProgramgateway**](ProgramGatewayFundingSourcesAPI.md#PostFundingsourcesProgramgateway) | **Post** /fundingsources/programgateway | Create program gateway source
[**PutFundingsourcesProgramgatewayCustomHeaderToken**](ProgramGatewayFundingSourcesAPI.md#PutFundingsourcesProgramgatewayCustomHeaderToken) | **Put** /fundingsources/programgateway/customheaders/{token} | Update program gateway source custom headers
[**PutFundingsourcesProgramgatewayToken**](ProgramGatewayFundingSourcesAPI.md#PutFundingsourcesProgramgatewayToken) | **Put** /fundingsources/programgateway/{token} | Update program gateway source



## GetFundingsourcesProgramgatewayToken

> GatewayProgramFundingSourceResponse GetFundingsourcesProgramgatewayToken(ctx, token).Execute()

Retrieve program gateway source



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
    token := "token_example" // string | Unique identifier of the program gateway funding source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramGatewayFundingSourcesAPI.GetFundingsourcesProgramgatewayToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewayFundingSourcesAPI.GetFundingsourcesProgramgatewayToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesProgramgatewayToken`: GatewayProgramFundingSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramGatewayFundingSourcesAPI.GetFundingsourcesProgramgatewayToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the program gateway funding source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesProgramgatewayTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayProgramFundingSourceResponse**](GatewayProgramFundingSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFundingsourcesProgramgateway

> GatewayProgramFundingSourceResponse PostFundingsourcesProgramgateway(ctx).GatewayProgramFundingSourceRequest(gatewayProgramFundingSourceRequest).Execute()

Create program gateway source



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
    gatewayProgramFundingSourceRequest := *openapiclient.NewGatewayProgramFundingSourceRequest("BasicAuthPassword_example", "BasicAuthUsername_example", "Name_example", "Url_example") // GatewayProgramFundingSourceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramGatewayFundingSourcesAPI.PostFundingsourcesProgramgateway(context.Background()).GatewayProgramFundingSourceRequest(gatewayProgramFundingSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewayFundingSourcesAPI.PostFundingsourcesProgramgateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFundingsourcesProgramgateway`: GatewayProgramFundingSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramGatewayFundingSourcesAPI.PostFundingsourcesProgramgateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFundingsourcesProgramgatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayProgramFundingSourceRequest** | [**GatewayProgramFundingSourceRequest**](GatewayProgramFundingSourceRequest.md) |  | 

### Return type

[**GatewayProgramFundingSourceResponse**](GatewayProgramFundingSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFundingsourcesProgramgatewayCustomHeaderToken

> GatewayProgramFundingSourceResponse PutFundingsourcesProgramgatewayCustomHeaderToken(ctx, token).GatewayProgramCustomHeaderUpdateRequest(gatewayProgramCustomHeaderUpdateRequest).Execute()

Update program gateway source custom headers



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
    token := "token_example" // string | Unique identifier of the program gateway funding source.
    gatewayProgramCustomHeaderUpdateRequest := *openapiclient.NewGatewayProgramCustomHeaderUpdateRequest() // GatewayProgramCustomHeaderUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramGatewayFundingSourcesAPI.PutFundingsourcesProgramgatewayCustomHeaderToken(context.Background(), token).GatewayProgramCustomHeaderUpdateRequest(gatewayProgramCustomHeaderUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewayFundingSourcesAPI.PutFundingsourcesProgramgatewayCustomHeaderToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFundingsourcesProgramgatewayCustomHeaderToken`: GatewayProgramFundingSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramGatewayFundingSourcesAPI.PutFundingsourcesProgramgatewayCustomHeaderToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the program gateway funding source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFundingsourcesProgramgatewayCustomHeaderTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayProgramCustomHeaderUpdateRequest** | [**GatewayProgramCustomHeaderUpdateRequest**](GatewayProgramCustomHeaderUpdateRequest.md) |  | 

### Return type

[**GatewayProgramFundingSourceResponse**](GatewayProgramFundingSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFundingsourcesProgramgatewayToken

> GatewayProgramFundingSourceResponse PutFundingsourcesProgramgatewayToken(ctx, token).GatewayProgramFundingSourceUpdateRequest(gatewayProgramFundingSourceUpdateRequest).Execute()

Update program gateway source



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
    token := "token_example" // string | Unique identifier of the program gateway funding source.
    gatewayProgramFundingSourceUpdateRequest := *openapiclient.NewGatewayProgramFundingSourceUpdateRequest("BasicAuthPassword_example", "BasicAuthUsername_example", "Url_example") // GatewayProgramFundingSourceUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramGatewayFundingSourcesAPI.PutFundingsourcesProgramgatewayToken(context.Background(), token).GatewayProgramFundingSourceUpdateRequest(gatewayProgramFundingSourceUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewayFundingSourcesAPI.PutFundingsourcesProgramgatewayToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFundingsourcesProgramgatewayToken`: GatewayProgramFundingSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramGatewayFundingSourcesAPI.PutFundingsourcesProgramgatewayToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the program gateway funding source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFundingsourcesProgramgatewayTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayProgramFundingSourceUpdateRequest** | [**GatewayProgramFundingSourceUpdateRequest**](GatewayProgramFundingSourceUpdateRequest.md) |  | 

### Return type

[**GatewayProgramFundingSourceResponse**](GatewayProgramFundingSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

