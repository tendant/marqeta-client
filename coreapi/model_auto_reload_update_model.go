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

// checks if the AutoReloadUpdateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoReloadUpdateModel{}

// AutoReloadUpdateModel struct for AutoReloadUpdateModel
type AutoReloadUpdateModel struct {
	// Specifies whether the auto reload is active.  Only one auto reload per level, per object, can be active.
	Active *bool `json:"active,omitempty"`
	Association *AutoReloadAssociation `json:"association,omitempty"`
	// Three-digit link:https://www.iso.org/iso-4217-currency-codes.html[ISO 4217 currency code, window=\"_blank\"].
	CurrencyCode *string `json:"currency_code,omitempty"`
	// Unique identifier of the funding source address to use for this auto reload.  If your funding source is an ACH account, then a `funding_source_address_token` is not required. If your funding source is a payment card, you must have at least one funding source address in order to create a GPA order.
	FundingSourceAddressToken *string `json:"funding_source_address_token,omitempty"`
	// Unique identifier of the funding source to use for this auto reload.
	FundingSourceToken *string `json:"funding_source_token,omitempty"`
	OrderScope *OrderScope `json:"order_scope,omitempty"`
	// The token in the path parameter takes precedence over the `token` body field.
	Token *string `json:"token,omitempty"`
}

// NewAutoReloadUpdateModel instantiates a new AutoReloadUpdateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoReloadUpdateModel() *AutoReloadUpdateModel {
	this := AutoReloadUpdateModel{}
	var active bool = true
	this.Active = &active
	return &this
}

// NewAutoReloadUpdateModelWithDefaults instantiates a new AutoReloadUpdateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoReloadUpdateModelWithDefaults() *AutoReloadUpdateModel {
	this := AutoReloadUpdateModel{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AutoReloadUpdateModel) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadUpdateModel) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AutoReloadUpdateModel) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AutoReloadUpdateModel) SetActive(v bool) {
	o.Active = &v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *AutoReloadUpdateModel) GetAssociation() AutoReloadAssociation {
	if o == nil || IsNil(o.Association) {
		var ret AutoReloadAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadUpdateModel) GetAssociationOk() (*AutoReloadAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *AutoReloadUpdateModel) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given AutoReloadAssociation and assigns it to the Association field.
func (o *AutoReloadUpdateModel) SetAssociation(v AutoReloadAssociation) {
	o.Association = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AutoReloadUpdateModel) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadUpdateModel) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AutoReloadUpdateModel) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AutoReloadUpdateModel) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetFundingSourceAddressToken returns the FundingSourceAddressToken field value if set, zero value otherwise.
func (o *AutoReloadUpdateModel) GetFundingSourceAddressToken() string {
	if o == nil || IsNil(o.FundingSourceAddressToken) {
		var ret string
		return ret
	}
	return *o.FundingSourceAddressToken
}

// GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadUpdateModel) GetFundingSourceAddressTokenOk() (*string, bool) {
	if o == nil || IsNil(o.FundingSourceAddressToken) {
		return nil, false
	}
	return o.FundingSourceAddressToken, true
}

// HasFundingSourceAddressToken returns a boolean if a field has been set.
func (o *AutoReloadUpdateModel) HasFundingSourceAddressToken() bool {
	if o != nil && !IsNil(o.FundingSourceAddressToken) {
		return true
	}

	return false
}

// SetFundingSourceAddressToken gets a reference to the given string and assigns it to the FundingSourceAddressToken field.
func (o *AutoReloadUpdateModel) SetFundingSourceAddressToken(v string) {
	o.FundingSourceAddressToken = &v
}

// GetFundingSourceToken returns the FundingSourceToken field value if set, zero value otherwise.
func (o *AutoReloadUpdateModel) GetFundingSourceToken() string {
	if o == nil || IsNil(o.FundingSourceToken) {
		var ret string
		return ret
	}
	return *o.FundingSourceToken
}

// GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadUpdateModel) GetFundingSourceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.FundingSourceToken) {
		return nil, false
	}
	return o.FundingSourceToken, true
}

// HasFundingSourceToken returns a boolean if a field has been set.
func (o *AutoReloadUpdateModel) HasFundingSourceToken() bool {
	if o != nil && !IsNil(o.FundingSourceToken) {
		return true
	}

	return false
}

// SetFundingSourceToken gets a reference to the given string and assigns it to the FundingSourceToken field.
func (o *AutoReloadUpdateModel) SetFundingSourceToken(v string) {
	o.FundingSourceToken = &v
}

// GetOrderScope returns the OrderScope field value if set, zero value otherwise.
func (o *AutoReloadUpdateModel) GetOrderScope() OrderScope {
	if o == nil || IsNil(o.OrderScope) {
		var ret OrderScope
		return ret
	}
	return *o.OrderScope
}

// GetOrderScopeOk returns a tuple with the OrderScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadUpdateModel) GetOrderScopeOk() (*OrderScope, bool) {
	if o == nil || IsNil(o.OrderScope) {
		return nil, false
	}
	return o.OrderScope, true
}

// HasOrderScope returns a boolean if a field has been set.
func (o *AutoReloadUpdateModel) HasOrderScope() bool {
	if o != nil && !IsNil(o.OrderScope) {
		return true
	}

	return false
}

// SetOrderScope gets a reference to the given OrderScope and assigns it to the OrderScope field.
func (o *AutoReloadUpdateModel) SetOrderScope(v OrderScope) {
	o.OrderScope = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AutoReloadUpdateModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoReloadUpdateModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AutoReloadUpdateModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AutoReloadUpdateModel) SetToken(v string) {
	o.Token = &v
}

func (o AutoReloadUpdateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoReloadUpdateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.FundingSourceAddressToken) {
		toSerialize["funding_source_address_token"] = o.FundingSourceAddressToken
	}
	if !IsNil(o.FundingSourceToken) {
		toSerialize["funding_source_token"] = o.FundingSourceToken
	}
	if !IsNil(o.OrderScope) {
		toSerialize["order_scope"] = o.OrderScope
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableAutoReloadUpdateModel struct {
	value *AutoReloadUpdateModel
	isSet bool
}

func (v NullableAutoReloadUpdateModel) Get() *AutoReloadUpdateModel {
	return v.value
}

func (v *NullableAutoReloadUpdateModel) Set(val *AutoReloadUpdateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoReloadUpdateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoReloadUpdateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoReloadUpdateModel(val *AutoReloadUpdateModel) *NullableAutoReloadUpdateModel {
	return &NullableAutoReloadUpdateModel{value: val, isSet: true}
}

func (v NullableAutoReloadUpdateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoReloadUpdateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


