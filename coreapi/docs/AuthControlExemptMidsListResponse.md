# AuthControlExemptMidsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of resources retrieved.  This field is returned if there are resources in your returned array. | [optional] 
**Data** | Pointer to [**[]AuthControlExemptMidsResponse**](AuthControlExemptMidsResponse.md) | Array of objects in a returned resource.  Objects are returned as appropriate to your query. | [optional] 
**EndIndex** | Pointer to **int32** | Sort order index of the last resource in the returned array. | [optional] 
**IsMore** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. A value of &#x60;false&#x60; indicates that no more unreturned resources exist.  This field is returned if there are resources in your returned array. | [optional] [default to false]
**StartIndex** | Pointer to **int32** | Sort order index of the first resource in the returned array. | [optional] 

## Methods

### NewAuthControlExemptMidsListResponse

`func NewAuthControlExemptMidsListResponse() *AuthControlExemptMidsListResponse`

NewAuthControlExemptMidsListResponse instantiates a new AuthControlExemptMidsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthControlExemptMidsListResponseWithDefaults

`func NewAuthControlExemptMidsListResponseWithDefaults() *AuthControlExemptMidsListResponse`

NewAuthControlExemptMidsListResponseWithDefaults instantiates a new AuthControlExemptMidsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AuthControlExemptMidsListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AuthControlExemptMidsListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AuthControlExemptMidsListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AuthControlExemptMidsListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *AuthControlExemptMidsListResponse) GetData() []AuthControlExemptMidsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthControlExemptMidsListResponse) GetDataOk() (*[]AuthControlExemptMidsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthControlExemptMidsListResponse) SetData(v []AuthControlExemptMidsResponse)`

SetData sets Data field to given value.

### HasData

`func (o *AuthControlExemptMidsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndIndex

`func (o *AuthControlExemptMidsListResponse) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *AuthControlExemptMidsListResponse) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *AuthControlExemptMidsListResponse) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *AuthControlExemptMidsListResponse) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetIsMore

`func (o *AuthControlExemptMidsListResponse) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *AuthControlExemptMidsListResponse) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *AuthControlExemptMidsListResponse) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.

### HasIsMore

`func (o *AuthControlExemptMidsListResponse) HasIsMore() bool`

HasIsMore returns a boolean if a field has been set.

### GetStartIndex

`func (o *AuthControlExemptMidsListResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *AuthControlExemptMidsListResponse) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *AuthControlExemptMidsListResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *AuthControlExemptMidsListResponse) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


