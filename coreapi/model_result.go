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

// checks if the Result type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Result{}

// Result Contains information about the KYC verification result.
type Result struct {
	// An array of KYC verification result code objects.
	Codes []ResultCode `json:"codes,omitempty"`
	// KYC verification status.
	Status *string `json:"status,omitempty"`
}

// NewResult instantiates a new Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResult() *Result {
	this := Result{}
	return &this
}

// NewResultWithDefaults instantiates a new Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultWithDefaults() *Result {
	this := Result{}
	return &this
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *Result) GetCodes() []ResultCode {
	if o == nil || IsNil(o.Codes) {
		var ret []ResultCode
		return ret
	}
	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetCodesOk() ([]ResultCode, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *Result) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given []ResultCode and assigns it to the Codes field.
func (o *Result) SetCodes(v []ResultCode) {
	o.Codes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Result) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Result) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Result) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Result) SetStatus(v string) {
	o.Status = &v
}

func (o Result) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Codes) {
		toSerialize["codes"] = o.Codes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResult struct {
	value *Result
	isSet bool
}

func (v NullableResult) Get() *Result {
	return v.value
}

func (v *NullableResult) Set(val *Result) {
	v.value = val
	v.isSet = true
}

func (v NullableResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResult(val *Result) *NullableResult {
	return &NullableResult{value: val, isSet: true}
}

func (v NullableResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


