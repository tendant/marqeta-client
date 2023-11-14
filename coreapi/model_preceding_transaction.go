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

// checks if the PrecedingTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrecedingTransaction{}

// PrecedingTransaction Returned for `authorization.clearing` transaction types following a financial advice.  Contains information about the preceding transaction.
type PrecedingTransaction struct {
	// Amount of the preceding transaction.
	Amount *float32 `json:"amount,omitempty"`
	// Unique identifier of the preceding transaction.
	Token *string `json:"token,omitempty"`
}

// NewPrecedingTransaction instantiates a new PrecedingTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrecedingTransaction() *PrecedingTransaction {
	this := PrecedingTransaction{}
	return &this
}

// NewPrecedingTransactionWithDefaults instantiates a new PrecedingTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrecedingTransactionWithDefaults() *PrecedingTransaction {
	this := PrecedingTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PrecedingTransaction) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecedingTransaction) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PrecedingTransaction) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *PrecedingTransaction) SetAmount(v float32) {
	o.Amount = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PrecedingTransaction) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecedingTransaction) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PrecedingTransaction) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PrecedingTransaction) SetToken(v string) {
	o.Token = &v
}

func (o PrecedingTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrecedingTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullablePrecedingTransaction struct {
	value *PrecedingTransaction
	isSet bool
}

func (v NullablePrecedingTransaction) Get() *PrecedingTransaction {
	return v.value
}

func (v *NullablePrecedingTransaction) Set(val *PrecedingTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecedingTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecedingTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecedingTransaction(val *PrecedingTransaction) *NullablePrecedingTransaction {
	return &NullablePrecedingTransaction{value: val, isSet: true}
}

func (v NullablePrecedingTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecedingTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


