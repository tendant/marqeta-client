# \CardsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCards**](CardsAPI.md#GetCards) | **Get** /cards | List cards by last 4 digits of PAN
[**GetCardsBarcodeBarcode**](CardsAPI.md#GetCardsBarcodeBarcode) | **Get** /cards/barcode/{barcode} | Retrieve card by barcode
[**GetCardsToken**](CardsAPI.md#GetCardsToken) | **Get** /cards/{token} | Retrieve card
[**GetCardsTokenShowpan**](CardsAPI.md#GetCardsTokenShowpan) | **Get** /cards/{token}/showpan | Show card PAN
[**GetCardsUserToken**](CardsAPI.md#GetCardsUserToken) | **Get** /cards/user/{token} | List cards for user
[**PostCards**](CardsAPI.md#PostCards) | **Post** /cards | Create card
[**PostCardsGetbypan**](CardsAPI.md#PostCardsGetbypan) | **Post** /cards/getbypan | Retrieve card by PAN
[**PutCardsToken**](CardsAPI.md#PutCardsToken) | **Put** /cards/{token} | Update card



## GetCards

> CardListResponse GetCards(ctx).LastFour(lastFour).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List cards by last 4 digits of PAN



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
    lastFour := "lastFour_example" // string | Last four digits of the primary account number (PAN) of the card you want to locate.
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.GetCards(context.Background()).LastFour(lastFour).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCards`: CardListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastFour** | **string** | Last four digits of the primary account number (PAN) of the card you want to locate. | 
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**CardListResponse**](CardListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardsBarcodeBarcode

> CardResponse GetCardsBarcodeBarcode(ctx, barcode).Fields(fields).Execute()

Retrieve card by barcode



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
    barcode := "barcode_example" // string | Barcode of the card to retrieve.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.GetCardsBarcodeBarcode(context.Background(), barcode).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardsBarcodeBarcode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardsBarcodeBarcode`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardsBarcodeBarcode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**barcode** | **string** | Barcode of the card to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardsBarcodeBarcodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardsToken

> CardResponse GetCardsToken(ctx, token).Fields(fields).Expand(expand).Execute()

Retrieve card



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
    token := "token_example" // string | Unique identifier of the card you want to retrieve.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    expand := "expand_example" // string | Embeds the associated object of the specified type into the response, for all `GET /cards` endpoints. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.GetCardsToken(context.Background(), token).Fields(fields).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardsToken`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **expand** | **string** | Embeds the associated object of the specified type into the response, for all &#x60;GET /cards&#x60; endpoints. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardsTokenShowpan

> CardResponse GetCardsTokenShowpan(ctx, token).Fields(fields).ShowCvvNumber(showCvvNumber).Execute()

Show card PAN



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
    token := "token_example" // string | Unique identifier of the card whose primary account number (PAN) you want to retrieve. Send a `GET` request to `/cards` to retrieve card tokens.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    showCvvNumber := true // bool | Set to `true` to show the CVV2 number in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.GetCardsTokenShowpan(context.Background(), token).Fields(fields).ShowCvvNumber(showCvvNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardsTokenShowpan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardsTokenShowpan`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardsTokenShowpan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card whose primary account number (PAN) you want to retrieve. Send a &#x60;GET&#x60; request to &#x60;/cards&#x60; to retrieve card tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardsTokenShowpanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **showCvvNumber** | **bool** | Set to &#x60;true&#x60; to show the CVV2 number in the response. | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardsUserToken

> CardListResponse GetCardsUserToken(ctx, token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List cards for user



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
    token := "token_example" // string | Unique identifier of the user whose cards you want to list. Send a `GET` request to `/users` to retrieve user tokens.
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.GetCardsUserToken(context.Background(), token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.GetCardsUserToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardsUserToken`: CardListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.GetCardsUserToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the user whose cards you want to list. Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardsUserTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**CardListResponse**](CardListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCards

> CardResponse PostCards(ctx).ShowCvvNumber(showCvvNumber).ShowPan(showPan).CardRequest(cardRequest).Execute()

Create card



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
    showCvvNumber := true // bool | Set to `true` to show the CVV2 number in the response. (optional) (default to false)
    showPan := true // bool | Set to `true` to show the full primary account number (PAN) in the response. (optional) (default to false)
    cardRequest := *openapiclient.NewCardRequest("CardProductToken_example", "UserToken_example") // CardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.PostCards(context.Background()).ShowCvvNumber(showCvvNumber).ShowPan(showPan).CardRequest(cardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.PostCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCards`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.PostCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showCvvNumber** | **bool** | Set to &#x60;true&#x60; to show the CVV2 number in the response. | [default to false]
 **showPan** | **bool** | Set to &#x60;true&#x60; to show the full primary account number (PAN) in the response. | [default to false]
 **cardRequest** | [**CardRequest**](CardRequest.md) |  | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCardsGetbypan

> PanResponse PostCardsGetbypan(ctx).PanRequest(panRequest).Execute()

Retrieve card by PAN



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
    panRequest := *openapiclient.NewPanRequest("Pan_example") // PanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.PostCardsGetbypan(context.Background()).PanRequest(panRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.PostCardsGetbypan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCardsGetbypan`: PanResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.PostCardsGetbypan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCardsGetbypanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **panRequest** | [**PanRequest**](PanRequest.md) |  | 

### Return type

[**PanResponse**](PanResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCardsToken

> CardResponse PutCardsToken(ctx, token).CardUpdateRequest(cardUpdateRequest).Execute()

Update card



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
    token := "token_example" // string | Unique identifier of the card you want to update.
    cardUpdateRequest := *openapiclient.NewCardUpdateRequest("Token_example") // CardUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsAPI.PutCardsToken(context.Background(), token).CardUpdateRequest(cardUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsAPI.PutCardsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCardsToken`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsAPI.PutCardsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCardsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardUpdateRequest** | [**CardUpdateRequest**](CardUpdateRequest.md) |  | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

