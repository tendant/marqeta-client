# ProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProductTokens** | Pointer to **[]string** | One or more associated card product tokens. | [optional] 
**Classification** | Pointer to [**ProductClassification**](ProductClassification.md) |  | [optional] [default to PRODUCTCLASSIFICATION_CONSUMER]
**Config** | Pointer to [**ProductConfig**](ProductConfig.md) |  | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the credit product was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**CreditLine** | Pointer to [**ProductCreditLine**](ProductCreditLine.md) |  | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] [default to CURRENCYCODE_USD]
**Description** | Pointer to **string** | Description of the credit product. | [optional] 
**InterestCalculation** | Pointer to [**InterestCalculation**](InterestCalculation.md) |  | [optional] 
**MinPaymentCalculation** | Pointer to [**ProductMinPaymentCalculation**](ProductMinPaymentCalculation.md) |  | [optional] 
**MinPaymentFlatAmount** | Pointer to **float32** | Minimum payment, expressed as a flat amount, due on the payment due day. | [optional] 
**MinPaymentPercentage** | Pointer to **float32** | Minimum payment, expressed as a percentage of the total statement balance, due on the payment due day. | [optional] 
**Name** | Pointer to **string** | Name of the credit product. | [optional] 
**ParentProductToken** | Pointer to **string** | Unique identifier of the parent credit product. | [optional] 
**PaymentAllocationOrder** | Pointer to [**[]PaymentAllocationOrderEnum**](PaymentAllocationOrderEnum.md) | Ordered list of balance types to which payments are allocated, from first to last. | [optional] [default to ["INTEREST","FEES","PRINCIPAL"]]
**ProductSubType** | Pointer to [**ProductSubType**](ProductSubType.md) |  | [optional] [default to PRODUCTSUBTYPE_CREDIT_CARD]
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] [default to PRODUCTTYPE_REVOLVING]
**Status** | Pointer to [**ResourceStatus**](ResourceStatus.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the credit product. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the credit product was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Usage** | Pointer to [**[]BalanceType**](BalanceType.md) | One or more usage types for the credit product. | [optional] 

## Methods

### NewProductResponse

`func NewProductResponse() *ProductResponse`

NewProductResponse instantiates a new ProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductResponseWithDefaults

`func NewProductResponseWithDefaults() *ProductResponse`

NewProductResponseWithDefaults instantiates a new ProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProductTokens

`func (o *ProductResponse) GetCardProductTokens() []string`

GetCardProductTokens returns the CardProductTokens field if non-nil, zero value otherwise.

### GetCardProductTokensOk

`func (o *ProductResponse) GetCardProductTokensOk() (*[]string, bool)`

GetCardProductTokensOk returns a tuple with the CardProductTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductTokens

`func (o *ProductResponse) SetCardProductTokens(v []string)`

SetCardProductTokens sets CardProductTokens field to given value.

### HasCardProductTokens

`func (o *ProductResponse) HasCardProductTokens() bool`

HasCardProductTokens returns a boolean if a field has been set.

### GetClassification

`func (o *ProductResponse) GetClassification() ProductClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ProductResponse) GetClassificationOk() (*ProductClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ProductResponse) SetClassification(v ProductClassification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ProductResponse) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetConfig

`func (o *ProductResponse) GetConfig() ProductConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ProductResponse) GetConfigOk() (*ProductConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ProductResponse) SetConfig(v ProductConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ProductResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ProductResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ProductResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ProductResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ProductResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreditLine

`func (o *ProductResponse) GetCreditLine() ProductCreditLine`

GetCreditLine returns the CreditLine field if non-nil, zero value otherwise.

### GetCreditLineOk

`func (o *ProductResponse) GetCreditLineOk() (*ProductCreditLine, bool)`

GetCreditLineOk returns a tuple with the CreditLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLine

`func (o *ProductResponse) SetCreditLine(v ProductCreditLine)`

SetCreditLine sets CreditLine field to given value.

### HasCreditLine

`func (o *ProductResponse) HasCreditLine() bool`

HasCreditLine returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ProductResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProductResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProductResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ProductResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDescription

`func (o *ProductResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterestCalculation

`func (o *ProductResponse) GetInterestCalculation() InterestCalculation`

GetInterestCalculation returns the InterestCalculation field if non-nil, zero value otherwise.

### GetInterestCalculationOk

`func (o *ProductResponse) GetInterestCalculationOk() (*InterestCalculation, bool)`

GetInterestCalculationOk returns a tuple with the InterestCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCalculation

`func (o *ProductResponse) SetInterestCalculation(v InterestCalculation)`

SetInterestCalculation sets InterestCalculation field to given value.

### HasInterestCalculation

`func (o *ProductResponse) HasInterestCalculation() bool`

HasInterestCalculation returns a boolean if a field has been set.

### GetMinPaymentCalculation

`func (o *ProductResponse) GetMinPaymentCalculation() ProductMinPaymentCalculation`

GetMinPaymentCalculation returns the MinPaymentCalculation field if non-nil, zero value otherwise.

### GetMinPaymentCalculationOk

`func (o *ProductResponse) GetMinPaymentCalculationOk() (*ProductMinPaymentCalculation, bool)`

GetMinPaymentCalculationOk returns a tuple with the MinPaymentCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentCalculation

`func (o *ProductResponse) SetMinPaymentCalculation(v ProductMinPaymentCalculation)`

SetMinPaymentCalculation sets MinPaymentCalculation field to given value.

### HasMinPaymentCalculation

`func (o *ProductResponse) HasMinPaymentCalculation() bool`

HasMinPaymentCalculation returns a boolean if a field has been set.

### GetMinPaymentFlatAmount

`func (o *ProductResponse) GetMinPaymentFlatAmount() float32`

GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field if non-nil, zero value otherwise.

### GetMinPaymentFlatAmountOk

`func (o *ProductResponse) GetMinPaymentFlatAmountOk() (*float32, bool)`

GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentFlatAmount

`func (o *ProductResponse) SetMinPaymentFlatAmount(v float32)`

SetMinPaymentFlatAmount sets MinPaymentFlatAmount field to given value.

### HasMinPaymentFlatAmount

`func (o *ProductResponse) HasMinPaymentFlatAmount() bool`

HasMinPaymentFlatAmount returns a boolean if a field has been set.

### GetMinPaymentPercentage

`func (o *ProductResponse) GetMinPaymentPercentage() float32`

GetMinPaymentPercentage returns the MinPaymentPercentage field if non-nil, zero value otherwise.

### GetMinPaymentPercentageOk

`func (o *ProductResponse) GetMinPaymentPercentageOk() (*float32, bool)`

GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentPercentage

`func (o *ProductResponse) SetMinPaymentPercentage(v float32)`

SetMinPaymentPercentage sets MinPaymentPercentage field to given value.

### HasMinPaymentPercentage

`func (o *ProductResponse) HasMinPaymentPercentage() bool`

HasMinPaymentPercentage returns a boolean if a field has been set.

### GetName

`func (o *ProductResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentProductToken

`func (o *ProductResponse) GetParentProductToken() string`

GetParentProductToken returns the ParentProductToken field if non-nil, zero value otherwise.

### GetParentProductTokenOk

`func (o *ProductResponse) GetParentProductTokenOk() (*string, bool)`

GetParentProductTokenOk returns a tuple with the ParentProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProductToken

`func (o *ProductResponse) SetParentProductToken(v string)`

SetParentProductToken sets ParentProductToken field to given value.

### HasParentProductToken

`func (o *ProductResponse) HasParentProductToken() bool`

HasParentProductToken returns a boolean if a field has been set.

### GetPaymentAllocationOrder

`func (o *ProductResponse) GetPaymentAllocationOrder() []PaymentAllocationOrderEnum`

GetPaymentAllocationOrder returns the PaymentAllocationOrder field if non-nil, zero value otherwise.

### GetPaymentAllocationOrderOk

`func (o *ProductResponse) GetPaymentAllocationOrderOk() (*[]PaymentAllocationOrderEnum, bool)`

GetPaymentAllocationOrderOk returns a tuple with the PaymentAllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAllocationOrder

`func (o *ProductResponse) SetPaymentAllocationOrder(v []PaymentAllocationOrderEnum)`

SetPaymentAllocationOrder sets PaymentAllocationOrder field to given value.

### HasPaymentAllocationOrder

`func (o *ProductResponse) HasPaymentAllocationOrder() bool`

HasPaymentAllocationOrder returns a boolean if a field has been set.

### GetProductSubType

`func (o *ProductResponse) GetProductSubType() ProductSubType`

GetProductSubType returns the ProductSubType field if non-nil, zero value otherwise.

### GetProductSubTypeOk

`func (o *ProductResponse) GetProductSubTypeOk() (*ProductSubType, bool)`

GetProductSubTypeOk returns a tuple with the ProductSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubType

`func (o *ProductResponse) SetProductSubType(v ProductSubType)`

SetProductSubType sets ProductSubType field to given value.

### HasProductSubType

`func (o *ProductResponse) HasProductSubType() bool`

HasProductSubType returns a boolean if a field has been set.

### GetProductType

`func (o *ProductResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ProductResponse) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetStatus

`func (o *ProductResponse) GetStatus() ResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductResponse) GetStatusOk() (*ResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductResponse) SetStatus(v ResourceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ProductResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProductResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProductResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProductResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *ProductResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ProductResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ProductResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ProductResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUsage

`func (o *ProductResponse) GetUsage() []BalanceType`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ProductResponse) GetUsageOk() (*[]BalanceType, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ProductResponse) SetUsage(v []BalanceType)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ProductResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


