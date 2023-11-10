# MsaAggregatedBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | **float32** |  | 
**Balances** | [**map[string]MsaAggregatedBalances**](MsaAggregatedBalances.md) |  | 
**CachedBalance** | **float32** |  | 
**CreditBalance** | **float32** |  | 
**CurrencyCode** | **string** |  | 
**ImpactedAmount** | Pointer to **float32** |  | [optional] 
**LastUpdatedTime** | **time.Time** |  | 
**LedgerBalance** | **float32** |  | 
**PendingCredits** | **float32** |  | 

## Methods

### NewMsaAggregatedBalances

`func NewMsaAggregatedBalances(availableBalance float32, balances map[string]MsaAggregatedBalances, cachedBalance float32, creditBalance float32, currencyCode string, lastUpdatedTime time.Time, ledgerBalance float32, pendingCredits float32, ) *MsaAggregatedBalances`

NewMsaAggregatedBalances instantiates a new MsaAggregatedBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsaAggregatedBalancesWithDefaults

`func NewMsaAggregatedBalancesWithDefaults() *MsaAggregatedBalances`

NewMsaAggregatedBalancesWithDefaults instantiates a new MsaAggregatedBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableBalance

`func (o *MsaAggregatedBalances) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *MsaAggregatedBalances) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *MsaAggregatedBalances) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalances

`func (o *MsaAggregatedBalances) GetBalances() map[string]MsaAggregatedBalances`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *MsaAggregatedBalances) GetBalancesOk() (*map[string]MsaAggregatedBalances, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *MsaAggregatedBalances) SetBalances(v map[string]MsaAggregatedBalances)`

SetBalances sets Balances field to given value.


### GetCachedBalance

`func (o *MsaAggregatedBalances) GetCachedBalance() float32`

GetCachedBalance returns the CachedBalance field if non-nil, zero value otherwise.

### GetCachedBalanceOk

`func (o *MsaAggregatedBalances) GetCachedBalanceOk() (*float32, bool)`

GetCachedBalanceOk returns a tuple with the CachedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedBalance

`func (o *MsaAggregatedBalances) SetCachedBalance(v float32)`

SetCachedBalance sets CachedBalance field to given value.


### GetCreditBalance

`func (o *MsaAggregatedBalances) GetCreditBalance() float32`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *MsaAggregatedBalances) GetCreditBalanceOk() (*float32, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *MsaAggregatedBalances) SetCreditBalance(v float32)`

SetCreditBalance sets CreditBalance field to given value.


### GetCurrencyCode

`func (o *MsaAggregatedBalances) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MsaAggregatedBalances) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MsaAggregatedBalances) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetImpactedAmount

`func (o *MsaAggregatedBalances) GetImpactedAmount() float32`

GetImpactedAmount returns the ImpactedAmount field if non-nil, zero value otherwise.

### GetImpactedAmountOk

`func (o *MsaAggregatedBalances) GetImpactedAmountOk() (*float32, bool)`

GetImpactedAmountOk returns a tuple with the ImpactedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedAmount

`func (o *MsaAggregatedBalances) SetImpactedAmount(v float32)`

SetImpactedAmount sets ImpactedAmount field to given value.

### HasImpactedAmount

`func (o *MsaAggregatedBalances) HasImpactedAmount() bool`

HasImpactedAmount returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *MsaAggregatedBalances) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *MsaAggregatedBalances) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *MsaAggregatedBalances) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetLedgerBalance

`func (o *MsaAggregatedBalances) GetLedgerBalance() float32`

GetLedgerBalance returns the LedgerBalance field if non-nil, zero value otherwise.

### GetLedgerBalanceOk

`func (o *MsaAggregatedBalances) GetLedgerBalanceOk() (*float32, bool)`

GetLedgerBalanceOk returns a tuple with the LedgerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerBalance

`func (o *MsaAggregatedBalances) SetLedgerBalance(v float32)`

SetLedgerBalance sets LedgerBalance field to given value.


### GetPendingCredits

`func (o *MsaAggregatedBalances) GetPendingCredits() float32`

GetPendingCredits returns the PendingCredits field if non-nil, zero value otherwise.

### GetPendingCreditsOk

`func (o *MsaAggregatedBalances) GetPendingCreditsOk() (*float32, bool)`

GetPendingCreditsOk returns a tuple with the PendingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCredits

`func (o *MsaAggregatedBalances) SetPendingCredits(v float32)`

SetPendingCredits sets PendingCredits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


