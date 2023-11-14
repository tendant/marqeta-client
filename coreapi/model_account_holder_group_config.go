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

// checks if the AccountHolderGroupConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountHolderGroupConfig{}

// AccountHolderGroupConfig Contains configuration fields for the account holder group.
type AccountHolderGroupConfig struct {
	// If set to `false`, this control prohibits an account holder's account from being reloaded with funds after the initial load.  This restriction applies to GPA orders, peer transfers, and direct deposits, but does not apply to operator adjustments.
	IsReloadable *bool `json:"is_reloadable,omitempty"`
	// If set to `ALWAYS`, new account holders are created in an `UNVERIFIED` status and must pass identity verification (KYC) before they can be active; if set to `CONDITIONAL`, new account holders begin in a `LIMITED` status and have limited actions available before passing identity verification; if set to `NEVER`, new account holders are created in an active state.
	KycRequired *string `json:"kyc_required,omitempty"`
	PreKycControls *PreKycControls `json:"pre_kyc_controls,omitempty"`
	// Associates the specified real-time fee group with the members of the account holder group.
	RealTimeFeeGroupToken *string `json:"real_time_fee_group_token,omitempty"`
}

// NewAccountHolderGroupConfig instantiates a new AccountHolderGroupConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderGroupConfig() *AccountHolderGroupConfig {
	this := AccountHolderGroupConfig{}
	var isReloadable bool = false
	this.IsReloadable = &isReloadable
	return &this
}

// NewAccountHolderGroupConfigWithDefaults instantiates a new AccountHolderGroupConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderGroupConfigWithDefaults() *AccountHolderGroupConfig {
	this := AccountHolderGroupConfig{}
	var isReloadable bool = false
	this.IsReloadable = &isReloadable
	return &this
}

// GetIsReloadable returns the IsReloadable field value if set, zero value otherwise.
func (o *AccountHolderGroupConfig) GetIsReloadable() bool {
	if o == nil || IsNil(o.IsReloadable) {
		var ret bool
		return ret
	}
	return *o.IsReloadable
}

// GetIsReloadableOk returns a tuple with the IsReloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderGroupConfig) GetIsReloadableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReloadable) {
		return nil, false
	}
	return o.IsReloadable, true
}

// HasIsReloadable returns a boolean if a field has been set.
func (o *AccountHolderGroupConfig) HasIsReloadable() bool {
	if o != nil && !IsNil(o.IsReloadable) {
		return true
	}

	return false
}

// SetIsReloadable gets a reference to the given bool and assigns it to the IsReloadable field.
func (o *AccountHolderGroupConfig) SetIsReloadable(v bool) {
	o.IsReloadable = &v
}

// GetKycRequired returns the KycRequired field value if set, zero value otherwise.
func (o *AccountHolderGroupConfig) GetKycRequired() string {
	if o == nil || IsNil(o.KycRequired) {
		var ret string
		return ret
	}
	return *o.KycRequired
}

// GetKycRequiredOk returns a tuple with the KycRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderGroupConfig) GetKycRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.KycRequired) {
		return nil, false
	}
	return o.KycRequired, true
}

// HasKycRequired returns a boolean if a field has been set.
func (o *AccountHolderGroupConfig) HasKycRequired() bool {
	if o != nil && !IsNil(o.KycRequired) {
		return true
	}

	return false
}

// SetKycRequired gets a reference to the given string and assigns it to the KycRequired field.
func (o *AccountHolderGroupConfig) SetKycRequired(v string) {
	o.KycRequired = &v
}

// GetPreKycControls returns the PreKycControls field value if set, zero value otherwise.
func (o *AccountHolderGroupConfig) GetPreKycControls() PreKycControls {
	if o == nil || IsNil(o.PreKycControls) {
		var ret PreKycControls
		return ret
	}
	return *o.PreKycControls
}

// GetPreKycControlsOk returns a tuple with the PreKycControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderGroupConfig) GetPreKycControlsOk() (*PreKycControls, bool) {
	if o == nil || IsNil(o.PreKycControls) {
		return nil, false
	}
	return o.PreKycControls, true
}

// HasPreKycControls returns a boolean if a field has been set.
func (o *AccountHolderGroupConfig) HasPreKycControls() bool {
	if o != nil && !IsNil(o.PreKycControls) {
		return true
	}

	return false
}

// SetPreKycControls gets a reference to the given PreKycControls and assigns it to the PreKycControls field.
func (o *AccountHolderGroupConfig) SetPreKycControls(v PreKycControls) {
	o.PreKycControls = &v
}

// GetRealTimeFeeGroupToken returns the RealTimeFeeGroupToken field value if set, zero value otherwise.
func (o *AccountHolderGroupConfig) GetRealTimeFeeGroupToken() string {
	if o == nil || IsNil(o.RealTimeFeeGroupToken) {
		var ret string
		return ret
	}
	return *o.RealTimeFeeGroupToken
}

// GetRealTimeFeeGroupTokenOk returns a tuple with the RealTimeFeeGroupToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderGroupConfig) GetRealTimeFeeGroupTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RealTimeFeeGroupToken) {
		return nil, false
	}
	return o.RealTimeFeeGroupToken, true
}

// HasRealTimeFeeGroupToken returns a boolean if a field has been set.
func (o *AccountHolderGroupConfig) HasRealTimeFeeGroupToken() bool {
	if o != nil && !IsNil(o.RealTimeFeeGroupToken) {
		return true
	}

	return false
}

// SetRealTimeFeeGroupToken gets a reference to the given string and assigns it to the RealTimeFeeGroupToken field.
func (o *AccountHolderGroupConfig) SetRealTimeFeeGroupToken(v string) {
	o.RealTimeFeeGroupToken = &v
}

func (o AccountHolderGroupConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountHolderGroupConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsReloadable) {
		toSerialize["is_reloadable"] = o.IsReloadable
	}
	if !IsNil(o.KycRequired) {
		toSerialize["kyc_required"] = o.KycRequired
	}
	if !IsNil(o.PreKycControls) {
		toSerialize["pre_kyc_controls"] = o.PreKycControls
	}
	if !IsNil(o.RealTimeFeeGroupToken) {
		toSerialize["real_time_fee_group_token"] = o.RealTimeFeeGroupToken
	}
	return toSerialize, nil
}

type NullableAccountHolderGroupConfig struct {
	value *AccountHolderGroupConfig
	isSet bool
}

func (v NullableAccountHolderGroupConfig) Get() *AccountHolderGroupConfig {
	return v.value
}

func (v *NullableAccountHolderGroupConfig) Set(val *AccountHolderGroupConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderGroupConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderGroupConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderGroupConfig(val *AccountHolderGroupConfig) *NullableAccountHolderGroupConfig {
	return &NullableAccountHolderGroupConfig{value: val, isSet: true}
}

func (v NullableAccountHolderGroupConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderGroupConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


