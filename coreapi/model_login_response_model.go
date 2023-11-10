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

// checks if the LoginResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginResponseModel{}

// LoginResponseModel struct for LoginResponseModel
type LoginResponseModel struct {
	AccessToken *AccessTokenResponse `json:"access_token,omitempty"`
	User *UserCardHolderResponse `json:"user,omitempty"`
}

// NewLoginResponseModel instantiates a new LoginResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponseModel() *LoginResponseModel {
	this := LoginResponseModel{}
	return &this
}

// NewLoginResponseModelWithDefaults instantiates a new LoginResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseModelWithDefaults() *LoginResponseModel {
	this := LoginResponseModel{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *LoginResponseModel) GetAccessToken() AccessTokenResponse {
	if o == nil || IsNil(o.AccessToken) {
		var ret AccessTokenResponse
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseModel) GetAccessTokenOk() (*AccessTokenResponse, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *LoginResponseModel) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given AccessTokenResponse and assigns it to the AccessToken field.
func (o *LoginResponseModel) SetAccessToken(v AccessTokenResponse) {
	o.AccessToken = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LoginResponseModel) GetUser() UserCardHolderResponse {
	if o == nil || IsNil(o.User) {
		var ret UserCardHolderResponse
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponseModel) GetUserOk() (*UserCardHolderResponse, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LoginResponseModel) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserCardHolderResponse and assigns it to the User field.
func (o *LoginResponseModel) SetUser(v UserCardHolderResponse) {
	o.User = &v
}

func (o LoginResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableLoginResponseModel struct {
	value *LoginResponseModel
	isSet bool
}

func (v NullableLoginResponseModel) Get() *LoginResponseModel {
	return v.value
}

func (v *NullableLoginResponseModel) Set(val *LoginResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponseModel(val *LoginResponseModel) *NullableLoginResponseModel {
	return &NullableLoginResponseModel{value: val, isSet: true}
}

func (v NullableLoginResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


