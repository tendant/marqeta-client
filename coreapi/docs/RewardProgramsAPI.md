# \RewardProgramsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveRedemptionsBySettlementDate**](RewardProgramsAPI.md#RetrieveRedemptionsBySettlementDate) | **Get** /credit/rewardprograms/redemptions | retrieves all completed redemptions by settlement date



## RetrieveRedemptionsBySettlementDate

> RedemptionsBySettlementDatePage RetrieveRedemptionsBySettlementDate(ctx).SettlementStartDate(settlementStartDate).SettlementEndDate(settlementEndDate).Destination(destination).Count(count).StartIndex(startIndex).Execute()

retrieves all completed redemptions by settlement date



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
    settlementStartDate := time.Now() // time.Time | Settlement start date to filter by.
    settlementEndDate := time.Now() // time.Time | Settlement end date to filter by.
    destination := openapiclient.DestinationType("INVESTMENT") // DestinationType | Specifies the destination for external redemptions to filter for. (optional)
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int64(789) // int64 | Sort order index of the first resource in the returned array. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsAPI.RetrieveRedemptionsBySettlementDate(context.Background()).SettlementStartDate(settlementStartDate).SettlementEndDate(settlementEndDate).Destination(destination).Count(count).StartIndex(startIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsAPI.RetrieveRedemptionsBySettlementDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRedemptionsBySettlementDate`: RedemptionsBySettlementDatePage
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsAPI.RetrieveRedemptionsBySettlementDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRedemptionsBySettlementDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settlementStartDate** | **time.Time** | Settlement start date to filter by. | 
 **settlementEndDate** | **time.Time** | Settlement end date to filter by. | 
 **destination** | [**DestinationType**](DestinationType.md) | Specifies the destination for external redemptions to filter for. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int64** | Sort order index of the first resource in the returned array. | [default to 0]

### Return type

[**RedemptionsBySettlementDatePage**](RedemptionsBySettlementDatePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

