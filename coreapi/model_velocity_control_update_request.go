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

// checks if the VelocityControlUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VelocityControlUpdateRequest{}

// VelocityControlUpdateRequest struct for VelocityControlUpdateRequest
type VelocityControlUpdateRequest struct {
	// Indicates whether the velocity control is active.
	Active *bool `json:"active,omitempty"`
	// Maximum monetary sum that can be cleared within the time period defined by the `velocity_window` field.
	AmountLimit *float32 `json:"amount_limit,omitempty"`
	// If set to `true`, only approved transactions are subject to control.
	ApprovalsOnly *bool `json:"approvals_only,omitempty"`
	Association *SpendControlAssociation `json:"association,omitempty"`
	// Three-character ISO 4217 currency code.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// If set to `true`, cashbacks from a point of sale are subject to control.
	IncludeCashback *bool `json:"include_cashback,omitempty"`
	// If set to `true`, original credit transactions are subject to control. Your request can contain either a `money_in_transaction` object or the `include_credits` field, not both.
	IncludeCredits *bool `json:"include_credits,omitempty"`
	// If set to `true`, purchases are subject to control.
	IncludePurchases *bool `json:"include_purchases,omitempty"`
	// If set to `true`, transfers are subject to control.
	IncludeTransfers *bool `json:"include_transfers,omitempty"`
	// If set to `true`, ATM withdrawals are subject to control.
	IncludeWithdrawals *bool `json:"include_withdrawals,omitempty"`
	MerchantScope *MerchantScope `json:"merchant_scope,omitempty"`
	MoneyInTransaction *MoneyInTransaction `json:"money_in_transaction,omitempty"`
	// Description of how the velocity control restricts spending. For example, \"Max spend of $500 per day\" or \"Max spend of $5000 per month for non-exempt employees\".  Ensure that the description you provide here adequately captures the kind of restriction exerted by this velocity control, because the Marqeta platform will send this information to you in a webhook in the event that the transaction authorization attempt is declined by the velocity control.  *NOTE:* This field is very important. If your program has multiple velocity controls in place, it is not always clear which one prevented the transaction from being approved. Adding specific details to this field gives you more contextual information when debugging or monitoring declined authorization attempts.
	Name *string `json:"name,omitempty"`
	// Unique identifier of the velocity control resource.
	Token string `json:"token"`
	// Maximum number of times a card can be used within the time period defined by the `velocity_window` field.  If `velocity_window` is set to `TRANSACTION`, do not include a `usage_limit` in your request.
	UsageLimit *int32 `json:"usage_limit,omitempty"`
	// Defines the time period to which the `amount_limit` and `usage_limit` fields apply:  * *DAY*  one day; days begin at 00:00:00 UTC. * *WEEK*  one week; weeks begin Sundays at 00:00:00 UTC. * *MONTH*  one month; months begin on the first day of month at 00:00:00 UTC. * *QUARTER* - three months; quarters begin on January 1, April 1, July 1, and October 1 at 00:00:00 UTC. * *YEAR* - twelve months; years begin on January 1 at 00:00:00 UTC. * *LIFETIME*  forever; time period never expires. * *TRANSACTION*  a single transaction.  // (2023-05-03): This statement was validated by Processing, as part of a customer inquiry. *NOTE:* If set to `DAY`, `WEEK`, or `MONTH`, the velocity control takes effect retroactively from the beginning of the specified period. The amount and usage data already collected within the first period is counted toward the limits.  // (2023-05-03): Commenting this note out, as it is untrue in testing as reported by customers and confirmed by transaction engine team //_*NOTE:* Editing any of the following fields on a velocity control resets its usage and amount count to 0:  //_* `merchant_scope.mcc` //_* `merchant_scope.mid` //_* `merchant_scope.mcc_group` //_* `association.user_token` //_* `association.card_product_token`
	VelocityWindow *string `json:"velocity_window,omitempty"`
}

// NewVelocityControlUpdateRequest instantiates a new VelocityControlUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVelocityControlUpdateRequest(token string) *VelocityControlUpdateRequest {
	this := VelocityControlUpdateRequest{}
	this.Token = token
	return &this
}

// NewVelocityControlUpdateRequestWithDefaults instantiates a new VelocityControlUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelocityControlUpdateRequestWithDefaults() *VelocityControlUpdateRequest {
	this := VelocityControlUpdateRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *VelocityControlUpdateRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAmountLimit returns the AmountLimit field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetAmountLimit() float32 {
	if o == nil || IsNil(o.AmountLimit) {
		var ret float32
		return ret
	}
	return *o.AmountLimit
}

// GetAmountLimitOk returns a tuple with the AmountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetAmountLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountLimit) {
		return nil, false
	}
	return o.AmountLimit, true
}

// HasAmountLimit returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasAmountLimit() bool {
	if o != nil && !IsNil(o.AmountLimit) {
		return true
	}

	return false
}

// SetAmountLimit gets a reference to the given float32 and assigns it to the AmountLimit field.
func (o *VelocityControlUpdateRequest) SetAmountLimit(v float32) {
	o.AmountLimit = &v
}

// GetApprovalsOnly returns the ApprovalsOnly field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetApprovalsOnly() bool {
	if o == nil || IsNil(o.ApprovalsOnly) {
		var ret bool
		return ret
	}
	return *o.ApprovalsOnly
}

// GetApprovalsOnlyOk returns a tuple with the ApprovalsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetApprovalsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ApprovalsOnly) {
		return nil, false
	}
	return o.ApprovalsOnly, true
}

// HasApprovalsOnly returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasApprovalsOnly() bool {
	if o != nil && !IsNil(o.ApprovalsOnly) {
		return true
	}

	return false
}

// SetApprovalsOnly gets a reference to the given bool and assigns it to the ApprovalsOnly field.
func (o *VelocityControlUpdateRequest) SetApprovalsOnly(v bool) {
	o.ApprovalsOnly = &v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetAssociation() SpendControlAssociation {
	if o == nil || IsNil(o.Association) {
		var ret SpendControlAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetAssociationOk() (*SpendControlAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given SpendControlAssociation and assigns it to the Association field.
func (o *VelocityControlUpdateRequest) SetAssociation(v SpendControlAssociation) {
	o.Association = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *VelocityControlUpdateRequest) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetIncludeCashback returns the IncludeCashback field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetIncludeCashback() bool {
	if o == nil || IsNil(o.IncludeCashback) {
		var ret bool
		return ret
	}
	return *o.IncludeCashback
}

// GetIncludeCashbackOk returns a tuple with the IncludeCashback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetIncludeCashbackOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCashback) {
		return nil, false
	}
	return o.IncludeCashback, true
}

// HasIncludeCashback returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasIncludeCashback() bool {
	if o != nil && !IsNil(o.IncludeCashback) {
		return true
	}

	return false
}

// SetIncludeCashback gets a reference to the given bool and assigns it to the IncludeCashback field.
func (o *VelocityControlUpdateRequest) SetIncludeCashback(v bool) {
	o.IncludeCashback = &v
}

// GetIncludeCredits returns the IncludeCredits field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetIncludeCredits() bool {
	if o == nil || IsNil(o.IncludeCredits) {
		var ret bool
		return ret
	}
	return *o.IncludeCredits
}

// GetIncludeCreditsOk returns a tuple with the IncludeCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetIncludeCreditsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCredits) {
		return nil, false
	}
	return o.IncludeCredits, true
}

// HasIncludeCredits returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasIncludeCredits() bool {
	if o != nil && !IsNil(o.IncludeCredits) {
		return true
	}

	return false
}

// SetIncludeCredits gets a reference to the given bool and assigns it to the IncludeCredits field.
func (o *VelocityControlUpdateRequest) SetIncludeCredits(v bool) {
	o.IncludeCredits = &v
}

// GetIncludePurchases returns the IncludePurchases field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetIncludePurchases() bool {
	if o == nil || IsNil(o.IncludePurchases) {
		var ret bool
		return ret
	}
	return *o.IncludePurchases
}

// GetIncludePurchasesOk returns a tuple with the IncludePurchases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetIncludePurchasesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludePurchases) {
		return nil, false
	}
	return o.IncludePurchases, true
}

// HasIncludePurchases returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasIncludePurchases() bool {
	if o != nil && !IsNil(o.IncludePurchases) {
		return true
	}

	return false
}

// SetIncludePurchases gets a reference to the given bool and assigns it to the IncludePurchases field.
func (o *VelocityControlUpdateRequest) SetIncludePurchases(v bool) {
	o.IncludePurchases = &v
}

// GetIncludeTransfers returns the IncludeTransfers field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetIncludeTransfers() bool {
	if o == nil || IsNil(o.IncludeTransfers) {
		var ret bool
		return ret
	}
	return *o.IncludeTransfers
}

// GetIncludeTransfersOk returns a tuple with the IncludeTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetIncludeTransfersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTransfers) {
		return nil, false
	}
	return o.IncludeTransfers, true
}

// HasIncludeTransfers returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasIncludeTransfers() bool {
	if o != nil && !IsNil(o.IncludeTransfers) {
		return true
	}

	return false
}

// SetIncludeTransfers gets a reference to the given bool and assigns it to the IncludeTransfers field.
func (o *VelocityControlUpdateRequest) SetIncludeTransfers(v bool) {
	o.IncludeTransfers = &v
}

// GetIncludeWithdrawals returns the IncludeWithdrawals field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetIncludeWithdrawals() bool {
	if o == nil || IsNil(o.IncludeWithdrawals) {
		var ret bool
		return ret
	}
	return *o.IncludeWithdrawals
}

// GetIncludeWithdrawalsOk returns a tuple with the IncludeWithdrawals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetIncludeWithdrawalsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeWithdrawals) {
		return nil, false
	}
	return o.IncludeWithdrawals, true
}

// HasIncludeWithdrawals returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasIncludeWithdrawals() bool {
	if o != nil && !IsNil(o.IncludeWithdrawals) {
		return true
	}

	return false
}

// SetIncludeWithdrawals gets a reference to the given bool and assigns it to the IncludeWithdrawals field.
func (o *VelocityControlUpdateRequest) SetIncludeWithdrawals(v bool) {
	o.IncludeWithdrawals = &v
}

// GetMerchantScope returns the MerchantScope field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetMerchantScope() MerchantScope {
	if o == nil || IsNil(o.MerchantScope) {
		var ret MerchantScope
		return ret
	}
	return *o.MerchantScope
}

// GetMerchantScopeOk returns a tuple with the MerchantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetMerchantScopeOk() (*MerchantScope, bool) {
	if o == nil || IsNil(o.MerchantScope) {
		return nil, false
	}
	return o.MerchantScope, true
}

// HasMerchantScope returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasMerchantScope() bool {
	if o != nil && !IsNil(o.MerchantScope) {
		return true
	}

	return false
}

// SetMerchantScope gets a reference to the given MerchantScope and assigns it to the MerchantScope field.
func (o *VelocityControlUpdateRequest) SetMerchantScope(v MerchantScope) {
	o.MerchantScope = &v
}

// GetMoneyInTransaction returns the MoneyInTransaction field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetMoneyInTransaction() MoneyInTransaction {
	if o == nil || IsNil(o.MoneyInTransaction) {
		var ret MoneyInTransaction
		return ret
	}
	return *o.MoneyInTransaction
}

// GetMoneyInTransactionOk returns a tuple with the MoneyInTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetMoneyInTransactionOk() (*MoneyInTransaction, bool) {
	if o == nil || IsNil(o.MoneyInTransaction) {
		return nil, false
	}
	return o.MoneyInTransaction, true
}

// HasMoneyInTransaction returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasMoneyInTransaction() bool {
	if o != nil && !IsNil(o.MoneyInTransaction) {
		return true
	}

	return false
}

// SetMoneyInTransaction gets a reference to the given MoneyInTransaction and assigns it to the MoneyInTransaction field.
func (o *VelocityControlUpdateRequest) SetMoneyInTransaction(v MoneyInTransaction) {
	o.MoneyInTransaction = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VelocityControlUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value
func (o *VelocityControlUpdateRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *VelocityControlUpdateRequest) SetToken(v string) {
	o.Token = v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetUsageLimit() int32 {
	if o == nil || IsNil(o.UsageLimit) {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetUsageLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UsageLimit) {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasUsageLimit() bool {
	if o != nil && !IsNil(o.UsageLimit) {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *VelocityControlUpdateRequest) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

// GetVelocityWindow returns the VelocityWindow field value if set, zero value otherwise.
func (o *VelocityControlUpdateRequest) GetVelocityWindow() string {
	if o == nil || IsNil(o.VelocityWindow) {
		var ret string
		return ret
	}
	return *o.VelocityWindow
}

// GetVelocityWindowOk returns a tuple with the VelocityWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelocityControlUpdateRequest) GetVelocityWindowOk() (*string, bool) {
	if o == nil || IsNil(o.VelocityWindow) {
		return nil, false
	}
	return o.VelocityWindow, true
}

// HasVelocityWindow returns a boolean if a field has been set.
func (o *VelocityControlUpdateRequest) HasVelocityWindow() bool {
	if o != nil && !IsNil(o.VelocityWindow) {
		return true
	}

	return false
}

// SetVelocityWindow gets a reference to the given string and assigns it to the VelocityWindow field.
func (o *VelocityControlUpdateRequest) SetVelocityWindow(v string) {
	o.VelocityWindow = &v
}

func (o VelocityControlUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VelocityControlUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.AmountLimit) {
		toSerialize["amount_limit"] = o.AmountLimit
	}
	if !IsNil(o.ApprovalsOnly) {
		toSerialize["approvals_only"] = o.ApprovalsOnly
	}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.IncludeCashback) {
		toSerialize["include_cashback"] = o.IncludeCashback
	}
	if !IsNil(o.IncludeCredits) {
		toSerialize["include_credits"] = o.IncludeCredits
	}
	if !IsNil(o.IncludePurchases) {
		toSerialize["include_purchases"] = o.IncludePurchases
	}
	if !IsNil(o.IncludeTransfers) {
		toSerialize["include_transfers"] = o.IncludeTransfers
	}
	if !IsNil(o.IncludeWithdrawals) {
		toSerialize["include_withdrawals"] = o.IncludeWithdrawals
	}
	if !IsNil(o.MerchantScope) {
		toSerialize["merchant_scope"] = o.MerchantScope
	}
	if !IsNil(o.MoneyInTransaction) {
		toSerialize["money_in_transaction"] = o.MoneyInTransaction
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.UsageLimit) {
		toSerialize["usage_limit"] = o.UsageLimit
	}
	if !IsNil(o.VelocityWindow) {
		toSerialize["velocity_window"] = o.VelocityWindow
	}
	return toSerialize, nil
}

type NullableVelocityControlUpdateRequest struct {
	value *VelocityControlUpdateRequest
	isSet bool
}

func (v NullableVelocityControlUpdateRequest) Get() *VelocityControlUpdateRequest {
	return v.value
}

func (v *NullableVelocityControlUpdateRequest) Set(val *VelocityControlUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityControlUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityControlUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityControlUpdateRequest(val *VelocityControlUpdateRequest) *NullableVelocityControlUpdateRequest {
	return &NullableVelocityControlUpdateRequest{value: val, isSet: true}
}

func (v NullableVelocityControlUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityControlUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


