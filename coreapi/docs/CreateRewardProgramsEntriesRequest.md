# CreateRewardProgramsEntriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | Date and time when the reward entry was created on the Marqeta platform, in UTC. | [optional] 
**Note** | **string** | A note explaining why the reward entry is being created manually. | 
**Token** | Pointer to **string** | Unique identifier of the reward entry. | [optional] 
**Value** | **float32** | Value of the reward granted to the account. | 

## Methods

### NewCreateRewardProgramsEntriesRequest

`func NewCreateRewardProgramsEntriesRequest(note string, value float32, ) *CreateRewardProgramsEntriesRequest`

NewCreateRewardProgramsEntriesRequest instantiates a new CreateRewardProgramsEntriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRewardProgramsEntriesRequestWithDefaults

`func NewCreateRewardProgramsEntriesRequestWithDefaults() *CreateRewardProgramsEntriesRequest`

NewCreateRewardProgramsEntriesRequestWithDefaults instantiates a new CreateRewardProgramsEntriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *CreateRewardProgramsEntriesRequest) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CreateRewardProgramsEntriesRequest) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CreateRewardProgramsEntriesRequest) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CreateRewardProgramsEntriesRequest) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetNote

`func (o *CreateRewardProgramsEntriesRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateRewardProgramsEntriesRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateRewardProgramsEntriesRequest) SetNote(v string)`

SetNote sets Note field to given value.


### GetToken

`func (o *CreateRewardProgramsEntriesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateRewardProgramsEntriesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateRewardProgramsEntriesRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateRewardProgramsEntriesRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetValue

`func (o *CreateRewardProgramsEntriesRequest) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateRewardProgramsEntriesRequest) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateRewardProgramsEntriesRequest) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


