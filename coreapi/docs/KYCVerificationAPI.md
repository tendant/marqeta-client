# \KYCVerificationAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKycBusinessBusinesstoken**](KYCVerificationAPI.md#GetKycBusinessBusinesstoken) | **Get** /kyc/business/{business_token} | List KYC results for a business
[**GetKycToken**](KYCVerificationAPI.md#GetKycToken) | **Get** /kyc/{token} | Retrieve KYC result
[**GetKycUserUsertoken**](KYCVerificationAPI.md#GetKycUserUsertoken) | **Get** /kyc/user/{user_token} | List KYC results for a user
[**PostKyc**](KYCVerificationAPI.md#PostKyc) | **Post** /kyc | Perform KYC verification



## GetKycBusinessBusinesstoken

> KYCListResponse GetKycBusinessBusinesstoken(ctx, businessToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List KYC results for a business



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
    businessToken := "businessToken_example" // string | The unique identifier of the business resource for which you want to retrieve KYC verification results.
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | The sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationAPI.GetKycBusinessBusinesstoken(context.Background(), businessToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationAPI.GetKycBusinessBusinesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKycBusinessBusinesstoken`: KYCListResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationAPI.GetKycBusinessBusinesstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessToken** | **string** | The unique identifier of the business resource for which you want to retrieve KYC verification results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKycBusinessBusinesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | The sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**KYCListResponse**](KYCListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKycToken

> KycResponse GetKycToken(ctx, token).Execute()

Retrieve KYC result



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
    token := "token_example" // string | Unique identifier of the KYC verification for which you want to retrieve the result.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationAPI.GetKycToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationAPI.GetKycToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKycToken`: KycResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationAPI.GetKycToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the KYC verification for which you want to retrieve the result. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKycTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KycResponse**](KycResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKycUserUsertoken

> KYCListResponse GetKycUserUsertoken(ctx, userToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List KYC results for a user



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
    userToken := "userToken_example" // string | Unique identifier of the user resource for which you want to retrieve KYC verification results.
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationAPI.GetKycUserUsertoken(context.Background(), userToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationAPI.GetKycUserUsertoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKycUserUsertoken`: KYCListResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationAPI.GetKycUserUsertoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userToken** | **string** | Unique identifier of the user resource for which you want to retrieve KYC verification results. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKycUserUsertokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**KYCListResponse**](KYCListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKyc

> KycResponse PostKyc(ctx).KycRequest(kycRequest).Execute()

Perform KYC verification



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
    kycRequest := *openapiclient.NewKycRequest() // KycRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KYCVerificationAPI.PostKyc(context.Background()).KycRequest(kycRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationAPI.PostKyc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostKyc`: KycResponse
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationAPI.PostKyc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKycRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kycRequest** | [**KycRequest**](KycRequest.md) |  | 

### Return type

[**KycResponse**](KycResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

