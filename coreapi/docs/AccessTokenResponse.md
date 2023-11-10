# AccessTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccesstokenId** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**Expires** | **time.Time** | Date and time when the access token expires. | 
**MasterRoles** | Pointer to **[]string** | Array of master roles. | [optional] 
**OneTime** | Pointer to **bool** | Indicates whether the access token is a single-use token. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the access token. | [optional] 
**TokenType** | Pointer to **string** | Specifies the type of access token. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user resource. | [optional] 

## Methods

### NewAccessTokenResponse

`func NewAccessTokenResponse(expires time.Time, ) *AccessTokenResponse`

NewAccessTokenResponse instantiates a new AccessTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenResponseWithDefaults

`func NewAccessTokenResponseWithDefaults() *AccessTokenResponse`

NewAccessTokenResponseWithDefaults instantiates a new AccessTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccesstokenId

`func (o *AccessTokenResponse) GetAccesstokenId() string`

GetAccesstokenId returns the AccesstokenId field if non-nil, zero value otherwise.

### GetAccesstokenIdOk

`func (o *AccessTokenResponse) GetAccesstokenIdOk() (*string, bool)`

GetAccesstokenIdOk returns a tuple with the AccesstokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesstokenId

`func (o *AccessTokenResponse) SetAccesstokenId(v string)`

SetAccesstokenId sets AccesstokenId field to given value.

### HasAccesstokenId

`func (o *AccessTokenResponse) HasAccesstokenId() bool`

HasAccesstokenId returns a boolean if a field has been set.

### GetApplication

`func (o *AccessTokenResponse) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *AccessTokenResponse) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *AccessTokenResponse) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *AccessTokenResponse) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetExpires

`func (o *AccessTokenResponse) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AccessTokenResponse) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AccessTokenResponse) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetMasterRoles

`func (o *AccessTokenResponse) GetMasterRoles() []string`

GetMasterRoles returns the MasterRoles field if non-nil, zero value otherwise.

### GetMasterRolesOk

`func (o *AccessTokenResponse) GetMasterRolesOk() (*[]string, bool)`

GetMasterRolesOk returns a tuple with the MasterRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRoles

`func (o *AccessTokenResponse) SetMasterRoles(v []string)`

SetMasterRoles sets MasterRoles field to given value.

### HasMasterRoles

`func (o *AccessTokenResponse) HasMasterRoles() bool`

HasMasterRoles returns a boolean if a field has been set.

### GetOneTime

`func (o *AccessTokenResponse) GetOneTime() bool`

GetOneTime returns the OneTime field if non-nil, zero value otherwise.

### GetOneTimeOk

`func (o *AccessTokenResponse) GetOneTimeOk() (*bool, bool)`

GetOneTimeOk returns a tuple with the OneTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTime

`func (o *AccessTokenResponse) SetOneTime(v bool)`

SetOneTime sets OneTime field to given value.

### HasOneTime

`func (o *AccessTokenResponse) HasOneTime() bool`

HasOneTime returns a boolean if a field has been set.

### GetToken

`func (o *AccessTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccessTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccessTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccessTokenResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenType

`func (o *AccessTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AccessTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AccessTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AccessTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUserToken

`func (o *AccessTokenResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AccessTokenResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AccessTokenResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *AccessTokenResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


