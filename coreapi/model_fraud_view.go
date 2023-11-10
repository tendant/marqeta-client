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

// checks if the FraudView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FraudView{}

// FraudView Contains one or more fraud determinations by the card network that apply to either the transaction or the cardholder's account.
type FraudView struct {
	IssuerProcessor *IssuerFraudView `json:"issuer_processor,omitempty"`
	Network *NetworkFraudView `json:"network,omitempty"`
	NetworkAccountIntelligenceScore *NetworkAccountIntelligenceScore `json:"network_account_intelligence_score,omitempty"`
}

// NewFraudView instantiates a new FraudView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFraudView() *FraudView {
	this := FraudView{}
	return &this
}

// NewFraudViewWithDefaults instantiates a new FraudView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFraudViewWithDefaults() *FraudView {
	this := FraudView{}
	return &this
}

// GetIssuerProcessor returns the IssuerProcessor field value if set, zero value otherwise.
func (o *FraudView) GetIssuerProcessor() IssuerFraudView {
	if o == nil || IsNil(o.IssuerProcessor) {
		var ret IssuerFraudView
		return ret
	}
	return *o.IssuerProcessor
}

// GetIssuerProcessorOk returns a tuple with the IssuerProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FraudView) GetIssuerProcessorOk() (*IssuerFraudView, bool) {
	if o == nil || IsNil(o.IssuerProcessor) {
		return nil, false
	}
	return o.IssuerProcessor, true
}

// HasIssuerProcessor returns a boolean if a field has been set.
func (o *FraudView) HasIssuerProcessor() bool {
	if o != nil && !IsNil(o.IssuerProcessor) {
		return true
	}

	return false
}

// SetIssuerProcessor gets a reference to the given IssuerFraudView and assigns it to the IssuerProcessor field.
func (o *FraudView) SetIssuerProcessor(v IssuerFraudView) {
	o.IssuerProcessor = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *FraudView) GetNetwork() NetworkFraudView {
	if o == nil || IsNil(o.Network) {
		var ret NetworkFraudView
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FraudView) GetNetworkOk() (*NetworkFraudView, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *FraudView) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NetworkFraudView and assigns it to the Network field.
func (o *FraudView) SetNetwork(v NetworkFraudView) {
	o.Network = &v
}

// GetNetworkAccountIntelligenceScore returns the NetworkAccountIntelligenceScore field value if set, zero value otherwise.
func (o *FraudView) GetNetworkAccountIntelligenceScore() NetworkAccountIntelligenceScore {
	if o == nil || IsNil(o.NetworkAccountIntelligenceScore) {
		var ret NetworkAccountIntelligenceScore
		return ret
	}
	return *o.NetworkAccountIntelligenceScore
}

// GetNetworkAccountIntelligenceScoreOk returns a tuple with the NetworkAccountIntelligenceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FraudView) GetNetworkAccountIntelligenceScoreOk() (*NetworkAccountIntelligenceScore, bool) {
	if o == nil || IsNil(o.NetworkAccountIntelligenceScore) {
		return nil, false
	}
	return o.NetworkAccountIntelligenceScore, true
}

// HasNetworkAccountIntelligenceScore returns a boolean if a field has been set.
func (o *FraudView) HasNetworkAccountIntelligenceScore() bool {
	if o != nil && !IsNil(o.NetworkAccountIntelligenceScore) {
		return true
	}

	return false
}

// SetNetworkAccountIntelligenceScore gets a reference to the given NetworkAccountIntelligenceScore and assigns it to the NetworkAccountIntelligenceScore field.
func (o *FraudView) SetNetworkAccountIntelligenceScore(v NetworkAccountIntelligenceScore) {
	o.NetworkAccountIntelligenceScore = &v
}

func (o FraudView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FraudView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IssuerProcessor) {
		toSerialize["issuer_processor"] = o.IssuerProcessor
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.NetworkAccountIntelligenceScore) {
		toSerialize["network_account_intelligence_score"] = o.NetworkAccountIntelligenceScore
	}
	return toSerialize, nil
}

type NullableFraudView struct {
	value *FraudView
	isSet bool
}

func (v NullableFraudView) Get() *FraudView {
	return v.value
}

func (v *NullableFraudView) Set(val *FraudView) {
	v.value = val
	v.isSet = true
}

func (v NullableFraudView) IsSet() bool {
	return v.isSet
}

func (v *NullableFraudView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFraudView(val *FraudView) *NullableFraudView {
	return &NullableFraudView{value: val, isSet: true}
}

func (v NullableFraudView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFraudView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


