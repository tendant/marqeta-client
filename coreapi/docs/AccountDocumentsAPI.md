# \AccountDocumentsAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountDocuments**](AccountDocumentsAPI.md#GetAccountDocuments) | **Get** /credit/accounts/{account_token}/documents | List documents
[**GetDocumentByAccountAndType**](AccountDocumentsAPI.md#GetDocumentByAccountAndType) | **Get** /credit/accounts/{account_token}/documents/{document_type} | Retrieve document
[**GetDocumentHistoryByAccountAndType**](AccountDocumentsAPI.md#GetDocumentHistoryByAccountAndType) | **Get** /credit/accounts/{account_token}/documents/{document_type}/history | Retrieve document history



## GetAccountDocuments

> AccountDocumentsResponse GetAccountDocuments(ctx, accountToken).Execute()

List documents



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to get documents.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountDocumentsAPI.GetAccountDocuments(context.Background(), accountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountDocumentsAPI.GetAccountDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountDocuments`: AccountDocumentsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountDocumentsAPI.GetAccountDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to get documents.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountDocumentsResponse**](AccountDocumentsResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentByAccountAndType

> AccountDocumentResponse GetDocumentByAccountAndType(ctx, accountToken, documentType).Execute()

Retrieve document



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to retrieve a specific type of document.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    documentType := openapiclient.AccountAndDocumentAssetType("TERMS_SCHEDULE") // AccountAndDocumentAssetType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountDocumentsAPI.GetDocumentByAccountAndType(context.Background(), accountToken, documentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountDocumentsAPI.GetDocumentByAccountAndType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentByAccountAndType`: AccountDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountDocumentsAPI.GetDocumentByAccountAndType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to retrieve a specific type of document.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**documentType** | [**AccountAndDocumentAssetType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentByAccountAndTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountDocumentResponse**](AccountDocumentResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentHistoryByAccountAndType

> AccountDocumentsPage GetDocumentHistoryByAccountAndType(ctx, accountToken, documentType).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

Retrieve document history



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which to get document history.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    documentType := openapiclient.AccountAndDocumentAssetType("TERMS_SCHEDULE") // AccountAndDocumentAssetType | 
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `effectiveDate`, and not by the field names appearing in response bodies such as `effective_date`. (optional) (default to "-effectiveDate")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountDocumentsAPI.GetDocumentHistoryByAccountAndType(context.Background(), accountToken, documentType).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountDocumentsAPI.GetDocumentHistoryByAccountAndType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentHistoryByAccountAndType`: AccountDocumentsPage
    fmt.Fprintf(os.Stdout, "Response from `AccountDocumentsAPI.GetDocumentHistoryByAccountAndType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which to get document history.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**documentType** | [**AccountAndDocumentAssetType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentHistoryByAccountAndTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;effectiveDate&#x60;, and not by the field names appearing in response bodies such as &#x60;effective_date&#x60;. | [default to &quot;-effectiveDate&quot;]

### Return type

[**AccountDocumentsPage**](AccountDocumentsPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

