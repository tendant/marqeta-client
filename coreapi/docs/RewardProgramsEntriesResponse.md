# RewardProgramsEntriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **time.Time** | Date and time when the reward entry was created on the Marqeta platform, in UTC. | 
**Mcc** | Pointer to **string** | Merchant category code (MCC) of the related journal entry. | [optional] 
**Mid** | Pointer to **string** | Merchant identifier (MID) of the related journal entry. | [optional] 
**Note** | **string** | A note providing information on the reward entry. | 
**RelatedJournalEntryToken** | **string** | Unique identifier of the related journal entry to which the reward rule was applied to trigger the reward entry. | 
**RelatedRedemptionToken** | Pointer to **string** | Unique identifier of the related redemption token that triggered the reward entry. | [optional] 
**RewardProgramToken** | **string** | Unique identifier of the reward program for which to return reward entries. | 
**RewardRulesConfigToken** | **string** | Unique identifier of the reward rules config used to determine the value of the reward entry. | 
**Status** | [**RewardEntryStatus**](RewardEntryStatus.md) |  | 
**Token** | **string** | Unique identifier of the reward entry. | 
**TransactionAmount** | **float32** | The transaction amount to which the reward rule was applied. Used to determine the value of the reward entry. | 
**Value** | Pointer to **float32** | Value of the reward entry. | [optional] 

## Methods

### NewRewardProgramsEntriesResponse

`func NewRewardProgramsEntriesResponse(createdTime time.Time, note string, relatedJournalEntryToken string, rewardProgramToken string, rewardRulesConfigToken string, status RewardEntryStatus, token string, transactionAmount float32, ) *RewardProgramsEntriesResponse`

NewRewardProgramsEntriesResponse instantiates a new RewardProgramsEntriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardProgramsEntriesResponseWithDefaults

`func NewRewardProgramsEntriesResponseWithDefaults() *RewardProgramsEntriesResponse`

NewRewardProgramsEntriesResponseWithDefaults instantiates a new RewardProgramsEntriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *RewardProgramsEntriesResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RewardProgramsEntriesResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RewardProgramsEntriesResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetMcc

`func (o *RewardProgramsEntriesResponse) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *RewardProgramsEntriesResponse) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *RewardProgramsEntriesResponse) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *RewardProgramsEntriesResponse) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMid

`func (o *RewardProgramsEntriesResponse) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *RewardProgramsEntriesResponse) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *RewardProgramsEntriesResponse) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *RewardProgramsEntriesResponse) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetNote

`func (o *RewardProgramsEntriesResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardProgramsEntriesResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardProgramsEntriesResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetRelatedJournalEntryToken

`func (o *RewardProgramsEntriesResponse) GetRelatedJournalEntryToken() string`

GetRelatedJournalEntryToken returns the RelatedJournalEntryToken field if non-nil, zero value otherwise.

### GetRelatedJournalEntryTokenOk

`func (o *RewardProgramsEntriesResponse) GetRelatedJournalEntryTokenOk() (*string, bool)`

GetRelatedJournalEntryTokenOk returns a tuple with the RelatedJournalEntryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedJournalEntryToken

`func (o *RewardProgramsEntriesResponse) SetRelatedJournalEntryToken(v string)`

SetRelatedJournalEntryToken sets RelatedJournalEntryToken field to given value.


### GetRelatedRedemptionToken

`func (o *RewardProgramsEntriesResponse) GetRelatedRedemptionToken() string`

GetRelatedRedemptionToken returns the RelatedRedemptionToken field if non-nil, zero value otherwise.

### GetRelatedRedemptionTokenOk

`func (o *RewardProgramsEntriesResponse) GetRelatedRedemptionTokenOk() (*string, bool)`

GetRelatedRedemptionTokenOk returns a tuple with the RelatedRedemptionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedRedemptionToken

`func (o *RewardProgramsEntriesResponse) SetRelatedRedemptionToken(v string)`

SetRelatedRedemptionToken sets RelatedRedemptionToken field to given value.

### HasRelatedRedemptionToken

`func (o *RewardProgramsEntriesResponse) HasRelatedRedemptionToken() bool`

HasRelatedRedemptionToken returns a boolean if a field has been set.

### GetRewardProgramToken

`func (o *RewardProgramsEntriesResponse) GetRewardProgramToken() string`

GetRewardProgramToken returns the RewardProgramToken field if non-nil, zero value otherwise.

### GetRewardProgramTokenOk

`func (o *RewardProgramsEntriesResponse) GetRewardProgramTokenOk() (*string, bool)`

GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardProgramToken

`func (o *RewardProgramsEntriesResponse) SetRewardProgramToken(v string)`

SetRewardProgramToken sets RewardProgramToken field to given value.


### GetRewardRulesConfigToken

`func (o *RewardProgramsEntriesResponse) GetRewardRulesConfigToken() string`

GetRewardRulesConfigToken returns the RewardRulesConfigToken field if non-nil, zero value otherwise.

### GetRewardRulesConfigTokenOk

`func (o *RewardProgramsEntriesResponse) GetRewardRulesConfigTokenOk() (*string, bool)`

GetRewardRulesConfigTokenOk returns a tuple with the RewardRulesConfigToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardRulesConfigToken

`func (o *RewardProgramsEntriesResponse) SetRewardRulesConfigToken(v string)`

SetRewardRulesConfigToken sets RewardRulesConfigToken field to given value.


### GetStatus

`func (o *RewardProgramsEntriesResponse) GetStatus() RewardEntryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RewardProgramsEntriesResponse) GetStatusOk() (*RewardEntryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RewardProgramsEntriesResponse) SetStatus(v RewardEntryStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *RewardProgramsEntriesResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RewardProgramsEntriesResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RewardProgramsEntriesResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransactionAmount

`func (o *RewardProgramsEntriesResponse) GetTransactionAmount() float32`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *RewardProgramsEntriesResponse) GetTransactionAmountOk() (*float32, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *RewardProgramsEntriesResponse) SetTransactionAmount(v float32)`

SetTransactionAmount sets TransactionAmount field to given value.


### GetValue

`func (o *RewardProgramsEntriesResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RewardProgramsEntriesResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RewardProgramsEntriesResponse) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RewardProgramsEntriesResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


