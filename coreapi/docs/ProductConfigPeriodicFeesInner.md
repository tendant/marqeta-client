# ProductConfigPeriodicFeesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **string** | How frequently the fee is charged. | 
**NumberOfDaysPostActivation** | **int32** | The number of days after an account is activated that the initial fee is charged. For example, if the value in this field is &#x60;30&#x60;, then the initial fee is charged 30 days after an account is activated. | 

## Methods

### NewProductConfigPeriodicFeesInner

`func NewProductConfigPeriodicFeesInner(frequency string, numberOfDaysPostActivation int32, ) *ProductConfigPeriodicFeesInner`

NewProductConfigPeriodicFeesInner instantiates a new ProductConfigPeriodicFeesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductConfigPeriodicFeesInnerWithDefaults

`func NewProductConfigPeriodicFeesInnerWithDefaults() *ProductConfigPeriodicFeesInner`

NewProductConfigPeriodicFeesInnerWithDefaults instantiates a new ProductConfigPeriodicFeesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *ProductConfigPeriodicFeesInner) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ProductConfigPeriodicFeesInner) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ProductConfigPeriodicFeesInner) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetNumberOfDaysPostActivation

`func (o *ProductConfigPeriodicFeesInner) GetNumberOfDaysPostActivation() int32`

GetNumberOfDaysPostActivation returns the NumberOfDaysPostActivation field if non-nil, zero value otherwise.

### GetNumberOfDaysPostActivationOk

`func (o *ProductConfigPeriodicFeesInner) GetNumberOfDaysPostActivationOk() (*int32, bool)`

GetNumberOfDaysPostActivationOk returns a tuple with the NumberOfDaysPostActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysPostActivation

`func (o *ProductConfigPeriodicFeesInner) SetNumberOfDaysPostActivation(v int32)`

SetNumberOfDaysPostActivation sets NumberOfDaysPostActivation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


