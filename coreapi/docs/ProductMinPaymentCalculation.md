# ProductMinPaymentCalculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludePastDueAmount** | **bool** | Whether to include the past due amount when calculating the minimum payment. | 
**MinPaymentFlatAmount** | **float32** | Minimum payment, expressed as a flat amount, due on the payment due day. | 
**MinPaymentPercentage** | [**ProductMinPaymentPercentage**](ProductMinPaymentPercentage.md) |  | 

## Methods

### NewProductMinPaymentCalculation

`func NewProductMinPaymentCalculation(includePastDueAmount bool, minPaymentFlatAmount float32, minPaymentPercentage ProductMinPaymentPercentage, ) *ProductMinPaymentCalculation`

NewProductMinPaymentCalculation instantiates a new ProductMinPaymentCalculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductMinPaymentCalculationWithDefaults

`func NewProductMinPaymentCalculationWithDefaults() *ProductMinPaymentCalculation`

NewProductMinPaymentCalculationWithDefaults instantiates a new ProductMinPaymentCalculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludePastDueAmount

`func (o *ProductMinPaymentCalculation) GetIncludePastDueAmount() bool`

GetIncludePastDueAmount returns the IncludePastDueAmount field if non-nil, zero value otherwise.

### GetIncludePastDueAmountOk

`func (o *ProductMinPaymentCalculation) GetIncludePastDueAmountOk() (*bool, bool)`

GetIncludePastDueAmountOk returns a tuple with the IncludePastDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePastDueAmount

`func (o *ProductMinPaymentCalculation) SetIncludePastDueAmount(v bool)`

SetIncludePastDueAmount sets IncludePastDueAmount field to given value.


### GetMinPaymentFlatAmount

`func (o *ProductMinPaymentCalculation) GetMinPaymentFlatAmount() float32`

GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field if non-nil, zero value otherwise.

### GetMinPaymentFlatAmountOk

`func (o *ProductMinPaymentCalculation) GetMinPaymentFlatAmountOk() (*float32, bool)`

GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentFlatAmount

`func (o *ProductMinPaymentCalculation) SetMinPaymentFlatAmount(v float32)`

SetMinPaymentFlatAmount sets MinPaymentFlatAmount field to given value.


### GetMinPaymentPercentage

`func (o *ProductMinPaymentCalculation) GetMinPaymentPercentage() ProductMinPaymentPercentage`

GetMinPaymentPercentage returns the MinPaymentPercentage field if non-nil, zero value otherwise.

### GetMinPaymentPercentageOk

`func (o *ProductMinPaymentCalculation) GetMinPaymentPercentageOk() (*ProductMinPaymentPercentage, bool)`

GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentPercentage

`func (o *ProductMinPaymentCalculation) SetMinPaymentPercentage(v ProductMinPaymentPercentage)`

SetMinPaymentPercentage sets MinPaymentPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


