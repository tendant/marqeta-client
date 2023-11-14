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

// checks if the ResetUserPasswordModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetUserPasswordModel{}

// ResetUserPasswordModel struct for ResetUserPasswordModel
type ResetUserPasswordModel struct {
	// New password to the cardholder's user account on the Marqeta platform.
	NewPassword string `json:"new_password"`
	// Unique identifier of the cardholder.
	UserToken string `json:"user_token"`
}

// NewResetUserPasswordModel instantiates a new ResetUserPasswordModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetUserPasswordModel(newPassword string, userToken string) *ResetUserPasswordModel {
	this := ResetUserPasswordModel{}
	this.NewPassword = newPassword
	this.UserToken = userToken
	return &this
}

// NewResetUserPasswordModelWithDefaults instantiates a new ResetUserPasswordModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetUserPasswordModelWithDefaults() *ResetUserPasswordModel {
	this := ResetUserPasswordModel{}
	return &this
}

// GetNewPassword returns the NewPassword field value
func (o *ResetUserPasswordModel) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *ResetUserPasswordModel) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *ResetUserPasswordModel) SetNewPassword(v string) {
	o.NewPassword = v
}

// GetUserToken returns the UserToken field value
func (o *ResetUserPasswordModel) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ResetUserPasswordModel) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *ResetUserPasswordModel) SetUserToken(v string) {
	o.UserToken = v
}

func (o ResetUserPasswordModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetUserPasswordModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["new_password"] = o.NewPassword
	toSerialize["user_token"] = o.UserToken
	return toSerialize, nil
}

type NullableResetUserPasswordModel struct {
	value *ResetUserPasswordModel
	isSet bool
}

func (v NullableResetUserPasswordModel) Get() *ResetUserPasswordModel {
	return v.value
}

func (v *NullableResetUserPasswordModel) Set(val *ResetUserPasswordModel) {
	v.value = val
	v.isSet = true
}

func (v NullableResetUserPasswordModel) IsSet() bool {
	return v.isSet
}

func (v *NullableResetUserPasswordModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetUserPasswordModel(val *ResetUserPasswordModel) *NullableResetUserPasswordModel {
	return &NullableResetUserPasswordModel{value: val, isSet: true}
}

func (v NullableResetUserPasswordModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetUserPasswordModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


