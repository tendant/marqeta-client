# RewardProgramsEntriesBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to **time.Time** | Date and time the reward entries balance was created on the Marqeta platform, in UTC. | [optional] 
**EndDate** | **time.Time** | The ending date (or date-time) of a date range from which to return accrued rewards, in UTC. Reward entries created on or before this date count toward the total reward balance. | 
**RewardProgramToken** | **string** | Unique identifier of the reward program for which to retrieve the reward entries balance. | 
**StartDate** | **time.Time** | The starting date (or date-time) of a date range from which to return accrued rewards, in UTC. Reward entries created on or after this date count toward the total reward balance. | 
**TotalRewardBalance** | **float32** | The total balance of rewards accrued within a date range. | 

## Methods

### NewRewardProgramsEntriesBalanceResponse

`func NewRewardProgramsEntriesBalanceResponse(endDate time.Time, rewardProgramToken string, startDate time.Time, totalRewardBalance float32, ) *RewardProgramsEntriesBalanceResponse`

NewRewardProgramsEntriesBalanceResponse instantiates a new RewardProgramsEntriesBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardProgramsEntriesBalanceResponseWithDefaults

`func NewRewardProgramsEntriesBalanceResponseWithDefaults() *RewardProgramsEntriesBalanceResponse`

NewRewardProgramsEntriesBalanceResponseWithDefaults instantiates a new RewardProgramsEntriesBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *RewardProgramsEntriesBalanceResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *RewardProgramsEntriesBalanceResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *RewardProgramsEntriesBalanceResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *RewardProgramsEntriesBalanceResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetEndDate

`func (o *RewardProgramsEntriesBalanceResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RewardProgramsEntriesBalanceResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RewardProgramsEntriesBalanceResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetRewardProgramToken

`func (o *RewardProgramsEntriesBalanceResponse) GetRewardProgramToken() string`

GetRewardProgramToken returns the RewardProgramToken field if non-nil, zero value otherwise.

### GetRewardProgramTokenOk

`func (o *RewardProgramsEntriesBalanceResponse) GetRewardProgramTokenOk() (*string, bool)`

GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardProgramToken

`func (o *RewardProgramsEntriesBalanceResponse) SetRewardProgramToken(v string)`

SetRewardProgramToken sets RewardProgramToken field to given value.


### GetStartDate

`func (o *RewardProgramsEntriesBalanceResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RewardProgramsEntriesBalanceResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RewardProgramsEntriesBalanceResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetTotalRewardBalance

`func (o *RewardProgramsEntriesBalanceResponse) GetTotalRewardBalance() float32`

GetTotalRewardBalance returns the TotalRewardBalance field if non-nil, zero value otherwise.

### GetTotalRewardBalanceOk

`func (o *RewardProgramsEntriesBalanceResponse) GetTotalRewardBalanceOk() (*float32, bool)`

GetTotalRewardBalanceOk returns a tuple with the TotalRewardBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRewardBalance

`func (o *RewardProgramsEntriesBalanceResponse) SetTotalRewardBalance(v float32)`

SetTotalRewardBalance sets TotalRewardBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


