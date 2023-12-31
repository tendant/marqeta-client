/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreapi

import (
	"encoding/json"
	"time"
)

// checks if the CardholderBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardholderBalance{}

// CardholderBalance struct for CardholderBalance
type CardholderBalance struct {
	AvailableBalance float32 `json:"available_balance"`
	Balances map[string]CardholderBalance `json:"balances"`
	CachedBalance float32 `json:"cached_balance"`
	CreditBalance float32 `json:"credit_balance"`
	CurrencyCode string `json:"currency_code"`
	ImpactedAmount *float32 `json:"impacted_amount,omitempty"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	LedgerBalance float32 `json:"ledger_balance"`
	PendingCredits float32 `json:"pending_credits"`
}

// NewCardholderBalance instantiates a new CardholderBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardholderBalance(availableBalance float32, balances map[string]CardholderBalance, cachedBalance float32, creditBalance float32, currencyCode string, lastUpdatedTime time.Time, ledgerBalance float32, pendingCredits float32) *CardholderBalance {
	this := CardholderBalance{}
	this.AvailableBalance = availableBalance
	this.Balances = balances
	this.CachedBalance = cachedBalance
	this.CreditBalance = creditBalance
	this.CurrencyCode = currencyCode
	this.LastUpdatedTime = lastUpdatedTime
	this.LedgerBalance = ledgerBalance
	this.PendingCredits = pendingCredits
	return &this
}

// NewCardholderBalanceWithDefaults instantiates a new CardholderBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardholderBalanceWithDefaults() *CardholderBalance {
	this := CardholderBalance{}
	return &this
}

// GetAvailableBalance returns the AvailableBalance field value
func (o *CardholderBalance) GetAvailableBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetAvailableBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableBalance, true
}

// SetAvailableBalance sets field value
func (o *CardholderBalance) SetAvailableBalance(v float32) {
	o.AvailableBalance = v
}

// GetBalances returns the Balances field value
func (o *CardholderBalance) GetBalances() map[string]CardholderBalance {
	if o == nil {
		var ret map[string]CardholderBalance
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetBalancesOk() (*map[string]CardholderBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *CardholderBalance) SetBalances(v map[string]CardholderBalance) {
	o.Balances = v
}

// GetCachedBalance returns the CachedBalance field value
func (o *CardholderBalance) GetCachedBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CachedBalance
}

// GetCachedBalanceOk returns a tuple with the CachedBalance field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetCachedBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CachedBalance, true
}

// SetCachedBalance sets field value
func (o *CardholderBalance) SetCachedBalance(v float32) {
	o.CachedBalance = v
}

// GetCreditBalance returns the CreditBalance field value
func (o *CardholderBalance) GetCreditBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreditBalance
}

// GetCreditBalanceOk returns a tuple with the CreditBalance field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetCreditBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditBalance, true
}

// SetCreditBalance sets field value
func (o *CardholderBalance) SetCreditBalance(v float32) {
	o.CreditBalance = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CardholderBalance) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CardholderBalance) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetImpactedAmount returns the ImpactedAmount field value if set, zero value otherwise.
func (o *CardholderBalance) GetImpactedAmount() float32 {
	if o == nil || IsNil(o.ImpactedAmount) {
		var ret float32
		return ret
	}
	return *o.ImpactedAmount
}

// GetImpactedAmountOk returns a tuple with the ImpactedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetImpactedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ImpactedAmount) {
		return nil, false
	}
	return o.ImpactedAmount, true
}

// HasImpactedAmount returns a boolean if a field has been set.
func (o *CardholderBalance) HasImpactedAmount() bool {
	if o != nil && !IsNil(o.ImpactedAmount) {
		return true
	}

	return false
}

// SetImpactedAmount gets a reference to the given float32 and assigns it to the ImpactedAmount field.
func (o *CardholderBalance) SetImpactedAmount(v float32) {
	o.ImpactedAmount = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *CardholderBalance) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *CardholderBalance) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetLedgerBalance returns the LedgerBalance field value
func (o *CardholderBalance) GetLedgerBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LedgerBalance
}

// GetLedgerBalanceOk returns a tuple with the LedgerBalance field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetLedgerBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerBalance, true
}

// SetLedgerBalance sets field value
func (o *CardholderBalance) SetLedgerBalance(v float32) {
	o.LedgerBalance = v
}

// GetPendingCredits returns the PendingCredits field value
func (o *CardholderBalance) GetPendingCredits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingCredits
}

// GetPendingCreditsOk returns a tuple with the PendingCredits field value
// and a boolean to check if the value has been set.
func (o *CardholderBalance) GetPendingCreditsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingCredits, true
}

// SetPendingCredits sets field value
func (o *CardholderBalance) SetPendingCredits(v float32) {
	o.PendingCredits = v
}

func (o CardholderBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardholderBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["available_balance"] = o.AvailableBalance
	toSerialize["balances"] = o.Balances
	toSerialize["cached_balance"] = o.CachedBalance
	toSerialize["credit_balance"] = o.CreditBalance
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.ImpactedAmount) {
		toSerialize["impacted_amount"] = o.ImpactedAmount
	}
	toSerialize["last_updated_time"] = o.LastUpdatedTime
	toSerialize["ledger_balance"] = o.LedgerBalance
	toSerialize["pending_credits"] = o.PendingCredits
	return toSerialize, nil
}

type NullableCardholderBalance struct {
	value *CardholderBalance
	isSet bool
}

func (v NullableCardholderBalance) Get() *CardholderBalance {
	return v.value
}

func (v *NullableCardholderBalance) Set(val *CardholderBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableCardholderBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableCardholderBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardholderBalance(val *CardholderBalance) *NullableCardholderBalance {
	return &NullableCardholderBalance{value: val, isSet: true}
}

func (v NullableCardholderBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardholderBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


