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
)

// checks if the CardProductConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardProductConfig{}

// CardProductConfig Defines the characteristics of the card product. Configurations are conditionally required based on program setup.  For more information, contact your Marqeta representative.
type CardProductConfig struct {
	CardLifeCycle *CardLifeCycle `json:"card_life_cycle,omitempty"`
	ClearingAndSettlement *ClearingAndSettlement `json:"clearing_and_settlement,omitempty"`
	DigitalWalletTokenization *DigitalWalletTokenization `json:"digital_wallet_tokenization,omitempty"`
	Fulfillment *CardProductFulfillment `json:"fulfillment,omitempty"`
	JitFunding *JitFunding `json:"jit_funding,omitempty"`
	Poi *Poi `json:"poi,omitempty"`
	SelectiveAuth *SelectiveAuth `json:"selective_auth,omitempty"`
	Special *Special `json:"special,omitempty"`
	TransactionControls *TransactionControls `json:"transaction_controls,omitempty"`
}

// NewCardProductConfig instantiates a new CardProductConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProductConfig() *CardProductConfig {
	this := CardProductConfig{}
	return &this
}

// NewCardProductConfigWithDefaults instantiates a new CardProductConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductConfigWithDefaults() *CardProductConfig {
	this := CardProductConfig{}
	return &this
}

// GetCardLifeCycle returns the CardLifeCycle field value if set, zero value otherwise.
func (o *CardProductConfig) GetCardLifeCycle() CardLifeCycle {
	if o == nil || IsNil(o.CardLifeCycle) {
		var ret CardLifeCycle
		return ret
	}
	return *o.CardLifeCycle
}

// GetCardLifeCycleOk returns a tuple with the CardLifeCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetCardLifeCycleOk() (*CardLifeCycle, bool) {
	if o == nil || IsNil(o.CardLifeCycle) {
		return nil, false
	}
	return o.CardLifeCycle, true
}

// HasCardLifeCycle returns a boolean if a field has been set.
func (o *CardProductConfig) HasCardLifeCycle() bool {
	if o != nil && !IsNil(o.CardLifeCycle) {
		return true
	}

	return false
}

// SetCardLifeCycle gets a reference to the given CardLifeCycle and assigns it to the CardLifeCycle field.
func (o *CardProductConfig) SetCardLifeCycle(v CardLifeCycle) {
	o.CardLifeCycle = &v
}

// GetClearingAndSettlement returns the ClearingAndSettlement field value if set, zero value otherwise.
func (o *CardProductConfig) GetClearingAndSettlement() ClearingAndSettlement {
	if o == nil || IsNil(o.ClearingAndSettlement) {
		var ret ClearingAndSettlement
		return ret
	}
	return *o.ClearingAndSettlement
}

// GetClearingAndSettlementOk returns a tuple with the ClearingAndSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetClearingAndSettlementOk() (*ClearingAndSettlement, bool) {
	if o == nil || IsNil(o.ClearingAndSettlement) {
		return nil, false
	}
	return o.ClearingAndSettlement, true
}

// HasClearingAndSettlement returns a boolean if a field has been set.
func (o *CardProductConfig) HasClearingAndSettlement() bool {
	if o != nil && !IsNil(o.ClearingAndSettlement) {
		return true
	}

	return false
}

// SetClearingAndSettlement gets a reference to the given ClearingAndSettlement and assigns it to the ClearingAndSettlement field.
func (o *CardProductConfig) SetClearingAndSettlement(v ClearingAndSettlement) {
	o.ClearingAndSettlement = &v
}

// GetDigitalWalletTokenization returns the DigitalWalletTokenization field value if set, zero value otherwise.
func (o *CardProductConfig) GetDigitalWalletTokenization() DigitalWalletTokenization {
	if o == nil || IsNil(o.DigitalWalletTokenization) {
		var ret DigitalWalletTokenization
		return ret
	}
	return *o.DigitalWalletTokenization
}

// GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool) {
	if o == nil || IsNil(o.DigitalWalletTokenization) {
		return nil, false
	}
	return o.DigitalWalletTokenization, true
}

// HasDigitalWalletTokenization returns a boolean if a field has been set.
func (o *CardProductConfig) HasDigitalWalletTokenization() bool {
	if o != nil && !IsNil(o.DigitalWalletTokenization) {
		return true
	}

	return false
}

// SetDigitalWalletTokenization gets a reference to the given DigitalWalletTokenization and assigns it to the DigitalWalletTokenization field.
func (o *CardProductConfig) SetDigitalWalletTokenization(v DigitalWalletTokenization) {
	o.DigitalWalletTokenization = &v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *CardProductConfig) GetFulfillment() CardProductFulfillment {
	if o == nil || IsNil(o.Fulfillment) {
		var ret CardProductFulfillment
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetFulfillmentOk() (*CardProductFulfillment, bool) {
	if o == nil || IsNil(o.Fulfillment) {
		return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *CardProductConfig) HasFulfillment() bool {
	if o != nil && !IsNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given CardProductFulfillment and assigns it to the Fulfillment field.
func (o *CardProductConfig) SetFulfillment(v CardProductFulfillment) {
	o.Fulfillment = &v
}

// GetJitFunding returns the JitFunding field value if set, zero value otherwise.
func (o *CardProductConfig) GetJitFunding() JitFunding {
	if o == nil || IsNil(o.JitFunding) {
		var ret JitFunding
		return ret
	}
	return *o.JitFunding
}

// GetJitFundingOk returns a tuple with the JitFunding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetJitFundingOk() (*JitFunding, bool) {
	if o == nil || IsNil(o.JitFunding) {
		return nil, false
	}
	return o.JitFunding, true
}

// HasJitFunding returns a boolean if a field has been set.
func (o *CardProductConfig) HasJitFunding() bool {
	if o != nil && !IsNil(o.JitFunding) {
		return true
	}

	return false
}

// SetJitFunding gets a reference to the given JitFunding and assigns it to the JitFunding field.
func (o *CardProductConfig) SetJitFunding(v JitFunding) {
	o.JitFunding = &v
}

// GetPoi returns the Poi field value if set, zero value otherwise.
func (o *CardProductConfig) GetPoi() Poi {
	if o == nil || IsNil(o.Poi) {
		var ret Poi
		return ret
	}
	return *o.Poi
}

// GetPoiOk returns a tuple with the Poi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetPoiOk() (*Poi, bool) {
	if o == nil || IsNil(o.Poi) {
		return nil, false
	}
	return o.Poi, true
}

// HasPoi returns a boolean if a field has been set.
func (o *CardProductConfig) HasPoi() bool {
	if o != nil && !IsNil(o.Poi) {
		return true
	}

	return false
}

// SetPoi gets a reference to the given Poi and assigns it to the Poi field.
func (o *CardProductConfig) SetPoi(v Poi) {
	o.Poi = &v
}

// GetSelectiveAuth returns the SelectiveAuth field value if set, zero value otherwise.
func (o *CardProductConfig) GetSelectiveAuth() SelectiveAuth {
	if o == nil || IsNil(o.SelectiveAuth) {
		var ret SelectiveAuth
		return ret
	}
	return *o.SelectiveAuth
}

// GetSelectiveAuthOk returns a tuple with the SelectiveAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetSelectiveAuthOk() (*SelectiveAuth, bool) {
	if o == nil || IsNil(o.SelectiveAuth) {
		return nil, false
	}
	return o.SelectiveAuth, true
}

// HasSelectiveAuth returns a boolean if a field has been set.
func (o *CardProductConfig) HasSelectiveAuth() bool {
	if o != nil && !IsNil(o.SelectiveAuth) {
		return true
	}

	return false
}

// SetSelectiveAuth gets a reference to the given SelectiveAuth and assigns it to the SelectiveAuth field.
func (o *CardProductConfig) SetSelectiveAuth(v SelectiveAuth) {
	o.SelectiveAuth = &v
}

// GetSpecial returns the Special field value if set, zero value otherwise.
func (o *CardProductConfig) GetSpecial() Special {
	if o == nil || IsNil(o.Special) {
		var ret Special
		return ret
	}
	return *o.Special
}

// GetSpecialOk returns a tuple with the Special field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetSpecialOk() (*Special, bool) {
	if o == nil || IsNil(o.Special) {
		return nil, false
	}
	return o.Special, true
}

// HasSpecial returns a boolean if a field has been set.
func (o *CardProductConfig) HasSpecial() bool {
	if o != nil && !IsNil(o.Special) {
		return true
	}

	return false
}

// SetSpecial gets a reference to the given Special and assigns it to the Special field.
func (o *CardProductConfig) SetSpecial(v Special) {
	o.Special = &v
}

// GetTransactionControls returns the TransactionControls field value if set, zero value otherwise.
func (o *CardProductConfig) GetTransactionControls() TransactionControls {
	if o == nil || IsNil(o.TransactionControls) {
		var ret TransactionControls
		return ret
	}
	return *o.TransactionControls
}

// GetTransactionControlsOk returns a tuple with the TransactionControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductConfig) GetTransactionControlsOk() (*TransactionControls, bool) {
	if o == nil || IsNil(o.TransactionControls) {
		return nil, false
	}
	return o.TransactionControls, true
}

// HasTransactionControls returns a boolean if a field has been set.
func (o *CardProductConfig) HasTransactionControls() bool {
	if o != nil && !IsNil(o.TransactionControls) {
		return true
	}

	return false
}

// SetTransactionControls gets a reference to the given TransactionControls and assigns it to the TransactionControls field.
func (o *CardProductConfig) SetTransactionControls(v TransactionControls) {
	o.TransactionControls = &v
}

func (o CardProductConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardProductConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardLifeCycle) {
		toSerialize["card_life_cycle"] = o.CardLifeCycle
	}
	if !IsNil(o.ClearingAndSettlement) {
		toSerialize["clearing_and_settlement"] = o.ClearingAndSettlement
	}
	if !IsNil(o.DigitalWalletTokenization) {
		toSerialize["digital_wallet_tokenization"] = o.DigitalWalletTokenization
	}
	if !IsNil(o.Fulfillment) {
		toSerialize["fulfillment"] = o.Fulfillment
	}
	if !IsNil(o.JitFunding) {
		toSerialize["jit_funding"] = o.JitFunding
	}
	if !IsNil(o.Poi) {
		toSerialize["poi"] = o.Poi
	}
	if !IsNil(o.SelectiveAuth) {
		toSerialize["selective_auth"] = o.SelectiveAuth
	}
	if !IsNil(o.Special) {
		toSerialize["special"] = o.Special
	}
	if !IsNil(o.TransactionControls) {
		toSerialize["transaction_controls"] = o.TransactionControls
	}
	return toSerialize, nil
}

type NullableCardProductConfig struct {
	value *CardProductConfig
	isSet bool
}

func (v NullableCardProductConfig) Get() *CardProductConfig {
	return v.value
}

func (v *NullableCardProductConfig) Set(val *CardProductConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductConfig(val *CardProductConfig) *NullableCardProductConfig {
	return &NullableCardProductConfig{value: val, isSet: true}
}

func (v NullableCardProductConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


