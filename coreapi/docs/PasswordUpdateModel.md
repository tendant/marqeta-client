# PasswordUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | **string** | Current password to the cardholder&#39;s user account on the Marqeta platform. | 
**NewPassword** | **string** | New password to the cardholder&#39;s user account on the Marqeta platform.  * Must contain at least one numeral + * Must contain at least one lowercase letter + * Must contain at least one uppercase letter + * Must contain at least one of these symbols: &#x60;@ # $ % ! ^ &amp; * ( ) \\ _ + ~ &#x60; - &#x3D; [ ] { } , ; : &#39; \&quot; , . / &lt; &gt; ?&#x60; | 

## Methods

### NewPasswordUpdateModel

`func NewPasswordUpdateModel(currentPassword string, newPassword string, ) *PasswordUpdateModel`

NewPasswordUpdateModel instantiates a new PasswordUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordUpdateModelWithDefaults

`func NewPasswordUpdateModelWithDefaults() *PasswordUpdateModel`

NewPasswordUpdateModelWithDefaults instantiates a new PasswordUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *PasswordUpdateModel) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *PasswordUpdateModel) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *PasswordUpdateModel) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.


### GetNewPassword

`func (o *PasswordUpdateModel) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PasswordUpdateModel) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PasswordUpdateModel) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


