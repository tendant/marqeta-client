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

// checks if the AccountHolderGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountHolderGroupRequest{}

// AccountHolderGroupRequest struct for AccountHolderGroupRequest
type AccountHolderGroupRequest struct {
	Config *AccountHolderGroupConfig `json:"config,omitempty"`
	// Descriptive name for the account holder group.
	Name *string `json:"name,omitempty"`
	// Unique identifier of the account holder group.
	Token *string `json:"token,omitempty"`
}

// NewAccountHolderGroupRequest instantiates a new AccountHolderGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderGroupRequest() *AccountHolderGroupRequest {
	this := AccountHolderGroupRequest{}
	return &this
}

// NewAccountHolderGroupRequestWithDefaults instantiates a new AccountHolderGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderGroupRequestWithDefaults() *AccountHolderGroupRequest {
	this := AccountHolderGroupRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AccountHolderGroupRequest) GetConfig() AccountHolderGroupConfig {
	if o == nil || IsNil(o.Config) {
		var ret AccountHolderGroupConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderGroupRequest) GetConfigOk() (*AccountHolderGroupConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AccountHolderGroupRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AccountHolderGroupConfig and assigns it to the Config field.
func (o *AccountHolderGroupRequest) SetConfig(v AccountHolderGroupConfig) {
	o.Config = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountHolderGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountHolderGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountHolderGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AccountHolderGroupRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderGroupRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AccountHolderGroupRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AccountHolderGroupRequest) SetToken(v string) {
	o.Token = &v
}

func (o AccountHolderGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountHolderGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableAccountHolderGroupRequest struct {
	value *AccountHolderGroupRequest
	isSet bool
}

func (v NullableAccountHolderGroupRequest) Get() *AccountHolderGroupRequest {
	return v.value
}

func (v *NullableAccountHolderGroupRequest) Set(val *AccountHolderGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderGroupRequest(val *AccountHolderGroupRequest) *NullableAccountHolderGroupRequest {
	return &NullableAccountHolderGroupRequest{value: val, isSet: true}
}

func (v NullableAccountHolderGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


