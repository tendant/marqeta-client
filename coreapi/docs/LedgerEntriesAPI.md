# \LedgerEntriesAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountLedgerEntry**](LedgerEntriesAPI.md#GetAccountLedgerEntry) | **Get** /credit/accounts/{account_token}/ledgerentries/{ledger_entry_token} | Retrieve account ledger entry
[**ListAccountLedgerEntries**](LedgerEntriesAPI.md#ListAccountLedgerEntries) | **Get** /credit/accounts/{account_token}/ledgerentries | List account ledger entries
[**ResendWebhookEvent**](LedgerEntriesAPI.md#ResendWebhookEvent) | **Post** /credit/webhooks/{event_type}/{resource_token} | Resend credit event notification



## GetAccountLedgerEntry

> LedgerEntry GetAccountLedgerEntry(ctx, accountToken, ledgerEntryToken).Execute()

Retrieve account ledger entry



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
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve ledger entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    ledgerEntryToken := "ledgerEntryToken_example" // string | The unique identifier of the ledger entry you want to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/ledgerentries` to retrieve existing ledger entry tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LedgerEntriesAPI.GetAccountLedgerEntry(context.Background(), accountToken, ledgerEntryToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerEntriesAPI.GetAccountLedgerEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountLedgerEntry`: LedgerEntry
    fmt.Fprintf(os.Stdout, "Response from `LedgerEntriesAPI.GetAccountLedgerEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve ledger entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**ledgerEntryToken** | **string** | The unique identifier of the ledger entry you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/ledgerentries&#x60; to retrieve existing ledger entry tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLedgerEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LedgerEntry**](LedgerEntry.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountLedgerEntries

> LedgerEntriesPage ListAccountLedgerEntries(ctx, accountToken).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).Statuses(statuses).Description(description).Groups(groups).Amount(amount).Expand(expand).SortBy(sortBy).Execute()

List account ledger entries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/tendant/marqeta-client/coreapi"
)

func main() {
    accountToken := "accountToken_example" // string | The unique identifier of the credit account for which you want to retrieve ledger entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
    count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    startDate := time.Now() // string | The starting date of the date range from which to return ledger entries. (optional)
    endDate := time.Now() // string | The ending date of the date range from which to return ledger entries. (optional)
    statuses := []string{"Inner_example"} // []string | An array of statuses by which to filter ledger entries. (optional)
    description := "description_example" // string | The description of the ledger entries to return. (optional)
    groups := []string{"Inner_example"} // []string | An array of groups by which to filter ledger entries.  To return all ledger entry groups, do not include this query parameter. (optional)
    amount := float32(8.14) // float32 | The number of ledger entries to return. (optional)
    expand := "expand_example" // string | Embeds the specified object into the response. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `createdTime`, and not by the field names appearing in response bodies such as `created_time`. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LedgerEntriesAPI.ListAccountLedgerEntries(context.Background(), accountToken).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).Statuses(statuses).Description(description).Groups(groups).Amount(amount).Expand(expand).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerEntriesAPI.ListAccountLedgerEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountLedgerEntries`: LedgerEntriesPage
    fmt.Fprintf(os.Stdout, "Response from `LedgerEntriesAPI.ListAccountLedgerEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | The unique identifier of the credit account for which you want to retrieve ledger entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountLedgerEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **startDate** | **string** | The starting date of the date range from which to return ledger entries. | 
 **endDate** | **string** | The ending date of the date range from which to return ledger entries. | 
 **statuses** | **[]string** | An array of statuses by which to filter ledger entries. | 
 **description** | **string** | The description of the ledger entries to return. | 
 **groups** | **[]string** | An array of groups by which to filter ledger entries.  To return all ledger entry groups, do not include this query parameter. | 
 **amount** | **float32** | The number of ledger entries to return. | 
 **expand** | **string** | Embeds the specified object into the response. | 
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]

### Return type

[**LedgerEntriesPage**](LedgerEntriesPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendWebhookEvent

> WebhookEventResendContainerResponse ResendWebhookEvent(ctx, eventType, resourceToken).Execute()

Resend credit event notification



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
    eventType := "eventType_example" // string | Specifies the type of event you want to resend.
    resourceToken := "resourceToken_example" // string | The unique identifier of the resource for which you want to resend a notification.  Send a `GET` request to `/credit/accounts/{account_token}/journalentries` to retrieve existing journal entry tokens.  Send a `GET` request to `/credit/accounts/{account_token}/ledgerentries` to retrieve existing ledger entry tokens.  Send a `GET` request to `/accounts/{account_token}/accounttransitions` to retrieve existing account transition tokens.  Send a `GET` request to `/credit/accounts/{account_token}/payments/{payment_token}` to retrieve existing payment transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements` to retrieve existing statement summary tokens.  Send a `GET` request to `/accounts/{account_token}/delinquencystate/transitions` to retrieve existing delinquency state transition tokens.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LedgerEntriesAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerEntriesAPI.ResendWebhookEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendWebhookEvent`: WebhookEventResendContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerEntriesAPI.ResendWebhookEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventType** | **string** | Specifies the type of event you want to resend. | 
**resourceToken** | **string** | The unique identifier of the resource for which you want to resend a notification.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/journalentries&#x60; to retrieve existing journal entry tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/ledgerentries&#x60; to retrieve existing ledger entry tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/accounttransitions&#x60; to retrieve existing account transition tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/payments/{payment_token}&#x60; to retrieve existing payment transition tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/statements&#x60; to retrieve existing statement summary tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/delinquencystate/transitions&#x60; to retrieve existing delinquency state transition tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebhookEventResendContainerResponse**](WebhookEventResendContainerResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

