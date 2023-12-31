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

// checks if the JitProgramResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JitProgramResponse{}

// JitProgramResponse Contains the gateway's information about the JIT Funding transaction.
type JitProgramResponse struct {
	JitFunding JitFundingApi `json:"jit_funding"`
	NetworkMetadata *NetworkMetadata `json:"network_metadata,omitempty"`
}

// NewJitProgramResponse instantiates a new JitProgramResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJitProgramResponse(jitFunding JitFundingApi) *JitProgramResponse {
	this := JitProgramResponse{}
	this.JitFunding = jitFunding
	return &this
}

// NewJitProgramResponseWithDefaults instantiates a new JitProgramResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJitProgramResponseWithDefaults() *JitProgramResponse {
	this := JitProgramResponse{}
	return &this
}

// GetJitFunding returns the JitFunding field value
func (o *JitProgramResponse) GetJitFunding() JitFundingApi {
	if o == nil {
		var ret JitFundingApi
		return ret
	}

	return o.JitFunding
}

// GetJitFundingOk returns a tuple with the JitFunding field value
// and a boolean to check if the value has been set.
func (o *JitProgramResponse) GetJitFundingOk() (*JitFundingApi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JitFunding, true
}

// SetJitFunding sets field value
func (o *JitProgramResponse) SetJitFunding(v JitFundingApi) {
	o.JitFunding = v
}

// GetNetworkMetadata returns the NetworkMetadata field value if set, zero value otherwise.
func (o *JitProgramResponse) GetNetworkMetadata() NetworkMetadata {
	if o == nil || IsNil(o.NetworkMetadata) {
		var ret NetworkMetadata
		return ret
	}
	return *o.NetworkMetadata
}

// GetNetworkMetadataOk returns a tuple with the NetworkMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JitProgramResponse) GetNetworkMetadataOk() (*NetworkMetadata, bool) {
	if o == nil || IsNil(o.NetworkMetadata) {
		return nil, false
	}
	return o.NetworkMetadata, true
}

// HasNetworkMetadata returns a boolean if a field has been set.
func (o *JitProgramResponse) HasNetworkMetadata() bool {
	if o != nil && !IsNil(o.NetworkMetadata) {
		return true
	}

	return false
}

// SetNetworkMetadata gets a reference to the given NetworkMetadata and assigns it to the NetworkMetadata field.
func (o *JitProgramResponse) SetNetworkMetadata(v NetworkMetadata) {
	o.NetworkMetadata = &v
}

func (o JitProgramResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JitProgramResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jit_funding"] = o.JitFunding
	if !IsNil(o.NetworkMetadata) {
		toSerialize["network_metadata"] = o.NetworkMetadata
	}
	return toSerialize, nil
}

type NullableJitProgramResponse struct {
	value *JitProgramResponse
	isSet bool
}

func (v NullableJitProgramResponse) Get() *JitProgramResponse {
	return v.value
}

func (v *NullableJitProgramResponse) Set(val *JitProgramResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJitProgramResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJitProgramResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJitProgramResponse(val *JitProgramResponse) *NullableJitProgramResponse {
	return &NullableJitProgramResponse{value: val, isSet: true}
}

func (v NullableJitProgramResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJitProgramResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


