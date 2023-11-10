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

// checks if the ProgramReserveTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramReserveTransactionResponse{}

// ProgramReserveTransactionResponse struct for ProgramReserveTransactionResponse
type ProgramReserveTransactionResponse struct {
	// Amount of the program reserve account credit or debit. Sometimes referred to as a _program funding account_.
	Amount *float32 `json:"amount,omitempty"`
	// Date and time when the resource was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Three-digit ISO 4217 currency code.
	CurrencyCode *string `json:"currency_code,omitempty"`
	IsCollateral *bool `json:"is_collateral,omitempty"`
	// The date and time when the resource was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Memo or note describing the transaction.
	Memo *string `json:"memo,omitempty"`
	// Transaction state.
	State *string `json:"state,omitempty"`
	// Comma-delimited list of tags describing the transaction.
	Tags *string `json:"tags,omitempty"`
	// The unique identifier of the transaction response.
	Token *string `json:"token,omitempty"`
	// Unique identifier of the transaction.
	TransactionToken *string `json:"transaction_token,omitempty"`
	// Transaction type.
	Type *string `json:"type,omitempty"`
}

// NewProgramReserveTransactionResponse instantiates a new ProgramReserveTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramReserveTransactionResponse(createdTime time.Time, lastModifiedTime time.Time) *ProgramReserveTransactionResponse {
	this := ProgramReserveTransactionResponse{}
	this.CreatedTime = createdTime
	this.LastModifiedTime = lastModifiedTime
	return &this
}

// NewProgramReserveTransactionResponseWithDefaults instantiates a new ProgramReserveTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramReserveTransactionResponseWithDefaults() *ProgramReserveTransactionResponse {
	this := ProgramReserveTransactionResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *ProgramReserveTransactionResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *ProgramReserveTransactionResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *ProgramReserveTransactionResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ProgramReserveTransactionResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetIsCollateral returns the IsCollateral field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetIsCollateral() bool {
	if o == nil || IsNil(o.IsCollateral) {
		var ret bool
		return ret
	}
	return *o.IsCollateral
}

// GetIsCollateralOk returns a tuple with the IsCollateral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetIsCollateralOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCollateral) {
		return nil, false
	}
	return o.IsCollateral, true
}

// HasIsCollateral returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasIsCollateral() bool {
	if o != nil && !IsNil(o.IsCollateral) {
		return true
	}

	return false
}

// SetIsCollateral gets a reference to the given bool and assigns it to the IsCollateral field.
func (o *ProgramReserveTransactionResponse) SetIsCollateral(v bool) {
	o.IsCollateral = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *ProgramReserveTransactionResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *ProgramReserveTransactionResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *ProgramReserveTransactionResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ProgramReserveTransactionResponse) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ProgramReserveTransactionResponse) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProgramReserveTransactionResponse) SetToken(v string) {
	o.Token = &v
}

// GetTransactionToken returns the TransactionToken field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetTransactionToken() string {
	if o == nil || IsNil(o.TransactionToken) {
		var ret string
		return ret
	}
	return *o.TransactionToken
}

// GetTransactionTokenOk returns a tuple with the TransactionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetTransactionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionToken) {
		return nil, false
	}
	return o.TransactionToken, true
}

// HasTransactionToken returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasTransactionToken() bool {
	if o != nil && !IsNil(o.TransactionToken) {
		return true
	}

	return false
}

// SetTransactionToken gets a reference to the given string and assigns it to the TransactionToken field.
func (o *ProgramReserveTransactionResponse) SetTransactionToken(v string) {
	o.TransactionToken = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProgramReserveTransactionResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramReserveTransactionResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProgramReserveTransactionResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProgramReserveTransactionResponse) SetType(v string) {
	o.Type = &v
}

func (o ProgramReserveTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramReserveTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.IsCollateral) {
		toSerialize["is_collateral"] = o.IsCollateral
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.TransactionToken) {
		toSerialize["transaction_token"] = o.TransactionToken
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableProgramReserveTransactionResponse struct {
	value *ProgramReserveTransactionResponse
	isSet bool
}

func (v NullableProgramReserveTransactionResponse) Get() *ProgramReserveTransactionResponse {
	return v.value
}

func (v *NullableProgramReserveTransactionResponse) Set(val *ProgramReserveTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramReserveTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramReserveTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramReserveTransactionResponse(val *ProgramReserveTransactionResponse) *NullableProgramReserveTransactionResponse {
	return &NullableProgramReserveTransactionResponse{value: val, isSet: true}
}

func (v NullableProgramReserveTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramReserveTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


