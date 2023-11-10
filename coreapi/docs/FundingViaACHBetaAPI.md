# \FundingViaACHBetaAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBanktransfersAch**](FundingViaACHBetaAPI.md#GetBanktransfersAch) | **Get** /banktransfers/ach | List ACH transfers
[**GetBanktransfersAchToken**](FundingViaACHBetaAPI.md#GetBanktransfersAchToken) | **Get** /banktransfers/ach/{token} | Retrieve ACH transfer
[**GetBanktransfersAchTransitions**](FundingViaACHBetaAPI.md#GetBanktransfersAchTransitions) | **Get** /banktransfers/ach/transitions | List ACH transfer transitions
[**PostBanktransfersAch**](FundingViaACHBetaAPI.md#PostBanktransfersAch) | **Post** /banktransfers/ach | Create ACH transfer
[**PostBanktransfersAchTransitions**](FundingViaACHBetaAPI.md#PostBanktransfersAchTransitions) | **Post** /banktransfers/ach/transitions | Create ACH transfer transition



## GetBanktransfersAch

> BankTransferListResponse GetBanktransfersAch(ctx).Count(count).StartIndex(startIndex).UserToken(userToken).BusinessToken(businessToken).FundingSourceToken(fundingSourceToken).Statuses(statuses).SortBy(sortBy).Expand(expand).FundingSourceType(fundingSourceType).Execute()

List ACH transfers



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
    userToken := "userToken_example" // string | Unique identifier of the user resource. (optional)
    businessToken := "businessToken_example" // string | Unique identifier of the business resource. (optional)
    fundingSourceToken := "fundingSourceToken_example" // string | Unique identifier of the funding source. (optional)
    statuses := "statuses_example" // string | Comma-delimited list of bank transfer statuses. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")
    expand := "expand_example" // string | Returns the full funding source object when `fundingsource` is passed. Otherwise, returns the funding source token. (optional)
    fundingSourceType := "fundingSourceType_example" // string | Funding source type to filter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingViaACHBetaAPI.GetBanktransfersAch(context.Background()).Count(count).StartIndex(startIndex).UserToken(userToken).BusinessToken(businessToken).FundingSourceToken(fundingSourceToken).Statuses(statuses).SortBy(sortBy).Expand(expand).FundingSourceType(fundingSourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingViaACHBetaAPI.GetBanktransfersAch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBanktransfersAch`: BankTransferListResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingViaACHBetaAPI.GetBanktransfersAch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBanktransfersAchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **userToken** | **string** | Unique identifier of the user resource. | 
 **businessToken** | **string** | Unique identifier of the business resource. | 
 **fundingSourceToken** | **string** | Unique identifier of the funding source. | 
 **statuses** | **string** | Comma-delimited list of bank transfer statuses. | 
 **sortBy** | **string** | Field on which to sort Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]
 **expand** | **string** | Returns the full funding source object when &#x60;fundingsource&#x60; is passed. Otherwise, returns the funding source token. | 
 **fundingSourceType** | **string** | Funding source type to filter. | 

### Return type

[**BankTransferListResponse**](BankTransferListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanktransfersAchToken

> BankTransferResponseModel GetBanktransfersAchToken(ctx, token).Execute()

Retrieve ACH transfer



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
    token := "token_example" // string | Unique identifier of the bank transfer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingViaACHBetaAPI.GetBanktransfersAchToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingViaACHBetaAPI.GetBanktransfersAchToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBanktransfersAchToken`: BankTransferResponseModel
    fmt.Fprintf(os.Stdout, "Response from `FundingViaACHBetaAPI.GetBanktransfersAchToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the bank transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBanktransfersAchTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BankTransferResponseModel**](BankTransferResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanktransfersAchTransitions

> BankTransferTransitionListResponse GetBanktransfersAchTransitions(ctx).Count(count).Token(token).BankTransferToken(bankTransferToken).StartIndex(startIndex).SortBy(sortBy).Statuses(statuses).Execute()

List ACH transfer transitions



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
    count := int32(56) // int32 | Number of bank transfer transitions to retrieve. (optional) (default to 5)
    token := "token_example" // string | Unique identifier of the bank transfer transition. (optional)
    bankTransferToken := "bankTransferToken_example" // string | Unique identifier of the bank transfer. (optional)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields lastModifiedTime or createdTime. Prefix the field name with a hyphen (-) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")
    statuses := "statuses_example" // string | Comma-delimited list of bank transfer states to display. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingViaACHBetaAPI.GetBanktransfersAchTransitions(context.Background()).Count(count).Token(token).BankTransferToken(bankTransferToken).StartIndex(startIndex).SortBy(sortBy).Statuses(statuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingViaACHBetaAPI.GetBanktransfersAchTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBanktransfersAchTransitions`: BankTransferTransitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingViaACHBetaAPI.GetBanktransfersAchTransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBanktransfersAchTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of bank transfer transitions to retrieve. | [default to 5]
 **token** | **string** | Unique identifier of the bank transfer transition. | 
 **bankTransferToken** | **string** | Unique identifier of the bank transfer. | 
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields lastModifiedTime or createdTime. Prefix the field name with a hyphen (-) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]
 **statuses** | **string** | Comma-delimited list of bank transfer states to display. | 

### Return type

[**BankTransferTransitionListResponse**](BankTransferTransitionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBanktransfersAch

> BankTransferResponseModel PostBanktransfersAch(ctx).BankTransferRequestModel(bankTransferRequestModel).Execute()

Create ACH transfer



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
    bankTransferRequestModel := *openapiclient.NewBankTransferRequestModel(float32(123), "FundingSourceToken_example", "Type_example") // BankTransferRequestModel | Create bank transfer request model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingViaACHBetaAPI.PostBanktransfersAch(context.Background()).BankTransferRequestModel(bankTransferRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingViaACHBetaAPI.PostBanktransfersAch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBanktransfersAch`: BankTransferResponseModel
    fmt.Fprintf(os.Stdout, "Response from `FundingViaACHBetaAPI.PostBanktransfersAch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBanktransfersAchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferRequestModel** | [**BankTransferRequestModel**](BankTransferRequestModel.md) | Create bank transfer request model | 

### Return type

[**BankTransferResponseModel**](BankTransferResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBanktransfersAchTransitions

> BankTransferTransitionResponseModel PostBanktransfersAchTransitions(ctx).BankTransferTransitionRequestModel(bankTransferTransitionRequestModel).Execute()

Create ACH transfer transition



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
    bankTransferTransitionRequestModel := *openapiclient.NewBankTransferTransitionRequestModel("BankTransferToken_example", "Channel_example", "Status_example") // BankTransferTransitionRequestModel | Create bank transfer transition request model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingViaACHBetaAPI.PostBanktransfersAchTransitions(context.Background()).BankTransferTransitionRequestModel(bankTransferTransitionRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingViaACHBetaAPI.PostBanktransfersAchTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBanktransfersAchTransitions`: BankTransferTransitionResponseModel
    fmt.Fprintf(os.Stdout, "Response from `FundingViaACHBetaAPI.PostBanktransfersAchTransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBanktransfersAchTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferTransitionRequestModel** | [**BankTransferTransitionRequestModel**](BankTransferTransitionRequestModel.md) | Create bank transfer transition request model | 

### Return type

[**BankTransferTransitionResponseModel**](BankTransferTransitionResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

