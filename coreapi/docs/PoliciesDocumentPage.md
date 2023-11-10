# PoliciesDocumentPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of resources returned. | 
**Data** | [**[]PolicyDocumentResponse**](PolicyDocumentResponse.md) | One or more document policies. | 
**EndIndex** | **int32** | Sort order index of the last resource in the returned array. | 
**IsMore** | **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. | 
**StartIndex** | **int32** | Sort order index of the first resource in the returned array. | 

## Methods

### NewPoliciesDocumentPage

`func NewPoliciesDocumentPage(count int32, data []PolicyDocumentResponse, endIndex int32, isMore bool, startIndex int32, ) *PoliciesDocumentPage`

NewPoliciesDocumentPage instantiates a new PoliciesDocumentPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesDocumentPageWithDefaults

`func NewPoliciesDocumentPageWithDefaults() *PoliciesDocumentPage`

NewPoliciesDocumentPageWithDefaults instantiates a new PoliciesDocumentPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PoliciesDocumentPage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PoliciesDocumentPage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PoliciesDocumentPage) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *PoliciesDocumentPage) GetData() []PolicyDocumentResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PoliciesDocumentPage) GetDataOk() (*[]PolicyDocumentResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PoliciesDocumentPage) SetData(v []PolicyDocumentResponse)`

SetData sets Data field to given value.


### GetEndIndex

`func (o *PoliciesDocumentPage) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *PoliciesDocumentPage) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *PoliciesDocumentPage) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.


### GetIsMore

`func (o *PoliciesDocumentPage) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *PoliciesDocumentPage) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *PoliciesDocumentPage) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.


### GetStartIndex

`func (o *PoliciesDocumentPage) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *PoliciesDocumentPage) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *PoliciesDocumentPage) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


