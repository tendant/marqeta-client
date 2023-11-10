# ProgramReserveAccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | Pointer to **float32** | Ledger balance, minus any authorized transactions that have not yet cleared. When using JIT Funding, this balance is usually equal to $0.00. | [optional] [readonly] 
**Balances** | Pointer to [**map[string]ProgramReserveAccountBalance**](ProgramReserveAccountBalance.md) | Contains program reserve account balance information, organized by currency code. Sometimes referred to as a _program funding account_. | [optional] 
**CreditBalance** | Pointer to **float32** | Not currently in use. | [optional] [readonly] 
**CurrencyCode** | Pointer to **string** | Three-digit ISO 4217 currency code. | [optional] 
**LedgerBalance** | Pointer to **float32** | When using standard funding: The funds that are available to spend immediately, including funds from any authorized transactions that have not yet cleared. When using Just-in-Time (JIT) Funding: Authorized funds that are currently on hold, but not yet cleared. | [optional] [readonly] 
**PendingCredits** | Pointer to **float32** | ACH loads that have been accepted, but for which the funding time has not yet elapsed. | [optional] [readonly] 

## Methods

### NewProgramReserveAccountBalance

`func NewProgramReserveAccountBalance() *ProgramReserveAccountBalance`

NewProgramReserveAccountBalance instantiates a new ProgramReserveAccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramReserveAccountBalanceWithDefaults

`func NewProgramReserveAccountBalanceWithDefaults() *ProgramReserveAccountBalance`

NewProgramReserveAccountBalanceWithDefaults instantiates a new ProgramReserveAccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableBalance

`func (o *ProgramReserveAccountBalance) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *ProgramReserveAccountBalance) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *ProgramReserveAccountBalance) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *ProgramReserveAccountBalance) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetBalances

`func (o *ProgramReserveAccountBalance) GetBalances() map[string]ProgramReserveAccountBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *ProgramReserveAccountBalance) GetBalancesOk() (*map[string]ProgramReserveAccountBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *ProgramReserveAccountBalance) SetBalances(v map[string]ProgramReserveAccountBalance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *ProgramReserveAccountBalance) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetCreditBalance

`func (o *ProgramReserveAccountBalance) GetCreditBalance() float32`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *ProgramReserveAccountBalance) GetCreditBalanceOk() (*float32, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *ProgramReserveAccountBalance) SetCreditBalance(v float32)`

SetCreditBalance sets CreditBalance field to given value.

### HasCreditBalance

`func (o *ProgramReserveAccountBalance) HasCreditBalance() bool`

HasCreditBalance returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ProgramReserveAccountBalance) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProgramReserveAccountBalance) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProgramReserveAccountBalance) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ProgramReserveAccountBalance) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetLedgerBalance

`func (o *ProgramReserveAccountBalance) GetLedgerBalance() float32`

GetLedgerBalance returns the LedgerBalance field if non-nil, zero value otherwise.

### GetLedgerBalanceOk

`func (o *ProgramReserveAccountBalance) GetLedgerBalanceOk() (*float32, bool)`

GetLedgerBalanceOk returns a tuple with the LedgerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerBalance

`func (o *ProgramReserveAccountBalance) SetLedgerBalance(v float32)`

SetLedgerBalance sets LedgerBalance field to given value.

### HasLedgerBalance

`func (o *ProgramReserveAccountBalance) HasLedgerBalance() bool`

HasLedgerBalance returns a boolean if a field has been set.

### GetPendingCredits

`func (o *ProgramReserveAccountBalance) GetPendingCredits() float32`

GetPendingCredits returns the PendingCredits field if non-nil, zero value otherwise.

### GetPendingCreditsOk

`func (o *ProgramReserveAccountBalance) GetPendingCreditsOk() (*float32, bool)`

GetPendingCreditsOk returns a tuple with the PendingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCredits

`func (o *ProgramReserveAccountBalance) SetPendingCredits(v float32)`

SetPendingCredits sets PendingCredits field to given value.

### HasPendingCredits

`func (o *ProgramReserveAccountBalance) HasPendingCredits() bool`

HasPendingCredits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


