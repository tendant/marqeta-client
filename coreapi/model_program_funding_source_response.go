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
	"time"
)

// checks if the ProgramFundingSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramFundingSourceResponse{}

// ProgramFundingSourceResponse struct for ProgramFundingSourceResponse
type ProgramFundingSourceResponse struct {
	// The account identifier.
	Account string `json:"account"`
	// Indicates whether the program funding source is active. This field is returned if it exists in the resource.
	Active *bool `json:"active,omitempty"`
	// Date and time when the resource was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Date and time when the resource was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Name of the program funding source.
	Name string `json:"name"`
	// Unique identifier of the funding source.
	Token string `json:"token"`
}

// NewProgramFundingSourceResponse instantiates a new ProgramFundingSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramFundingSourceResponse(account string, createdTime time.Time, lastModifiedTime time.Time, name string, token string) *ProgramFundingSourceResponse {
	this := ProgramFundingSourceResponse{}
	this.Account = account
	this.CreatedTime = createdTime
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	this.Token = token
	return &this
}

// NewProgramFundingSourceResponseWithDefaults instantiates a new ProgramFundingSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramFundingSourceResponseWithDefaults() *ProgramFundingSourceResponse {
	this := ProgramFundingSourceResponse{}
	return &this
}

// GetAccount returns the Account field value
func (o *ProgramFundingSourceResponse) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceResponse) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ProgramFundingSourceResponse) SetAccount(v string) {
	o.Account = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ProgramFundingSourceResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ProgramFundingSourceResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ProgramFundingSourceResponse) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *ProgramFundingSourceResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *ProgramFundingSourceResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *ProgramFundingSourceResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *ProgramFundingSourceResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetName returns the Name field value
func (o *ProgramFundingSourceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProgramFundingSourceResponse) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value
func (o *ProgramFundingSourceResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ProgramFundingSourceResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ProgramFundingSourceResponse) SetToken(v string) {
	o.Token = v
}

func (o ProgramFundingSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramFundingSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["name"] = o.Name
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableProgramFundingSourceResponse struct {
	value *ProgramFundingSourceResponse
	isSet bool
}

func (v NullableProgramFundingSourceResponse) Get() *ProgramFundingSourceResponse {
	return v.value
}

func (v *NullableProgramFundingSourceResponse) Set(val *ProgramFundingSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramFundingSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramFundingSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramFundingSourceResponse(val *ProgramFundingSourceResponse) *NullableProgramFundingSourceResponse {
	return &NullableProgramFundingSourceResponse{value: val, isSet: true}
}

func (v NullableProgramFundingSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramFundingSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


