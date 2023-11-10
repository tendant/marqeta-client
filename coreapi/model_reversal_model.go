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

// checks if the ReversalModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReversalModel{}

// ReversalModel struct for ReversalModel
type ReversalModel struct {
	Amount float32 `json:"amount"`
	FindOriginalWindowDays *int32 `json:"find_original_window_days,omitempty"`
	IsAdvice *bool `json:"is_advice,omitempty"`
	NetworkFees []NetworkFeeModel `json:"network_fees,omitempty"`
	OriginalTransactionToken string `json:"original_transaction_token"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

// NewReversalModel instantiates a new ReversalModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReversalModel(amount float32, originalTransactionToken string) *ReversalModel {
	this := ReversalModel{}
	this.Amount = amount
	var isAdvice bool = false
	this.IsAdvice = &isAdvice
	this.OriginalTransactionToken = originalTransactionToken
	return &this
}

// NewReversalModelWithDefaults instantiates a new ReversalModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReversalModelWithDefaults() *ReversalModel {
	this := ReversalModel{}
	var isAdvice bool = false
	this.IsAdvice = &isAdvice
	return &this
}

// GetAmount returns the Amount field value
func (o *ReversalModel) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ReversalModel) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ReversalModel) SetAmount(v float32) {
	o.Amount = v
}

// GetFindOriginalWindowDays returns the FindOriginalWindowDays field value if set, zero value otherwise.
func (o *ReversalModel) GetFindOriginalWindowDays() int32 {
	if o == nil || IsNil(o.FindOriginalWindowDays) {
		var ret int32
		return ret
	}
	return *o.FindOriginalWindowDays
}

// GetFindOriginalWindowDaysOk returns a tuple with the FindOriginalWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReversalModel) GetFindOriginalWindowDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.FindOriginalWindowDays) {
		return nil, false
	}
	return o.FindOriginalWindowDays, true
}

// HasFindOriginalWindowDays returns a boolean if a field has been set.
func (o *ReversalModel) HasFindOriginalWindowDays() bool {
	if o != nil && !IsNil(o.FindOriginalWindowDays) {
		return true
	}

	return false
}

// SetFindOriginalWindowDays gets a reference to the given int32 and assigns it to the FindOriginalWindowDays field.
func (o *ReversalModel) SetFindOriginalWindowDays(v int32) {
	o.FindOriginalWindowDays = &v
}

// GetIsAdvice returns the IsAdvice field value if set, zero value otherwise.
func (o *ReversalModel) GetIsAdvice() bool {
	if o == nil || IsNil(o.IsAdvice) {
		var ret bool
		return ret
	}
	return *o.IsAdvice
}

// GetIsAdviceOk returns a tuple with the IsAdvice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReversalModel) GetIsAdviceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdvice) {
		return nil, false
	}
	return o.IsAdvice, true
}

// HasIsAdvice returns a boolean if a field has been set.
func (o *ReversalModel) HasIsAdvice() bool {
	if o != nil && !IsNil(o.IsAdvice) {
		return true
	}

	return false
}

// SetIsAdvice gets a reference to the given bool and assigns it to the IsAdvice field.
func (o *ReversalModel) SetIsAdvice(v bool) {
	o.IsAdvice = &v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *ReversalModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || IsNil(o.NetworkFees) {
		var ret []NetworkFeeModel
		return ret
	}
	return o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReversalModel) GetNetworkFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || IsNil(o.NetworkFees) {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *ReversalModel) HasNetworkFees() bool {
	if o != nil && !IsNil(o.NetworkFees) {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *ReversalModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = v
}

// GetOriginalTransactionToken returns the OriginalTransactionToken field value
func (o *ReversalModel) GetOriginalTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalTransactionToken
}

// GetOriginalTransactionTokenOk returns a tuple with the OriginalTransactionToken field value
// and a boolean to check if the value has been set.
func (o *ReversalModel) GetOriginalTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalTransactionToken, true
}

// SetOriginalTransactionToken sets field value
func (o *ReversalModel) SetOriginalTransactionToken(v string) {
	o.OriginalTransactionToken = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *ReversalModel) GetWebhook() Webhook {
	if o == nil || IsNil(o.Webhook) {
		var ret Webhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReversalModel) GetWebhookOk() (*Webhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *ReversalModel) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given Webhook and assigns it to the Webhook field.
func (o *ReversalModel) SetWebhook(v Webhook) {
	o.Webhook = &v
}

func (o ReversalModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReversalModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.FindOriginalWindowDays) {
		toSerialize["find_original_window_days"] = o.FindOriginalWindowDays
	}
	if !IsNil(o.IsAdvice) {
		toSerialize["is_advice"] = o.IsAdvice
	}
	if !IsNil(o.NetworkFees) {
		toSerialize["network_fees"] = o.NetworkFees
	}
	toSerialize["original_transaction_token"] = o.OriginalTransactionToken
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

type NullableReversalModel struct {
	value *ReversalModel
	isSet bool
}

func (v NullableReversalModel) Get() *ReversalModel {
	return v.value
}

func (v *NullableReversalModel) Set(val *ReversalModel) {
	v.value = val
	v.isSet = true
}

func (v NullableReversalModel) IsSet() bool {
	return v.isSet
}

func (v *NullableReversalModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReversalModel(val *ReversalModel) *NullableReversalModel {
	return &NullableReversalModel{value: val, isSet: true}
}

func (v NullableReversalModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReversalModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

