# StatementPaymentWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclosure** | Pointer to **string** | Statement disclosure in the case of negative or no amortization, or no lifetime repayment for the minimum payment warning type.  * &#x60;NEGATIVE_OR_NO_AMORTIZATION&#x60; - Occurs when the interest amount is higher than the minimum payment; results in the outstanding balance remaining in perpetuity. * &#x60;NO_LIFETIME_REPAYMENT&#x60; - Occurs when the interest amount is just below the minimum payment; results in the outstanding balance taking longer than a lifetime to pay off. | [optional] 
**InterestPaid** | Pointer to **float32** | For the minimum payment warning type, this value represents the total amount of interest to pay off the statement balance if only making the minimum payment each month.  For the 3 Year warning type, this value represents the total amount of interest if paying off the statement balance in three years. | [optional] 
**MonthlyPayment** | Pointer to **float32** | For the minimum payment warning type, this value is 0.  For the 3 Year warning type, this value represents the fixed monthly payment amount required to pay off the statement balance in three years. | [optional] 
**PayOffPeriod** | Pointer to **int32** | For the minimum payment warning type, this value represents the number of periods required to pay off the statement balance.  For the 3 Year warning type, this value is 36 (months). | [optional] 
**PeriodType** | Pointer to **string** | Time unit of the pay off period. | [optional] 
**TotalPaid** | Pointer to **float32** | For the minimum payment warning type, this value represents the total amount of principal and interest to pay off the statement balance if only making the minimum payment each month.  For the 3 Year warning type, this value represents the total amount of principal and interest if paying off the statement balance in three years. | [optional] 
**Type** | Pointer to **string** | Type of statement warning.  * &#x60;MIN_PAYMENT&#x60; - Displays the total estimated payment amount and how long it would take to pay off the statement balance if only making minimum payments. * &#x60;3_YEAR&#x60; - Displays the monthly payment amount and total estimated payment amount if paying off the statement balance in three years. | [optional] 

## Methods

### NewStatementPaymentWarning

`func NewStatementPaymentWarning() *StatementPaymentWarning`

NewStatementPaymentWarning instantiates a new StatementPaymentWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementPaymentWarningWithDefaults

`func NewStatementPaymentWarningWithDefaults() *StatementPaymentWarning`

NewStatementPaymentWarningWithDefaults instantiates a new StatementPaymentWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclosure

`func (o *StatementPaymentWarning) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *StatementPaymentWarning) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *StatementPaymentWarning) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *StatementPaymentWarning) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetInterestPaid

`func (o *StatementPaymentWarning) GetInterestPaid() float32`

GetInterestPaid returns the InterestPaid field if non-nil, zero value otherwise.

### GetInterestPaidOk

`func (o *StatementPaymentWarning) GetInterestPaidOk() (*float32, bool)`

GetInterestPaidOk returns a tuple with the InterestPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPaid

`func (o *StatementPaymentWarning) SetInterestPaid(v float32)`

SetInterestPaid sets InterestPaid field to given value.

### HasInterestPaid

`func (o *StatementPaymentWarning) HasInterestPaid() bool`

HasInterestPaid returns a boolean if a field has been set.

### GetMonthlyPayment

`func (o *StatementPaymentWarning) GetMonthlyPayment() float32`

GetMonthlyPayment returns the MonthlyPayment field if non-nil, zero value otherwise.

### GetMonthlyPaymentOk

`func (o *StatementPaymentWarning) GetMonthlyPaymentOk() (*float32, bool)`

GetMonthlyPaymentOk returns a tuple with the MonthlyPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPayment

`func (o *StatementPaymentWarning) SetMonthlyPayment(v float32)`

SetMonthlyPayment sets MonthlyPayment field to given value.

### HasMonthlyPayment

`func (o *StatementPaymentWarning) HasMonthlyPayment() bool`

HasMonthlyPayment returns a boolean if a field has been set.

### GetPayOffPeriod

`func (o *StatementPaymentWarning) GetPayOffPeriod() int32`

GetPayOffPeriod returns the PayOffPeriod field if non-nil, zero value otherwise.

### GetPayOffPeriodOk

`func (o *StatementPaymentWarning) GetPayOffPeriodOk() (*int32, bool)`

GetPayOffPeriodOk returns a tuple with the PayOffPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOffPeriod

`func (o *StatementPaymentWarning) SetPayOffPeriod(v int32)`

SetPayOffPeriod sets PayOffPeriod field to given value.

### HasPayOffPeriod

`func (o *StatementPaymentWarning) HasPayOffPeriod() bool`

HasPayOffPeriod returns a boolean if a field has been set.

### GetPeriodType

`func (o *StatementPaymentWarning) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *StatementPaymentWarning) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *StatementPaymentWarning) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *StatementPaymentWarning) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### GetTotalPaid

`func (o *StatementPaymentWarning) GetTotalPaid() float32`

GetTotalPaid returns the TotalPaid field if non-nil, zero value otherwise.

### GetTotalPaidOk

`func (o *StatementPaymentWarning) GetTotalPaidOk() (*float32, bool)`

GetTotalPaidOk returns a tuple with the TotalPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaid

`func (o *StatementPaymentWarning) SetTotalPaid(v float32)`

SetTotalPaid sets TotalPaid field to given value.

### HasTotalPaid

`func (o *StatementPaymentWarning) HasTotalPaid() bool`

HasTotalPaid returns a boolean if a field has been set.

### GetType

`func (o *StatementPaymentWarning) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatementPaymentWarning) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatementPaymentWarning) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StatementPaymentWarning) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


