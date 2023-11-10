# \ApplicationsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditApplicationsPost**](ApplicationsAPI.md#CreditApplicationsPost) | **Post** /credit/applications | Create application
[**GetFileByType**](ApplicationsAPI.md#GetFileByType) | **Get** /credit/applications/files/{type} | Retrieve file on a bundle or application
[**PageApplicationTransitions**](ApplicationsAPI.md#PageApplicationTransitions) | **Get** /credit/applications/{token}/transitions | List application transitions
[**RetrieveApplication**](ApplicationsAPI.md#RetrieveApplication) | **Get** /credit/applications/{token} | Retrieve application
[**RetrieveFiles**](ApplicationsAPI.md#RetrieveFiles) | **Get** /credit/applications/files | List files on a bundle or application
[**TransitionApplication**](ApplicationsAPI.md#TransitionApplication) | **Post** /credit/applications/{token}/transitions | Transition application state



## CreditApplicationsPost

> ApplicationsResponse CreditApplicationsPost(ctx).CreateApplicationsRequest(createApplicationsRequest).Execute()

Create application



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
    createApplicationsRequest := *openapiclient.NewCreateApplicationsRequest(false, "BundleToken_example", "EDisclosureTrackingToken_example", float32(123), "PrimaryIncomeSource_example", "PrivacyPolicyTrackingToken_example", "ResidenceType_example", "SoctTrackingToken_example", float32(123), "UserToken_example") // CreateApplicationsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.CreditApplicationsPost(context.Background()).CreateApplicationsRequest(createApplicationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreditApplicationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreditApplicationsPost`: ApplicationsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreditApplicationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditApplicationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApplicationsRequest** | [**CreateApplicationsRequest**](CreateApplicationsRequest.md) |  | 

### Return type

[**ApplicationsResponse**](ApplicationsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileByType

> FileResponse GetFileByType(ctx, type_).BundleToken(bundleToken).ApplicationToken(applicationToken).Execute()

Retrieve file on a bundle or application



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
    type_ := openapiclient.FileType("SOCT") // FileType | The type of file to retrieve.  * `SOCT` - The Summary of Credit Terms (SOCT), which outlines the interest rates, interest charges, and fees associated with the credit account being offered to the user. * `REWARDS_DISCLOSURE_PRE_TERMS` - The Rewards Disclosure Pre-terms, which discloses detailed information about the rewards program on the credit account being offered to the user before a decision is rendered on their application. * `REWARDS_DISCLOSURE_POST_TERMS` - The Rewards Disclosure Post-terms, which discloses detailed information about the rewards program on the user's credit account if their application is approved. * `BENEFITS_DISCLOSURE` - The Benefits Disclosure, which which is given to a user if their application is approved and discloses detailed information about the benefits on the user's credit account. * `CARD_MEMBER_AGREEMENT` - The Card Member Agreement, which specifies the terms and conditions of the user's credit account, including the interest rates, interest charges, fees, minimum payment calculations, and more. * `PRIVACY_POLICY` - The Privacy Policy, which explains how the information on the user's application is collected, handled, and processed. * `E_DISCLOSURE` - The eDisclosure, which states that the user has agreed to receive disclosures electronically. * `TERMS_SCHEDULE` - The Terms Schedule, which is given to a user if their application is approved and specifies the interest rate details on the user's credit account. * `NOAA` - The Notice of Adverse Action (NOAA), which is given to a user if their application is declined and informs them of the specific reasons why they were denied a credit account.
    bundleToken := "bundleToken_example" // string | Unique identifier of the bundle on which you want to retrieve a file.  Required if retrieving one of the following file types:  * `CARD_MEMBER_AGREEMENT` * `E_DISCLOSURE` * `PRIVACY_POLICY` * `REWARDS_DISCLOSURE_PRE_TERMS` * `REWARDS_DISCLOSURE_POST_TERMS` * `SOCT` (optional)
    applicationToken := "applicationToken_example" // string | Unique identifier of the application on which you want to retrieve a file.  Required if retrieving one of the following files:  * `BENEFITS_DISCLOSURE` * `NOAA` * `TERMS_SCHEDULE` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetFileByType(context.Background(), type_).BundleToken(bundleToken).ApplicationToken(applicationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetFileByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileByType`: FileResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetFileByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**FileType**](.md) | The type of file to retrieve.  * &#x60;SOCT&#x60; - The Summary of Credit Terms (SOCT), which outlines the interest rates, interest charges, and fees associated with the credit account being offered to the user. * &#x60;REWARDS_DISCLOSURE_PRE_TERMS&#x60; - The Rewards Disclosure Pre-terms, which discloses detailed information about the rewards program on the credit account being offered to the user before a decision is rendered on their application. * &#x60;REWARDS_DISCLOSURE_POST_TERMS&#x60; - The Rewards Disclosure Post-terms, which discloses detailed information about the rewards program on the user&#39;s credit account if their application is approved. * &#x60;BENEFITS_DISCLOSURE&#x60; - The Benefits Disclosure, which which is given to a user if their application is approved and discloses detailed information about the benefits on the user&#39;s credit account. * &#x60;CARD_MEMBER_AGREEMENT&#x60; - The Card Member Agreement, which specifies the terms and conditions of the user&#39;s credit account, including the interest rates, interest charges, fees, minimum payment calculations, and more. * &#x60;PRIVACY_POLICY&#x60; - The Privacy Policy, which explains how the information on the user&#39;s application is collected, handled, and processed. * &#x60;E_DISCLOSURE&#x60; - The eDisclosure, which states that the user has agreed to receive disclosures electronically. * &#x60;TERMS_SCHEDULE&#x60; - The Terms Schedule, which is given to a user if their application is approved and specifies the interest rate details on the user&#39;s credit account. * &#x60;NOAA&#x60; - The Notice of Adverse Action (NOAA), which is given to a user if their application is declined and informs them of the specific reasons why they were denied a credit account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleToken** | **string** | Unique identifier of the bundle on which you want to retrieve a file.  Required if retrieving one of the following file types:  * &#x60;CARD_MEMBER_AGREEMENT&#x60; * &#x60;E_DISCLOSURE&#x60; * &#x60;PRIVACY_POLICY&#x60; * &#x60;REWARDS_DISCLOSURE_PRE_TERMS&#x60; * &#x60;REWARDS_DISCLOSURE_POST_TERMS&#x60; * &#x60;SOCT&#x60; | 
 **applicationToken** | **string** | Unique identifier of the application on which you want to retrieve a file.  Required if retrieving one of the following files:  * &#x60;BENEFITS_DISCLOSURE&#x60; * &#x60;NOAA&#x60; * &#x60;TERMS_SCHEDULE&#x60; | 

### Return type

[**FileResponse**](FileResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageApplicationTransitions

> ApplicationsTransitionPage PageApplicationTransitions(ctx, token).Execute()

List application transitions



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
    token := "token_example" // string | The unique identifier of the application for which you want to retrieve transitions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.PageApplicationTransitions(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.PageApplicationTransitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageApplicationTransitions`: ApplicationsTransitionPage
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.PageApplicationTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The unique identifier of the application for which you want to retrieve transitions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageApplicationTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationsTransitionPage**](ApplicationsTransitionPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveApplication

> ApplicationsResponse RetrieveApplication(ctx, token).Expand(expand).Execute()

Retrieve application



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
    token := "token_example" // string | Unique identifier of the application to retrieve.
    expand := []openapiclient.ExpandObjects{openapiclient.ExpandObjects("DEVICE_DATA")} // []ExpandObjects | Embeds the specified object into the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.RetrieveApplication(context.Background(), token).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.RetrieveApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveApplication`: ApplicationsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.RetrieveApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the application to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | [**[]ExpandObjects**](ExpandObjects.md) | Embeds the specified object into the response. | 

### Return type

[**ApplicationsResponse**](ApplicationsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFiles

> map[string]FileResponse RetrieveFiles(ctx).BundleToken(bundleToken).ApplicationToken(applicationToken).Execute()

List files on a bundle or application



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
    bundleToken := "bundleToken_example" // string | Unique identifier of the bundle whose files you want to retrieve.  The following file types are returned with the `bundle_token`:  * `CARD_MEMBER_AGREEMENT` * `E_DISCLOSURE` * `PRIVACY_POLICY` * `REWARDS_DISCLOSURE_PRE_TERMS` * `REWARDS_DISCLOSURE_POST_TERMS` * `SOCT` (optional)
    applicationToken := "applicationToken_example" // string | Unique identifier of the application whose files you want to retrieve.  The following file types are returned with the `application_token`:  * `BENEFITS_DISCLOSURE` * `CARD_MEMBER_AGREEMENT` * `E_DISCLOSURE` * `NOAA` * `PRIVACY_POLICY` * `REWARDS_DISCLOSURE_PRE_TERMS` * `REWARDS_DISCLOSURE_POST_TERMS` * `SOCT` * `TERMS_SCHEDULE` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.RetrieveFiles(context.Background()).BundleToken(bundleToken).ApplicationToken(applicationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.RetrieveFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveFiles`: map[string]FileResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.RetrieveFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleToken** | **string** | Unique identifier of the bundle whose files you want to retrieve.  The following file types are returned with the &#x60;bundle_token&#x60;:  * &#x60;CARD_MEMBER_AGREEMENT&#x60; * &#x60;E_DISCLOSURE&#x60; * &#x60;PRIVACY_POLICY&#x60; * &#x60;REWARDS_DISCLOSURE_PRE_TERMS&#x60; * &#x60;REWARDS_DISCLOSURE_POST_TERMS&#x60; * &#x60;SOCT&#x60; | 
 **applicationToken** | **string** | Unique identifier of the application whose files you want to retrieve.  The following file types are returned with the &#x60;application_token&#x60;:  * &#x60;BENEFITS_DISCLOSURE&#x60; * &#x60;CARD_MEMBER_AGREEMENT&#x60; * &#x60;E_DISCLOSURE&#x60; * &#x60;NOAA&#x60; * &#x60;PRIVACY_POLICY&#x60; * &#x60;REWARDS_DISCLOSURE_PRE_TERMS&#x60; * &#x60;REWARDS_DISCLOSURE_POST_TERMS&#x60; * &#x60;SOCT&#x60; * &#x60;TERMS_SCHEDULE&#x60; | 

### Return type

[**map[string]FileResponse**](FileResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionApplication

> ApplicationTransitionResponse TransitionApplication(ctx, token).ApplicationTransitionRequest(applicationTransitionRequest).Execute()

Transition application state



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
    token := "token_example" // string | Unique identifier of the application whose state you want to transition.
    applicationTransitionRequest := *openapiclient.NewApplicationTransitionRequest(openapiclient.ApplicationResourceState("CREATED")) // ApplicationTransitionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.TransitionApplication(context.Background(), token).ApplicationTransitionRequest(applicationTransitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.TransitionApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransitionApplication`: ApplicationTransitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.TransitionApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the application whose state you want to transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationTransitionRequest** | [**ApplicationTransitionRequest**](ApplicationTransitionRequest.md) |  | 

### Return type

[**ApplicationTransitionResponse**](ApplicationTransitionResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

