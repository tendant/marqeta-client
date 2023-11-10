# ResetUserPasswordModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | **string** | New password to the cardholder&#39;s user account on the Marqeta platform. | 
**UserToken** | **string** | Unique identifier of the cardholder. | 

## Methods

### NewResetUserPasswordModel

`func NewResetUserPasswordModel(newPassword string, userToken string, ) *ResetUserPasswordModel`

NewResetUserPasswordModel instantiates a new ResetUserPasswordModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetUserPasswordModelWithDefaults

`func NewResetUserPasswordModelWithDefaults() *ResetUserPasswordModel`

NewResetUserPasswordModelWithDefaults instantiates a new ResetUserPasswordModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *ResetUserPasswordModel) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetUserPasswordModel) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetUserPasswordModel) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetUserToken

`func (o *ResetUserPasswordModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *ResetUserPasswordModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *ResetUserPasswordModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


