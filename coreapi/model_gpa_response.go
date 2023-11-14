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
	"time"
)

// checks if the GpaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GpaResponse{}

// GpaResponse Contains information about a GPA order, including fees, funding sources, and addresses. See <</core-api/gpa-orders, GPA Orders>> for more information.
type GpaResponse struct {
	// Amount funded.
	Amount float32 `json:"amount"`
	// Unique identifier of the business.  This field is returned if it exists in the resource.
	BusinessToken *string `json:"business_token,omitempty"`
	// Date and time when the GPA order was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Three-digit ISO 4217 currency code.
	CurrencyCode string `json:"currency_code"`
	// List of fees associated with the funding transaction.  This array is returned if it exists in the resource.
	Fees []FeeDetail `json:"fees,omitempty"`
	Funding Funding `json:"funding"`
	// Unique identifier of the funding source address to use for this order.
	FundingSourceAddressToken *string `json:"funding_source_address_token,omitempty"`
	// Unique identifier of the funding source to use for this order.
	FundingSourceToken string `json:"funding_source_token"`
	// Message about the status of the funding request. Useful for determining whether it was approved and completed successfully, declined by the gateway, or timed out.  This field is returned if it exists in the resource.
	GatewayMessage *string `json:"gateway_message,omitempty"`
	// Unique identifier of the JIT Funding request and response.  This field is returned if it exists in the resource.
	GatewayToken *int64 `json:"gateway_token,omitempty"`
	JitFunding *JitFundingApi `json:"jit_funding,omitempty"`
	// Date and time when the GPA order was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Additional descriptive text.  This field is returned if it exists in the resource.
	Memo *string `json:"memo,omitempty"`
	Response Response `json:"response"`
	// Current status of the funding transaction.
	State string `json:"state"`
	// Comma-delimited list of tags describing the GPA order.  This field is returned if it exists in the resource.
	Tags *string `json:"tags,omitempty"`
	// Unique identifier of the GPA order.
	Token string `json:"token"`
	// Unique identifier of the transaction being funded.
	TransactionToken string `json:"transaction_token"`
	// Unique identifier of the user resource.  This field is returned if it exists in the resource.
	UserToken *string `json:"user_token,omitempty"`
}

// NewGpaResponse instantiates a new GpaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpaResponse(amount float32, createdTime time.Time, currencyCode string, funding Funding, fundingSourceToken string, lastModifiedTime time.Time, response Response, state string, token string, transactionToken string) *GpaResponse {
	this := GpaResponse{}
	this.Amount = amount
	this.CreatedTime = createdTime
	this.CurrencyCode = currencyCode
	this.Funding = funding
	this.FundingSourceToken = fundingSourceToken
	this.LastModifiedTime = lastModifiedTime
	this.Response = response
	this.State = state
	this.Token = token
	this.TransactionToken = transactionToken
	return &this
}

// NewGpaResponseWithDefaults instantiates a new GpaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpaResponseWithDefaults() *GpaResponse {
	this := GpaResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GpaResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GpaResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *GpaResponse) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *GpaResponse) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *GpaResponse) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *GpaResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *GpaResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *GpaResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *GpaResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *GpaResponse) GetFees() []FeeDetail {
	if o == nil || IsNil(o.Fees) {
		var ret []FeeDetail
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetFeesOk() ([]FeeDetail, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *GpaResponse) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []FeeDetail and assigns it to the Fees field.
func (o *GpaResponse) SetFees(v []FeeDetail) {
	o.Fees = v
}

// GetFunding returns the Funding field value
func (o *GpaResponse) GetFunding() Funding {
	if o == nil {
		var ret Funding
		return ret
	}

	return o.Funding
}

// GetFundingOk returns a tuple with the Funding field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetFundingOk() (*Funding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Funding, true
}

// SetFunding sets field value
func (o *GpaResponse) SetFunding(v Funding) {
	o.Funding = v
}

// GetFundingSourceAddressToken returns the FundingSourceAddressToken field value if set, zero value otherwise.
func (o *GpaResponse) GetFundingSourceAddressToken() string {
	if o == nil || IsNil(o.FundingSourceAddressToken) {
		var ret string
		return ret
	}
	return *o.FundingSourceAddressToken
}

// GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetFundingSourceAddressTokenOk() (*string, bool) {
	if o == nil || IsNil(o.FundingSourceAddressToken) {
		return nil, false
	}
	return o.FundingSourceAddressToken, true
}

// HasFundingSourceAddressToken returns a boolean if a field has been set.
func (o *GpaResponse) HasFundingSourceAddressToken() bool {
	if o != nil && !IsNil(o.FundingSourceAddressToken) {
		return true
	}

	return false
}

// SetFundingSourceAddressToken gets a reference to the given string and assigns it to the FundingSourceAddressToken field.
func (o *GpaResponse) SetFundingSourceAddressToken(v string) {
	o.FundingSourceAddressToken = &v
}

// GetFundingSourceToken returns the FundingSourceToken field value
func (o *GpaResponse) GetFundingSourceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingSourceToken
}

// GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetFundingSourceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingSourceToken, true
}

// SetFundingSourceToken sets field value
func (o *GpaResponse) SetFundingSourceToken(v string) {
	o.FundingSourceToken = v
}

// GetGatewayMessage returns the GatewayMessage field value if set, zero value otherwise.
func (o *GpaResponse) GetGatewayMessage() string {
	if o == nil || IsNil(o.GatewayMessage) {
		var ret string
		return ret
	}
	return *o.GatewayMessage
}

// GetGatewayMessageOk returns a tuple with the GatewayMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetGatewayMessageOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayMessage) {
		return nil, false
	}
	return o.GatewayMessage, true
}

// HasGatewayMessage returns a boolean if a field has been set.
func (o *GpaResponse) HasGatewayMessage() bool {
	if o != nil && !IsNil(o.GatewayMessage) {
		return true
	}

	return false
}

// SetGatewayMessage gets a reference to the given string and assigns it to the GatewayMessage field.
func (o *GpaResponse) SetGatewayMessage(v string) {
	o.GatewayMessage = &v
}

// GetGatewayToken returns the GatewayToken field value if set, zero value otherwise.
func (o *GpaResponse) GetGatewayToken() int64 {
	if o == nil || IsNil(o.GatewayToken) {
		var ret int64
		return ret
	}
	return *o.GatewayToken
}

// GetGatewayTokenOk returns a tuple with the GatewayToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetGatewayTokenOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayToken) {
		return nil, false
	}
	return o.GatewayToken, true
}

// HasGatewayToken returns a boolean if a field has been set.
func (o *GpaResponse) HasGatewayToken() bool {
	if o != nil && !IsNil(o.GatewayToken) {
		return true
	}

	return false
}

// SetGatewayToken gets a reference to the given int64 and assigns it to the GatewayToken field.
func (o *GpaResponse) SetGatewayToken(v int64) {
	o.GatewayToken = &v
}

// GetJitFunding returns the JitFunding field value if set, zero value otherwise.
func (o *GpaResponse) GetJitFunding() JitFundingApi {
	if o == nil || IsNil(o.JitFunding) {
		var ret JitFundingApi
		return ret
	}
	return *o.JitFunding
}

// GetJitFundingOk returns a tuple with the JitFunding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetJitFundingOk() (*JitFundingApi, bool) {
	if o == nil || IsNil(o.JitFunding) {
		return nil, false
	}
	return o.JitFunding, true
}

// HasJitFunding returns a boolean if a field has been set.
func (o *GpaResponse) HasJitFunding() bool {
	if o != nil && !IsNil(o.JitFunding) {
		return true
	}

	return false
}

// SetJitFunding gets a reference to the given JitFundingApi and assigns it to the JitFunding field.
func (o *GpaResponse) SetJitFunding(v JitFundingApi) {
	o.JitFunding = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *GpaResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *GpaResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *GpaResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *GpaResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *GpaResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetResponse returns the Response field value
func (o *GpaResponse) GetResponse() Response {
	if o == nil {
		var ret Response
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetResponseOk() (*Response, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *GpaResponse) SetResponse(v Response) {
	o.Response = v
}

// GetState returns the State field value
func (o *GpaResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *GpaResponse) SetState(v string) {
	o.State = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GpaResponse) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GpaResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *GpaResponse) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value
func (o *GpaResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GpaResponse) SetToken(v string) {
	o.Token = v
}

// GetTransactionToken returns the TransactionToken field value
func (o *GpaResponse) GetTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionToken
}

// GetTransactionTokenOk returns a tuple with the TransactionToken field value
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionToken, true
}

// SetTransactionToken sets field value
func (o *GpaResponse) SetTransactionToken(v string) {
	o.TransactionToken = v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *GpaResponse) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpaResponse) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *GpaResponse) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *GpaResponse) SetUserToken(v string) {
	o.UserToken = &v
}

func (o GpaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GpaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	toSerialize["funding"] = o.Funding
	if !IsNil(o.FundingSourceAddressToken) {
		toSerialize["funding_source_address_token"] = o.FundingSourceAddressToken
	}
	toSerialize["funding_source_token"] = o.FundingSourceToken
	if !IsNil(o.GatewayMessage) {
		toSerialize["gateway_message"] = o.GatewayMessage
	}
	if !IsNil(o.GatewayToken) {
		toSerialize["gateway_token"] = o.GatewayToken
	}
	if !IsNil(o.JitFunding) {
		toSerialize["jit_funding"] = o.JitFunding
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	toSerialize["response"] = o.Response
	toSerialize["state"] = o.State
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["token"] = o.Token
	toSerialize["transaction_token"] = o.TransactionToken
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

type NullableGpaResponse struct {
	value *GpaResponse
	isSet bool
}

func (v NullableGpaResponse) Get() *GpaResponse {
	return v.value
}

func (v *NullableGpaResponse) Set(val *GpaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGpaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGpaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpaResponse(val *GpaResponse) *NullableGpaResponse {
	return &NullableGpaResponse{value: val, isSet: true}
}

func (v NullableGpaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


