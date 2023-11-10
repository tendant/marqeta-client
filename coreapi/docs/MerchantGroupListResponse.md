# MerchantGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of resources retrieved.  This field is returned if there are resources in your returned array. | [optional] 
**Data** | Pointer to [**[]MerchantGroupResponse**](MerchantGroupResponse.md) | Array of merchant group resources.  Resources are returned as appropriate to your query. | [optional] 
**EndIndex** | Pointer to **int32** | The sort order index of the last resource in the returned array.  This field is returned if there are resources in your returned array. | [optional] 
**IsMore** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. A value of &#x60;false&#x60; indicates that no more unreturned resources exist.  This field is returned if there are resources in your returned array. | [optional] [default to false]
**StartIndex** | Pointer to **int32** | Sort order index of the last resource in the returned array.  This field is returned if there are resources in your returned array. | [optional] 

## Methods

### NewMerchantGroupListResponse

`func NewMerchantGroupListResponse() *MerchantGroupListResponse`

NewMerchantGroupListResponse instantiates a new MerchantGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantGroupListResponseWithDefaults

`func NewMerchantGroupListResponseWithDefaults() *MerchantGroupListResponse`

NewMerchantGroupListResponseWithDefaults instantiates a new MerchantGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MerchantGroupListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MerchantGroupListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MerchantGroupListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MerchantGroupListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *MerchantGroupListResponse) GetData() []MerchantGroupResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MerchantGroupListResponse) GetDataOk() (*[]MerchantGroupResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MerchantGroupListResponse) SetData(v []MerchantGroupResponse)`

SetData sets Data field to given value.

### HasData

`func (o *MerchantGroupListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *MerchantGroupListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *MerchantGroupListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *MerchantGroupListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *MerchantGroupListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *MerchantGroupListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *MerchantGroupListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *MerchantGroupListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *MerchantGroupListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *MerchantGroupListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *MerchantGroupListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *MerchantGroupListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *MerchantGroupListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


