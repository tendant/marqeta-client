# \SimulationsDirectDepositsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectDepositsCreditEvent**](SimulationsDirectDepositsAPI.md#DirectDepositsCreditEvent) | **Post** /simulations/directdeposits/credit | Simulate credit
[**DirectDepositsDebitEvent**](SimulationsDirectDepositsAPI.md#DirectDepositsDebitEvent) | **Post** /simulations/directdeposits/debit | Simulate debit



## DirectDepositsCreditEvent

> DirectDepositSimulationResponse DirectDepositsCreditEvent(ctx).DirectDepositSimulationRequest(directDepositSimulationRequest).Execute()

Simulate credit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    directDepositSimulationRequest := *openapiclient.NewDirectDepositSimulationRequest("AccountNumber_example", float32(123), time.Now()) // DirectDepositSimulationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsDirectDepositsAPI.DirectDepositsCreditEvent(context.Background()).DirectDepositSimulationRequest(directDepositSimulationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsDirectDepositsAPI.DirectDepositsCreditEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectDepositsCreditEvent`: DirectDepositSimulationResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsDirectDepositsAPI.DirectDepositsCreditEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectDepositsCreditEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directDepositSimulationRequest** | [**DirectDepositSimulationRequest**](DirectDepositSimulationRequest.md) |  | 

### Return type

[**DirectDepositSimulationResponse**](DirectDepositSimulationResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectDepositsDebitEvent

> DirectDepositSimulationResponse DirectDepositsDebitEvent(ctx).DirectDepositSimulationRequest(directDepositSimulationRequest).Execute()

Simulate debit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    directDepositSimulationRequest := *openapiclient.NewDirectDepositSimulationRequest("AccountNumber_example", float32(123), time.Now()) // DirectDepositSimulationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimulationsDirectDepositsAPI.DirectDepositsDebitEvent(context.Background()).DirectDepositSimulationRequest(directDepositSimulationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimulationsDirectDepositsAPI.DirectDepositsDebitEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectDepositsDebitEvent`: DirectDepositSimulationResponse
    fmt.Fprintf(os.Stdout, "Response from `SimulationsDirectDepositsAPI.DirectDepositsDebitEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectDepositsDebitEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directDepositSimulationRequest** | [**DirectDepositSimulationRequest**](DirectDepositSimulationRequest.md) |  | 

### Return type

[**DirectDepositSimulationResponse**](DirectDepositSimulationResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

