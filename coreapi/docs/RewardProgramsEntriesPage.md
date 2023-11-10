# RewardProgramsEntriesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of resources returned. | 
**Data** | [**[]RewardProgramsEntriesResponse**](RewardProgramsEntriesResponse.md) | An array of reward entry objects. | 
**EndIndex** | **int64** | Sort order index of the last resource in the returned array. | 
**IsMore** | **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. | 
**StartIndex** | **int64** | Sort order index of the first resource in the returned array. | 

## Methods

### NewRewardProgramsEntriesPage

`func NewRewardProgramsEntriesPage(count int32, data []RewardProgramsEntriesResponse, endIndex int64, isMore bool, startIndex int64, ) *RewardProgramsEntriesPage`

NewRewardProgramsEntriesPage instantiates a new RewardProgramsEntriesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardProgramsEntriesPageWithDefaults

`func NewRewardProgramsEntriesPageWithDefaults() *RewardProgramsEntriesPage`

NewRewardProgramsEntriesPageWithDefaults instantiates a new RewardProgramsEntriesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RewardProgramsEntriesPage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RewardProgramsEntriesPage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RewardProgramsEntriesPage) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *RewardProgramsEntriesPage) GetData() []RewardProgramsEntriesResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RewardProgramsEntriesPage) GetDataOk() (*[]RewardProgramsEntriesResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RewardProgramsEntriesPage) SetData(v []RewardProgramsEntriesResponse)`

SetData sets Data field to given value.


### GetEndIndex

`func (o *RewardProgramsEntriesPage) GetEndIndex() int64`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *RewardProgramsEntriesPage) GetEndIndexOk() (*int64, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *RewardProgramsEntriesPage) SetEndIndex(v int64)`

SetEndIndex sets EndIndex field to given value.


### GetIsMore

`func (o *RewardProgramsEntriesPage) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *RewardProgramsEntriesPage) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *RewardProgramsEntriesPage) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.


### GetStartIndex

`func (o *RewardProgramsEntriesPage) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RewardProgramsEntriesPage) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RewardProgramsEntriesPage) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


