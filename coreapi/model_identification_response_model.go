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

// checks if the IdentificationResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentificationResponseModel{}

// IdentificationResponseModel Contains identifications associated with the cardholder.
type IdentificationResponseModel struct {
	// Expiration date for the identification, if applicable.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// Type of identification.
	Type *string `json:"type,omitempty"`
	// Number associated with the identification.
	Value *string `json:"value,omitempty"`
}

// NewIdentificationResponseModel instantiates a new IdentificationResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationResponseModel() *IdentificationResponseModel {
	this := IdentificationResponseModel{}
	return &this
}

// NewIdentificationResponseModelWithDefaults instantiates a new IdentificationResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationResponseModelWithDefaults() *IdentificationResponseModel {
	this := IdentificationResponseModel{}
	return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *IdentificationResponseModel) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationResponseModel) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *IdentificationResponseModel) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *IdentificationResponseModel) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentificationResponseModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationResponseModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentificationResponseModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentificationResponseModel) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IdentificationResponseModel) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationResponseModel) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IdentificationResponseModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IdentificationResponseModel) SetValue(v string) {
	o.Value = &v
}

func (o IdentificationResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentificationResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableIdentificationResponseModel struct {
	value *IdentificationResponseModel
	isSet bool
}

func (v NullableIdentificationResponseModel) Get() *IdentificationResponseModel {
	return v.value
}

func (v *NullableIdentificationResponseModel) Set(val *IdentificationResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationResponseModel(val *IdentificationResponseModel) *NullableIdentificationResponseModel {
	return &NullableIdentificationResponseModel{value: val, isSet: true}
}

func (v NullableIdentificationResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


