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

// checks if the OrignalcreditRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrignalcreditRequestModel{}

// OrignalcreditRequestModel struct for OrignalcreditRequestModel
type OrignalcreditRequestModel struct {
	Amount float32 `json:"amount"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	CardToken string `json:"card_token"`
	Mid string `json:"mid"`
	ScreeningScore *string `json:"screening_score,omitempty"`
	SenderData *OriginalCreditSenderData `json:"sender_data,omitempty"`
	TransactionPurpose *string `json:"transactionPurpose,omitempty"`
	Type string `json:"type"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

// NewOrignalcreditRequestModel instantiates a new OrignalcreditRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrignalcreditRequestModel(amount float32, cardToken string, mid string, type_ string) *OrignalcreditRequestModel {
	this := OrignalcreditRequestModel{}
	this.Amount = amount
	this.CardToken = cardToken
	this.Mid = mid
	this.Type = type_
	return &this
}

// NewOrignalcreditRequestModelWithDefaults instantiates a new OrignalcreditRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrignalcreditRequestModelWithDefaults() *OrignalcreditRequestModel {
	this := OrignalcreditRequestModel{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OrignalcreditRequestModel) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrignalcreditRequestModel) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *OrignalcreditRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *OrignalcreditRequestModel) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *OrignalcreditRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetCardToken returns the CardToken field value
func (o *OrignalcreditRequestModel) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *OrignalcreditRequestModel) SetCardToken(v string) {
	o.CardToken = v
}

// GetMid returns the Mid field value
func (o *OrignalcreditRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *OrignalcreditRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetScreeningScore returns the ScreeningScore field value if set, zero value otherwise.
func (o *OrignalcreditRequestModel) GetScreeningScore() string {
	if o == nil || IsNil(o.ScreeningScore) {
		var ret string
		return ret
	}
	return *o.ScreeningScore
}

// GetScreeningScoreOk returns a tuple with the ScreeningScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetScreeningScoreOk() (*string, bool) {
	if o == nil || IsNil(o.ScreeningScore) {
		return nil, false
	}
	return o.ScreeningScore, true
}

// HasScreeningScore returns a boolean if a field has been set.
func (o *OrignalcreditRequestModel) HasScreeningScore() bool {
	if o != nil && !IsNil(o.ScreeningScore) {
		return true
	}

	return false
}

// SetScreeningScore gets a reference to the given string and assigns it to the ScreeningScore field.
func (o *OrignalcreditRequestModel) SetScreeningScore(v string) {
	o.ScreeningScore = &v
}

// GetSenderData returns the SenderData field value if set, zero value otherwise.
func (o *OrignalcreditRequestModel) GetSenderData() OriginalCreditSenderData {
	if o == nil || IsNil(o.SenderData) {
		var ret OriginalCreditSenderData
		return ret
	}
	return *o.SenderData
}

// GetSenderDataOk returns a tuple with the SenderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetSenderDataOk() (*OriginalCreditSenderData, bool) {
	if o == nil || IsNil(o.SenderData) {
		return nil, false
	}
	return o.SenderData, true
}

// HasSenderData returns a boolean if a field has been set.
func (o *OrignalcreditRequestModel) HasSenderData() bool {
	if o != nil && !IsNil(o.SenderData) {
		return true
	}

	return false
}

// SetSenderData gets a reference to the given OriginalCreditSenderData and assigns it to the SenderData field.
func (o *OrignalcreditRequestModel) SetSenderData(v OriginalCreditSenderData) {
	o.SenderData = &v
}

// GetTransactionPurpose returns the TransactionPurpose field value if set, zero value otherwise.
func (o *OrignalcreditRequestModel) GetTransactionPurpose() string {
	if o == nil || IsNil(o.TransactionPurpose) {
		var ret string
		return ret
	}
	return *o.TransactionPurpose
}

// GetTransactionPurposeOk returns a tuple with the TransactionPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetTransactionPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionPurpose) {
		return nil, false
	}
	return o.TransactionPurpose, true
}

// HasTransactionPurpose returns a boolean if a field has been set.
func (o *OrignalcreditRequestModel) HasTransactionPurpose() bool {
	if o != nil && !IsNil(o.TransactionPurpose) {
		return true
	}

	return false
}

// SetTransactionPurpose gets a reference to the given string and assigns it to the TransactionPurpose field.
func (o *OrignalcreditRequestModel) SetTransactionPurpose(v string) {
	o.TransactionPurpose = &v
}

// GetType returns the Type field value
func (o *OrignalcreditRequestModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrignalcreditRequestModel) SetType(v string) {
	o.Type = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *OrignalcreditRequestModel) GetWebhook() Webhook {
	if o == nil || IsNil(o.Webhook) {
		var ret Webhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrignalcreditRequestModel) GetWebhookOk() (*Webhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *OrignalcreditRequestModel) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given Webhook and assigns it to the Webhook field.
func (o *OrignalcreditRequestModel) SetWebhook(v Webhook) {
	o.Webhook = &v
}

func (o OrignalcreditRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrignalcreditRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	toSerialize["card_token"] = o.CardToken
	toSerialize["mid"] = o.Mid
	if !IsNil(o.ScreeningScore) {
		toSerialize["screening_score"] = o.ScreeningScore
	}
	if !IsNil(o.SenderData) {
		toSerialize["sender_data"] = o.SenderData
	}
	if !IsNil(o.TransactionPurpose) {
		toSerialize["transactionPurpose"] = o.TransactionPurpose
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

type NullableOrignalcreditRequestModel struct {
	value *OrignalcreditRequestModel
	isSet bool
}

func (v NullableOrignalcreditRequestModel) Get() *OrignalcreditRequestModel {
	return v.value
}

func (v *NullableOrignalcreditRequestModel) Set(val *OrignalcreditRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrignalcreditRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrignalcreditRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrignalcreditRequestModel(val *OrignalcreditRequestModel) *NullableOrignalcreditRequestModel {
	return &NullableOrignalcreditRequestModel{value: val, isSet: true}
}

func (v NullableOrignalcreditRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrignalcreditRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


