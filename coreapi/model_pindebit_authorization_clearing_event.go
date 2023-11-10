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

// checks if the PindebitAuthorizationClearingEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PindebitAuthorizationClearingEvent{}

// PindebitAuthorizationClearingEvent struct for PindebitAuthorizationClearingEvent
type PindebitAuthorizationClearingEvent struct {
	CardTransactionSimulation
	// Amount of the transaction.
	Amount float32 `json:"amount"`
	CardAcceptor *TransactionCardAcceptor `json:"card_acceptor,omitempty"`
}

// NewPindebitAuthorizationClearingEvent instantiates a new PindebitAuthorizationClearingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPindebitAuthorizationClearingEvent(amount float32) *PindebitAuthorizationClearingEvent {
	this := PindebitAuthorizationClearingEvent{}
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// NewPindebitAuthorizationClearingEventWithDefaults instantiates a new PindebitAuthorizationClearingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPindebitAuthorizationClearingEventWithDefaults() *PindebitAuthorizationClearingEvent {
	this := PindebitAuthorizationClearingEvent{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PindebitAuthorizationClearingEvent) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PindebitAuthorizationClearingEvent) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PindebitAuthorizationClearingEvent) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *PindebitAuthorizationClearingEvent) GetCardAcceptor() TransactionCardAcceptor {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret TransactionCardAcceptor
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PindebitAuthorizationClearingEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *PindebitAuthorizationClearingEvent) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given TransactionCardAcceptor and assigns it to the CardAcceptor field.
func (o *PindebitAuthorizationClearingEvent) SetCardAcceptor(v TransactionCardAcceptor) {
	o.CardAcceptor = &v
}

func (o PindebitAuthorizationClearingEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PindebitAuthorizationClearingEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCardTransactionSimulation, errCardTransactionSimulation := json.Marshal(o.CardTransactionSimulation)
	if errCardTransactionSimulation != nil {
		return map[string]interface{}{}, errCardTransactionSimulation
	}
	errCardTransactionSimulation = json.Unmarshal([]byte(serializedCardTransactionSimulation), &toSerialize)
	if errCardTransactionSimulation != nil {
		return map[string]interface{}{}, errCardTransactionSimulation
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	return toSerialize, nil
}

type NullablePindebitAuthorizationClearingEvent struct {
	value *PindebitAuthorizationClearingEvent
	isSet bool
}

func (v NullablePindebitAuthorizationClearingEvent) Get() *PindebitAuthorizationClearingEvent {
	return v.value
}

func (v *NullablePindebitAuthorizationClearingEvent) Set(val *PindebitAuthorizationClearingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePindebitAuthorizationClearingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePindebitAuthorizationClearingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePindebitAuthorizationClearingEvent(val *PindebitAuthorizationClearingEvent) *NullablePindebitAuthorizationClearingEvent {
	return &NullablePindebitAuthorizationClearingEvent{value: val, isSet: true}
}

func (v NullablePindebitAuthorizationClearingEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePindebitAuthorizationClearingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

