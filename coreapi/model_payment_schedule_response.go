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

// checks if the PaymentScheduleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentScheduleResponse{}

// PaymentScheduleResponse A future one-time or recurring payment schedule.
type PaymentScheduleResponse struct {
	// Unique identifier of the credit account on which the payment schedule is made.
	AccountToken string `json:"account_token"`
	// Amount of the payment.  Returned if the `amount_category` is `FIXED`.
	Amount NullableFloat32 `json:"amount,omitempty"`
	AmountCategory PaymentScheduleAmountCategory `json:"amount_category"`
	// Date and time when the payment schedule was created on Marqeta's credit platform, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Description of the payment schedule.
	Description *string `json:"description,omitempty"`
	Frequency PaymentScheduleFrequency `json:"frequency"`
	// The date to make a one-time payment.  Returned if `frequency` is `ONCE`.
	NextPaymentImpactDate *string `json:"next_payment_impact_date,omitempty"`
	// Day on which monthly payments are made.  Returned if the `frequency` is `MONTHLY`.
	PaymentDay *string `json:"payment_day,omitempty"`
	// Unique identifier of a payment source.
	PaymentSourceToken string `json:"payment_source_token"`
	Status PaymentScheduleStatus `json:"status"`
	// Unique identifier of the payment schedule.
	Token string `json:"token"`
	// Date and time when the payment schedule was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}

// NewPaymentScheduleResponse instantiates a new PaymentScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentScheduleResponse(accountToken string, amountCategory PaymentScheduleAmountCategory, currencyCode CurrencyCode, frequency PaymentScheduleFrequency, paymentSourceToken string, status PaymentScheduleStatus, token string) *PaymentScheduleResponse {
	this := PaymentScheduleResponse{}
	this.AccountToken = accountToken
	this.AmountCategory = amountCategory
	this.CurrencyCode = currencyCode
	this.Frequency = frequency
	this.PaymentSourceToken = paymentSourceToken
	this.Status = status
	this.Token = token
	return &this
}

// NewPaymentScheduleResponseWithDefaults instantiates a new PaymentScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentScheduleResponseWithDefaults() *PaymentScheduleResponse {
	this := PaymentScheduleResponse{}
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = currencyCode
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *PaymentScheduleResponse) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *PaymentScheduleResponse) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentScheduleResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentScheduleResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentScheduleResponse) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *PaymentScheduleResponse) SetAmount(v float32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *PaymentScheduleResponse) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *PaymentScheduleResponse) UnsetAmount() {
	o.Amount.Unset()
}

// GetAmountCategory returns the AmountCategory field value
func (o *PaymentScheduleResponse) GetAmountCategory() PaymentScheduleAmountCategory {
	if o == nil {
		var ret PaymentScheduleAmountCategory
		return ret
	}

	return o.AmountCategory
}

// GetAmountCategoryOk returns a tuple with the AmountCategory field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetAmountCategoryOk() (*PaymentScheduleAmountCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCategory, true
}

// SetAmountCategory sets field value
func (o *PaymentScheduleResponse) SetAmountCategory(v PaymentScheduleAmountCategory) {
	o.AmountCategory = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *PaymentScheduleResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *PaymentScheduleResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *PaymentScheduleResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *PaymentScheduleResponse) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *PaymentScheduleResponse) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentScheduleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentScheduleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentScheduleResponse) SetDescription(v string) {
	o.Description = &v
}

// GetFrequency returns the Frequency field value
func (o *PaymentScheduleResponse) GetFrequency() PaymentScheduleFrequency {
	if o == nil {
		var ret PaymentScheduleFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetFrequencyOk() (*PaymentScheduleFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *PaymentScheduleResponse) SetFrequency(v PaymentScheduleFrequency) {
	o.Frequency = v
}

// GetNextPaymentImpactDate returns the NextPaymentImpactDate field value if set, zero value otherwise.
func (o *PaymentScheduleResponse) GetNextPaymentImpactDate() string {
	if o == nil || IsNil(o.NextPaymentImpactDate) {
		var ret string
		return ret
	}
	return *o.NextPaymentImpactDate
}

// GetNextPaymentImpactDateOk returns a tuple with the NextPaymentImpactDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetNextPaymentImpactDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextPaymentImpactDate) {
		return nil, false
	}
	return o.NextPaymentImpactDate, true
}

// HasNextPaymentImpactDate returns a boolean if a field has been set.
func (o *PaymentScheduleResponse) HasNextPaymentImpactDate() bool {
	if o != nil && !IsNil(o.NextPaymentImpactDate) {
		return true
	}

	return false
}

// SetNextPaymentImpactDate gets a reference to the given string and assigns it to the NextPaymentImpactDate field.
func (o *PaymentScheduleResponse) SetNextPaymentImpactDate(v string) {
	o.NextPaymentImpactDate = &v
}

// GetPaymentDay returns the PaymentDay field value if set, zero value otherwise.
func (o *PaymentScheduleResponse) GetPaymentDay() string {
	if o == nil || IsNil(o.PaymentDay) {
		var ret string
		return ret
	}
	return *o.PaymentDay
}

// GetPaymentDayOk returns a tuple with the PaymentDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetPaymentDayOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentDay) {
		return nil, false
	}
	return o.PaymentDay, true
}

// HasPaymentDay returns a boolean if a field has been set.
func (o *PaymentScheduleResponse) HasPaymentDay() bool {
	if o != nil && !IsNil(o.PaymentDay) {
		return true
	}

	return false
}

// SetPaymentDay gets a reference to the given string and assigns it to the PaymentDay field.
func (o *PaymentScheduleResponse) SetPaymentDay(v string) {
	o.PaymentDay = &v
}

// GetPaymentSourceToken returns the PaymentSourceToken field value
func (o *PaymentScheduleResponse) GetPaymentSourceToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentSourceToken
}

// GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetPaymentSourceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentSourceToken, true
}

// SetPaymentSourceToken sets field value
func (o *PaymentScheduleResponse) SetPaymentSourceToken(v string) {
	o.PaymentSourceToken = v
}

// GetStatus returns the Status field value
func (o *PaymentScheduleResponse) GetStatus() PaymentScheduleStatus {
	if o == nil {
		var ret PaymentScheduleStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetStatusOk() (*PaymentScheduleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentScheduleResponse) SetStatus(v PaymentScheduleStatus) {
	o.Status = v
}

// GetToken returns the Token field value
func (o *PaymentScheduleResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PaymentScheduleResponse) SetToken(v string) {
	o.Token = v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *PaymentScheduleResponse) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentScheduleResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *PaymentScheduleResponse) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *PaymentScheduleResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

func (o PaymentScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_token"] = o.AccountToken
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	toSerialize["amount_category"] = o.AmountCategory
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["frequency"] = o.Frequency
	if !IsNil(o.NextPaymentImpactDate) {
		toSerialize["next_payment_impact_date"] = o.NextPaymentImpactDate
	}
	if !IsNil(o.PaymentDay) {
		toSerialize["payment_day"] = o.PaymentDay
	}
	toSerialize["payment_source_token"] = o.PaymentSourceToken
	toSerialize["status"] = o.Status
	toSerialize["token"] = o.Token
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}
	return toSerialize, nil
}

type NullablePaymentScheduleResponse struct {
	value *PaymentScheduleResponse
	isSet bool
}

func (v NullablePaymentScheduleResponse) Get() *PaymentScheduleResponse {
	return v.value
}

func (v *NullablePaymentScheduleResponse) Set(val *PaymentScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentScheduleResponse(val *PaymentScheduleResponse) *NullablePaymentScheduleResponse {
	return &NullablePaymentScheduleResponse{value: val, isSet: true}
}

func (v NullablePaymentScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


