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

// checks if the PolicyProductMinPaymentCalculation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyProductMinPaymentCalculation{}

// PolicyProductMinPaymentCalculation Contains information used to calculate the minimum payment amount on a credit product policy.
type PolicyProductMinPaymentCalculation struct {
	// Whether to include the overlimit amount when calculating the minimum payment.
	IncludeOverlimitAmount bool `json:"include_overlimit_amount"`
	// Whether to include the past due amount when calculating the minimum payment.
	IncludePastDueAmount bool `json:"include_past_due_amount"`
	// Minimum payment, expressed as a flat amount, due on the payment due day.
	MinPaymentFlatAmount *float32 `json:"min_payment_flat_amount,omitempty"`
	MinPaymentPercentage *PolicyProductMinPaymentPercentage `json:"min_payment_percentage,omitempty"`
}

// NewPolicyProductMinPaymentCalculation instantiates a new PolicyProductMinPaymentCalculation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyProductMinPaymentCalculation(includeOverlimitAmount bool, includePastDueAmount bool) *PolicyProductMinPaymentCalculation {
	this := PolicyProductMinPaymentCalculation{}
	this.IncludeOverlimitAmount = includeOverlimitAmount
	this.IncludePastDueAmount = includePastDueAmount
	return &this
}

// NewPolicyProductMinPaymentCalculationWithDefaults instantiates a new PolicyProductMinPaymentCalculation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyProductMinPaymentCalculationWithDefaults() *PolicyProductMinPaymentCalculation {
	this := PolicyProductMinPaymentCalculation{}
	return &this
}

// GetIncludeOverlimitAmount returns the IncludeOverlimitAmount field value
func (o *PolicyProductMinPaymentCalculation) GetIncludeOverlimitAmount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeOverlimitAmount
}

// GetIncludeOverlimitAmountOk returns a tuple with the IncludeOverlimitAmount field value
// and a boolean to check if the value has been set.
func (o *PolicyProductMinPaymentCalculation) GetIncludeOverlimitAmountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeOverlimitAmount, true
}

// SetIncludeOverlimitAmount sets field value
func (o *PolicyProductMinPaymentCalculation) SetIncludeOverlimitAmount(v bool) {
	o.IncludeOverlimitAmount = v
}

// GetIncludePastDueAmount returns the IncludePastDueAmount field value
func (o *PolicyProductMinPaymentCalculation) GetIncludePastDueAmount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludePastDueAmount
}

// GetIncludePastDueAmountOk returns a tuple with the IncludePastDueAmount field value
// and a boolean to check if the value has been set.
func (o *PolicyProductMinPaymentCalculation) GetIncludePastDueAmountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludePastDueAmount, true
}

// SetIncludePastDueAmount sets field value
func (o *PolicyProductMinPaymentCalculation) SetIncludePastDueAmount(v bool) {
	o.IncludePastDueAmount = v
}

// GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field value if set, zero value otherwise.
func (o *PolicyProductMinPaymentCalculation) GetMinPaymentFlatAmount() float32 {
	if o == nil || IsNil(o.MinPaymentFlatAmount) {
		var ret float32
		return ret
	}
	return *o.MinPaymentFlatAmount
}

// GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyProductMinPaymentCalculation) GetMinPaymentFlatAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.MinPaymentFlatAmount) {
		return nil, false
	}
	return o.MinPaymentFlatAmount, true
}

// HasMinPaymentFlatAmount returns a boolean if a field has been set.
func (o *PolicyProductMinPaymentCalculation) HasMinPaymentFlatAmount() bool {
	if o != nil && !IsNil(o.MinPaymentFlatAmount) {
		return true
	}

	return false
}

// SetMinPaymentFlatAmount gets a reference to the given float32 and assigns it to the MinPaymentFlatAmount field.
func (o *PolicyProductMinPaymentCalculation) SetMinPaymentFlatAmount(v float32) {
	o.MinPaymentFlatAmount = &v
}

// GetMinPaymentPercentage returns the MinPaymentPercentage field value if set, zero value otherwise.
func (o *PolicyProductMinPaymentCalculation) GetMinPaymentPercentage() PolicyProductMinPaymentPercentage {
	if o == nil || IsNil(o.MinPaymentPercentage) {
		var ret PolicyProductMinPaymentPercentage
		return ret
	}
	return *o.MinPaymentPercentage
}

// GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyProductMinPaymentCalculation) GetMinPaymentPercentageOk() (*PolicyProductMinPaymentPercentage, bool) {
	if o == nil || IsNil(o.MinPaymentPercentage) {
		return nil, false
	}
	return o.MinPaymentPercentage, true
}

// HasMinPaymentPercentage returns a boolean if a field has been set.
func (o *PolicyProductMinPaymentCalculation) HasMinPaymentPercentage() bool {
	if o != nil && !IsNil(o.MinPaymentPercentage) {
		return true
	}

	return false
}

// SetMinPaymentPercentage gets a reference to the given PolicyProductMinPaymentPercentage and assigns it to the MinPaymentPercentage field.
func (o *PolicyProductMinPaymentCalculation) SetMinPaymentPercentage(v PolicyProductMinPaymentPercentage) {
	o.MinPaymentPercentage = &v
}

func (o PolicyProductMinPaymentCalculation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyProductMinPaymentCalculation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include_overlimit_amount"] = o.IncludeOverlimitAmount
	toSerialize["include_past_due_amount"] = o.IncludePastDueAmount
	if !IsNil(o.MinPaymentFlatAmount) {
		toSerialize["min_payment_flat_amount"] = o.MinPaymentFlatAmount
	}
	if !IsNil(o.MinPaymentPercentage) {
		toSerialize["min_payment_percentage"] = o.MinPaymentPercentage
	}
	return toSerialize, nil
}

type NullablePolicyProductMinPaymentCalculation struct {
	value *PolicyProductMinPaymentCalculation
	isSet bool
}

func (v NullablePolicyProductMinPaymentCalculation) Get() *PolicyProductMinPaymentCalculation {
	return v.value
}

func (v *NullablePolicyProductMinPaymentCalculation) Set(val *PolicyProductMinPaymentCalculation) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyProductMinPaymentCalculation) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyProductMinPaymentCalculation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyProductMinPaymentCalculation(val *PolicyProductMinPaymentCalculation) *NullablePolicyProductMinPaymentCalculation {
	return &NullablePolicyProductMinPaymentCalculation{value: val, isSet: true}
}

func (v NullablePolicyProductMinPaymentCalculation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyProductMinPaymentCalculation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


