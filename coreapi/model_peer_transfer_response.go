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
	"time"
)

// checks if the PeerTransferResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeerTransferResponse{}

// PeerTransferResponse Contains information about a peer transfer, including sender and recipient tokens, transfer amount, and currency code.
type PeerTransferResponse struct {
	// Amount of the transfer.
	Amount float32 `json:"amount"`
	CreatedTime time.Time `json:"created_time"`
	// Three-digit ISO 4217 currency code.
	CurrencyCode string `json:"currency_code"`
	// Additional descriptive text about the peer transfer.
	Memo *string `json:"memo,omitempty"`
	// Specifies the business account holder that receives funds.
	RecipientBusinessToken *string `json:"recipient_business_token,omitempty"`
	// Specifies the user account holder that receives funds.
	RecipientUserToken *string `json:"recipient_user_token,omitempty"`
	// Specifies the business account holder that sends funds.
	SenderBusinessToken *string `json:"sender_business_token,omitempty"`
	// Specifies the user account holder that sends funds.
	SenderUserToken *string `json:"sender_user_token,omitempty"`
	// Metadata about the peer transfer.
	Tags *string `json:"tags,omitempty"`
	// Unique identifier of the peer transfer request.
	Token string `json:"token"`
}

// NewPeerTransferResponse instantiates a new PeerTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeerTransferResponse(amount float32, createdTime time.Time, currencyCode string, token string) *PeerTransferResponse {
	this := PeerTransferResponse{}
	this.Amount = amount
	this.CreatedTime = createdTime
	this.CurrencyCode = currencyCode
	this.Token = token
	return &this
}

// NewPeerTransferResponseWithDefaults instantiates a new PeerTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeerTransferResponseWithDefaults() *PeerTransferResponse {
	this := PeerTransferResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PeerTransferResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PeerTransferResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *PeerTransferResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *PeerTransferResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *PeerTransferResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *PeerTransferResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *PeerTransferResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *PeerTransferResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *PeerTransferResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetRecipientBusinessToken returns the RecipientBusinessToken field value if set, zero value otherwise.
func (o *PeerTransferResponse) GetRecipientBusinessToken() string {
	if o == nil || IsNil(o.RecipientBusinessToken) {
		var ret string
		return ret
	}
	return *o.RecipientBusinessToken
}

// GetRecipientBusinessTokenOk returns a tuple with the RecipientBusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetRecipientBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientBusinessToken) {
		return nil, false
	}
	return o.RecipientBusinessToken, true
}

// HasRecipientBusinessToken returns a boolean if a field has been set.
func (o *PeerTransferResponse) HasRecipientBusinessToken() bool {
	if o != nil && !IsNil(o.RecipientBusinessToken) {
		return true
	}

	return false
}

// SetRecipientBusinessToken gets a reference to the given string and assigns it to the RecipientBusinessToken field.
func (o *PeerTransferResponse) SetRecipientBusinessToken(v string) {
	o.RecipientBusinessToken = &v
}

// GetRecipientUserToken returns the RecipientUserToken field value if set, zero value otherwise.
func (o *PeerTransferResponse) GetRecipientUserToken() string {
	if o == nil || IsNil(o.RecipientUserToken) {
		var ret string
		return ret
	}
	return *o.RecipientUserToken
}

// GetRecipientUserTokenOk returns a tuple with the RecipientUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetRecipientUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientUserToken) {
		return nil, false
	}
	return o.RecipientUserToken, true
}

// HasRecipientUserToken returns a boolean if a field has been set.
func (o *PeerTransferResponse) HasRecipientUserToken() bool {
	if o != nil && !IsNil(o.RecipientUserToken) {
		return true
	}

	return false
}

// SetRecipientUserToken gets a reference to the given string and assigns it to the RecipientUserToken field.
func (o *PeerTransferResponse) SetRecipientUserToken(v string) {
	o.RecipientUserToken = &v
}

// GetSenderBusinessToken returns the SenderBusinessToken field value if set, zero value otherwise.
func (o *PeerTransferResponse) GetSenderBusinessToken() string {
	if o == nil || IsNil(o.SenderBusinessToken) {
		var ret string
		return ret
	}
	return *o.SenderBusinessToken
}

// GetSenderBusinessTokenOk returns a tuple with the SenderBusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetSenderBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SenderBusinessToken) {
		return nil, false
	}
	return o.SenderBusinessToken, true
}

// HasSenderBusinessToken returns a boolean if a field has been set.
func (o *PeerTransferResponse) HasSenderBusinessToken() bool {
	if o != nil && !IsNil(o.SenderBusinessToken) {
		return true
	}

	return false
}

// SetSenderBusinessToken gets a reference to the given string and assigns it to the SenderBusinessToken field.
func (o *PeerTransferResponse) SetSenderBusinessToken(v string) {
	o.SenderBusinessToken = &v
}

// GetSenderUserToken returns the SenderUserToken field value if set, zero value otherwise.
func (o *PeerTransferResponse) GetSenderUserToken() string {
	if o == nil || IsNil(o.SenderUserToken) {
		var ret string
		return ret
	}
	return *o.SenderUserToken
}

// GetSenderUserTokenOk returns a tuple with the SenderUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetSenderUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SenderUserToken) {
		return nil, false
	}
	return o.SenderUserToken, true
}

// HasSenderUserToken returns a boolean if a field has been set.
func (o *PeerTransferResponse) HasSenderUserToken() bool {
	if o != nil && !IsNil(o.SenderUserToken) {
		return true
	}

	return false
}

// SetSenderUserToken gets a reference to the given string and assigns it to the SenderUserToken field.
func (o *PeerTransferResponse) SetSenderUserToken(v string) {
	o.SenderUserToken = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PeerTransferResponse) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PeerTransferResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PeerTransferResponse) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value
func (o *PeerTransferResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PeerTransferResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PeerTransferResponse) SetToken(v string) {
	o.Token = v
}

func (o PeerTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeerTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.RecipientBusinessToken) {
		toSerialize["recipient_business_token"] = o.RecipientBusinessToken
	}
	if !IsNil(o.RecipientUserToken) {
		toSerialize["recipient_user_token"] = o.RecipientUserToken
	}
	if !IsNil(o.SenderBusinessToken) {
		toSerialize["sender_business_token"] = o.SenderBusinessToken
	}
	if !IsNil(o.SenderUserToken) {
		toSerialize["sender_user_token"] = o.SenderUserToken
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullablePeerTransferResponse struct {
	value *PeerTransferResponse
	isSet bool
}

func (v NullablePeerTransferResponse) Get() *PeerTransferResponse {
	return v.value
}

func (v *NullablePeerTransferResponse) Set(val *PeerTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePeerTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePeerTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeerTransferResponse(val *PeerTransferResponse) *NullablePeerTransferResponse {
	return &NullablePeerTransferResponse{value: val, isSet: true}
}

func (v NullablePeerTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeerTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


