# \PingAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPing**](PingAPI.md#GetPing) | **Get** /ping | Returns a heartbeat to the consumer
[**PostPing**](PingAPI.md#PostPing) | **Post** /ping | Echo test for sending payload to server



## GetPing

> PingResponse GetPing(ctx).Execute()

Returns a heartbeat to the consumer



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingAPI.GetPing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingAPI.GetPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPing`: PingResponse
    fmt.Fprintf(os.Stdout, "Response from `PingAPI.GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingRequest struct via the builder pattern


### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPing

> EchoPingResponse PostPing(ctx).EchoPingRequest(echoPingRequest).Execute()

Echo test for sending payload to server

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
    echoPingRequest := *openapiclient.NewEchoPingRequest() // EchoPingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingAPI.PostPing(context.Background()).EchoPingRequest(echoPingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingAPI.PostPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPing`: EchoPingResponse
    fmt.Fprintf(os.Stdout, "Response from `PingAPI.PostPing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **echoPingRequest** | [**EchoPingRequest**](EchoPingRequest.md) |  | 

### Return type

[**EchoPingResponse**](EchoPingResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

