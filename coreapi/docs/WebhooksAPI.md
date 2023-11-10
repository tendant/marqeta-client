# \WebhooksAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhooks**](WebhooksAPI.md#GetWebhooks) | **Get** /webhooks | List webhooks
[**GetWebhooksToken**](WebhooksAPI.md#GetWebhooksToken) | **Get** /webhooks/{token} | Retrieve webhook
[**PostWebhooks**](WebhooksAPI.md#PostWebhooks) | **Post** /webhooks | Create webhook
[**PostWebhooksTokenEventtypeEventtoken**](WebhooksAPI.md#PostWebhooksTokenEventtypeEventtoken) | **Post** /webhooks/{token}/{event_type}/{event_token} | Resend event notification
[**PostWebhooksTokenPing**](WebhooksAPI.md#PostWebhooksTokenPing) | **Post** /webhooks/{token}/ping | Ping webhook
[**PutWebhooksCustomHeadersToken**](WebhooksAPI.md#PutWebhooksCustomHeadersToken) | **Put** /webhooks/customheaders/{token} | Update webhook custom headers
[**PutWebhooksToken**](WebhooksAPI.md#PutWebhooksToken) | **Put** /webhooks/{token} | Update webhook



## GetWebhooks

> WebhookResponseModelListResponse GetWebhooks(ctx).Active(active).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List webhooks



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
    active := true // bool | Set to `true` to return only active webhook configurations. (optional) (default to false)
    count := int32(56) // int32 | Number of resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.GetWebhooks(context.Background()).Active(active).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooks`: WebhookResponseModelListResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | Set to &#x60;true&#x60; to return only active webhook configurations. | [default to false]
 **count** | **int32** | Number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**WebhookResponseModelListResponse**](WebhookResponseModelListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksToken

> WebhookResponseModel GetWebhooksToken(ctx, token).Execute()

Retrieve webhook



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
    token := "token_example" // string | Unique identifier of the webhook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.GetWebhooksToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhooksToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksToken`: WebhookResponseModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhooksToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the webhook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookResponseModel**](WebhookResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooks

> WebhookResponseModel PostWebhooks(ctx).WebhookRequestModel(webhookRequestModel).Execute()

Create webhook



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
    webhookRequestModel := *openapiclient.NewWebhookRequestModel(*openapiclient.NewWebhookConfigModel("BasicAuthPassword_example", "BasicAuthUsername_example", "Url_example"), []string{"Events_example"}, "Name_example") // WebhookRequestModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.PostWebhooks(context.Background()).WebhookRequestModel(webhookRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWebhooks`: WebhookResponseModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PostWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookRequestModel** | [**WebhookRequestModel**](WebhookRequestModel.md) |  | 

### Return type

[**WebhookResponseModel**](WebhookResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksTokenEventtypeEventtoken

> PostWebhooksTokenEventtypeEventtoken(ctx, token, eventType, eventToken).Execute()

Resend event notification



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
    token := "token_example" // string | Unique identifier of the webhook.
    eventType := "eventType_example" // string | Specifies the type of event you want to resend.
    eventToken := "eventToken_example" // string | Unique identifier of the event for which you want to resend a notification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksAPI.PostWebhooksTokenEventtypeEventtoken(context.Background(), token, eventType, eventToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostWebhooksTokenEventtypeEventtoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the webhook. | 
**eventType** | **string** | Specifies the type of event you want to resend. | 
**eventToken** | **string** | Unique identifier of the event for which you want to resend a notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksTokenEventtypeEventtokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksTokenPing

> WebhookPingModel PostWebhooksTokenPing(ctx, token).Execute()

Ping webhook



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
    token := "token_example" // string | Unique identifier of the webhook.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.PostWebhooksTokenPing(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostWebhooksTokenPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWebhooksTokenPing`: WebhookPingModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PostWebhooksTokenPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the webhook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksTokenPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookPingModel**](WebhookPingModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhooksCustomHeadersToken

> WebhookResponseModel PutWebhooksCustomHeadersToken(ctx, token).WebhookUpdateCustomHeaderRequest(webhookUpdateCustomHeaderRequest).Execute()

Update webhook custom headers



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
    token := "token_example" // string | Unique identifier of the webhook.
    webhookUpdateCustomHeaderRequest := *openapiclient.NewWebhookUpdateCustomHeaderRequest() // WebhookUpdateCustomHeaderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.PutWebhooksCustomHeadersToken(context.Background(), token).WebhookUpdateCustomHeaderRequest(webhookUpdateCustomHeaderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PutWebhooksCustomHeadersToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWebhooksCustomHeadersToken`: WebhookResponseModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PutWebhooksCustomHeadersToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the webhook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksCustomHeadersTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookUpdateCustomHeaderRequest** | [**WebhookUpdateCustomHeaderRequest**](WebhookUpdateCustomHeaderRequest.md) |  | 

### Return type

[**WebhookResponseModel**](WebhookResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhooksToken

> WebhookResponseModel PutWebhooksToken(ctx, token).WebhookBaseModel(webhookBaseModel).Execute()

Update webhook



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
    token := "token_example" // string | Unique identifier of the webhook.
    webhookBaseModel := *openapiclient.NewWebhookBaseModel(*openapiclient.NewWebhookConfigModel("BasicAuthPassword_example", "BasicAuthUsername_example", "Url_example"), []string{"Events_example"}, "Name_example") // WebhookBaseModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.PutWebhooksToken(context.Background(), token).WebhookBaseModel(webhookBaseModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PutWebhooksToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutWebhooksToken`: WebhookResponseModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PutWebhooksToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the webhook. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookBaseModel** | [**WebhookBaseModel**](WebhookBaseModel.md) |  | 

### Return type

[**WebhookResponseModel**](WebhookResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

