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

// checks if the Msa type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Msa{}

// Msa struct for Msa
type Msa struct {
	CampaignToken string `json:"campaign_token"`
	ReloadAmount float32 `json:"reload_amount"`
	TriggerAmount float32 `json:"trigger_amount"`
}

// NewMsa instantiates a new Msa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsa(campaignToken string, reloadAmount float32, triggerAmount float32) *Msa {
	this := Msa{}
	this.CampaignToken = campaignToken
	this.ReloadAmount = reloadAmount
	this.TriggerAmount = triggerAmount
	return &this
}

// NewMsaWithDefaults instantiates a new Msa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsaWithDefaults() *Msa {
	this := Msa{}
	return &this
}

// GetCampaignToken returns the CampaignToken field value
func (o *Msa) GetCampaignToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignToken
}

// GetCampaignTokenOk returns a tuple with the CampaignToken field value
// and a boolean to check if the value has been set.
func (o *Msa) GetCampaignTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignToken, true
}

// SetCampaignToken sets field value
func (o *Msa) SetCampaignToken(v string) {
	o.CampaignToken = v
}

// GetReloadAmount returns the ReloadAmount field value
func (o *Msa) GetReloadAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ReloadAmount
}

// GetReloadAmountOk returns a tuple with the ReloadAmount field value
// and a boolean to check if the value has been set.
func (o *Msa) GetReloadAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReloadAmount, true
}

// SetReloadAmount sets field value
func (o *Msa) SetReloadAmount(v float32) {
	o.ReloadAmount = v
}

// GetTriggerAmount returns the TriggerAmount field value
func (o *Msa) GetTriggerAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TriggerAmount
}

// GetTriggerAmountOk returns a tuple with the TriggerAmount field value
// and a boolean to check if the value has been set.
func (o *Msa) GetTriggerAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerAmount, true
}

// SetTriggerAmount sets field value
func (o *Msa) SetTriggerAmount(v float32) {
	o.TriggerAmount = v
}

func (o Msa) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Msa) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign_token"] = o.CampaignToken
	toSerialize["reload_amount"] = o.ReloadAmount
	toSerialize["trigger_amount"] = o.TriggerAmount
	return toSerialize, nil
}

type NullableMsa struct {
	value *Msa
	isSet bool
}

func (v NullableMsa) Get() *Msa {
	return v.value
}

func (v *NullableMsa) Set(val *Msa) {
	v.value = val
	v.isSet = true
}

func (v NullableMsa) IsSet() bool {
	return v.isSet
}

func (v *NullableMsa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsa(val *Msa) *NullableMsa {
	return &NullableMsa{value: val, isSet: true}
}

func (v NullableMsa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


