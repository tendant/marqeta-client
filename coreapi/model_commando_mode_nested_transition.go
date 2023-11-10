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

// checks if the CommandoModeNestedTransition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandoModeNestedTransition{}

// CommandoModeNestedTransition Describes the Commando Mode control set's `current_state` object.
type CommandoModeNestedTransition struct {
	// Mechanism that changed the Commando Mode control set's state.
	Channel string `json:"channel"`
	// Indicates whether Commando Mode is enabled.  * If `commando_enabled` is `true` and `COMMANDO_MANUAL` is configured, all transactions are processed via Commando Mode. * If `commando_enabled` is `true` and `COMMANDO_AUTO` is configured, Commando Mode is ready to intervene only when a transaction times out or encounters an error.
	CommandoEnabled bool `json:"commando_enabled"`
	// Describes the reason why the current state of the Commando Mode control set was last changed.
	Reason *string `json:"reason,omitempty"`
	// Identifies the user who last changed the Commando Mode control set's state.
	Username *string `json:"username,omitempty"`
}

// NewCommandoModeNestedTransition instantiates a new CommandoModeNestedTransition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandoModeNestedTransition(channel string, commandoEnabled bool) *CommandoModeNestedTransition {
	this := CommandoModeNestedTransition{}
	this.Channel = channel
	this.CommandoEnabled = commandoEnabled
	return &this
}

// NewCommandoModeNestedTransitionWithDefaults instantiates a new CommandoModeNestedTransition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandoModeNestedTransitionWithDefaults() *CommandoModeNestedTransition {
	this := CommandoModeNestedTransition{}
	var commandoEnabled bool = false
	this.CommandoEnabled = commandoEnabled
	return &this
}

// GetChannel returns the Channel field value
func (o *CommandoModeNestedTransition) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *CommandoModeNestedTransition) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *CommandoModeNestedTransition) SetChannel(v string) {
	o.Channel = v
}

// GetCommandoEnabled returns the CommandoEnabled field value
func (o *CommandoModeNestedTransition) GetCommandoEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CommandoEnabled
}

// GetCommandoEnabledOk returns a tuple with the CommandoEnabled field value
// and a boolean to check if the value has been set.
func (o *CommandoModeNestedTransition) GetCommandoEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandoEnabled, true
}

// SetCommandoEnabled sets field value
func (o *CommandoModeNestedTransition) SetCommandoEnabled(v bool) {
	o.CommandoEnabled = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CommandoModeNestedTransition) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeNestedTransition) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CommandoModeNestedTransition) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CommandoModeNestedTransition) SetReason(v string) {
	o.Reason = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CommandoModeNestedTransition) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandoModeNestedTransition) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CommandoModeNestedTransition) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CommandoModeNestedTransition) SetUsername(v string) {
	o.Username = &v
}

func (o CommandoModeNestedTransition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandoModeNestedTransition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["commando_enabled"] = o.CommandoEnabled
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCommandoModeNestedTransition struct {
	value *CommandoModeNestedTransition
	isSet bool
}

func (v NullableCommandoModeNestedTransition) Get() *CommandoModeNestedTransition {
	return v.value
}

func (v *NullableCommandoModeNestedTransition) Set(val *CommandoModeNestedTransition) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandoModeNestedTransition) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandoModeNestedTransition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandoModeNestedTransition(val *CommandoModeNestedTransition) *NullableCommandoModeNestedTransition {
	return &NullableCommandoModeNestedTransition{value: val, isSet: true}
}

func (v NullableCommandoModeNestedTransition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandoModeNestedTransition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

