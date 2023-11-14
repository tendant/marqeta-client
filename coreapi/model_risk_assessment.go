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

// checks if the RiskAssessment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskAssessment{}

// RiskAssessment Contains the digital wallet provider's risk assessment for provisioning the digital wallet token.
type RiskAssessment struct {
	// Wallet provider's decision as to whether the digital wallet token should be provisioned.
	Score *string `json:"score,omitempty"`
	// Wallet provider's risk assessment version.
	Version *string `json:"version,omitempty"`
}

// NewRiskAssessment instantiates a new RiskAssessment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskAssessment() *RiskAssessment {
	this := RiskAssessment{}
	return &this
}

// NewRiskAssessmentWithDefaults instantiates a new RiskAssessment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskAssessmentWithDefaults() *RiskAssessment {
	this := RiskAssessment{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *RiskAssessment) GetScore() string {
	if o == nil || IsNil(o.Score) {
		var ret string
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessment) GetScoreOk() (*string, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *RiskAssessment) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given string and assigns it to the Score field.
func (o *RiskAssessment) SetScore(v string) {
	o.Score = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RiskAssessment) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAssessment) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RiskAssessment) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RiskAssessment) SetVersion(v string) {
	o.Version = &v
}

func (o RiskAssessment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskAssessment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRiskAssessment struct {
	value *RiskAssessment
	isSet bool
}

func (v NullableRiskAssessment) Get() *RiskAssessment {
	return v.value
}

func (v *NullableRiskAssessment) Set(val *RiskAssessment) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskAssessment) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskAssessment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskAssessment(val *RiskAssessment) *NullableRiskAssessment {
	return &NullableRiskAssessment{value: val, isSet: true}
}

func (v NullableRiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskAssessment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


