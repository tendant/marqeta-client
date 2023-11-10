# AprScheduleCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**[]AprScheduleEntryCreateReq**](AprScheduleEntryCreateReq.md) | Contains one or more &#x60;schedule&#x60; objects, which contain information on the annual percentage rates (APRs) associated with the type of balance on the credit account and when they are effective. | 
**Type** | [**AccountAprType**](AccountAprType.md) |  | 

## Methods

### NewAprScheduleCreateReq

`func NewAprScheduleCreateReq(schedule []AprScheduleEntryCreateReq, type_ AccountAprType, ) *AprScheduleCreateReq`

NewAprScheduleCreateReq instantiates a new AprScheduleCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAprScheduleCreateReqWithDefaults

`func NewAprScheduleCreateReqWithDefaults() *AprScheduleCreateReq`

NewAprScheduleCreateReqWithDefaults instantiates a new AprScheduleCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *AprScheduleCreateReq) GetSchedule() []AprScheduleEntryCreateReq`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AprScheduleCreateReq) GetScheduleOk() (*[]AprScheduleEntryCreateReq, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AprScheduleCreateReq) SetSchedule(v []AprScheduleEntryCreateReq)`

SetSchedule sets Schedule field to given value.


### GetType

`func (o *AprScheduleCreateReq) GetType() AccountAprType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AprScheduleCreateReq) GetTypeOk() (*AccountAprType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AprScheduleCreateReq) SetType(v AccountAprType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


