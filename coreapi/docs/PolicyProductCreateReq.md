# PolicyProductCreateReq

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
**Token** | Pointer to **string** | Unique identifier of the credit product policy. | [optional] 
**Usage** | [**[]BalanceType**](BalanceType.md) | One or more usage types for the credit product policy. | 

## Methods

### NewPolicyProductCreateReq

`func NewPolicyProductCreateReq(cardProducts []PolicyProductCardProductReq, classification ProductClassification, creditLine ProductCreditLine, currencyCode CurrencyCode, interestCalculation InterestCalculation, name string, payments PolicyProductPaymentConfiguration, productSubType ProductSubType, productType ProductType, usage []BalanceType, ) *PolicyProductCreateReq`

NewPolicyProductCreateReq instantiates a new PolicyProductCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductCreateReqWithDefaults

`func NewPolicyProductCreateReqWithDefaults() *PolicyProductCreateReq`

NewPolicyProductCreateReqWithDefaults instantiates a new PolicyProductCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *PolicyProductCreateReq) GetCardProducts() []PolicyProductCardProductReq`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *PolicyProductCreateReq) GetCardProductsOk() (*[]PolicyProductCardProductReq, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *PolicyProductCreateReq) SetCardProducts(v []PolicyProductCardProductReq)`

SetCardProducts sets CardProducts field to given value.


### GetClassification

`func (o *PolicyProductCreateReq) GetClassification() ProductClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *PolicyProductCreateReq) GetClassificationOk() (*ProductClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *PolicyProductCreateReq) SetClassification(v ProductClassification)`

SetClassification sets Classification field to given value.


### GetCreditLine

`func (o *PolicyProductCreateReq) GetCreditLine() ProductCreditLine`

GetCreditLine returns the CreditLine field if non-nil, zero value otherwise.

### GetCreditLineOk

`func (o *PolicyProductCreateReq) GetCreditLineOk() (*ProductCreditLine, bool)`

GetCreditLineOk returns a tuple with the CreditLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLine

`func (o *PolicyProductCreateReq) SetCreditLine(v ProductCreditLine)`

SetCreditLine sets CreditLine field to given value.


### GetCurrencyCode

`func (o *PolicyProductCreateReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PolicyProductCreateReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PolicyProductCreateReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *PolicyProductCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyProductCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyProductCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyProductCreateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterestCalculation

`func (o *PolicyProductCreateReq) GetInterestCalculation() InterestCalculation`

GetInterestCalculation returns the InterestCalculation field if non-nil, zero value otherwise.

### GetInterestCalculationOk

`func (o *PolicyProductCreateReq) GetInterestCalculationOk() (*InterestCalculation, bool)`

GetInterestCalculationOk returns a tuple with the InterestCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCalculation

`func (o *PolicyProductCreateReq) SetInterestCalculation(v InterestCalculation)`

SetInterestCalculation sets InterestCalculation field to given value.


### GetName

`func (o *PolicyProductCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyProductCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyProductCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetPayments

`func (o *PolicyProductCreateReq) GetPayments() PolicyProductPaymentConfiguration`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *PolicyProductCreateReq) GetPaymentsOk() (*PolicyProductPaymentConfiguration, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *PolicyProductCreateReq) SetPayments(v PolicyProductPaymentConfiguration)`

SetPayments sets Payments field to given value.


### GetProductSubType

`func (o *PolicyProductCreateReq) GetProductSubType() ProductSubType`

GetProductSubType returns the ProductSubType field if non-nil, zero value otherwise.

### GetProductSubTypeOk

`func (o *PolicyProductCreateReq) GetProductSubTypeOk() (*ProductSubType, bool)`

GetProductSubTypeOk returns a tuple with the ProductSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubType

`func (o *PolicyProductCreateReq) SetProductSubType(v ProductSubType)`

SetProductSubType sets ProductSubType field to given value.


### GetProductType

`func (o *PolicyProductCreateReq) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PolicyProductCreateReq) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PolicyProductCreateReq) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetToken

`func (o *PolicyProductCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyProductCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyProductCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyProductCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsage

`func (o *PolicyProductCreateReq) GetUsage() []BalanceType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PolicyProductCreateReq) GetUsageOk() (*[]BalanceType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PolicyProductCreateReq) SetUsage(v []BalanceType)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


