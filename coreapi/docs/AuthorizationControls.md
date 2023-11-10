# AuthorizationControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldExpirationDays** | Pointer to **int32** | Specifies the number of days after which an authorization associated with this group expires. | [optional] [default to 7]
**HoldIncrease** | Pointer to [**HoldIncrease**](HoldIncrease.md) |  | [optional] 

## Methods

### NewAuthorizationControls

`func NewAuthorizationControls() *AuthorizationControls`

NewAuthorizationControls instantiates a new AuthorizationControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationControlsWithDefaults

`func NewAuthorizationControlsWithDefaults() *AuthorizationControls`

NewAuthorizationControlsWithDefaults instantiates a new AuthorizationControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldExpirationDays

`func (o *AuthorizationControls) GetHoldExpirationDays() int32`

GetHoldExpirationDays returns the HoldExpirationDays field if non-nil, zero value otherwise.

### GetHoldExpirationDaysOk

`func (o *AuthorizationControls) GetHoldExpirationDaysOk() (*int32, bool)`

GetHoldExpirationDaysOk returns a tuple with the HoldExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldExpirationDays

`func (o *AuthorizationControls) SetHoldExpirationDays(v int32)`

SetHoldExpirationDays sets HoldExpirationDays field to given value.

### HasHoldExpirationDays

`func (o *AuthorizationControls) HasHoldExpirationDays() bool`

HasHoldExpirationDays returns a boolean if a field has been set.

### GetHoldIncrease

`func (o *AuthorizationControls) GetHoldIncrease() HoldIncrease`

GetHoldIncrease returns the HoldIncrease field if non-nil, zero value otherwise.

### GetHoldIncreaseOk

`func (o *AuthorizationControls) GetHoldIncreaseOk() (*HoldIncrease, bool)`

GetHoldIncreaseOk returns a tuple with the HoldIncrease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldIncrease

`func (o *AuthorizationControls) SetHoldIncrease(v HoldIncrease)`

SetHoldIncrease sets HoldIncrease field to given value.

### HasHoldIncrease

`func (o *AuthorizationControls) HasHoldIncrease() bool`

HasHoldIncrease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


