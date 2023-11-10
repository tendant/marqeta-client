# LoginRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Cardholder email address. | [optional] 
**Password** | Pointer to **string** | Password to the cardholder&#39;s user account on the Marqeta platform. | [optional] 
**UserToken** | Pointer to **string** | Identifies the cardholder to log in.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | [optional] 

## Methods

### NewLoginRequestModel

`func NewLoginRequestModel() *LoginRequestModel`

NewLoginRequestModel instantiates a new LoginRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginRequestModelWithDefaults

`func NewLoginRequestModelWithDefaults() *LoginRequestModel`

NewLoginRequestModelWithDefaults instantiates a new LoginRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LoginRequestModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoginRequestModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoginRequestModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoginRequestModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *LoginRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoginRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserToken

`func (o *LoginRequestModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *LoginRequestModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *LoginRequestModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *LoginRequestModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


