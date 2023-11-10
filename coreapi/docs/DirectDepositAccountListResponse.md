# DirectDepositAccountListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]DirectDepositAccountResponse**](DirectDepositAccountResponse.md) |  | [optional] 
**EndIndex** | Pointer to **int32** |  | [optional] 
**IsMore** | Pointer to **bool** |  | [optional] [default to false]
**StartIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewDirectDepositAccountListResponse

`func NewDirectDepositAccountListResponse() *DirectDepositAccountListResponse`

NewDirectDepositAccountListResponse instantiates a new DirectDepositAccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDepositAccountListResponseWithDefaults

`func NewDirectDepositAccountListResponseWithDefaults() *DirectDepositAccountListResponse`

NewDirectDepositAccountListResponseWithDefaults instantiates a new DirectDepositAccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DirectDepositAccountListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DirectDepositAccountListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DirectDepositAccountListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DirectDepositAccountListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *DirectDepositAccountListResponse) GetData() []DirectDepositAccountResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DirectDepositAccountListResponse) GetDataOk() (*[]DirectDepositAccountResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DirectDepositAccountListResponse) SetData(v []DirectDepositAccountResponse)`

SetData sets Data field to given value.

### HasData

`func (o *DirectDepositAccountListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *DirectDepositAccountListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *DirectDepositAccountListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *DirectDepositAccountListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *DirectDepositAccountListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *DirectDepositAccountListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *DirectDepositAccountListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *DirectDepositAccountListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *DirectDepositAccountListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *DirectDepositAccountListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *DirectDepositAccountListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *DirectDepositAccountListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *DirectDepositAccountListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


