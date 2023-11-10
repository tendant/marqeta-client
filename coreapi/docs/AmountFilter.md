# AmountFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GreaterThan** | Pointer to **float32** | The minimum amount that a balance for a billing cycle can be to earn the reward. | [optional] 
**LessThan** | Pointer to **float32** | The maximum amount that a balance for a billing cycle can be to earn the reward. | [optional] 

## Methods

### NewAmountFilter

`func NewAmountFilter() *AmountFilter`

NewAmountFilter instantiates a new AmountFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountFilterWithDefaults

`func NewAmountFilterWithDefaults() *AmountFilter`

NewAmountFilterWithDefaults instantiates a new AmountFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreaterThan

`func (o *AmountFilter) GetGreaterThan() float32`

GetGreaterThan returns the GreaterThan field if non-nil, zero value otherwise.

### GetGreaterThanOk

`func (o *AmountFilter) GetGreaterThanOk() (*float32, bool)`

GetGreaterThanOk returns a tuple with the GreaterThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreaterThan

`func (o *AmountFilter) SetGreaterThan(v float32)`

SetGreaterThan sets GreaterThan field to given value.

### HasGreaterThan

`func (o *AmountFilter) HasGreaterThan() bool`

HasGreaterThan returns a boolean if a field has been set.

### GetLessThan

`func (o *AmountFilter) GetLessThan() float32`

GetLessThan returns the LessThan field if non-nil, zero value otherwise.

### GetLessThanOk

`func (o *AmountFilter) GetLessThanOk() (*float32, bool)`

GetLessThanOk returns a tuple with the LessThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessThan

`func (o *AmountFilter) SetLessThan(v float32)`

SetLessThan sets LessThan field to given value.

### HasLessThan

`func (o *AmountFilter) HasLessThan() bool`

HasLessThan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


