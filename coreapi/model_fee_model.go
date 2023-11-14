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

// checks if the FeeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeModel{}

// FeeModel Contains attributes that define characteristics of one or more fees. This array is returned in the response when it is included in the request.
type FeeModel struct {
	// Additional text that describes the fee.
	Memo *string `json:"memo,omitempty"`
	// Descriptive metadata about the fee.
	Tags *string `json:"tags,omitempty"`
	// Unique identifier of the fee.
	Token string `json:"token"`
}

// NewFeeModel instantiates a new FeeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeModel(token string) *FeeModel {
	this := FeeModel{}
	this.Token = token
	return &this
}

// NewFeeModelWithDefaults instantiates a new FeeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeModelWithDefaults() *FeeModel {
	this := FeeModel{}
	return &this
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *FeeModel) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeModel) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *FeeModel) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *FeeModel) SetMemo(v string) {
	o.Memo = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FeeModel) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeModel) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FeeModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *FeeModel) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value
func (o *FeeModel) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FeeModel) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FeeModel) SetToken(v string) {
	o.Token = v
}

func (o FeeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableFeeModel struct {
	value *FeeModel
	isSet bool
}

func (v NullableFeeModel) Get() *FeeModel {
	return v.value
}

func (v *NullableFeeModel) Set(val *FeeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeModel(val *FeeModel) *NullableFeeModel {
	return &NullableFeeModel{value: val, isSet: true}
}

func (v NullableFeeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


