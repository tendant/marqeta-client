# VelocityControlBalanceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of velocity control resources retrieved.  This field is returned if there are resources in your returned array. | [optional] 
**Data** | Pointer to [**[]VelocityControlBalanceResponse**](VelocityControlBalanceResponse.md) | Array of velocity control objects that include available balances.  Objects are returned as appropriate to your query. | [optional] 
**EndIndex** | Pointer to **int32** | Sort order index of the last resource in the returned array.  This field is returned if there are resources in your returned array. | [optional] 
**IsMore** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. A value of &#x60;false&#x60; indicates that no more unreturned resources exist.  This field is returned if there are resources in your returned array. | [optional] [default to false]
**StartIndex** | Pointer to **int32** | Sort order index of the last resource in the returned array.  This field is returned if there are resources in your returned array. | [optional] 

## Methods

### NewVelocityControlBalanceListResponse

`func NewVelocityControlBalanceListResponse() *VelocityControlBalanceListResponse`

NewVelocityControlBalanceListResponse instantiates a new VelocityControlBalanceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityControlBalanceListResponseWithDefaults

`func NewVelocityControlBalanceListResponseWithDefaults() *VelocityControlBalanceListResponse`

NewVelocityControlBalanceListResponseWithDefaults instantiates a new VelocityControlBalanceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VelocityControlBalanceListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VelocityControlBalanceListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VelocityControlBalanceListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VelocityControlBalanceListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *VelocityControlBalanceListResponse) GetData() []VelocityControlBalanceResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VelocityControlBalanceListResponse) GetDataOk() (*[]VelocityControlBalanceResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VelocityControlBalanceListResponse) SetData(v []VelocityControlBalanceResponse)`

SetData sets Data field to given value.

### HasData

`func (o *VelocityControlBalanceListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *VelocityControlBalanceListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *VelocityControlBalanceListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *VelocityControlBalanceListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *VelocityControlBalanceListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *VelocityControlBalanceListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *VelocityControlBalanceListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *VelocityControlBalanceListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *VelocityControlBalanceListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *VelocityControlBalanceListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *VelocityControlBalanceListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *VelocityControlBalanceListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *VelocityControlBalanceListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


