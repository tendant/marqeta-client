# \AccountHolderFundingSourcesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFundingsourcesAchFundingsourcetoken**](AccountHolderFundingSourcesAPI.md#GetFundingsourcesAchFundingsourcetoken) | **Get** /fundingsources/ach/{funding_source_token} | Retrieve ACH source
[**GetFundingsourcesAchFundingsourcetokenVerificationamounts**](AccountHolderFundingSourcesAPI.md#GetFundingsourcesAchFundingsourcetokenVerificationamounts) | **Get** /fundingsources/ach/{funding_source_token}/verificationamounts | Retrieve ACH verification amounts
[**GetFundingsourcesBusinessBusinesstoken**](AccountHolderFundingSourcesAPI.md#GetFundingsourcesBusinessBusinesstoken) | **Get** /fundingsources/business/{business_token} | List sources for business
[**GetFundingsourcesPaymentcardFundingsourcetoken**](AccountHolderFundingSourcesAPI.md#GetFundingsourcesPaymentcardFundingsourcetoken) | **Get** /fundingsources/paymentcard/{funding_source_token} | Retrieve payment card source
[**GetFundingsourcesUserUsertoken**](AccountHolderFundingSourcesAPI.md#GetFundingsourcesUserUsertoken) | **Get** /fundingsources/user/{user_token} | List sources for user
[**PostFundingsourcesAch**](AccountHolderFundingSourcesAPI.md#PostFundingsourcesAch) | **Post** /fundingsources/ach | Create ACH source
[**PostFundingsourcesAchPartner**](AccountHolderFundingSourcesAPI.md#PostFundingsourcesAchPartner) | **Post** /fundingsources/ach/partner | Create ACH source via a partner integration
[**PostFundingsourcesPaymentcard**](AccountHolderFundingSourcesAPI.md#PostFundingsourcesPaymentcard) | **Post** /fundingsources/paymentcard | Create payment card source
[**PutFundingsourcesAchFundingsourcetoken**](AccountHolderFundingSourcesAPI.md#PutFundingsourcesAchFundingsourcetoken) | **Put** /fundingsources/ach/{funding_source_token} | Verify or update ACH source
[**PutFundingsourcesFundingsourcetoken**](AccountHolderFundingSourcesAPI.md#PutFundingsourcesFundingsourcetoken) | **Put** /fundingsources/paymentcard/{funding_source_token} | Update payment card source
[**PutFundingsourcesFundingsourcetokenDefault**](AccountHolderFundingSourcesAPI.md#PutFundingsourcesFundingsourcetokenDefault) | **Put** /fundingsources/{funding_source_token}/default | Set default source



## GetFundingsourcesAchFundingsourcetoken

> AchResponseModel GetFundingsourcesAchFundingsourcetoken(ctx, fundingSourceToken).Execute()

Retrieve ACH source



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
    fundingSourceToken := "fundingSourceToken_example" // string | Unique identifier of the funding source.  Send a `GET` request to `/fundingsources/user/{user_token}` to retrieve existing funding source tokens for a user or to `/fundingsources/business/{business_token}` to retrieve existing funding source tokens for a business.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetoken(context.Background(), fundingSourceToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesAchFundingsourcetoken`: AchResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceToken** | **string** | Unique identifier of the funding source.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve existing funding source tokens for a user or to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve existing funding source tokens for a business. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesAchFundingsourcetokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AchResponseModel**](AchResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsourcesAchFundingsourcetokenVerificationamounts

> AchVerificationModel GetFundingsourcesAchFundingsourcetokenVerificationamounts(ctx, fundingSourceToken).Execute()

Retrieve ACH verification amounts



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
    fundingSourceToken := "fundingSourceToken_example" // string | Unique identifier of the funding source.  Send a `GET` request to `/fundingsources/user/{user_token}` to retrieve existing funding source tokens for a user or to `/fundingsources/business/{business_token}` to retrieve existing funding source tokens for a business.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetokenVerificationamounts(context.Background(), fundingSourceToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetokenVerificationamounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesAchFundingsourcetokenVerificationamounts`: AchVerificationModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetokenVerificationamounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceToken** | **string** | Unique identifier of the funding source.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve existing funding source tokens for a user or to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve existing funding source tokens for a business. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesAchFundingsourcetokenVerificationamountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AchVerificationModel**](AchVerificationModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsourcesBusinessBusinesstoken

> FundingAccountListResponse GetFundingsourcesBusinessBusinesstoken(ctx, businessToken).Type_(type_).Fields(fields).Execute()

List sources for business



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
    type_ := "type__example" // string | Type of funding source to return: ACH or payment card. Leave unspecified to return both types. (optional)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesBusinessBusinesstoken(context.Background(), businessToken).Type_(type_).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.GetFundingsourcesBusinessBusinesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesBusinessBusinesstoken`: FundingAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.GetFundingsourcesBusinessBusinesstoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessToken** | **string** | Unique identifier of the business account holder.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesBusinessBusinesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of funding source to return: ACH or payment card. Leave unspecified to return both types. | 
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**FundingAccountListResponse**](FundingAccountListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsourcesPaymentcardFundingsourcetoken

> PaymentCardResponseModel GetFundingsourcesPaymentcardFundingsourcetoken(ctx, fundingSourceToken).Execute()

Retrieve payment card source



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
    fundingSourceToken := "fundingSourceToken_example" // string | U|- Unique identifier of the funding source. Send a `GET` request to `/fundingsources/user/{user_token}` to retrieve existing funding source tokens for a user or to `/fundingsources/business/{business_token}` to retrieve existing funding source tokens for a business.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesPaymentcardFundingsourcetoken(context.Background(), fundingSourceToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.GetFundingsourcesPaymentcardFundingsourcetoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesPaymentcardFundingsourcetoken`: PaymentCardResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.GetFundingsourcesPaymentcardFundingsourcetoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceToken** | **string** | U|- Unique identifier of the funding source. Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve existing funding source tokens for a user or to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve existing funding source tokens for a business. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesPaymentcardFundingsourcetokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentCardResponseModel**](PaymentCardResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsourcesUserUsertoken

> FundingAccountListResponse GetFundingsourcesUserUsertoken(ctx, userToken).Type_(type_).Fields(fields).Execute()

List sources for user



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
    type_ := "type__example" // string | Type of funding source to retrieve: ACH or payment card. Leave unspecified to return both types. (optional)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesUserUsertoken(context.Background(), userToken).Type_(type_).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.GetFundingsourcesUserUsertoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsourcesUserUsertoken`: FundingAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.GetFundingsourcesUserUsertoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userToken** | **string** | Unique identifier of the user account holder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsourcesUserUsertokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of funding source to retrieve: ACH or payment card. Leave unspecified to return both types. | 
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**FundingAccountListResponse**](FundingAccountListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFundingsourcesAch

> AchResponseModel PostFundingsourcesAch(ctx).AchModel(achModel).Execute()

Create ACH source



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
    achModel := *openapiclient.NewAchModel("AccountNumber_example", "AccountType_example", "NameOnAccount_example", "RoutingNumber_example") // AchModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.PostFundingsourcesAch(context.Background()).AchModel(achModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.PostFundingsourcesAch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFundingsourcesAch`: AchResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.PostFundingsourcesAch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFundingsourcesAchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achModel** | [**AchModel**](AchModel.md) |  | 

### Return type

[**AchResponseModel**](AchResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFundingsourcesAchPartner

> AchResponseModel PostFundingsourcesAchPartner(ctx).AchPartnerRequestModel(achPartnerRequestModel).Execute()

Create ACH source via a partner integration



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
    achPartnerRequestModel := *openapiclient.NewAchPartnerRequestModel("Partner_example", "PartnerAccountLinkReferenceToken_example") // AchPartnerRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.PostFundingsourcesAchPartner(context.Background()).AchPartnerRequestModel(achPartnerRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.PostFundingsourcesAchPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFundingsourcesAchPartner`: AchResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.PostFundingsourcesAchPartner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFundingsourcesAchPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achPartnerRequestModel** | [**AchPartnerRequestModel**](AchPartnerRequestModel.md) |  | 

### Return type

[**AchResponseModel**](AchResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFundingsourcesPaymentcard

> PaymentCardResponseModel PostFundingsourcesPaymentcard(ctx).TokenRequest(tokenRequest).Execute()

Create payment card source



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
    tokenRequest := *openapiclient.NewTokenRequest("AccountNumber_example", "CvvNumber_example", "ExpDate_example") // TokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.PostFundingsourcesPaymentcard(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.PostFundingsourcesPaymentcard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFundingsourcesPaymentcard`: PaymentCardResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.PostFundingsourcesPaymentcard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFundingsourcesPaymentcardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**PaymentCardResponseModel**](PaymentCardResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFundingsourcesAchFundingsourcetoken

> AchResponseModel PutFundingsourcesAchFundingsourcetoken(ctx, fundingSourceToken).AchVerificationModel(achVerificationModel).Execute()

Verify or update ACH source



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
    fundingSourceToken := "fundingSourceToken_example" // string | Unique identifier of the funding source.  Send a `GET` request to `/fundingsources/user/{user_token}` to retrieve existing funding source tokens for a user or to `/fundingsources/business/{business_token}` to retrieve existing funding source tokens for a business.
    achVerificationModel := *openapiclient.NewAchVerificationModel() // AchVerificationModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.PutFundingsourcesAchFundingsourcetoken(context.Background(), fundingSourceToken).AchVerificationModel(achVerificationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.PutFundingsourcesAchFundingsourcetoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFundingsourcesAchFundingsourcetoken`: AchResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.PutFundingsourcesAchFundingsourcetoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceToken** | **string** | Unique identifier of the funding source.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve existing funding source tokens for a user or to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve existing funding source tokens for a business. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFundingsourcesAchFundingsourcetokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **achVerificationModel** | [**AchVerificationModel**](AchVerificationModel.md) |  | 

### Return type

[**AchResponseModel**](AchResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFundingsourcesFundingsourcetoken

> PaymentCardResponseModel PutFundingsourcesFundingsourcetoken(ctx, fundingSourceToken).TokenUpdateRequest(tokenUpdateRequest).Execute()

Update payment card source



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
    fundingSourceToken := "fundingSourceToken_example" // string | Unique identifier of the funding source.  Send a `GET` request to `/fundingsources/user/{user_token}` to retrieve existing funding source tokens for a user or to `/fundingsources/business/{business_token}` to retrieve existing funding source tokens for a business.
    tokenUpdateRequest := *openapiclient.NewTokenUpdateRequest("ExpDate_example") // TokenUpdateRequest | Payment card

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetoken(context.Background(), fundingSourceToken).TokenUpdateRequest(tokenUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFundingsourcesFundingsourcetoken`: PaymentCardResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceToken** | **string** | Unique identifier of the funding source.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve existing funding source tokens for a user or to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve existing funding source tokens for a business. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFundingsourcesFundingsourcetokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenUpdateRequest** | [**TokenUpdateRequest**](TokenUpdateRequest.md) | Payment card | 

### Return type

[**PaymentCardResponseModel**](PaymentCardResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFundingsourcesFundingsourcetokenDefault

> PaymentCardResponseModel PutFundingsourcesFundingsourcetokenDefault(ctx, fundingSourceToken).Execute()

Set default source



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
    fundingSourceToken := "fundingSourceToken_example" // string | Unique identifier of the funding source.  Send a `GET` request to `/fundingsources/user/{user_token}` to retrieve existing funding source tokens for a user or to `/fundingsources/business/{business_token}` to retrieve existing funding source tokens for a business.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetokenDefault(context.Background(), fundingSourceToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetokenDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFundingsourcesFundingsourcetokenDefault`: PaymentCardResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetokenDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingSourceToken** | **string** | Unique identifier of the funding source.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve existing funding source tokens for a user or to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve existing funding source tokens for a business. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFundingsourcesFundingsourcetokenDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentCardResponseModel**](PaymentCardResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

