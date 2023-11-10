# OneTimeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Cardholder email address.  Required when neither the user token nor the admin access token is provided as the Basic Authentication password (case #3). | [optional] 
**Password** | Pointer to **string** | Password to the cardholder&#39;s user account on the Marqeta platform.  Required when neither the user token nor the admin access token is provided as the Basic Authentication password (case #3). | [optional] 
**UserToken** | Pointer to **string** | Identifies the cardholder whose data is accessed. Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve cardholder tokens.  Required when the Basic Authentication password is set to an admin access token (case #2). | [optional] 

## Methods

### NewOneTimeRequestModel

`func NewOneTimeRequestModel() *OneTimeRequestModel`

NewOneTimeRequestModel instantiates a new OneTimeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneTimeRequestModelWithDefaults

`func NewOneTimeRequestModelWithDefaults() *OneTimeRequestModel`

NewOneTimeRequestModelWithDefaults instantiates a new OneTimeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OneTimeRequestModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OneTimeRequestModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OneTimeRequestModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OneTimeRequestModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *OneTimeRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OneTimeRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OneTimeRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OneTimeRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserToken

`func (o *OneTimeRequestModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *OneTimeRequestModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *OneTimeRequestModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *OneTimeRequestModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


