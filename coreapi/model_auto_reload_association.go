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

// checks if the AutoReloadAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoReloadAssociation{}

// AutoReloadAssociation Specifies the scope of the auto reload.  Input no more than one field. If no value is supplied, the auto reload applies at the program level.
type AutoReloadAssociation struct {
	// Unique identifier of the business for which the auto reload is configured.  Send a `GET` request to `/businesses` to retrieve business tokens.
	BusinessToken *string `json:"business_token,omitempty"`
	// Unique identifier of the card product for which the auto reload is configured.  Send a `GET` request to `/cardproducts` to retrieve card product tokens.
	CardProductToken *string `json:"card_product_token,omitempty"`
	// Unique identifier of the user for which the auto reload is configured.  Send a `GET` request to `/users` to retrieve user tokens.
	UserToken *string `json:"user_token,omitempty"`
}

// NewAutoReloadAssociation instantiates a new AutoReloadAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoReloadAssociation() *AutoReloadAssociation {
	this := AutoReloadAssociation{}
	return &this
}

// NewAutoReloadAssociationWithDefaults instantiates a new AutoReloadAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoReloadAssociationWithDefaults() *AutoReloadAssociation {
	this := AutoReloadAssociation{}
	return &this
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *AutoReloadAssociation) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadAssociation) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *AutoReloadAssociation) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *AutoReloadAssociation) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetCardProductToken returns the CardProductToken field value if set, zero value otherwise.
func (o *AutoReloadAssociation) GetCardProductToken() string {
	if o == nil || IsNil(o.CardProductToken) {
		var ret string
		return ret
	}
	return *o.CardProductToken
}

// GetCardProductTokenOk returns a tuple with the CardProductToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadAssociation) GetCardProductTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CardProductToken) {
		return nil, false
	}
	return o.CardProductToken, true
}

// HasCardProductToken returns a boolean if a field has been set.
func (o *AutoReloadAssociation) HasCardProductToken() bool {
	if o != nil && !IsNil(o.CardProductToken) {
		return true
	}

	return false
}

// SetCardProductToken gets a reference to the given string and assigns it to the CardProductToken field.
func (o *AutoReloadAssociation) SetCardProductToken(v string) {
	o.CardProductToken = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *AutoReloadAssociation) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadAssociation) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *AutoReloadAssociation) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *AutoReloadAssociation) SetUserToken(v string) {
	o.UserToken = &v
}

func (o AutoReloadAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoReloadAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	if !IsNil(o.CardProductToken) {
		toSerialize["card_product_token"] = o.CardProductToken
	}
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

type NullableAutoReloadAssociation struct {
	value *AutoReloadAssociation
	isSet bool
}

func (v NullableAutoReloadAssociation) Get() *AutoReloadAssociation {
	return v.value
}

func (v *NullableAutoReloadAssociation) Set(val *AutoReloadAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoReloadAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoReloadAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoReloadAssociation(val *AutoReloadAssociation) *NullableAutoReloadAssociation {
	return &NullableAutoReloadAssociation{value: val, isSet: true}
}

func (v NullableAutoReloadAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoReloadAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


