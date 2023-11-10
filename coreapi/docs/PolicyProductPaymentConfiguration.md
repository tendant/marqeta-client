# PolicyProductPaymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationOrder** | [**[]PaymentAllocationOrderEnum**](PaymentAllocationOrderEnum.md) | Ordered list of balance types to which payments are allocated, from first to last. | [default to ["INTEREST","FEES","PRINCIPAL"]]
**BillingCycleDay** | **int32** | Day of month the billing cycle starts. | 
**DueDay** | **int32** | Day of month the payment for the previous billing cycle is due. | 
**MinPaymentCalculation** | [**PolicyProductMinPaymentCalculation**](PolicyProductMinPaymentCalculation.md) |  | 

## Methods

### NewPolicyProductPaymentConfiguration

`func NewPolicyProductPaymentConfiguration(allocationOrder []PaymentAllocationOrderEnum, billingCycleDay int32, dueDay int32, minPaymentCalculation PolicyProductMinPaymentCalculation, ) *PolicyProductPaymentConfiguration`

NewPolicyProductPaymentConfiguration instantiates a new PolicyProductPaymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductPaymentConfigurationWithDefaults

`func NewPolicyProductPaymentConfigurationWithDefaults() *PolicyProductPaymentConfiguration`

NewPolicyProductPaymentConfigurationWithDefaults instantiates a new PolicyProductPaymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationOrder

`func (o *PolicyProductPaymentConfiguration) GetAllocationOrder() []PaymentAllocationOrderEnum`

GetAllocationOrder returns the AllocationOrder field if non-nil, zero value otherwise.

### GetAllocationOrderOk

`func (o *PolicyProductPaymentConfiguration) GetAllocationOrderOk() (*[]PaymentAllocationOrderEnum, bool)`

GetAllocationOrderOk returns a tuple with the AllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationOrder

`func (o *PolicyProductPaymentConfiguration) SetAllocationOrder(v []PaymentAllocationOrderEnum)`

SetAllocationOrder sets AllocationOrder field to given value.


### GetBillingCycleDay

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleDay() int32`

GetBillingCycleDay returns the BillingCycleDay field if non-nil, zero value otherwise.

### GetBillingCycleDayOk

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleDayOk() (*int32, bool)`

GetBillingCycleDayOk returns a tuple with the BillingCycleDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleDay

`func (o *PolicyProductPaymentConfiguration) SetBillingCycleDay(v int32)`

SetBillingCycleDay sets BillingCycleDay field to given value.


### GetDueDay

`func (o *PolicyProductPaymentConfiguration) GetDueDay() int32`

GetDueDay returns the DueDay field if non-nil, zero value otherwise.

### GetDueDayOk

`func (o *PolicyProductPaymentConfiguration) GetDueDayOk() (*int32, bool)`

GetDueDayOk returns a tuple with the DueDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDay

`func (o *PolicyProductPaymentConfiguration) SetDueDay(v int32)`

SetDueDay sets DueDay field to given value.


### GetMinPaymentCalculation

`func (o *PolicyProductPaymentConfiguration) GetMinPaymentCalculation() PolicyProductMinPaymentCalculation`

GetMinPaymentCalculation returns the MinPaymentCalculation field if non-nil, zero value otherwise.

### GetMinPaymentCalculationOk

`func (o *PolicyProductPaymentConfiguration) GetMinPaymentCalculationOk() (*PolicyProductMinPaymentCalculation, bool)`

GetMinPaymentCalculationOk returns a tuple with the MinPaymentCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentCalculation

`func (o *PolicyProductPaymentConfiguration) SetMinPaymentCalculation(v PolicyProductMinPaymentCalculation)`

SetMinPaymentCalculation sets MinPaymentCalculation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


