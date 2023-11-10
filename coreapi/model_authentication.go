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

// checks if the Authentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Authentication{}

// Authentication Contains the cardholder's email address and password information.
type Authentication struct {
	// Specifies whether the email address has been verified.
	EmailVerified *bool `json:"email_verified,omitempty"`
	// Date and time when the email address was verified.
	EmailVerifiedTime *time.Time `json:"email_verified_time,omitempty"`
	// Specifies the channel through which the password was last changed.
	LastPasswordUpdateChannel *string `json:"last_password_update_channel,omitempty"`
	// Date and time when the password was last changed.
	LastPasswordUpdateTime *time.Time `json:"last_password_update_time,omitempty"`
}

// NewAuthentication instantiates a new Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthentication() *Authentication {
	this := Authentication{}
	var emailVerified bool = false
	this.EmailVerified = &emailVerified
	return &this
}

// NewAuthenticationWithDefaults instantiates a new Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationWithDefaults() *Authentication {
	this := Authentication{}
	var emailVerified bool = false
	this.EmailVerified = &emailVerified
	return &this
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *Authentication) GetEmailVerified() bool {
	if o == nil || IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *Authentication) HasEmailVerified() bool {
	if o != nil && !IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *Authentication) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetEmailVerifiedTime returns the EmailVerifiedTime field value if set, zero value otherwise.
func (o *Authentication) GetEmailVerifiedTime() time.Time {
	if o == nil || IsNil(o.EmailVerifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.EmailVerifiedTime
}

// GetEmailVerifiedTimeOk returns a tuple with the EmailVerifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetEmailVerifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EmailVerifiedTime) {
		return nil, false
	}
	return o.EmailVerifiedTime, true
}

// HasEmailVerifiedTime returns a boolean if a field has been set.
func (o *Authentication) HasEmailVerifiedTime() bool {
	if o != nil && !IsNil(o.EmailVerifiedTime) {
		return true
	}

	return false
}

// SetEmailVerifiedTime gets a reference to the given time.Time and assigns it to the EmailVerifiedTime field.
func (o *Authentication) SetEmailVerifiedTime(v time.Time) {
	o.EmailVerifiedTime = &v
}

// GetLastPasswordUpdateChannel returns the LastPasswordUpdateChannel field value if set, zero value otherwise.
func (o *Authentication) GetLastPasswordUpdateChannel() string {
	if o == nil || IsNil(o.LastPasswordUpdateChannel) {
		var ret string
		return ret
	}
	return *o.LastPasswordUpdateChannel
}

// GetLastPasswordUpdateChannelOk returns a tuple with the LastPasswordUpdateChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetLastPasswordUpdateChannelOk() (*string, bool) {
	if o == nil || IsNil(o.LastPasswordUpdateChannel) {
		return nil, false
	}
	return o.LastPasswordUpdateChannel, true
}

// HasLastPasswordUpdateChannel returns a boolean if a field has been set.
func (o *Authentication) HasLastPasswordUpdateChannel() bool {
	if o != nil && !IsNil(o.LastPasswordUpdateChannel) {
		return true
	}

	return false
}

// SetLastPasswordUpdateChannel gets a reference to the given string and assigns it to the LastPasswordUpdateChannel field.
func (o *Authentication) SetLastPasswordUpdateChannel(v string) {
	o.LastPasswordUpdateChannel = &v
}

// GetLastPasswordUpdateTime returns the LastPasswordUpdateTime field value if set, zero value otherwise.
func (o *Authentication) GetLastPasswordUpdateTime() time.Time {
	if o == nil || IsNil(o.LastPasswordUpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordUpdateTime
}

// GetLastPasswordUpdateTimeOk returns a tuple with the LastPasswordUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetLastPasswordUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPasswordUpdateTime) {
		return nil, false
	}
	return o.LastPasswordUpdateTime, true
}

// HasLastPasswordUpdateTime returns a boolean if a field has been set.
func (o *Authentication) HasLastPasswordUpdateTime() bool {
	if o != nil && !IsNil(o.LastPasswordUpdateTime) {
		return true
	}

	return false
}

// SetLastPasswordUpdateTime gets a reference to the given time.Time and assigns it to the LastPasswordUpdateTime field.
func (o *Authentication) SetLastPasswordUpdateTime(v time.Time) {
	o.LastPasswordUpdateTime = &v
}

func (o Authentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Authentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailVerified) {
		toSerialize["email_verified"] = o.EmailVerified
	}
	if !IsNil(o.EmailVerifiedTime) {
		toSerialize["email_verified_time"] = o.EmailVerifiedTime
	}
	if !IsNil(o.LastPasswordUpdateChannel) {
		toSerialize["last_password_update_channel"] = o.LastPasswordUpdateChannel
	}
	if !IsNil(o.LastPasswordUpdateTime) {
		toSerialize["last_password_update_time"] = o.LastPasswordUpdateTime
	}
	return toSerialize, nil
}

type NullableAuthentication struct {
	value *Authentication
	isSet bool
}

func (v NullableAuthentication) Get() *Authentication {
	return v.value
}

func (v *NullableAuthentication) Set(val *Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication(val *Authentication) *NullableAuthentication {
	return &NullableAuthentication{value: val, isSet: true}
}

func (v NullableAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


