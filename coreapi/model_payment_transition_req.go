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

// checks if the PaymentTransitionReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentTransitionReq{}

// PaymentTransitionReq Request used to transition the status of a payment.
type PaymentTransitionReq struct {
	RefundDetails *RefundDetails `json:"refund_details,omitempty"`
	Status PaymentStatus `json:"status"`
	// Unique identifier of the payment status transition.
	Token *string `json:"token,omitempty"`
}

// NewPaymentTransitionReq instantiates a new PaymentTransitionReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentTransitionReq(status PaymentStatus) *PaymentTransitionReq {
	this := PaymentTransitionReq{}
	this.Status = status
	return &this
}

// NewPaymentTransitionReqWithDefaults instantiates a new PaymentTransitionReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentTransitionReqWithDefaults() *PaymentTransitionReq {
	this := PaymentTransitionReq{}
	return &this
}

// GetRefundDetails returns the RefundDetails field value if set, zero value otherwise.
func (o *PaymentTransitionReq) GetRefundDetails() RefundDetails {
	if o == nil || IsNil(o.RefundDetails) {
		var ret RefundDetails
		return ret
	}
	return *o.RefundDetails
}

// GetRefundDetailsOk returns a tuple with the RefundDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransitionReq) GetRefundDetailsOk() (*RefundDetails, bool) {
	if o == nil || IsNil(o.RefundDetails) {
		return nil, false
	}
	return o.RefundDetails, true
}

// HasRefundDetails returns a boolean if a field has been set.
func (o *PaymentTransitionReq) HasRefundDetails() bool {
	if o != nil && !IsNil(o.RefundDetails) {
		return true
	}

	return false
}

// SetRefundDetails gets a reference to the given RefundDetails and assigns it to the RefundDetails field.
func (o *PaymentTransitionReq) SetRefundDetails(v RefundDetails) {
	o.RefundDetails = &v
}

// GetStatus returns the Status field value
func (o *PaymentTransitionReq) GetStatus() PaymentStatus {
	if o == nil {
		var ret PaymentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentTransitionReq) GetStatusOk() (*PaymentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentTransitionReq) SetStatus(v PaymentStatus) {
	o.Status = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PaymentTransitionReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransitionReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PaymentTransitionReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PaymentTransitionReq) SetToken(v string) {
	o.Token = &v
}

func (o PaymentTransitionReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentTransitionReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefundDetails) {
		toSerialize["refund_details"] = o.RefundDetails
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullablePaymentTransitionReq struct {
	value *PaymentTransitionReq
	isSet bool
}

func (v NullablePaymentTransitionReq) Get() *PaymentTransitionReq {
	return v.value
}

func (v *NullablePaymentTransitionReq) Set(val *PaymentTransitionReq) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTransitionReq) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTransitionReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTransitionReq(val *PaymentTransitionReq) *NullablePaymentTransitionReq {
	return &NullablePaymentTransitionReq{value: val, isSet: true}
}

func (v NullablePaymentTransitionReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTransitionReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


