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

// checks if the CommandoModeEnables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandoModeEnables{}

// CommandoModeEnables Defines program behavior when Commando Mode is enabled.
type CommandoModeEnables struct {
	// Unique identifiers of the authorization controls enabled while in Commando Mode.
	AuthControls []string `json:"auth_controls,omitempty"`
	// If set to `true`, transactions conducted while Commando Mode is enabled proceed even when the card is suspended. If set to `false`, transactions conducted while Commando Mode is enabled are declined if the card is suspended.
	IgnoreCardSuspendedState *bool `json:"ignore_card_suspended_state,omitempty"`
	// Unique identifier of the program funding source that substitutes for the program gateway funding source upon Commando Mode enablement.
	ProgramFundingSource string `json:"program_funding_source"`
	// This field is not currently in use.
	UseCacheBalance *bool `json:"use_cache_balance,omitempty"`
	// Unique identifiers of the velocity controls enabled while in Commando Mode.  Velocity controls that are enabled in Commando Mode are inactive until a Commando Mode event occurs. When Commando Mode velocity controls are activated, they conform to the `velocity_window` specified in that velocity control. For example, a `velocity_window` of `DAY` is one calendar day starting at 00:00:00 UTC. If a Commando Mode event occurs at 11:59:59 UTC, the `DAY` window includes all transactions that occurred between 00:00:00 and 11:59:59 on that calendar day.
	VelocityControls []string `json:"velocity_controls,omitempty"`
}

// NewCommandoModeEnables instantiates a new CommandoModeEnables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandoModeEnables(programFundingSource string) *CommandoModeEnables {
	this := CommandoModeEnables{}
	var ignoreCardSuspendedState bool = false
	this.IgnoreCardSuspendedState = &ignoreCardSuspendedState
	this.ProgramFundingSource = programFundingSource
	var useCacheBalance bool = false
	this.UseCacheBalance = &useCacheBalance
	return &this
}

// NewCommandoModeEnablesWithDefaults instantiates a new CommandoModeEnables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandoModeEnablesWithDefaults() *CommandoModeEnables {
	this := CommandoModeEnables{}
	var ignoreCardSuspendedState bool = false
	this.IgnoreCardSuspendedState = &ignoreCardSuspendedState
	var useCacheBalance bool = false
	this.UseCacheBalance = &useCacheBalance
	return &this
}

// GetAuthControls returns the AuthControls field value if set, zero value otherwise.
func (o *CommandoModeEnables) GetAuthControls() []string {
	if o == nil || IsNil(o.AuthControls) {
		var ret []string
		return ret
	}
	return o.AuthControls
}

// GetAuthControlsOk returns a tuple with the AuthControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeEnables) GetAuthControlsOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthControls) {
		return nil, false
	}
	return o.AuthControls, true
}

// HasAuthControls returns a boolean if a field has been set.
func (o *CommandoModeEnables) HasAuthControls() bool {
	if o != nil && !IsNil(o.AuthControls) {
		return true
	}

	return false
}

// SetAuthControls gets a reference to the given []string and assigns it to the AuthControls field.
func (o *CommandoModeEnables) SetAuthControls(v []string) {
	o.AuthControls = v
}

// GetIgnoreCardSuspendedState returns the IgnoreCardSuspendedState field value if set, zero value otherwise.
func (o *CommandoModeEnables) GetIgnoreCardSuspendedState() bool {
	if o == nil || IsNil(o.IgnoreCardSuspendedState) {
		var ret bool
		return ret
	}
	return *o.IgnoreCardSuspendedState
}

// GetIgnoreCardSuspendedStateOk returns a tuple with the IgnoreCardSuspendedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeEnables) GetIgnoreCardSuspendedStateOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreCardSuspendedState) {
		return nil, false
	}
	return o.IgnoreCardSuspendedState, true
}

// HasIgnoreCardSuspendedState returns a boolean if a field has been set.
func (o *CommandoModeEnables) HasIgnoreCardSuspendedState() bool {
	if o != nil && !IsNil(o.IgnoreCardSuspendedState) {
		return true
	}

	return false
}

// SetIgnoreCardSuspendedState gets a reference to the given bool and assigns it to the IgnoreCardSuspendedState field.
func (o *CommandoModeEnables) SetIgnoreCardSuspendedState(v bool) {
	o.IgnoreCardSuspendedState = &v
}

// GetProgramFundingSource returns the ProgramFundingSource field value
func (o *CommandoModeEnables) GetProgramFundingSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProgramFundingSource
}

// GetProgramFundingSourceOk returns a tuple with the ProgramFundingSource field value
// and a boolean to check if the value has been set.
func (o *CommandoModeEnables) GetProgramFundingSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramFundingSource, true
}

// SetProgramFundingSource sets field value
func (o *CommandoModeEnables) SetProgramFundingSource(v string) {
	o.ProgramFundingSource = v
}

// GetUseCacheBalance returns the UseCacheBalance field value if set, zero value otherwise.
func (o *CommandoModeEnables) GetUseCacheBalance() bool {
	if o == nil || IsNil(o.UseCacheBalance) {
		var ret bool
		return ret
	}
	return *o.UseCacheBalance
}

// GetUseCacheBalanceOk returns a tuple with the UseCacheBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeEnables) GetUseCacheBalanceOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCacheBalance) {
		return nil, false
	}
	return o.UseCacheBalance, true
}

// HasUseCacheBalance returns a boolean if a field has been set.
func (o *CommandoModeEnables) HasUseCacheBalance() bool {
	if o != nil && !IsNil(o.UseCacheBalance) {
		return true
	}

	return false
}

// SetUseCacheBalance gets a reference to the given bool and assigns it to the UseCacheBalance field.
func (o *CommandoModeEnables) SetUseCacheBalance(v bool) {
	o.UseCacheBalance = &v
}

// GetVelocityControls returns the VelocityControls field value if set, zero value otherwise.
func (o *CommandoModeEnables) GetVelocityControls() []string {
	if o == nil || IsNil(o.VelocityControls) {
		var ret []string
		return ret
	}
	return o.VelocityControls
}

// GetVelocityControlsOk returns a tuple with the VelocityControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeEnables) GetVelocityControlsOk() ([]string, bool) {
	if o == nil || IsNil(o.VelocityControls) {
		return nil, false
	}
	return o.VelocityControls, true
}

// HasVelocityControls returns a boolean if a field has been set.
func (o *CommandoModeEnables) HasVelocityControls() bool {
	if o != nil && !IsNil(o.VelocityControls) {
		return true
	}

	return false
}

// SetVelocityControls gets a reference to the given []string and assigns it to the VelocityControls field.
func (o *CommandoModeEnables) SetVelocityControls(v []string) {
	o.VelocityControls = v
}

func (o CommandoModeEnables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandoModeEnables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthControls) {
		toSerialize["auth_controls"] = o.AuthControls
	}
	if !IsNil(o.IgnoreCardSuspendedState) {
		toSerialize["ignore_card_suspended_state"] = o.IgnoreCardSuspendedState
	}
	toSerialize["program_funding_source"] = o.ProgramFundingSource
	if !IsNil(o.UseCacheBalance) {
		toSerialize["use_cache_balance"] = o.UseCacheBalance
	}
	if !IsNil(o.VelocityControls) {
		toSerialize["velocity_controls"] = o.VelocityControls
	}
	return toSerialize, nil
}

type NullableCommandoModeEnables struct {
	value *CommandoModeEnables
	isSet bool
}

func (v NullableCommandoModeEnables) Get() *CommandoModeEnables {
	return v.value
}

func (v *NullableCommandoModeEnables) Set(val *CommandoModeEnables) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandoModeEnables) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandoModeEnables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandoModeEnables(val *CommandoModeEnables) *NullableCommandoModeEnables {
	return &NullableCommandoModeEnables{value: val, isSet: true}
}

func (v NullableCommandoModeEnables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandoModeEnables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


