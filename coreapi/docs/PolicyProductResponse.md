# PolicyProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | Pointer to [**[]PolicyProductCardProductResponse**](PolicyProductCardProductResponse.md) | One or more card products associated with the credit product policy. | [optional] 
**Classification** | Pointer to [**ProductClassification**](ProductClassification.md) |  | [optional] [default to PRODUCTCLASSIFICATION_CONSUMER]
**CreatedTime** | Pointer to **time.Time** | Date and time when the credit product policy was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**CreditLine** | Pointer to [**ProductCreditLine**](ProductCreditLine.md) |  | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] [default to CURRENCYCODE_USD]
**Description** | Pointer to **string** | Description of the credit product policy. | [optional] 
**InterestCalculation** | Pointer to [**InterestCalculation**](InterestCalculation.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the credit product policy. | [optional] 
**Payments** | Pointer to [**PolicyProductPaymentConfiguration**](PolicyProductPaymentConfiguration.md) |  | [optional] 
**ProductSubType** | Pointer to [**ProductSubType**](ProductSubType.md) |  | [optional] [default to PRODUCTSUBTYPE_CREDIT_CARD]
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] [default to PRODUCTTYPE_REVOLVING]
**Token** | Pointer to **string** | Unique identifier of the credit product policy. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the credit product policy was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Usage** | Pointer to [**[]BalanceType**](BalanceType.md) | One or more usage types for the credit product policy. | [optional] 

## Methods

### NewPolicyProductResponse

`func NewPolicyProductResponse() *PolicyProductResponse`

NewPolicyProductResponse instantiates a new PolicyProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductResponseWithDefaults

`func NewPolicyProductResponseWithDefaults() *PolicyProductResponse`

NewPolicyProductResponseWithDefaults instantiates a new PolicyProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *PolicyProductResponse) GetCardProducts() []PolicyProductCardProductResponse`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *PolicyProductResponse) GetCardProductsOk() (*[]PolicyProductCardProductResponse, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *PolicyProductResponse) SetCardProducts(v []PolicyProductCardProductResponse)`

SetCardProducts sets CardProducts field to given value.

### HasCardProducts

`func (o *PolicyProductResponse) HasCardProducts() bool`

HasCardProducts returns a boolean if a field has been set.

### GetClassification

`func (o *PolicyProductResponse) GetClassification() ProductClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *PolicyProductResponse) GetClassificationOk() (*ProductClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *PolicyProductResponse) SetClassification(v ProductClassification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *PolicyProductResponse) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PolicyProductResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PolicyProductResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PolicyProductResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PolicyProductResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreditLine

`func (o *PolicyProductResponse) GetCreditLine() ProductCreditLine`

GetCreditLine returns the CreditLine field if non-nil, zero value otherwise.

### GetCreditLineOk

`func (o *PolicyProductResponse) GetCreditLineOk() (*ProductCreditLine, bool)`

GetCreditLineOk returns a tuple with the CreditLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLine

`func (o *PolicyProductResponse) SetCreditLine(v ProductCreditLine)`

SetCreditLine sets CreditLine field to given value.

### HasCreditLine

`func (o *PolicyProductResponse) HasCreditLine() bool`

HasCreditLine returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PolicyProductResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PolicyProductResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PolicyProductResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PolicyProductResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyProductResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyProductResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyProductResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyProductResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterestCalculation

`func (o *PolicyProductResponse) GetInterestCalculation() InterestCalculation`

GetInterestCalculation returns the InterestCalculation field if non-nil, zero value otherwise.

### GetInterestCalculationOk

`func (o *PolicyProductResponse) GetInterestCalculationOk() (*InterestCalculation, bool)`

GetInterestCalculationOk returns a tuple with the InterestCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCalculation

`func (o *PolicyProductResponse) SetInterestCalculation(v InterestCalculation)`

SetInterestCalculation sets InterestCalculation field to given value.

### HasInterestCalculation

`func (o *PolicyProductResponse) HasInterestCalculation() bool`

HasInterestCalculation returns a boolean if a field has been set.

### GetName

`func (o *PolicyProductResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyProductResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyProductResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyProductResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayments

`func (o *PolicyProductResponse) GetPayments() PolicyProductPaymentConfiguration`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *PolicyProductResponse) GetPaymentsOk() (*PolicyProductPaymentConfiguration, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *PolicyProductResponse) SetPayments(v PolicyProductPaymentConfiguration)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *PolicyProductResponse) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetProductSubType

`func (o *PolicyProductResponse) GetProductSubType() ProductSubType`

GetProductSubType returns the ProductSubType field if non-nil, zero value otherwise.

### GetProductSubTypeOk

`func (o *PolicyProductResponse) GetProductSubTypeOk() (*ProductSubType, bool)`

GetProductSubTypeOk returns a tuple with the ProductSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubType

`func (o *PolicyProductResponse) SetProductSubType(v ProductSubType)`

SetProductSubType sets ProductSubType field to given value.

### HasProductSubType

`func (o *PolicyProductResponse) HasProductSubType() bool`

HasProductSubType returns a boolean if a field has been set.

### GetProductType

`func (o *PolicyProductResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PolicyProductResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PolicyProductResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PolicyProductResponse) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetToken

`func (o *PolicyProductResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyProductResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyProductResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyProductResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *PolicyProductResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PolicyProductResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PolicyProductResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PolicyProductResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUsage

`func (o *PolicyProductResponse) GetUsage() []BalanceType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PolicyProductResponse) GetUsageOk() (*[]BalanceType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PolicyProductResponse) SetUsage(v []BalanceType)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PolicyProductResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


