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

// checks if the DirectDepositAccountTransitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectDepositAccountTransitionRequest{}

// DirectDepositAccountTransitionRequest struct for DirectDepositAccountTransitionRequest
type DirectDepositAccountTransitionRequest struct {
	AccountToken string `json:"account_token"`
	Channel string `json:"channel"`
	Reason *string `json:"reason,omitempty"`
	State *string `json:"state,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewDirectDepositAccountTransitionRequest instantiates a new DirectDepositAccountTransitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDepositAccountTransitionRequest(accountToken string, channel string) *DirectDepositAccountTransitionRequest {
	this := DirectDepositAccountTransitionRequest{}
	this.AccountToken = accountToken
	this.Channel = channel
	return &this
}

// NewDirectDepositAccountTransitionRequestWithDefaults instantiates a new DirectDepositAccountTransitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDepositAccountTransitionRequestWithDefaults() *DirectDepositAccountTransitionRequest {
	this := DirectDepositAccountTransitionRequest{}
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *DirectDepositAccountTransitionRequest) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountTransitionRequest) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *DirectDepositAccountTransitionRequest) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetChannel returns the Channel field value
func (o *DirectDepositAccountTransitionRequest) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountTransitionRequest) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *DirectDepositAccountTransitionRequest) SetChannel(v string) {
	o.Channel = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DirectDepositAccountTransitionRequest) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountTransitionRequest) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DirectDepositAccountTransitionRequest) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DirectDepositAccountTransitionRequest) SetReason(v string) {
	o.Reason = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DirectDepositAccountTransitionRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountTransitionRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DirectDepositAccountTransitionRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DirectDepositAccountTransitionRequest) SetState(v string) {
	o.State = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DirectDepositAccountTransitionRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountTransitionRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DirectDepositAccountTransitionRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DirectDepositAccountTransitionRequest) SetToken(v string) {
	o.Token = &v
}

func (o DirectDepositAccountTransitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDepositAccountTransitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_token"] = o.AccountToken
	toSerialize["channel"] = o.Channel
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableDirectDepositAccountTransitionRequest struct {
	value *DirectDepositAccountTransitionRequest
	isSet bool
}

func (v NullableDirectDepositAccountTransitionRequest) Get() *DirectDepositAccountTransitionRequest {
	return v.value
}

func (v *NullableDirectDepositAccountTransitionRequest) Set(val *DirectDepositAccountTransitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDepositAccountTransitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDepositAccountTransitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDepositAccountTransitionRequest(val *DirectDepositAccountTransitionRequest) *NullableDirectDepositAccountTransitionRequest {
	return &NullableDirectDepositAccountTransitionRequest{value: val, isSet: true}
}

func (v NullableDirectDepositAccountTransitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDepositAccountTransitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

