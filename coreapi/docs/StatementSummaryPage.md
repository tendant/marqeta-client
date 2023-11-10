# StatementSummaryPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of resources returned. | 
**Data** | [**[]StatementSummary**](StatementSummary.md) | One or more statement summaries. | 
**EndIndex** | **int32** | Sort order index of the last resource in the returned array. | 
**IsMore** | **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. | 
**StartIndex** | **int32** | Sort order index of the first resource in the returned array. | 

## Methods

### NewStatementSummaryPage

`func NewStatementSummaryPage(count int32, data []StatementSummary, endIndex int32, isMore bool, startIndex int32, ) *StatementSummaryPage`

NewStatementSummaryPage instantiates a new StatementSummaryPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementSummaryPageWithDefaults

`func NewStatementSummaryPageWithDefaults() *StatementSummaryPage`

NewStatementSummaryPageWithDefaults instantiates a new StatementSummaryPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *StatementSummaryPage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StatementSummaryPage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StatementSummaryPage) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *StatementSummaryPage) GetData() []StatementSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StatementSummaryPage) GetDataOk() (*[]StatementSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StatementSummaryPage) SetData(v []StatementSummary)`

SetData sets Data field to given value.


### GetEndIndex

`func (o *StatementSummaryPage) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *StatementSummaryPage) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *StatementSummaryPage) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.


### GetIsMore

`func (o *StatementSummaryPage) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *StatementSummaryPage) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *StatementSummaryPage) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.


### GetStartIndex

`func (o *StatementSummaryPage) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *StatementSummaryPage) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *StatementSummaryPage) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


