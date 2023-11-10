# \ProgramFundingSourcesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllACHFundingSources**](ProgramFundingSourcesAPI.md#GetAllACHFundingSources) | **Get** /fundingsources/program/ach | List ACH program sources
[**GetFundingsourcesProgramToken**](ProgramFundingSourcesAPI.md#GetFundingsourcesProgramToken) | **Get** /fundingsources/program/{token} | Retrieve program source
[**PostFundingsourcesProgram**](ProgramFundingSourcesAPI.md#PostFundingsourcesProgram) | **Post** /fundingsources/program | Create program source
[**PostFundingsourcesProgramAch**](ProgramFundingSourcesAPI.md#PostFundingsourcesProgramAch) | **Post** /fundingsources/program/ach | Create ACH program source
[**PutFundingsourcesProgramToken**](ProgramFundingSourcesAPI.md#PutFundingsourcesProgramToken) | **Put** /fundingsources/program/{token} | Update program source



## GetAllACHFundingSources

> ACHListResponse GetAllACHFundingSources(ctx).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Active(active).Execute()

List ACH program sources



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
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")
    active := true // bool | If `true`, returns ACH funding sources from active programs only. If `false`, returns all ACH funding sources. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramFundingSourcesAPI.GetAllACHFundingSources(context.Background()).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Active(active).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramFundingSourcesAPI.GetAllACHFundingSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllACHFundingSources`: ACHListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramFundingSourcesAPI.GetAllACHFundingSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllACHFundingSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]
 **active** | **bool** | If &#x60;true&#x60;, returns ACH funding sources from active programs only. If &#x60;false&#x60;, returns all ACH funding sources. | 

### Return type

[**ACHListResponse**](ACHListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsourcesProgramToken

> ProgramFundingSourceResponse GetFundingsourcesProgramToken(ctx, token).Execute()

Retrieve program source



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
    token := "token_example" // string | Unique identifier of the program funding source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramFundingSourcesAPI.GetFundingsourcesProgramToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramFundingSourcesAPI.GetFundingsourcesProgramToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesProgramToken`: ProgramFundingSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramFundingSourcesAPI.GetFundingsourcesProgramToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the program funding source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesProgramTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramFundingSourceResponse**](ProgramFundingSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFundingsourcesProgram

> ProgramFundingSourceResponse PostFundingsourcesProgram(ctx).ProgramFundingSourceRequest(programFundingSourceRequest).Execute()

Create program source



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
    programFundingSourceRequest := *openapiclient.NewProgramFundingSourceRequest("Name_example") // ProgramFundingSourceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramFundingSourcesAPI.PostFundingsourcesProgram(context.Background()).ProgramFundingSourceRequest(programFundingSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramFundingSourcesAPI.PostFundingsourcesProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFundingsourcesProgram`: ProgramFundingSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramFundingSourcesAPI.PostFundingsourcesProgram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFundingsourcesProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programFundingSourceRequest** | [**ProgramFundingSourceRequest**](ProgramFundingSourceRequest.md) |  | 

### Return type

[**ProgramFundingSourceResponse**](ProgramFundingSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFundingsourcesProgramAch

> BaseAchResponseModel PostFundingsourcesProgramAch(ctx).BaseAchRequestModel(baseAchRequestModel).Execute()

Create ACH program source



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
    baseAchRequestModel := *openapiclient.NewBaseAchRequestModel("AccountNumber_example", "AccountType_example", "NameOnAccount_example", "RoutingNumber_example") // BaseAchRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramFundingSourcesAPI.PostFundingsourcesProgramAch(context.Background()).BaseAchRequestModel(baseAchRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramFundingSourcesAPI.PostFundingsourcesProgramAch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFundingsourcesProgramAch`: BaseAchResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ProgramFundingSourcesAPI.PostFundingsourcesProgramAch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFundingsourcesProgramAchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseAchRequestModel** | [**BaseAchRequestModel**](BaseAchRequestModel.md) |  | 

### Return type

[**BaseAchResponseModel**](BaseAchResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFundingsourcesProgramToken

> ProgramFundingSourceResponse PutFundingsourcesProgramToken(ctx, token).ProgramFundingSourceUpdateRequest(programFundingSourceUpdateRequest).Execute()

Update program source



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
    token := "token_example" // string | Unique identifier of the program funding source.
    programFundingSourceUpdateRequest := *openapiclient.NewProgramFundingSourceUpdateRequest() // ProgramFundingSourceUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramFundingSourcesAPI.PutFundingsourcesProgramToken(context.Background(), token).ProgramFundingSourceUpdateRequest(programFundingSourceUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramFundingSourcesAPI.PutFundingsourcesProgramToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFundingsourcesProgramToken`: ProgramFundingSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramFundingSourcesAPI.PutFundingsourcesProgramToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the program funding source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFundingsourcesProgramTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **programFundingSourceUpdateRequest** | [**ProgramFundingSourceUpdateRequest**](ProgramFundingSourceUpdateRequest.md) |  | 

### Return type

[**ProgramFundingSourceResponse**](ProgramFundingSourceResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

