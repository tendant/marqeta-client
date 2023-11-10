# TransactionModelGpa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedgerBalance** | Pointer to **float32** | When using standard funding: The funds that are available to spend immediately, including funds from any authorized transactions that have not yet cleared. When using Just-in-Time (JIT) Funding: Authorized funds that are currently on hold, but not yet cleared. | [optional] 
**AvailableBalance** | Pointer to **float32** | Ledger balance minus any authorized transactions that have not yet cleared. Also known as the cardholder&#39;s purchasing power. When using JIT Funding, this balance is usually equal to $0.00. | [optional] 
**ImpactedAmount** | Pointer to **float32** | Balance change based on the amount of the transaction. | [optional] 

## Methods

### NewTransactionModelGpa

`func NewTransactionModelGpa() *TransactionModelGpa`

NewTransactionModelGpa instantiates a new TransactionModelGpa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionModelGpaWithDefaults

`func NewTransactionModelGpaWithDefaults() *TransactionModelGpa`

NewTransactionModelGpaWithDefaults instantiates a new TransactionModelGpa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedgerBalance

`func (o *TransactionModelGpa) GetLedgerBalance() float32`

GetLedgerBalance returns the LedgerBalance field if non-nil, zero value otherwise.

### GetLedgerBalanceOk

`func (o *TransactionModelGpa) GetLedgerBalanceOk() (*float32, bool)`

GetLedgerBalanceOk returns a tuple with the LedgerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerBalance

`func (o *TransactionModelGpa) SetLedgerBalance(v float32)`

SetLedgerBalance sets LedgerBalance field to given value.

### HasLedgerBalance

`func (o *TransactionModelGpa) HasLedgerBalance() bool`

HasLedgerBalance returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *TransactionModelGpa) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *TransactionModelGpa) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *TransactionModelGpa) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *TransactionModelGpa) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetImpactedAmount

`func (o *TransactionModelGpa) GetImpactedAmount() float32`

GetImpactedAmount returns the ImpactedAmount field if non-nil, zero value otherwise.

### GetImpactedAmountOk

`func (o *TransactionModelGpa) GetImpactedAmountOk() (*float32, bool)`

GetImpactedAmountOk returns a tuple with the ImpactedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedAmount

`func (o *TransactionModelGpa) SetImpactedAmount(v float32)`

SetImpactedAmount sets ImpactedAmount field to given value.

### HasImpactedAmount

`func (o *TransactionModelGpa) HasImpactedAmount() bool`

HasImpactedAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


