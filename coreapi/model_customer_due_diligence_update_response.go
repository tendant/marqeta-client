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

// checks if the CustomerDueDiligenceUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDueDiligenceUpdateResponse{}

// CustomerDueDiligenceUpdateResponse struct for CustomerDueDiligenceUpdateResponse
type CustomerDueDiligenceUpdateResponse struct {
	Answer *string `json:"answer,omitempty"`
}

// NewCustomerDueDiligenceUpdateResponse instantiates a new CustomerDueDiligenceUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDueDiligenceUpdateResponse() *CustomerDueDiligenceUpdateResponse {
	this := CustomerDueDiligenceUpdateResponse{}
	return &this
}

// NewCustomerDueDiligenceUpdateResponseWithDefaults instantiates a new CustomerDueDiligenceUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDueDiligenceUpdateResponseWithDefaults() *CustomerDueDiligenceUpdateResponse {
	this := CustomerDueDiligenceUpdateResponse{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *CustomerDueDiligenceUpdateResponse) GetAnswer() string {
	if o == nil || IsNil(o.Answer) {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDueDiligenceUpdateResponse) GetAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *CustomerDueDiligenceUpdateResponse) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *CustomerDueDiligenceUpdateResponse) SetAnswer(v string) {
	o.Answer = &v
}

func (o CustomerDueDiligenceUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDueDiligenceUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	return toSerialize, nil
}

type NullableCustomerDueDiligenceUpdateResponse struct {
	value *CustomerDueDiligenceUpdateResponse
	isSet bool
}

func (v NullableCustomerDueDiligenceUpdateResponse) Get() *CustomerDueDiligenceUpdateResponse {
	return v.value
}

func (v *NullableCustomerDueDiligenceUpdateResponse) Set(val *CustomerDueDiligenceUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDueDiligenceUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDueDiligenceUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDueDiligenceUpdateResponse(val *CustomerDueDiligenceUpdateResponse) *NullableCustomerDueDiligenceUpdateResponse {
	return &NullableCustomerDueDiligenceUpdateResponse{value: val, isSet: true}
}

func (v NullableCustomerDueDiligenceUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDueDiligenceUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

