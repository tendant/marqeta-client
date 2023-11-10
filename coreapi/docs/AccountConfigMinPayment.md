# AccountConfigMinPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the minimum payment override is currently active. | [optional] 
**MinPaymentFlatAmount** | Pointer to **float32** | Flat amount of the minimum payment override. | [optional] 
**MinPaymentPercentage** | Pointer to **float32** | Percentage of the total statement balance used to calculate the minimum payment override amount. | [optional] 
**OverrideEndTime** | Pointer to **time.Time** | Date and time when the minimum payment override ends, in UTC. | [optional] 
**OverrideStartTime** | Pointer to **time.Time** | Date and time when the minimum payment override starts, in UTC. | [optional] 

## Methods

### NewAccountConfigMinPayment

`func NewAccountConfigMinPayment() *AccountConfigMinPayment`

NewAccountConfigMinPayment instantiates a new AccountConfigMinPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigMinPaymentWithDefaults

`func NewAccountConfigMinPaymentWithDefaults() *AccountConfigMinPayment`

NewAccountConfigMinPaymentWithDefaults instantiates a new AccountConfigMinPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AccountConfigMinPayment) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AccountConfigMinPayment) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AccountConfigMinPayment) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AccountConfigMinPayment) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMinPaymentFlatAmount

`func (o *AccountConfigMinPayment) GetMinPaymentFlatAmount() float32`

GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field if non-nil, zero value otherwise.

### GetMinPaymentFlatAmountOk

`func (o *AccountConfigMinPayment) GetMinPaymentFlatAmountOk() (*float32, bool)`

GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentFlatAmount

`func (o *AccountConfigMinPayment) SetMinPaymentFlatAmount(v float32)`

SetMinPaymentFlatAmount sets MinPaymentFlatAmount field to given value.

### HasMinPaymentFlatAmount

`func (o *AccountConfigMinPayment) HasMinPaymentFlatAmount() bool`

HasMinPaymentFlatAmount returns a boolean if a field has been set.

### GetMinPaymentPercentage

`func (o *AccountConfigMinPayment) GetMinPaymentPercentage() float32`

GetMinPaymentPercentage returns the MinPaymentPercentage field if non-nil, zero value otherwise.

### GetMinPaymentPercentageOk

`func (o *AccountConfigMinPayment) GetMinPaymentPercentageOk() (*float32, bool)`

GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentPercentage

`func (o *AccountConfigMinPayment) SetMinPaymentPercentage(v float32)`

SetMinPaymentPercentage sets MinPaymentPercentage field to given value.

### HasMinPaymentPercentage

`func (o *AccountConfigMinPayment) HasMinPaymentPercentage() bool`

HasMinPaymentPercentage returns a boolean if a field has been set.

### GetOverrideEndTime

`func (o *AccountConfigMinPayment) GetOverrideEndTime() time.Time`

GetOverrideEndTime returns the OverrideEndTime field if non-nil, zero value otherwise.

### GetOverrideEndTimeOk

`func (o *AccountConfigMinPayment) GetOverrideEndTimeOk() (*time.Time, bool)`

GetOverrideEndTimeOk returns a tuple with the OverrideEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideEndTime

`func (o *AccountConfigMinPayment) SetOverrideEndTime(v time.Time)`

SetOverrideEndTime sets OverrideEndTime field to given value.

### HasOverrideEndTime

`func (o *AccountConfigMinPayment) HasOverrideEndTime() bool`

HasOverrideEndTime returns a boolean if a field has been set.

### GetOverrideStartTime

`func (o *AccountConfigMinPayment) GetOverrideStartTime() time.Time`

GetOverrideStartTime returns the OverrideStartTime field if non-nil, zero value otherwise.

### GetOverrideStartTimeOk

`func (o *AccountConfigMinPayment) GetOverrideStartTimeOk() (*time.Time, bool)`

GetOverrideStartTimeOk returns a tuple with the OverrideStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideStartTime

`func (o *AccountConfigMinPayment) SetOverrideStartTime(v time.Time)`

SetOverrideStartTime sets OverrideStartTime field to given value.

### HasOverrideStartTime

`func (o *AccountConfigMinPayment) HasOverrideStartTime() bool`

HasOverrideStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


