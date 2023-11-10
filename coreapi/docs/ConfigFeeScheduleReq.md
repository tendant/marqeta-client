# ConfigFeeScheduleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**[]ConfigFeeScheduleEntry**](ConfigFeeScheduleEntry.md) | Contains one or more fee schedules. | 
**Type** | [**AccountProductFeeType**](AccountProductFeeType.md) |  | 

## Methods

### NewConfigFeeScheduleReq

`func NewConfigFeeScheduleReq(schedule []ConfigFeeScheduleEntry, type_ AccountProductFeeType, ) *ConfigFeeScheduleReq`

NewConfigFeeScheduleReq instantiates a new ConfigFeeScheduleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFeeScheduleReqWithDefaults

`func NewConfigFeeScheduleReqWithDefaults() *ConfigFeeScheduleReq`

NewConfigFeeScheduleReqWithDefaults instantiates a new ConfigFeeScheduleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *ConfigFeeScheduleReq) GetSchedule() []ConfigFeeScheduleEntry`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ConfigFeeScheduleReq) GetScheduleOk() (*[]ConfigFeeScheduleEntry, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ConfigFeeScheduleReq) SetSchedule(v []ConfigFeeScheduleEntry)`

SetSchedule sets Schedule field to given value.


### GetType

`func (o *ConfigFeeScheduleReq) GetType() AccountProductFeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigFeeScheduleReq) GetTypeOk() (*AccountProductFeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigFeeScheduleReq) SetType(v AccountProductFeeType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


