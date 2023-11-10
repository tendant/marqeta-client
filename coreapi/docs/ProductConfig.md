# ProductConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleDay** | **int32** | Day of month the billing cycle starts. | 
**Fees** | Pointer to [**[]ProductFeeType**](ProductFeeType.md) | One or more fee types. | [optional] 
**PaymentDueDay** | **int32** | Day of month the payment for the previous billing cycle is due. | 
**PeriodicFees** | Pointer to [**[]ProductConfigPeriodicFeesInner**](ProductConfigPeriodicFeesInner.md) | Contains one or more periodic fees. | [optional] 

## Methods

### NewProductConfig

`func NewProductConfig(billingCycleDay int32, paymentDueDay int32, ) *ProductConfig`

NewProductConfig instantiates a new ProductConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductConfigWithDefaults

`func NewProductConfigWithDefaults() *ProductConfig`

NewProductConfigWithDefaults instantiates a new ProductConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleDay

`func (o *ProductConfig) GetBillingCycleDay() int32`

GetBillingCycleDay returns the BillingCycleDay field if non-nil, zero value otherwise.

### GetBillingCycleDayOk

`func (o *ProductConfig) GetBillingCycleDayOk() (*int32, bool)`

GetBillingCycleDayOk returns a tuple with the BillingCycleDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleDay

`func (o *ProductConfig) SetBillingCycleDay(v int32)`

SetBillingCycleDay sets BillingCycleDay field to given value.


### GetFees

`func (o *ProductConfig) GetFees() []ProductFeeType`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ProductConfig) GetFeesOk() (*[]ProductFeeType, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ProductConfig) SetFees(v []ProductFeeType)`

SetFees sets Fees field to given value.

### HasFees

`func (o *ProductConfig) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetPaymentDueDay

`func (o *ProductConfig) GetPaymentDueDay() int32`

GetPaymentDueDay returns the PaymentDueDay field if non-nil, zero value otherwise.

### GetPaymentDueDayOk

`func (o *ProductConfig) GetPaymentDueDayOk() (*int32, bool)`

GetPaymentDueDayOk returns a tuple with the PaymentDueDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDay

`func (o *ProductConfig) SetPaymentDueDay(v int32)`

SetPaymentDueDay sets PaymentDueDay field to given value.


### GetPeriodicFees

`func (o *ProductConfig) GetPeriodicFees() []ProductConfigPeriodicFeesInner`

GetPeriodicFees returns the PeriodicFees field if non-nil, zero value otherwise.

### GetPeriodicFeesOk

`func (o *ProductConfig) GetPeriodicFeesOk() (*[]ProductConfigPeriodicFeesInner, bool)`

GetPeriodicFeesOk returns a tuple with the PeriodicFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicFees

`func (o *ProductConfig) SetPeriodicFees(v []ProductConfigPeriodicFeesInner)`

SetPeriodicFees sets PeriodicFees field to given value.

### HasPeriodicFees

`func (o *ProductConfig) HasPeriodicFees() bool`

HasPeriodicFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


