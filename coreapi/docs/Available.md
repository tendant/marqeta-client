# Available

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Total amount of spend remaining in the velocity control. | 
**DaysRemaining** | Pointer to **int32** | Number of days remaining in the velocity control time window. | [optional] 
**Uses** | **int32** | Number of uses remaining in the velocity control. | [default to 0]

## Methods

### NewAvailable

`func NewAvailable(amount float32, uses int32, ) *Available`

NewAvailable instantiates a new Available object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableWithDefaults

`func NewAvailableWithDefaults() *Available`

NewAvailableWithDefaults instantiates a new Available object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Available) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Available) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Available) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDaysRemaining

`func (o *Available) GetDaysRemaining() int32`

GetDaysRemaining returns the DaysRemaining field if non-nil, zero value otherwise.

### GetDaysRemainingOk

`func (o *Available) GetDaysRemainingOk() (*int32, bool)`

GetDaysRemainingOk returns a tuple with the DaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRemaining

`func (o *Available) SetDaysRemaining(v int32)`

SetDaysRemaining sets DaysRemaining field to given value.

### HasDaysRemaining

`func (o *Available) HasDaysRemaining() bool`

HasDaysRemaining returns a boolean if a field has been set.

### GetUses

`func (o *Available) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *Available) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *Available) SetUses(v int32)`

SetUses sets Uses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


