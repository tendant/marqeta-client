# \BundlesBetaAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneBundle**](BundlesBetaAPI.md#CloneBundle) | **Post** /credit/bundles/{token}/clone | Clone bundle
[**CreateBundle**](BundlesBetaAPI.md#CreateBundle) | **Post** /credit/bundles | Create bundle
[**ListBundles**](BundlesBetaAPI.md#ListBundles) | **Get** /credit/bundles | List bundles
[**ListRelatedBundles**](BundlesBetaAPI.md#ListRelatedBundles) | **Get** /credit/bundles/{token}/lineage | List related bundles
[**PromoteBundle**](BundlesBetaAPI.md#PromoteBundle) | **Put** /credit/bundles/{token}/promote | Promote bundle
[**RetrieveBundle**](BundlesBetaAPI.md#RetrieveBundle) | **Get** /credit/bundles/{token} | Retrieve bundle
[**TransitionBundle**](BundlesBetaAPI.md#TransitionBundle) | **Post** /credit/bundles/{token}/transitions | Transition a bundle
[**UpdateBundle**](BundlesBetaAPI.md#UpdateBundle) | **Put** /credit/bundles/{token} | Update bundle



## CloneBundle

> BundleResponse CloneBundle(ctx, token).Execute()

Clone bundle



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
    token := "token_example" // string | The unique identifier of the bundle to clone.  Send a `GET` request to `/bundles` to retrieve existing bundle tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.CloneBundle(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.CloneBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneBundle`: BundleResponse
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.CloneBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the bundle to clone.  Send a &#x60;GET&#x60; request to &#x60;/bundles&#x60; to retrieve existing bundle tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BundleResponse**](BundleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBundle

> BundleResponse CreateBundle(ctx).BundleCreateReq(bundleCreateReq).Execute()

Create bundle



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
    bundleCreateReq := *openapiclient.NewBundleCreateReq("AprPolicyToken_example", "CreditProductPolicyToken_example", "Description_example", "DocumentPolicyToken_example", "FeePolicyToken_example", "Name_example", openapiclient.BundleResourceStatus("DRAFT")) // BundleCreateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.CreateBundle(context.Background()).BundleCreateReq(bundleCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.CreateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBundle`: BundleResponse
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.CreateBundle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleCreateReq** | [**BundleCreateReq**](BundleCreateReq.md) |  | 

### Return type

[**BundleResponse**](BundleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundles

> BundleResponsePage ListBundles(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Status(status).Execute()

List bundles



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
    status := []openapiclient.BundleResourceStatus{openapiclient.BundleResourceStatus("DRAFT")} // []BundleResourceStatus | An array of statuses by which to filter bundles. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.ListBundles(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.ListBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBundles`: BundleResponsePage
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.ListBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]
 **status** | [**[]BundleResourceStatus**](BundleResourceStatus.md) | An array of statuses by which to filter bundles. | 

### Return type

[**BundleResponsePage**](BundleResponsePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRelatedBundles

> BundleResponsePage ListRelatedBundles(ctx, token).Count(count).StartIndex(startIndex).SortBy(sortBy).Status(status).Execute()

List related bundles



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
    token := "token_example" // string | The unique identifier of the parent product bundle.  Send a `GET` request to `/bundles` to retrieve existing product bundles tokens.
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")
    status := []openapiclient.BundleResourceStatus{openapiclient.BundleResourceStatus("DRAFT")} // []BundleResourceStatus | An array of statuses by which to filter bundles. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.ListRelatedBundles(context.Background(), token).Count(count).StartIndex(startIndex).SortBy(sortBy).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.ListRelatedBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRelatedBundles`: BundleResponsePage
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.ListRelatedBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the parent product bundle.  Send a &#x60;GET&#x60; request to &#x60;/bundles&#x60; to retrieve existing product bundles tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRelatedBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]
 **status** | [**[]BundleResourceStatus**](BundleResourceStatus.md) | An array of statuses by which to filter bundles. | 

### Return type

[**BundleResponsePage**](BundleResponsePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteBundle

> BundleResponse PromoteBundle(ctx, token).Execute()

Promote bundle



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
    token := "token_example" // string | The unique identifier of the bundle to promote.  Send a `GET` request to `/bundles` to retrieve existing bundle tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.PromoteBundle(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.PromoteBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromoteBundle`: BundleResponse
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.PromoteBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the bundle to promote.  Send a &#x60;GET&#x60; request to &#x60;/bundles&#x60; to retrieve existing bundle tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BundleResponse**](BundleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBundle

> BundleResponse RetrieveBundle(ctx, token).ExpandObjects(expandObjects).Execute()

Retrieve bundle



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
    token := "token_example" // string | The unique identifier of the bundle to retrieve.  Send a `GET` request to `/credit/bundles` to retrieve existing  tokens.
    expandObjects := []openapiclient.PolicyType{openapiclient.PolicyType("APR")} // []PolicyType | Embeds the associated object of the specified type into the response. For more, see <</core-api/object-expansion, object expansion>>. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.RetrieveBundle(context.Background(), token).ExpandObjects(expandObjects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.RetrieveBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveBundle`: BundleResponse
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.RetrieveBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the bundle to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/bundles&#x60; to retrieve existing  tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expandObjects** | [**[]PolicyType**](PolicyType.md) | Embeds the associated object of the specified type into the response. For more, see &lt;&lt;/core-api/object-expansion, object expansion&gt;&gt;. | 

### Return type

[**BundleResponse**](BundleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionBundle

> BundleTransitionResponse TransitionBundle(ctx, token).BundleTransitionReq(bundleTransitionReq).Execute()

Transition a bundle



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
    token := "token_example" // string | Token of the bundle whose status you want to transition.  Send a `GET` request to `/credit/bundles` to retrieve existing  tokens.
    bundleTransitionReq := *openapiclient.NewBundleTransitionReq(openapiclient.BundleResourceStatus("DRAFT")) // BundleTransitionReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.TransitionBundle(context.Background(), token).BundleTransitionReq(bundleTransitionReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.TransitionBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransitionBundle`: BundleTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.TransitionBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Token of the bundle whose status you want to transition.  Send a &#x60;GET&#x60; request to &#x60;/credit/bundles&#x60; to retrieve existing  tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleTransitionReq** | [**BundleTransitionReq**](BundleTransitionReq.md) |  | 

### Return type

[**BundleTransitionResponse**](BundleTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBundle

> BundleResponse UpdateBundle(ctx, token).BundleUpdateReq(bundleUpdateReq).Execute()

Update bundle



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
    token := "token_example" // string | Token of the bundle to update.  Send a `GET` request to `/credit/bundles` to retrieve existing  tokens.
    bundleUpdateReq := *openapiclient.NewBundleUpdateReq("AprPolicyToken_example", "CreditProductPolicyToken_example", "Description_example", "DocumentPolicyToken_example", "FeePolicyToken_example", "Name_example") // BundleUpdateReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesBetaAPI.UpdateBundle(context.Background(), token).BundleUpdateReq(bundleUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesBetaAPI.UpdateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBundle`: BundleResponse
    fmt.Fprintf(os.Stdout, "Response from `BundlesBetaAPI.UpdateBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Token of the bundle to update.  Send a &#x60;GET&#x60; request to &#x60;/credit/bundles&#x60; to retrieve existing  tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleUpdateReq** | [**BundleUpdateReq**](BundleUpdateReq.md) |  | 

### Return type

[**BundleResponse**](BundleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

