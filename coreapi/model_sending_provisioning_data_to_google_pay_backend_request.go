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

// checks if the SendingProvisioningDataToGooglePayBackendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendingProvisioningDataToGooglePayBackendRequest{}

// SendingProvisioningDataToGooglePayBackendRequest struct for SendingProvisioningDataToGooglePayBackendRequest
type SendingProvisioningDataToGooglePayBackendRequest struct {
	// Indicates if the Funding Primary Account Number (FPAN) will be attempted.  * *1* - FPAN save will be attempted. * *0* - FPAN save will not be attempted.
	CardSetting int32 `json:"card_setting"`
	// Unique identifier of the card resource.
	CardToken string `json:"card_token"`
	// String provided by Google Wallet that identifies the client session.
	ClientSessionId string `json:"client_session_id"`
	// Google-assigned string that uniquely identifies both the integrator that is initiating the session and the issuer of the card.
	IntegratorId string `json:"integrator_id"`
	// String provided by Google Wallet that identifies the Android device that will receive the digital wallet token.
	PublicDeviceId string `json:"public_device_id"`
	// String provided by Google Wallet that identifies the device-scoped wallet that receives the digital wallet token.
	PublicWalletId string `json:"public_wallet_id"`
	// String provided by Google Wallet that identifies the backend session.
	ServerSessionId string `json:"server_session_id"`
	// Indicates if tokenization will be attempted.  * *1* - Tokenization will be attempted. * *0* - Tokenization will not be attempted.
	TokenSetting int32 `json:"token_setting"`
}

// NewSendingProvisioningDataToGooglePayBackendRequest instantiates a new SendingProvisioningDataToGooglePayBackendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendingProvisioningDataToGooglePayBackendRequest(cardSetting int32, cardToken string, clientSessionId string, integratorId string, publicDeviceId string, publicWalletId string, serverSessionId string, tokenSetting int32) *SendingProvisioningDataToGooglePayBackendRequest {
	this := SendingProvisioningDataToGooglePayBackendRequest{}
	this.CardSetting = cardSetting
	this.CardToken = cardToken
	this.ClientSessionId = clientSessionId
	this.IntegratorId = integratorId
	this.PublicDeviceId = publicDeviceId
	this.PublicWalletId = publicWalletId
	this.ServerSessionId = serverSessionId
	this.TokenSetting = tokenSetting
	return &this
}

// NewSendingProvisioningDataToGooglePayBackendRequestWithDefaults instantiates a new SendingProvisioningDataToGooglePayBackendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendingProvisioningDataToGooglePayBackendRequestWithDefaults() *SendingProvisioningDataToGooglePayBackendRequest {
	this := SendingProvisioningDataToGooglePayBackendRequest{}
	return &this
}

// GetCardSetting returns the CardSetting field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardSetting() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CardSetting
}

// GetCardSettingOk returns a tuple with the CardSetting field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardSettingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardSetting, true
}

// SetCardSetting sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetCardSetting(v int32) {
	o.CardSetting = v
}

// GetCardToken returns the CardToken field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetCardToken(v string) {
	o.CardToken = v
}

// GetClientSessionId returns the ClientSessionId field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetClientSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSessionId
}

// GetClientSessionIdOk returns a tuple with the ClientSessionId field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetClientSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSessionId, true
}

// SetClientSessionId sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetClientSessionId(v string) {
	o.ClientSessionId = v
}

// GetIntegratorId returns the IntegratorId field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetIntegratorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegratorId
}

// GetIntegratorIdOk returns a tuple with the IntegratorId field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetIntegratorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegratorId, true
}

// SetIntegratorId sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetIntegratorId(v string) {
	o.IntegratorId = v
}

// GetPublicDeviceId returns the PublicDeviceId field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicDeviceId
}

// GetPublicDeviceIdOk returns a tuple with the PublicDeviceId field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicDeviceId, true
}

// SetPublicDeviceId sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetPublicDeviceId(v string) {
	o.PublicDeviceId = v
}

// GetPublicWalletId returns the PublicWalletId field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicWalletId
}

// GetPublicWalletIdOk returns a tuple with the PublicWalletId field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicWalletId, true
}

// SetPublicWalletId sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetPublicWalletId(v string) {
	o.PublicWalletId = v
}

// GetServerSessionId returns the ServerSessionId field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetServerSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerSessionId
}

// GetServerSessionIdOk returns a tuple with the ServerSessionId field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetServerSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerSessionId, true
}

// SetServerSessionId sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetServerSessionId(v string) {
	o.ServerSessionId = v
}

// GetTokenSetting returns the TokenSetting field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetTokenSetting() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenSetting
}

// GetTokenSettingOk returns a tuple with the TokenSetting field value
// and a boolean to check if the value has been set.
func (o *SendingProvisioningDataToGooglePayBackendRequest) GetTokenSettingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenSetting, true
}

// SetTokenSetting sets field value
func (o *SendingProvisioningDataToGooglePayBackendRequest) SetTokenSetting(v int32) {
	o.TokenSetting = v
}

func (o SendingProvisioningDataToGooglePayBackendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendingProvisioningDataToGooglePayBackendRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_setting"] = o.CardSetting
	toSerialize["card_token"] = o.CardToken
	toSerialize["client_session_id"] = o.ClientSessionId
	toSerialize["integrator_id"] = o.IntegratorId
	toSerialize["public_device_id"] = o.PublicDeviceId
	toSerialize["public_wallet_id"] = o.PublicWalletId
	toSerialize["server_session_id"] = o.ServerSessionId
	toSerialize["token_setting"] = o.TokenSetting
	return toSerialize, nil
}

type NullableSendingProvisioningDataToGooglePayBackendRequest struct {
	value *SendingProvisioningDataToGooglePayBackendRequest
	isSet bool
}

func (v NullableSendingProvisioningDataToGooglePayBackendRequest) Get() *SendingProvisioningDataToGooglePayBackendRequest {
	return v.value
}

func (v *NullableSendingProvisioningDataToGooglePayBackendRequest) Set(val *SendingProvisioningDataToGooglePayBackendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendingProvisioningDataToGooglePayBackendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendingProvisioningDataToGooglePayBackendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendingProvisioningDataToGooglePayBackendRequest(val *SendingProvisioningDataToGooglePayBackendRequest) *NullableSendingProvisioningDataToGooglePayBackendRequest {
	return &NullableSendingProvisioningDataToGooglePayBackendRequest{value: val, isSet: true}
}

func (v NullableSendingProvisioningDataToGooglePayBackendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendingProvisioningDataToGooglePayBackendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


