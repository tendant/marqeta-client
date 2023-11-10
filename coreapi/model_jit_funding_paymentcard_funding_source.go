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

// checks if the JitFundingPaymentcardFundingSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JitFundingPaymentcardFundingSource{}

// JitFundingPaymentcardFundingSource struct for JitFundingPaymentcardFundingSource
type JitFundingPaymentcardFundingSource struct {
	// Specifies whether JIT Funding is enabled or disabled for the payment card funding source. A value of `true` indicates that the payment card funding source is enabled and will be debited when swipes occur.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies the return destination for refunds in the case of a transaction reversal.
	RefundsDestination *string `json:"refunds_destination,omitempty"`
}

// NewJitFundingPaymentcardFundingSource instantiates a new JitFundingPaymentcardFundingSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJitFundingPaymentcardFundingSource() *JitFundingPaymentcardFundingSource {
	this := JitFundingPaymentcardFundingSource{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewJitFundingPaymentcardFundingSourceWithDefaults instantiates a new JitFundingPaymentcardFundingSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJitFundingPaymentcardFundingSourceWithDefaults() *JitFundingPaymentcardFundingSource {
	this := JitFundingPaymentcardFundingSource{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *JitFundingPaymentcardFundingSource) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JitFundingPaymentcardFundingSource) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *JitFundingPaymentcardFundingSource) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *JitFundingPaymentcardFundingSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRefundsDestination returns the RefundsDestination field value if set, zero value otherwise.
func (o *JitFundingPaymentcardFundingSource) GetRefundsDestination() string {
	if o == nil || IsNil(o.RefundsDestination) {
		var ret string
		return ret
	}
	return *o.RefundsDestination
}

// GetRefundsDestinationOk returns a tuple with the RefundsDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JitFundingPaymentcardFundingSource) GetRefundsDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.RefundsDestination) {
		return nil, false
	}
	return o.RefundsDestination, true
}

// HasRefundsDestination returns a boolean if a field has been set.
func (o *JitFundingPaymentcardFundingSource) HasRefundsDestination() bool {
	if o != nil && !IsNil(o.RefundsDestination) {
		return true
	}

	return false
}

// SetRefundsDestination gets a reference to the given string and assigns it to the RefundsDestination field.
func (o *JitFundingPaymentcardFundingSource) SetRefundsDestination(v string) {
	o.RefundsDestination = &v
}

func (o JitFundingPaymentcardFundingSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JitFundingPaymentcardFundingSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.RefundsDestination) {
		toSerialize["refunds_destination"] = o.RefundsDestination
	}
	return toSerialize, nil
}

type NullableJitFundingPaymentcardFundingSource struct {
	value *JitFundingPaymentcardFundingSource
	isSet bool
}

func (v NullableJitFundingPaymentcardFundingSource) Get() *JitFundingPaymentcardFundingSource {
	return v.value
}

func (v *NullableJitFundingPaymentcardFundingSource) Set(val *JitFundingPaymentcardFundingSource) {
	v.value = val
	v.isSet = true
}

func (v NullableJitFundingPaymentcardFundingSource) IsSet() bool {
	return v.isSet
}

func (v *NullableJitFundingPaymentcardFundingSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJitFundingPaymentcardFundingSource(val *JitFundingPaymentcardFundingSource) *NullableJitFundingPaymentcardFundingSource {
	return &NullableJitFundingPaymentcardFundingSource{value: val, isSet: true}
}

func (v NullableJitFundingPaymentcardFundingSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJitFundingPaymentcardFundingSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

