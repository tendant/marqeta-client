# \AccountRewardsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReward**](AccountRewardsAPI.md#CreateReward) | **Post** /credit/accounts/{account_token}/rewards | Create account reward



## CreateReward

> RewardResponse CreateReward(ctx, accountToken).RewardCreateReq(rewardCreateReq).Execute()

Create account reward



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to create a reward.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    rewardCreateReq := *openapiclient.NewRewardCreateReq(float32(123), openapiclient.CurrencyCode("USD"), "Description_example") // RewardCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountRewardsAPI.CreateReward(context.Background(), accountToken).RewardCreateReq(rewardCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountRewardsAPI.CreateReward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReward`: RewardResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountRewardsAPI.CreateReward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to create a reward.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRewardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rewardCreateReq** | [**RewardCreateReq**](RewardCreateReq.md) |  | 

### Return type

[**RewardResponse**](RewardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

