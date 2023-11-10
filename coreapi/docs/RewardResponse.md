# RewardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | Unique identifier of the account on which the reward exists. | [optional] 
**Amount** | **float32** | Amount of the reward. | 
**AppliedToAmount** | Pointer to **float32** | The total amount to which a percentage reward method is applied (for example, if a 3% reward is applied to 100, then &#x60;100&#x60; is the &#x60;applied_to_amount&#x60; value).  This field is not applicable for a flat fee method.  Returned for auto-cash back reward types only. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the reward was created on Marqeta&#39;s credit platform, in UTC. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | **string** | Description of the reward. | 
**Method** | Pointer to [**Method**](Method.md) |  | [optional] 
**Note** | Pointer to **string** | Additional information about the reward. | [optional] 
**Token** | **string** | Unique identifier of the reward.  If in the &#x60;detail_object&#x60;, unique identifier of the detail object. | 
**Type** | [**RewardType**](RewardType.md) |  | 
**UpdatedTime** | **time.Time** | Date and time when the reward was last updated on Marqeta&#39;s credit platform, in UTC. | 
**Value** | Pointer to **float32** | Value of the percentage or flat amount.  Returned for auto-cash back reward types only. | [optional] 

## Methods

### NewRewardResponse

`func NewRewardResponse(amount float32, createdTime time.Time, currencyCode CurrencyCode, description string, token string, type_ RewardType, updatedTime time.Time, ) *RewardResponse`

NewRewardResponse instantiates a new RewardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardResponseWithDefaults

`func NewRewardResponseWithDefaults() *RewardResponse`

NewRewardResponseWithDefaults instantiates a new RewardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *RewardResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *RewardResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *RewardResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *RewardResponse) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetAmount

`func (o *RewardResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAppliedToAmount

`func (o *RewardResponse) GetAppliedToAmount() float32`

GetAppliedToAmount returns the AppliedToAmount field if non-nil, zero value otherwise.

### GetAppliedToAmountOk

`func (o *RewardResponse) GetAppliedToAmountOk() (*float32, bool)`

GetAppliedToAmountOk returns a tuple with the AppliedToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedToAmount

`func (o *RewardResponse) SetAppliedToAmount(v float32)`

SetAppliedToAmount sets AppliedToAmount field to given value.

### HasAppliedToAmount

`func (o *RewardResponse) HasAppliedToAmount() bool`

HasAppliedToAmount returns a boolean if a field has been set.

### GetCreatedTime

`func (o *RewardResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RewardResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RewardResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *RewardResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RewardResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RewardResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *RewardResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMethod

`func (o *RewardResponse) GetMethod() Method`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RewardResponse) GetMethodOk() (*Method, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RewardResponse) SetMethod(v Method)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RewardResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNote

`func (o *RewardResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RewardResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetToken

`func (o *RewardResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RewardResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RewardResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *RewardResponse) GetType() RewardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RewardResponse) GetTypeOk() (*RewardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RewardResponse) SetType(v RewardType)`

SetType sets Type field to given value.


### GetUpdatedTime

`func (o *RewardResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *RewardResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *RewardResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetValue

`func (o *RewardResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RewardResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RewardResponse) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RewardResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


