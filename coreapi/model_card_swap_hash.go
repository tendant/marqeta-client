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

// checks if the CardSwapHash type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardSwapHash{}

// CardSwapHash Contains identifiers for swapping digital wallet tokens between cards.
type CardSwapHash struct {
	// Unique identifier of the new card resource to which the digital wallet tokens are assigned.
	NewCardToken string `json:"new_card_token"`
	// Unique identifier of the existing card resource that has digital wallet tokens assigned to it.
	PreviousCardToken string `json:"previous_card_token"`
}

// NewCardSwapHash instantiates a new CardSwapHash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardSwapHash(newCardToken string, previousCardToken string) *CardSwapHash {
	this := CardSwapHash{}
	this.NewCardToken = newCardToken
	this.PreviousCardToken = previousCardToken
	return &this
}

// NewCardSwapHashWithDefaults instantiates a new CardSwapHash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardSwapHashWithDefaults() *CardSwapHash {
	this := CardSwapHash{}
	return &this
}

// GetNewCardToken returns the NewCardToken field value
func (o *CardSwapHash) GetNewCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewCardToken
}

// GetNewCardTokenOk returns a tuple with the NewCardToken field value
// and a boolean to check if the value has been set.
func (o *CardSwapHash) GetNewCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewCardToken, true
}

// SetNewCardToken sets field value
func (o *CardSwapHash) SetNewCardToken(v string) {
	o.NewCardToken = v
}

// GetPreviousCardToken returns the PreviousCardToken field value
func (o *CardSwapHash) GetPreviousCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviousCardToken
}

// GetPreviousCardTokenOk returns a tuple with the PreviousCardToken field value
// and a boolean to check if the value has been set.
func (o *CardSwapHash) GetPreviousCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousCardToken, true
}

// SetPreviousCardToken sets field value
func (o *CardSwapHash) SetPreviousCardToken(v string) {
	o.PreviousCardToken = v
}

func (o CardSwapHash) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardSwapHash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["new_card_token"] = o.NewCardToken
	toSerialize["previous_card_token"] = o.PreviousCardToken
	return toSerialize, nil
}

type NullableCardSwapHash struct {
	value *CardSwapHash
	isSet bool
}

func (v NullableCardSwapHash) Get() *CardSwapHash {
	return v.value
}

func (v *NullableCardSwapHash) Set(val *CardSwapHash) {
	v.value = val
	v.isSet = true
}

func (v NullableCardSwapHash) IsSet() bool {
	return v.isSet
}

func (v *NullableCardSwapHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardSwapHash(val *CardSwapHash) *NullableCardSwapHash {
	return &NullableCardSwapHash{value: val, isSet: true}
}

func (v NullableCardSwapHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardSwapHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


