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

// checks if the PolicyDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyDocumentResponse{}

// PolicyDocumentResponse Contains information on a document policy.
type PolicyDocumentResponse struct {
	AccountStatement *PolicyDocumentTemplateResponse `json:"account_statement,omitempty"`
	BenefitsDisclosurePremium *PolicyDocumentAssetResponse `json:"benefits_disclosure_premium,omitempty"`
	BenefitsDisclosureTraditional *PolicyDocumentAssetResponse `json:"benefits_disclosure_traditional,omitempty"`
	CardMemberAgreement *PolicyDocumentAssetResponse `json:"card_member_agreement,omitempty"`
	// Date and time when the document policy was created on Marqeta's credit platform, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	EDisclosure *PolicyDocumentAssetResponse `json:"e_disclosure,omitempty"`
	// Name of the document policy.
	Name *string `json:"name,omitempty"`
	NoaaMultipleReasonWithDoddFrank *PolicyDocumentTemplateResponse `json:"noaa_multiple_reason_with_dodd_frank,omitempty"`
	NoaaSingleReason *PolicyDocumentTemplateResponse `json:"noaa_single_reason,omitempty"`
	NoaaSingleReasonWithDoddFrank *PolicyDocumentTemplateResponse `json:"noaa_single_reason_with_dodd_frank,omitempty"`
	PrivacyPolicy *PolicyDocumentAssetResponse `json:"privacy_policy,omitempty"`
	RewardsDisclosure *PolicyDocumentAssetAndTemplateResponse `json:"rewards_disclosure,omitempty"`
	SummaryOfCreditTerms *PolicyDocumentAssetAndTemplateResponse `json:"summary_of_credit_terms,omitempty"`
	TermsSchedule *PolicyDocumentTemplateResponse `json:"terms_schedule,omitempty"`
	// Unique identifier of the document policy.
	Token *string `json:"token,omitempty"`
	// Date and time when the document policy was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}

// NewPolicyDocumentResponse instantiates a new PolicyDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyDocumentResponse() *PolicyDocumentResponse {
	this := PolicyDocumentResponse{}
	return &this
}

// NewPolicyDocumentResponseWithDefaults instantiates a new PolicyDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyDocumentResponseWithDefaults() *PolicyDocumentResponse {
	this := PolicyDocumentResponse{}
	return &this
}

// GetAccountStatement returns the AccountStatement field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetAccountStatement() PolicyDocumentTemplateResponse {
	if o == nil || IsNil(o.AccountStatement) {
		var ret PolicyDocumentTemplateResponse
		return ret
	}
	return *o.AccountStatement
}

// GetAccountStatementOk returns a tuple with the AccountStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetAccountStatementOk() (*PolicyDocumentTemplateResponse, bool) {
	if o == nil || IsNil(o.AccountStatement) {
		return nil, false
	}
	return o.AccountStatement, true
}

// HasAccountStatement returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasAccountStatement() bool {
	if o != nil && !IsNil(o.AccountStatement) {
		return true
	}

	return false
}

// SetAccountStatement gets a reference to the given PolicyDocumentTemplateResponse and assigns it to the AccountStatement field.
func (o *PolicyDocumentResponse) SetAccountStatement(v PolicyDocumentTemplateResponse) {
	o.AccountStatement = &v
}

// GetBenefitsDisclosurePremium returns the BenefitsDisclosurePremium field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetBenefitsDisclosurePremium() PolicyDocumentAssetResponse {
	if o == nil || IsNil(o.BenefitsDisclosurePremium) {
		var ret PolicyDocumentAssetResponse
		return ret
	}
	return *o.BenefitsDisclosurePremium
}

// GetBenefitsDisclosurePremiumOk returns a tuple with the BenefitsDisclosurePremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetBenefitsDisclosurePremiumOk() (*PolicyDocumentAssetResponse, bool) {
	if o == nil || IsNil(o.BenefitsDisclosurePremium) {
		return nil, false
	}
	return o.BenefitsDisclosurePremium, true
}

// HasBenefitsDisclosurePremium returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasBenefitsDisclosurePremium() bool {
	if o != nil && !IsNil(o.BenefitsDisclosurePremium) {
		return true
	}

	return false
}

// SetBenefitsDisclosurePremium gets a reference to the given PolicyDocumentAssetResponse and assigns it to the BenefitsDisclosurePremium field.
func (o *PolicyDocumentResponse) SetBenefitsDisclosurePremium(v PolicyDocumentAssetResponse) {
	o.BenefitsDisclosurePremium = &v
}

// GetBenefitsDisclosureTraditional returns the BenefitsDisclosureTraditional field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetBenefitsDisclosureTraditional() PolicyDocumentAssetResponse {
	if o == nil || IsNil(o.BenefitsDisclosureTraditional) {
		var ret PolicyDocumentAssetResponse
		return ret
	}
	return *o.BenefitsDisclosureTraditional
}

// GetBenefitsDisclosureTraditionalOk returns a tuple with the BenefitsDisclosureTraditional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetBenefitsDisclosureTraditionalOk() (*PolicyDocumentAssetResponse, bool) {
	if o == nil || IsNil(o.BenefitsDisclosureTraditional) {
		return nil, false
	}
	return o.BenefitsDisclosureTraditional, true
}

// HasBenefitsDisclosureTraditional returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasBenefitsDisclosureTraditional() bool {
	if o != nil && !IsNil(o.BenefitsDisclosureTraditional) {
		return true
	}

	return false
}

// SetBenefitsDisclosureTraditional gets a reference to the given PolicyDocumentAssetResponse and assigns it to the BenefitsDisclosureTraditional field.
func (o *PolicyDocumentResponse) SetBenefitsDisclosureTraditional(v PolicyDocumentAssetResponse) {
	o.BenefitsDisclosureTraditional = &v
}

// GetCardMemberAgreement returns the CardMemberAgreement field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetCardMemberAgreement() PolicyDocumentAssetResponse {
	if o == nil || IsNil(o.CardMemberAgreement) {
		var ret PolicyDocumentAssetResponse
		return ret
	}
	return *o.CardMemberAgreement
}

// GetCardMemberAgreementOk returns a tuple with the CardMemberAgreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetCardMemberAgreementOk() (*PolicyDocumentAssetResponse, bool) {
	if o == nil || IsNil(o.CardMemberAgreement) {
		return nil, false
	}
	return o.CardMemberAgreement, true
}

// HasCardMemberAgreement returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasCardMemberAgreement() bool {
	if o != nil && !IsNil(o.CardMemberAgreement) {
		return true
	}

	return false
}

// SetCardMemberAgreement gets a reference to the given PolicyDocumentAssetResponse and assigns it to the CardMemberAgreement field.
func (o *PolicyDocumentResponse) SetCardMemberAgreement(v PolicyDocumentAssetResponse) {
	o.CardMemberAgreement = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *PolicyDocumentResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetEDisclosure returns the EDisclosure field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetEDisclosure() PolicyDocumentAssetResponse {
	if o == nil || IsNil(o.EDisclosure) {
		var ret PolicyDocumentAssetResponse
		return ret
	}
	return *o.EDisclosure
}

// GetEDisclosureOk returns a tuple with the EDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetEDisclosureOk() (*PolicyDocumentAssetResponse, bool) {
	if o == nil || IsNil(o.EDisclosure) {
		return nil, false
	}
	return o.EDisclosure, true
}

// HasEDisclosure returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasEDisclosure() bool {
	if o != nil && !IsNil(o.EDisclosure) {
		return true
	}

	return false
}

// SetEDisclosure gets a reference to the given PolicyDocumentAssetResponse and assigns it to the EDisclosure field.
func (o *PolicyDocumentResponse) SetEDisclosure(v PolicyDocumentAssetResponse) {
	o.EDisclosure = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyDocumentResponse) SetName(v string) {
	o.Name = &v
}

// GetNoaaMultipleReasonWithDoddFrank returns the NoaaMultipleReasonWithDoddFrank field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetNoaaMultipleReasonWithDoddFrank() PolicyDocumentTemplateResponse {
	if o == nil || IsNil(o.NoaaMultipleReasonWithDoddFrank) {
		var ret PolicyDocumentTemplateResponse
		return ret
	}
	return *o.NoaaMultipleReasonWithDoddFrank
}

// GetNoaaMultipleReasonWithDoddFrankOk returns a tuple with the NoaaMultipleReasonWithDoddFrank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetNoaaMultipleReasonWithDoddFrankOk() (*PolicyDocumentTemplateResponse, bool) {
	if o == nil || IsNil(o.NoaaMultipleReasonWithDoddFrank) {
		return nil, false
	}
	return o.NoaaMultipleReasonWithDoddFrank, true
}

// HasNoaaMultipleReasonWithDoddFrank returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasNoaaMultipleReasonWithDoddFrank() bool {
	if o != nil && !IsNil(o.NoaaMultipleReasonWithDoddFrank) {
		return true
	}

	return false
}

// SetNoaaMultipleReasonWithDoddFrank gets a reference to the given PolicyDocumentTemplateResponse and assigns it to the NoaaMultipleReasonWithDoddFrank field.
func (o *PolicyDocumentResponse) SetNoaaMultipleReasonWithDoddFrank(v PolicyDocumentTemplateResponse) {
	o.NoaaMultipleReasonWithDoddFrank = &v
}

// GetNoaaSingleReason returns the NoaaSingleReason field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetNoaaSingleReason() PolicyDocumentTemplateResponse {
	if o == nil || IsNil(o.NoaaSingleReason) {
		var ret PolicyDocumentTemplateResponse
		return ret
	}
	return *o.NoaaSingleReason
}

// GetNoaaSingleReasonOk returns a tuple with the NoaaSingleReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetNoaaSingleReasonOk() (*PolicyDocumentTemplateResponse, bool) {
	if o == nil || IsNil(o.NoaaSingleReason) {
		return nil, false
	}
	return o.NoaaSingleReason, true
}

// HasNoaaSingleReason returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasNoaaSingleReason() bool {
	if o != nil && !IsNil(o.NoaaSingleReason) {
		return true
	}

	return false
}

// SetNoaaSingleReason gets a reference to the given PolicyDocumentTemplateResponse and assigns it to the NoaaSingleReason field.
func (o *PolicyDocumentResponse) SetNoaaSingleReason(v PolicyDocumentTemplateResponse) {
	o.NoaaSingleReason = &v
}

// GetNoaaSingleReasonWithDoddFrank returns the NoaaSingleReasonWithDoddFrank field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetNoaaSingleReasonWithDoddFrank() PolicyDocumentTemplateResponse {
	if o == nil || IsNil(o.NoaaSingleReasonWithDoddFrank) {
		var ret PolicyDocumentTemplateResponse
		return ret
	}
	return *o.NoaaSingleReasonWithDoddFrank
}

// GetNoaaSingleReasonWithDoddFrankOk returns a tuple with the NoaaSingleReasonWithDoddFrank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetNoaaSingleReasonWithDoddFrankOk() (*PolicyDocumentTemplateResponse, bool) {
	if o == nil || IsNil(o.NoaaSingleReasonWithDoddFrank) {
		return nil, false
	}
	return o.NoaaSingleReasonWithDoddFrank, true
}

// HasNoaaSingleReasonWithDoddFrank returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasNoaaSingleReasonWithDoddFrank() bool {
	if o != nil && !IsNil(o.NoaaSingleReasonWithDoddFrank) {
		return true
	}

	return false
}

// SetNoaaSingleReasonWithDoddFrank gets a reference to the given PolicyDocumentTemplateResponse and assigns it to the NoaaSingleReasonWithDoddFrank field.
func (o *PolicyDocumentResponse) SetNoaaSingleReasonWithDoddFrank(v PolicyDocumentTemplateResponse) {
	o.NoaaSingleReasonWithDoddFrank = &v
}

// GetPrivacyPolicy returns the PrivacyPolicy field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetPrivacyPolicy() PolicyDocumentAssetResponse {
	if o == nil || IsNil(o.PrivacyPolicy) {
		var ret PolicyDocumentAssetResponse
		return ret
	}
	return *o.PrivacyPolicy
}

// GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetPrivacyPolicyOk() (*PolicyDocumentAssetResponse, bool) {
	if o == nil || IsNil(o.PrivacyPolicy) {
		return nil, false
	}
	return o.PrivacyPolicy, true
}

// HasPrivacyPolicy returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasPrivacyPolicy() bool {
	if o != nil && !IsNil(o.PrivacyPolicy) {
		return true
	}

	return false
}

// SetPrivacyPolicy gets a reference to the given PolicyDocumentAssetResponse and assigns it to the PrivacyPolicy field.
func (o *PolicyDocumentResponse) SetPrivacyPolicy(v PolicyDocumentAssetResponse) {
	o.PrivacyPolicy = &v
}

// GetRewardsDisclosure returns the RewardsDisclosure field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetRewardsDisclosure() PolicyDocumentAssetAndTemplateResponse {
	if o == nil || IsNil(o.RewardsDisclosure) {
		var ret PolicyDocumentAssetAndTemplateResponse
		return ret
	}
	return *o.RewardsDisclosure
}

// GetRewardsDisclosureOk returns a tuple with the RewardsDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetRewardsDisclosureOk() (*PolicyDocumentAssetAndTemplateResponse, bool) {
	if o == nil || IsNil(o.RewardsDisclosure) {
		return nil, false
	}
	return o.RewardsDisclosure, true
}

// HasRewardsDisclosure returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasRewardsDisclosure() bool {
	if o != nil && !IsNil(o.RewardsDisclosure) {
		return true
	}

	return false
}

// SetRewardsDisclosure gets a reference to the given PolicyDocumentAssetAndTemplateResponse and assigns it to the RewardsDisclosure field.
func (o *PolicyDocumentResponse) SetRewardsDisclosure(v PolicyDocumentAssetAndTemplateResponse) {
	o.RewardsDisclosure = &v
}

// GetSummaryOfCreditTerms returns the SummaryOfCreditTerms field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetSummaryOfCreditTerms() PolicyDocumentAssetAndTemplateResponse {
	if o == nil || IsNil(o.SummaryOfCreditTerms) {
		var ret PolicyDocumentAssetAndTemplateResponse
		return ret
	}
	return *o.SummaryOfCreditTerms
}

// GetSummaryOfCreditTermsOk returns a tuple with the SummaryOfCreditTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetSummaryOfCreditTermsOk() (*PolicyDocumentAssetAndTemplateResponse, bool) {
	if o == nil || IsNil(o.SummaryOfCreditTerms) {
		return nil, false
	}
	return o.SummaryOfCreditTerms, true
}

// HasSummaryOfCreditTerms returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasSummaryOfCreditTerms() bool {
	if o != nil && !IsNil(o.SummaryOfCreditTerms) {
		return true
	}

	return false
}

// SetSummaryOfCreditTerms gets a reference to the given PolicyDocumentAssetAndTemplateResponse and assigns it to the SummaryOfCreditTerms field.
func (o *PolicyDocumentResponse) SetSummaryOfCreditTerms(v PolicyDocumentAssetAndTemplateResponse) {
	o.SummaryOfCreditTerms = &v
}

// GetTermsSchedule returns the TermsSchedule field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetTermsSchedule() PolicyDocumentTemplateResponse {
	if o == nil || IsNil(o.TermsSchedule) {
		var ret PolicyDocumentTemplateResponse
		return ret
	}
	return *o.TermsSchedule
}

// GetTermsScheduleOk returns a tuple with the TermsSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetTermsScheduleOk() (*PolicyDocumentTemplateResponse, bool) {
	if o == nil || IsNil(o.TermsSchedule) {
		return nil, false
	}
	return o.TermsSchedule, true
}

// HasTermsSchedule returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasTermsSchedule() bool {
	if o != nil && !IsNil(o.TermsSchedule) {
		return true
	}

	return false
}

// SetTermsSchedule gets a reference to the given PolicyDocumentTemplateResponse and assigns it to the TermsSchedule field.
func (o *PolicyDocumentResponse) SetTermsSchedule(v PolicyDocumentTemplateResponse) {
	o.TermsSchedule = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PolicyDocumentResponse) SetToken(v string) {
	o.Token = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *PolicyDocumentResponse) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *PolicyDocumentResponse) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *PolicyDocumentResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

func (o PolicyDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountStatement) {
		toSerialize["account_statement"] = o.AccountStatement
	}
	if !IsNil(o.BenefitsDisclosurePremium) {
		toSerialize["benefits_disclosure_premium"] = o.BenefitsDisclosurePremium
	}
	if !IsNil(o.BenefitsDisclosureTraditional) {
		toSerialize["benefits_disclosure_traditional"] = o.BenefitsDisclosureTraditional
	}
	if !IsNil(o.CardMemberAgreement) {
		toSerialize["card_member_agreement"] = o.CardMemberAgreement
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.EDisclosure) {
		toSerialize["e_disclosure"] = o.EDisclosure
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NoaaMultipleReasonWithDoddFrank) {
		toSerialize["noaa_multiple_reason_with_dodd_frank"] = o.NoaaMultipleReasonWithDoddFrank
	}
	if !IsNil(o.NoaaSingleReason) {
		toSerialize["noaa_single_reason"] = o.NoaaSingleReason
	}
	if !IsNil(o.NoaaSingleReasonWithDoddFrank) {
		toSerialize["noaa_single_reason_with_dodd_frank"] = o.NoaaSingleReasonWithDoddFrank
	}
	if !IsNil(o.PrivacyPolicy) {
		toSerialize["privacy_policy"] = o.PrivacyPolicy
	}
	if !IsNil(o.RewardsDisclosure) {
		toSerialize["rewards_disclosure"] = o.RewardsDisclosure
	}
	if !IsNil(o.SummaryOfCreditTerms) {
		toSerialize["summary_of_credit_terms"] = o.SummaryOfCreditTerms
	}
	if !IsNil(o.TermsSchedule) {
		toSerialize["terms_schedule"] = o.TermsSchedule
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}
	return toSerialize, nil
}

type NullablePolicyDocumentResponse struct {
	value *PolicyDocumentResponse
	isSet bool
}

func (v NullablePolicyDocumentResponse) Get() *PolicyDocumentResponse {
	return v.value
}

func (v *NullablePolicyDocumentResponse) Set(val *PolicyDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDocumentResponse(val *PolicyDocumentResponse) *NullablePolicyDocumentResponse {
	return &NullablePolicyDocumentResponse{value: val, isSet: true}
}

func (v NullablePolicyDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


