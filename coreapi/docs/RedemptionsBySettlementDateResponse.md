# RedemptionsBySettlementDateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | token of account the redemption is for. | 
**Amount** | **float32** |  | 
**CompletionTime** | Pointer to **time.Time** | yyyy-MM-ddThh:mm:ssZ | [optional] 
**CreatedTime** | **time.Time** | yyyy-MM-ddThh:mm:ssZ | 
**Destination** | [**DestinationType**](DestinationType.md) |  | 
**Note** | **string** |  | 
**RedemptionToken** | **string** | Identifier of the redemption. | 
**RewardProgramToken** | **string** | Token of reward program the redemption is for. | 
**Status** | [**RedemptionStatus**](RedemptionStatus.md) |  | 
**Type** | [**RedemptionType**](RedemptionType.md) |  | 
**UpdatedTime** | **time.Time** | yyyy-MM-ddThh:mm:ssZ | 

## Methods

### NewRedemptionsBySettlementDateResponse

`func NewRedemptionsBySettlementDateResponse(accountToken string, amount float32, createdTime time.Time, destination DestinationType, note string, redemptionToken string, rewardProgramToken string, status RedemptionStatus, type_ RedemptionType, updatedTime time.Time, ) *RedemptionsBySettlementDateResponse`

NewRedemptionsBySettlementDateResponse instantiates a new RedemptionsBySettlementDateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedemptionsBySettlementDateResponseWithDefaults

`func NewRedemptionsBySettlementDateResponseWithDefaults() *RedemptionsBySettlementDateResponse`

NewRedemptionsBySettlementDateResponseWithDefaults instantiates a new RedemptionsBySettlementDateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *RedemptionsBySettlementDateResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *RedemptionsBySettlementDateResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *RedemptionsBySettlementDateResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAmount

`func (o *RedemptionsBySettlementDateResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RedemptionsBySettlementDateResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RedemptionsBySettlementDateResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCompletionTime

`func (o *RedemptionsBySettlementDateResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *RedemptionsBySettlementDateResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *RedemptionsBySettlementDateResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *RedemptionsBySettlementDateResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetCreatedTime

`func (o *RedemptionsBySettlementDateResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RedemptionsBySettlementDateResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RedemptionsBySettlementDateResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDestination

`func (o *RedemptionsBySettlementDateResponse) GetDestination() DestinationType`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RedemptionsBySettlementDateResponse) GetDestinationOk() (*DestinationType, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RedemptionsBySettlementDateResponse) SetDestination(v DestinationType)`

SetDestination sets Destination field to given value.


### GetNote

`func (o *RedemptionsBySettlementDateResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RedemptionsBySettlementDateResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RedemptionsBySettlementDateResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetRedemptionToken

`func (o *RedemptionsBySettlementDateResponse) GetRedemptionToken() string`

GetRedemptionToken returns the RedemptionToken field if non-nil, zero value otherwise.

### GetRedemptionTokenOk

`func (o *RedemptionsBySettlementDateResponse) GetRedemptionTokenOk() (*string, bool)`

GetRedemptionTokenOk returns a tuple with the RedemptionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionToken

`func (o *RedemptionsBySettlementDateResponse) SetRedemptionToken(v string)`

SetRedemptionToken sets RedemptionToken field to given value.


### GetRewardProgramToken

`func (o *RedemptionsBySettlementDateResponse) GetRewardProgramToken() string`

GetRewardProgramToken returns the RewardProgramToken field if non-nil, zero value otherwise.

### GetRewardProgramTokenOk

`func (o *RedemptionsBySettlementDateResponse) GetRewardProgramTokenOk() (*string, bool)`

GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardProgramToken

`func (o *RedemptionsBySettlementDateResponse) SetRewardProgramToken(v string)`

SetRewardProgramToken sets RewardProgramToken field to given value.


### GetStatus

`func (o *RedemptionsBySettlementDateResponse) GetStatus() RedemptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RedemptionsBySettlementDateResponse) GetStatusOk() (*RedemptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RedemptionsBySettlementDateResponse) SetStatus(v RedemptionStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *RedemptionsBySettlementDateResponse) GetType() RedemptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RedemptionsBySettlementDateResponse) GetTypeOk() (*RedemptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RedemptionsBySettlementDateResponse) SetType(v RedemptionType)`

SetType sets Type field to given value.


### GetUpdatedTime

`func (o *RedemptionsBySettlementDateResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *RedemptionsBySettlementDateResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *RedemptionsBySettlementDateResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


