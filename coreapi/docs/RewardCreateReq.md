# RewardCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the reward. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | **string** | Description of the reward. | 
**Note** | Pointer to **string** | Additional information about the reward. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the reward. | [optional] 

## Methods

### NewRewardCreateReq

`func NewRewardCreateReq(amount float32, currencyCode CurrencyCode, description string, ) *RewardCreateReq`

NewRewardCreateReq instantiates a new RewardCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardCreateReqWithDefaults

`func NewRewardCreateReqWithDefaults() *RewardCreateReq`

NewRewardCreateReqWithDefaults instantiates a new RewardCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RewardCreateReq) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardCreateReq) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardCreateReq) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrencyCode

`func (o *RewardCreateReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RewardCreateReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RewardCreateReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *RewardCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNote

`func (o *RewardCreateReq) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardCreateReq) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardCreateReq) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RewardCreateReq) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetToken

`func (o *RewardCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RewardCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RewardCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RewardCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


