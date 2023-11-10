# ClientAccessTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**CardToken** | Pointer to **string** | Unique identifier of the card whose sensitive information you want to display. | [optional] 
**Created** | **time.Time** | Date and time that the client access token was created, in UTC. | 
**Expires** | **time.Time** | Date and time that the client access token expires, in UTC. | 
**Token** | Pointer to **string** | Value of the client access token to redeem when displaying sensitive card data. | [optional] 

## Methods

### NewClientAccessTokenResponse

`func NewClientAccessTokenResponse(created time.Time, expires time.Time, ) *ClientAccessTokenResponse`

NewClientAccessTokenResponse instantiates a new ClientAccessTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAccessTokenResponseWithDefaults

`func NewClientAccessTokenResponseWithDefaults() *ClientAccessTokenResponse`

NewClientAccessTokenResponseWithDefaults instantiates a new ClientAccessTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ClientAccessTokenResponse) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ClientAccessTokenResponse) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ClientAccessTokenResponse) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ClientAccessTokenResponse) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCardToken

`func (o *ClientAccessTokenResponse) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *ClientAccessTokenResponse) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *ClientAccessTokenResponse) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.

### HasCardToken

`func (o *ClientAccessTokenResponse) HasCardToken() bool`

HasCardToken returns a boolean if a field has been set.

### GetCreated

`func (o *ClientAccessTokenResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClientAccessTokenResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClientAccessTokenResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *ClientAccessTokenResponse) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ClientAccessTokenResponse) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ClientAccessTokenResponse) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetToken

`func (o *ClientAccessTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientAccessTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientAccessTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClientAccessTokenResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


