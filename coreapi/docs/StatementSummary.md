# StatementSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the statement summary is generated. | 
**AvailableCredit** | **float32** | Amount available to spend on the credit account, as of the statement closing date. | 
**ClosingBalance** | **float32** | Balance of the credit account when the statement period ended. | 
**ClosingDate** | **time.Time** | Date and time when the statement period ended. | 
**CreatedTime** | **time.Time** | Date and time when the statement summary was created on Marqeta&#39;s credit platform, in UTC. | 
**CreditLimit** | Pointer to **float32** | Maximum balance the credit account can carry, as of the statement closing date. | [optional] 
**Credits** | **float32** | Total amount of credits received during the statement period. | 
**CycleType** | [**CycleType**](CycleType.md) |  | 
**DaysInBillingCycle** | **int32** | Number of days in the billing cycle, also known as the statement period. | 
**Fees** | **float32** | Total amount of fees charged during the statement period. Does not include periodic fees. | 
**Interest** | **float32** | Total amount of interest charged during the statement period. | 
**OpeningBalance** | **float32** | Balance of the credit account when the statement period began. | 
**OpeningDate** | **time.Time** | Date and time when the statement period began. | 
**PastDueAmount** | **float32** | Total payment amount, required to make the account current. | 
**Payments** | **float32** | Total amount of payments made during the statement period. | 
**Purchases** | **float32** | Total amount of purchases made during the statement period. | 
**Token** | **string** | Unique identifier of the statement summary. | 

## Methods

### NewStatementSummary

`func NewStatementSummary(accountToken string, availableCredit float32, closingBalance float32, closingDate time.Time, createdTime time.Time, credits float32, cycleType CycleType, daysInBillingCycle int32, fees float32, interest float32, openingBalance float32, openingDate time.Time, pastDueAmount float32, payments float32, purchases float32, token string, ) *StatementSummary`

NewStatementSummary instantiates a new StatementSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementSummaryWithDefaults

`func NewStatementSummaryWithDefaults() *StatementSummary`

NewStatementSummaryWithDefaults instantiates a new StatementSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *StatementSummary) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *StatementSummary) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *StatementSummary) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAvailableCredit

`func (o *StatementSummary) GetAvailableCredit() float32`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *StatementSummary) GetAvailableCreditOk() (*float32, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *StatementSummary) SetAvailableCredit(v float32)`

SetAvailableCredit sets AvailableCredit field to given value.


### GetClosingBalance

`func (o *StatementSummary) GetClosingBalance() float32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *StatementSummary) GetClosingBalanceOk() (*float32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *StatementSummary) SetClosingBalance(v float32)`

SetClosingBalance sets ClosingBalance field to given value.


### GetClosingDate

`func (o *StatementSummary) GetClosingDate() time.Time`

GetClosingDate returns the ClosingDate field if non-nil, zero value otherwise.

### GetClosingDateOk

`func (o *StatementSummary) GetClosingDateOk() (*time.Time, bool)`

GetClosingDateOk returns a tuple with the ClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDate

`func (o *StatementSummary) SetClosingDate(v time.Time)`

SetClosingDate sets ClosingDate field to given value.


### GetCreatedTime

`func (o *StatementSummary) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StatementSummary) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StatementSummary) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreditLimit

`func (o *StatementSummary) GetCreditLimit() float32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *StatementSummary) GetCreditLimitOk() (*float32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *StatementSummary) SetCreditLimit(v float32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *StatementSummary) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCredits

`func (o *StatementSummary) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *StatementSummary) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *StatementSummary) SetCredits(v float32)`

SetCredits sets Credits field to given value.


### GetCycleType

`func (o *StatementSummary) GetCycleType() CycleType`

GetCycleType returns the CycleType field if non-nil, zero value otherwise.

### GetCycleTypeOk

`func (o *StatementSummary) GetCycleTypeOk() (*CycleType, bool)`

GetCycleTypeOk returns a tuple with the CycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleType

`func (o *StatementSummary) SetCycleType(v CycleType)`

SetCycleType sets CycleType field to given value.


### GetDaysInBillingCycle

`func (o *StatementSummary) GetDaysInBillingCycle() int32`

GetDaysInBillingCycle returns the DaysInBillingCycle field if non-nil, zero value otherwise.

### GetDaysInBillingCycleOk

`func (o *StatementSummary) GetDaysInBillingCycleOk() (*int32, bool)`

GetDaysInBillingCycleOk returns a tuple with the DaysInBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInBillingCycle

`func (o *StatementSummary) SetDaysInBillingCycle(v int32)`

SetDaysInBillingCycle sets DaysInBillingCycle field to given value.


### GetFees

`func (o *StatementSummary) GetFees() float32`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *StatementSummary) GetFeesOk() (*float32, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *StatementSummary) SetFees(v float32)`

SetFees sets Fees field to given value.


### GetInterest

`func (o *StatementSummary) GetInterest() float32`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *StatementSummary) GetInterestOk() (*float32, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *StatementSummary) SetInterest(v float32)`

SetInterest sets Interest field to given value.


### GetOpeningBalance

`func (o *StatementSummary) GetOpeningBalance() float32`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *StatementSummary) GetOpeningBalanceOk() (*float32, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *StatementSummary) SetOpeningBalance(v float32)`

SetOpeningBalance sets OpeningBalance field to given value.


### GetOpeningDate

`func (o *StatementSummary) GetOpeningDate() time.Time`

GetOpeningDate returns the OpeningDate field if non-nil, zero value otherwise.

### GetOpeningDateOk

`func (o *StatementSummary) GetOpeningDateOk() (*time.Time, bool)`

GetOpeningDateOk returns a tuple with the OpeningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDate

`func (o *StatementSummary) SetOpeningDate(v time.Time)`

SetOpeningDate sets OpeningDate field to given value.


### GetPastDueAmount

`func (o *StatementSummary) GetPastDueAmount() float32`

GetPastDueAmount returns the PastDueAmount field if non-nil, zero value otherwise.

### GetPastDueAmountOk

`func (o *StatementSummary) GetPastDueAmountOk() (*float32, bool)`

GetPastDueAmountOk returns a tuple with the PastDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastDueAmount

`func (o *StatementSummary) SetPastDueAmount(v float32)`

SetPastDueAmount sets PastDueAmount field to given value.


### GetPayments

`func (o *StatementSummary) GetPayments() float32`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *StatementSummary) GetPaymentsOk() (*float32, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *StatementSummary) SetPayments(v float32)`

SetPayments sets Payments field to given value.


### GetPurchases

`func (o *StatementSummary) GetPurchases() float32`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *StatementSummary) GetPurchasesOk() (*float32, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *StatementSummary) SetPurchases(v float32)`

SetPurchases sets Purchases field to given value.


### GetToken

`func (o *StatementSummary) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StatementSummary) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StatementSummary) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


