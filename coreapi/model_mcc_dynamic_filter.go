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

// checks if the MccDynamicFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MccDynamicFilter{}

// MccDynamicFilter Contains information on the dynamic merchant category code (MCC) for a reward.
type MccDynamicFilter struct {
	// One or more dynamic MCCs.
	Includes []DynamicMccType `json:"includes"`
}

// NewMccDynamicFilter instantiates a new MccDynamicFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMccDynamicFilter(includes []DynamicMccType) *MccDynamicFilter {
	this := MccDynamicFilter{}
	this.Includes = includes
	return &this
}

// NewMccDynamicFilterWithDefaults instantiates a new MccDynamicFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMccDynamicFilterWithDefaults() *MccDynamicFilter {
	this := MccDynamicFilter{}
	return &this
}

// GetIncludes returns the Includes field value
func (o *MccDynamicFilter) GetIncludes() []DynamicMccType {
	if o == nil {
		var ret []DynamicMccType
		return ret
	}

	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value
// and a boolean to check if the value has been set.
func (o *MccDynamicFilter) GetIncludesOk() ([]DynamicMccType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Includes, true
}

// SetIncludes sets field value
func (o *MccDynamicFilter) SetIncludes(v []DynamicMccType) {
	o.Includes = v
}

func (o MccDynamicFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MccDynamicFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["includes"] = o.Includes
	return toSerialize, nil
}

type NullableMccDynamicFilter struct {
	value *MccDynamicFilter
	isSet bool
}

func (v NullableMccDynamicFilter) Get() *MccDynamicFilter {
	return v.value
}

func (v *NullableMccDynamicFilter) Set(val *MccDynamicFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMccDynamicFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMccDynamicFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMccDynamicFilter(val *MccDynamicFilter) *NullableMccDynamicFilter {
	return &NullableMccDynamicFilter{value: val, isSet: true}
}

func (v NullableMccDynamicFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMccDynamicFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


