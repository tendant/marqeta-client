# StatementReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | Date and time when the statement reward was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**CurrentBillingCycleReward** | Pointer to **float32** | Amount of rewards received in the current billing cycle. | [optional] 
**PreviousBillingCycleReward** | Pointer to **float32** | Amount of rewards received in the previous billing cycle. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the rewards for a specific statement. | [optional] 

## Methods

### NewStatementReward

`func NewStatementReward() *StatementReward`

NewStatementReward instantiates a new StatementReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementRewardWithDefaults

`func NewStatementRewardWithDefaults() *StatementReward`

NewStatementRewardWithDefaults instantiates a new StatementReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *StatementReward) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StatementReward) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StatementReward) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StatementReward) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrentBillingCycleReward

`func (o *StatementReward) GetCurrentBillingCycleReward() float32`

GetCurrentBillingCycleReward returns the CurrentBillingCycleReward field if non-nil, zero value otherwise.

### GetCurrentBillingCycleRewardOk

`func (o *StatementReward) GetCurrentBillingCycleRewardOk() (*float32, bool)`

GetCurrentBillingCycleRewardOk returns a tuple with the CurrentBillingCycleReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBillingCycleReward

`func (o *StatementReward) SetCurrentBillingCycleReward(v float32)`

SetCurrentBillingCycleReward sets CurrentBillingCycleReward field to given value.

### HasCurrentBillingCycleReward

`func (o *StatementReward) HasCurrentBillingCycleReward() bool`

HasCurrentBillingCycleReward returns a boolean if a field has been set.

### GetPreviousBillingCycleReward

`func (o *StatementReward) GetPreviousBillingCycleReward() float32`

GetPreviousBillingCycleReward returns the PreviousBillingCycleReward field if non-nil, zero value otherwise.

### GetPreviousBillingCycleRewardOk

`func (o *StatementReward) GetPreviousBillingCycleRewardOk() (*float32, bool)`

GetPreviousBillingCycleRewardOk returns a tuple with the PreviousBillingCycleReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBillingCycleReward

`func (o *StatementReward) SetPreviousBillingCycleReward(v float32)`

SetPreviousBillingCycleReward sets PreviousBillingCycleReward field to given value.

### HasPreviousBillingCycleReward

`func (o *StatementReward) HasPreviousBillingCycleReward() bool`

HasPreviousBillingCycleReward returns a boolean if a field has been set.

### GetToken

`func (o *StatementReward) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StatementReward) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StatementReward) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *StatementReward) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


