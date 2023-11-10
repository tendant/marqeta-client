# \RewardProgramsBetaAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRewardProgramEntry**](RewardProgramsBetaAPI.md#PostRewardProgramEntry) | **Post** /credit/rewardprograms/{token}/entries | Create reward entry
[**RetrieveRewardProgram**](RewardProgramsBetaAPI.md#RetrieveRewardProgram) | **Get** /credit/rewardprograms/{token} | Retrieve reward program
[**RetrieveRewardProgramBalance**](RewardProgramsBetaAPI.md#RetrieveRewardProgramBalance) | **Get** /credit/rewardprograms/{token}/balances | Retrieve reward program balances
[**RetrieveRewardProgramEntries**](RewardProgramsBetaAPI.md#RetrieveRewardProgramEntries) | **Get** /credit/rewardprograms/{token}/entries | List reward entries
[**RetrieveRewardProgramEntriesBalance**](RewardProgramsBetaAPI.md#RetrieveRewardProgramEntriesBalance) | **Get** /credit/rewardprograms/{token}/entries/balance | Retrieve reward entries balance
[**RetrieveRewardProgramEntry**](RewardProgramsBetaAPI.md#RetrieveRewardProgramEntry) | **Get** /credit/rewardprograms/{token}/entries/{entry_token} | Retrieve reward entry
[**RetrieveRewardPrograms**](RewardProgramsBetaAPI.md#RetrieveRewardPrograms) | **Get** /credit/rewardprograms | List reward programs
[**RetrieveRewardProgramsRulesConfig**](RewardProgramsBetaAPI.md#RetrieveRewardProgramsRulesConfig) | **Get** /credit/rewardprograms/{token}/rulesconfigs | List rules configurations
[**RetrieveRewardProgramsRulesConfigApplied**](RewardProgramsBetaAPI.md#RetrieveRewardProgramsRulesConfigApplied) | **Get** /credit/rewardprograms/{token}/rulesconfigs/applied | Retrieve last rules configuration applied
[**UpdateRewardProgram**](RewardProgramsBetaAPI.md#UpdateRewardProgram) | **Put** /credit/rewardprograms/{token} | Activate or deactivate reward program



## PostRewardProgramEntry

> RewardProgramsEntriesResponse PostRewardProgramEntry(ctx, token).CreateRewardProgramsEntriesRequest(createRewardProgramsEntriesRequest).Execute()

Create reward entry



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
    createRewardProgramsEntriesRequest := *openapiclient.NewCreateRewardProgramsEntriesRequest("Note_example", float32(123)) // CreateRewardProgramsEntriesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.PostRewardProgramEntry(context.Background(), token).CreateRewardProgramsEntriesRequest(createRewardProgramsEntriesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.PostRewardProgramEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRewardProgramEntry`: RewardProgramsEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.PostRewardProgramEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardProgramEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRewardProgramsEntriesRequest** | [**CreateRewardProgramsEntriesRequest**](CreateRewardProgramsEntriesRequest.md) |  | 

### Return type

[**RewardProgramsEntriesResponse**](RewardProgramsEntriesResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardProgram

> RewardProgramsResponse RetrieveRewardProgram(ctx, token).Execute()

Retrieve reward program



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgram(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardProgram`: RewardProgramsResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardProgram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RewardProgramsResponse**](RewardProgramsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardProgramBalance

> RewardProgramsBalancesResponse RetrieveRewardProgramBalance(ctx, token).Execute()

Retrieve reward program balances



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramBalance(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardProgramBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardProgramBalance`: RewardProgramsBalancesResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardProgramBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RewardProgramsBalancesResponse**](RewardProgramsBalancesResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardProgramEntries

> RewardProgramsEntriesPage RetrieveRewardProgramEntries(ctx, token).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Status(status).Execute()

List reward entries



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
    status := []openapiclient.RewardEntryStatus{openapiclient.RewardEntryStatus("PENDING")} // []RewardEntryStatus | Status of the reward entries in the returned array. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramEntries(context.Background(), token).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardProgramEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardProgramEntries`: RewardProgramsEntriesPage
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardProgramEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | The starting date (or date-time) of a date range from which to return resources, in UTC. | 
 **endDate** | **time.Time** | The ending date (or date-time) of a date range from which to return resources, in UTC. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int64** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]
 **status** | [**[]RewardEntryStatus**](RewardEntryStatus.md) | Status of the reward entries in the returned array. | 

### Return type

[**RewardProgramsEntriesPage**](RewardProgramsEntriesPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardProgramEntriesBalance

> RewardProgramsEntriesBalanceResponse RetrieveRewardProgramEntriesBalance(ctx, token).StartDate(startDate).EndDate(endDate).Execute()

Retrieve reward entries balance



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
    startDate := time.Now() // time.Time | The starting date (or date-time) of a date range from which to return resources, in UTC.
    endDate := time.Now() // time.Time | The ending date (or date-time) of a date range from which to return resources, in UTC.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramEntriesBalance(context.Background(), token).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardProgramEntriesBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardProgramEntriesBalance`: RewardProgramsEntriesBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardProgramEntriesBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramEntriesBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | The starting date (or date-time) of a date range from which to return resources, in UTC. | 
 **endDate** | **time.Time** | The ending date (or date-time) of a date range from which to return resources, in UTC. | 

### Return type

[**RewardProgramsEntriesBalanceResponse**](RewardProgramsEntriesBalanceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardProgramEntry

> RewardProgramsEntriesResponse RetrieveRewardProgramEntry(ctx, token, entryToken).Execute()

Retrieve reward entry



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
    entryToken := "entryToken_example" // string | Unique identifier of the reward entry to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramEntry(context.Background(), token, entryToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardProgramEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardProgramEntry`: RewardProgramsEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardProgramEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 
**entryToken** | **string** | Unique identifier of the reward entry to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RewardProgramsEntriesResponse**](RewardProgramsEntriesResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardPrograms

> RewardProgramsPageResponse RetrieveRewardPrograms(ctx).AccountToken(accountToken).IsActive(isActive).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List reward programs



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve reward programs. (optional)
    isActive := true // bool | A value of `true` returns active resources; `false` returns inactive resources. (optional)
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int64(789) // int64 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as `updatedTime`, and not by the field names appearing in response bodies such as `updated_time`. (optional) (default to "-updatedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardPrograms(context.Background()).AccountToken(accountToken).IsActive(isActive).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardPrograms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardPrograms`: RewardProgramsPageResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardPrograms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve reward programs. | 
 **isActive** | **bool** | A value of &#x60;true&#x60; returns active resources; &#x60;false&#x60; returns inactive resources. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int64** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as &#x60;updatedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;updated_time&#x60;. | [default to &quot;-updatedTime&quot;]

### Return type

[**RewardProgramsPageResponse**](RewardProgramsPageResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardProgramsRulesConfig

> RewardProgramsRulesConfigsPage RetrieveRewardProgramsRulesConfig(ctx, token).IsActive(isActive).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List rules configurations



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
    isActive := true // bool | A value of `true` returns active resources; `false` returns inactive resources. (optional)
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int64(789) // int64 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as `updatedTime`, and not by the field names appearing in response bodies such as `updated_time`. (optional) (default to "-updatedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfig(context.Background(), token).IsActive(isActive).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardProgramsRulesConfig`: RewardProgramsRulesConfigsPage
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramsRulesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isActive** | **bool** | A value of &#x60;true&#x60; returns active resources; &#x60;false&#x60; returns inactive resources. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int64** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE*: You must sort using system field names such as &#x60;updatedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;updated_time&#x60;. | [default to &quot;-updatedTime&quot;]

### Return type

[**RewardProgramsRulesConfigsPage**](RewardProgramsRulesConfigsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardProgramsRulesConfigApplied

> RewardProgramsRulesConfigsResponse RetrieveRewardProgramsRulesConfigApplied(ctx, token).Execute()

Retrieve last rules configuration applied



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfigApplied(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfigApplied``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardProgramsRulesConfigApplied`: RewardProgramsRulesConfigsResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.RetrieveRewardProgramsRulesConfigApplied`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardProgramsRulesConfigAppliedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RewardProgramsRulesConfigsResponse**](RewardProgramsRulesConfigsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRewardProgram

> RewardProgramsResponse UpdateRewardProgram(ctx, token).PutRewardProgramsRequest(putRewardProgramsRequest).Execute()

Activate or deactivate reward program



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
    putRewardProgramsRequest := *openapiclient.NewPutRewardProgramsRequest(false, "Note_example") // PutRewardProgramsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RewardProgramsBetaAPI.UpdateRewardProgram(context.Background(), token).PutRewardProgramsRequest(putRewardProgramsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RewardProgramsBetaAPI.UpdateRewardProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRewardProgram`: RewardProgramsResponse
    fmt.Fprintf(os.Stdout, "Response from `RewardProgramsBetaAPI.UpdateRewardProgram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the reward program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRewardProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putRewardProgramsRequest** | [**PutRewardProgramsRequest**](PutRewardProgramsRequest.md) |  | 

### Return type

[**RewardProgramsResponse**](RewardProgramsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

