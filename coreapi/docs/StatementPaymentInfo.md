# StatementPaymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | Date and time when the statement payment information was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**MinimumPaymentDue** | Pointer to **float32** | Minimum payment amount for the current statement period, based on the associated credit product settings. | [optional] 
**NewStatementBalance** | Pointer to **float32** | Balance on the credit account when the statement period ended. | [optional] 
**PaymentCutoffDate** | Pointer to **time.Time** | Last day a payment can be made before interest and fees are charged to the account. | [optional] 
**PaymentDueDate** | Pointer to **time.Time** | Payment due date, based on the credit account settings. | [optional] 
**StatementSummaryToken** | Pointer to **string** | Unique identifier of the statement summary. | [optional] 
**ThreeYearSavings** | Pointer to **float32** | Savings amount if the balance is paid off in three years versus only making minimum payments. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the statement payment. | [optional] 
**Warnings** | Pointer to [**[]StatementPaymentWarning**](StatementPaymentWarning.md) | One or more payoff warnings. | [optional] 

## Methods

### NewStatementPaymentInfo

`func NewStatementPaymentInfo() *StatementPaymentInfo`

NewStatementPaymentInfo instantiates a new StatementPaymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementPaymentInfoWithDefaults

`func NewStatementPaymentInfoWithDefaults() *StatementPaymentInfo`

NewStatementPaymentInfoWithDefaults instantiates a new StatementPaymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *StatementPaymentInfo) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StatementPaymentInfo) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StatementPaymentInfo) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *StatementPaymentInfo) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetMinimumPaymentDue

`func (o *StatementPaymentInfo) GetMinimumPaymentDue() float32`

GetMinimumPaymentDue returns the MinimumPaymentDue field if non-nil, zero value otherwise.

### GetMinimumPaymentDueOk

`func (o *StatementPaymentInfo) GetMinimumPaymentDueOk() (*float32, bool)`

GetMinimumPaymentDueOk returns a tuple with the MinimumPaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPaymentDue

`func (o *StatementPaymentInfo) SetMinimumPaymentDue(v float32)`

SetMinimumPaymentDue sets MinimumPaymentDue field to given value.

### HasMinimumPaymentDue

`func (o *StatementPaymentInfo) HasMinimumPaymentDue() bool`

HasMinimumPaymentDue returns a boolean if a field has been set.

### GetNewStatementBalance

`func (o *StatementPaymentInfo) GetNewStatementBalance() float32`

GetNewStatementBalance returns the NewStatementBalance field if non-nil, zero value otherwise.

### GetNewStatementBalanceOk

`func (o *StatementPaymentInfo) GetNewStatementBalanceOk() (*float32, bool)`

GetNewStatementBalanceOk returns a tuple with the NewStatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatementBalance

`func (o *StatementPaymentInfo) SetNewStatementBalance(v float32)`

SetNewStatementBalance sets NewStatementBalance field to given value.

### HasNewStatementBalance

`func (o *StatementPaymentInfo) HasNewStatementBalance() bool`

HasNewStatementBalance returns a boolean if a field has been set.

### GetPaymentCutoffDate

`func (o *StatementPaymentInfo) GetPaymentCutoffDate() time.Time`

GetPaymentCutoffDate returns the PaymentCutoffDate field if non-nil, zero value otherwise.

### GetPaymentCutoffDateOk

`func (o *StatementPaymentInfo) GetPaymentCutoffDateOk() (*time.Time, bool)`

GetPaymentCutoffDateOk returns a tuple with the PaymentCutoffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCutoffDate

`func (o *StatementPaymentInfo) SetPaymentCutoffDate(v time.Time)`

SetPaymentCutoffDate sets PaymentCutoffDate field to given value.

### HasPaymentCutoffDate

`func (o *StatementPaymentInfo) HasPaymentCutoffDate() bool`

HasPaymentCutoffDate returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *StatementPaymentInfo) GetPaymentDueDate() time.Time`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *StatementPaymentInfo) GetPaymentDueDateOk() (*time.Time, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *StatementPaymentInfo) SetPaymentDueDate(v time.Time)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *StatementPaymentInfo) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetStatementSummaryToken

`func (o *StatementPaymentInfo) GetStatementSummaryToken() string`

GetStatementSummaryToken returns the StatementSummaryToken field if non-nil, zero value otherwise.

### GetStatementSummaryTokenOk

`func (o *StatementPaymentInfo) GetStatementSummaryTokenOk() (*string, bool)`

GetStatementSummaryTokenOk returns a tuple with the StatementSummaryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementSummaryToken

`func (o *StatementPaymentInfo) SetStatementSummaryToken(v string)`

SetStatementSummaryToken sets StatementSummaryToken field to given value.

### HasStatementSummaryToken

`func (o *StatementPaymentInfo) HasStatementSummaryToken() bool`

HasStatementSummaryToken returns a boolean if a field has been set.

### GetThreeYearSavings

`func (o *StatementPaymentInfo) GetThreeYearSavings() float32`

GetThreeYearSavings returns the ThreeYearSavings field if non-nil, zero value otherwise.

### GetThreeYearSavingsOk

`func (o *StatementPaymentInfo) GetThreeYearSavingsOk() (*float32, bool)`

GetThreeYearSavingsOk returns a tuple with the ThreeYearSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeYearSavings

`func (o *StatementPaymentInfo) SetThreeYearSavings(v float32)`

SetThreeYearSavings sets ThreeYearSavings field to given value.

### HasThreeYearSavings

`func (o *StatementPaymentInfo) HasThreeYearSavings() bool`

HasThreeYearSavings returns a boolean if a field has been set.

### GetToken

`func (o *StatementPaymentInfo) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StatementPaymentInfo) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StatementPaymentInfo) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *StatementPaymentInfo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetWarnings

`func (o *StatementPaymentInfo) GetWarnings() []StatementPaymentWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *StatementPaymentInfo) GetWarningsOk() (*[]StatementPaymentWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *StatementPaymentInfo) SetWarnings(v []StatementPaymentWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *StatementPaymentInfo) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


