# \ProgramTransfersAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProgramtransfers**](ProgramTransfersAPI.md#GetProgramtransfers) | **Get** /programtransfers | List program transfers
[**GetProgramtransfersToken**](ProgramTransfersAPI.md#GetProgramtransfersToken) | **Get** /programtransfers/{token} | Retrieve program transfer
[**GetProgramtransfersTypes**](ProgramTransfersAPI.md#GetProgramtransfersTypes) | **Get** /programtransfers/types | List program transfer types
[**GetProgramtransfersTypesTypetoken**](ProgramTransfersAPI.md#GetProgramtransfersTypesTypetoken) | **Get** /programtransfers/types/{type_token} | Retrieve program transfer type
[**PostProgramtransfers**](ProgramTransfersAPI.md#PostProgramtransfers) | **Post** /programtransfers | Create program transfer
[**PostProgramtransfersTypes**](ProgramTransfersAPI.md#PostProgramtransfersTypes) | **Post** /programtransfers/types | Create program transfer type
[**PutProgramtransfersTypesTypetoken**](ProgramTransfersAPI.md#PutProgramtransfersTypesTypetoken) | **Put** /programtransfers/types/{type_token} | Update program transfer type



## GetProgramtransfers

> ProgramTransferListResponse GetProgramtransfers(ctx).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).UserToken(userToken).BusinessToken(businessToken).TypeToken(typeToken).Execute()

List program transfers



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
    count := int32(56) // int32 | Number of program transfers to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")
    userToken := "userToken_example" // string | Unique identifier of the user account holder whose program transfers you want to retrieve.  Send a `GET` request to `/users` to retrieve user tokens. (optional)
    businessToken := "businessToken_example" // string | Unique identifier of the business account holder whose program transfers you want to retrieve.  Send a `GET` request to `/businesses` to retrieve business tokens. (optional)
    typeToken := "typeToken_example" // string | Unique identifier of the program transfer type to retrieve. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramTransfersAPI.GetProgramtransfers(context.Background()).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).UserToken(userToken).BusinessToken(businessToken).TypeToken(typeToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramTransfersAPI.GetProgramtransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProgramtransfers`: ProgramTransferListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramTransfersAPI.GetProgramtransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramtransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of program transfers to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]
 **userToken** | **string** | Unique identifier of the user account holder whose program transfers you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | 
 **businessToken** | **string** | Unique identifier of the business account holder whose program transfers you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | 
 **typeToken** | **string** | Unique identifier of the program transfer type to retrieve. | 

### Return type

[**ProgramTransferListResponse**](ProgramTransferListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramtransfersToken

> ProgramTransferResponse GetProgramtransfersToken(ctx, token).Execute()

Retrieve program transfer



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
    token := "token_example" // string | Unique identifier of the program transfer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramTransfersAPI.GetProgramtransfersToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramTransfersAPI.GetProgramtransfersToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProgramtransfersToken`: ProgramTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramTransfersAPI.GetProgramtransfersToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the program transfer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramtransfersTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramTransferResponse**](ProgramTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramtransfersTypes

> ProgramTransferTypeListResponse GetProgramtransfersTypes(ctx).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List program transfer types



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
    count := int32(56) // int32 | Number of program transfer types to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramTransfersAPI.GetProgramtransfersTypes(context.Background()).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramTransfersAPI.GetProgramtransfersTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProgramtransfersTypes`: ProgramTransferTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramTransfersAPI.GetProgramtransfersTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramtransfersTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of program transfer types to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**ProgramTransferTypeListResponse**](ProgramTransferTypeListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramtransfersTypesTypetoken

> ProgramTransferTypeResponse GetProgramtransfersTypesTypetoken(ctx, typeToken).Execute()

Retrieve program transfer type



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
    typeToken := "typeToken_example" // string | Unique identifier of the program transfer type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramTransfersAPI.GetProgramtransfersTypesTypetoken(context.Background(), typeToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramTransfersAPI.GetProgramtransfersTypesTypetoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProgramtransfersTypesTypetoken`: ProgramTransferTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramTransfersAPI.GetProgramtransfersTypesTypetoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeToken** | **string** | Unique identifier of the program transfer type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramtransfersTypesTypetokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramTransferTypeResponse**](ProgramTransferTypeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProgramtransfers

> ProgramTransferResponse PostProgramtransfers(ctx).ProgramTransfer(programTransfer).Execute()

Create program transfer



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
    programTransfer := *openapiclient.NewProgramTransfer(float32(123), "CurrencyCode_example", "TypeToken_example") // ProgramTransfer |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramTransfersAPI.PostProgramtransfers(context.Background()).ProgramTransfer(programTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramTransfersAPI.PostProgramtransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostProgramtransfers`: ProgramTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramTransfersAPI.PostProgramtransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProgramtransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programTransfer** | [**ProgramTransfer**](ProgramTransfer.md) |  | 

### Return type

[**ProgramTransferResponse**](ProgramTransferResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProgramtransfersTypes

> ProgramTransferTypeResponse PostProgramtransfersTypes(ctx).ProgramTransferTypeRequest(programTransferTypeRequest).Execute()

Create program transfer type



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
    programTransferTypeRequest := *openapiclient.NewProgramTransferTypeRequest("Memo_example", "ProgramFundingSourceToken_example") // ProgramTransferTypeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramTransfersAPI.PostProgramtransfersTypes(context.Background()).ProgramTransferTypeRequest(programTransferTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramTransfersAPI.PostProgramtransfersTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostProgramtransfersTypes`: ProgramTransferTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramTransfersAPI.PostProgramtransfersTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProgramtransfersTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programTransferTypeRequest** | [**ProgramTransferTypeRequest**](ProgramTransferTypeRequest.md) |  | 

### Return type

[**ProgramTransferTypeResponse**](ProgramTransferTypeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProgramtransfersTypesTypetoken

> ProgramTransferTypeResponse PutProgramtransfersTypesTypetoken(ctx, typeToken).ProgramTransferTypeRequest(programTransferTypeRequest).Execute()

Update program transfer type



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
    typeToken := "typeToken_example" // string | Unique identifier of the program transfer type.
    programTransferTypeRequest := *openapiclient.NewProgramTransferTypeRequest("Memo_example", "ProgramFundingSourceToken_example") // ProgramTransferTypeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgramTransfersAPI.PutProgramtransfersTypesTypetoken(context.Background(), typeToken).ProgramTransferTypeRequest(programTransferTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramTransfersAPI.PutProgramtransfersTypesTypetoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutProgramtransfersTypesTypetoken`: ProgramTransferTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProgramTransfersAPI.PutProgramtransfersTypesTypetoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeToken** | **string** | Unique identifier of the program transfer type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProgramtransfersTypesTypetokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **programTransferTypeRequest** | [**ProgramTransferTypeRequest**](ProgramTransferTypeRequest.md) |  | 

### Return type

[**ProgramTransferTypeResponse**](ProgramTransferTypeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

