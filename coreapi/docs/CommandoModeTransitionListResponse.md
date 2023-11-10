# CommandoModeTransitionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of Commando Mode control set transition objects to retrieve. | [optional] 
**Data** | Pointer to [**[]CommandoModeTransitionResponse**](CommandoModeTransitionResponse.md) | Array of Commando Mode control set transition objects. | [optional] 
**EndIndex** | Pointer to **int32** | Sort order index of the last resource in the returned array. | [optional] 
**IsMore** | Pointer to **bool** | Value of &#x60;true&#x60; indicates that more unreturned resources exist. | [optional] [default to false]
**StartIndex** | Pointer to **int32** | Sort order index of the first resource in the returned array. | [optional] 

## Methods

### NewCommandoModeTransitionListResponse

`func NewCommandoModeTransitionListResponse() *CommandoModeTransitionListResponse`

NewCommandoModeTransitionListResponse instantiates a new CommandoModeTransitionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandoModeTransitionListResponseWithDefaults

`func NewCommandoModeTransitionListResponseWithDefaults() *CommandoModeTransitionListResponse`

NewCommandoModeTransitionListResponseWithDefaults instantiates a new CommandoModeTransitionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CommandoModeTransitionListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CommandoModeTransitionListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CommandoModeTransitionListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CommandoModeTransitionListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *CommandoModeTransitionListResponse) GetData() []CommandoModeTransitionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CommandoModeTransitionListResponse) GetDataOk() (*[]CommandoModeTransitionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CommandoModeTransitionListResponse) SetData(v []CommandoModeTransitionResponse)`

SetData sets Data field to given value.

### HasData

`func (o *CommandoModeTransitionListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *CommandoModeTransitionListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *CommandoModeTransitionListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *CommandoModeTransitionListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *CommandoModeTransitionListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *CommandoModeTransitionListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *CommandoModeTransitionListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *CommandoModeTransitionListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *CommandoModeTransitionListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *CommandoModeTransitionListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *CommandoModeTransitionListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *CommandoModeTransitionListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *CommandoModeTransitionListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


