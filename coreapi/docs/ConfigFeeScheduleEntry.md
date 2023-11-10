# ConfigFeeScheduleEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveDate** | Pointer to **time.Time** | Date and time when the fee goes into effect, in UTC. | [optional] 
**Method** | [**FeeMethod**](FeeMethod.md) |  | 
**Value** | **float32** | Amount of the fee. | 

## Methods

### NewConfigFeeScheduleEntry

`func NewConfigFeeScheduleEntry(method FeeMethod, value float32, ) *ConfigFeeScheduleEntry`

NewConfigFeeScheduleEntry instantiates a new ConfigFeeScheduleEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFeeScheduleEntryWithDefaults

`func NewConfigFeeScheduleEntryWithDefaults() *ConfigFeeScheduleEntry`

NewConfigFeeScheduleEntryWithDefaults instantiates a new ConfigFeeScheduleEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveDate

`func (o *ConfigFeeScheduleEntry) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ConfigFeeScheduleEntry) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ConfigFeeScheduleEntry) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ConfigFeeScheduleEntry) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetMethod

`func (o *ConfigFeeScheduleEntry) GetMethod() FeeMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConfigFeeScheduleEntry) GetMethodOk() (*FeeMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConfigFeeScheduleEntry) SetMethod(v FeeMethod)`

SetMethod sets Method field to given value.


### GetValue

`func (o *ConfigFeeScheduleEntry) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigFeeScheduleEntry) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigFeeScheduleEntry) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


