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

// checks if the RewardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewardResponse{}

// RewardResponse struct for RewardResponse
type RewardResponse struct {
	// Unique identifier of the account on which the reward exists.
	AccountToken *string `json:"account_token,omitempty"`
	// Amount of the reward.
	Amount float32 `json:"amount"`
	// The total amount to which a percentage reward method is applied (for example, if a 3% reward is applied to 100, then `100` is the `applied_to_amount` value).  This field is not applicable for a flat fee method.  Returned for auto-cash back reward types only.
	AppliedToAmount *float32 `json:"applied_to_amount,omitempty"`
	// Date and time when the reward was created on Marqeta's credit platform, in UTC.
	CreatedTime time.Time `json:"created_time"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Description of the reward.
	Description string `json:"description"`
	Method *Method `json:"method,omitempty"`
	// Additional information about the reward.
	Note *string `json:"note,omitempty"`
	// Unique identifier of the reward.  If in the `detail_object`, unique identifier of the detail object.
	Token string `json:"token"`
	Type RewardType `json:"type"`
	// Date and time when the reward was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime time.Time `json:"updated_time"`
	// Value of the percentage or flat amount.  Returned for auto-cash back reward types only.
	Value *float32 `json:"value,omitempty"`
}

// NewRewardResponse instantiates a new RewardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardResponse(amount float32, createdTime time.Time, currencyCode CurrencyCode, description string, token string, type_ RewardType, updatedTime time.Time) *RewardResponse {
	this := RewardResponse{}
	this.Amount = amount
	this.CreatedTime = createdTime
	this.CurrencyCode = currencyCode
	this.Description = description
	this.Token = token
	this.Type = type_
	this.UpdatedTime = updatedTime
	return &this
}

// NewRewardResponseWithDefaults instantiates a new RewardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardResponseWithDefaults() *RewardResponse {
	this := RewardResponse{}
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = currencyCode
	return &this
}

// GetAccountToken returns the AccountToken field value if set, zero value otherwise.
func (o *RewardResponse) GetAccountToken() string {
	if o == nil || IsNil(o.AccountToken) {
		var ret string
		return ret
	}
	return *o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetAccountTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccountToken) {
		return nil, false
	}
	return o.AccountToken, true
}

// HasAccountToken returns a boolean if a field has been set.
func (o *RewardResponse) HasAccountToken() bool {
	if o != nil && !IsNil(o.AccountToken) {
		return true
	}

	return false
}

// SetAccountToken gets a reference to the given string and assigns it to the AccountToken field.
func (o *RewardResponse) SetAccountToken(v string) {
	o.AccountToken = &v
}

// GetAmount returns the Amount field value
func (o *RewardResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RewardResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetAppliedToAmount returns the AppliedToAmount field value if set, zero value otherwise.
func (o *RewardResponse) GetAppliedToAmount() float32 {
	if o == nil || IsNil(o.AppliedToAmount) {
		var ret float32
		return ret
	}
	return *o.AppliedToAmount
}

// GetAppliedToAmountOk returns a tuple with the AppliedToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetAppliedToAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.AppliedToAmount) {
		return nil, false
	}
	return o.AppliedToAmount, true
}

// HasAppliedToAmount returns a boolean if a field has been set.
func (o *RewardResponse) HasAppliedToAmount() bool {
	if o != nil && !IsNil(o.AppliedToAmount) {
		return true
	}

	return false
}

// SetAppliedToAmount gets a reference to the given float32 and assigns it to the AppliedToAmount field.
func (o *RewardResponse) SetAppliedToAmount(v float32) {
	o.AppliedToAmount = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *RewardResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *RewardResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *RewardResponse) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *RewardResponse) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetDescription returns the Description field value
func (o *RewardResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RewardResponse) SetDescription(v string) {
	o.Description = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RewardResponse) GetMethod() Method {
	if o == nil || IsNil(o.Method) {
		var ret Method
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetMethodOk() (*Method, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RewardResponse) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given Method and assigns it to the Method field.
func (o *RewardResponse) SetMethod(v Method) {
	o.Method = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *RewardResponse) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *RewardResponse) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *RewardResponse) SetNote(v string) {
	o.Note = &v
}

// GetToken returns the Token field value
func (o *RewardResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RewardResponse) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value
func (o *RewardResponse) GetType() RewardType {
	if o == nil {
		var ret RewardType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetTypeOk() (*RewardType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RewardResponse) SetType(v RewardType) {
	o.Type = v
}

// GetUpdatedTime returns the UpdatedTime field value
func (o *RewardResponse) GetUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTime, true
}

// SetUpdatedTime sets field value
func (o *RewardResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RewardResponse) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardResponse) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RewardResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *RewardResponse) SetValue(v float32) {
	o.Value = &v
}

func (o RewardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountToken) {
		toSerialize["account_token"] = o.AccountToken
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.AppliedToAmount) {
		toSerialize["applied_to_amount"] = o.AppliedToAmount
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["description"] = o.Description
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	toSerialize["token"] = o.Token
	toSerialize["type"] = o.Type
	toSerialize["updated_time"] = o.UpdatedTime
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRewardResponse struct {
	value *RewardResponse
	isSet bool
}

func (v NullableRewardResponse) Get() *RewardResponse {
	return v.value
}

func (v *NullableRewardResponse) Set(val *RewardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardResponse(val *RewardResponse) *NullableRewardResponse {
	return &NullableRewardResponse{value: val, isSet: true}
}

func (v NullableRewardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

