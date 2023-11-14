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

// checks if the ReturnedDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnedDetails{}

// ReturnedDetails Contains information on a returned payment.
type ReturnedDetails struct {
	// Return code for the returned payment. For more, see <</developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes>>.
	ReturnCode string `json:"return_code"`
	// Reason the payment was returned. For more, see <</developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes>>.
	ReturnReason string `json:"return_reason"`
}

// NewReturnedDetails instantiates a new ReturnedDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnedDetails(returnCode string, returnReason string) *ReturnedDetails {
	this := ReturnedDetails{}
	this.ReturnCode = returnCode
	this.ReturnReason = returnReason
	return &this
}

// NewReturnedDetailsWithDefaults instantiates a new ReturnedDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnedDetailsWithDefaults() *ReturnedDetails {
	this := ReturnedDetails{}
	return &this
}

// GetReturnCode returns the ReturnCode field value
func (o *ReturnedDetails) GetReturnCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value
// and a boolean to check if the value has been set.
func (o *ReturnedDetails) GetReturnCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnCode, true
}

// SetReturnCode sets field value
func (o *ReturnedDetails) SetReturnCode(v string) {
	o.ReturnCode = v
}

// GetReturnReason returns the ReturnReason field value
func (o *ReturnedDetails) GetReturnReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value
// and a boolean to check if the value has been set.
func (o *ReturnedDetails) GetReturnReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnReason, true
}

// SetReturnReason sets field value
func (o *ReturnedDetails) SetReturnReason(v string) {
	o.ReturnReason = v
}

func (o ReturnedDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnedDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["return_code"] = o.ReturnCode
	toSerialize["return_reason"] = o.ReturnReason
	return toSerialize, nil
}

type NullableReturnedDetails struct {
	value *ReturnedDetails
	isSet bool
}

func (v NullableReturnedDetails) Get() *ReturnedDetails {
	return v.value
}

func (v *NullableReturnedDetails) Set(val *ReturnedDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnedDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnedDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnedDetails(val *ReturnedDetails) *NullableReturnedDetails {
	return &NullableReturnedDetails{value: val, isSet: true}
}

func (v NullableReturnedDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnedDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


