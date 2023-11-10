# RedemptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount to redeem. | 
**CreatedTime** | **time.Time** | Date and time when the reward redemption was created on the Marqeta platform, in UTC. | 
**Destination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional] 
**ExternalSettlementDateTime** | Pointer to **time.Time** | Date and time when the reward redemption was settled on your external platform.  This field is returned if you handled the reward redemption outside of Marqeta&#39;s credit platform. | [optional] 
**Note** | **string** |  | 
**ReceivingAccountToken** | Pointer to **string** | Unique identifier of the external account receiving the reward redemption. This token is equivalent to the &lt;&lt;/core-api/payment-sources, payment source&gt;&gt; token. | [optional] 
**RelatedRewardEntries** | Pointer to [**[]RewardProgramsEntriesResponse**](RewardProgramsEntriesResponse.md) | Contains one or more reward entries related to the redemption. | [optional] 
**RewardProgramToken** | **string** | Unique identifier of the reward program for which to redeem rewards. | 
**SorRewardToken** | Pointer to **string** | Unique identifier of the system of reward (SOR) reward that was created to represent the reward redemption as a &#x60;STATEMENT_CREDIT&#x60; on a credit account. The SOR entry is a positive amount that is added to the account balance. | [optional] 
**Status** | [**RedemptionStatus**](RedemptionStatus.md) |  | 
**Token** | **string** | Unique identifier of the reward redemption. | 
**Type** | [**RedemptionType**](RedemptionType.md) |  | 
**UpdatedTime** | **time.Time** | Date and time when the reward redemption was last updated on the Marqeta platform, in UTC. | 

## Methods

### NewRedemptionsResponse

`func NewRedemptionsResponse(amount float32, createdTime time.Time, note string, rewardProgramToken string, status RedemptionStatus, token string, type_ RedemptionType, updatedTime time.Time, ) *RedemptionsResponse`

NewRedemptionsResponse instantiates a new RedemptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedemptionsResponseWithDefaults

`func NewRedemptionsResponseWithDefaults() *RedemptionsResponse`

NewRedemptionsResponseWithDefaults instantiates a new RedemptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RedemptionsResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RedemptionsResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RedemptionsResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *RedemptionsResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RedemptionsResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RedemptionsResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDestination

`func (o *RedemptionsResponse) GetDestination() DestinationType`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RedemptionsResponse) GetDestinationOk() (*DestinationType, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RedemptionsResponse) SetDestination(v DestinationType)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *RedemptionsResponse) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetExternalSettlementDateTime

`func (o *RedemptionsResponse) GetExternalSettlementDateTime() time.Time`

GetExternalSettlementDateTime returns the ExternalSettlementDateTime field if non-nil, zero value otherwise.

### GetExternalSettlementDateTimeOk

`func (o *RedemptionsResponse) GetExternalSettlementDateTimeOk() (*time.Time, bool)`

GetExternalSettlementDateTimeOk returns a tuple with the ExternalSettlementDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSettlementDateTime

`func (o *RedemptionsResponse) SetExternalSettlementDateTime(v time.Time)`

SetExternalSettlementDateTime sets ExternalSettlementDateTime field to given value.

### HasExternalSettlementDateTime

`func (o *RedemptionsResponse) HasExternalSettlementDateTime() bool`

HasExternalSettlementDateTime returns a boolean if a field has been set.

### GetNote

`func (o *RedemptionsResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RedemptionsResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RedemptionsResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetReceivingAccountToken

`func (o *RedemptionsResponse) GetReceivingAccountToken() string`

GetReceivingAccountToken returns the ReceivingAccountToken field if non-nil, zero value otherwise.

### GetReceivingAccountTokenOk

`func (o *RedemptionsResponse) GetReceivingAccountTokenOk() (*string, bool)`

GetReceivingAccountTokenOk returns a tuple with the ReceivingAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountToken

`func (o *RedemptionsResponse) SetReceivingAccountToken(v string)`

SetReceivingAccountToken sets ReceivingAccountToken field to given value.

### HasReceivingAccountToken

`func (o *RedemptionsResponse) HasReceivingAccountToken() bool`

HasReceivingAccountToken returns a boolean if a field has been set.

### GetRelatedRewardEntries

`func (o *RedemptionsResponse) GetRelatedRewardEntries() []RewardProgramsEntriesResponse`

GetRelatedRewardEntries returns the RelatedRewardEntries field if non-nil, zero value otherwise.

### GetRelatedRewardEntriesOk

`func (o *RedemptionsResponse) GetRelatedRewardEntriesOk() (*[]RewardProgramsEntriesResponse, bool)`

GetRelatedRewardEntriesOk returns a tuple with the RelatedRewardEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedRewardEntries

`func (o *RedemptionsResponse) SetRelatedRewardEntries(v []RewardProgramsEntriesResponse)`

SetRelatedRewardEntries sets RelatedRewardEntries field to given value.

### HasRelatedRewardEntries

`func (o *RedemptionsResponse) HasRelatedRewardEntries() bool`

HasRelatedRewardEntries returns a boolean if a field has been set.

### GetRewardProgramToken

`func (o *RedemptionsResponse) GetRewardProgramToken() string`

GetRewardProgramToken returns the RewardProgramToken field if non-nil, zero value otherwise.

### GetRewardProgramTokenOk

`func (o *RedemptionsResponse) GetRewardProgramTokenOk() (*string, bool)`

GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardProgramToken

`func (o *RedemptionsResponse) SetRewardProgramToken(v string)`

SetRewardProgramToken sets RewardProgramToken field to given value.


### GetSorRewardToken

`func (o *RedemptionsResponse) GetSorRewardToken() string`

GetSorRewardToken returns the SorRewardToken field if non-nil, zero value otherwise.

### GetSorRewardTokenOk

`func (o *RedemptionsResponse) GetSorRewardTokenOk() (*string, bool)`

GetSorRewardTokenOk returns a tuple with the SorRewardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorRewardToken

`func (o *RedemptionsResponse) SetSorRewardToken(v string)`

SetSorRewardToken sets SorRewardToken field to given value.

### HasSorRewardToken

`func (o *RedemptionsResponse) HasSorRewardToken() bool`

HasSorRewardToken returns a boolean if a field has been set.

### GetStatus

`func (o *RedemptionsResponse) GetStatus() RedemptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RedemptionsResponse) GetStatusOk() (*RedemptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RedemptionsResponse) SetStatus(v RedemptionStatus)`

SetStatus sets Status field to given value.


### GetToken

`func (o *RedemptionsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RedemptionsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RedemptionsResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *RedemptionsResponse) GetType() RedemptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RedemptionsResponse) GetTypeOk() (*RedemptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RedemptionsResponse) SetType(v RedemptionType)`

SetType sets Type field to given value.


### GetUpdatedTime

`func (o *RedemptionsResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *RedemptionsResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *RedemptionsResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


