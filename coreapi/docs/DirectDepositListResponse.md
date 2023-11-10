# DirectDepositListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]DepositDepositResponse**](DepositDepositResponse.md) |  | [optional] 
**EndIndex** | Pointer to **int32** |  | [optional] 
**IsMore** | Pointer to **bool** |  | [optional] [default to false]
**StartIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewDirectDepositListResponse

`func NewDirectDepositListResponse() *DirectDepositListResponse`

NewDirectDepositListResponse instantiates a new DirectDepositListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositListResponseWithDefaults

`func NewDirectDepositListResponseWithDefaults() *DirectDepositListResponse`

NewDirectDepositListResponseWithDefaults instantiates a new DirectDepositListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DirectDepositListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DirectDepositListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DirectDepositListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DirectDepositListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *DirectDepositListResponse) GetData() []DepositDepositResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DirectDepositListResponse) GetDataOk() (*[]DepositDepositResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DirectDepositListResponse) SetData(v []DepositDepositResponse)`

SetData sets Data field to given value.

### HasData

`func (o *DirectDepositListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *DirectDepositListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *DirectDepositListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *DirectDepositListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *DirectDepositListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *DirectDepositListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *DirectDepositListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *DirectDepositListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *DirectDepositListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *DirectDepositListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *DirectDepositListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *DirectDepositListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *DirectDepositListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


