# ConfigFeeScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the fee is active. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date and time when the fee was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Schedule** | Pointer to [**[]ConfigFeeScheduleEntry**](ConfigFeeScheduleEntry.md) | Contains one or more fee schedules. | [optional] 
**Type** | Pointer to [**AccountProductFeeType**](AccountProductFeeType.md) |  | [optional] 
**UpdatedDate** | Pointer to **time.Time** | Date and time when the fee was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewConfigFeeScheduleResponse

`func NewConfigFeeScheduleResponse() *ConfigFeeScheduleResponse`

NewConfigFeeScheduleResponse instantiates a new ConfigFeeScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFeeScheduleResponseWithDefaults

`func NewConfigFeeScheduleResponseWithDefaults() *ConfigFeeScheduleResponse`

NewConfigFeeScheduleResponseWithDefaults instantiates a new ConfigFeeScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ConfigFeeScheduleResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConfigFeeScheduleResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConfigFeeScheduleResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConfigFeeScheduleResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ConfigFeeScheduleResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ConfigFeeScheduleResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ConfigFeeScheduleResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ConfigFeeScheduleResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSchedule

`func (o *ConfigFeeScheduleResponse) GetSchedule() []ConfigFeeScheduleEntry`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ConfigFeeScheduleResponse) GetScheduleOk() (*[]ConfigFeeScheduleEntry, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ConfigFeeScheduleResponse) SetSchedule(v []ConfigFeeScheduleEntry)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ConfigFeeScheduleResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetType

`func (o *ConfigFeeScheduleResponse) GetType() AccountProductFeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigFeeScheduleResponse) GetTypeOk() (*AccountProductFeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigFeeScheduleResponse) SetType(v AccountProductFeeType)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigFeeScheduleResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ConfigFeeScheduleResponse) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ConfigFeeScheduleResponse) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ConfigFeeScheduleResponse) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ConfigFeeScheduleResponse) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


