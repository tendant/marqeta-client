# PutRewardProgramsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | **bool** | A value of &#x60;true&#x60; indicates that the reward program is active and rewards can be accrued for the associated account. | 
**Note** | **string** | A note explaining why the reward program is being activated or deactivated. | 

## Methods

### NewPutRewardProgramsRequest

`func NewPutRewardProgramsRequest(isActive bool, note string, ) *PutRewardProgramsRequest`

NewPutRewardProgramsRequest instantiates a new PutRewardProgramsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRewardProgramsRequestWithDefaults

`func NewPutRewardProgramsRequestWithDefaults() *PutRewardProgramsRequest`

NewPutRewardProgramsRequestWithDefaults instantiates a new PutRewardProgramsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *PutRewardProgramsRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PutRewardProgramsRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PutRewardProgramsRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetNote

`func (o *PutRewardProgramsRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PutRewardProgramsRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PutRewardProgramsRequest) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


