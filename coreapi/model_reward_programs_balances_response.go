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

// checks if the RewardProgramsBalancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewardProgramsBalancesResponse{}

// RewardProgramsBalancesResponse struct for RewardProgramsBalancesResponse
type RewardProgramsBalancesResponse struct {
	// Closing date of the billing cycle for which rewards were accrued, in UTC.
	BillingCycleClosingDate time.Time `json:"billing_cycle_closing_date"`
	// Opening date of the billing cycle for which rewards were accrued, in UTC.
	BillingCycleOpeningDate time.Time `json:"billing_cycle_opening_date"`
	// The net balance for a billing cycle, which is total amount spent during a billing cycle, minus any refunds or reversals. Used to determine reward accrual.
	NetBalance float32 `json:"net_balance"`
	// The pending balance of the rewards accrued for the current billing cycle. Pending rewards cannot be redeemed.
	PendingRewardBalance float32 `json:"pending_reward_balance"`
	// The reward percentage applied to the balance for the current billing cycle. Determined by the reward rules config.
	Percentage int32 `json:"percentage"`
	// Unique identifier of reward program for which to return balances.
	RewardProgramToken string `json:"reward_program_token"`
	// The total balance of the rewards accrued to date minus the rewards redeemed to date.
	TotalRewardBalance float32 `json:"total_reward_balance"`
}

// NewRewardProgramsBalancesResponse instantiates a new RewardProgramsBalancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardProgramsBalancesResponse(billingCycleClosingDate time.Time, billingCycleOpeningDate time.Time, netBalance float32, pendingRewardBalance float32, percentage int32, rewardProgramToken string, totalRewardBalance float32) *RewardProgramsBalancesResponse {
	this := RewardProgramsBalancesResponse{}
	this.BillingCycleClosingDate = billingCycleClosingDate
	this.BillingCycleOpeningDate = billingCycleOpeningDate
	this.NetBalance = netBalance
	this.PendingRewardBalance = pendingRewardBalance
	this.Percentage = percentage
	this.RewardProgramToken = rewardProgramToken
	this.TotalRewardBalance = totalRewardBalance
	return &this
}

// NewRewardProgramsBalancesResponseWithDefaults instantiates a new RewardProgramsBalancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardProgramsBalancesResponseWithDefaults() *RewardProgramsBalancesResponse {
	this := RewardProgramsBalancesResponse{}
	return &this
}

// GetBillingCycleClosingDate returns the BillingCycleClosingDate field value
func (o *RewardProgramsBalancesResponse) GetBillingCycleClosingDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BillingCycleClosingDate
}

// GetBillingCycleClosingDateOk returns a tuple with the BillingCycleClosingDate field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsBalancesResponse) GetBillingCycleClosingDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingCycleClosingDate, true
}

// SetBillingCycleClosingDate sets field value
func (o *RewardProgramsBalancesResponse) SetBillingCycleClosingDate(v time.Time) {
	o.BillingCycleClosingDate = v
}

// GetBillingCycleOpeningDate returns the BillingCycleOpeningDate field value
func (o *RewardProgramsBalancesResponse) GetBillingCycleOpeningDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BillingCycleOpeningDate
}

// GetBillingCycleOpeningDateOk returns a tuple with the BillingCycleOpeningDate field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsBalancesResponse) GetBillingCycleOpeningDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingCycleOpeningDate, true
}

// SetBillingCycleOpeningDate sets field value
func (o *RewardProgramsBalancesResponse) SetBillingCycleOpeningDate(v time.Time) {
	o.BillingCycleOpeningDate = v
}

// GetNetBalance returns the NetBalance field value
func (o *RewardProgramsBalancesResponse) GetNetBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetBalance
}

// GetNetBalanceOk returns a tuple with the NetBalance field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsBalancesResponse) GetNetBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetBalance, true
}

// SetNetBalance sets field value
func (o *RewardProgramsBalancesResponse) SetNetBalance(v float32) {
	o.NetBalance = v
}

// GetPendingRewardBalance returns the PendingRewardBalance field value
func (o *RewardProgramsBalancesResponse) GetPendingRewardBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingRewardBalance
}

// GetPendingRewardBalanceOk returns a tuple with the PendingRewardBalance field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsBalancesResponse) GetPendingRewardBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingRewardBalance, true
}

// SetPendingRewardBalance sets field value
func (o *RewardProgramsBalancesResponse) SetPendingRewardBalance(v float32) {
	o.PendingRewardBalance = v
}

// GetPercentage returns the Percentage field value
func (o *RewardProgramsBalancesResponse) GetPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsBalancesResponse) GetPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *RewardProgramsBalancesResponse) SetPercentage(v int32) {
	o.Percentage = v
}

// GetRewardProgramToken returns the RewardProgramToken field value
func (o *RewardProgramsBalancesResponse) GetRewardProgramToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardProgramToken
}

// GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsBalancesResponse) GetRewardProgramTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardProgramToken, true
}

// SetRewardProgramToken sets field value
func (o *RewardProgramsBalancesResponse) SetRewardProgramToken(v string) {
	o.RewardProgramToken = v
}

// GetTotalRewardBalance returns the TotalRewardBalance field value
func (o *RewardProgramsBalancesResponse) GetTotalRewardBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalRewardBalance
}

// GetTotalRewardBalanceOk returns a tuple with the TotalRewardBalance field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsBalancesResponse) GetTotalRewardBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRewardBalance, true
}

// SetTotalRewardBalance sets field value
func (o *RewardProgramsBalancesResponse) SetTotalRewardBalance(v float32) {
	o.TotalRewardBalance = v
}

func (o RewardProgramsBalancesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewardProgramsBalancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billing_cycle_closing_date"] = o.BillingCycleClosingDate
	toSerialize["billing_cycle_opening_date"] = o.BillingCycleOpeningDate
	toSerialize["net_balance"] = o.NetBalance
	toSerialize["pending_reward_balance"] = o.PendingRewardBalance
	toSerialize["percentage"] = o.Percentage
	toSerialize["reward_program_token"] = o.RewardProgramToken
	toSerialize["total_reward_balance"] = o.TotalRewardBalance
	return toSerialize, nil
}

type NullableRewardProgramsBalancesResponse struct {
	value *RewardProgramsBalancesResponse
	isSet bool
}

func (v NullableRewardProgramsBalancesResponse) Get() *RewardProgramsBalancesResponse {
	return v.value
}

func (v *NullableRewardProgramsBalancesResponse) Set(val *RewardProgramsBalancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardProgramsBalancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardProgramsBalancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardProgramsBalancesResponse(val *RewardProgramsBalancesResponse) *NullableRewardProgramsBalancesResponse {
	return &NullableRewardProgramsBalancesResponse{value: val, isSet: true}
}

func (v NullableRewardProgramsBalancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardProgramsBalancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

