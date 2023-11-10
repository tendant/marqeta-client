# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailVerified** | Pointer to **bool** | Specifies whether the email address has been verified. | [optional] [default to false]
**EmailVerifiedTime** | Pointer to **time.Time** | Date and time when the email address was verified. | [optional] 
**LastPasswordUpdateChannel** | Pointer to **string** | Specifies the channel through which the password was last changed. | [optional] 
**LastPasswordUpdateTime** | Pointer to **time.Time** | Date and time when the password was last changed. | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailVerified

`func (o *Authentication) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *Authentication) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *Authentication) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *Authentication) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetEmailVerifiedTime

`func (o *Authentication) GetEmailVerifiedTime() time.Time`

GetEmailVerifiedTime returns the EmailVerifiedTime field if non-nil, zero value otherwise.

### GetEmailVerifiedTimeOk

`func (o *Authentication) GetEmailVerifiedTimeOk() (*time.Time, bool)`

GetEmailVerifiedTimeOk returns a tuple with the EmailVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerifiedTime

`func (o *Authentication) SetEmailVerifiedTime(v time.Time)`

SetEmailVerifiedTime sets EmailVerifiedTime field to given value.

### HasEmailVerifiedTime

`func (o *Authentication) HasEmailVerifiedTime() bool`

HasEmailVerifiedTime returns a boolean if a field has been set.

### GetLastPasswordUpdateChannel

`func (o *Authentication) GetLastPasswordUpdateChannel() string`

GetLastPasswordUpdateChannel returns the LastPasswordUpdateChannel field if non-nil, zero value otherwise.

### GetLastPasswordUpdateChannelOk

`func (o *Authentication) GetLastPasswordUpdateChannelOk() (*string, bool)`

GetLastPasswordUpdateChannelOk returns a tuple with the LastPasswordUpdateChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordUpdateChannel

`func (o *Authentication) SetLastPasswordUpdateChannel(v string)`

SetLastPasswordUpdateChannel sets LastPasswordUpdateChannel field to given value.

### HasLastPasswordUpdateChannel

`func (o *Authentication) HasLastPasswordUpdateChannel() bool`

HasLastPasswordUpdateChannel returns a boolean if a field has been set.

### GetLastPasswordUpdateTime

`func (o *Authentication) GetLastPasswordUpdateTime() time.Time`

GetLastPasswordUpdateTime returns the LastPasswordUpdateTime field if non-nil, zero value otherwise.

### GetLastPasswordUpdateTimeOk

`func (o *Authentication) GetLastPasswordUpdateTimeOk() (*time.Time, bool)`

GetLastPasswordUpdateTimeOk returns a tuple with the LastPasswordUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordUpdateTime

`func (o *Authentication) SetLastPasswordUpdateTime(v time.Time)`

SetLastPasswordUpdateTime sets LastPasswordUpdateTime field to given value.

### HasLastPasswordUpdateTime

`func (o *Authentication) HasLastPasswordUpdateTime() bool`

HasLastPasswordUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


