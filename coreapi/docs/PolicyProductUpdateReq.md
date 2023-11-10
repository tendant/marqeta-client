# PolicyProductUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | [**[]PolicyProductCardProductReq**](PolicyProductCardProductReq.md) | One or more card products associated with the credit product policy. | 
**Classification** | [**ProductClassification**](ProductClassification.md) |  | [default to PRODUCTCLASSIFICATION_CONSUMER]
**CreditLine** | [**ProductCreditLine**](ProductCreditLine.md) |  | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | Pointer to **string** | Description of the credit product policy. | [optional] 
**InterestCalculation** | [**InterestCalculation**](InterestCalculation.md) |  | 
**Name** | **string** | Name of the credit product policy. | 
**Payments** | [**PolicyProductPaymentConfiguration**](PolicyProductPaymentConfiguration.md) |  | 
**ProductSubType** | [**ProductSubType**](ProductSubType.md) |  | [default to PRODUCTSUBTYPE_CREDIT_CARD]
**ProductType** | [**ProductType**](ProductType.md) |  | [default to PRODUCTTYPE_REVOLVING]
**Usage** | [**[]BalanceType**](BalanceType.md) | One or more usage types for the credit product policy. | 

## Methods

### NewPolicyProductUpdateReq

`func NewPolicyProductUpdateReq(cardProducts []PolicyProductCardProductReq, classification ProductClassification, creditLine ProductCreditLine, currencyCode CurrencyCode, interestCalculation InterestCalculation, name string, payments PolicyProductPaymentConfiguration, productSubType ProductSubType, productType ProductType, usage []BalanceType, ) *PolicyProductUpdateReq`

NewPolicyProductUpdateReq instantiates a new PolicyProductUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductUpdateReqWithDefaults

`func NewPolicyProductUpdateReqWithDefaults() *PolicyProductUpdateReq`

NewPolicyProductUpdateReqWithDefaults instantiates a new PolicyProductUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *PolicyProductUpdateReq) GetCardProducts() []PolicyProductCardProductReq`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *PolicyProductUpdateReq) GetCardProductsOk() (*[]PolicyProductCardProductReq, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *PolicyProductUpdateReq) SetCardProducts(v []PolicyProductCardProductReq)`

SetCardProducts sets CardProducts field to given value.


### GetClassification

`func (o *PolicyProductUpdateReq) GetClassification() ProductClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *PolicyProductUpdateReq) GetClassificationOk() (*ProductClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *PolicyProductUpdateReq) SetClassification(v ProductClassification)`

SetClassification sets Classification field to given value.


### GetCreditLine

`func (o *PolicyProductUpdateReq) GetCreditLine() ProductCreditLine`

GetCreditLine returns the CreditLine field if non-nil, zero value otherwise.

### GetCreditLineOk

`func (o *PolicyProductUpdateReq) GetCreditLineOk() (*ProductCreditLine, bool)`

GetCreditLineOk returns a tuple with the CreditLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLine

`func (o *PolicyProductUpdateReq) SetCreditLine(v ProductCreditLine)`

SetCreditLine sets CreditLine field to given value.


### GetCurrencyCode

`func (o *PolicyProductUpdateReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PolicyProductUpdateReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PolicyProductUpdateReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *PolicyProductUpdateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyProductUpdateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyProductUpdateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyProductUpdateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterestCalculation

`func (o *PolicyProductUpdateReq) GetInterestCalculation() InterestCalculation`

GetInterestCalculation returns the InterestCalculation field if non-nil, zero value otherwise.

### GetInterestCalculationOk

`func (o *PolicyProductUpdateReq) GetInterestCalculationOk() (*InterestCalculation, bool)`

GetInterestCalculationOk returns a tuple with the InterestCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCalculation

`func (o *PolicyProductUpdateReq) SetInterestCalculation(v InterestCalculation)`

SetInterestCalculation sets InterestCalculation field to given value.


### GetName

`func (o *PolicyProductUpdateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyProductUpdateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyProductUpdateReq) SetName(v string)`

SetName sets Name field to given value.


### GetPayments

`func (o *PolicyProductUpdateReq) GetPayments() PolicyProductPaymentConfiguration`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *PolicyProductUpdateReq) GetPaymentsOk() (*PolicyProductPaymentConfiguration, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *PolicyProductUpdateReq) SetPayments(v PolicyProductPaymentConfiguration)`

SetPayments sets Payments field to given value.


### GetProductSubType

`func (o *PolicyProductUpdateReq) GetProductSubType() ProductSubType`

GetProductSubType returns the ProductSubType field if non-nil, zero value otherwise.

### GetProductSubTypeOk

`func (o *PolicyProductUpdateReq) GetProductSubTypeOk() (*ProductSubType, bool)`

GetProductSubTypeOk returns a tuple with the ProductSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubType

`func (o *PolicyProductUpdateReq) SetProductSubType(v ProductSubType)`

SetProductSubType sets ProductSubType field to given value.


### GetProductType

`func (o *PolicyProductUpdateReq) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PolicyProductUpdateReq) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PolicyProductUpdateReq) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetUsage

`func (o *PolicyProductUpdateReq) GetUsage() []BalanceType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PolicyProductUpdateReq) GetUsageOk() (*[]BalanceType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PolicyProductUpdateReq) SetUsage(v []BalanceType)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


