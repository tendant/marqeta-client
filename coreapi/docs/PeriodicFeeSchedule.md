# PeriodicFeeSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | fee amount | [optional] 
**EffectiveDate** | Pointer to **string** | date the fee becomes effective | [optional] 
**NextFeeImpactDate** | Pointer to **string** | date of the next time fee will be charged | [optional] 
**Type** | Pointer to **string** | type of fee to be charged | [optional] 

## Methods

### NewPeriodicFeeSchedule

`func NewPeriodicFeeSchedule() *PeriodicFeeSchedule`

NewPeriodicFeeSchedule instantiates a new PeriodicFeeSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodicFeeScheduleWithDefaults

`func NewPeriodicFeeScheduleWithDefaults() *PeriodicFeeSchedule`

NewPeriodicFeeScheduleWithDefaults instantiates a new PeriodicFeeSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PeriodicFeeSchedule) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PeriodicFeeSchedule) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PeriodicFeeSchedule) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PeriodicFeeSchedule) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *PeriodicFeeSchedule) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *PeriodicFeeSchedule) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *PeriodicFeeSchedule) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *PeriodicFeeSchedule) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetNextFeeImpactDate

`func (o *PeriodicFeeSchedule) GetNextFeeImpactDate() string`

GetNextFeeImpactDate returns the NextFeeImpactDate field if non-nil, zero value otherwise.

### GetNextFeeImpactDateOk

`func (o *PeriodicFeeSchedule) GetNextFeeImpactDateOk() (*string, bool)`

GetNextFeeImpactDateOk returns a tuple with the NextFeeImpactDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFeeImpactDate

`func (o *PeriodicFeeSchedule) SetNextFeeImpactDate(v string)`

SetNextFeeImpactDate sets NextFeeImpactDate field to given value.

### HasNextFeeImpactDate

`func (o *PeriodicFeeSchedule) HasNextFeeImpactDate() bool`

HasNextFeeImpactDate returns a boolean if a field has been set.

### GetType

`func (o *PeriodicFeeSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PeriodicFeeSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PeriodicFeeSchedule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PeriodicFeeSchedule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


