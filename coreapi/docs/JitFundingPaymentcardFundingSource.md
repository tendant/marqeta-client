# JitFundingPaymentcardFundingSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether JIT Funding is enabled or disabled for the payment card funding source. A value of &#x60;true&#x60; indicates that the payment card funding source is enabled and will be debited when swipes occur. | [optional] [default to false]
**RefundsDestination** | Pointer to **string** | Specifies the return destination for refunds in the case of a transaction reversal. | [optional] 

## Methods

### NewJitFundingPaymentcardFundingSource

`func NewJitFundingPaymentcardFundingSource() *JitFundingPaymentcardFundingSource`

NewJitFundingPaymentcardFundingSource instantiates a new JitFundingPaymentcardFundingSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitFundingPaymentcardFundingSourceWithDefaults

`func NewJitFundingPaymentcardFundingSourceWithDefaults() *JitFundingPaymentcardFundingSource`

NewJitFundingPaymentcardFundingSourceWithDefaults instantiates a new JitFundingPaymentcardFundingSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *JitFundingPaymentcardFundingSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JitFundingPaymentcardFundingSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JitFundingPaymentcardFundingSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JitFundingPaymentcardFundingSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefundsDestination

`func (o *JitFundingPaymentcardFundingSource) GetRefundsDestination() string`

GetRefundsDestination returns the RefundsDestination field if non-nil, zero value otherwise.

### GetRefundsDestinationOk

`func (o *JitFundingPaymentcardFundingSource) GetRefundsDestinationOk() (*string, bool)`

GetRefundsDestinationOk returns a tuple with the RefundsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundsDestination

`func (o *JitFundingPaymentcardFundingSource) SetRefundsDestination(v string)`

SetRefundsDestination sets RefundsDestination field to given value.

### HasRefundsDestination

`func (o *JitFundingPaymentcardFundingSource) HasRefundsDestination() bool`

HasRefundsDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


