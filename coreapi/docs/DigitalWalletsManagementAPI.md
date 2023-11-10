# \DigitalWalletsManagementAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateApplePayWPPJWT**](DigitalWalletsManagementAPI.md#GenerateApplePayWPPJWT) | **Post** /digitalwallets/wpp/applePayJWT | Create request for Apple Wallet web push provisioning
[**GetDigitalwallettokens**](DigitalWalletsManagementAPI.md#GetDigitalwallettokens) | **Get** /digitalwallettokens | List digital wallet tokens
[**GetDigitalwallettokensCardCardtoken**](DigitalWalletsManagementAPI.md#GetDigitalwallettokensCardCardtoken) | **Get** /digitalwallettokens/card/{card_token} | List digital wallet tokens for card
[**GetDigitalwallettokensToken**](DigitalWalletsManagementAPI.md#GetDigitalwallettokensToken) | **Get** /digitalwallettokens/{token} | Retrieve digital wallet token
[**GetDigitalwallettokensTokenShowtokenpan**](DigitalWalletsManagementAPI.md#GetDigitalwallettokensTokenShowtokenpan) | **Get** /digitalwallettokens/{token}/showtokenpan | Retrieve digital wallet token PAN
[**GetDigitalwallettokentransitionsDigitalwallettokenToken**](DigitalWalletsManagementAPI.md#GetDigitalwallettokentransitionsDigitalwallettokenToken) | **Get** /digitalwallettokentransitions/digitalwallettoken/{token} | List transitions for digital wallet token
[**GetDigitalwallettokentransitionsToken**](DigitalWalletsManagementAPI.md#GetDigitalwallettokentransitionsToken) | **Get** /digitalwallettokentransitions/{token} | Retrieve digital wallet token transition
[**PostDigitalwalletprovisionrequestsAndroidpay**](DigitalWalletsManagementAPI.md#PostDigitalwalletprovisionrequestsAndroidpay) | **Post** /digitalwalletprovisionrequests/androidpay | Create digital wallet token provisioning request for Google Wallet
[**PostDigitalwalletprovisionrequestsApplepay**](DigitalWalletsManagementAPI.md#PostDigitalwalletprovisionrequestsApplepay) | **Post** /digitalwalletprovisionrequests/applepay | Create digital wallet token provisioning request for Apple Wallet
[**PostDigitalwalletprovisionrequestsSamsungpay**](DigitalWalletsManagementAPI.md#PostDigitalwalletprovisionrequestsSamsungpay) | **Post** /digitalwalletprovisionrequests/samsungpay | Create digital wallet token provisioning request for Samsung Wallet
[**PostDigitalwalletprovisionrequestsXPay**](DigitalWalletsManagementAPI.md#PostDigitalwalletprovisionrequestsXPay) | **Post** /digitalwalletprovisionrequests/xpay | Create digital wallet token provisioning request for XPay
[**PostDigitalwallettokentransitions**](DigitalWalletsManagementAPI.md#PostDigitalwallettokentransitions) | **Post** /digitalwallettokentransitions | Create digital wallet token transition
[**SendOPCDataToGooglePay**](DigitalWalletsManagementAPI.md#SendOPCDataToGooglePay) | **Post** /digitalwallets/wpp/googlePayPushProvisioningNotification | Create request for Google Wallet web push provisioning



## GenerateApplePayWPPJWT

> WebPushProvisioningApplePayJWTResponse GenerateApplePayWPPJWT(ctx).ReqSysId(reqSysId).RequestForApplePayWppJWT(requestForApplePayWppJWT).Execute()

Create request for Apple Wallet web push provisioning



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
    reqSysId := "123d837e-958a-4e9f-bc97-4843ec948123" // string | Random pseudo-unique value used for troubleshooting between multiple parties.
    requestForApplePayWppJWT := *openapiclient.NewRequestForApplePayWppJWT("5855opl9-abcc-4cb7-a330-xyze5799ea5") // RequestForApplePayWppJWT | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.GenerateApplePayWPPJWT(context.Background()).ReqSysId(reqSysId).RequestForApplePayWppJWT(requestForApplePayWppJWT).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.GenerateApplePayWPPJWT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateApplePayWPPJWT`: WebPushProvisioningApplePayJWTResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.GenerateApplePayWPPJWT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateApplePayWPPJWTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reqSysId** | **string** | Random pseudo-unique value used for troubleshooting between multiple parties. | 
 **requestForApplePayWppJWT** | [**RequestForApplePayWppJWT**](RequestForApplePayWppJWT.md) |  | 

### Return type

[**WebPushProvisioningApplePayJWTResponse**](WebPushProvisioningApplePayJWTResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDigitalwallettokens

> DigitalWalletTokenListResponse GetDigitalwallettokens(ctx).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).PanReferenceId(panReferenceId).TokenReferenceId(tokenReferenceId).CorrelationId(correlationId).TokenType(tokenType).TokenRequestorName(tokenRequestorName).State(state).Embed(embed).Execute()

List digital wallet tokens



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
    count := int32(56) // int32 | Number of digital wallet token resources to retrieve. (optional) (default to 10)
    startIndex := int32(56) // int32 | Sort order index of the first digital wallet token resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")
    startDate := "startDate_example" // string | Date when the digital wallet token becomes active. (optional)
    endDate := "endDate_example" // string | Expiration date of the digital wallet token. (optional)
    panReferenceId := "panReferenceId_example" // string | Unique identifier of the digital wallet token primary account number (PAN) within the card network. This value may vary, depending on the digital wallet. For example, the `pan_reference_id` may be different in Apple Wallet and Google Wallet for the same digital wallet token. (optional)
    tokenReferenceId := "tokenReferenceId_example" // string | Unique identifier of the digital wallet token within the card network. The `token_reference_id` is unique at the card network level. (optional)
    correlationId := "correlationId_example" // string | Unique value representing a tokenization request (Mastercard only). (optional)
    tokenType := "tokenType_example" // string | Comma-delimited list of digital wallet token types to display. (optional)
    tokenRequestorName := "tokenRequestorName_example" // string | Name of the token requestor within the card network.  *NOTE:* The list of example values for this field is maintained by the card networks and is subject to change. (optional)
    state := "state_example" // string | Comma-delimited list of digital wallet token states to display. (optional)
    embed := "embed_example" // string | An optional embedded user object. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.GetDigitalwallettokens(context.Background()).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).StartDate(startDate).EndDate(endDate).PanReferenceId(panReferenceId).TokenReferenceId(tokenReferenceId).CorrelationId(correlationId).TokenType(tokenType).TokenRequestorName(tokenRequestorName).State(state).Embed(embed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.GetDigitalwallettokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalwallettokens`: DigitalWalletTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.GetDigitalwallettokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalwallettokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of digital wallet token resources to retrieve. | [default to 10]
 **startIndex** | **int32** | Sort order index of the first digital wallet token resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]
 **startDate** | **string** | Date when the digital wallet token becomes active. | 
 **endDate** | **string** | Expiration date of the digital wallet token. | 
 **panReferenceId** | **string** | Unique identifier of the digital wallet token primary account number (PAN) within the card network. This value may vary, depending on the digital wallet. For example, the &#x60;pan_reference_id&#x60; may be different in Apple Wallet and Google Wallet for the same digital wallet token. | 
 **tokenReferenceId** | **string** | Unique identifier of the digital wallet token within the card network. The &#x60;token_reference_id&#x60; is unique at the card network level. | 
 **correlationId** | **string** | Unique value representing a tokenization request (Mastercard only). | 
 **tokenType** | **string** | Comma-delimited list of digital wallet token types to display. | 
 **tokenRequestorName** | **string** | Name of the token requestor within the card network.  *NOTE:* The list of example values for this field is maintained by the card networks and is subject to change. | 
 **state** | **string** | Comma-delimited list of digital wallet token states to display. | 
 **embed** | **string** | An optional embedded user object. | 

### Return type

[**DigitalWalletTokenListResponse**](DigitalWalletTokenListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDigitalwallettokensCardCardtoken

> DigitalWalletTokenListResponse GetDigitalwallettokensCardCardtoken(ctx, cardToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List digital wallet tokens for card



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
    cardToken := "cardToken_example" // string | Unique identifier of the card. Used to minimize the need to exchange card details during subsequent calls, and also for troubleshooting.
    count := int32(56) // int32 | Number of digital wallet token resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first digital wallet token resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.GetDigitalwallettokensCardCardtoken(context.Background(), cardToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.GetDigitalwallettokensCardCardtoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalwallettokensCardCardtoken`: DigitalWalletTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.GetDigitalwallettokensCardCardtoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardToken** | **string** | Unique identifier of the card. Used to minimize the need to exchange card details during subsequent calls, and also for troubleshooting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalwallettokensCardCardtokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of digital wallet token resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first digital wallet token resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**DigitalWalletTokenListResponse**](DigitalWalletTokenListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDigitalwallettokensToken

> DigitalWalletToken GetDigitalwallettokensToken(ctx, token).Execute()

Retrieve digital wallet token



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
    token := "token_example" // string | Unique identifier of the digital wallet token (DWT).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.GetDigitalwallettokensToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.GetDigitalwallettokensToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalwallettokensToken`: DigitalWalletToken
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.GetDigitalwallettokensToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the digital wallet token (DWT). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalwallettokensTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigitalWalletToken**](DigitalWalletToken.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDigitalwallettokensTokenShowtokenpan

> DigitalWalletToken GetDigitalwallettokensTokenShowtokenpan(ctx, token).Execute()

Retrieve digital wallet token PAN



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
    token := "token_example" // string | Unique identifier of the digital wallet token (DWT).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.GetDigitalwallettokensTokenShowtokenpan(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.GetDigitalwallettokensTokenShowtokenpan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalwallettokensTokenShowtokenpan`: DigitalWalletToken
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.GetDigitalwallettokensTokenShowtokenpan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the digital wallet token (DWT). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalwallettokensTokenShowtokenpanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigitalWalletToken**](DigitalWalletToken.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDigitalwallettokentransitionsDigitalwallettokenToken

> DigitalWalletTokenTransitionListResponse GetDigitalwallettokentransitionsDigitalwallettokenToken(ctx, token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List transitions for digital wallet token



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
    token := "token_example" // string | Unique identifier of the digital wallet token (DWT).
    count := int32(56) // int32 | Number of digital wallet transitions to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | The sort order index of the first digital wallet token in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-id")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.GetDigitalwallettokentransitionsDigitalwallettokenToken(context.Background(), token).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.GetDigitalwallettokentransitionsDigitalwallettokenToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalwallettokentransitionsDigitalwallettokenToken`: DigitalWalletTokenTransitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.GetDigitalwallettokentransitionsDigitalwallettokenToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the digital wallet token (DWT). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalwallettokentransitionsDigitalwallettokenTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of digital wallet transitions to retrieve. | [default to 5]
 **startIndex** | **int32** | The sort order index of the first digital wallet token in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-id&quot;]

### Return type

[**DigitalWalletTokenTransitionListResponse**](DigitalWalletTokenTransitionListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDigitalwallettokentransitionsToken

> DigitalWalletTokenTransitionResponse GetDigitalwallettokentransitionsToken(ctx, token).Fields(fields).Execute()

Retrieve digital wallet token transition



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
    token := "token_example" // string | Unique identifier of the digital wallet token (DWT) transition.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.GetDigitalwallettokentransitionsToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.GetDigitalwallettokentransitionsToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalwallettokentransitionsToken`: DigitalWalletTokenTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.GetDigitalwallettokentransitionsToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the digital wallet token (DWT) transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalwallettokentransitionsTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**DigitalWalletTokenTransitionResponse**](DigitalWalletTokenTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDigitalwalletprovisionrequestsAndroidpay

> DigitalWalletAndroidPayProvisionResponse PostDigitalwalletprovisionrequestsAndroidpay(ctx).DigitalWalletAndroidPayProvisionRequest(digitalWalletAndroidPayProvisionRequest).Execute()

Create digital wallet token provisioning request for Google Wallet



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
    digitalWalletAndroidPayProvisionRequest := *openapiclient.NewDigitalWalletAndroidPayProvisionRequest("CardToken_example", "DeviceId_example", "DeviceType_example", "ProvisioningAppVersion_example", "WalletAccountId_example") // DigitalWalletAndroidPayProvisionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsAndroidpay(context.Background()).DigitalWalletAndroidPayProvisionRequest(digitalWalletAndroidPayProvisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsAndroidpay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDigitalwalletprovisionrequestsAndroidpay`: DigitalWalletAndroidPayProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsAndroidpay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDigitalwalletprovisionrequestsAndroidpayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digitalWalletAndroidPayProvisionRequest** | [**DigitalWalletAndroidPayProvisionRequest**](DigitalWalletAndroidPayProvisionRequest.md) |  | 

### Return type

[**DigitalWalletAndroidPayProvisionResponse**](DigitalWalletAndroidPayProvisionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDigitalwalletprovisionrequestsApplepay

> DigitalWalletApplePayProvisionResponse PostDigitalwalletprovisionrequestsApplepay(ctx).DigitalWalletApplePayProvisionRequest(digitalWalletApplePayProvisionRequest).Execute()

Create digital wallet token provisioning request for Apple Wallet



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
    digitalWalletApplePayProvisionRequest := *openapiclient.NewDigitalWalletApplePayProvisionRequest("CardToken_example", []string{"Certificates_example"}, "DeviceType_example", "Nonce_example", "NonceSignature_example", "ProvisioningAppVersion_example") // DigitalWalletApplePayProvisionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsApplepay(context.Background()).DigitalWalletApplePayProvisionRequest(digitalWalletApplePayProvisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsApplepay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDigitalwalletprovisionrequestsApplepay`: DigitalWalletApplePayProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsApplepay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDigitalwalletprovisionrequestsApplepayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digitalWalletApplePayProvisionRequest** | [**DigitalWalletApplePayProvisionRequest**](DigitalWalletApplePayProvisionRequest.md) |  | 

### Return type

[**DigitalWalletApplePayProvisionResponse**](DigitalWalletApplePayProvisionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDigitalwalletprovisionrequestsSamsungpay

> DigitalWalletSamsungPayProvisionResponse PostDigitalwalletprovisionrequestsSamsungpay(ctx).DigitalWalletSamsungPayProvisionRequest(digitalWalletSamsungPayProvisionRequest).Execute()

Create digital wallet token provisioning request for Samsung Wallet



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
    digitalWalletSamsungPayProvisionRequest := *openapiclient.NewDigitalWalletSamsungPayProvisionRequest("CardToken_example", "DeviceId_example", "DeviceType_example", "ProvisioningAppVersion_example", "WalletUserId_example") // DigitalWalletSamsungPayProvisionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsSamsungpay(context.Background()).DigitalWalletSamsungPayProvisionRequest(digitalWalletSamsungPayProvisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsSamsungpay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDigitalwalletprovisionrequestsSamsungpay`: DigitalWalletSamsungPayProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsSamsungpay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDigitalwalletprovisionrequestsSamsungpayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digitalWalletSamsungPayProvisionRequest** | [**DigitalWalletSamsungPayProvisionRequest**](DigitalWalletSamsungPayProvisionRequest.md) |  | 

### Return type

[**DigitalWalletSamsungPayProvisionResponse**](DigitalWalletSamsungPayProvisionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDigitalwalletprovisionrequestsXPay

> DigitalWalletXPayProvisionResponse PostDigitalwalletprovisionrequestsXPay(ctx).DigitalWalletXPayProvisionRequest(digitalWalletXPayProvisionRequest).Execute()

Create digital wallet token provisioning request for XPay



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
    digitalWalletXPayProvisionRequest := *openapiclient.NewDigitalWalletXPayProvisionRequest("CardToken_example", "DeviceId_example", "DeviceType_example", "ProvisioningAppVersion_example", "TokenRequestorId_example", "WalletAccountId_example") // DigitalWalletXPayProvisionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsXPay(context.Background()).DigitalWalletXPayProvisionRequest(digitalWalletXPayProvisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsXPay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDigitalwalletprovisionrequestsXPay`: DigitalWalletXPayProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.PostDigitalwalletprovisionrequestsXPay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDigitalwalletprovisionrequestsXPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digitalWalletXPayProvisionRequest** | [**DigitalWalletXPayProvisionRequest**](DigitalWalletXPayProvisionRequest.md) |  | 

### Return type

[**DigitalWalletXPayProvisionResponse**](DigitalWalletXPayProvisionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDigitalwallettokentransitions

> DigitalWalletTokenTransitionResponse PostDigitalwallettokentransitions(ctx).DigitalWalletTokenTransitionRequest(digitalWalletTokenTransitionRequest).Execute()

Create digital wallet token transition



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
    digitalWalletTokenTransitionRequest := *openapiclient.NewDigitalWalletTokenTransitionRequest(*openapiclient.NewDigitalWalletTokenHash("Token_example"), "State_example") // DigitalWalletTokenTransitionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsManagementAPI.PostDigitalwallettokentransitions(context.Background()).DigitalWalletTokenTransitionRequest(digitalWalletTokenTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.PostDigitalwallettokentransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDigitalwallettokentransitions`: DigitalWalletTokenTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsManagementAPI.PostDigitalwallettokentransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDigitalwallettokentransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digitalWalletTokenTransitionRequest** | [**DigitalWalletTokenTransitionRequest**](DigitalWalletTokenTransitionRequest.md) |  | 

### Return type

[**DigitalWalletTokenTransitionResponse**](DigitalWalletTokenTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendOPCDataToGooglePay

> SendOPCDataToGooglePay(ctx).ReqSysId(reqSysId).SendingProvisioningDataToGooglePayBackendRequest(sendingProvisioningDataToGooglePayBackendRequest).Execute()

Create request for Google Wallet web push provisioning



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
    reqSysId := "123d837e-958a-4e9f-bc97-4843ec948123" // string | Random pseudo-unique value used for troubleshooting between multiple parties.
    sendingProvisioningDataToGooglePayBackendRequest := *openapiclient.NewSendingProvisioningDataToGooglePayBackendRequest(int32(0), "5855opl9-abcc-4cb7-a330-xyze5799ea5", "5855opl9-abcc-4cb7-a330-xyze5799ea5", "ACMEISSUER_1", "5855opl9-abcc-4cb7-a330-xyze5799ea5", "5855opl9-abcc-4cb7-a330-xyze5799ea5", "8c84fab9-889c-4cb7-a330-accac5799ea5", int32(0)) // SendingProvisioningDataToGooglePayBackendRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DigitalWalletsManagementAPI.SendOPCDataToGooglePay(context.Background()).ReqSysId(reqSysId).SendingProvisioningDataToGooglePayBackendRequest(sendingProvisioningDataToGooglePayBackendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsManagementAPI.SendOPCDataToGooglePay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendOPCDataToGooglePayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reqSysId** | **string** | Random pseudo-unique value used for troubleshooting between multiple parties. | 
 **sendingProvisioningDataToGooglePayBackendRequest** | [**SendingProvisioningDataToGooglePayBackendRequest**](SendingProvisioningDataToGooglePayBackendRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

