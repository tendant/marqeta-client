# \AccountAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeriodicFeeSchedules**](AccountAPI.md#GetPeriodicFeeSchedules) | **Get** /credit/accounts/{account_token}/periodicfeeschedules | Get all active and upcoming periodic fee schedules of an account



## GetPeriodicFeeSchedules

> PeriodicFeeSchedulePage GetPeriodicFeeSchedules(ctx, accountToken).Count(count).StartIndex(startIndex).Execute()

Get all active and upcoming periodic fee schedules of an account



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
    accountToken := "accountToken_example" // string | account token
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.GetPeriodicFeeSchedules(context.Background(), accountToken).Count(count).StartIndex(startIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetPeriodicFeeSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeriodicFeeSchedules`: PeriodicFeeSchedulePage
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetPeriodicFeeSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | account token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeriodicFeeSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]

### Return type

[**PeriodicFeeSchedulePage**](PeriodicFeeSchedulePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

