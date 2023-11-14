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

// checks if the DigitalWalletTokenMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletTokenMetadata{}

// DigitalWalletTokenMetadata Contains additional information about the digital wallet token.
type DigitalWalletTokenMetadata struct {
	// Language specified in the `config.transaction_controls.notification_language` field of the card product: `ces` (Czech), `deu` (German), `eng` (English), `fra` (French), `ita` (Italian), `pol` (Polish), `spa` (Spanish), `swe` (Swedish)  The ISO maintains the link:https://www.iso.org/iso-3166-country-codes.html[full list of ISO 3166 two- and three-digit numeric country codes, window=\"_blank\"].
	CardproductPreferredNotificationLanguage *string `json:"cardproduct_preferred_notification_language,omitempty"`
	// Unique identifier of the product configuration on the Marqeta platform.
	IssuerProductConfigId *string `json:"issuer_product_config_id,omitempty"`
}

// NewDigitalWalletTokenMetadata instantiates a new DigitalWalletTokenMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletTokenMetadata() *DigitalWalletTokenMetadata {
	this := DigitalWalletTokenMetadata{}
	return &this
}

// NewDigitalWalletTokenMetadataWithDefaults instantiates a new DigitalWalletTokenMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenMetadataWithDefaults() *DigitalWalletTokenMetadata {
	this := DigitalWalletTokenMetadata{}
	return &this
}

// GetCardproductPreferredNotificationLanguage returns the CardproductPreferredNotificationLanguage field value if set, zero value otherwise.
func (o *DigitalWalletTokenMetadata) GetCardproductPreferredNotificationLanguage() string {
	if o == nil || IsNil(o.CardproductPreferredNotificationLanguage) {
		var ret string
		return ret
	}
	return *o.CardproductPreferredNotificationLanguage
}

// GetCardproductPreferredNotificationLanguageOk returns a tuple with the CardproductPreferredNotificationLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenMetadata) GetCardproductPreferredNotificationLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.CardproductPreferredNotificationLanguage) {
		return nil, false
	}
	return o.CardproductPreferredNotificationLanguage, true
}

// HasCardproductPreferredNotificationLanguage returns a boolean if a field has been set.
func (o *DigitalWalletTokenMetadata) HasCardproductPreferredNotificationLanguage() bool {
	if o != nil && !IsNil(o.CardproductPreferredNotificationLanguage) {
		return true
	}

	return false
}

// SetCardproductPreferredNotificationLanguage gets a reference to the given string and assigns it to the CardproductPreferredNotificationLanguage field.
func (o *DigitalWalletTokenMetadata) SetCardproductPreferredNotificationLanguage(v string) {
	o.CardproductPreferredNotificationLanguage = &v
}

// GetIssuerProductConfigId returns the IssuerProductConfigId field value if set, zero value otherwise.
func (o *DigitalWalletTokenMetadata) GetIssuerProductConfigId() string {
	if o == nil || IsNil(o.IssuerProductConfigId) {
		var ret string
		return ret
	}
	return *o.IssuerProductConfigId
}

// GetIssuerProductConfigIdOk returns a tuple with the IssuerProductConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenMetadata) GetIssuerProductConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerProductConfigId) {
		return nil, false
	}
	return o.IssuerProductConfigId, true
}

// HasIssuerProductConfigId returns a boolean if a field has been set.
func (o *DigitalWalletTokenMetadata) HasIssuerProductConfigId() bool {
	if o != nil && !IsNil(o.IssuerProductConfigId) {
		return true
	}

	return false
}

// SetIssuerProductConfigId gets a reference to the given string and assigns it to the IssuerProductConfigId field.
func (o *DigitalWalletTokenMetadata) SetIssuerProductConfigId(v string) {
	o.IssuerProductConfigId = &v
}

func (o DigitalWalletTokenMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletTokenMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardproductPreferredNotificationLanguage) {
		toSerialize["cardproduct_preferred_notification_language"] = o.CardproductPreferredNotificationLanguage
	}
	if !IsNil(o.IssuerProductConfigId) {
		toSerialize["issuer_product_config_id"] = o.IssuerProductConfigId
	}
	return toSerialize, nil
}

type NullableDigitalWalletTokenMetadata struct {
	value *DigitalWalletTokenMetadata
	isSet bool
}

func (v NullableDigitalWalletTokenMetadata) Get() *DigitalWalletTokenMetadata {
	return v.value
}

func (v *NullableDigitalWalletTokenMetadata) Set(val *DigitalWalletTokenMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenMetadata(val *DigitalWalletTokenMetadata) *NullableDigitalWalletTokenMetadata {
	return &NullableDigitalWalletTokenMetadata{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


