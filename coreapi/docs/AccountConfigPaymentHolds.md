# AccountConfigPaymentHolds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchHoldDays** | Pointer to **int32** | Number of days to hold an ACH payment. | [optional] 
**CheckHoldDays** | Pointer to **int32** | Number of days to hold a check payment. | [optional] 

## Methods

### NewAccountConfigPaymentHolds

`func NewAccountConfigPaymentHolds() *AccountConfigPaymentHolds`

NewAccountConfigPaymentHolds instantiates a new AccountConfigPaymentHolds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigPaymentHoldsWithDefaults

`func NewAccountConfigPaymentHoldsWithDefaults() *AccountConfigPaymentHolds`

NewAccountConfigPaymentHoldsWithDefaults instantiates a new AccountConfigPaymentHolds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchHoldDays

`func (o *AccountConfigPaymentHolds) GetAchHoldDays() int32`

GetAchHoldDays returns the AchHoldDays field if non-nil, zero value otherwise.

### GetAchHoldDaysOk

`func (o *AccountConfigPaymentHolds) GetAchHoldDaysOk() (*int32, bool)`

GetAchHoldDaysOk returns a tuple with the AchHoldDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchHoldDays

`func (o *AccountConfigPaymentHolds) SetAchHoldDays(v int32)`

SetAchHoldDays sets AchHoldDays field to given value.

### HasAchHoldDays

`func (o *AccountConfigPaymentHolds) HasAchHoldDays() bool`

HasAchHoldDays returns a boolean if a field has been set.

### GetCheckHoldDays

`func (o *AccountConfigPaymentHolds) GetCheckHoldDays() int32`

GetCheckHoldDays returns the CheckHoldDays field if non-nil, zero value otherwise.

### GetCheckHoldDaysOk

`func (o *AccountConfigPaymentHolds) GetCheckHoldDaysOk() (*int32, bool)`

GetCheckHoldDaysOk returns a tuple with the CheckHoldDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHoldDays

`func (o *AccountConfigPaymentHolds) SetCheckHoldDays(v int32)`

SetCheckHoldDays sets CheckHoldDays field to given value.

### HasCheckHoldDays

`func (o *AccountConfigPaymentHolds) HasCheckHoldDays() bool`

HasCheckHoldDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


