# \FeeChargesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeeChargeToken**](FeeChargesAPI.md#GetFeeChargeToken) | **Get** /feecharges/{token} | Retrieve fee charge
[**PostFeeCharge**](FeeChargesAPI.md#PostFeeCharge) | **Post** /feecharges | Create fee charge



## GetFeeChargeToken

> FeeTransferResponse GetFeeChargeToken(ctx, token).Execute()

Retrieve fee charge



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
    token := "token_example" // string | Unique identifier of the fee charge to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeeChargesAPI.GetFeeChargeToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeeChargesAPI.GetFeeChargeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeeChargeToken`: FeeTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `FeeChargesAPI.GetFeeChargeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the fee charge to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeeChargeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeeTransferResponse**](FeeTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeeCharge

> FeeTransferResponse PostFeeCharge(ctx).FeeTransferRequest(feeTransferRequest).Execute()

Create fee charge



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
    feeTransferRequest := *openapiclient.NewFeeTransferRequest("BusinessToken_example", []openapiclient.FeeModel{*openapiclient.NewFeeModel("Token_example")}, "UserToken_example") // FeeTransferRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeeChargesAPI.PostFeeCharge(context.Background()).FeeTransferRequest(feeTransferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeeChargesAPI.PostFeeCharge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFeeCharge`: FeeTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `FeeChargesAPI.PostFeeCharge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFeeChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeTransferRequest** | [**FeeTransferRequest**](FeeTransferRequest.md) |  | 

### Return type

[**FeeTransferResponse**](FeeTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

