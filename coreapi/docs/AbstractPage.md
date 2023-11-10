# AbstractPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of resources returned. | 
**EndIndex** | **int64** | Sort order index of the last resource in the returned array. | 
**IsMore** | **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. | 
**StartIndex** | **int64** | Sort order index of the first resource in the returned array. | 

## Methods

### NewAbstractPage

`func NewAbstractPage(count int32, endIndex int64, isMore bool, startIndex int64, ) *AbstractPage`

NewAbstractPage instantiates a new AbstractPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractPageWithDefaults

`func NewAbstractPageWithDefaults() *AbstractPage`

NewAbstractPageWithDefaults instantiates a new AbstractPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AbstractPage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AbstractPage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AbstractPage) SetCount(v int32)`

SetCount sets Count field to given value.


### GetEndIndex

`func (o *AbstractPage) GetEndIndex() int64`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *AbstractPage) GetEndIndexOk() (*int64, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *AbstractPage) SetEndIndex(v int64)`

SetEndIndex sets EndIndex field to given value.


### GetIsMore

`func (o *AbstractPage) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *AbstractPage) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *AbstractPage) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.


### GetStartIndex

`func (o *AbstractPage) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *AbstractPage) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *AbstractPage) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


