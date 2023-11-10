# \AcceptedCountriesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAcceptedcountries**](AcceptedCountriesAPI.md#GetAcceptedcountries) | **Get** /acceptedcountries | List accepted countries objects
[**GetAcceptedcountriesToken**](AcceptedCountriesAPI.md#GetAcceptedcountriesToken) | **Get** /acceptedcountries/{token} | Retrieve an accepted countries object



## GetAcceptedcountries

> AcceptedCountriesListResponse GetAcceptedcountries(ctx).Count(count).StartIndex(startIndex).Name(name).Whitelist(whitelist).SearchType(searchType).Fields(fields).SortBy(sortBy).Execute()

List accepted countries objects



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
    count := int32(56) // int32 | Number of accepted countries objects to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    name := "name_example" // string | Name of the accepted countries object. (optional)
    whitelist := true // bool | Specifies if the accepted countries object is an allow list. (optional)
    searchType := "searchType_example" // string | Specifies the type of search you want to perform. (optional)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on).  Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort.  Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`.  Prefix the field name with a hyphen (`-`) to sort in descending order.  Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcceptedCountriesAPI.GetAcceptedcountries(context.Background()).Count(count).StartIndex(startIndex).Name(name).Whitelist(whitelist).SearchType(searchType).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcceptedCountriesAPI.GetAcceptedcountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcceptedcountries`: AcceptedCountriesListResponse
    fmt.Fprintf(os.Stdout, "Response from `AcceptedCountriesAPI.GetAcceptedcountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAcceptedcountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of accepted countries objects to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **name** | **string** | Name of the accepted countries object. | 
 **whitelist** | **bool** | Specifies if the accepted countries object is an allow list. | 
 **searchType** | **string** | Specifies the type of search you want to perform. | 
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on).  Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort.  Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;.  Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order.  Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**AcceptedCountriesListResponse**](AcceptedCountriesListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcceptedcountriesToken

> AcceptedCountriesModel GetAcceptedcountriesToken(ctx, token).Fields(fields).Execute()

Retrieve an accepted countries object



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
    token := "token_example" // string | Unique identifier of the accepted countries object.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AcceptedCountriesAPI.GetAcceptedcountriesToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AcceptedCountriesAPI.GetAcceptedcountriesToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAcceptedcountriesToken`: AcceptedCountriesModel
    fmt.Fprintf(os.Stdout, "Response from `AcceptedCountriesAPI.GetAcceptedcountriesToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the accepted countries object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcceptedcountriesTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**AcceptedCountriesModel**](AcceptedCountriesModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

