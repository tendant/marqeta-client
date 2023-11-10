# PushToCardListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]PushToCardResponse**](PushToCardResponse.md) |  | [optional] 
**EndIndex** | Pointer to **int32** |  | [optional] 
**IsMore** | Pointer to **bool** |  | [optional] [default to false]
**StartIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewPushToCardListResponse

`func NewPushToCardListResponse() *PushToCardListResponse`

NewPushToCardListResponse instantiates a new PushToCardListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushToCardListResponseWithDefaults

`func NewPushToCardListResponseWithDefaults() *PushToCardListResponse`

NewPushToCardListResponseWithDefaults instantiates a new PushToCardListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PushToCardListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PushToCardListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PushToCardListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PushToCardListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *PushToCardListResponse) GetData() []PushToCardResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PushToCardListResponse) GetDataOk() (*[]PushToCardResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PushToCardListResponse) SetData(v []PushToCardResponse)`

SetData sets Data field to given value.

### HasData

`func (o *PushToCardListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *PushToCardListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *PushToCardListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *PushToCardListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *PushToCardListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *PushToCardListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *PushToCardListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *PushToCardListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *PushToCardListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *PushToCardListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *PushToCardListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *PushToCardListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *PushToCardListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


