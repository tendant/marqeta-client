# AccountReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**Method**](Method.md) |  | 
**Type** | [**RewardType**](RewardType.md) |  | 
**Value** | Pointer to **float32** | Value of the reward, either a flat reward amount or percentage value. | [optional] 

## Methods

### NewAccountReward

`func NewAccountReward(method Method, type_ RewardType, ) *AccountReward`

NewAccountReward instantiates a new AccountReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRewardWithDefaults

`func NewAccountRewardWithDefaults() *AccountReward`

NewAccountRewardWithDefaults instantiates a new AccountReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *AccountReward) GetMethod() Method`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccountReward) GetMethodOk() (*Method, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccountReward) SetMethod(v Method)`

SetMethod sets Method field to given value.


### GetType

`func (o *AccountReward) GetType() RewardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountReward) GetTypeOk() (*RewardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountReward) SetType(v RewardType)`

SetType sets Type field to given value.


### GetValue

`func (o *AccountReward) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AccountReward) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AccountReward) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *AccountReward) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


