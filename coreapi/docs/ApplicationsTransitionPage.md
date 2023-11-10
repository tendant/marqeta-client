# ApplicationsTransitionPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of resources to retrieve. | 
**Data** | [**[]ApplicationTransitionResponse**](ApplicationTransitionResponse.md) | An array of application transition objects. | 
**EndIndex** | **int32** | Sort order index of the last resource in the returned array. | 
**IsMore** | **bool** | A value of &#x60;true&#x60; indicates that more unreturned resources exist. | 
**StartIndex** | **int32** | Sort order index of the first resource in the returned array. | 

## Methods

### NewApplicationsTransitionPage

`func NewApplicationsTransitionPage(count int32, data []ApplicationTransitionResponse, endIndex int32, isMore bool, startIndex int32, ) *ApplicationsTransitionPage`

NewApplicationsTransitionPage instantiates a new ApplicationsTransitionPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsTransitionPageWithDefaults

`func NewApplicationsTransitionPageWithDefaults() *ApplicationsTransitionPage`

NewApplicationsTransitionPageWithDefaults instantiates a new ApplicationsTransitionPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ApplicationsTransitionPage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApplicationsTransitionPage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApplicationsTransitionPage) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *ApplicationsTransitionPage) GetData() []ApplicationTransitionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationsTransitionPage) GetDataOk() (*[]ApplicationTransitionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationsTransitionPage) SetData(v []ApplicationTransitionResponse)`

SetData sets Data field to given value.


### GetEndIndex

`func (o *ApplicationsTransitionPage) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *ApplicationsTransitionPage) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *ApplicationsTransitionPage) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.


### GetIsMore

`func (o *ApplicationsTransitionPage) GetIsMore() bool`

GetIsMore returns the IsMore field if non-nil, zero value otherwise.

### GetIsMoreOk

`func (o *ApplicationsTransitionPage) GetIsMoreOk() (*bool, bool)`

GetIsMoreOk returns a tuple with the IsMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMore

`func (o *ApplicationsTransitionPage) SetIsMore(v bool)`

SetIsMore sets IsMore field to given value.


### GetStartIndex

`func (o *ApplicationsTransitionPage) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ApplicationsTransitionPage) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ApplicationsTransitionPage) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


