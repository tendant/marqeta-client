# RewardProgramsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the associated credit account. | 
**BundleToken** | **string** | Unique identifier of the associated bundle that contains the reward policy on which the reward program is based. | 
**CalculationType** | [**CalculationType**](CalculationType.md) |  | 
**CreatedTime** | **time.Time** | Date and time when the reward program was created on the Marqeta platform, in UTC. | 
**IsActive** | **bool** | A value of &#x60;true&#x60; indicates that the reward program is active. | 
**Note** | **string** | A note that provides information on the reward program. | 
**Token** | **string** | Unique identifier of the reward program. | 
**UpdatedTime** | **time.Time** | Date and time when the reward program was last updated on the Marqeta platform, in UTC. | 

## Methods

### NewRewardProgramsResponse

`func NewRewardProgramsResponse(accountToken string, bundleToken string, calculationType CalculationType, createdTime time.Time, isActive bool, note string, token string, updatedTime time.Time, ) *RewardProgramsResponse`

NewRewardProgramsResponse instantiates a new RewardProgramsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardProgramsResponseWithDefaults

`func NewRewardProgramsResponseWithDefaults() *RewardProgramsResponse`

NewRewardProgramsResponseWithDefaults instantiates a new RewardProgramsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *RewardProgramsResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *RewardProgramsResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *RewardProgramsResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetBundleToken

`func (o *RewardProgramsResponse) GetBundleToken() string`

GetBundleToken returns the BundleToken field if non-nil, zero value otherwise.

### GetBundleTokenOk

`func (o *RewardProgramsResponse) GetBundleTokenOk() (*string, bool)`

GetBundleTokenOk returns a tuple with the BundleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleToken

`func (o *RewardProgramsResponse) SetBundleToken(v string)`

SetBundleToken sets BundleToken field to given value.


### GetCalculationType

`func (o *RewardProgramsResponse) GetCalculationType() CalculationType`

GetCalculationType returns the CalculationType field if non-nil, zero value otherwise.

### GetCalculationTypeOk

`func (o *RewardProgramsResponse) GetCalculationTypeOk() (*CalculationType, bool)`

GetCalculationTypeOk returns a tuple with the CalculationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationType

`func (o *RewardProgramsResponse) SetCalculationType(v CalculationType)`

SetCalculationType sets CalculationType field to given value.


### GetCreatedTime

`func (o *RewardProgramsResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RewardProgramsResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RewardProgramsResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetIsActive

`func (o *RewardProgramsResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RewardProgramsResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RewardProgramsResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetNote

`func (o *RewardProgramsResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardProgramsResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardProgramsResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetToken

`func (o *RewardProgramsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RewardProgramsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RewardProgramsResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *RewardProgramsResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *RewardProgramsResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *RewardProgramsResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


