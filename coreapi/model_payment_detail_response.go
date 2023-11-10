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

// checks if the PaymentDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentDetailResponse{}

// PaymentDetailResponse Response containing payment details with transition history
type PaymentDetailResponse struct {
	// Unique identifier of the credit account on which the payment is made.
	AccountToken string `json:"account_token"`
	// Total amount of the payment.
	Amount float32 `json:"amount"`
	// Date and time when the payment was created on Marqeta's credit platform, in UTC.
	CreatedTime time.Time `json:"created_time"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Description of the payment.
	Description string `json:"description"`
	// After a payment completes, the number of days to hold the available credit on the account before increasing it.
	HoldDays int32 `json:"hold_days"`
	// Date and time when the available credit hold is released.
	HoldEndTime *time.Time `json:"hold_end_time,omitempty"`
	// Whether the available credit hold was manually released for this payment.
	IsManuallyReleased *bool `json:"is_manually_released,omitempty"`
	// Customer-defined additional information about the payment (for example, a check number).
	Metadata *string `json:"metadata,omitempty"`
	// Method of payment.
	Method string `json:"method"`
	// Whether the available credit is on hold for this payment.
	OnHold *bool `json:"on_hold,omitempty"`
	// Unique identifier of the payment schedule.
	PaymentScheduleToken *string `json:"payment_schedule_token,omitempty"`
	// Unique identifier of the payment source. Required for ACH payments.
	PaymentSourceToken *string `json:"payment_source_token,omitempty"`
	RefundDetails NullableRefundDetailsResponse `json:"refund_details,omitempty"`
	ReturnedDetails NullableReturnedDetails `json:"returned_details,omitempty"`
	Status PaymentStatus `json:"status"`
	// Unique identifier of the payment.  If in the `detail_object`, unique identifier of the detail object.
	Token string `json:"token"`
	// Contains one or more `transitions` objects, which contain information on a payment status transition.
	Transitions []PaymentTransitionResponse `json:"transitions"`
	// Date and time when the payment was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime time.Time `json:"updated_time"`
}

// NewPaymentDetailResponse instantiates a new PaymentDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDetailResponse(accountToken string, amount float32, createdTime time.Time, currencyCode CurrencyCode, description string, holdDays int32, method string, status PaymentStatus, token string, transitions []PaymentTransitionResponse, updatedTime time.Time) *PaymentDetailResponse {
	this := PaymentDetailResponse{}
	this.AccountToken = accountToken
	this.Amount = amount
	this.CreatedTime = createdTime
	this.CurrencyCode = currencyCode
	this.Description = description
	this.HoldDays = holdDays
	var isManuallyReleased bool = false
	this.IsManuallyReleased = &isManuallyReleased
	this.Method = method
	var onHold bool = false
	this.OnHold = &onHold
	this.Status = status
	this.Token = token
	this.Transitions = transitions
	this.UpdatedTime = updatedTime
	return &this
}

// NewPaymentDetailResponseWithDefaults instantiates a new PaymentDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDetailResponseWithDefaults() *PaymentDetailResponse {
	this := PaymentDetailResponse{}
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = currencyCode
	var holdDays int32 = 0
	this.HoldDays = holdDays
	var isManuallyReleased bool = false
	this.IsManuallyReleased = &isManuallyReleased
	var onHold bool = false
	this.OnHold = &onHold
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *PaymentDetailResponse) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *PaymentDetailResponse) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetAmount returns the Amount field value
func (o *PaymentDetailResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentDetailResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *PaymentDetailResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *PaymentDetailResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *PaymentDetailResponse) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *PaymentDetailResponse) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetDescription returns the Description field value
func (o *PaymentDetailResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PaymentDetailResponse) SetDescription(v string) {
	o.Description = v
}

// GetHoldDays returns the HoldDays field value
func (o *PaymentDetailResponse) GetHoldDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HoldDays
}

// GetHoldDaysOk returns a tuple with the HoldDays field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetHoldDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HoldDays, true
}

// SetHoldDays sets field value
func (o *PaymentDetailResponse) SetHoldDays(v int32) {
	o.HoldDays = v
}

// GetHoldEndTime returns the HoldEndTime field value if set, zero value otherwise.
func (o *PaymentDetailResponse) GetHoldEndTime() time.Time {
	if o == nil || IsNil(o.HoldEndTime) {
		var ret time.Time
		return ret
	}
	return *o.HoldEndTime
}

// GetHoldEndTimeOk returns a tuple with the HoldEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetHoldEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.HoldEndTime) {
		return nil, false
	}
	return o.HoldEndTime, true
}

// HasHoldEndTime returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasHoldEndTime() bool {
	if o != nil && !IsNil(o.HoldEndTime) {
		return true
	}

	return false
}

// SetHoldEndTime gets a reference to the given time.Time and assigns it to the HoldEndTime field.
func (o *PaymentDetailResponse) SetHoldEndTime(v time.Time) {
	o.HoldEndTime = &v
}

// GetIsManuallyReleased returns the IsManuallyReleased field value if set, zero value otherwise.
func (o *PaymentDetailResponse) GetIsManuallyReleased() bool {
	if o == nil || IsNil(o.IsManuallyReleased) {
		var ret bool
		return ret
	}
	return *o.IsManuallyReleased
}

// GetIsManuallyReleasedOk returns a tuple with the IsManuallyReleased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetIsManuallyReleasedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManuallyReleased) {
		return nil, false
	}
	return o.IsManuallyReleased, true
}

// HasIsManuallyReleased returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasIsManuallyReleased() bool {
	if o != nil && !IsNil(o.IsManuallyReleased) {
		return true
	}

	return false
}

// SetIsManuallyReleased gets a reference to the given bool and assigns it to the IsManuallyReleased field.
func (o *PaymentDetailResponse) SetIsManuallyReleased(v bool) {
	o.IsManuallyReleased = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PaymentDetailResponse) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *PaymentDetailResponse) SetMetadata(v string) {
	o.Metadata = &v
}

// GetMethod returns the Method field value
func (o *PaymentDetailResponse) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *PaymentDetailResponse) SetMethod(v string) {
	o.Method = v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise.
func (o *PaymentDetailResponse) GetOnHold() bool {
	if o == nil || IsNil(o.OnHold) {
		var ret bool
		return ret
	}
	return *o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetOnHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.OnHold) {
		return nil, false
	}
	return o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasOnHold() bool {
	if o != nil && !IsNil(o.OnHold) {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given bool and assigns it to the OnHold field.
func (o *PaymentDetailResponse) SetOnHold(v bool) {
	o.OnHold = &v
}

// GetPaymentScheduleToken returns the PaymentScheduleToken field value if set, zero value otherwise.
func (o *PaymentDetailResponse) GetPaymentScheduleToken() string {
	if o == nil || IsNil(o.PaymentScheduleToken) {
		var ret string
		return ret
	}
	return *o.PaymentScheduleToken
}

// GetPaymentScheduleTokenOk returns a tuple with the PaymentScheduleToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetPaymentScheduleTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentScheduleToken) {
		return nil, false
	}
	return o.PaymentScheduleToken, true
}

// HasPaymentScheduleToken returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasPaymentScheduleToken() bool {
	if o != nil && !IsNil(o.PaymentScheduleToken) {
		return true
	}

	return false
}

// SetPaymentScheduleToken gets a reference to the given string and assigns it to the PaymentScheduleToken field.
func (o *PaymentDetailResponse) SetPaymentScheduleToken(v string) {
	o.PaymentScheduleToken = &v
}

// GetPaymentSourceToken returns the PaymentSourceToken field value if set, zero value otherwise.
func (o *PaymentDetailResponse) GetPaymentSourceToken() string {
	if o == nil || IsNil(o.PaymentSourceToken) {
		var ret string
		return ret
	}
	return *o.PaymentSourceToken
}

// GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetPaymentSourceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentSourceToken) {
		return nil, false
	}
	return o.PaymentSourceToken, true
}

// HasPaymentSourceToken returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasPaymentSourceToken() bool {
	if o != nil && !IsNil(o.PaymentSourceToken) {
		return true
	}

	return false
}

// SetPaymentSourceToken gets a reference to the given string and assigns it to the PaymentSourceToken field.
func (o *PaymentDetailResponse) SetPaymentSourceToken(v string) {
	o.PaymentSourceToken = &v
}

// GetRefundDetails returns the RefundDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDetailResponse) GetRefundDetails() RefundDetailsResponse {
	if o == nil || IsNil(o.RefundDetails.Get()) {
		var ret RefundDetailsResponse
		return ret
	}
	return *o.RefundDetails.Get()
}

// GetRefundDetailsOk returns a tuple with the RefundDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDetailResponse) GetRefundDetailsOk() (*RefundDetailsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefundDetails.Get(), o.RefundDetails.IsSet()
}

// HasRefundDetails returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasRefundDetails() bool {
	if o != nil && o.RefundDetails.IsSet() {
		return true
	}

	return false
}

// SetRefundDetails gets a reference to the given NullableRefundDetailsResponse and assigns it to the RefundDetails field.
func (o *PaymentDetailResponse) SetRefundDetails(v RefundDetailsResponse) {
	o.RefundDetails.Set(&v)
}
// SetRefundDetailsNil sets the value for RefundDetails to be an explicit nil
func (o *PaymentDetailResponse) SetRefundDetailsNil() {
	o.RefundDetails.Set(nil)
}

// UnsetRefundDetails ensures that no value is present for RefundDetails, not even an explicit nil
func (o *PaymentDetailResponse) UnsetRefundDetails() {
	o.RefundDetails.Unset()
}

// GetReturnedDetails returns the ReturnedDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentDetailResponse) GetReturnedDetails() ReturnedDetails {
	if o == nil || IsNil(o.ReturnedDetails.Get()) {
		var ret ReturnedDetails
		return ret
	}
	return *o.ReturnedDetails.Get()
}

// GetReturnedDetailsOk returns a tuple with the ReturnedDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentDetailResponse) GetReturnedDetailsOk() (*ReturnedDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnedDetails.Get(), o.ReturnedDetails.IsSet()
}

// HasReturnedDetails returns a boolean if a field has been set.
func (o *PaymentDetailResponse) HasReturnedDetails() bool {
	if o != nil && o.ReturnedDetails.IsSet() {
		return true
	}

	return false
}

// SetReturnedDetails gets a reference to the given NullableReturnedDetails and assigns it to the ReturnedDetails field.
func (o *PaymentDetailResponse) SetReturnedDetails(v ReturnedDetails) {
	o.ReturnedDetails.Set(&v)
}
// SetReturnedDetailsNil sets the value for ReturnedDetails to be an explicit nil
func (o *PaymentDetailResponse) SetReturnedDetailsNil() {
	o.ReturnedDetails.Set(nil)
}

// UnsetReturnedDetails ensures that no value is present for ReturnedDetails, not even an explicit nil
func (o *PaymentDetailResponse) UnsetReturnedDetails() {
	o.ReturnedDetails.Unset()
}

// GetStatus returns the Status field value
func (o *PaymentDetailResponse) GetStatus() PaymentStatus {
	if o == nil {
		var ret PaymentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetStatusOk() (*PaymentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentDetailResponse) SetStatus(v PaymentStatus) {
	o.Status = v
}

// GetToken returns the Token field value
func (o *PaymentDetailResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PaymentDetailResponse) SetToken(v string) {
	o.Token = v
}

// GetTransitions returns the Transitions field value
func (o *PaymentDetailResponse) GetTransitions() []PaymentTransitionResponse {
	if o == nil {
		var ret []PaymentTransitionResponse
		return ret
	}

	return o.Transitions
}

// GetTransitionsOk returns a tuple with the Transitions field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetTransitionsOk() ([]PaymentTransitionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transitions, true
}

// SetTransitions sets field value
func (o *PaymentDetailResponse) SetTransitions(v []PaymentTransitionResponse) {
	o.Transitions = v
}

// GetUpdatedTime returns the UpdatedTime field value
func (o *PaymentDetailResponse) GetUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value
// and a boolean to check if the value has been set.
func (o *PaymentDetailResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTime, true
}

// SetUpdatedTime sets field value
func (o *PaymentDetailResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = v
}

func (o PaymentDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_token"] = o.AccountToken
	toSerialize["amount"] = o.Amount
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["description"] = o.Description
	toSerialize["hold_days"] = o.HoldDays
	if !IsNil(o.HoldEndTime) {
		toSerialize["hold_end_time"] = o.HoldEndTime
	}
	if !IsNil(o.IsManuallyReleased) {
		toSerialize["is_manually_released"] = o.IsManuallyReleased
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["method"] = o.Method
	if !IsNil(o.OnHold) {
		toSerialize["on_hold"] = o.OnHold
	}
	if !IsNil(o.PaymentScheduleToken) {
		toSerialize["payment_schedule_token"] = o.PaymentScheduleToken
	}
	if !IsNil(o.PaymentSourceToken) {
		toSerialize["payment_source_token"] = o.PaymentSourceToken
	}
	if o.RefundDetails.IsSet() {
		toSerialize["refund_details"] = o.RefundDetails.Get()
	}
	if o.ReturnedDetails.IsSet() {
		toSerialize["returned_details"] = o.ReturnedDetails.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["token"] = o.Token
	toSerialize["transitions"] = o.Transitions
	toSerialize["updated_time"] = o.UpdatedTime
	return toSerialize, nil
}

type NullablePaymentDetailResponse struct {
	value *PaymentDetailResponse
	isSet bool
}

func (v NullablePaymentDetailResponse) Get() *PaymentDetailResponse {
	return v.value
}

func (v *NullablePaymentDetailResponse) Set(val *PaymentDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDetailResponse(val *PaymentDetailResponse) *NullablePaymentDetailResponse {
	return &NullablePaymentDetailResponse{value: val, isSet: true}
}

func (v NullablePaymentDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


