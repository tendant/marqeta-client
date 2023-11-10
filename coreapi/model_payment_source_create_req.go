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

// checks if the PaymentSourceCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentSourceCreateReq{}

// PaymentSourceCreateReq Contains information about a payment source.
type PaymentSourceCreateReq struct {
	// Account number of the payment source.
	AccountNumber string `json:"account_number"`
	// Unique identifier of the credit account receiving the payment.
	AccountToken string `json:"account_token"`
	// Name of the individual or business who owns the payment source.
	Name string `json:"name"`
	// Type of payment source owner.
	Owner *string `json:"owner,omitempty"`
	// Routing number of the payment source.
	RoutingNumber string `json:"routing_number"`
	// Type of payment source.
	SourceType string `json:"source_type"`
	// Unique identifier of the payment source.
	Token *string `json:"token,omitempty"`
	// Unique identifier of the user making the payment.
	UserToken string `json:"user_token"`
	// Additional information on the verification.
	VerificationNotes *string `json:"verification_notes,omitempty"`
	// Whether to override the verification process.
	VerificationOverride bool `json:"verification_override"`
}

// NewPaymentSourceCreateReq instantiates a new PaymentSourceCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentSourceCreateReq(accountNumber string, accountToken string, name string, routingNumber string, sourceType string, userToken string, verificationOverride bool) *PaymentSourceCreateReq {
	this := PaymentSourceCreateReq{}
	this.AccountNumber = accountNumber
	this.AccountToken = accountToken
	this.Name = name
	this.RoutingNumber = routingNumber
	this.SourceType = sourceType
	this.UserToken = userToken
	this.VerificationOverride = verificationOverride
	return &this
}

// NewPaymentSourceCreateReqWithDefaults instantiates a new PaymentSourceCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentSourceCreateReqWithDefaults() *PaymentSourceCreateReq {
	this := PaymentSourceCreateReq{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *PaymentSourceCreateReq) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *PaymentSourceCreateReq) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAccountToken returns the AccountToken field value
func (o *PaymentSourceCreateReq) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *PaymentSourceCreateReq) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetName returns the Name field value
func (o *PaymentSourceCreateReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentSourceCreateReq) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PaymentSourceCreateReq) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PaymentSourceCreateReq) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *PaymentSourceCreateReq) SetOwner(v string) {
	o.Owner = &v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *PaymentSourceCreateReq) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *PaymentSourceCreateReq) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetSourceType returns the SourceType field value
func (o *PaymentSourceCreateReq) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *PaymentSourceCreateReq) SetSourceType(v string) {
	o.SourceType = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PaymentSourceCreateReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PaymentSourceCreateReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PaymentSourceCreateReq) SetToken(v string) {
	o.Token = &v
}

// GetUserToken returns the UserToken field value
func (o *PaymentSourceCreateReq) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *PaymentSourceCreateReq) SetUserToken(v string) {
	o.UserToken = v
}

// GetVerificationNotes returns the VerificationNotes field value if set, zero value otherwise.
func (o *PaymentSourceCreateReq) GetVerificationNotes() string {
	if o == nil || IsNil(o.VerificationNotes) {
		var ret string
		return ret
	}
	return *o.VerificationNotes
}

// GetVerificationNotesOk returns a tuple with the VerificationNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetVerificationNotesOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationNotes) {
		return nil, false
	}
	return o.VerificationNotes, true
}

// HasVerificationNotes returns a boolean if a field has been set.
func (o *PaymentSourceCreateReq) HasVerificationNotes() bool {
	if o != nil && !IsNil(o.VerificationNotes) {
		return true
	}

	return false
}

// SetVerificationNotes gets a reference to the given string and assigns it to the VerificationNotes field.
func (o *PaymentSourceCreateReq) SetVerificationNotes(v string) {
	o.VerificationNotes = &v
}

// GetVerificationOverride returns the VerificationOverride field value
func (o *PaymentSourceCreateReq) GetVerificationOverride() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VerificationOverride
}

// GetVerificationOverrideOk returns a tuple with the VerificationOverride field value
// and a boolean to check if the value has been set.
func (o *PaymentSourceCreateReq) GetVerificationOverrideOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationOverride, true
}

// SetVerificationOverride sets field value
func (o *PaymentSourceCreateReq) SetVerificationOverride(v bool) {
	o.VerificationOverride = v
}

func (o PaymentSourceCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentSourceCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["account_token"] = o.AccountToken
	toSerialize["name"] = o.Name
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	toSerialize["routing_number"] = o.RoutingNumber
	toSerialize["source_type"] = o.SourceType
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["user_token"] = o.UserToken
	if !IsNil(o.VerificationNotes) {
		toSerialize["verification_notes"] = o.VerificationNotes
	}
	toSerialize["verification_override"] = o.VerificationOverride
	return toSerialize, nil
}

type NullablePaymentSourceCreateReq struct {
	value *PaymentSourceCreateReq
	isSet bool
}

func (v NullablePaymentSourceCreateReq) Get() *PaymentSourceCreateReq {
	return v.value
}

func (v *NullablePaymentSourceCreateReq) Set(val *PaymentSourceCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentSourceCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentSourceCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentSourceCreateReq(val *PaymentSourceCreateReq) *NullablePaymentSourceCreateReq {
	return &NullablePaymentSourceCreateReq{value: val, isSet: true}
}

func (v NullablePaymentSourceCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentSourceCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


