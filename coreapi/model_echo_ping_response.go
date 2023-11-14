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

// checks if the EchoPingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EchoPingResponse{}

// EchoPingResponse struct for EchoPingResponse
type EchoPingResponse struct {
	Id *string `json:"id,omitempty"`
	Payload *string `json:"payload,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewEchoPingResponse instantiates a new EchoPingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEchoPingResponse() *EchoPingResponse {
	this := EchoPingResponse{}
	var success bool = false
	this.Success = &success
	return &this
}

// NewEchoPingResponseWithDefaults instantiates a new EchoPingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEchoPingResponseWithDefaults() *EchoPingResponse {
	this := EchoPingResponse{}
	var success bool = false
	this.Success = &success
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EchoPingResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EchoPingResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EchoPingResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EchoPingResponse) SetId(v string) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *EchoPingResponse) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EchoPingResponse) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *EchoPingResponse) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *EchoPingResponse) SetPayload(v string) {
	o.Payload = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *EchoPingResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EchoPingResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *EchoPingResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *EchoPingResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o EchoPingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EchoPingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableEchoPingResponse struct {
	value *EchoPingResponse
	isSet bool
}

func (v NullableEchoPingResponse) Get() *EchoPingResponse {
	return v.value
}

func (v *NullableEchoPingResponse) Set(val *EchoPingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEchoPingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEchoPingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEchoPingResponse(val *EchoPingResponse) *NullableEchoPingResponse {
	return &NullableEchoPingResponse{value: val, isSet: true}
}

func (v NullableEchoPingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEchoPingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


