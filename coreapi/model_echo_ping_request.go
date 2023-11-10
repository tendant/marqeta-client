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

// checks if the EchoPingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EchoPingRequest{}

// EchoPingRequest struct for EchoPingRequest
type EchoPingRequest struct {
	// Payload of the ping request.
	Payload *string `json:"payload,omitempty"`
	// Unique identifier of the ping request.
	Token *string `json:"token,omitempty"`
}

// NewEchoPingRequest instantiates a new EchoPingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEchoPingRequest() *EchoPingRequest {
	this := EchoPingRequest{}
	return &this
}

// NewEchoPingRequestWithDefaults instantiates a new EchoPingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEchoPingRequestWithDefaults() *EchoPingRequest {
	this := EchoPingRequest{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *EchoPingRequest) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EchoPingRequest) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *EchoPingRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *EchoPingRequest) SetPayload(v string) {
	o.Payload = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *EchoPingRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EchoPingRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *EchoPingRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *EchoPingRequest) SetToken(v string) {
	o.Token = &v
}

func (o EchoPingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EchoPingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableEchoPingRequest struct {
	value *EchoPingRequest
	isSet bool
}

func (v NullableEchoPingRequest) Get() *EchoPingRequest {
	return v.value
}

func (v *NullableEchoPingRequest) Set(val *EchoPingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEchoPingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEchoPingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEchoPingRequest(val *EchoPingRequest) *NullableEchoPingRequest {
	return &NullableEchoPingRequest{value: val, isSet: true}
}

func (v NullableEchoPingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEchoPingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


