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

// checks if the ResultCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultCode{}

// ResultCode Contains the KYC result code and a descriptive message about that codes.
type ResultCode struct {
	// For any `PENDING` or `FAILURE` outcome, see the <<user_kyc_failure_codes, User KYC failure codes>> table, the <<outcome_reasons_for_the_business, Outcome reasons for the business>> table, or the <<outcome_reasons_for_individuals_associated_with_a_business, Outcome reasons for individuals associated with a business>> table.
	Code *string `json:"code,omitempty"`
	// Result code description.
	Message *string `json:"message,omitempty"`
}

// NewResultCode instantiates a new ResultCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultCode() *ResultCode {
	this := ResultCode{}
	return &this
}

// NewResultCodeWithDefaults instantiates a new ResultCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultCodeWithDefaults() *ResultCode {
	this := ResultCode{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResultCode) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCode) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResultCode) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ResultCode) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResultCode) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultCode) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResultCode) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResultCode) SetMessage(v string) {
	o.Message = &v
}

func (o ResultCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableResultCode struct {
	value *ResultCode
	isSet bool
}

func (v NullableResultCode) Get() *ResultCode {
	return v.value
}

func (v *NullableResultCode) Set(val *ResultCode) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCode) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCode(val *ResultCode) *NullableResultCode {
	return &NullableResultCode{value: val, isSet: true}
}

func (v NullableResultCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


