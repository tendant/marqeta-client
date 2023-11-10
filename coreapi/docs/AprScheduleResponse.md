# AprScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the APR is active. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date and time when the APR was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Schedule** | [**[]AprScheduleEntryResponse**](AprScheduleEntryResponse.md) | Contains one or more &#x60;schedule&#x60; objects, which contain information about the annual percentage rates (APRs) associated with the type of balance on the credit account and when they are effective. | 
**Type** | [**AccountAprType**](AccountAprType.md) |  | 
**UpdatedDate** | Pointer to **time.Time** | Date and time when the APR was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewAprScheduleResponse

`func NewAprScheduleResponse(schedule []AprScheduleEntryResponse, type_ AccountAprType, ) *AprScheduleResponse`

NewAprScheduleResponse instantiates a new AprScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAprScheduleResponseWithDefaults

`func NewAprScheduleResponseWithDefaults() *AprScheduleResponse`

NewAprScheduleResponseWithDefaults instantiates a new AprScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AprScheduleResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AprScheduleResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AprScheduleResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AprScheduleResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedDate

`func (o *AprScheduleResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AprScheduleResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AprScheduleResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AprScheduleResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSchedule

`func (o *AprScheduleResponse) GetSchedule() []AprScheduleEntryResponse`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AprScheduleResponse) GetScheduleOk() (*[]AprScheduleEntryResponse, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AprScheduleResponse) SetSchedule(v []AprScheduleEntryResponse)`

SetSchedule sets Schedule field to given value.


### GetType

`func (o *AprScheduleResponse) GetType() AccountAprType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AprScheduleResponse) GetTypeOk() (*AccountAprType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AprScheduleResponse) SetType(v AccountAprType)`

SetType sets Type field to given value.


### GetUpdatedDate

`func (o *AprScheduleResponse) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *AprScheduleResponse) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *AprScheduleResponse) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *AprScheduleResponse) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


