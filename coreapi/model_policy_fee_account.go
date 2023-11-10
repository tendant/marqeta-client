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

// checks if the PolicyFeeAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyFeeAccount{}

// PolicyFeeAccount Contains information on the fee policy on a credit account.
type PolicyFeeAccount struct {
	LatePayment *PolicyFeePayment `json:"late_payment,omitempty"`
	ReturnedPayment *PolicyFeePayment `json:"returned_payment,omitempty"`
}

// NewPolicyFeeAccount instantiates a new PolicyFeeAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyFeeAccount() *PolicyFeeAccount {
	this := PolicyFeeAccount{}
	return &this
}

// NewPolicyFeeAccountWithDefaults instantiates a new PolicyFeeAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyFeeAccountWithDefaults() *PolicyFeeAccount {
	this := PolicyFeeAccount{}
	return &this
}

// GetLatePayment returns the LatePayment field value if set, zero value otherwise.
func (o *PolicyFeeAccount) GetLatePayment() PolicyFeePayment {
	if o == nil || IsNil(o.LatePayment) {
		var ret PolicyFeePayment
		return ret
	}
	return *o.LatePayment
}

// GetLatePaymentOk returns a tuple with the LatePayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeAccount) GetLatePaymentOk() (*PolicyFeePayment, bool) {
	if o == nil || IsNil(o.LatePayment) {
		return nil, false
	}
	return o.LatePayment, true
}

// HasLatePayment returns a boolean if a field has been set.
func (o *PolicyFeeAccount) HasLatePayment() bool {
	if o != nil && !IsNil(o.LatePayment) {
		return true
	}

	return false
}

// SetLatePayment gets a reference to the given PolicyFeePayment and assigns it to the LatePayment field.
func (o *PolicyFeeAccount) SetLatePayment(v PolicyFeePayment) {
	o.LatePayment = &v
}

// GetReturnedPayment returns the ReturnedPayment field value if set, zero value otherwise.
func (o *PolicyFeeAccount) GetReturnedPayment() PolicyFeePayment {
	if o == nil || IsNil(o.ReturnedPayment) {
		var ret PolicyFeePayment
		return ret
	}
	return *o.ReturnedPayment
}

// GetReturnedPaymentOk returns a tuple with the ReturnedPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeAccount) GetReturnedPaymentOk() (*PolicyFeePayment, bool) {
	if o == nil || IsNil(o.ReturnedPayment) {
		return nil, false
	}
	return o.ReturnedPayment, true
}

// HasReturnedPayment returns a boolean if a field has been set.
func (o *PolicyFeeAccount) HasReturnedPayment() bool {
	if o != nil && !IsNil(o.ReturnedPayment) {
		return true
	}

	return false
}

// SetReturnedPayment gets a reference to the given PolicyFeePayment and assigns it to the ReturnedPayment field.
func (o *PolicyFeeAccount) SetReturnedPayment(v PolicyFeePayment) {
	o.ReturnedPayment = &v
}

func (o PolicyFeeAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyFeeAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LatePayment) {
		toSerialize["late_payment"] = o.LatePayment
	}
	if !IsNil(o.ReturnedPayment) {
		toSerialize["returned_payment"] = o.ReturnedPayment
	}
	return toSerialize, nil
}

type NullablePolicyFeeAccount struct {
	value *PolicyFeeAccount
	isSet bool
}

func (v NullablePolicyFeeAccount) Get() *PolicyFeeAccount {
	return v.value
}

func (v *NullablePolicyFeeAccount) Set(val *PolicyFeeAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyFeeAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyFeeAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyFeeAccount(val *PolicyFeeAccount) *NullablePolicyFeeAccount {
	return &NullablePolicyFeeAccount{value: val, isSet: true}
}

func (v NullablePolicyFeeAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyFeeAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


