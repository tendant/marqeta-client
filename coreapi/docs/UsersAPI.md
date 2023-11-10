# \UsersAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsers**](UsersAPI.md#GetUsers) | **Get** /users | List users
[**GetUsersAuthClientaccesstokenToken**](UsersAPI.md#GetUsersAuthClientaccesstokenToken) | **Get** /users/auth/clientaccesstoken/{token} | Retrieve client access token
[**GetUsersParenttokenChildren**](UsersAPI.md#GetUsersParenttokenChildren) | **Get** /users/{parent_token}/children | List user child accounts
[**GetUsersToken**](UsersAPI.md#GetUsersToken) | **Get** /users/{token} | Retrieve user
[**GetUsersTokenSsn**](UsersAPI.md#GetUsersTokenSsn) | **Get** /users/{token}/ssn | Retrieve user identification number
[**PostUsers**](UsersAPI.md#PostUsers) | **Post** /users | Create user
[**PostUsersAuthChangepassword**](UsersAPI.md#PostUsersAuthChangepassword) | **Post** /users/auth/changepassword | Update user password
[**PostUsersAuthClientaccesstoken**](UsersAPI.md#PostUsersAuthClientaccesstoken) | **Post** /users/auth/clientaccesstoken | Create client access token
[**PostUsersAuthLogin**](UsersAPI.md#PostUsersAuthLogin) | **Post** /users/auth/login | Log in user
[**PostUsersAuthLogout**](UsersAPI.md#PostUsersAuthLogout) | **Post** /users/auth/logout | Log out user
[**PostUsersAuthOnetime**](UsersAPI.md#PostUsersAuthOnetime) | **Post** /users/auth/onetime | Create single-use token
[**PostUsersAuthResetpassword**](UsersAPI.md#PostUsersAuthResetpassword) | **Post** /users/auth/resetpassword | Request user password reset token
[**PostUsersAuthResetpasswordToken**](UsersAPI.md#PostUsersAuthResetpasswordToken) | **Post** /users/auth/resetpassword/{token} | Reset user password
[**PostUsersAuthVerifyemail**](UsersAPI.md#PostUsersAuthVerifyemail) | **Post** /users/auth/verifyemail | Request email verification token
[**PostUsersAuthVerifyemailToken**](UsersAPI.md#PostUsersAuthVerifyemailToken) | **Post** /users/auth/verifyemail/{token} | Verify email address
[**PostUsersLookup**](UsersAPI.md#PostUsersLookup) | **Post** /users/lookup | Search users
[**PutUsersToken**](UsersAPI.md#PutUsersToken) | **Put** /users/{token} | Update user



## GetUsers

> UserCardHolderListResponse GetUsers(ctx).Count(count).StartIndex(startIndex).SearchType(searchType).Fields(fields).SortBy(sortBy).Execute()

List users



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
    count := int32(56) // int32 | Number of user resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    searchType := "searchType_example" // string | Search type. (optional)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.GetUsers(context.Background()).Count(count).StartIndex(startIndex).SearchType(searchType).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: UserCardHolderListResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of user resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **searchType** | **string** | Search type. | 
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**UserCardHolderListResponse**](UserCardHolderListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersAuthClientaccesstokenToken

> ClientAccessTokenResponse GetUsersAuthClientaccesstokenToken(ctx, token).ApplicationToken(applicationToken).Execute()

Retrieve client access token



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
    token := "token_example" // string | Client access token.
    applicationToken := "applicationToken_example" // string | Unique identifier of the `application` object. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.GetUsersAuthClientaccesstokenToken(context.Background(), token).ApplicationToken(applicationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersAuthClientaccesstokenToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersAuthClientaccesstokenToken`: ClientAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersAuthClientaccesstokenToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Client access token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersAuthClientaccesstokenTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationToken** | **string** | Unique identifier of the &#x60;application&#x60; object. | 

### Return type

[**ClientAccessTokenResponse**](ClientAccessTokenResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersParenttokenChildren

> UserCardHolderListResponse GetUsersParenttokenChildren(ctx, parentToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()

List user child accounts



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
    parentToken := "parentToken_example" // string | Unique identifier of the parent account holder.
    count := int32(56) // int32 | Number of user resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.GetUsersParenttokenChildren(context.Background(), parentToken).Count(count).StartIndex(startIndex).Fields(fields).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersParenttokenChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersParenttokenChildren`: UserCardHolderListResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersParenttokenChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentToken** | **string** | Unique identifier of the parent account holder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersParenttokenChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of user resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**UserCardHolderListResponse**](UserCardHolderListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersToken

> UserCardHolderResponse GetUsersToken(ctx, token).Fields(fields).Execute()

Retrieve user



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
    token := "token_example" // string | Unique identifier of the user resource.
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.GetUsersToken(context.Background(), token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersToken`: UserCardHolderResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the user resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 

### Return type

[**UserCardHolderResponse**](UserCardHolderResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersTokenSsn

> SsnResponseModel GetUsersTokenSsn(ctx, token).FullSsn(fullSsn).Execute()

Retrieve user identification number



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
    token := "token_example" // string | Unique identifier of the user resource.
    fullSsn := true // bool | To return the full identification number, set to `true`. To return only the last four digits, set to `false`.  If the identifications array contains only the last four digits of the identification number, the `/users/{token}/ssn` endpoint will return only the last four digits, regardless of the `full_ssn` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.GetUsersTokenSsn(context.Background(), token).FullSsn(fullSsn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersTokenSsn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersTokenSsn`: SsnResponseModel
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersTokenSsn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the user resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersTokenSsnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullSsn** | **bool** | To return the full identification number, set to &#x60;true&#x60;. To return only the last four digits, set to &#x60;false&#x60;.  If the identifications array contains only the last four digits of the identification number, the &#x60;/users/{token}/ssn&#x60; endpoint will return only the last four digits, regardless of the &#x60;full_ssn&#x60; parameter. | 

### Return type

[**SsnResponseModel**](SsnResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsers

> UserCardHolderResponse PostUsers(ctx).CardHolderModel(cardHolderModel).Execute()

Create user



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
    cardHolderModel := *openapiclient.NewCardHolderModel() // CardHolderModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.PostUsers(context.Background()).CardHolderModel(cardHolderModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsers`: UserCardHolderResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardHolderModel** | [**CardHolderModel**](CardHolderModel.md) |  | 

### Return type

[**UserCardHolderResponse**](UserCardHolderResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthChangepassword

> PostUsersAuthChangepassword(ctx).PasswordUpdateModel(passwordUpdateModel).Execute()

Update user password



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
    passwordUpdateModel := *openapiclient.NewPasswordUpdateModel("CurrentPassword_example", "NewPassword_example") // PasswordUpdateModel | Password update object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.PostUsersAuthChangepassword(context.Background()).PasswordUpdateModel(passwordUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthChangepassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthChangepasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordUpdateModel** | [**PasswordUpdateModel**](PasswordUpdateModel.md) | Password update object | 

### Return type

 (empty response body)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthClientaccesstoken

> ClientAccessTokenResponse PostUsersAuthClientaccesstoken(ctx).ClientAccessTokenRequest(clientAccessTokenRequest).Execute()

Create client access token



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
    clientAccessTokenRequest := *openapiclient.NewClientAccessTokenRequest("CardToken_example") // ClientAccessTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.PostUsersAuthClientaccesstoken(context.Background()).ClientAccessTokenRequest(clientAccessTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthClientaccesstoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsersAuthClientaccesstoken`: ClientAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PostUsersAuthClientaccesstoken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthClientaccesstokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientAccessTokenRequest** | [**ClientAccessTokenRequest**](ClientAccessTokenRequest.md) |  | 

### Return type

[**ClientAccessTokenResponse**](ClientAccessTokenResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthLogin

> LoginResponseModel PostUsersAuthLogin(ctx).LoginRequestModel(loginRequestModel).Execute()

Log in user



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
    loginRequestModel := *openapiclient.NewLoginRequestModel() // LoginRequestModel | User login object (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.PostUsersAuthLogin(context.Background()).LoginRequestModel(loginRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsersAuthLogin`: LoginResponseModel
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PostUsersAuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequestModel** | [**LoginRequestModel**](LoginRequestModel.md) | User login object | 

### Return type

[**LoginResponseModel**](LoginResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthLogout

> PostUsersAuthLogout(ctx).Execute()

Log out user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.PostUsersAuthLogout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthLogout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthLogoutRequest struct via the builder pattern


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


## PostUsersAuthOnetime

> AccessTokenResponse PostUsersAuthOnetime(ctx).OneTimeRequestModel(oneTimeRequestModel).Execute()

Create single-use token



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
    oneTimeRequestModel := *openapiclient.NewOneTimeRequestModel() // OneTimeRequestModel | One-time object (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.PostUsersAuthOnetime(context.Background()).OneTimeRequestModel(oneTimeRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthOnetime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsersAuthOnetime`: AccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PostUsersAuthOnetime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthOnetimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oneTimeRequestModel** | [**OneTimeRequestModel**](OneTimeRequestModel.md) | One-time object | 

### Return type

[**AccessTokenResponse**](AccessTokenResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthResetpassword

> PostUsersAuthResetpassword(ctx).ResetUserPasswordEmailModel(resetUserPasswordEmailModel).Execute()

Request user password reset token



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
    resetUserPasswordEmailModel := *openapiclient.NewResetUserPasswordEmailModel("Email_example") // ResetUserPasswordEmailModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.PostUsersAuthResetpassword(context.Background()).ResetUserPasswordEmailModel(resetUserPasswordEmailModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthResetpassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthResetpasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetUserPasswordEmailModel** | [**ResetUserPasswordEmailModel**](ResetUserPasswordEmailModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthResetpasswordToken

> PostUsersAuthResetpasswordToken(ctx, token).ResetUserPasswordModel(resetUserPasswordModel).Execute()

Reset user password



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
    token := "token_example" // string | Password reset token generated using the `POST /users/auth/resetpassword` operation.
    resetUserPasswordModel := *openapiclient.NewResetUserPasswordModel("NewPassword_example", "UserToken_example") // ResetUserPasswordModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.PostUsersAuthResetpasswordToken(context.Background(), token).ResetUserPasswordModel(resetUserPasswordModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthResetpasswordToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Password reset token generated using the &#x60;POST /users/auth/resetpassword&#x60; operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthResetpasswordTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetUserPasswordModel** | [**ResetUserPasswordModel**](ResetUserPasswordModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthVerifyemail

> PostUsersAuthVerifyemail(ctx).Execute()

Request email verification token



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.PostUsersAuthVerifyemail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthVerifyemail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthVerifyemailRequest struct via the builder pattern


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


## PostUsersAuthVerifyemailToken

> PostUsersAuthVerifyemailToken(ctx, token).Execute()

Verify email address



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
    token := "token_example" // string | Email verification token generated using the `POST /users/auth/verifyemail` operation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.PostUsersAuthVerifyemailToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersAuthVerifyemailToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Email verification token generated using the &#x60;POST /users/auth/verifyemail&#x60; operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthVerifyemailTokenRequest struct via the builder pattern


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


## PostUsersLookup

> UserCardHolderListResponse PostUsersLookup(ctx).Count(count).StartIndex(startIndex).SearchType(searchType).Fields(fields).SortBy(sortBy).UserCardHolderSearchModel(userCardHolderSearchModel).Execute()

Search users



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
    count := int32(56) // int32 | Number of user resources to retrieve. (optional) (default to 5)
    startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
    searchType := "searchType_example" // string | Search type. (optional)
    fields := "fields_example" // string | Comma-delimited list of fields to return (`field_1,field_2`, and so on). Leave blank to return all fields. (optional)
    sortBy := "sortBy_example" // string | Field on which to sort. Use any field in the resource model, or one of the system fields `lastModifiedTime` or `createdTime`. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-lastModifiedTime")
    userCardHolderSearchModel := *openapiclient.NewUserCardHolderSearchModel() // UserCardHolderSearchModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.PostUsersLookup(context.Background()).Count(count).StartIndex(startIndex).SearchType(searchType).Fields(fields).SortBy(sortBy).UserCardHolderSearchModel(userCardHolderSearchModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PostUsersLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsersLookup`: UserCardHolderListResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PostUsersLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of user resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **searchType** | **string** | Search type. | 
 **fields** | **string** | Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields. | 
 **sortBy** | **string** | Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-lastModifiedTime&quot;]
 **userCardHolderSearchModel** | [**UserCardHolderSearchModel**](UserCardHolderSearchModel.md) |  | 

### Return type

[**UserCardHolderListResponse**](UserCardHolderListResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUsersToken

> CardHolderModel PutUsersToken(ctx, token).UserCardHolderUpdateModel(userCardHolderUpdateModel).Execute()

Update user



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
    token := "token_example" // string | Unique identifier of the user resource you want to update.
    userCardHolderUpdateModel := *openapiclient.NewUserCardHolderUpdateModel() // UserCardHolderUpdateModel | User object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.PutUsersToken(context.Background(), token).UserCardHolderUpdateModel(userCardHolderUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PutUsersToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUsersToken`: CardHolderModel
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PutUsersToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the user resource you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUsersTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCardHolderUpdateModel** | [**UserCardHolderUpdateModel**](UserCardHolderUpdateModel.md) | User object | 

### Return type

[**CardHolderModel**](CardHolderModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

