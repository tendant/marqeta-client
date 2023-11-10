# RealTimeFeeGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of resources to retrieve.  This field is returned if there are resources in your returned array. | [optional] 
**Data** | Pointer to [**[]RealTimeFeeGroup**](RealTimeFeeGroup.md) | Array of real time fee group objects.  Objects are returned as appropriate to your query. | [optional] 
**EndIndex** | Pointer to **int32** | Sort order index of the last resource in the returned array.  This field is returned if there are resources in your returned array. | [optional] 
**IsMore** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. A value of &#x60;false&#x60; indicates that no more unreturned resources exist.  This field is returned if there are resources in your returned array. | [optional] [default to false]
**StartIndex** | Pointer to **int32** | Sort order index of the first resource in the returned array.  This field is returned if there are resources in your returned array. | [optional] 

## Methods

### NewRealTimeFeeGroupListResponse

`func NewRealTimeFeeGroupListResponse() *RealTimeFeeGroupListResponse`

NewRealTimeFeeGroupListResponse instantiates a new RealTimeFeeGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeFeeGroupListResponseWithDefaults

`func NewRealTimeFeeGroupListResponseWithDefaults() *RealTimeFeeGroupListResponse`

NewRealTimeFeeGroupListResponseWithDefaults instantiates a new RealTimeFeeGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RealTimeFeeGroupListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RealTimeFeeGroupListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RealTimeFeeGroupListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RealTimeFeeGroupListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *RealTimeFeeGroupListResponse) GetData() []RealTimeFeeGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RealTimeFeeGroupListResponse) GetDataOk() (*[]RealTimeFeeGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RealTimeFeeGroupListResponse) SetData(v []RealTimeFeeGroup)`

SetData sets Data field to given value.

### HasData

`func (o *RealTimeFeeGroupListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *RealTimeFeeGroupListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *RealTimeFeeGroupListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *RealTimeFeeGroupListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *RealTimeFeeGroupListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *RealTimeFeeGroupListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *RealTimeFeeGroupListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *RealTimeFeeGroupListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *RealTimeFeeGroupListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *RealTimeFeeGroupListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RealTimeFeeGroupListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RealTimeFeeGroupListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *RealTimeFeeGroupListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


