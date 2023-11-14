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

// checks if the UnloadRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnloadRequestModel{}

// UnloadRequestModel struct for UnloadRequestModel
type UnloadRequestModel struct {
	// Amount of funds to return to the funding source.
	Amount float32 `json:"amount"`
	// Unique identifier of the funding source to use for this GPA unload order.  Send a `GET` request to `/fundingsources/addresses/user/{token}` to retrieve addresses for a specific user.
	FundingSourceAddressToken *string `json:"funding_source_address_token,omitempty"`
	// Additional descriptive text about the GPA unload.
	Memo *string `json:"memo,omitempty"`
	// Unique identifier of the original GPA order.
	OriginalOrderToken string `json:"original_order_token"`
	// Comma-delimited list of tags describing the GPA unload order.
	Tags *string `json:"tags,omitempty"`
	// Unique identifier of the GPA unload order.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
}

// NewUnloadRequestModel instantiates a new UnloadRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnloadRequestModel(amount float32, originalOrderToken string) *UnloadRequestModel {
	this := UnloadRequestModel{}
	this.Amount = amount
	this.OriginalOrderToken = originalOrderToken
	return &this
}

// NewUnloadRequestModelWithDefaults instantiates a new UnloadRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnloadRequestModelWithDefaults() *UnloadRequestModel {
	this := UnloadRequestModel{}
	return &this
}

// GetAmount returns the Amount field value
func (o *UnloadRequestModel) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UnloadRequestModel) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UnloadRequestModel) SetAmount(v float32) {
	o.Amount = v
}

// GetFundingSourceAddressToken returns the FundingSourceAddressToken field value if set, zero value otherwise.
func (o *UnloadRequestModel) GetFundingSourceAddressToken() string {
	if o == nil || IsNil(o.FundingSourceAddressToken) {
		var ret string
		return ret
	}
	return *o.FundingSourceAddressToken
}

// GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnloadRequestModel) GetFundingSourceAddressTokenOk() (*string, bool) {
	if o == nil || IsNil(o.FundingSourceAddressToken) {
		return nil, false
	}
	return o.FundingSourceAddressToken, true
}

// HasFundingSourceAddressToken returns a boolean if a field has been set.
func (o *UnloadRequestModel) HasFundingSourceAddressToken() bool {
	if o != nil && !IsNil(o.FundingSourceAddressToken) {
		return true
	}

	return false
}

// SetFundingSourceAddressToken gets a reference to the given string and assigns it to the FundingSourceAddressToken field.
func (o *UnloadRequestModel) SetFundingSourceAddressToken(v string) {
	o.FundingSourceAddressToken = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *UnloadRequestModel) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnloadRequestModel) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *UnloadRequestModel) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *UnloadRequestModel) SetMemo(v string) {
	o.Memo = &v
}

// GetOriginalOrderToken returns the OriginalOrderToken field value
func (o *UnloadRequestModel) GetOriginalOrderToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalOrderToken
}

// GetOriginalOrderTokenOk returns a tuple with the OriginalOrderToken field value
// and a boolean to check if the value has been set.
func (o *UnloadRequestModel) GetOriginalOrderTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalOrderToken, true
}

// SetOriginalOrderToken sets field value
func (o *UnloadRequestModel) SetOriginalOrderToken(v string) {
	o.OriginalOrderToken = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UnloadRequestModel) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnloadRequestModel) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UnloadRequestModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *UnloadRequestModel) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UnloadRequestModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnloadRequestModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UnloadRequestModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UnloadRequestModel) SetToken(v string) {
	o.Token = &v
}

func (o UnloadRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnloadRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.FundingSourceAddressToken) {
		toSerialize["funding_source_address_token"] = o.FundingSourceAddressToken
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	toSerialize["original_order_token"] = o.OriginalOrderToken
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUnloadRequestModel struct {
	value *UnloadRequestModel
	isSet bool
}

func (v NullableUnloadRequestModel) Get() *UnloadRequestModel {
	return v.value
}

func (v *NullableUnloadRequestModel) Set(val *UnloadRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUnloadRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUnloadRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnloadRequestModel(val *UnloadRequestModel) *NullableUnloadRequestModel {
	return &NullableUnloadRequestModel{value: val, isSet: true}
}

func (v NullableUnloadRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnloadRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


