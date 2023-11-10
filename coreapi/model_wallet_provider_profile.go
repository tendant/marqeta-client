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

// checks if the WalletProviderProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletProviderProfile{}

// WalletProviderProfile Contains information held and provided by the digital wallet provider.
type WalletProviderProfile struct {
	Account *Account `json:"account,omitempty"`
	// Score from the device.
	DeviceScore *string `json:"device_score,omitempty"`
	// Source from which the digital wallet provider obtained the card primary account number (PAN).
	PanSource *string `json:"pan_source,omitempty"`
	// Reason for the wallet provider's provisioning decision.  * *01*  Cardholder's wallet account is too new relative to launch. * *02*  Cardholder's wallet account is too new relative to provisioning request. * *03*  Cardholder's wallet account/card pair is newer than date threshold. * *04*  Changes made to account data within the account threshold. * *05*  Suspicious transactions linked to this account. * *06*  Account has not had activity in the last year. * *07*  Suspended cards in the secure element. * *08*  Device was put in lost mode in the last seven days for longer than the duration threshold. * *09*  The number of provisioning attempts on this device in 24 hours exceeds threshold. * *0A*  There have been more than the threshold number of different cards attempted at provisioning to this phone in 24 hours. * *0B*  The card provisioning attempt contains a distinct name in excess of the threshold. * *0C*  The device score is less than 3. * *0D*  The account score is less than 4. * *0E*  Device provisioning location outside of the cardholder's wallet account home country. * *0G*  Suspect fraud.
	ReasonCode *string `json:"reason_code,omitempty"`
	// Array of recommendation reasons from the digital wallet provider.
	RecommendationReasons []string `json:"recommendation_reasons,omitempty"`
	RiskAssessment *RiskAssessment `json:"risk_assessment,omitempty"`
}

// NewWalletProviderProfile instantiates a new WalletProviderProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletProviderProfile() *WalletProviderProfile {
	this := WalletProviderProfile{}
	return &this
}

// NewWalletProviderProfileWithDefaults instantiates a new WalletProviderProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletProviderProfileWithDefaults() *WalletProviderProfile {
	this := WalletProviderProfile{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *WalletProviderProfile) GetAccount() Account {
	if o == nil || IsNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletProviderProfile) GetAccountOk() (*Account, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *WalletProviderProfile) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *WalletProviderProfile) SetAccount(v Account) {
	o.Account = &v
}

// GetDeviceScore returns the DeviceScore field value if set, zero value otherwise.
func (o *WalletProviderProfile) GetDeviceScore() string {
	if o == nil || IsNil(o.DeviceScore) {
		var ret string
		return ret
	}
	return *o.DeviceScore
}

// GetDeviceScoreOk returns a tuple with the DeviceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletProviderProfile) GetDeviceScoreOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceScore) {
		return nil, false
	}
	return o.DeviceScore, true
}

// HasDeviceScore returns a boolean if a field has been set.
func (o *WalletProviderProfile) HasDeviceScore() bool {
	if o != nil && !IsNil(o.DeviceScore) {
		return true
	}

	return false
}

// SetDeviceScore gets a reference to the given string and assigns it to the DeviceScore field.
func (o *WalletProviderProfile) SetDeviceScore(v string) {
	o.DeviceScore = &v
}

// GetPanSource returns the PanSource field value if set, zero value otherwise.
func (o *WalletProviderProfile) GetPanSource() string {
	if o == nil || IsNil(o.PanSource) {
		var ret string
		return ret
	}
	return *o.PanSource
}

// GetPanSourceOk returns a tuple with the PanSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletProviderProfile) GetPanSourceOk() (*string, bool) {
	if o == nil || IsNil(o.PanSource) {
		return nil, false
	}
	return o.PanSource, true
}

// HasPanSource returns a boolean if a field has been set.
func (o *WalletProviderProfile) HasPanSource() bool {
	if o != nil && !IsNil(o.PanSource) {
		return true
	}

	return false
}

// SetPanSource gets a reference to the given string and assigns it to the PanSource field.
func (o *WalletProviderProfile) SetPanSource(v string) {
	o.PanSource = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *WalletProviderProfile) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletProviderProfile) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *WalletProviderProfile) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *WalletProviderProfile) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetRecommendationReasons returns the RecommendationReasons field value if set, zero value otherwise.
func (o *WalletProviderProfile) GetRecommendationReasons() []string {
	if o == nil || IsNil(o.RecommendationReasons) {
		var ret []string
		return ret
	}
	return o.RecommendationReasons
}

// GetRecommendationReasonsOk returns a tuple with the RecommendationReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletProviderProfile) GetRecommendationReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.RecommendationReasons) {
		return nil, false
	}
	return o.RecommendationReasons, true
}

// HasRecommendationReasons returns a boolean if a field has been set.
func (o *WalletProviderProfile) HasRecommendationReasons() bool {
	if o != nil && !IsNil(o.RecommendationReasons) {
		return true
	}

	return false
}

// SetRecommendationReasons gets a reference to the given []string and assigns it to the RecommendationReasons field.
func (o *WalletProviderProfile) SetRecommendationReasons(v []string) {
	o.RecommendationReasons = v
}

// GetRiskAssessment returns the RiskAssessment field value if set, zero value otherwise.
func (o *WalletProviderProfile) GetRiskAssessment() RiskAssessment {
	if o == nil || IsNil(o.RiskAssessment) {
		var ret RiskAssessment
		return ret
	}
	return *o.RiskAssessment
}

// GetRiskAssessmentOk returns a tuple with the RiskAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletProviderProfile) GetRiskAssessmentOk() (*RiskAssessment, bool) {
	if o == nil || IsNil(o.RiskAssessment) {
		return nil, false
	}
	return o.RiskAssessment, true
}

// HasRiskAssessment returns a boolean if a field has been set.
func (o *WalletProviderProfile) HasRiskAssessment() bool {
	if o != nil && !IsNil(o.RiskAssessment) {
		return true
	}

	return false
}

// SetRiskAssessment gets a reference to the given RiskAssessment and assigns it to the RiskAssessment field.
func (o *WalletProviderProfile) SetRiskAssessment(v RiskAssessment) {
	o.RiskAssessment = &v
}

func (o WalletProviderProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletProviderProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.DeviceScore) {
		toSerialize["device_score"] = o.DeviceScore
	}
	if !IsNil(o.PanSource) {
		toSerialize["pan_source"] = o.PanSource
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reason_code"] = o.ReasonCode
	}
	if !IsNil(o.RecommendationReasons) {
		toSerialize["recommendation_reasons"] = o.RecommendationReasons
	}
	if !IsNil(o.RiskAssessment) {
		toSerialize["risk_assessment"] = o.RiskAssessment
	}
	return toSerialize, nil
}

type NullableWalletProviderProfile struct {
	value *WalletProviderProfile
	isSet bool
}

func (v NullableWalletProviderProfile) Get() *WalletProviderProfile {
	return v.value
}

func (v *NullableWalletProviderProfile) Set(val *WalletProviderProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletProviderProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletProviderProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletProviderProfile(val *WalletProviderProfile) *NullableWalletProviderProfile {
	return &NullableWalletProviderProfile{value: val, isSet: true}
}

func (v NullableWalletProviderProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletProviderProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


