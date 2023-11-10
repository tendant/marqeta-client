# \PINsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPinsControltoken**](PINsAPI.md#PostPinsControltoken) | **Post** /pins/controltoken | Create PIN control token
[**PutPins**](PINsAPI.md#PutPins) | **Put** /pins | Create or update PIN
[**RevealPins**](PINsAPI.md#RevealPins) | **Post** /pins/reveal | Reveal PIN



## PostPinsControltoken

> ControlTokenResponse PostPinsControltoken(ctx).ControlTokenRequest(controlTokenRequest).Execute()

Create PIN control token



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
    controlTokenRequest := *openapiclient.NewControlTokenRequest("CardToken_example") // ControlTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PINsAPI.PostPinsControltoken(context.Background()).ControlTokenRequest(controlTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PINsAPI.PostPinsControltoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPinsControltoken`: ControlTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PINsAPI.PostPinsControltoken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPinsControltokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlTokenRequest** | [**ControlTokenRequest**](ControlTokenRequest.md) |  | 

### Return type

[**ControlTokenResponse**](ControlTokenResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPins

> PutPins(ctx).PinRequest(pinRequest).Execute()

Create or update PIN



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
    pinRequest := *openapiclient.NewPinRequest("ControlToken_example", "Pin_example") // PinRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PINsAPI.PutPins(context.Background()).PinRequest(pinRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PINsAPI.PutPins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pinRequest** | [**PinRequest**](PinRequest.md) |  | 

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


## RevealPins

> RevealPins(ctx).PinRevealRequest(pinRevealRequest).Execute()

Reveal PIN



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
    pinRevealRequest := *openapiclient.NewPinRevealRequest("CardholderVerificationMethod_example", "ControlToken_example") // PinRevealRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PINsAPI.RevealPins(context.Background()).PinRevealRequest(pinRevealRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PINsAPI.RevealPins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevealPinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pinRevealRequest** | [**PinRevealRequest**](PinRevealRequest.md) |  | 

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

