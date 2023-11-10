# \AddressesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFundingsourcesAddressesBusinessBusinesstoken**](AddressesAPI.md#GetFundingsourcesAddressesBusinessBusinesstoken) | **Get** /fundingsources/addresses/business/{business_token} | List business addresses
[**GetFundingsourcesAddressesFundingsourceaddresstoken**](AddressesAPI.md#GetFundingsourcesAddressesFundingsourceaddresstoken) | **Get** /fundingsources/addresses/{funding_source_address_token} | Retrieve address
[**GetFundingsourcesAddressesUserUsertoken**](AddressesAPI.md#GetFundingsourcesAddressesUserUsertoken) | **Get** /fundingsources/addresses/user/{user_token} | Lists all addresses for a user
[**PostFundingsourcesAddresses**](AddressesAPI.md#PostFundingsourcesAddresses) | **Post** /fundingsources/addresses | Create address
[**PutFundingsourcesAddressesFundingsourceaddresstoken**](AddressesAPI.md#PutFundingsourcesAddressesFundingsourceaddresstoken) | **Put** /fundingsources/addresses/{funding_source_address_token} | Update address



## GetFundingsourcesAddressesBusinessBusinesstoken

> CardholderAddressListResponse GetFundingsourcesAddressesBusinessBusinesstoken(ctx, businessToken).Fields(fields).Execute()

List business addresses



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
    businessToken := "businessToken_example" // string | Unique identifier of the business account holder.  Send a `GET` request to `/businesses` to retrieve business tokens.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesAPI.GetFundingsourcesAddressesBusinessBusinesstoken(context.Background(), businessToken).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetFundingsourcesAddressesBusinessBusinesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesAddressesBusinessBusinesstoken`: CardholderAddressListResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetFundingsourcesAddressesBusinessBusinesstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessToken** | **string** | Unique identifier of the business account holder.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesAddressesBusinessBusinesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**CardholderAddressListResponse**](CardholderAddressListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsourcesAddressesFundingsourceaddresstoken

> CardholderAddressResponse GetFundingsourcesAddressesFundingsourceaddresstoken(ctx, fundingSourceAddressToken).Execute()

Retrieve address



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
    fundingSourceAddressToken := "fundingSourceAddressToken_example" // string | Unique identifier of the funding source address.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesAPI.GetFundingsourcesAddressesFundingsourceaddresstoken(context.Background(), fundingSourceAddressToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetFundingsourcesAddressesFundingsourceaddresstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesAddressesFundingsourceaddresstoken`: CardholderAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetFundingsourcesAddressesFundingsourceaddresstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceAddressToken** | **string** | Unique identifier of the funding source address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesAddressesFundingsourceaddresstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardholderAddressResponse**](CardholderAddressResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsourcesAddressesUserUsertoken

> CardholderAddressListResponse GetFundingsourcesAddressesUserUsertoken(ctx, userToken).Fields(fields).Execute()

Lists all addresses for a user



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
    userToken := "userToken_example" // string | Unique identifier of the user account holder.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesAPI.GetFundingsourcesAddressesUserUsertoken(context.Background(), userToken).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetFundingsourcesAddressesUserUsertoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesAddressesUserUsertoken`: CardholderAddressListResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetFundingsourcesAddressesUserUsertoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userToken** | **string** | Unique identifier of the user account holder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesAddressesUserUsertokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**CardholderAddressListResponse**](CardholderAddressListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFundingsourcesAddresses

> CardholderAddressResponse PostFundingsourcesAddresses(ctx).CardHolderAddressModel(cardHolderAddressModel).Execute()

Create address



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
    cardHolderAddressModel := *openapiclient.NewCardHolderAddressModel("Address1_example", "City_example", "Country_example", "FirstName_example", "LastName_example", "State_example") // CardHolderAddressModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesAPI.PostFundingsourcesAddresses(context.Background()).CardHolderAddressModel(cardHolderAddressModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.PostFundingsourcesAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFundingsourcesAddresses`: CardholderAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.PostFundingsourcesAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFundingsourcesAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardHolderAddressModel** | [**CardHolderAddressModel**](CardHolderAddressModel.md) |  | 

### Return type

[**CardholderAddressResponse**](CardholderAddressResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFundingsourcesAddressesFundingsourceaddresstoken

> CardholderAddressResponse PutFundingsourcesAddressesFundingsourceaddresstoken(ctx, fundingSourceAddressToken).CardHolderAddressUpdateModel(cardHolderAddressUpdateModel).Execute()

Update address



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
    fundingSourceAddressToken := "fundingSourceAddressToken_example" // string | Unique identifier of the funding source address.
    cardHolderAddressUpdateModel := *openapiclient.NewCardHolderAddressUpdateModel() // CardHolderAddressUpdateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressesAPI.PutFundingsourcesAddressesFundingsourceaddresstoken(context.Background(), fundingSourceAddressToken).CardHolderAddressUpdateModel(cardHolderAddressUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.PutFundingsourcesAddressesFundingsourceaddresstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFundingsourcesAddressesFundingsourceaddresstoken`: CardholderAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.PutFundingsourcesAddressesFundingsourceaddresstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceAddressToken** | **string** | Unique identifier of the funding source address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFundingsourcesAddressesFundingsourceaddresstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardHolderAddressUpdateModel** | [**CardHolderAddressUpdateModel**](CardHolderAddressUpdateModel.md) |  | 

### Return type

[**CardholderAddressResponse**](CardholderAddressResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

