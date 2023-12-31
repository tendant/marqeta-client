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

// checks if the AuthorizationAdviceEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationAdviceEvent{}

// AuthorizationAdviceEvent struct for AuthorizationAdviceEvent
type AuthorizationAdviceEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	PrecedingRelatedTransactionToken string `json:"preceding_related_transaction_token"`
	// Amount of the transaction.
	Amount float32 `json:"amount"`
	CardAcceptor *TransactionCardAcceptor `json:"card_acceptor,omitempty"`
}

// NewAuthorizationAdviceEvent instantiates a new AuthorizationAdviceEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationAdviceEvent(precedingRelatedTransactionToken string, amount float32) *AuthorizationAdviceEvent {
	this := AuthorizationAdviceEvent{}
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// NewAuthorizationAdviceEventWithDefaults instantiates a new AuthorizationAdviceEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationAdviceEventWithDefaults() *AuthorizationAdviceEvent {
	this := AuthorizationAdviceEvent{}
	return &this
}

// GetPrecedingRelatedTransactionToken returns the PrecedingRelatedTransactionToken field value
func (o *AuthorizationAdviceEvent) GetPrecedingRelatedTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrecedingRelatedTransactionToken
}

// GetPrecedingRelatedTransactionTokenOk returns a tuple with the PrecedingRelatedTransactionToken field value
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceEvent) GetPrecedingRelatedTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecedingRelatedTransactionToken, true
}

// SetPrecedingRelatedTransactionToken sets field value
func (o *AuthorizationAdviceEvent) SetPrecedingRelatedTransactionToken(v string) {
	o.PrecedingRelatedTransactionToken = v
}

// GetAmount returns the Amount field value
func (o *AuthorizationAdviceEvent) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceEvent) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AuthorizationAdviceEvent) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *AuthorizationAdviceEvent) GetCardAcceptor() TransactionCardAcceptor {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret TransactionCardAcceptor
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationAdviceEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *AuthorizationAdviceEvent) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given TransactionCardAcceptor and assigns it to the CardAcceptor field.
func (o *AuthorizationAdviceEvent) SetCardAcceptor(v TransactionCardAcceptor) {
	o.CardAcceptor = &v
}

func (o AuthorizationAdviceEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationAdviceEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCardTransactionSimulation, errCardTransactionSimulation := json.Marshal(o.CardTransactionSimulation)
	if errCardTransactionSimulation != nil {
		return map[string]interface{}{}, errCardTransactionSimulation
	}
	errCardTransactionSimulation = json.Unmarshal([]byte(serializedCardTransactionSimulation), &toSerialize)
	if errCardTransactionSimulation != nil {
		return map[string]interface{}{}, errCardTransactionSimulation
	}
	toSerialize["preceding_related_transaction_token"] = o.PrecedingRelatedTransactionToken
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	return toSerialize, nil
}

type NullableAuthorizationAdviceEvent struct {
	value *AuthorizationAdviceEvent
	isSet bool
}

func (v NullableAuthorizationAdviceEvent) Get() *AuthorizationAdviceEvent {
	return v.value
}

func (v *NullableAuthorizationAdviceEvent) Set(val *AuthorizationAdviceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationAdviceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationAdviceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationAdviceEvent(val *AuthorizationAdviceEvent) *NullableAuthorizationAdviceEvent {
	return &NullableAuthorizationAdviceEvent{value: val, isSet: true}
}

func (v NullableAuthorizationAdviceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationAdviceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


