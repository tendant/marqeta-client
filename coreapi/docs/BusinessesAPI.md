# \BusinessesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBusinesses**](BusinessesAPI.md#GetBusinesses) | **Get** /businesses | List businesses
[**GetBusinessesParenttokenChildren**](BusinessesAPI.md#GetBusinessesParenttokenChildren) | **Get** /businesses/{parent_token}/children | List business children
[**GetBusinessesToken**](BusinessesAPI.md#GetBusinessesToken) | **Get** /businesses/{token} | Retrieve business
[**GetBusinessesTokenSsn**](BusinessesAPI.md#GetBusinessesTokenSsn) | **Get** /businesses/{token}/ssn | Retrieve business identification number
[**PostBusinesses**](BusinessesAPI.md#PostBusinesses) | **Post** /businesses | Create business
[**PostBusinessesLookup**](BusinessesAPI.md#PostBusinessesLookup) | **Post** /businesses/lookup | Search businesses
[**PutBusinessesToken**](BusinessesAPI.md#PutBusinessesToken) | **Put** /businesses/{token} | Update business



## GetBusinesses

> BusinessCardHolderListResponse GetBusinesses(ctx).Count(count).StartIndex(startIndex).BusinessNameDba(businessNameDba).BusinessNameLegal(businessNameLegal).SearchType(searchType).Fields(fields).SortBy(sortBy).Execute()

List businesses



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
    count := int32(56) // int32 | Number of business resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    businessNameDba := "businessNameDba_example" // string | Fictitious or \"doing business as (DBA)\" name of the business. (optional)
    businessNameLegal := "businessNameLegal_example" // string | Legal name of the business. (optional)
    searchType := "searchType_example" // string | Specifies the search type for the query. (optional)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessesAPI.GetBusinesses(context.Background()).Count(count).StartIndex(startIndex).BusinessNameDba(businessNameDba).BusinessNameLegal(businessNameLegal).SearchType(searchType).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessesAPI.GetBusinesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinesses`: BusinessCardHolderListResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessesAPI.GetBusinesses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of business resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **businessNameDba** | **string** | Fictitious or \&quot;doing business as (DBA)\&quot; name of the business. | 
 **businessNameLegal** | **string** | Legal name of the business. | 
 **searchType** | **string** | Specifies the search type for the query. | 
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**BusinessCardHolderListResponse**](BusinessCardHolderListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessesParenttokenChildren

> BusinessUserCardHolderListResponse GetBusinessesParenttokenChildren(ctx, parentToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List business children



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
    parentToken := "parentToken_example" // string | Unique identifier of the parent business.
    count := int32(56) // int32 | Number of child cardholders to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessesAPI.GetBusinessesParenttokenChildren(context.Background(), parentToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessesAPI.GetBusinessesParenttokenChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessesParenttokenChildren`: BusinessUserCardHolderListResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessesAPI.GetBusinessesParenttokenChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentToken** | **string** | Unique identifier of the parent business. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessesParenttokenChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of child cardholders to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**BusinessUserCardHolderListResponse**](BusinessUserCardHolderListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessesToken

> BusinessCardHolderResponse GetBusinessesToken(ctx, token).Fields(fields).Execute()

Retrieve business



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
    token := "token_example" // string | Unique identifier of the business resource.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessesAPI.GetBusinessesToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessesAPI.GetBusinessesToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessesToken`: BusinessCardHolderResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessesAPI.GetBusinessesToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the business resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessesTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**BusinessCardHolderResponse**](BusinessCardHolderResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessesTokenSsn

> SsnResponseModel GetBusinessesTokenSsn(ctx, token).FullSsn(fullSsn).Execute()

Retrieve business identification number



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
    token := "token_example" // string | Unique identifier of the business resource.
    fullSsn := true // bool | To return the full identification number, set to `true`. To return only the last four digits, set to `false`. If the `proprietor_or_officer.identifications` array contains only the last four digits of the identification number, the `/businesses/{token}/ssn` endpoint will return only the last four digits regardless of the `full_ssn` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessesAPI.GetBusinessesTokenSsn(context.Background(), token).FullSsn(fullSsn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessesAPI.GetBusinessesTokenSsn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessesTokenSsn`: SsnResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BusinessesAPI.GetBusinessesTokenSsn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the business resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessesTokenSsnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullSsn** | **bool** | To return the full identification number, set to &#x60;true&#x60;. To return only the last four digits, set to &#x60;false&#x60;. If the &#x60;proprietor_or_officer.identifications&#x60; array contains only the last four digits of the identification number, the &#x60;/businesses/{token}/ssn&#x60; endpoint will return only the last four digits regardless of the &#x60;full_ssn&#x60; parameter. | 

### Return type

[**SsnResponseModel**](SsnResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBusinesses

> BusinessCardHolderResponse PostBusinesses(ctx).BusinessCardholder(businessCardholder).Execute()

Create business



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
    businessCardholder := *openapiclient.NewBusinessCardholder() // BusinessCardholder |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessesAPI.PostBusinesses(context.Background()).BusinessCardholder(businessCardholder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessesAPI.PostBusinesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBusinesses`: BusinessCardHolderResponse
    fmt.Fprintf(os.Stdout, "Response from `BusinessesAPI.PostBusinesses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBusinessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessCardholder** | [**BusinessCardholder**](BusinessCardholder.md) |  | 

### Return type

[**BusinessCardHolderResponse**](BusinessCardHolderResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBusinessesLookup

> BusinessCardholder PostBusinessesLookup(ctx).DDARequest(dDARequest).Execute()

Search businesses



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
    dDARequest := *openapiclient.NewDDARequest("Dda_example") // DDARequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessesAPI.PostBusinessesLookup(context.Background()).DDARequest(dDARequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessesAPI.PostBusinessesLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBusinessesLookup`: BusinessCardholder
    fmt.Fprintf(os.Stdout, "Response from `BusinessesAPI.PostBusinessesLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBusinessesLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dDARequest** | [**DDARequest**](DDARequest.md) |  | 

### Return type

[**BusinessCardholder**](BusinessCardholder.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBusinessesToken

> BusinessCardholder PutBusinessesToken(ctx, token).BusinessCardHolderUpdate(businessCardHolderUpdate).Execute()

Update business



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
    token := "token_example" // string | Unique identifier of the business resource
    businessCardHolderUpdate := *openapiclient.NewBusinessCardHolderUpdate() // BusinessCardHolderUpdate | Business object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessesAPI.PutBusinessesToken(context.Background(), token).BusinessCardHolderUpdate(businessCardHolderUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessesAPI.PutBusinessesToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutBusinessesToken`: BusinessCardholder
    fmt.Fprintf(os.Stdout, "Response from `BusinessesAPI.PutBusinessesToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the business resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBusinessesTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **businessCardHolderUpdate** | [**BusinessCardHolderUpdate**](BusinessCardHolderUpdate.md) | Business object | 

### Return type

[**BusinessCardholder**](BusinessCardholder.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

