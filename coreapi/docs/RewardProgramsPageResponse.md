# RewardProgramsPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of resources returned. | 
**Data** | [**[]RewardProgramsResponse**](RewardProgramsResponse.md) | An array of reward program objects. | 
**EndIndex** | **int64** | Sort order index of the last resource in the returned array. | 
**IsMore** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. | [optional] 
**StartIndex** | **int64** | Sort order index of the first resource in the returned array. | 

## Methods

### NewRewardProgramsPageResponse

`func NewRewardProgramsPageResponse(count int32, data []RewardProgramsResponse, endIndex int64, startIndex int64, ) *RewardProgramsPageResponse`

NewRewardProgramsPageResponse instantiates a new RewardProgramsPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardProgramsPageResponseWithDefaults

`func NewRewardProgramsPageResponseWithDefaults() *RewardProgramsPageResponse`

NewRewardProgramsPageResponseWithDefaults instantiates a new RewardProgramsPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RewardProgramsPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RewardProgramsPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RewardProgramsPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *RewardProgramsPageResponse) GetData() []RewardProgramsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RewardProgramsPageResponse) GetDataOk() (*[]RewardProgramsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RewardProgramsPageResponse) SetData(v []RewardProgramsResponse)`

SetData sets Data field to given value.


### GetEndIndex

`func (o *RewardProgramsPageResponse) GetEndIndex() int64`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *RewardProgramsPageResponse) GetEndIndexOk() (*int64, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *RewardProgramsPageResponse) SetEndIndex(v int64)`

SetEndIndex sets EndIndex field to given value.


### GetIsMore

`func (o *RewardProgramsPageResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *RewardProgramsPageResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *RewardProgramsPageResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *RewardProgramsPageResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *RewardProgramsPageResponse) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RewardProgramsPageResponse) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RewardProgramsPageResponse) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


