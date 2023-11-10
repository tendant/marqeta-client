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

// checks if the PushToCardDisburseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushToCardDisburseRequest{}

// PushToCardDisburseRequest struct for PushToCardDisburseRequest
type PushToCardDisburseRequest struct {
	Amount float32 `json:"amount"`
	CurrencyCode string `json:"currency_code"`
	Memo *string `json:"memo,omitempty"`
	PaymentInstrumentToken string `json:"payment_instrument_token"`
	SoftDescriptor *PTCSoftDescriptor `json:"soft_descriptor,omitempty"`
	Tags *string `json:"tags,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewPushToCardDisburseRequest instantiates a new PushToCardDisburseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushToCardDisburseRequest(amount float32, currencyCode string, paymentInstrumentToken string) *PushToCardDisburseRequest {
	this := PushToCardDisburseRequest{}
	this.Amount = amount
	this.CurrencyCode = currencyCode
	this.PaymentInstrumentToken = paymentInstrumentToken
	return &this
}

// NewPushToCardDisburseRequestWithDefaults instantiates a new PushToCardDisburseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushToCardDisburseRequestWithDefaults() *PushToCardDisburseRequest {
	this := PushToCardDisburseRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PushToCardDisburseRequest) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PushToCardDisburseRequest) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PushToCardDisburseRequest) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *PushToCardDisburseRequest) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *PushToCardDisburseRequest) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *PushToCardDisburseRequest) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *PushToCardDisburseRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisburseRequest) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *PushToCardDisburseRequest) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *PushToCardDisburseRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetPaymentInstrumentToken returns the PaymentInstrumentToken field value
func (o *PushToCardDisburseRequest) GetPaymentInstrumentToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentInstrumentToken
}

// GetPaymentInstrumentTokenOk returns a tuple with the PaymentInstrumentToken field value
// and a boolean to check if the value has been set.
func (o *PushToCardDisburseRequest) GetPaymentInstrumentTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentInstrumentToken, true
}

// SetPaymentInstrumentToken sets field value
func (o *PushToCardDisburseRequest) SetPaymentInstrumentToken(v string) {
	o.PaymentInstrumentToken = v
}

// GetSoftDescriptor returns the SoftDescriptor field value if set, zero value otherwise.
func (o *PushToCardDisburseRequest) GetSoftDescriptor() PTCSoftDescriptor {
	if o == nil || IsNil(o.SoftDescriptor) {
		var ret PTCSoftDescriptor
		return ret
	}
	return *o.SoftDescriptor
}

// GetSoftDescriptorOk returns a tuple with the SoftDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisburseRequest) GetSoftDescriptorOk() (*PTCSoftDescriptor, bool) {
	if o == nil || IsNil(o.SoftDescriptor) {
		return nil, false
	}
	return o.SoftDescriptor, true
}

// HasSoftDescriptor returns a boolean if a field has been set.
func (o *PushToCardDisburseRequest) HasSoftDescriptor() bool {
	if o != nil && !IsNil(o.SoftDescriptor) {
		return true
	}

	return false
}

// SetSoftDescriptor gets a reference to the given PTCSoftDescriptor and assigns it to the SoftDescriptor field.
func (o *PushToCardDisburseRequest) SetSoftDescriptor(v PTCSoftDescriptor) {
	o.SoftDescriptor = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PushToCardDisburseRequest) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisburseRequest) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PushToCardDisburseRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PushToCardDisburseRequest) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PushToCardDisburseRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisburseRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PushToCardDisburseRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PushToCardDisburseRequest) SetToken(v string) {
	o.Token = &v
}

func (o PushToCardDisburseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushToCardDisburseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	toSerialize["payment_instrument_token"] = o.PaymentInstrumentToken
	if !IsNil(o.SoftDescriptor) {
		toSerialize["soft_descriptor"] = o.SoftDescriptor
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullablePushToCardDisburseRequest struct {
	value *PushToCardDisburseRequest
	isSet bool
}

func (v NullablePushToCardDisburseRequest) Get() *PushToCardDisburseRequest {
	return v.value
}

func (v *NullablePushToCardDisburseRequest) Set(val *PushToCardDisburseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePushToCardDisburseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePushToCardDisburseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushToCardDisburseRequest(val *PushToCardDisburseRequest) *NullablePushToCardDisburseRequest {
	return &NullablePushToCardDisburseRequest{value: val, isSet: true}
}

func (v NullablePushToCardDisburseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushToCardDisburseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


