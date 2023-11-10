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

// checks if the PolicyDocumentUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyDocumentUpdateReq{}

// PolicyDocumentUpdateReq Request details to update a document policy.
type PolicyDocumentUpdateReq struct {
	AccountStatement PolicyDocumentTemplateReq `json:"account_statement"`
	BenefitsDisclosurePremium PolicyDocumentAssetReq `json:"benefits_disclosure_premium"`
	BenefitsDisclosureTraditional PolicyDocumentAssetReq `json:"benefits_disclosure_traditional"`
	CardMemberAgreement PolicyDocumentAssetReq `json:"card_member_agreement"`
	EDisclosure PolicyDocumentAssetReq `json:"e_disclosure"`
	// Name of the document policy.
	Name string `json:"name"`
	NoaaMultipleReasonWithDoddFrank PolicyDocumentTemplateReq `json:"noaa_multiple_reason_with_dodd_frank"`
	NoaaSingleReason PolicyDocumentTemplateReq `json:"noaa_single_reason"`
	NoaaSingleReasonWithDoddFrank PolicyDocumentTemplateReq `json:"noaa_single_reason_with_dodd_frank"`
	PrivacyPolicy PolicyDocumentAssetReq `json:"privacy_policy"`
	RewardsDisclosure *PolicyDocumentAssetAndTemplateReq `json:"rewards_disclosure,omitempty"`
	SummaryOfCreditTerms PolicyDocumentAssetAndTemplateReq `json:"summary_of_credit_terms"`
	TermsSchedule PolicyDocumentTemplateReq `json:"terms_schedule"`
}

// NewPolicyDocumentUpdateReq instantiates a new PolicyDocumentUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyDocumentUpdateReq(accountStatement PolicyDocumentTemplateReq, benefitsDisclosurePremium PolicyDocumentAssetReq, benefitsDisclosureTraditional PolicyDocumentAssetReq, cardMemberAgreement PolicyDocumentAssetReq, eDisclosure PolicyDocumentAssetReq, name string, noaaMultipleReasonWithDoddFrank PolicyDocumentTemplateReq, noaaSingleReason PolicyDocumentTemplateReq, noaaSingleReasonWithDoddFrank PolicyDocumentTemplateReq, privacyPolicy PolicyDocumentAssetReq, summaryOfCreditTerms PolicyDocumentAssetAndTemplateReq, termsSchedule PolicyDocumentTemplateReq) *PolicyDocumentUpdateReq {
	this := PolicyDocumentUpdateReq{}
	this.AccountStatement = accountStatement
	this.BenefitsDisclosurePremium = benefitsDisclosurePremium
	this.BenefitsDisclosureTraditional = benefitsDisclosureTraditional
	this.CardMemberAgreement = cardMemberAgreement
	this.EDisclosure = eDisclosure
	this.Name = name
	this.NoaaMultipleReasonWithDoddFrank = noaaMultipleReasonWithDoddFrank
	this.NoaaSingleReason = noaaSingleReason
	this.NoaaSingleReasonWithDoddFrank = noaaSingleReasonWithDoddFrank
	this.PrivacyPolicy = privacyPolicy
	this.SummaryOfCreditTerms = summaryOfCreditTerms
	this.TermsSchedule = termsSchedule
	return &this
}

// NewPolicyDocumentUpdateReqWithDefaults instantiates a new PolicyDocumentUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyDocumentUpdateReqWithDefaults() *PolicyDocumentUpdateReq {
	this := PolicyDocumentUpdateReq{}
	return &this
}

// GetAccountStatement returns the AccountStatement field value
func (o *PolicyDocumentUpdateReq) GetAccountStatement() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.AccountStatement
}

// GetAccountStatementOk returns a tuple with the AccountStatement field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetAccountStatementOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountStatement, true
}

// SetAccountStatement sets field value
func (o *PolicyDocumentUpdateReq) SetAccountStatement(v PolicyDocumentTemplateReq) {
	o.AccountStatement = v
}

// GetBenefitsDisclosurePremium returns the BenefitsDisclosurePremium field value
func (o *PolicyDocumentUpdateReq) GetBenefitsDisclosurePremium() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.BenefitsDisclosurePremium
}

// GetBenefitsDisclosurePremiumOk returns a tuple with the BenefitsDisclosurePremium field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetBenefitsDisclosurePremiumOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitsDisclosurePremium, true
}

// SetBenefitsDisclosurePremium sets field value
func (o *PolicyDocumentUpdateReq) SetBenefitsDisclosurePremium(v PolicyDocumentAssetReq) {
	o.BenefitsDisclosurePremium = v
}

// GetBenefitsDisclosureTraditional returns the BenefitsDisclosureTraditional field value
func (o *PolicyDocumentUpdateReq) GetBenefitsDisclosureTraditional() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.BenefitsDisclosureTraditional
}

// GetBenefitsDisclosureTraditionalOk returns a tuple with the BenefitsDisclosureTraditional field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetBenefitsDisclosureTraditionalOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitsDisclosureTraditional, true
}

// SetBenefitsDisclosureTraditional sets field value
func (o *PolicyDocumentUpdateReq) SetBenefitsDisclosureTraditional(v PolicyDocumentAssetReq) {
	o.BenefitsDisclosureTraditional = v
}

// GetCardMemberAgreement returns the CardMemberAgreement field value
func (o *PolicyDocumentUpdateReq) GetCardMemberAgreement() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.CardMemberAgreement
}

// GetCardMemberAgreementOk returns a tuple with the CardMemberAgreement field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetCardMemberAgreementOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardMemberAgreement, true
}

// SetCardMemberAgreement sets field value
func (o *PolicyDocumentUpdateReq) SetCardMemberAgreement(v PolicyDocumentAssetReq) {
	o.CardMemberAgreement = v
}

// GetEDisclosure returns the EDisclosure field value
func (o *PolicyDocumentUpdateReq) GetEDisclosure() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.EDisclosure
}

// GetEDisclosureOk returns a tuple with the EDisclosure field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetEDisclosureOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EDisclosure, true
}

// SetEDisclosure sets field value
func (o *PolicyDocumentUpdateReq) SetEDisclosure(v PolicyDocumentAssetReq) {
	o.EDisclosure = v
}

// GetName returns the Name field value
func (o *PolicyDocumentUpdateReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyDocumentUpdateReq) SetName(v string) {
	o.Name = v
}

// GetNoaaMultipleReasonWithDoddFrank returns the NoaaMultipleReasonWithDoddFrank field value
func (o *PolicyDocumentUpdateReq) GetNoaaMultipleReasonWithDoddFrank() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.NoaaMultipleReasonWithDoddFrank
}

// GetNoaaMultipleReasonWithDoddFrankOk returns a tuple with the NoaaMultipleReasonWithDoddFrank field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetNoaaMultipleReasonWithDoddFrankOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoaaMultipleReasonWithDoddFrank, true
}

// SetNoaaMultipleReasonWithDoddFrank sets field value
func (o *PolicyDocumentUpdateReq) SetNoaaMultipleReasonWithDoddFrank(v PolicyDocumentTemplateReq) {
	o.NoaaMultipleReasonWithDoddFrank = v
}

// GetNoaaSingleReason returns the NoaaSingleReason field value
func (o *PolicyDocumentUpdateReq) GetNoaaSingleReason() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.NoaaSingleReason
}

// GetNoaaSingleReasonOk returns a tuple with the NoaaSingleReason field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetNoaaSingleReasonOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoaaSingleReason, true
}

// SetNoaaSingleReason sets field value
func (o *PolicyDocumentUpdateReq) SetNoaaSingleReason(v PolicyDocumentTemplateReq) {
	o.NoaaSingleReason = v
}

// GetNoaaSingleReasonWithDoddFrank returns the NoaaSingleReasonWithDoddFrank field value
func (o *PolicyDocumentUpdateReq) GetNoaaSingleReasonWithDoddFrank() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.NoaaSingleReasonWithDoddFrank
}

// GetNoaaSingleReasonWithDoddFrankOk returns a tuple with the NoaaSingleReasonWithDoddFrank field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetNoaaSingleReasonWithDoddFrankOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoaaSingleReasonWithDoddFrank, true
}

// SetNoaaSingleReasonWithDoddFrank sets field value
func (o *PolicyDocumentUpdateReq) SetNoaaSingleReasonWithDoddFrank(v PolicyDocumentTemplateReq) {
	o.NoaaSingleReasonWithDoddFrank = v
}

// GetPrivacyPolicy returns the PrivacyPolicy field value
func (o *PolicyDocumentUpdateReq) GetPrivacyPolicy() PolicyDocumentAssetReq {
	if o == nil {
		var ret PolicyDocumentAssetReq
		return ret
	}

	return o.PrivacyPolicy
}

// GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetPrivacyPolicyOk() (*PolicyDocumentAssetReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyPolicy, true
}

// SetPrivacyPolicy sets field value
func (o *PolicyDocumentUpdateReq) SetPrivacyPolicy(v PolicyDocumentAssetReq) {
	o.PrivacyPolicy = v
}

// GetRewardsDisclosure returns the RewardsDisclosure field value if set, zero value otherwise.
func (o *PolicyDocumentUpdateReq) GetRewardsDisclosure() PolicyDocumentAssetAndTemplateReq {
	if o == nil || IsNil(o.RewardsDisclosure) {
		var ret PolicyDocumentAssetAndTemplateReq
		return ret
	}
	return *o.RewardsDisclosure
}

// GetRewardsDisclosureOk returns a tuple with the RewardsDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetRewardsDisclosureOk() (*PolicyDocumentAssetAndTemplateReq, bool) {
	if o == nil || IsNil(o.RewardsDisclosure) {
		return nil, false
	}
	return o.RewardsDisclosure, true
}

// HasRewardsDisclosure returns a boolean if a field has been set.
func (o *PolicyDocumentUpdateReq) HasRewardsDisclosure() bool {
	if o != nil && !IsNil(o.RewardsDisclosure) {
		return true
	}

	return false
}

// SetRewardsDisclosure gets a reference to the given PolicyDocumentAssetAndTemplateReq and assigns it to the RewardsDisclosure field.
func (o *PolicyDocumentUpdateReq) SetRewardsDisclosure(v PolicyDocumentAssetAndTemplateReq) {
	o.RewardsDisclosure = &v
}

// GetSummaryOfCreditTerms returns the SummaryOfCreditTerms field value
func (o *PolicyDocumentUpdateReq) GetSummaryOfCreditTerms() PolicyDocumentAssetAndTemplateReq {
	if o == nil {
		var ret PolicyDocumentAssetAndTemplateReq
		return ret
	}

	return o.SummaryOfCreditTerms
}

// GetSummaryOfCreditTermsOk returns a tuple with the SummaryOfCreditTerms field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetSummaryOfCreditTermsOk() (*PolicyDocumentAssetAndTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SummaryOfCreditTerms, true
}

// SetSummaryOfCreditTerms sets field value
func (o *PolicyDocumentUpdateReq) SetSummaryOfCreditTerms(v PolicyDocumentAssetAndTemplateReq) {
	o.SummaryOfCreditTerms = v
}

// GetTermsSchedule returns the TermsSchedule field value
func (o *PolicyDocumentUpdateReq) GetTermsSchedule() PolicyDocumentTemplateReq {
	if o == nil {
		var ret PolicyDocumentTemplateReq
		return ret
	}

	return o.TermsSchedule
}

// GetTermsScheduleOk returns a tuple with the TermsSchedule field value
// and a boolean to check if the value has been set.
func (o *PolicyDocumentUpdateReq) GetTermsScheduleOk() (*PolicyDocumentTemplateReq, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermsSchedule, true
}

// SetTermsSchedule sets field value
func (o *PolicyDocumentUpdateReq) SetTermsSchedule(v PolicyDocumentTemplateReq) {
	o.TermsSchedule = v
}

func (o PolicyDocumentUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyDocumentUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_statement"] = o.AccountStatement
	toSerialize["benefits_disclosure_premium"] = o.BenefitsDisclosurePremium
	toSerialize["benefits_disclosure_traditional"] = o.BenefitsDisclosureTraditional
	toSerialize["card_member_agreement"] = o.CardMemberAgreement
	toSerialize["e_disclosure"] = o.EDisclosure
	toSerialize["name"] = o.Name
	toSerialize["noaa_multiple_reason_with_dodd_frank"] = o.NoaaMultipleReasonWithDoddFrank
	toSerialize["noaa_single_reason"] = o.NoaaSingleReason
	toSerialize["noaa_single_reason_with_dodd_frank"] = o.NoaaSingleReasonWithDoddFrank
	toSerialize["privacy_policy"] = o.PrivacyPolicy
	if !IsNil(o.RewardsDisclosure) {
		toSerialize["rewards_disclosure"] = o.RewardsDisclosure
	}
	toSerialize["summary_of_credit_terms"] = o.SummaryOfCreditTerms
	toSerialize["terms_schedule"] = o.TermsSchedule
	return toSerialize, nil
}

type NullablePolicyDocumentUpdateReq struct {
	value *PolicyDocumentUpdateReq
	isSet bool
}

func (v NullablePolicyDocumentUpdateReq) Get() *PolicyDocumentUpdateReq {
	return v.value
}

func (v *NullablePolicyDocumentUpdateReq) Set(val *PolicyDocumentUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDocumentUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDocumentUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDocumentUpdateReq(val *PolicyDocumentUpdateReq) *NullablePolicyDocumentUpdateReq {
	return &NullablePolicyDocumentUpdateReq{value: val, isSet: true}
}

func (v NullablePolicyDocumentUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDocumentUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


