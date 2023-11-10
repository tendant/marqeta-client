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

// checks if the GatewayResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayResponse{}

// GatewayResponse Contains information from the gateway in response to a funding request.
type GatewayResponse struct {
	// Code received from the gateway.
	Code string `json:"code"`
	Data *JitProgramResponse `json:"data,omitempty"`
}

// NewGatewayResponse instantiates a new GatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayResponse(code string) *GatewayResponse {
	this := GatewayResponse{}
	this.Code = code
	return &this
}

// NewGatewayResponseWithDefaults instantiates a new GatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayResponseWithDefaults() *GatewayResponse {
	this := GatewayResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *GatewayResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GatewayResponse) SetCode(v string) {
	o.Code = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GatewayResponse) GetData() JitProgramResponse {
	if o == nil || IsNil(o.Data) {
		var ret JitProgramResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetDataOk() (*JitProgramResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GatewayResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given JitProgramResponse and assigns it to the Data field.
func (o *GatewayResponse) SetData(v JitProgramResponse) {
	o.Data = &v
}

func (o GatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGatewayResponse struct {
	value *GatewayResponse
	isSet bool
}

func (v NullableGatewayResponse) Get() *GatewayResponse {
	return v.value
}

func (v *NullableGatewayResponse) Set(val *GatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayResponse(val *GatewayResponse) *NullableGatewayResponse {
	return &NullableGatewayResponse{value: val, isSet: true}
}

func (v NullableGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


