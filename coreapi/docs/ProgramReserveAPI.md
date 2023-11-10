# \ProgramReserveAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProgramreserveBalances**](ProgramReserveAPI.md#GetProgramreserveBalances) | **Get** /programreserve/balances | Retrieve reserve account balances
[**GetProgramreserveTransactions**](ProgramReserveAPI.md#GetProgramreserveTransactions) | **Get** /programreserve/transactions | List program reserve transactions



## GetProgramreserveBalances

> ProgramReserveAccountBalance GetProgramreserveBalances(ctx).Execute()

Retrieve reserve account balances



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
    resp, r, err := apiClient.ProgramReserveAPI.GetProgramreserveBalances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramReserveAPI.GetProgramreserveBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProgramreserveBalances`: ProgramReserveAccountBalance
    fmt.Fprintf(os.Stdout, "Response from `ProgramReserveAPI.GetProgramreserveBalances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramreserveBalancesRequest struct via the builder pattern


### Return type

[**ProgramReserveAccountBalance**](ProgramReserveAccountBalance.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramreserveTransactions

> ProgramReserveTransactionListResponse GetProgramreserveTransactions(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List program reserve transactions



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
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramReserveAPI.GetProgramreserveTransactions(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramReserveAPI.GetProgramreserveTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProgramreserveTransactions`: ProgramReserveTransactionListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramReserveAPI.GetProgramreserveTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramreserveTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**ProgramReserveTransactionListResponse**](ProgramReserveTransactionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

