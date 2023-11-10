# LoginResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to [**AccessTokenResponse**](AccessTokenResponse.md) |  | [optional] 
**User** | Pointer to [**UserCardHolderResponse**](UserCardHolderResponse.md) |  | [optional] 

## Methods

### NewLoginResponseModel

`func NewLoginResponseModel() *LoginResponseModel`

NewLoginResponseModel instantiates a new LoginResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginResponseModelWithDefaults

`func NewLoginResponseModelWithDefaults() *LoginResponseModel`

NewLoginResponseModelWithDefaults instantiates a new LoginResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *LoginResponseModel) GetAccessToken() AccessTokenResponse`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *LoginResponseModel) GetAccessTokenOk() (*AccessTokenResponse, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *LoginResponseModel) SetAccessToken(v AccessTokenResponse)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *LoginResponseModel) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetUser

`func (o *LoginResponseModel) GetUser() UserCardHolderResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoginResponseModel) GetUserOk() (*UserCardHolderResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoginResponseModel) SetUser(v UserCardHolderResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *LoginResponseModel) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


