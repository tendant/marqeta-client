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

// checks if the CardSecurityCodeVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardSecurityCodeVerification{}

// CardSecurityCodeVerification Contains information about a verification check performed on the card's security code.
type CardSecurityCodeVerification struct {
	Response Response `json:"response"`
	// Indicates the type of security code. Can have these possible values:  * *CVV1*  the security code stored in the magnetic stripe on the card. * *CVV2*  the security code printed on the card. * *ICVV*  the security code stored on the chip of the card. * *DCVV*  a dynamic security code used in some contactless payments when a card or device is tapped against the card reader.
	Type string `json:"type"`
}

// NewCardSecurityCodeVerification instantiates a new CardSecurityCodeVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardSecurityCodeVerification(response Response, type_ string) *CardSecurityCodeVerification {
	this := CardSecurityCodeVerification{}
	this.Response = response
	this.Type = type_
	return &this
}

// NewCardSecurityCodeVerificationWithDefaults instantiates a new CardSecurityCodeVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardSecurityCodeVerificationWithDefaults() *CardSecurityCodeVerification {
	this := CardSecurityCodeVerification{}
	return &this
}

// GetResponse returns the Response field value
func (o *CardSecurityCodeVerification) GetResponse() Response {
	if o == nil {
		var ret Response
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *CardSecurityCodeVerification) GetResponseOk() (*Response, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *CardSecurityCodeVerification) SetResponse(v Response) {
	o.Response = v
}

// GetType returns the Type field value
func (o *CardSecurityCodeVerification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CardSecurityCodeVerification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CardSecurityCodeVerification) SetType(v string) {
	o.Type = v
}

func (o CardSecurityCodeVerification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardSecurityCodeVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["response"] = o.Response
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCardSecurityCodeVerification struct {
	value *CardSecurityCodeVerification
	isSet bool
}

func (v NullableCardSecurityCodeVerification) Get() *CardSecurityCodeVerification {
	return v.value
}

func (v *NullableCardSecurityCodeVerification) Set(val *CardSecurityCodeVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableCardSecurityCodeVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableCardSecurityCodeVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardSecurityCodeVerification(val *CardSecurityCodeVerification) *NullableCardSecurityCodeVerification {
	return &NullableCardSecurityCodeVerification{value: val, isSet: true}
}

func (v NullableCardSecurityCodeVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardSecurityCodeVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


