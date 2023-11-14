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

// checks if the AccountFundingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountFundingRequest{}

// AccountFundingRequest Information used when funding an account.
type AccountFundingRequest struct {
	// Account holder's first name.
	FirstName string `json:"first_name"`
	// Account holder's last name.
	LastName *string `json:"last_name,omitempty"`
	// Recipient's name.
	ReceiverName *string `json:"receiver_name,omitempty"`
	// Describes the source of the funding.
	FundingSource string `json:"funding_source"`
	// Identifies the account type used for the funding request.
	ReceiverAccountType string `json:"receiver_account_type"`
	// Identifies the purpose of the transaction.
	TransactionPurpose *string `json:"transaction_purpose,omitempty"`
	// Describes the type of transaction.
	TransactionType string `json:"transaction_type"`
}

// NewAccountFundingRequest instantiates a new AccountFundingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountFundingRequest(firstName string, fundingSource string, receiverAccountType string, transactionType string) *AccountFundingRequest {
	this := AccountFundingRequest{}
	this.FirstName = firstName
	this.FundingSource = fundingSource
	this.ReceiverAccountType = receiverAccountType
	this.TransactionType = transactionType
	return &this
}

// NewAccountFundingRequestWithDefaults instantiates a new AccountFundingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountFundingRequestWithDefaults() *AccountFundingRequest {
	this := AccountFundingRequest{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *AccountFundingRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *AccountFundingRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *AccountFundingRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AccountFundingRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AccountFundingRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AccountFundingRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetReceiverName returns the ReceiverName field value if set, zero value otherwise.
func (o *AccountFundingRequest) GetReceiverName() string {
	if o == nil || IsNil(o.ReceiverName) {
		var ret string
		return ret
	}
	return *o.ReceiverName
}

// GetReceiverNameOk returns a tuple with the ReceiverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingRequest) GetReceiverNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiverName) {
		return nil, false
	}
	return o.ReceiverName, true
}

// HasReceiverName returns a boolean if a field has been set.
func (o *AccountFundingRequest) HasReceiverName() bool {
	if o != nil && !IsNil(o.ReceiverName) {
		return true
	}

	return false
}

// SetReceiverName gets a reference to the given string and assigns it to the ReceiverName field.
func (o *AccountFundingRequest) SetReceiverName(v string) {
	o.ReceiverName = &v
}

// GetFundingSource returns the FundingSource field value
func (o *AccountFundingRequest) GetFundingSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value
// and a boolean to check if the value has been set.
func (o *AccountFundingRequest) GetFundingSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingSource, true
}

// SetFundingSource sets field value
func (o *AccountFundingRequest) SetFundingSource(v string) {
	o.FundingSource = v
}

// GetReceiverAccountType returns the ReceiverAccountType field value
func (o *AccountFundingRequest) GetReceiverAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverAccountType
}

// GetReceiverAccountTypeOk returns a tuple with the ReceiverAccountType field value
// and a boolean to check if the value has been set.
func (o *AccountFundingRequest) GetReceiverAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverAccountType, true
}

// SetReceiverAccountType sets field value
func (o *AccountFundingRequest) SetReceiverAccountType(v string) {
	o.ReceiverAccountType = v
}

// GetTransactionPurpose returns the TransactionPurpose field value if set, zero value otherwise.
func (o *AccountFundingRequest) GetTransactionPurpose() string {
	if o == nil || IsNil(o.TransactionPurpose) {
		var ret string
		return ret
	}
	return *o.TransactionPurpose
}

// GetTransactionPurposeOk returns a tuple with the TransactionPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFundingRequest) GetTransactionPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionPurpose) {
		return nil, false
	}
	return o.TransactionPurpose, true
}

// HasTransactionPurpose returns a boolean if a field has been set.
func (o *AccountFundingRequest) HasTransactionPurpose() bool {
	if o != nil && !IsNil(o.TransactionPurpose) {
		return true
	}

	return false
}

// SetTransactionPurpose gets a reference to the given string and assigns it to the TransactionPurpose field.
func (o *AccountFundingRequest) SetTransactionPurpose(v string) {
	o.TransactionPurpose = &v
}

// GetTransactionType returns the TransactionType field value
func (o *AccountFundingRequest) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *AccountFundingRequest) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *AccountFundingRequest) SetTransactionType(v string) {
	o.TransactionType = v
}

func (o AccountFundingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountFundingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.ReceiverName) {
		toSerialize["receiver_name"] = o.ReceiverName
	}
	toSerialize["funding_source"] = o.FundingSource
	toSerialize["receiver_account_type"] = o.ReceiverAccountType
	if !IsNil(o.TransactionPurpose) {
		toSerialize["transaction_purpose"] = o.TransactionPurpose
	}
	toSerialize["transaction_type"] = o.TransactionType
	return toSerialize, nil
}

type NullableAccountFundingRequest struct {
	value *AccountFundingRequest
	isSet bool
}

func (v NullableAccountFundingRequest) Get() *AccountFundingRequest {
	return v.value
}

func (v *NullableAccountFundingRequest) Set(val *AccountFundingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountFundingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountFundingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountFundingRequest(val *AccountFundingRequest) *NullableAccountFundingRequest {
	return &NullableAccountFundingRequest{value: val, isSet: true}
}

func (v NullableAccountFundingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountFundingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


