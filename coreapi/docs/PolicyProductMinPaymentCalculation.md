# PolicyProductMinPaymentCalculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeOverlimitAmount** | **bool** | Whether to include the overlimit amount when calculating the minimum payment. | 
**IncludePastDueAmount** | **bool** | Whether to include the past due amount when calculating the minimum payment. | 
**MinPaymentFlatAmount** | Pointer to **float32** | Minimum payment, expressed as a flat amount, due on the payment due day. | [optional] 
**MinPaymentPercentage** | Pointer to [**PolicyProductMinPaymentPercentage**](PolicyProductMinPaymentPercentage.md) |  | [optional] 

## Methods

### NewPolicyProductMinPaymentCalculation

`func NewPolicyProductMinPaymentCalculation(includeOverlimitAmount bool, includePastDueAmount bool, ) *PolicyProductMinPaymentCalculation`

NewPolicyProductMinPaymentCalculation instantiates a new PolicyProductMinPaymentCalculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductMinPaymentCalculationWithDefaults

`func NewPolicyProductMinPaymentCalculationWithDefaults() *PolicyProductMinPaymentCalculation`

NewPolicyProductMinPaymentCalculationWithDefaults instantiates a new PolicyProductMinPaymentCalculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeOverlimitAmount

`func (o *PolicyProductMinPaymentCalculation) GetIncludeOverlimitAmount() bool`

GetIncludeOverlimitAmount returns the IncludeOverlimitAmount field if non-nil, zero value otherwise.

### GetIncludeOverlimitAmountOk

`func (o *PolicyProductMinPaymentCalculation) GetIncludeOverlimitAmountOk() (*bool, bool)`

GetIncludeOverlimitAmountOk returns a tuple with the IncludeOverlimitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOverlimitAmount

`func (o *PolicyProductMinPaymentCalculation) SetIncludeOverlimitAmount(v bool)`

SetIncludeOverlimitAmount sets IncludeOverlimitAmount field to given value.


### GetIncludePastDueAmount

`func (o *PolicyProductMinPaymentCalculation) GetIncludePastDueAmount() bool`

GetIncludePastDueAmount returns the IncludePastDueAmount field if non-nil, zero value otherwise.

### GetIncludePastDueAmountOk

`func (o *PolicyProductMinPaymentCalculation) GetIncludePastDueAmountOk() (*bool, bool)`

GetIncludePastDueAmountOk returns a tuple with the IncludePastDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePastDueAmount

`func (o *PolicyProductMinPaymentCalculation) SetIncludePastDueAmount(v bool)`

SetIncludePastDueAmount sets IncludePastDueAmount field to given value.


### GetMinPaymentFlatAmount

`func (o *PolicyProductMinPaymentCalculation) GetMinPaymentFlatAmount() float32`

GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field if non-nil, zero value otherwise.

### GetMinPaymentFlatAmountOk

`func (o *PolicyProductMinPaymentCalculation) GetMinPaymentFlatAmountOk() (*float32, bool)`

GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentFlatAmount

`func (o *PolicyProductMinPaymentCalculation) SetMinPaymentFlatAmount(v float32)`

SetMinPaymentFlatAmount sets MinPaymentFlatAmount field to given value.

### HasMinPaymentFlatAmount

`func (o *PolicyProductMinPaymentCalculation) HasMinPaymentFlatAmount() bool`

HasMinPaymentFlatAmount returns a boolean if a field has been set.

### GetMinPaymentPercentage

`func (o *PolicyProductMinPaymentCalculation) GetMinPaymentPercentage() PolicyProductMinPaymentPercentage`

GetMinPaymentPercentage returns the MinPaymentPercentage field if non-nil, zero value otherwise.

### GetMinPaymentPercentageOk

`func (o *PolicyProductMinPaymentCalculation) GetMinPaymentPercentageOk() (*PolicyProductMinPaymentPercentage, bool)`

GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentPercentage

`func (o *PolicyProductMinPaymentCalculation) SetMinPaymentPercentage(v PolicyProductMinPaymentPercentage)`

SetMinPaymentPercentage sets MinPaymentPercentage field to given value.

### HasMinPaymentPercentage

`func (o *PolicyProductMinPaymentCalculation) HasMinPaymentPercentage() bool`

HasMinPaymentPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


