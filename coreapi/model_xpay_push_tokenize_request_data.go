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

// checks if the XpayPushTokenizeRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XpayPushTokenizeRequestData{}

// XpayPushTokenizeRequestData Contains details about a card tokenization push request.
type XpayPushTokenizeRequestData struct {
	// Specifies the type of card.
	CardType *string `json:"card_type,omitempty"`
	// Name of the card as displayed in the digital wallet, typically showing the card brand and last four digits of the primary account number (PAN). `Visa 5678`, for example.
	DisplayName *string `json:"display_name,omitempty"`
	// Encrypted card or cardholder details.
	ExtraProvisionPayload *string `json:"extra_provision_payload,omitempty"`
	// Last four digits of the primary account number of the physical or virtual card.
	LastDigits *string `json:"last_digits,omitempty"`
	// Specifies the card network of the physical or virtual card.
	Network *string `json:"network,omitempty"`
	// Specifies the network that provides the digital wallet token service.
	TokenServiceProvider *string `json:"token_service_provider,omitempty"`
}

// NewXpayPushTokenizeRequestData instantiates a new XpayPushTokenizeRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXpayPushTokenizeRequestData() *XpayPushTokenizeRequestData {
	this := XpayPushTokenizeRequestData{}
	return &this
}

// NewXpayPushTokenizeRequestDataWithDefaults instantiates a new XpayPushTokenizeRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXpayPushTokenizeRequestDataWithDefaults() *XpayPushTokenizeRequestData {
	this := XpayPushTokenizeRequestData{}
	return &this
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *XpayPushTokenizeRequestData) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpayPushTokenizeRequestData) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *XpayPushTokenizeRequestData) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *XpayPushTokenizeRequestData) SetCardType(v string) {
	o.CardType = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *XpayPushTokenizeRequestData) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpayPushTokenizeRequestData) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *XpayPushTokenizeRequestData) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *XpayPushTokenizeRequestData) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtraProvisionPayload returns the ExtraProvisionPayload field value if set, zero value otherwise.
func (o *XpayPushTokenizeRequestData) GetExtraProvisionPayload() string {
	if o == nil || IsNil(o.ExtraProvisionPayload) {
		var ret string
		return ret
	}
	return *o.ExtraProvisionPayload
}

// GetExtraProvisionPayloadOk returns a tuple with the ExtraProvisionPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpayPushTokenizeRequestData) GetExtraProvisionPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraProvisionPayload) {
		return nil, false
	}
	return o.ExtraProvisionPayload, true
}

// HasExtraProvisionPayload returns a boolean if a field has been set.
func (o *XpayPushTokenizeRequestData) HasExtraProvisionPayload() bool {
	if o != nil && !IsNil(o.ExtraProvisionPayload) {
		return true
	}

	return false
}

// SetExtraProvisionPayload gets a reference to the given string and assigns it to the ExtraProvisionPayload field.
func (o *XpayPushTokenizeRequestData) SetExtraProvisionPayload(v string) {
	o.ExtraProvisionPayload = &v
}

// GetLastDigits returns the LastDigits field value if set, zero value otherwise.
func (o *XpayPushTokenizeRequestData) GetLastDigits() string {
	if o == nil || IsNil(o.LastDigits) {
		var ret string
		return ret
	}
	return *o.LastDigits
}

// GetLastDigitsOk returns a tuple with the LastDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpayPushTokenizeRequestData) GetLastDigitsOk() (*string, bool) {
	if o == nil || IsNil(o.LastDigits) {
		return nil, false
	}
	return o.LastDigits, true
}

// HasLastDigits returns a boolean if a field has been set.
func (o *XpayPushTokenizeRequestData) HasLastDigits() bool {
	if o != nil && !IsNil(o.LastDigits) {
		return true
	}

	return false
}

// SetLastDigits gets a reference to the given string and assigns it to the LastDigits field.
func (o *XpayPushTokenizeRequestData) SetLastDigits(v string) {
	o.LastDigits = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *XpayPushTokenizeRequestData) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpayPushTokenizeRequestData) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *XpayPushTokenizeRequestData) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *XpayPushTokenizeRequestData) SetNetwork(v string) {
	o.Network = &v
}

// GetTokenServiceProvider returns the TokenServiceProvider field value if set, zero value otherwise.
func (o *XpayPushTokenizeRequestData) GetTokenServiceProvider() string {
	if o == nil || IsNil(o.TokenServiceProvider) {
		var ret string
		return ret
	}
	return *o.TokenServiceProvider
}

// GetTokenServiceProviderOk returns a tuple with the TokenServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpayPushTokenizeRequestData) GetTokenServiceProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TokenServiceProvider) {
		return nil, false
	}
	return o.TokenServiceProvider, true
}

// HasTokenServiceProvider returns a boolean if a field has been set.
func (o *XpayPushTokenizeRequestData) HasTokenServiceProvider() bool {
	if o != nil && !IsNil(o.TokenServiceProvider) {
		return true
	}

	return false
}

// SetTokenServiceProvider gets a reference to the given string and assigns it to the TokenServiceProvider field.
func (o *XpayPushTokenizeRequestData) SetTokenServiceProvider(v string) {
	o.TokenServiceProvider = &v
}

func (o XpayPushTokenizeRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XpayPushTokenizeRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.ExtraProvisionPayload) {
		toSerialize["extra_provision_payload"] = o.ExtraProvisionPayload
	}
	if !IsNil(o.LastDigits) {
		toSerialize["last_digits"] = o.LastDigits
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.TokenServiceProvider) {
		toSerialize["token_service_provider"] = o.TokenServiceProvider
	}
	return toSerialize, nil
}

type NullableXpayPushTokenizeRequestData struct {
	value *XpayPushTokenizeRequestData
	isSet bool
}

func (v NullableXpayPushTokenizeRequestData) Get() *XpayPushTokenizeRequestData {
	return v.value
}

func (v *NullableXpayPushTokenizeRequestData) Set(val *XpayPushTokenizeRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableXpayPushTokenizeRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableXpayPushTokenizeRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXpayPushTokenizeRequestData(val *XpayPushTokenizeRequestData) *NullableXpayPushTokenizeRequestData {
	return &NullableXpayPushTokenizeRequestData{value: val, isSet: true}
}

func (v NullableXpayPushTokenizeRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXpayPushTokenizeRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


