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

// checks if the RedemptionsBySettlementDateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedemptionsBySettlementDateResponse{}

// RedemptionsBySettlementDateResponse Return redemptions.
type RedemptionsBySettlementDateResponse struct {
	// token of account the redemption is for.
	AccountToken string `json:"account_token"`
	Amount float32 `json:"amount"`
	// yyyy-MM-ddThh:mm:ssZ
	CompletionTime *time.Time `json:"completion_time,omitempty"`
	// yyyy-MM-ddThh:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	Destination DestinationType `json:"destination"`
	Note string `json:"note"`
	// Identifier of the redemption.
	RedemptionToken string `json:"redemption_token"`
	// Token of reward program the redemption is for.
	RewardProgramToken string `json:"reward_program_token"`
	Status RedemptionStatus `json:"status"`
	Type RedemptionType `json:"type"`
	// yyyy-MM-ddThh:mm:ssZ
	UpdatedTime time.Time `json:"updated_time"`
}

// NewRedemptionsBySettlementDateResponse instantiates a new RedemptionsBySettlementDateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedemptionsBySettlementDateResponse(accountToken string, amount float32, createdTime time.Time, destination DestinationType, note string, redemptionToken string, rewardProgramToken string, status RedemptionStatus, type_ RedemptionType, updatedTime time.Time) *RedemptionsBySettlementDateResponse {
	this := RedemptionsBySettlementDateResponse{}
	this.AccountToken = accountToken
	this.Amount = amount
	this.CreatedTime = createdTime
	this.Destination = destination
	this.Note = note
	this.RedemptionToken = redemptionToken
	this.RewardProgramToken = rewardProgramToken
	this.Status = status
	this.Type = type_
	this.UpdatedTime = updatedTime
	return &this
}

// NewRedemptionsBySettlementDateResponseWithDefaults instantiates a new RedemptionsBySettlementDateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedemptionsBySettlementDateResponseWithDefaults() *RedemptionsBySettlementDateResponse {
	this := RedemptionsBySettlementDateResponse{}
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *RedemptionsBySettlementDateResponse) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *RedemptionsBySettlementDateResponse) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetAmount returns the Amount field value
func (o *RedemptionsBySettlementDateResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RedemptionsBySettlementDateResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *RedemptionsBySettlementDateResponse) GetCompletionTime() time.Time {
	if o == nil || IsNil(o.CompletionTime) {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *RedemptionsBySettlementDateResponse) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *RedemptionsBySettlementDateResponse) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *RedemptionsBySettlementDateResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *RedemptionsBySettlementDateResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetDestination returns the Destination field value
func (o *RedemptionsBySettlementDateResponse) GetDestination() DestinationType {
	if o == nil {
		var ret DestinationType
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetDestinationOk() (*DestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *RedemptionsBySettlementDateResponse) SetDestination(v DestinationType) {
	o.Destination = v
}

// GetNote returns the Note field value
func (o *RedemptionsBySettlementDateResponse) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *RedemptionsBySettlementDateResponse) SetNote(v string) {
	o.Note = v
}

// GetRedemptionToken returns the RedemptionToken field value
func (o *RedemptionsBySettlementDateResponse) GetRedemptionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedemptionToken
}

// GetRedemptionTokenOk returns a tuple with the RedemptionToken field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetRedemptionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedemptionToken, true
}

// SetRedemptionToken sets field value
func (o *RedemptionsBySettlementDateResponse) SetRedemptionToken(v string) {
	o.RedemptionToken = v
}

// GetRewardProgramToken returns the RewardProgramToken field value
func (o *RedemptionsBySettlementDateResponse) GetRewardProgramToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardProgramToken
}

// GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetRewardProgramTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardProgramToken, true
}

// SetRewardProgramToken sets field value
func (o *RedemptionsBySettlementDateResponse) SetRewardProgramToken(v string) {
	o.RewardProgramToken = v
}

// GetStatus returns the Status field value
func (o *RedemptionsBySettlementDateResponse) GetStatus() RedemptionStatus {
	if o == nil {
		var ret RedemptionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetStatusOk() (*RedemptionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RedemptionsBySettlementDateResponse) SetStatus(v RedemptionStatus) {
	o.Status = v
}

// GetType returns the Type field value
func (o *RedemptionsBySettlementDateResponse) GetType() RedemptionType {
	if o == nil {
		var ret RedemptionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetTypeOk() (*RedemptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RedemptionsBySettlementDateResponse) SetType(v RedemptionType) {
	o.Type = v
}

// GetUpdatedTime returns the UpdatedTime field value
func (o *RedemptionsBySettlementDateResponse) GetUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value
// and a boolean to check if the value has been set.
func (o *RedemptionsBySettlementDateResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTime, true
}

// SetUpdatedTime sets field value
func (o *RedemptionsBySettlementDateResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = v
}

func (o RedemptionsBySettlementDateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedemptionsBySettlementDateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_token"] = o.AccountToken
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CompletionTime) {
		toSerialize["completion_time"] = o.CompletionTime
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["destination"] = o.Destination
	toSerialize["note"] = o.Note
	toSerialize["redemption_token"] = o.RedemptionToken
	toSerialize["reward_program_token"] = o.RewardProgramToken
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["updated_time"] = o.UpdatedTime
	return toSerialize, nil
}

type NullableRedemptionsBySettlementDateResponse struct {
	value *RedemptionsBySettlementDateResponse
	isSet bool
}

func (v NullableRedemptionsBySettlementDateResponse) Get() *RedemptionsBySettlementDateResponse {
	return v.value
}

func (v *NullableRedemptionsBySettlementDateResponse) Set(val *RedemptionsBySettlementDateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedemptionsBySettlementDateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedemptionsBySettlementDateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedemptionsBySettlementDateResponse(val *RedemptionsBySettlementDateResponse) *NullableRedemptionsBySettlementDateResponse {
	return &NullableRedemptionsBySettlementDateResponse{value: val, isSet: true}
}

func (v NullableRedemptionsBySettlementDateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedemptionsBySettlementDateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


