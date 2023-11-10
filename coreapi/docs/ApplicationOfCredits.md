# ApplicationOfCredits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CycleType** | [**CycleType**](CycleType.md) |  | 
**Day** | **int32** | Day of the billing cycle when credits are applied. | 

## Methods

### NewApplicationOfCredits

`func NewApplicationOfCredits(cycleType CycleType, day int32, ) *ApplicationOfCredits`

NewApplicationOfCredits instantiates a new ApplicationOfCredits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOfCreditsWithDefaults

`func NewApplicationOfCreditsWithDefaults() *ApplicationOfCredits`

NewApplicationOfCreditsWithDefaults instantiates a new ApplicationOfCredits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycleType

`func (o *ApplicationOfCredits) GetCycleType() CycleType`

GetCycleType returns the CycleType field if non-nil, zero value otherwise.

### GetCycleTypeOk

`func (o *ApplicationOfCredits) GetCycleTypeOk() (*CycleType, bool)`

GetCycleTypeOk returns a tuple with the CycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleType

`func (o *ApplicationOfCredits) SetCycleType(v CycleType)`

SetCycleType sets CycleType field to given value.


### GetDay

`func (o *ApplicationOfCredits) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *ApplicationOfCredits) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *ApplicationOfCredits) SetDay(v int32)`

SetDay sets Day field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


