# \PoliciesBetaAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneAprPolicy**](PoliciesBetaAPI.md#CloneAprPolicy) | **Post** /credit/policies/aprs/{token}/clone | Clone APR policy
[**CloneDocumentPolicy**](PoliciesBetaAPI.md#CloneDocumentPolicy) | **Post** /credit/policies/documents/{token}/clone | Clone document policy
[**CloneFeePolicy**](PoliciesBetaAPI.md#CloneFeePolicy) | **Post** /credit/policies/fees/{token}/clone | Clone fee policy
[**CloneProductPolicy**](PoliciesBetaAPI.md#CloneProductPolicy) | **Post** /credit/policies/products/{token}/clone | Clone credit product policy
[**CloneRewardPolicy**](PoliciesBetaAPI.md#CloneRewardPolicy) | **Post** /credit/policies/rewards/{token}/clone | Clone reward policy
[**CreateAprPolicy**](PoliciesBetaAPI.md#CreateAprPolicy) | **Post** /credit/policies/aprs | Create APR policy
[**CreateDocumentPolicy**](PoliciesBetaAPI.md#CreateDocumentPolicy) | **Post** /credit/policies/documents | Create document policy
[**CreateFeePolicy**](PoliciesBetaAPI.md#CreateFeePolicy) | **Post** /credit/policies/fees | Create fee policy
[**CreateProductPolicy**](PoliciesBetaAPI.md#CreateProductPolicy) | **Post** /credit/policies/products | Create credit product policy
[**CreateRewardPolicy**](PoliciesBetaAPI.md#CreateRewardPolicy) | **Post** /credit/policies/rewards | Create reward policy
[**GetAprPolicies**](PoliciesBetaAPI.md#GetAprPolicies) | **Get** /credit/policies/aprs | List APR policies
[**GetAprPolicyByToken**](PoliciesBetaAPI.md#GetAprPolicyByToken) | **Get** /credit/policies/aprs/{token} | Retrieve APR policy
[**GetAprPolicySchedulesWithToken**](PoliciesBetaAPI.md#GetAprPolicySchedulesWithToken) | **Get** /credit/policies/aprs/{token}/schedule | List APR schedules
[**GetFeePolicies**](PoliciesBetaAPI.md#GetFeePolicies) | **Get** /credit/policies/fees | List fee policies
[**GetFeePolicyByToken**](PoliciesBetaAPI.md#GetFeePolicyByToken) | **Get** /credit/policies/fees/{token} | Retrieve fee policy
[**ListDocumentPolicies**](PoliciesBetaAPI.md#ListDocumentPolicies) | **Get** /credit/policies/documents | List document policies
[**ListProductPolicies**](PoliciesBetaAPI.md#ListProductPolicies) | **Get** /credit/policies/products | List credit product policies
[**ListRewardPolicies**](PoliciesBetaAPI.md#ListRewardPolicies) | **Get** /credit/policies/rewards | List reward policies
[**RetrieveDocumentPolicy**](PoliciesBetaAPI.md#RetrieveDocumentPolicy) | **Get** /credit/policies/documents/{token} | Retrieve document policy
[**RetrieveProductPolicy**](PoliciesBetaAPI.md#RetrieveProductPolicy) | **Get** /credit/policies/products/{token} | Retrieve credit product policy
[**RetrieveRewardPolicy**](PoliciesBetaAPI.md#RetrieveRewardPolicy) | **Get** /credit/policies/rewards/{token} | Retrieve reward policy
[**UpdateAprPolicyWithToken**](PoliciesBetaAPI.md#UpdateAprPolicyWithToken) | **Put** /credit/policies/aprs/{token} | Update APR policy
[**UpdateDocumentPolicy**](PoliciesBetaAPI.md#UpdateDocumentPolicy) | **Put** /credit/policies/documents/{token} | Update document policy
[**UpdateFeePolicyWithToken**](PoliciesBetaAPI.md#UpdateFeePolicyWithToken) | **Put** /credit/policies/fees/{token} | Update fee policy
[**UpdateProductPolicy**](PoliciesBetaAPI.md#UpdateProductPolicy) | **Put** /credit/policies/products/{token} | Update credit product policy
[**UpdateRewardPolicy**](PoliciesBetaAPI.md#UpdateRewardPolicy) | **Put** /credit/policies/rewards/{token} | Update reward policy



## CloneAprPolicy

> PolicyAprResponse CloneAprPolicy(ctx, token).Execute()

Clone APR policy



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
    token := "token_example" // string | The unique identifier of the APR policy to clone.  Send a `GET` request to `/policies/aprs` to retrieve existing APR policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CloneAprPolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CloneAprPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneAprPolicy`: PolicyAprResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CloneAprPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the APR policy to clone.  Send a &#x60;GET&#x60; request to &#x60;/policies/aprs&#x60; to retrieve existing APR policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneAprPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyAprResponse**](PolicyAprResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneDocumentPolicy

> PolicyDocumentResponse CloneDocumentPolicy(ctx, token).Execute()

Clone document policy



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
    token := "token_example" // string | The unique identifier of the document policy to clone.  Send a `GET` request to `/policies/documents` to retrieve existing document policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CloneDocumentPolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CloneDocumentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneDocumentPolicy`: PolicyDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CloneDocumentPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the document policy to clone.  Send a &#x60;GET&#x60; request to &#x60;/policies/documents&#x60; to retrieve existing document policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDocumentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyDocumentResponse**](PolicyDocumentResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneFeePolicy

> PolicyFeeResponse CloneFeePolicy(ctx, token).Execute()

Clone fee policy



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
    token := "token_example" // string | The unique identifier of the fee policy to clone.  Send a `GET` request to `/policies/fee` to retrieve existing fee policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CloneFeePolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CloneFeePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneFeePolicy`: PolicyFeeResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CloneFeePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the fee policy to clone.  Send a &#x60;GET&#x60; request to &#x60;/policies/fee&#x60; to retrieve existing fee policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneFeePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyFeeResponse**](PolicyFeeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneProductPolicy

> PolicyProductResponse CloneProductPolicy(ctx, token).Execute()

Clone credit product policy



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
    token := "token_example" // string | The unique identifier of the credit product policy to clone.  Send a `GET` request to `/policies/products` to retrieve existing credit product policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CloneProductPolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CloneProductPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneProductPolicy`: PolicyProductResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CloneProductPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the credit product policy to clone.  Send a &#x60;GET&#x60; request to &#x60;/policies/products&#x60; to retrieve existing credit product policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneProductPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyProductResponse**](PolicyProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneRewardPolicy

> PolicyRewardResponse CloneRewardPolicy(ctx, token).Execute()

Clone reward policy



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
    token := "token_example" // string | The unique identifier of the reward policy to clone.  Send a `GET` request to `/policies/rewards` to retrieve existing reward policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CloneRewardPolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CloneRewardPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneRewardPolicy`: PolicyRewardResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CloneRewardPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the reward policy to clone.  Send a &#x60;GET&#x60; request to &#x60;/policies/rewards&#x60; to retrieve existing reward policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneRewardPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyRewardResponse**](PolicyRewardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAprPolicy

> PolicyAprResponse CreateAprPolicy(ctx).PolicyAprCreateReq(policyAprCreateReq).Execute()

Create APR policy



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
    policyAprCreateReq := *openapiclient.NewPolicyAprCreateReq("Name_example", *openapiclient.NewPolicyAprPurchaseReq([]openapiclient.PolicyAprTierReq{*openapiclient.NewPolicyAprTierReq(float32(123))})) // PolicyAprCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CreateAprPolicy(context.Background()).PolicyAprCreateReq(policyAprCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CreateAprPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAprPolicy`: PolicyAprResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CreateAprPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAprPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyAprCreateReq** | [**PolicyAprCreateReq**](PolicyAprCreateReq.md) |  | 

### Return type

[**PolicyAprResponse**](PolicyAprResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocumentPolicy

> PolicyDocumentResponse CreateDocumentPolicy(ctx).PolicyDocumentCreateReq(policyDocumentCreateReq).Execute()

Create document policy



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
    policyDocumentCreateReq := *openapiclient.NewPolicyDocumentCreateReq(*openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), "Name_example", *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetAndTemplateReq("AssetToken_example", "TemplateToken_example"), *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example")) // PolicyDocumentCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CreateDocumentPolicy(context.Background()).PolicyDocumentCreateReq(policyDocumentCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CreateDocumentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocumentPolicy`: PolicyDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CreateDocumentPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyDocumentCreateReq** | [**PolicyDocumentCreateReq**](PolicyDocumentCreateReq.md) |  | 

### Return type

[**PolicyDocumentResponse**](PolicyDocumentResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeePolicy

> PolicyFeeResponse CreateFeePolicy(ctx).PolicyFeeCreateReq(policyFeeCreateReq).Execute()

Create fee policy



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
    policyFeeCreateReq := *openapiclient.NewPolicyFeeCreateReq(*openapiclient.NewPolicyFeeAccount(), "Name_example") // PolicyFeeCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CreateFeePolicy(context.Background()).PolicyFeeCreateReq(policyFeeCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CreateFeePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeePolicy`: PolicyFeeResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CreateFeePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyFeeCreateReq** | [**PolicyFeeCreateReq**](PolicyFeeCreateReq.md) |  | 

### Return type

[**PolicyFeeResponse**](PolicyFeeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProductPolicy

> PolicyProductResponse CreateProductPolicy(ctx).PolicyProductCreateReq(policyProductCreateReq).Execute()

Create credit product policy



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
    policyProductCreateReq := *openapiclient.NewPolicyProductCreateReq([]openapiclient.PolicyProductCardProductReq{*openapiclient.NewPolicyProductCardProductReq(openapiclient.PolicyProductCardProductLevel("PREMIUM"), "Token_example")}, openapiclient.ProductClassification("CONSUMER"), *openapiclient.NewProductCreditLine(float32(123), float32(123)), openapiclient.CurrencyCode("USD"), *openapiclient.NewInterestCalculation(*openapiclient.NewApplicationOfCredits(openapiclient.CycleType("BEGINNING_REVOLVING"), int32(123)), "DayCount_example", "GraceDaysApplication_example", []string{"InterestApplication_example"}, openapiclient.InterestOnGraceReactivationEnum("ACCRUE_FULL_CYCLE"), "Method_example", float32(123)), "Name_example", *openapiclient.NewPolicyProductPaymentConfiguration([]openapiclient.PaymentAllocationOrderEnum{openapiclient.PaymentAllocationOrderEnum("INTEREST")}, int32(123), int32(123), *openapiclient.NewPolicyProductMinPaymentCalculation(false, false)), openapiclient.ProductSubType("CREDIT_CARD"), openapiclient.ProductType("REVOLVING"), []openapiclient.BalanceType{openapiclient.BalanceType("PURCHASE")}) // PolicyProductCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CreateProductPolicy(context.Background()).PolicyProductCreateReq(policyProductCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CreateProductPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductPolicy`: PolicyProductResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CreateProductPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyProductCreateReq** | [**PolicyProductCreateReq**](PolicyProductCreateReq.md) |  | 

### Return type

[**PolicyProductResponse**](PolicyProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRewardPolicy

> PolicyRewardResponse CreateRewardPolicy(ctx).PolicyRewardReq(policyRewardReq).Execute()

Create reward policy



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
    policyRewardReq := *openapiclient.NewPolicyRewardReq("Name_example", []openapiclient.PolicyRewardRule{*openapiclient.NewPolicyRewardRule(*openapiclient.NewPolicyRewardRuleFilters(*openapiclient.NewAmountFilter()), *openapiclient.NewPolicyRewardRuleOutcome(float32(123)), openapiclient.PolicyRewardRuleType("CASHBACK"))}) // PolicyRewardReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.CreateRewardPolicy(context.Background()).PolicyRewardReq(policyRewardReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.CreateRewardPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRewardPolicy`: PolicyRewardResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.CreateRewardPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRewardPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyRewardReq** | [**PolicyRewardReq**](PolicyRewardReq.md) |  | 

### Return type

[**PolicyRewardResponse**](PolicyRewardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAprPolicies

> PolicyAprsPage GetAprPolicies(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List APR policies



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
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `effectiveDate`, and not by the field names appearing in response bodies such as `effective_date`. (optional) (default to "-effectiveDate")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.GetAprPolicies(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.GetAprPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAprPolicies`: PolicyAprsPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.GetAprPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAprPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;effectiveDate&#x60;, and not by the field names appearing in response bodies such as &#x60;effective_date&#x60;. | [default to &quot;-effectiveDate&quot;]

### Return type

[**PolicyAprsPage**](PolicyAprsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAprPolicyByToken

> PolicyAprResponse GetAprPolicyByToken(ctx, token).Execute()

Retrieve APR policy



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
    token := "token_example" // string | The unique identifier of the APR policy to retrieve.  Send a `GET` request to `/policies/aprs` to retrieve existing APR policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.GetAprPolicyByToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.GetAprPolicyByToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAprPolicyByToken`: PolicyAprResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.GetAprPolicyByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the APR policy to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/policies/aprs&#x60; to retrieve existing APR policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAprPolicyByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyAprResponse**](PolicyAprResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAprPolicySchedulesWithToken

> PolicyAprsPage GetAprPolicySchedulesWithToken(ctx, token).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List APR schedules



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
    token := "token_example" // string | The unique identifier of the APR policy on which to retrieve APR schedules.  Send a `GET` request to `/policies/aprs` to retrieve existing product policy tokens.
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.GetAprPolicySchedulesWithToken(context.Background(), token).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.GetAprPolicySchedulesWithToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAprPolicySchedulesWithToken`: PolicyAprsPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.GetAprPolicySchedulesWithToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the APR policy on which to retrieve APR schedules.  Send a &#x60;GET&#x60; request to &#x60;/policies/aprs&#x60; to retrieve existing product policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAprPolicySchedulesWithTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**PolicyAprsPage**](PolicyAprsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeePolicies

> PolicyFeesPage GetFeePolicies(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List fee policies



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
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.GetFeePolicies(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.GetFeePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeePolicies`: PolicyFeesPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.GetFeePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**PolicyFeesPage**](PolicyFeesPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeePolicyByToken

> PolicyFeeResponse GetFeePolicyByToken(ctx, token).Execute()

Retrieve fee policy



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
    token := "token_example" // string | The unique identifier of the fee policy to retrieve.  Send a `GET` request to `/policies/fee` to retrieve existing fee policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.GetFeePolicyByToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.GetFeePolicyByToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeePolicyByToken`: PolicyFeeResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.GetFeePolicyByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the fee policy to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/policies/fee&#x60; to retrieve existing fee policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeePolicyByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyFeeResponse**](PolicyFeeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentPolicies

> PoliciesDocumentPage ListDocumentPolicies(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List document policies



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
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.ListDocumentPolicies(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.ListDocumentPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocumentPolicies`: PoliciesDocumentPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.ListDocumentPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**PoliciesDocumentPage**](PoliciesDocumentPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductPolicies

> PoliciesProductPage ListProductPolicies(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List credit product policies



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
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.ListProductPolicies(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.ListProductPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProductPolicies`: PoliciesProductPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.ListProductPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**PoliciesProductPage**](PoliciesProductPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRewardPolicies

> PolicyRewardPage ListRewardPolicies(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List reward policies



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
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.ListRewardPolicies(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.ListRewardPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRewardPolicies`: PolicyRewardPage
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.ListRewardPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRewardPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**PolicyRewardPage**](PolicyRewardPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDocumentPolicy

> PolicyDocumentResponse RetrieveDocumentPolicy(ctx, token).Execute()

Retrieve document policy



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
    token := "token_example" // string | The unique identifier of the document policy to retrieve.  Send a `GET` request to `/policies/documents` to retrieve existing document policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.RetrieveDocumentPolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.RetrieveDocumentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDocumentPolicy`: PolicyDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.RetrieveDocumentPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the document policy to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/policies/documents&#x60; to retrieve existing document policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDocumentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyDocumentResponse**](PolicyDocumentResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveProductPolicy

> PolicyProductResponse RetrieveProductPolicy(ctx, token).Execute()

Retrieve credit product policy



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
    token := "token_example" // string | The unique identifier of the credit product policy to retrieve.  Send a `GET` request to `/policies/products` to retrieve existing credit product policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.RetrieveProductPolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.RetrieveProductPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveProductPolicy`: PolicyProductResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.RetrieveProductPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the credit product policy to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/policies/products&#x60; to retrieve existing credit product policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveProductPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyProductResponse**](PolicyProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveRewardPolicy

> PolicyRewardResponse RetrieveRewardPolicy(ctx, token).Execute()

Retrieve reward policy



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
    token := "token_example" // string | The unique identifier of the reward policy to retrieve.  Send a `GET` request to `/policies/rewards` to retrieve existing reward policy tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.RetrieveRewardPolicy(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.RetrieveRewardPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRewardPolicy`: PolicyRewardResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.RetrieveRewardPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the reward policy to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/policies/rewards&#x60; to retrieve existing reward policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRewardPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyRewardResponse**](PolicyRewardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAprPolicyWithToken

> PolicyAprResponse UpdateAprPolicyWithToken(ctx, token).PolicyAprUpdateReq(policyAprUpdateReq).Execute()

Update APR policy



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
    token := "token_example" // string | The unique identifier of the APR policy to update.  Send a `GET` request to `/policies/aprs` to retrieve existing APR policy tokens.
    policyAprUpdateReq := *openapiclient.NewPolicyAprUpdateReq("Name_example") // PolicyAprUpdateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.UpdateAprPolicyWithToken(context.Background(), token).PolicyAprUpdateReq(policyAprUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.UpdateAprPolicyWithToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAprPolicyWithToken`: PolicyAprResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.UpdateAprPolicyWithToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the APR policy to update.  Send a &#x60;GET&#x60; request to &#x60;/policies/aprs&#x60; to retrieve existing APR policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAprPolicyWithTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyAprUpdateReq** | [**PolicyAprUpdateReq**](PolicyAprUpdateReq.md) |  | 

### Return type

[**PolicyAprResponse**](PolicyAprResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPolicy

> PolicyDocumentResponse UpdateDocumentPolicy(ctx, token).PolicyDocumentUpdateReq(policyDocumentUpdateReq).Execute()

Update document policy



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
    token := "token_example" // string | The unique identifier of the document policy to update.  Send a `GET` request to `/policies/documents` to retrieve existing document policy tokens.
    policyDocumentUpdateReq := *openapiclient.NewPolicyDocumentUpdateReq(*openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), "Name_example", *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example"), *openapiclient.NewPolicyDocumentAssetReq("AssetToken_example"), *openapiclient.NewPolicyDocumentAssetAndTemplateReq("AssetToken_example", "TemplateToken_example"), *openapiclient.NewPolicyDocumentTemplateReq("TemplateToken_example")) // PolicyDocumentUpdateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.UpdateDocumentPolicy(context.Background(), token).PolicyDocumentUpdateReq(policyDocumentUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.UpdateDocumentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocumentPolicy`: PolicyDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.UpdateDocumentPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the document policy to update.  Send a &#x60;GET&#x60; request to &#x60;/policies/documents&#x60; to retrieve existing document policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyDocumentUpdateReq** | [**PolicyDocumentUpdateReq**](PolicyDocumentUpdateReq.md) |  | 

### Return type

[**PolicyDocumentResponse**](PolicyDocumentResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeePolicyWithToken

> PolicyFeeResponse UpdateFeePolicyWithToken(ctx, token).PolicyFeeUpdateReq(policyFeeUpdateReq).Execute()

Update fee policy



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
    token := "token_example" // string | The unique identifier of the fee policy to retrieve.  Send a `GET` request to `/policies/fee` to retrieve existing fee policy tokens.
    policyFeeUpdateReq := *openapiclient.NewPolicyFeeUpdateReq("Name_example") // PolicyFeeUpdateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.UpdateFeePolicyWithToken(context.Background(), token).PolicyFeeUpdateReq(policyFeeUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.UpdateFeePolicyWithToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeePolicyWithToken`: PolicyFeeResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.UpdateFeePolicyWithToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the fee policy to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/policies/fee&#x60; to retrieve existing fee policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeePolicyWithTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyFeeUpdateReq** | [**PolicyFeeUpdateReq**](PolicyFeeUpdateReq.md) |  | 

### Return type

[**PolicyFeeResponse**](PolicyFeeResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductPolicy

> PolicyProductResponse UpdateProductPolicy(ctx, token).PolicyProductUpdateReq(policyProductUpdateReq).Execute()

Update credit product policy



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
    token := "token_example" // string | The unique identifier of the credit product policy to retrieve.  Send a `GET` request to `/policies/products` to retrieve existing credit product policy tokens.
    policyProductUpdateReq := *openapiclient.NewPolicyProductUpdateReq([]openapiclient.PolicyProductCardProductReq{*openapiclient.NewPolicyProductCardProductReq(openapiclient.PolicyProductCardProductLevel("PREMIUM"), "Token_example")}, openapiclient.ProductClassification("CONSUMER"), *openapiclient.NewProductCreditLine(float32(123), float32(123)), openapiclient.CurrencyCode("USD"), *openapiclient.NewInterestCalculation(*openapiclient.NewApplicationOfCredits(openapiclient.CycleType("BEGINNING_REVOLVING"), int32(123)), "DayCount_example", "GraceDaysApplication_example", []string{"InterestApplication_example"}, openapiclient.InterestOnGraceReactivationEnum("ACCRUE_FULL_CYCLE"), "Method_example", float32(123)), "Name_example", *openapiclient.NewPolicyProductPaymentConfiguration([]openapiclient.PaymentAllocationOrderEnum{openapiclient.PaymentAllocationOrderEnum("INTEREST")}, int32(123), int32(123), *openapiclient.NewPolicyProductMinPaymentCalculation(false, false)), openapiclient.ProductSubType("CREDIT_CARD"), openapiclient.ProductType("REVOLVING"), []openapiclient.BalanceType{openapiclient.BalanceType("PURCHASE")}) // PolicyProductUpdateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.UpdateProductPolicy(context.Background(), token).PolicyProductUpdateReq(policyProductUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.UpdateProductPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProductPolicy`: PolicyProductResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.UpdateProductPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the credit product policy to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/policies/products&#x60; to retrieve existing credit product policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyProductUpdateReq** | [**PolicyProductUpdateReq**](PolicyProductUpdateReq.md) |  | 

### Return type

[**PolicyProductResponse**](PolicyProductResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRewardPolicy

> PolicyRewardResponse UpdateRewardPolicy(ctx, token).PolicyRewardReq(policyRewardReq).Execute()

Update reward policy



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
    token := "token_example" // string | The unique identifier of the reward policy to update.  Send a `GET` request to `/policies/rewards` to retrieve existing reward policy tokens.
    policyRewardReq := *openapiclient.NewPolicyRewardReq("Name_example", []openapiclient.PolicyRewardRule{*openapiclient.NewPolicyRewardRule(*openapiclient.NewPolicyRewardRuleFilters(*openapiclient.NewAmountFilter()), *openapiclient.NewPolicyRewardRuleOutcome(float32(123)), openapiclient.PolicyRewardRuleType("CASHBACK"))}) // PolicyRewardReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesBetaAPI.UpdateRewardPolicy(context.Background(), token).PolicyRewardReq(policyRewardReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesBetaAPI.UpdateRewardPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRewardPolicy`: PolicyRewardResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesBetaAPI.UpdateRewardPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the reward policy to update.  Send a &#x60;GET&#x60; request to &#x60;/policies/rewards&#x60; to retrieve existing reward policy tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRewardPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyRewardReq** | [**PolicyRewardReq**](PolicyRewardReq.md) |  | 

### Return type

[**PolicyRewardResponse**](PolicyRewardResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

