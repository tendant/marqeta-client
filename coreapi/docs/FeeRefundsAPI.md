# \FeeRefundsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostFeeRefunds**](FeeRefundsAPI.md#PostFeeRefunds) | **Post** /feerefunds | Create fee refund



## PostFeeRefunds

> FeeTransferResponse PostFeeRefunds(ctx).FeeRefundRequest(feeRefundRequest).Execute()

Create fee refund



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
    feeRefundRequest := *openapiclient.NewFeeRefundRequest() // FeeRefundRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeeRefundsAPI.PostFeeRefunds(context.Background()).FeeRefundRequest(feeRefundRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeeRefundsAPI.PostFeeRefunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFeeRefunds`: FeeTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `FeeRefundsAPI.PostFeeRefunds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFeeRefundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeRefundRequest** | [**FeeRefundRequest**](FeeRefundRequest.md) |  | 

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

