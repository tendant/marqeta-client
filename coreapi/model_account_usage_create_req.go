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

// checks if the AccountUsageCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUsageCreateReq{}

// AccountUsageCreateReq Contains information on how a credit account is used and what types of balances are permitted on the account.
type AccountUsageCreateReq struct {
	// Contains one or more annual percentage rates (APRs) associated with the type of balance on the credit account.
	Aprs []AprScheduleCreateReq `json:"aprs"`
	// Contains one or more fees associated with the usage type.
	Fees []AccountFee `json:"fees,omitempty"`
	Type BalanceType `json:"type"`
}

// NewAccountUsageCreateReq instantiates a new AccountUsageCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUsageCreateReq(aprs []AprScheduleCreateReq, type_ BalanceType) *AccountUsageCreateReq {
	this := AccountUsageCreateReq{}
	this.Aprs = aprs
	this.Type = type_
	return &this
}

// NewAccountUsageCreateReqWithDefaults instantiates a new AccountUsageCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUsageCreateReqWithDefaults() *AccountUsageCreateReq {
	this := AccountUsageCreateReq{}
	return &this
}

// GetAprs returns the Aprs field value
func (o *AccountUsageCreateReq) GetAprs() []AprScheduleCreateReq {
	if o == nil {
		var ret []AprScheduleCreateReq
		return ret
	}

	return o.Aprs
}

// GetAprsOk returns a tuple with the Aprs field value
// and a boolean to check if the value has been set.
func (o *AccountUsageCreateReq) GetAprsOk() ([]AprScheduleCreateReq, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aprs, true
}

// SetAprs sets field value
func (o *AccountUsageCreateReq) SetAprs(v []AprScheduleCreateReq) {
	o.Aprs = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *AccountUsageCreateReq) GetFees() []AccountFee {
	if o == nil || IsNil(o.Fees) {
		var ret []AccountFee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUsageCreateReq) GetFeesOk() ([]AccountFee, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *AccountUsageCreateReq) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []AccountFee and assigns it to the Fees field.
func (o *AccountUsageCreateReq) SetFees(v []AccountFee) {
	o.Fees = v
}

// GetType returns the Type field value
func (o *AccountUsageCreateReq) GetType() BalanceType {
	if o == nil {
		var ret BalanceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountUsageCreateReq) GetTypeOk() (*BalanceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountUsageCreateReq) SetType(v BalanceType) {
	o.Type = v
}

func (o AccountUsageCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUsageCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aprs"] = o.Aprs
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAccountUsageCreateReq struct {
	value *AccountUsageCreateReq
	isSet bool
}

func (v NullableAccountUsageCreateReq) Get() *AccountUsageCreateReq {
	return v.value
}

func (v *NullableAccountUsageCreateReq) Set(val *AccountUsageCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUsageCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUsageCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUsageCreateReq(val *AccountUsageCreateReq) *NullableAccountUsageCreateReq {
	return &NullableAccountUsageCreateReq{value: val, isSet: true}
}

func (v NullableAccountUsageCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUsageCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


