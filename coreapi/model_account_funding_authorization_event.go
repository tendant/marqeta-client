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

// checks if the AccountFundingAuthorizationEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountFundingAuthorizationEvent{}

// AccountFundingAuthorizationEvent struct for AccountFundingAuthorizationEvent
type AccountFundingAuthorizationEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	CardToken string `json:"card_token"`
	// Amount of the transaction.
	Amount float32 `json:"amount"`
	CardAcceptor *TransactionCardAcceptor `json:"card_acceptor,omitempty"`
}

// NewAccountFundingAuthorizationEvent instantiates a new AccountFundingAuthorizationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountFundingAuthorizationEvent(cardToken string, amount float32) *AccountFundingAuthorizationEvent {
	this := AccountFundingAuthorizationEvent{}
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// NewAccountFundingAuthorizationEventWithDefaults instantiates a new AccountFundingAuthorizationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountFundingAuthorizationEventWithDefaults() *AccountFundingAuthorizationEvent {
	this := AccountFundingAuthorizationEvent{}
	return &this
}

// GetCardToken returns the CardToken field value
func (o *AccountFundingAuthorizationEvent) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *AccountFundingAuthorizationEvent) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *AccountFundingAuthorizationEvent) SetCardToken(v string) {
	o.CardToken = v
}

// GetAmount returns the Amount field value
func (o *AccountFundingAuthorizationEvent) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AccountFundingAuthorizationEvent) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AccountFundingAuthorizationEvent) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *AccountFundingAuthorizationEvent) GetCardAcceptor() TransactionCardAcceptor {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret TransactionCardAcceptor
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingAuthorizationEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *AccountFundingAuthorizationEvent) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given TransactionCardAcceptor and assigns it to the CardAcceptor field.
func (o *AccountFundingAuthorizationEvent) SetCardAcceptor(v TransactionCardAcceptor) {
	o.CardAcceptor = &v
}

func (o AccountFundingAuthorizationEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountFundingAuthorizationEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCardTransactionSimulation, errCardTransactionSimulation := json.Marshal(o.CardTransactionSimulation)
	if errCardTransactionSimulation != nil {
		return map[string]interface{}{}, errCardTransactionSimulation
	}
	errCardTransactionSimulation = json.Unmarshal([]byte(serializedCardTransactionSimulation), &toSerialize)
	if errCardTransactionSimulation != nil {
		return map[string]interface{}{}, errCardTransactionSimulation
	}
	toSerialize["card_token"] = o.CardToken
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	return toSerialize, nil
}

type NullableAccountFundingAuthorizationEvent struct {
	value *AccountFundingAuthorizationEvent
	isSet bool
}

func (v NullableAccountFundingAuthorizationEvent) Get() *AccountFundingAuthorizationEvent {
	return v.value
}

func (v *NullableAccountFundingAuthorizationEvent) Set(val *AccountFundingAuthorizationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountFundingAuthorizationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountFundingAuthorizationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountFundingAuthorizationEvent(val *AccountFundingAuthorizationEvent) *NullableAccountFundingAuthorizationEvent {
	return &NullableAccountFundingAuthorizationEvent{value: val, isSet: true}
}

func (v NullableAccountFundingAuthorizationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountFundingAuthorizationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


