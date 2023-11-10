# ProductCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProductTokens** | **[]string** | One or more associated card product tokens. | 
**Classification** | [**ProductClassification**](ProductClassification.md) |  | [default to PRODUCTCLASSIFICATION_CONSUMER]
**Config** | [**ProductConfig**](ProductConfig.md) |  | 
**CreditLine** | [**ProductCreditLine**](ProductCreditLine.md) |  | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Description** | Pointer to **string** | Description of the credit product. | [optional] 
**InterestCalculation** | [**InterestCalculation**](InterestCalculation.md) |  | 
**MinPaymentCalculation** | Pointer to [**ProductMinPaymentCalculation**](ProductMinPaymentCalculation.md) |  | [optional] 
**MinPaymentFlatAmount** | **float32** | Minimum payment, expressed as a flat amount, due on the payment due day. | 
**MinPaymentPercentage** | **float32** | Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day. | 
**Name** | **string** | Name of the credit product. | 
**PaymentAllocationOrder** | [**[]PaymentAllocationOrderEnum**](PaymentAllocationOrderEnum.md) | Ordered list of balance types to which payments are allocated, from first to last. | [default to ["INTEREST","FEES","PRINCIPAL"]]
**ProductSubType** | [**ProductSubType**](ProductSubType.md) |  | [default to PRODUCTSUBTYPE_CREDIT_CARD]
**ProductType** | [**ProductType**](ProductType.md) |  | [default to PRODUCTTYPE_REVOLVING]
**Status** | Pointer to [**ResourceStatus**](ResourceStatus.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the credit product. | [optional] 
**Usage** | [**[]BalanceType**](BalanceType.md) | One or more usage types for the credit product. | 

## Methods

### NewProductCreateReq

`func NewProductCreateReq(cardProductTokens []string, classification ProductClassification, config ProductConfig, creditLine ProductCreditLine, currencyCode CurrencyCode, interestCalculation InterestCalculation, minPaymentFlatAmount float32, minPaymentPercentage float32, name string, paymentAllocationOrder []PaymentAllocationOrderEnum, productSubType ProductSubType, productType ProductType, usage []BalanceType, ) *ProductCreateReq`

NewProductCreateReq instantiates a new ProductCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCreateReqWithDefaults

`func NewProductCreateReqWithDefaults() *ProductCreateReq`

NewProductCreateReqWithDefaults instantiates a new ProductCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProductTokens

`func (o *ProductCreateReq) GetCardProductTokens() []string`

GetCardProductTokens returns the CardProductTokens field if non-nil, zero value otherwise.

### GetCardProductTokensOk

`func (o *ProductCreateReq) GetCardProductTokensOk() (*[]string, bool)`

GetCardProductTokensOk returns a tuple with the CardProductTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductTokens

`func (o *ProductCreateReq) SetCardProductTokens(v []string)`

SetCardProductTokens sets CardProductTokens field to given value.


### GetClassification

`func (o *ProductCreateReq) GetClassification() ProductClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ProductCreateReq) GetClassificationOk() (*ProductClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ProductCreateReq) SetClassification(v ProductClassification)`

SetClassification sets Classification field to given value.


### GetConfig

`func (o *ProductCreateReq) GetConfig() ProductConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ProductCreateReq) GetConfigOk() (*ProductConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ProductCreateReq) SetConfig(v ProductConfig)`

SetConfig sets Config field to given value.


### GetCreditLine

`func (o *ProductCreateReq) GetCreditLine() ProductCreditLine`

GetCreditLine returns the CreditLine field if non-nil, zero value otherwise.

### GetCreditLineOk

`func (o *ProductCreateReq) GetCreditLineOk() (*ProductCreditLine, bool)`

GetCreditLineOk returns a tuple with the CreditLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLine

`func (o *ProductCreateReq) SetCreditLine(v ProductCreditLine)`

SetCreditLine sets CreditLine field to given value.


### GetCurrencyCode

`func (o *ProductCreateReq) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProductCreateReq) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProductCreateReq) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetDescription

`func (o *ProductCreateReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductCreateReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductCreateReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductCreateReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterestCalculation

`func (o *ProductCreateReq) GetInterestCalculation() InterestCalculation`

GetInterestCalculation returns the InterestCalculation field if non-nil, zero value otherwise.

### GetInterestCalculationOk

`func (o *ProductCreateReq) GetInterestCalculationOk() (*InterestCalculation, bool)`

GetInterestCalculationOk returns a tuple with the InterestCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCalculation

`func (o *ProductCreateReq) SetInterestCalculation(v InterestCalculation)`

SetInterestCalculation sets InterestCalculation field to given value.


### GetMinPaymentCalculation

`func (o *ProductCreateReq) GetMinPaymentCalculation() ProductMinPaymentCalculation`

GetMinPaymentCalculation returns the MinPaymentCalculation field if non-nil, zero value otherwise.

### GetMinPaymentCalculationOk

`func (o *ProductCreateReq) GetMinPaymentCalculationOk() (*ProductMinPaymentCalculation, bool)`

GetMinPaymentCalculationOk returns a tuple with the MinPaymentCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentCalculation

`func (o *ProductCreateReq) SetMinPaymentCalculation(v ProductMinPaymentCalculation)`

SetMinPaymentCalculation sets MinPaymentCalculation field to given value.

### HasMinPaymentCalculation

`func (o *ProductCreateReq) HasMinPaymentCalculation() bool`

HasMinPaymentCalculation returns a boolean if a field has been set.

### GetMinPaymentFlatAmount

`func (o *ProductCreateReq) GetMinPaymentFlatAmount() float32`

GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field if non-nil, zero value otherwise.

### GetMinPaymentFlatAmountOk

`func (o *ProductCreateReq) GetMinPaymentFlatAmountOk() (*float32, bool)`

GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentFlatAmount

`func (o *ProductCreateReq) SetMinPaymentFlatAmount(v float32)`

SetMinPaymentFlatAmount sets MinPaymentFlatAmount field to given value.


### GetMinPaymentPercentage

`func (o *ProductCreateReq) GetMinPaymentPercentage() float32`

GetMinPaymentPercentage returns the MinPaymentPercentage field if non-nil, zero value otherwise.

### GetMinPaymentPercentageOk

`func (o *ProductCreateReq) GetMinPaymentPercentageOk() (*float32, bool)`

GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentPercentage

`func (o *ProductCreateReq) SetMinPaymentPercentage(v float32)`

SetMinPaymentPercentage sets MinPaymentPercentage field to given value.


### GetName

`func (o *ProductCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentAllocationOrder

`func (o *ProductCreateReq) GetPaymentAllocationOrder() []PaymentAllocationOrderEnum`

GetPaymentAllocationOrder returns the PaymentAllocationOrder field if non-nil, zero value otherwise.

### GetPaymentAllocationOrderOk

`func (o *ProductCreateReq) GetPaymentAllocationOrderOk() (*[]PaymentAllocationOrderEnum, bool)`

GetPaymentAllocationOrderOk returns a tuple with the PaymentAllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAllocationOrder

`func (o *ProductCreateReq) SetPaymentAllocationOrder(v []PaymentAllocationOrderEnum)`

SetPaymentAllocationOrder sets PaymentAllocationOrder field to given value.


### GetProductSubType

`func (o *ProductCreateReq) GetProductSubType() ProductSubType`

GetProductSubType returns the ProductSubType field if non-nil, zero value otherwise.

### GetProductSubTypeOk

`func (o *ProductCreateReq) GetProductSubTypeOk() (*ProductSubType, bool)`

GetProductSubTypeOk returns a tuple with the ProductSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubType

`func (o *ProductCreateReq) SetProductSubType(v ProductSubType)`

SetProductSubType sets ProductSubType field to given value.


### GetProductType

`func (o *ProductCreateReq) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductCreateReq) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductCreateReq) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetStatus

`func (o *ProductCreateReq) GetStatus() ResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductCreateReq) GetStatusOk() (*ResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductCreateReq) SetStatus(v ResourceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductCreateReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ProductCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProductCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProductCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProductCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsage

`func (o *ProductCreateReq) GetUsage() []BalanceType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ProductCreateReq) GetUsageOk() (*[]BalanceType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ProductCreateReq) SetUsage(v []BalanceType)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


