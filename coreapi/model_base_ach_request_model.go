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

// checks if the BaseAchRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseAchRequestModel{}

// BaseAchRequestModel struct for BaseAchRequestModel
type BaseAchRequestModel struct {
	// ACH account number.
	AccountNumber string `json:"account_number"`
	// Type of account.
	AccountType string `json:"account_type"`
	// Name of the bank holding the account.
	BankName *string `json:"bank_name,omitempty"`
	// If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source.
	IsDefaultAccount *bool `json:"is_default_account,omitempty"`
	// Name on the ACH account.
	NameOnAccount string `json:"name_on_account"`
	// Routing number for the ACH account.
	RoutingNumber string `json:"routing_number"`
	// Unique identifier of the funding source. If you do not include a token, the system will generate one automatically. This token is necessary for use in other calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
	// Free-form text field for holding notes about verification. This field is returned only if `verification_override = true`.
	VerificationNotes *string `json:"verification_notes,omitempty"`
	// Allows the ACH funding source to be used regardless of its verification status.
	VerificationOverride *bool `json:"verification_override,omitempty"`
}

// NewBaseAchRequestModel instantiates a new BaseAchRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseAchRequestModel(accountNumber string, accountType string, nameOnAccount string, routingNumber string) *BaseAchRequestModel {
	this := BaseAchRequestModel{}
	this.AccountNumber = accountNumber
	this.AccountType = accountType
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	this.NameOnAccount = nameOnAccount
	this.RoutingNumber = routingNumber
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// NewBaseAchRequestModelWithDefaults instantiates a new BaseAchRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseAchRequestModelWithDefaults() *BaseAchRequestModel {
	this := BaseAchRequestModel{}
	var isDefaultAccount bool = false
	this.IsDefaultAccount = &isDefaultAccount
	var verificationOverride bool = false
	this.VerificationOverride = &verificationOverride
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *BaseAchRequestModel) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *BaseAchRequestModel) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAccountType returns the AccountType field value
func (o *BaseAchRequestModel) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *BaseAchRequestModel) SetAccountType(v string) {
	o.AccountType = v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *BaseAchRequestModel) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *BaseAchRequestModel) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *BaseAchRequestModel) SetBankName(v string) {
	o.BankName = &v
}

// GetIsDefaultAccount returns the IsDefaultAccount field value if set, zero value otherwise.
func (o *BaseAchRequestModel) GetIsDefaultAccount() bool {
	if o == nil || IsNil(o.IsDefaultAccount) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccount
}

// GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetIsDefaultAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultAccount) {
		return nil, false
	}
	return o.IsDefaultAccount, true
}

// HasIsDefaultAccount returns a boolean if a field has been set.
func (o *BaseAchRequestModel) HasIsDefaultAccount() bool {
	if o != nil && !IsNil(o.IsDefaultAccount) {
		return true
	}

	return false
}

// SetIsDefaultAccount gets a reference to the given bool and assigns it to the IsDefaultAccount field.
func (o *BaseAchRequestModel) SetIsDefaultAccount(v bool) {
	o.IsDefaultAccount = &v
}

// GetNameOnAccount returns the NameOnAccount field value
func (o *BaseAchRequestModel) GetNameOnAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameOnAccount
}

// GetNameOnAccountOk returns a tuple with the NameOnAccount field value
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetNameOnAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameOnAccount, true
}

// SetNameOnAccount sets field value
func (o *BaseAchRequestModel) SetNameOnAccount(v string) {
	o.NameOnAccount = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *BaseAchRequestModel) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *BaseAchRequestModel) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *BaseAchRequestModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *BaseAchRequestModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *BaseAchRequestModel) SetToken(v string) {
	o.Token = &v
}

// GetVerificationNotes returns the VerificationNotes field value if set, zero value otherwise.
func (o *BaseAchRequestModel) GetVerificationNotes() string {
	if o == nil || IsNil(o.VerificationNotes) {
		var ret string
		return ret
	}
	return *o.VerificationNotes
}

// GetVerificationNotesOk returns a tuple with the VerificationNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetVerificationNotesOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationNotes) {
		return nil, false
	}
	return o.VerificationNotes, true
}

// HasVerificationNotes returns a boolean if a field has been set.
func (o *BaseAchRequestModel) HasVerificationNotes() bool {
	if o != nil && !IsNil(o.VerificationNotes) {
		return true
	}

	return false
}

// SetVerificationNotes gets a reference to the given string and assigns it to the VerificationNotes field.
func (o *BaseAchRequestModel) SetVerificationNotes(v string) {
	o.VerificationNotes = &v
}

// GetVerificationOverride returns the VerificationOverride field value if set, zero value otherwise.
func (o *BaseAchRequestModel) GetVerificationOverride() bool {
	if o == nil || IsNil(o.VerificationOverride) {
		var ret bool
		return ret
	}
	return *o.VerificationOverride
}

// GetVerificationOverrideOk returns a tuple with the VerificationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAchRequestModel) GetVerificationOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.VerificationOverride) {
		return nil, false
	}
	return o.VerificationOverride, true
}

// HasVerificationOverride returns a boolean if a field has been set.
func (o *BaseAchRequestModel) HasVerificationOverride() bool {
	if o != nil && !IsNil(o.VerificationOverride) {
		return true
	}

	return false
}

// SetVerificationOverride gets a reference to the given bool and assigns it to the VerificationOverride field.
func (o *BaseAchRequestModel) SetVerificationOverride(v bool) {
	o.VerificationOverride = &v
}

func (o BaseAchRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseAchRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["account_type"] = o.AccountType
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.IsDefaultAccount) {
		toSerialize["is_default_account"] = o.IsDefaultAccount
	}
	toSerialize["name_on_account"] = o.NameOnAccount
	toSerialize["routing_number"] = o.RoutingNumber
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.VerificationNotes) {
		toSerialize["verification_notes"] = o.VerificationNotes
	}
	if !IsNil(o.VerificationOverride) {
		toSerialize["verification_override"] = o.VerificationOverride
	}
	return toSerialize, nil
}

type NullableBaseAchRequestModel struct {
	value *BaseAchRequestModel
	isSet bool
}

func (v NullableBaseAchRequestModel) Get() *BaseAchRequestModel {
	return v.value
}

func (v *NullableBaseAchRequestModel) Set(val *BaseAchRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseAchRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseAchRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseAchRequestModel(val *BaseAchRequestModel) *NullableBaseAchRequestModel {
	return &NullableBaseAchRequestModel{value: val, isSet: true}
}

func (v NullableBaseAchRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseAchRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

