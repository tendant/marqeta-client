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
)

// checks if the CardTransactionSimulation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardTransactionSimulation{}

// CardTransactionSimulation struct for CardTransactionSimulation
type CardTransactionSimulation struct {
	EventType *string `json:"event_type,omitempty"`
	// The cashback amount requested.
	CashBackAmount *float32 `json:"cash_back_amount,omitempty"`
	// Indicates if the transaction is a pre-authorization. Set to `true` to mark the amount as an authorization only.
	IsPreauthorization *bool `json:"is_preauthorization,omitempty"`
	// Set to `true` to simulate a forced clearing transaction. *NOTE:* If you set this field to `true`, you must also set the `original_transaction_token` field to an existing transaction token (the transaction does not need to be related).
	ForcePost *bool `json:"force_post,omitempty"`
	// List of fees associated with the transaction. This array is returned if it exists in the resource.
	Fees []NetworkFeeModel `json:"fees,omitempty"`
	CardOptions *CardOptions `json:"card_options,omitempty"`
	Pos *Pos `json:"pos,omitempty"`
	// Indicates which card network was used to complete the transaction.
	Network *string `json:"network,omitempty"`
	// Indicates which subnetwork used to complete the transaction.
	SubNetwork *string `json:"sub_network,omitempty"`
	CurrencyConversion *CurrencyConversion `json:"currency_conversion,omitempty"`
	// Currency type of the transaction.
	CurrencyCode *string `json:"currency_code,omitempty"`
	AccountFunding *AccountFundingRequest `json:"account_funding,omitempty"`
	OriginalCredit *OriginalCredit `json:"original_credit,omitempty"`
}

// NewCardTransactionSimulation instantiates a new CardTransactionSimulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTransactionSimulation() *CardTransactionSimulation {
	this := CardTransactionSimulation{}
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// NewCardTransactionSimulationWithDefaults instantiates a new CardTransactionSimulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTransactionSimulationWithDefaults() *CardTransactionSimulation {
	this := CardTransactionSimulation{}
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CardTransactionSimulation) SetEventType(v string) {
	o.EventType = &v
}

// GetCashBackAmount returns the CashBackAmount field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetCashBackAmount() float32 {
	if o == nil || IsNil(o.CashBackAmount) {
		var ret float32
		return ret
	}
	return *o.CashBackAmount
}

// GetCashBackAmountOk returns a tuple with the CashBackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetCashBackAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CashBackAmount) {
		return nil, false
	}
	return o.CashBackAmount, true
}

// HasCashBackAmount returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasCashBackAmount() bool {
	if o != nil && !IsNil(o.CashBackAmount) {
		return true
	}

	return false
}

// SetCashBackAmount gets a reference to the given float32 and assigns it to the CashBackAmount field.
func (o *CardTransactionSimulation) SetCashBackAmount(v float32) {
	o.CashBackAmount = &v
}

// GetIsPreauthorization returns the IsPreauthorization field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetIsPreauthorization() bool {
	if o == nil || IsNil(o.IsPreauthorization) {
		var ret bool
		return ret
	}
	return *o.IsPreauthorization
}

// GetIsPreauthorizationOk returns a tuple with the IsPreauthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetIsPreauthorizationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPreauthorization) {
		return nil, false
	}
	return o.IsPreauthorization, true
}

// HasIsPreauthorization returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasIsPreauthorization() bool {
	if o != nil && !IsNil(o.IsPreauthorization) {
		return true
	}

	return false
}

// SetIsPreauthorization gets a reference to the given bool and assigns it to the IsPreauthorization field.
func (o *CardTransactionSimulation) SetIsPreauthorization(v bool) {
	o.IsPreauthorization = &v
}

// GetForcePost returns the ForcePost field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetForcePost() bool {
	if o == nil || IsNil(o.ForcePost) {
		var ret bool
		return ret
	}
	return *o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetForcePostOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePost) {
		return nil, false
	}
	return o.ForcePost, true
}

// HasForcePost returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasForcePost() bool {
	if o != nil && !IsNil(o.ForcePost) {
		return true
	}

	return false
}

// SetForcePost gets a reference to the given bool and assigns it to the ForcePost field.
func (o *CardTransactionSimulation) SetForcePost(v bool) {
	o.ForcePost = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetFees() []NetworkFeeModel {
	if o == nil || IsNil(o.Fees) {
		var ret []NetworkFeeModel
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []NetworkFeeModel and assigns it to the Fees field.
func (o *CardTransactionSimulation) SetFees(v []NetworkFeeModel) {
	o.Fees = v
}

// GetCardOptions returns the CardOptions field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetCardOptions() CardOptions {
	if o == nil || IsNil(o.CardOptions) {
		var ret CardOptions
		return ret
	}
	return *o.CardOptions
}

// GetCardOptionsOk returns a tuple with the CardOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetCardOptionsOk() (*CardOptions, bool) {
	if o == nil || IsNil(o.CardOptions) {
		return nil, false
	}
	return o.CardOptions, true
}

// HasCardOptions returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasCardOptions() bool {
	if o != nil && !IsNil(o.CardOptions) {
		return true
	}

	return false
}

// SetCardOptions gets a reference to the given CardOptions and assigns it to the CardOptions field.
func (o *CardTransactionSimulation) SetCardOptions(v CardOptions) {
	o.CardOptions = &v
}

// GetPos returns the Pos field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetPos() Pos {
	if o == nil || IsNil(o.Pos) {
		var ret Pos
		return ret
	}
	return *o.Pos
}

// GetPosOk returns a tuple with the Pos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetPosOk() (*Pos, bool) {
	if o == nil || IsNil(o.Pos) {
		return nil, false
	}
	return o.Pos, true
}

// HasPos returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasPos() bool {
	if o != nil && !IsNil(o.Pos) {
		return true
	}

	return false
}

// SetPos gets a reference to the given Pos and assigns it to the Pos field.
func (o *CardTransactionSimulation) SetPos(v Pos) {
	o.Pos = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *CardTransactionSimulation) SetNetwork(v string) {
	o.Network = &v
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetSubNetwork() string {
	if o == nil || IsNil(o.SubNetwork) {
		var ret string
		return ret
	}
	return *o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetSubNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.SubNetwork) {
		return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasSubNetwork() bool {
	if o != nil && !IsNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given string and assigns it to the SubNetwork field.
func (o *CardTransactionSimulation) SetSubNetwork(v string) {
	o.SubNetwork = &v
}

// GetCurrencyConversion returns the CurrencyConversion field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetCurrencyConversion() CurrencyConversion {
	if o == nil || IsNil(o.CurrencyConversion) {
		var ret CurrencyConversion
		return ret
	}
	return *o.CurrencyConversion
}

// GetCurrencyConversionOk returns a tuple with the CurrencyConversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetCurrencyConversionOk() (*CurrencyConversion, bool) {
	if o == nil || IsNil(o.CurrencyConversion) {
		return nil, false
	}
	return o.CurrencyConversion, true
}

// HasCurrencyConversion returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasCurrencyConversion() bool {
	if o != nil && !IsNil(o.CurrencyConversion) {
		return true
	}

	return false
}

// SetCurrencyConversion gets a reference to the given CurrencyConversion and assigns it to the CurrencyConversion field.
func (o *CardTransactionSimulation) SetCurrencyConversion(v CurrencyConversion) {
	o.CurrencyConversion = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *CardTransactionSimulation) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetAccountFunding returns the AccountFunding field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetAccountFunding() AccountFundingRequest {
	if o == nil || IsNil(o.AccountFunding) {
		var ret AccountFundingRequest
		return ret
	}
	return *o.AccountFunding
}

// GetAccountFundingOk returns a tuple with the AccountFunding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetAccountFundingOk() (*AccountFundingRequest, bool) {
	if o == nil || IsNil(o.AccountFunding) {
		return nil, false
	}
	return o.AccountFunding, true
}

// HasAccountFunding returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasAccountFunding() bool {
	if o != nil && !IsNil(o.AccountFunding) {
		return true
	}

	return false
}

// SetAccountFunding gets a reference to the given AccountFundingRequest and assigns it to the AccountFunding field.
func (o *CardTransactionSimulation) SetAccountFunding(v AccountFundingRequest) {
	o.AccountFunding = &v
}

// GetOriginalCredit returns the OriginalCredit field value if set, zero value otherwise.
func (o *CardTransactionSimulation) GetOriginalCredit() OriginalCredit {
	if o == nil || IsNil(o.OriginalCredit) {
		var ret OriginalCredit
		return ret
	}
	return *o.OriginalCredit
}

// GetOriginalCreditOk returns a tuple with the OriginalCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionSimulation) GetOriginalCreditOk() (*OriginalCredit, bool) {
	if o == nil || IsNil(o.OriginalCredit) {
		return nil, false
	}
	return o.OriginalCredit, true
}

// HasOriginalCredit returns a boolean if a field has been set.
func (o *CardTransactionSimulation) HasOriginalCredit() bool {
	if o != nil && !IsNil(o.OriginalCredit) {
		return true
	}

	return false
}

// SetOriginalCredit gets a reference to the given OriginalCredit and assigns it to the OriginalCredit field.
func (o *CardTransactionSimulation) SetOriginalCredit(v OriginalCredit) {
	o.OriginalCredit = &v
}

func (o CardTransactionSimulation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardTransactionSimulation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.CashBackAmount) {
		toSerialize["cash_back_amount"] = o.CashBackAmount
	}
	if !IsNil(o.IsPreauthorization) {
		toSerialize["is_preauthorization"] = o.IsPreauthorization
	}
	if !IsNil(o.ForcePost) {
		toSerialize["force_post"] = o.ForcePost
	}
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !IsNil(o.CardOptions) {
		toSerialize["card_options"] = o.CardOptions
	}
	if !IsNil(o.Pos) {
		toSerialize["pos"] = o.Pos
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.SubNetwork) {
		toSerialize["sub_network"] = o.SubNetwork
	}
	if !IsNil(o.CurrencyConversion) {
		toSerialize["currency_conversion"] = o.CurrencyConversion
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.AccountFunding) {
		toSerialize["account_funding"] = o.AccountFunding
	}
	if !IsNil(o.OriginalCredit) {
		toSerialize["original_credit"] = o.OriginalCredit
	}
	return toSerialize, nil
}

type NullableCardTransactionSimulation struct {
	value *CardTransactionSimulation
	isSet bool
}

func (v NullableCardTransactionSimulation) Get() *CardTransactionSimulation {
	return v.value
}

func (v *NullableCardTransactionSimulation) Set(val *CardTransactionSimulation) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTransactionSimulation) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTransactionSimulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTransactionSimulation(val *CardTransactionSimulation) *NullableCardTransactionSimulation {
	return &NullableCardTransactionSimulation{value: val, isSet: true}
}

func (v NullableCardTransactionSimulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTransactionSimulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


