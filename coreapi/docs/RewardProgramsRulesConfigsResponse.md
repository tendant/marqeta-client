# RewardProgramsRulesConfigsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualType** | [**AccrualType**](AccrualType.md) |  | 
**CreatedTime** | **time.Time** | Date and time when the reward rules config was created on the Marqeta platform, in UTC. | 
**GreaterThan** | Pointer to **float32** | Minimum amount that the balance for a billing cycle can be to apply the specified reward percentage. For example, if the &#x60;greater_than&#x60; value is &#x60;500&#x60;, the account holder earns _x_% of the account balance if they spend over $500 during a billing cycle. | [optional] 
**IsActive** | **bool** | A value of &#x60;true&#x60; indicates that the reward rules config is active. | 
**LessThan** | Pointer to **float32** | Maximum amount that the balance for a billing cycle can be to apply the specified reward percentage. For example, if the &#x60;less_than&#x60; value is &#x60;1500&#x60;, the account holder earns _x_% of the account balance if they spend under $1500 during a billing cycle. | [optional] 
**Percentage** | **int32** | The reward percentage applied when the balance for a billing cycle is within the range specified in the &#x60;less_than&#x60; and &#x60;greater_than&#x60; fields. For example, if the &#x60;percentage&#x60; is &#x60;1&#x60;, the account holder earns 1% of the account balance if they spend between the &#x60;less_than&#x60; and &#x60;greater_than&#x60; amounts during a billing cycle. | 
**RewardProgramToken** | **string** | Unique identifier of the reward program on which the rules config is applied. | 
**Token** | **string** | Unique identifier of the reward rules config. | 
**UpdatedTime** | **time.Time** | Date and time when the reward rules config was last updated on the Marqeta platform, in UTC. | 

## Methods

### NewRewardProgramsRulesConfigsResponse

`func NewRewardProgramsRulesConfigsResponse(accrualType AccrualType, createdTime time.Time, isActive bool, percentage int32, rewardProgramToken string, token string, updatedTime time.Time, ) *RewardProgramsRulesConfigsResponse`

NewRewardProgramsRulesConfigsResponse instantiates a new RewardProgramsRulesConfigsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardProgramsRulesConfigsResponseWithDefaults

`func NewRewardProgramsRulesConfigsResponseWithDefaults() *RewardProgramsRulesConfigsResponse`

NewRewardProgramsRulesConfigsResponseWithDefaults instantiates a new RewardProgramsRulesConfigsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualType

`func (o *RewardProgramsRulesConfigsResponse) GetAccrualType() AccrualType`

GetAccrualType returns the AccrualType field if non-nil, zero value otherwise.

### GetAccrualTypeOk

`func (o *RewardProgramsRulesConfigsResponse) GetAccrualTypeOk() (*AccrualType, bool)`

GetAccrualTypeOk returns a tuple with the AccrualType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualType

`func (o *RewardProgramsRulesConfigsResponse) SetAccrualType(v AccrualType)`

SetAccrualType sets AccrualType field to given value.


### GetCreatedTime

`func (o *RewardProgramsRulesConfigsResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RewardProgramsRulesConfigsResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RewardProgramsRulesConfigsResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetGreaterThan

`func (o *RewardProgramsRulesConfigsResponse) GetGreaterThan() float32`

GetGreaterThan returns the GreaterThan field if non-nil, zero value otherwise.

### GetGreaterThanOk

`func (o *RewardProgramsRulesConfigsResponse) GetGreaterThanOk() (*float32, bool)`

GetGreaterThanOk returns a tuple with the GreaterThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreaterThan

`func (o *RewardProgramsRulesConfigsResponse) SetGreaterThan(v float32)`

SetGreaterThan sets GreaterThan field to given value.

### HasGreaterThan

`func (o *RewardProgramsRulesConfigsResponse) HasGreaterThan() bool`

HasGreaterThan returns a boolean if a field has been set.

### GetIsActive

`func (o *RewardProgramsRulesConfigsResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RewardProgramsRulesConfigsResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RewardProgramsRulesConfigsResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLessThan

`func (o *RewardProgramsRulesConfigsResponse) GetLessThan() float32`

GetLessThan returns the LessThan field if non-nil, zero value otherwise.

### GetLessThanOk

`func (o *RewardProgramsRulesConfigsResponse) GetLessThanOk() (*float32, bool)`

GetLessThanOk returns a tuple with the LessThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessThan

`func (o *RewardProgramsRulesConfigsResponse) SetLessThan(v float32)`

SetLessThan sets LessThan field to given value.

### HasLessThan

`func (o *RewardProgramsRulesConfigsResponse) HasLessThan() bool`

HasLessThan returns a boolean if a field has been set.

### GetPercentage

`func (o *RewardProgramsRulesConfigsResponse) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *RewardProgramsRulesConfigsResponse) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *RewardProgramsRulesConfigsResponse) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetRewardProgramToken

`func (o *RewardProgramsRulesConfigsResponse) GetRewardProgramToken() string`

GetRewardProgramToken returns the RewardProgramToken field if non-nil, zero value otherwise.

### GetRewardProgramTokenOk

`func (o *RewardProgramsRulesConfigsResponse) GetRewardProgramTokenOk() (*string, bool)`

GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardProgramToken

`func (o *RewardProgramsRulesConfigsResponse) SetRewardProgramToken(v string)`

SetRewardProgramToken sets RewardProgramToken field to given value.


### GetToken

`func (o *RewardProgramsRulesConfigsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RewardProgramsRulesConfigsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RewardProgramsRulesConfigsResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *RewardProgramsRulesConfigsResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *RewardProgramsRulesConfigsResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *RewardProgramsRulesConfigsResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


