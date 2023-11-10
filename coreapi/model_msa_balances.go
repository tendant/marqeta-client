/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marqeta_coreapi_client

import (
	"encoding/json"
	"time"
)

// checks if the MsaBalances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MsaBalances{}

// MsaBalances struct for MsaBalances
type MsaBalances struct {
	AvailableBalance float32 `json:"available_balance"`
	Balances map[string]MsaBalances `json:"balances"`
	CachedBalance float32 `json:"cached_balance"`
	CreditBalance float32 `json:"credit_balance"`
	CurrencyCode string `json:"currency_code"`
	ImpactedAmount *float32 `json:"impacted_amount,omitempty"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	LedgerBalance float32 `json:"ledger_balance"`
	PendingCredits float32 `json:"pending_credits"`
}

// NewMsaBalances instantiates a new MsaBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsaBalances(availableBalance float32, balances map[string]MsaBalances, cachedBalance float32, creditBalance float32, currencyCode string, lastUpdatedTime time.Time, ledgerBalance float32, pendingCredits float32) *MsaBalances {
	this := MsaBalances{}
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

// NewMsaBalancesWithDefaults instantiates a new MsaBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsaBalancesWithDefaults() *MsaBalances {
	this := MsaBalances{}
	return &this
}

// GetAvailableBalance returns the AvailableBalance field value
func (o *MsaBalances) GetAvailableBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AvailableBalance
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetAvailableBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableBalance, true
}

// SetAvailableBalance sets field value
func (o *MsaBalances) SetAvailableBalance(v float32) {
	o.AvailableBalance = v
}

// GetBalances returns the Balances field value
func (o *MsaBalances) GetBalances() map[string]MsaBalances {
	if o == nil {
		var ret map[string]MsaBalances
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetBalancesOk() (*map[string]MsaBalances, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *MsaBalances) SetBalances(v map[string]MsaBalances) {
	o.Balances = v
}

// GetCachedBalance returns the CachedBalance field value
func (o *MsaBalances) GetCachedBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CachedBalance
}

// GetCachedBalanceOk returns a tuple with the CachedBalance field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetCachedBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CachedBalance, true
}

// SetCachedBalance sets field value
func (o *MsaBalances) SetCachedBalance(v float32) {
	o.CachedBalance = v
}

// GetCreditBalance returns the CreditBalance field value
func (o *MsaBalances) GetCreditBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreditBalance
}

// GetCreditBalanceOk returns a tuple with the CreditBalance field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetCreditBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditBalance, true
}

// SetCreditBalance sets field value
func (o *MsaBalances) SetCreditBalance(v float32) {
	o.CreditBalance = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *MsaBalances) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *MsaBalances) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetImpactedAmount returns the ImpactedAmount field value if set, zero value otherwise.
func (o *MsaBalances) GetImpactedAmount() float32 {
	if o == nil || IsNil(o.ImpactedAmount) {
		var ret float32
		return ret
	}
	return *o.ImpactedAmount
}

// GetImpactedAmountOk returns a tuple with the ImpactedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetImpactedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ImpactedAmount) {
		return nil, false
	}
	return o.ImpactedAmount, true
}

// HasImpactedAmount returns a boolean if a field has been set.
func (o *MsaBalances) HasImpactedAmount() bool {
	if o != nil && !IsNil(o.ImpactedAmount) {
		return true
	}

	return false
}

// SetImpactedAmount gets a reference to the given float32 and assigns it to the ImpactedAmount field.
func (o *MsaBalances) SetImpactedAmount(v float32) {
	o.ImpactedAmount = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *MsaBalances) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *MsaBalances) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetLedgerBalance returns the LedgerBalance field value
func (o *MsaBalances) GetLedgerBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LedgerBalance
}

// GetLedgerBalanceOk returns a tuple with the LedgerBalance field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetLedgerBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerBalance, true
}

// SetLedgerBalance sets field value
func (o *MsaBalances) SetLedgerBalance(v float32) {
	o.LedgerBalance = v
}

// GetPendingCredits returns the PendingCredits field value
func (o *MsaBalances) GetPendingCredits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingCredits
}

// GetPendingCreditsOk returns a tuple with the PendingCredits field value
// and a boolean to check if the value has been set.
func (o *MsaBalances) GetPendingCreditsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingCredits, true
}

// SetPendingCredits sets field value
func (o *MsaBalances) SetPendingCredits(v float32) {
	o.PendingCredits = v
}

func (o MsaBalances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MsaBalances) ToMap() (map[string]interface{}, error) {
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

type NullableMsaBalances struct {
	value *MsaBalances
	isSet bool
}

func (v NullableMsaBalances) Get() *MsaBalances {
	return v.value
}

func (v *NullableMsaBalances) Set(val *MsaBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableMsaBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableMsaBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsaBalances(val *MsaBalances) *NullableMsaBalances {
	return &NullableMsaBalances{value: val, isSet: true}
}

func (v NullableMsaBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsaBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

