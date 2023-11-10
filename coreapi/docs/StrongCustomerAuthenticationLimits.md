# StrongCustomerAuthenticationLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CavvAuthenticationAmountIncrementalPercentage** | Pointer to **string** |  | [optional] [readonly] 
**EnableCavvAuthenticationAmountValidation** | Pointer to **bool** |  | [optional] [readonly] 
**ScaContactlessCumulativeAmountLimit** | Pointer to **float32** | Specifies the cumulative limit of transactions the cardholder can perform before receiving an SCA challenge.  A value of &#x60;0&#x60; in this field means that the cumulative amount spent in contactless POS transactions performed by the cardholder does not impact the decision of whether or not an SCA challenge is served. | [optional] [readonly] 
**ScaContactlessTransactionLimit** | Pointer to **float32** | Specifies the maximum allowable amount for a single contactless point-of-sale (POS) transaction, above which the cardholder receives a strong customer authentication (SCA) challenge.  A value of &#x60;0&#x60; in this field means that the amount of any single contactless POS transaction performed by the cardholder does not impact the decision of whether or not an SCA challenge is served. | [optional] [readonly] 
**ScaContactlessTransactionsCountLimit** | Pointer to **int32** | Specifies the number of contactless POS transactions the cardholder can perform before receiving an SCA challenge.  A value of &#x60;0&#x60; in this field means that the number of contactless POS transactions performed by the cardholder does not impact the decision of whether or not an SCA challenge is served. | [optional] [readonly] 
**ScaContactlessTransactionsCurrency** | Pointer to **string** | Specifies the currency type for contactless POS transactions.  This field is required if either the &#x60;sca_contactless_transaction_limit&#x60; field or the &#x60;sca_contactless_cumulative_amount_limit&#x60; field in this object contains a value, even if that value is &#x60;0&#x60;. | [optional] [readonly] 
**ScaLvpCumulativeAmountLimit** | Pointer to **float32** | Specifies the cumulative limit of LVP e-commerce transactions the cardholder can perform before receiving an SCA challenge.  For example, if you set the value of this field to &#x60;100.00&#x60;, your cardholder can perform two transactions for &#x60;30.00&#x60; and two transactions for &#x60;20.00&#x60; before receiving an SCA challenge.  If you leave this field blank, the cumulative amount spent in LVP e-commerce transactions performed by the cardholder does not impact the decision of whether or not an SCA challenge is served. | [optional] [readonly] 
**ScaLvpTransactionLimit** | Pointer to **float32** | Specifies the maximum allowable amount for a single low-value payment (LVP) e-commerce transaction, above which the cardholder receives a strong customer authentication (SCA) challenge.  If you leave this field blank, the amount of any single LVP e-commerce transaction performed by the cardholder does not impact the decision of whether or not an SCA challenge is served. | [optional] [readonly] 
**ScaLvpTransactionsCountLimit** | Pointer to **int32** | Specifies the number of LVP e-commerce transactions the cardholder can perform before receiving an SCA challenge.  If you leave this field blank, the total number of LVP e-commerce transactions performed by the cardholder does not impact the decision of whether or not an SCA challenge is served. | [optional] [readonly] 
**ScaLvpTransactionsCurrency** | Pointer to **string** | Specifies the currency type for LVP e-commerce transaction limits.  This field is required if any one of the &#x60;sca_lvp_transaction_limit&#x60;, &#x60;sca_lvp_cumulative_amount_limit&#x60;, or &#x60;sca_lvp_transactions_count_limit&#x60; fields in this object contains a value, even if that value is &#x60;0&#x60;. | [optional] [readonly] 
**ScaTraExemptionAmountLimit** | Pointer to **float32** | Specifies the maximum allowable amount for a single low-value payment (LVP) e-commerce transaction with transaction risk analysis (TRA) exemption sent by the merchant or acquirer. If the transaction amount exceeds the specified limit, then the transaction is either approved or it receives a strong customer authentication (SCA) challenge based on &#x60;sca_lvp_transaction_limit&#x60; and &#x60;sca_lvp_transactions_currency&#x60;. | [optional] [readonly] 

## Methods

### NewStrongCustomerAuthenticationLimits

`func NewStrongCustomerAuthenticationLimits() *StrongCustomerAuthenticationLimits`

NewStrongCustomerAuthenticationLimits instantiates a new StrongCustomerAuthenticationLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrongCustomerAuthenticationLimitsWithDefaults

`func NewStrongCustomerAuthenticationLimitsWithDefaults() *StrongCustomerAuthenticationLimits`

NewStrongCustomerAuthenticationLimitsWithDefaults instantiates a new StrongCustomerAuthenticationLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCavvAuthenticationAmountIncrementalPercentage

`func (o *StrongCustomerAuthenticationLimits) GetCavvAuthenticationAmountIncrementalPercentage() string`

GetCavvAuthenticationAmountIncrementalPercentage returns the CavvAuthenticationAmountIncrementalPercentage field if non-nil, zero value otherwise.

### GetCavvAuthenticationAmountIncrementalPercentageOk

`func (o *StrongCustomerAuthenticationLimits) GetCavvAuthenticationAmountIncrementalPercentageOk() (*string, bool)`

GetCavvAuthenticationAmountIncrementalPercentageOk returns a tuple with the CavvAuthenticationAmountIncrementalPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAuthenticationAmountIncrementalPercentage

`func (o *StrongCustomerAuthenticationLimits) SetCavvAuthenticationAmountIncrementalPercentage(v string)`

SetCavvAuthenticationAmountIncrementalPercentage sets CavvAuthenticationAmountIncrementalPercentage field to given value.

### HasCavvAuthenticationAmountIncrementalPercentage

`func (o *StrongCustomerAuthenticationLimits) HasCavvAuthenticationAmountIncrementalPercentage() bool`

HasCavvAuthenticationAmountIncrementalPercentage returns a boolean if a field has been set.

### GetEnableCavvAuthenticationAmountValidation

`func (o *StrongCustomerAuthenticationLimits) GetEnableCavvAuthenticationAmountValidation() bool`

GetEnableCavvAuthenticationAmountValidation returns the EnableCavvAuthenticationAmountValidation field if non-nil, zero value otherwise.

### GetEnableCavvAuthenticationAmountValidationOk

`func (o *StrongCustomerAuthenticationLimits) GetEnableCavvAuthenticationAmountValidationOk() (*bool, bool)`

GetEnableCavvAuthenticationAmountValidationOk returns a tuple with the EnableCavvAuthenticationAmountValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCavvAuthenticationAmountValidation

`func (o *StrongCustomerAuthenticationLimits) SetEnableCavvAuthenticationAmountValidation(v bool)`

SetEnableCavvAuthenticationAmountValidation sets EnableCavvAuthenticationAmountValidation field to given value.

### HasEnableCavvAuthenticationAmountValidation

`func (o *StrongCustomerAuthenticationLimits) HasEnableCavvAuthenticationAmountValidation() bool`

HasEnableCavvAuthenticationAmountValidation returns a boolean if a field has been set.

### GetScaContactlessCumulativeAmountLimit

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessCumulativeAmountLimit() float32`

GetScaContactlessCumulativeAmountLimit returns the ScaContactlessCumulativeAmountLimit field if non-nil, zero value otherwise.

### GetScaContactlessCumulativeAmountLimitOk

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessCumulativeAmountLimitOk() (*float32, bool)`

GetScaContactlessCumulativeAmountLimitOk returns a tuple with the ScaContactlessCumulativeAmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaContactlessCumulativeAmountLimit

`func (o *StrongCustomerAuthenticationLimits) SetScaContactlessCumulativeAmountLimit(v float32)`

SetScaContactlessCumulativeAmountLimit sets ScaContactlessCumulativeAmountLimit field to given value.

### HasScaContactlessCumulativeAmountLimit

`func (o *StrongCustomerAuthenticationLimits) HasScaContactlessCumulativeAmountLimit() bool`

HasScaContactlessCumulativeAmountLimit returns a boolean if a field has been set.

### GetScaContactlessTransactionLimit

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessTransactionLimit() float32`

GetScaContactlessTransactionLimit returns the ScaContactlessTransactionLimit field if non-nil, zero value otherwise.

### GetScaContactlessTransactionLimitOk

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessTransactionLimitOk() (*float32, bool)`

GetScaContactlessTransactionLimitOk returns a tuple with the ScaContactlessTransactionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaContactlessTransactionLimit

`func (o *StrongCustomerAuthenticationLimits) SetScaContactlessTransactionLimit(v float32)`

SetScaContactlessTransactionLimit sets ScaContactlessTransactionLimit field to given value.

### HasScaContactlessTransactionLimit

`func (o *StrongCustomerAuthenticationLimits) HasScaContactlessTransactionLimit() bool`

HasScaContactlessTransactionLimit returns a boolean if a field has been set.

### GetScaContactlessTransactionsCountLimit

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessTransactionsCountLimit() int32`

GetScaContactlessTransactionsCountLimit returns the ScaContactlessTransactionsCountLimit field if non-nil, zero value otherwise.

### GetScaContactlessTransactionsCountLimitOk

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessTransactionsCountLimitOk() (*int32, bool)`

GetScaContactlessTransactionsCountLimitOk returns a tuple with the ScaContactlessTransactionsCountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaContactlessTransactionsCountLimit

`func (o *StrongCustomerAuthenticationLimits) SetScaContactlessTransactionsCountLimit(v int32)`

SetScaContactlessTransactionsCountLimit sets ScaContactlessTransactionsCountLimit field to given value.

### HasScaContactlessTransactionsCountLimit

`func (o *StrongCustomerAuthenticationLimits) HasScaContactlessTransactionsCountLimit() bool`

HasScaContactlessTransactionsCountLimit returns a boolean if a field has been set.

### GetScaContactlessTransactionsCurrency

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessTransactionsCurrency() string`

GetScaContactlessTransactionsCurrency returns the ScaContactlessTransactionsCurrency field if non-nil, zero value otherwise.

### GetScaContactlessTransactionsCurrencyOk

`func (o *StrongCustomerAuthenticationLimits) GetScaContactlessTransactionsCurrencyOk() (*string, bool)`

GetScaContactlessTransactionsCurrencyOk returns a tuple with the ScaContactlessTransactionsCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaContactlessTransactionsCurrency

`func (o *StrongCustomerAuthenticationLimits) SetScaContactlessTransactionsCurrency(v string)`

SetScaContactlessTransactionsCurrency sets ScaContactlessTransactionsCurrency field to given value.

### HasScaContactlessTransactionsCurrency

`func (o *StrongCustomerAuthenticationLimits) HasScaContactlessTransactionsCurrency() bool`

HasScaContactlessTransactionsCurrency returns a boolean if a field has been set.

### GetScaLvpCumulativeAmountLimit

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpCumulativeAmountLimit() float32`

GetScaLvpCumulativeAmountLimit returns the ScaLvpCumulativeAmountLimit field if non-nil, zero value otherwise.

### GetScaLvpCumulativeAmountLimitOk

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpCumulativeAmountLimitOk() (*float32, bool)`

GetScaLvpCumulativeAmountLimitOk returns a tuple with the ScaLvpCumulativeAmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaLvpCumulativeAmountLimit

`func (o *StrongCustomerAuthenticationLimits) SetScaLvpCumulativeAmountLimit(v float32)`

SetScaLvpCumulativeAmountLimit sets ScaLvpCumulativeAmountLimit field to given value.

### HasScaLvpCumulativeAmountLimit

`func (o *StrongCustomerAuthenticationLimits) HasScaLvpCumulativeAmountLimit() bool`

HasScaLvpCumulativeAmountLimit returns a boolean if a field has been set.

### GetScaLvpTransactionLimit

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpTransactionLimit() float32`

GetScaLvpTransactionLimit returns the ScaLvpTransactionLimit field if non-nil, zero value otherwise.

### GetScaLvpTransactionLimitOk

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpTransactionLimitOk() (*float32, bool)`

GetScaLvpTransactionLimitOk returns a tuple with the ScaLvpTransactionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaLvpTransactionLimit

`func (o *StrongCustomerAuthenticationLimits) SetScaLvpTransactionLimit(v float32)`

SetScaLvpTransactionLimit sets ScaLvpTransactionLimit field to given value.

### HasScaLvpTransactionLimit

`func (o *StrongCustomerAuthenticationLimits) HasScaLvpTransactionLimit() bool`

HasScaLvpTransactionLimit returns a boolean if a field has been set.

### GetScaLvpTransactionsCountLimit

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpTransactionsCountLimit() int32`

GetScaLvpTransactionsCountLimit returns the ScaLvpTransactionsCountLimit field if non-nil, zero value otherwise.

### GetScaLvpTransactionsCountLimitOk

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpTransactionsCountLimitOk() (*int32, bool)`

GetScaLvpTransactionsCountLimitOk returns a tuple with the ScaLvpTransactionsCountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaLvpTransactionsCountLimit

`func (o *StrongCustomerAuthenticationLimits) SetScaLvpTransactionsCountLimit(v int32)`

SetScaLvpTransactionsCountLimit sets ScaLvpTransactionsCountLimit field to given value.

### HasScaLvpTransactionsCountLimit

`func (o *StrongCustomerAuthenticationLimits) HasScaLvpTransactionsCountLimit() bool`

HasScaLvpTransactionsCountLimit returns a boolean if a field has been set.

### GetScaLvpTransactionsCurrency

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpTransactionsCurrency() string`

GetScaLvpTransactionsCurrency returns the ScaLvpTransactionsCurrency field if non-nil, zero value otherwise.

### GetScaLvpTransactionsCurrencyOk

`func (o *StrongCustomerAuthenticationLimits) GetScaLvpTransactionsCurrencyOk() (*string, bool)`

GetScaLvpTransactionsCurrencyOk returns a tuple with the ScaLvpTransactionsCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaLvpTransactionsCurrency

`func (o *StrongCustomerAuthenticationLimits) SetScaLvpTransactionsCurrency(v string)`

SetScaLvpTransactionsCurrency sets ScaLvpTransactionsCurrency field to given value.

### HasScaLvpTransactionsCurrency

`func (o *StrongCustomerAuthenticationLimits) HasScaLvpTransactionsCurrency() bool`

HasScaLvpTransactionsCurrency returns a boolean if a field has been set.

### GetScaTraExemptionAmountLimit

`func (o *StrongCustomerAuthenticationLimits) GetScaTraExemptionAmountLimit() float32`

GetScaTraExemptionAmountLimit returns the ScaTraExemptionAmountLimit field if non-nil, zero value otherwise.

### GetScaTraExemptionAmountLimitOk

`func (o *StrongCustomerAuthenticationLimits) GetScaTraExemptionAmountLimitOk() (*float32, bool)`

GetScaTraExemptionAmountLimitOk returns a tuple with the ScaTraExemptionAmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaTraExemptionAmountLimit

`func (o *StrongCustomerAuthenticationLimits) SetScaTraExemptionAmountLimit(v float32)`

SetScaTraExemptionAmountLimit sets ScaTraExemptionAmountLimit field to given value.

### HasScaTraExemptionAmountLimit

`func (o *StrongCustomerAuthenticationLimits) HasScaTraExemptionAmountLimit() bool`

HasScaTraExemptionAmountLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


