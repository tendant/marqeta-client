# \RewardRedemptionsBetaAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRedemption**](RewardRedemptionsBetaAPI.md#GetRedemption) | **Get** /credit/rewardprograms/{token}/redemptions/{redemption_token} | Retrieve reward redemption
[**PostRedemptionTransition**](RewardRedemptionsBetaAPI.md#PostRedemptionTransition) | **Post** /credit/rewardprograms/{token}/redemptions/{redemption_token}/transitions | Transition reward redemption status
[**PostRewardProgramRedemption**](RewardRedemptionsBetaAPI.md#PostRewardProgramRedemption) | **Post** /credit/rewardprograms/{token}/redemptions | Create reward redemption
[**RetrieveRedemptions**](RewardRedemptionsBetaAPI.md#RetrieveRedemptions) | **Get** /credit/rewardprograms/{token}/redemptions | List reward redemptions
[**RetrieveRedemptionsBalance**](RewardRedemptionsBetaAPI.md#RetrieveRedemptionsBalance) | **Get** /credit/rewardprograms/{token}/redemptions/balance | Retrieve reward redemption balance



## GetRedemption

> RedemptionsResponse GetRedemption(ctx, token, redemptionToken).Execute()

Retrieve reward redemption



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
    token := "token_example" // string | Unique identifier of the reward program.
    redemptionToken := "redemptionToken_example" // string | Unique identifier of the reward redemption.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardRedemptionsBetaAPI.GetRedemption(context.Background(), token, redemptionToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardRedemptionsBetaAPI.GetRedemption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRedemption`: RedemptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardRedemptionsBetaAPI.GetRedemption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 
**redemptionToken** | **string** | Unique identifier of the reward redemption. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedemptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RedemptionsResponse**](RedemptionsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRedemptionTransition

> RedemptionTransitionsResponse PostRedemptionTransition(ctx, token, redemptionToken).CreateRedemptionTransitionsRequest(createRedemptionTransitionsRequest).Execute()

Transition reward redemption status



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
    token := "token_example" // string | Unique identifier of the reward program.
    redemptionToken := "redemptionToken_example" // string | Unique identifier of the reward redemption.
    createRedemptionTransitionsRequest := *openapiclient.NewCreateRedemptionTransitionsRequest(openapiclient.RedemptionStatus("INITIATED")) // CreateRedemptionTransitionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardRedemptionsBetaAPI.PostRedemptionTransition(context.Background(), token, redemptionToken).CreateRedemptionTransitionsRequest(createRedemptionTransitionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardRedemptionsBetaAPI.PostRedemptionTransition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRedemptionTransition`: RedemptionTransitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardRedemptionsBetaAPI.PostRedemptionTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 
**redemptionToken** | **string** | Unique identifier of the reward redemption. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRedemptionTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRedemptionTransitionsRequest** | [**CreateRedemptionTransitionsRequest**](CreateRedemptionTransitionsRequest.md) |  | 

### Return type

[**RedemptionTransitionsResponse**](RedemptionTransitionsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRewardProgramRedemption

> RedemptionsResponse PostRewardProgramRedemption(ctx, token).CreateRedemptionsRequest(createRedemptionsRequest).Execute()

Create reward redemption



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
    token := "token_example" // string | Unique identifier of the reward program.
    createRedemptionsRequest := *openapiclient.NewCreateRedemptionsRequest(float32(123), openapiclient.RedemptionType("EXTERNAL")) // CreateRedemptionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardRedemptionsBetaAPI.PostRewardProgramRedemption(context.Background(), token).CreateRedemptionsRequest(createRedemptionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardRedemptionsBetaAPI.PostRewardProgramRedemption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRewardProgramRedemption`: RedemptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardRedemptionsBetaAPI.PostRewardProgramRedemption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardProgramRedemptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRedemptionsRequest** | [**CreateRedemptionsRequest**](CreateRedemptionsRequest.md) |  | 

### Return type

[**RedemptionsResponse**](RedemptionsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRedemptions

> RedemptionsPage RetrieveRedemptions(ctx, token).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Type_(type_).Execute()

List reward redemptions



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
    token := "token_example" // string | Unique identifier of the reward program.
    startDate := time.Now() // time.Time | The starting date (or date-time) of a date range from which to return resources, in UTC. (optional)
    endDate := time.Now() // time.Time | The ending date (or date-time) of a date range from which to return resources, in UTC. (optional)
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int64(789) // int64 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as `createdTime`, and not by the field names appearing in response bodies such as `created_time`. (optional) (default to "-createdTime")
    type_ := openapiclient.RedemptionType("EXTERNAL") // RedemptionType | Type of reward redemptions in the returned array. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardRedemptionsBetaAPI.RetrieveRedemptions(context.Background(), token).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardRedemptionsBetaAPI.RetrieveRedemptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRedemptions`: RedemptionsPage
    fmt.Fprintf(os.Stdout, "Response from `RewardRedemptionsBetaAPI.RetrieveRedemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRedemptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | The starting date (or date-time) of a date range from which to return resources, in UTC. | 
 **endDate** | **time.Time** | The ending date (or date-time) of a date range from which to return resources, in UTC. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int64** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]
 **type_** | [**RedemptionType**](RedemptionType.md) | Type of reward redemptions in the returned array. | 

### Return type

[**RedemptionsPage**](RedemptionsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRedemptionsBalance

> RedemptionsBalanceResponse RetrieveRedemptionsBalance(ctx, token).StartDate(startDate).EndDate(endDate).Type_(type_).Execute()

Retrieve reward redemption balance



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
    token := "token_example" // string | Unique identifier of the reward program.
    startDate := time.Now() // time.Time | The starting date (or date-time) of a date range from which to return resources, in UTC. (optional)
    endDate := time.Now() // time.Time | The ending date (or date-time) of a date range from which to return resources, in UTC. (optional)
    type_ := []openapiclient.RedemptionType{openapiclient.RedemptionType("EXTERNAL")} // []RedemptionType | Type of reward redemptions in the returned array. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardRedemptionsBetaAPI.RetrieveRedemptionsBalance(context.Background(), token).StartDate(startDate).EndDate(endDate).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardRedemptionsBetaAPI.RetrieveRedemptionsBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRedemptionsBalance`: RedemptionsBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardRedemptionsBetaAPI.RetrieveRedemptionsBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRedemptionsBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | The starting date (or date-time) of a date range from which to return resources, in UTC. | 
 **endDate** | **time.Time** | The ending date (or date-time) of a date range from which to return resources, in UTC. | 
 **type_** | [**[]RedemptionType**](RedemptionType.md) | Type of reward redemptions in the returned array. | 

### Return type

[**RedemptionsBalanceResponse**](RedemptionsBalanceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

