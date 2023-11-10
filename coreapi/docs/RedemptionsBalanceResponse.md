# RedemptionsBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **time.Time** | The ending date (or date-time) of a date range from which to return the redemption balance, in UTC. | 
**RedemptionTotalAmount** | **float32** | Total amount of rewards redeemed within a specified date range. | 
**RewardProgramToken** | **string** | Unique identifier of the reward program for which to return the redemption balance. | 
**StartDate** | **time.Time** | The starting date (or date-time) of a date range from which to return the redemption balance, in UTC. | 

## Methods

### NewRedemptionsBalanceResponse

`func NewRedemptionsBalanceResponse(endDate time.Time, redemptionTotalAmount float32, rewardProgramToken string, startDate time.Time, ) *RedemptionsBalanceResponse`

NewRedemptionsBalanceResponse instantiates a new RedemptionsBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedemptionsBalanceResponseWithDefaults

`func NewRedemptionsBalanceResponseWithDefaults() *RedemptionsBalanceResponse`

NewRedemptionsBalanceResponseWithDefaults instantiates a new RedemptionsBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *RedemptionsBalanceResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RedemptionsBalanceResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RedemptionsBalanceResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetRedemptionTotalAmount

`func (o *RedemptionsBalanceResponse) GetRedemptionTotalAmount() float32`

GetRedemptionTotalAmount returns the RedemptionTotalAmount field if non-nil, zero value otherwise.

### GetRedemptionTotalAmountOk

`func (o *RedemptionsBalanceResponse) GetRedemptionTotalAmountOk() (*float32, bool)`

GetRedemptionTotalAmountOk returns a tuple with the RedemptionTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionTotalAmount

`func (o *RedemptionsBalanceResponse) SetRedemptionTotalAmount(v float32)`

SetRedemptionTotalAmount sets RedemptionTotalAmount field to given value.


### GetRewardProgramToken

`func (o *RedemptionsBalanceResponse) GetRewardProgramToken() string`

GetRewardProgramToken returns the RewardProgramToken field if non-nil, zero value otherwise.

### GetRewardProgramTokenOk

`func (o *RedemptionsBalanceResponse) GetRewardProgramTokenOk() (*string, bool)`

GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardProgramToken

`func (o *RedemptionsBalanceResponse) SetRewardProgramToken(v string)`

SetRewardProgramToken sets RewardProgramToken field to given value.


### GetStartDate

`func (o *RedemptionsBalanceResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RedemptionsBalanceResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RedemptionsBalanceResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


