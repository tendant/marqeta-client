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

// checks if the SpendControlAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendControlAssociation{}

// SpendControlAssociation Defines the group of users to which the velocity control applies.
type SpendControlAssociation struct {
	// Unique identifier of the card product.  Pass either `card_product_token` or `user_token`, not both.
	CardProductToken *string `json:"card_product_token,omitempty"`
	// Unique identifier of the cardholder.  Pass either `card_product_token` or `user_token`, not both.
	UserToken *string `json:"user_token,omitempty"`
}

// NewSpendControlAssociation instantiates a new SpendControlAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendControlAssociation() *SpendControlAssociation {
	this := SpendControlAssociation{}
	return &this
}

// NewSpendControlAssociationWithDefaults instantiates a new SpendControlAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendControlAssociationWithDefaults() *SpendControlAssociation {
	this := SpendControlAssociation{}
	return &this
}

// GetCardProductToken returns the CardProductToken field value if set, zero value otherwise.
func (o *SpendControlAssociation) GetCardProductToken() string {
	if o == nil || IsNil(o.CardProductToken) {
		var ret string
		return ret
	}
	return *o.CardProductToken
}

// GetCardProductTokenOk returns a tuple with the CardProductToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControlAssociation) GetCardProductTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CardProductToken) {
		return nil, false
	}
	return o.CardProductToken, true
}

// HasCardProductToken returns a boolean if a field has been set.
func (o *SpendControlAssociation) HasCardProductToken() bool {
	if o != nil && !IsNil(o.CardProductToken) {
		return true
	}

	return false
}

// SetCardProductToken gets a reference to the given string and assigns it to the CardProductToken field.
func (o *SpendControlAssociation) SetCardProductToken(v string) {
	o.CardProductToken = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *SpendControlAssociation) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControlAssociation) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *SpendControlAssociation) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *SpendControlAssociation) SetUserToken(v string) {
	o.UserToken = &v
}

func (o SpendControlAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendControlAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardProductToken) {
		toSerialize["card_product_token"] = o.CardProductToken
	}
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

type NullableSpendControlAssociation struct {
	value *SpendControlAssociation
	isSet bool
}

func (v NullableSpendControlAssociation) Get() *SpendControlAssociation {
	return v.value
}

func (v *NullableSpendControlAssociation) Set(val *SpendControlAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendControlAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendControlAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendControlAssociation(val *SpendControlAssociation) *NullableSpendControlAssociation {
	return &NullableSpendControlAssociation{value: val, isSet: true}
}

func (v NullableSpendControlAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendControlAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

