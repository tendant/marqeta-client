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

// checks if the ProgramFundingSourceUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramFundingSourceUpdateRequest{}

// ProgramFundingSourceUpdateRequest struct for ProgramFundingSourceUpdateRequest
type ProgramFundingSourceUpdateRequest struct {
	// Indicates whether the program funding source is active.
	Active *bool `json:"active,omitempty"`
	// Name of the program funding source.
	Name *string `json:"name,omitempty"`
}

// NewProgramFundingSourceUpdateRequest instantiates a new ProgramFundingSourceUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramFundingSourceUpdateRequest() *ProgramFundingSourceUpdateRequest {
	this := ProgramFundingSourceUpdateRequest{}
	return &this
}

// NewProgramFundingSourceUpdateRequestWithDefaults instantiates a new ProgramFundingSourceUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramFundingSourceUpdateRequestWithDefaults() *ProgramFundingSourceUpdateRequest {
	this := ProgramFundingSourceUpdateRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ProgramFundingSourceUpdateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ProgramFundingSourceUpdateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ProgramFundingSourceUpdateRequest) SetActive(v bool) {
	o.Active = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProgramFundingSourceUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProgramFundingSourceUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProgramFundingSourceUpdateRequest) SetName(v string) {
	o.Name = &v
}

func (o ProgramFundingSourceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramFundingSourceUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableProgramFundingSourceUpdateRequest struct {
	value *ProgramFundingSourceUpdateRequest
	isSet bool
}

func (v NullableProgramFundingSourceUpdateRequest) Get() *ProgramFundingSourceUpdateRequest {
	return v.value
}

func (v *NullableProgramFundingSourceUpdateRequest) Set(val *ProgramFundingSourceUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramFundingSourceUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramFundingSourceUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramFundingSourceUpdateRequest(val *ProgramFundingSourceUpdateRequest) *NullableProgramFundingSourceUpdateRequest {
	return &NullableProgramFundingSourceUpdateRequest{value: val, isSet: true}
}

func (v NullableProgramFundingSourceUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramFundingSourceUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

