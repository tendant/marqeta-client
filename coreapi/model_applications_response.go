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

// checks if the ApplicationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationsResponse{}

// ApplicationsResponse struct for ApplicationsResponse
type ApplicationsResponse struct {
	// Unique identifier of the credit account for which the user is applying.  Returned when retrieving an application in the `APPROVED` state.
	AccountToken *string `json:"account_token,omitempty"`
	// A value of `true` indicates that the user has a non-taxable income source.
	AnyNonTaxableIncome *bool `json:"any_non_taxable_income,omitempty"`
	// Date and time when the application was accepted on the Marqeta platform, in UTC.  Returned if the user accepted their approved application.
	ApplicationAcceptedAt *time.Time `json:"application_accepted_at,omitempty"`
	// Date and time when the application was canceled on the Marqeta platform, in UTC.
	ApplicationCanceledAt *time.Time `json:"application_canceled_at,omitempty"`
	// Date and time when Marqeta accepted the Benefits Disclosure, in UTC.  Returned if the user accepted their approved application.
	BenefitsDisclosureAcceptedAt *time.Time `json:"benefits_disclosure_accepted_at,omitempty"`
	// Unique identifier of the bundle associated with the application.
	BundleToken string `json:"bundle_token"`
	// Date and time when Marqeta accepted the Card Membership Agreement, in UTC.  Returned if the user accepted their approved application.
	CardMembershipAgreementAcceptedAt *time.Time `json:"card_membership_agreement_accepted_at,omitempty"`
	// Date and time when the application was created on the Marqeta platform, in UTC.
	CreatedDate time.Time `json:"created_date"`
	DecisionModel *DecisionsResponse `json:"decision_model,omitempty"`
	// Unique identifier of the decision made on the application.
	DecisionToken *string `json:"decision_token,omitempty"`
	DeviceData *DeviceData `json:"device_data,omitempty"`
	// Date and time when Marqeta accepted the e-Disclosure, in UTC.  Returned if the user accepted their approved application.
	EDisclosureAcceptedAt time.Time `json:"e_disclosure_accepted_at"`
	ErrorDetails *ErrorDetailsResponse `json:"error_details,omitempty"`
	// Customer-defined additional information about the application.
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	// Monthly amount of the mortgage or rent that the user currently pays.
	MonthlyMortgageOrRent *float32 `json:"monthly_mortgage_or_rent,omitempty"`
	// Unique identifier of the offer for a pre-screened applicant.
	OfferId *string `json:"offer_id,omitempty"`
	// Date and time when Marqeta accepted the Pre-qualified Offer Pre-terms, in UTC.  Returned if the user accepted their approved application.
	PrequalifiedOfferPreTermsAcceptedAt *time.Time `json:"prequalified_offer_pre_terms_accepted_at,omitempty"`
	// Whether the primary income source comes from the user being employed, unemployed, self-employment, or another situation.
	PrimaryIncomeSource *string `json:"primary_income_source,omitempty"`
	// Date and time when Marqeta accepted the Privacy Policy, in UTC.  Returned if the user accepted their approved application.
	PrivacyPolicyAcceptedAt time.Time `json:"privacy_policy_accepted_at"`
	// Whether the user owns or rents their residence, or has another situation.
	ResidenceType *string `json:"residence_type,omitempty"`
	// Date and time when Marqeta accepted the Rewards Disclosure, in UTC.  Returned if the user accepted their approved application.
	RewardsDisclosurePostTermsAcceptedAt *time.Time `json:"rewards_disclosure_post_terms_accepted_at,omitempty"`
	// Date and time when Marqeta accepted the Rewards Disclosure, in UTC.  Returned if the user accepted their approved application.
	RewardsDisclosurePreTermsAcceptedAt time.Time `json:"rewards_disclosure_pre_terms_accepted_at"`
	// Date and time when Marqeta accepted the Summary of Credit Terms (SOCT), in UTC.  Returned if the user accepted their approved application.
	SoctAcceptedAt time.Time `json:"soct_accepted_at"`
	State ApplicationResourceState `json:"state"`
	// Date and time when Marqeta accepted the Terms Schedule, in UTC.  Returned if the user accepted their approved application.
	TermScheduleInformationAcceptedAt *time.Time `json:"term_schedule_information_accepted_at,omitempty"`
	// Unique identifier of the application.
	Token string `json:"token"`
	// The total amount of the user's annual income.
	TotalAnnualIncome *float32 `json:"total_annual_income,omitempty"`
	Type ApplicationType `json:"type"`
	// Date and time when the application was last updated on the Marqeta platform, in UTC.
	UpdatedDate time.Time `json:"updated_date"`
	// Unique identifier of the applicant, the user applying for a credit account.
	UserToken string `json:"user_token"`
}

// NewApplicationsResponse instantiates a new ApplicationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationsResponse(bundleToken string, createdDate time.Time, eDisclosureAcceptedAt time.Time, privacyPolicyAcceptedAt time.Time, rewardsDisclosurePreTermsAcceptedAt time.Time, soctAcceptedAt time.Time, state ApplicationResourceState, token string, type_ ApplicationType, updatedDate time.Time, userToken string) *ApplicationsResponse {
	this := ApplicationsResponse{}
	this.BundleToken = bundleToken
	this.CreatedDate = createdDate
	this.EDisclosureAcceptedAt = eDisclosureAcceptedAt
	this.PrivacyPolicyAcceptedAt = privacyPolicyAcceptedAt
	this.RewardsDisclosurePreTermsAcceptedAt = rewardsDisclosurePreTermsAcceptedAt
	this.SoctAcceptedAt = soctAcceptedAt
	this.State = state
	this.Token = token
	this.Type = type_
	this.UpdatedDate = updatedDate
	this.UserToken = userToken
	return &this
}

// NewApplicationsResponseWithDefaults instantiates a new ApplicationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationsResponseWithDefaults() *ApplicationsResponse {
	this := ApplicationsResponse{}
	return &this
}

// GetAccountToken returns the AccountToken field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetAccountToken() string {
	if o == nil || IsNil(o.AccountToken) {
		var ret string
		return ret
	}
	return *o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetAccountTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccountToken) {
		return nil, false
	}
	return o.AccountToken, true
}

// HasAccountToken returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasAccountToken() bool {
	if o != nil && !IsNil(o.AccountToken) {
		return true
	}

	return false
}

// SetAccountToken gets a reference to the given string and assigns it to the AccountToken field.
func (o *ApplicationsResponse) SetAccountToken(v string) {
	o.AccountToken = &v
}

// GetAnyNonTaxableIncome returns the AnyNonTaxableIncome field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetAnyNonTaxableIncome() bool {
	if o == nil || IsNil(o.AnyNonTaxableIncome) {
		var ret bool
		return ret
	}
	return *o.AnyNonTaxableIncome
}

// GetAnyNonTaxableIncomeOk returns a tuple with the AnyNonTaxableIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetAnyNonTaxableIncomeOk() (*bool, bool) {
	if o == nil || IsNil(o.AnyNonTaxableIncome) {
		return nil, false
	}
	return o.AnyNonTaxableIncome, true
}

// HasAnyNonTaxableIncome returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasAnyNonTaxableIncome() bool {
	if o != nil && !IsNil(o.AnyNonTaxableIncome) {
		return true
	}

	return false
}

// SetAnyNonTaxableIncome gets a reference to the given bool and assigns it to the AnyNonTaxableIncome field.
func (o *ApplicationsResponse) SetAnyNonTaxableIncome(v bool) {
	o.AnyNonTaxableIncome = &v
}

// GetApplicationAcceptedAt returns the ApplicationAcceptedAt field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetApplicationAcceptedAt() time.Time {
	if o == nil || IsNil(o.ApplicationAcceptedAt) {
		var ret time.Time
		return ret
	}
	return *o.ApplicationAcceptedAt
}

// GetApplicationAcceptedAtOk returns a tuple with the ApplicationAcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetApplicationAcceptedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ApplicationAcceptedAt) {
		return nil, false
	}
	return o.ApplicationAcceptedAt, true
}

// HasApplicationAcceptedAt returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasApplicationAcceptedAt() bool {
	if o != nil && !IsNil(o.ApplicationAcceptedAt) {
		return true
	}

	return false
}

// SetApplicationAcceptedAt gets a reference to the given time.Time and assigns it to the ApplicationAcceptedAt field.
func (o *ApplicationsResponse) SetApplicationAcceptedAt(v time.Time) {
	o.ApplicationAcceptedAt = &v
}

// GetApplicationCanceledAt returns the ApplicationCanceledAt field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetApplicationCanceledAt() time.Time {
	if o == nil || IsNil(o.ApplicationCanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.ApplicationCanceledAt
}

// GetApplicationCanceledAtOk returns a tuple with the ApplicationCanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetApplicationCanceledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ApplicationCanceledAt) {
		return nil, false
	}
	return o.ApplicationCanceledAt, true
}

// HasApplicationCanceledAt returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasApplicationCanceledAt() bool {
	if o != nil && !IsNil(o.ApplicationCanceledAt) {
		return true
	}

	return false
}

// SetApplicationCanceledAt gets a reference to the given time.Time and assigns it to the ApplicationCanceledAt field.
func (o *ApplicationsResponse) SetApplicationCanceledAt(v time.Time) {
	o.ApplicationCanceledAt = &v
}

// GetBenefitsDisclosureAcceptedAt returns the BenefitsDisclosureAcceptedAt field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetBenefitsDisclosureAcceptedAt() time.Time {
	if o == nil || IsNil(o.BenefitsDisclosureAcceptedAt) {
		var ret time.Time
		return ret
	}
	return *o.BenefitsDisclosureAcceptedAt
}

// GetBenefitsDisclosureAcceptedAtOk returns a tuple with the BenefitsDisclosureAcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetBenefitsDisclosureAcceptedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BenefitsDisclosureAcceptedAt) {
		return nil, false
	}
	return o.BenefitsDisclosureAcceptedAt, true
}

// HasBenefitsDisclosureAcceptedAt returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasBenefitsDisclosureAcceptedAt() bool {
	if o != nil && !IsNil(o.BenefitsDisclosureAcceptedAt) {
		return true
	}

	return false
}

// SetBenefitsDisclosureAcceptedAt gets a reference to the given time.Time and assigns it to the BenefitsDisclosureAcceptedAt field.
func (o *ApplicationsResponse) SetBenefitsDisclosureAcceptedAt(v time.Time) {
	o.BenefitsDisclosureAcceptedAt = &v
}

// GetBundleToken returns the BundleToken field value
func (o *ApplicationsResponse) GetBundleToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BundleToken
}

// GetBundleTokenOk returns a tuple with the BundleToken field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetBundleTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleToken, true
}

// SetBundleToken sets field value
func (o *ApplicationsResponse) SetBundleToken(v string) {
	o.BundleToken = v
}

// GetCardMembershipAgreementAcceptedAt returns the CardMembershipAgreementAcceptedAt field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetCardMembershipAgreementAcceptedAt() time.Time {
	if o == nil || IsNil(o.CardMembershipAgreementAcceptedAt) {
		var ret time.Time
		return ret
	}
	return *o.CardMembershipAgreementAcceptedAt
}

// GetCardMembershipAgreementAcceptedAtOk returns a tuple with the CardMembershipAgreementAcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetCardMembershipAgreementAcceptedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CardMembershipAgreementAcceptedAt) {
		return nil, false
	}
	return o.CardMembershipAgreementAcceptedAt, true
}

// HasCardMembershipAgreementAcceptedAt returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasCardMembershipAgreementAcceptedAt() bool {
	if o != nil && !IsNil(o.CardMembershipAgreementAcceptedAt) {
		return true
	}

	return false
}

// SetCardMembershipAgreementAcceptedAt gets a reference to the given time.Time and assigns it to the CardMembershipAgreementAcceptedAt field.
func (o *ApplicationsResponse) SetCardMembershipAgreementAcceptedAt(v time.Time) {
	o.CardMembershipAgreementAcceptedAt = &v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ApplicationsResponse) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ApplicationsResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetDecisionModel returns the DecisionModel field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetDecisionModel() DecisionsResponse {
	if o == nil || IsNil(o.DecisionModel) {
		var ret DecisionsResponse
		return ret
	}
	return *o.DecisionModel
}

// GetDecisionModelOk returns a tuple with the DecisionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetDecisionModelOk() (*DecisionsResponse, bool) {
	if o == nil || IsNil(o.DecisionModel) {
		return nil, false
	}
	return o.DecisionModel, true
}

// HasDecisionModel returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasDecisionModel() bool {
	if o != nil && !IsNil(o.DecisionModel) {
		return true
	}

	return false
}

// SetDecisionModel gets a reference to the given DecisionsResponse and assigns it to the DecisionModel field.
func (o *ApplicationsResponse) SetDecisionModel(v DecisionsResponse) {
	o.DecisionModel = &v
}

// GetDecisionToken returns the DecisionToken field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetDecisionToken() string {
	if o == nil || IsNil(o.DecisionToken) {
		var ret string
		return ret
	}
	return *o.DecisionToken
}

// GetDecisionTokenOk returns a tuple with the DecisionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetDecisionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.DecisionToken) {
		return nil, false
	}
	return o.DecisionToken, true
}

// HasDecisionToken returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasDecisionToken() bool {
	if o != nil && !IsNil(o.DecisionToken) {
		return true
	}

	return false
}

// SetDecisionToken gets a reference to the given string and assigns it to the DecisionToken field.
func (o *ApplicationsResponse) SetDecisionToken(v string) {
	o.DecisionToken = &v
}

// GetDeviceData returns the DeviceData field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetDeviceData() DeviceData {
	if o == nil || IsNil(o.DeviceData) {
		var ret DeviceData
		return ret
	}
	return *o.DeviceData
}

// GetDeviceDataOk returns a tuple with the DeviceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetDeviceDataOk() (*DeviceData, bool) {
	if o == nil || IsNil(o.DeviceData) {
		return nil, false
	}
	return o.DeviceData, true
}

// HasDeviceData returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasDeviceData() bool {
	if o != nil && !IsNil(o.DeviceData) {
		return true
	}

	return false
}

// SetDeviceData gets a reference to the given DeviceData and assigns it to the DeviceData field.
func (o *ApplicationsResponse) SetDeviceData(v DeviceData) {
	o.DeviceData = &v
}

// GetEDisclosureAcceptedAt returns the EDisclosureAcceptedAt field value
func (o *ApplicationsResponse) GetEDisclosureAcceptedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EDisclosureAcceptedAt
}

// GetEDisclosureAcceptedAtOk returns a tuple with the EDisclosureAcceptedAt field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetEDisclosureAcceptedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EDisclosureAcceptedAt, true
}

// SetEDisclosureAcceptedAt sets field value
func (o *ApplicationsResponse) SetEDisclosureAcceptedAt(v time.Time) {
	o.EDisclosureAcceptedAt = v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetErrorDetails() ErrorDetailsResponse {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetailsResponse
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetErrorDetailsOk() (*ErrorDetailsResponse, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetailsResponse and assigns it to the ErrorDetails field.
func (o *ApplicationsResponse) SetErrorDetails(v ErrorDetailsResponse) {
	o.ErrorDetails = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetMetaData() map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return map[string]interface{}{}, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *ApplicationsResponse) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetMonthlyMortgageOrRent returns the MonthlyMortgageOrRent field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetMonthlyMortgageOrRent() float32 {
	if o == nil || IsNil(o.MonthlyMortgageOrRent) {
		var ret float32
		return ret
	}
	return *o.MonthlyMortgageOrRent
}

// GetMonthlyMortgageOrRentOk returns a tuple with the MonthlyMortgageOrRent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetMonthlyMortgageOrRentOk() (*float32, bool) {
	if o == nil || IsNil(o.MonthlyMortgageOrRent) {
		return nil, false
	}
	return o.MonthlyMortgageOrRent, true
}

// HasMonthlyMortgageOrRent returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasMonthlyMortgageOrRent() bool {
	if o != nil && !IsNil(o.MonthlyMortgageOrRent) {
		return true
	}

	return false
}

// SetMonthlyMortgageOrRent gets a reference to the given float32 and assigns it to the MonthlyMortgageOrRent field.
func (o *ApplicationsResponse) SetMonthlyMortgageOrRent(v float32) {
	o.MonthlyMortgageOrRent = &v
}

// GetOfferId returns the OfferId field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetOfferId() string {
	if o == nil || IsNil(o.OfferId) {
		var ret string
		return ret
	}
	return *o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetOfferIdOk() (*string, bool) {
	if o == nil || IsNil(o.OfferId) {
		return nil, false
	}
	return o.OfferId, true
}

// HasOfferId returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasOfferId() bool {
	if o != nil && !IsNil(o.OfferId) {
		return true
	}

	return false
}

// SetOfferId gets a reference to the given string and assigns it to the OfferId field.
func (o *ApplicationsResponse) SetOfferId(v string) {
	o.OfferId = &v
}

// GetPrequalifiedOfferPreTermsAcceptedAt returns the PrequalifiedOfferPreTermsAcceptedAt field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetPrequalifiedOfferPreTermsAcceptedAt() time.Time {
	if o == nil || IsNil(o.PrequalifiedOfferPreTermsAcceptedAt) {
		var ret time.Time
		return ret
	}
	return *o.PrequalifiedOfferPreTermsAcceptedAt
}

// GetPrequalifiedOfferPreTermsAcceptedAtOk returns a tuple with the PrequalifiedOfferPreTermsAcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetPrequalifiedOfferPreTermsAcceptedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PrequalifiedOfferPreTermsAcceptedAt) {
		return nil, false
	}
	return o.PrequalifiedOfferPreTermsAcceptedAt, true
}

// HasPrequalifiedOfferPreTermsAcceptedAt returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasPrequalifiedOfferPreTermsAcceptedAt() bool {
	if o != nil && !IsNil(o.PrequalifiedOfferPreTermsAcceptedAt) {
		return true
	}

	return false
}

// SetPrequalifiedOfferPreTermsAcceptedAt gets a reference to the given time.Time and assigns it to the PrequalifiedOfferPreTermsAcceptedAt field.
func (o *ApplicationsResponse) SetPrequalifiedOfferPreTermsAcceptedAt(v time.Time) {
	o.PrequalifiedOfferPreTermsAcceptedAt = &v
}

// GetPrimaryIncomeSource returns the PrimaryIncomeSource field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetPrimaryIncomeSource() string {
	if o == nil || IsNil(o.PrimaryIncomeSource) {
		var ret string
		return ret
	}
	return *o.PrimaryIncomeSource
}

// GetPrimaryIncomeSourceOk returns a tuple with the PrimaryIncomeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetPrimaryIncomeSourceOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryIncomeSource) {
		return nil, false
	}
	return o.PrimaryIncomeSource, true
}

// HasPrimaryIncomeSource returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasPrimaryIncomeSource() bool {
	if o != nil && !IsNil(o.PrimaryIncomeSource) {
		return true
	}

	return false
}

// SetPrimaryIncomeSource gets a reference to the given string and assigns it to the PrimaryIncomeSource field.
func (o *ApplicationsResponse) SetPrimaryIncomeSource(v string) {
	o.PrimaryIncomeSource = &v
}

// GetPrivacyPolicyAcceptedAt returns the PrivacyPolicyAcceptedAt field value
func (o *ApplicationsResponse) GetPrivacyPolicyAcceptedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PrivacyPolicyAcceptedAt
}

// GetPrivacyPolicyAcceptedAtOk returns a tuple with the PrivacyPolicyAcceptedAt field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetPrivacyPolicyAcceptedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyPolicyAcceptedAt, true
}

// SetPrivacyPolicyAcceptedAt sets field value
func (o *ApplicationsResponse) SetPrivacyPolicyAcceptedAt(v time.Time) {
	o.PrivacyPolicyAcceptedAt = v
}

// GetResidenceType returns the ResidenceType field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetResidenceType() string {
	if o == nil || IsNil(o.ResidenceType) {
		var ret string
		return ret
	}
	return *o.ResidenceType
}

// GetResidenceTypeOk returns a tuple with the ResidenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetResidenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResidenceType) {
		return nil, false
	}
	return o.ResidenceType, true
}

// HasResidenceType returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasResidenceType() bool {
	if o != nil && !IsNil(o.ResidenceType) {
		return true
	}

	return false
}

// SetResidenceType gets a reference to the given string and assigns it to the ResidenceType field.
func (o *ApplicationsResponse) SetResidenceType(v string) {
	o.ResidenceType = &v
}

// GetRewardsDisclosurePostTermsAcceptedAt returns the RewardsDisclosurePostTermsAcceptedAt field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetRewardsDisclosurePostTermsAcceptedAt() time.Time {
	if o == nil || IsNil(o.RewardsDisclosurePostTermsAcceptedAt) {
		var ret time.Time
		return ret
	}
	return *o.RewardsDisclosurePostTermsAcceptedAt
}

// GetRewardsDisclosurePostTermsAcceptedAtOk returns a tuple with the RewardsDisclosurePostTermsAcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetRewardsDisclosurePostTermsAcceptedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RewardsDisclosurePostTermsAcceptedAt) {
		return nil, false
	}
	return o.RewardsDisclosurePostTermsAcceptedAt, true
}

// HasRewardsDisclosurePostTermsAcceptedAt returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasRewardsDisclosurePostTermsAcceptedAt() bool {
	if o != nil && !IsNil(o.RewardsDisclosurePostTermsAcceptedAt) {
		return true
	}

	return false
}

// SetRewardsDisclosurePostTermsAcceptedAt gets a reference to the given time.Time and assigns it to the RewardsDisclosurePostTermsAcceptedAt field.
func (o *ApplicationsResponse) SetRewardsDisclosurePostTermsAcceptedAt(v time.Time) {
	o.RewardsDisclosurePostTermsAcceptedAt = &v
}

// GetRewardsDisclosurePreTermsAcceptedAt returns the RewardsDisclosurePreTermsAcceptedAt field value
func (o *ApplicationsResponse) GetRewardsDisclosurePreTermsAcceptedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RewardsDisclosurePreTermsAcceptedAt
}

// GetRewardsDisclosurePreTermsAcceptedAtOk returns a tuple with the RewardsDisclosurePreTermsAcceptedAt field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetRewardsDisclosurePreTermsAcceptedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardsDisclosurePreTermsAcceptedAt, true
}

// SetRewardsDisclosurePreTermsAcceptedAt sets field value
func (o *ApplicationsResponse) SetRewardsDisclosurePreTermsAcceptedAt(v time.Time) {
	o.RewardsDisclosurePreTermsAcceptedAt = v
}

// GetSoctAcceptedAt returns the SoctAcceptedAt field value
func (o *ApplicationsResponse) GetSoctAcceptedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SoctAcceptedAt
}

// GetSoctAcceptedAtOk returns a tuple with the SoctAcceptedAt field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetSoctAcceptedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoctAcceptedAt, true
}

// SetSoctAcceptedAt sets field value
func (o *ApplicationsResponse) SetSoctAcceptedAt(v time.Time) {
	o.SoctAcceptedAt = v
}

// GetState returns the State field value
func (o *ApplicationsResponse) GetState() ApplicationResourceState {
	if o == nil {
		var ret ApplicationResourceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetStateOk() (*ApplicationResourceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ApplicationsResponse) SetState(v ApplicationResourceState) {
	o.State = v
}

// GetTermScheduleInformationAcceptedAt returns the TermScheduleInformationAcceptedAt field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetTermScheduleInformationAcceptedAt() time.Time {
	if o == nil || IsNil(o.TermScheduleInformationAcceptedAt) {
		var ret time.Time
		return ret
	}
	return *o.TermScheduleInformationAcceptedAt
}

// GetTermScheduleInformationAcceptedAtOk returns a tuple with the TermScheduleInformationAcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetTermScheduleInformationAcceptedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TermScheduleInformationAcceptedAt) {
		return nil, false
	}
	return o.TermScheduleInformationAcceptedAt, true
}

// HasTermScheduleInformationAcceptedAt returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasTermScheduleInformationAcceptedAt() bool {
	if o != nil && !IsNil(o.TermScheduleInformationAcceptedAt) {
		return true
	}

	return false
}

// SetTermScheduleInformationAcceptedAt gets a reference to the given time.Time and assigns it to the TermScheduleInformationAcceptedAt field.
func (o *ApplicationsResponse) SetTermScheduleInformationAcceptedAt(v time.Time) {
	o.TermScheduleInformationAcceptedAt = &v
}

// GetToken returns the Token field value
func (o *ApplicationsResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ApplicationsResponse) SetToken(v string) {
	o.Token = v
}

// GetTotalAnnualIncome returns the TotalAnnualIncome field value if set, zero value otherwise.
func (o *ApplicationsResponse) GetTotalAnnualIncome() float32 {
	if o == nil || IsNil(o.TotalAnnualIncome) {
		var ret float32
		return ret
	}
	return *o.TotalAnnualIncome
}

// GetTotalAnnualIncomeOk returns a tuple with the TotalAnnualIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetTotalAnnualIncomeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalAnnualIncome) {
		return nil, false
	}
	return o.TotalAnnualIncome, true
}

// HasTotalAnnualIncome returns a boolean if a field has been set.
func (o *ApplicationsResponse) HasTotalAnnualIncome() bool {
	if o != nil && !IsNil(o.TotalAnnualIncome) {
		return true
	}

	return false
}

// SetTotalAnnualIncome gets a reference to the given float32 and assigns it to the TotalAnnualIncome field.
func (o *ApplicationsResponse) SetTotalAnnualIncome(v float32) {
	o.TotalAnnualIncome = &v
}

// GetType returns the Type field value
func (o *ApplicationsResponse) GetType() ApplicationType {
	if o == nil {
		var ret ApplicationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetTypeOk() (*ApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationsResponse) SetType(v ApplicationType) {
	o.Type = v
}

// GetUpdatedDate returns the UpdatedDate field value
func (o *ApplicationsResponse) GetUpdatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedDate, true
}

// SetUpdatedDate sets field value
func (o *ApplicationsResponse) SetUpdatedDate(v time.Time) {
	o.UpdatedDate = v
}

// GetUserToken returns the UserToken field value
func (o *ApplicationsResponse) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ApplicationsResponse) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *ApplicationsResponse) SetUserToken(v string) {
	o.UserToken = v
}

func (o ApplicationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountToken) {
		toSerialize["account_token"] = o.AccountToken
	}
	if !IsNil(o.AnyNonTaxableIncome) {
		toSerialize["any_non_taxable_income"] = o.AnyNonTaxableIncome
	}
	if !IsNil(o.ApplicationAcceptedAt) {
		toSerialize["application_accepted_at"] = o.ApplicationAcceptedAt
	}
	if !IsNil(o.ApplicationCanceledAt) {
		toSerialize["application_canceled_at"] = o.ApplicationCanceledAt
	}
	if !IsNil(o.BenefitsDisclosureAcceptedAt) {
		toSerialize["benefits_disclosure_accepted_at"] = o.BenefitsDisclosureAcceptedAt
	}
	toSerialize["bundle_token"] = o.BundleToken
	if !IsNil(o.CardMembershipAgreementAcceptedAt) {
		toSerialize["card_membership_agreement_accepted_at"] = o.CardMembershipAgreementAcceptedAt
	}
	toSerialize["created_date"] = o.CreatedDate
	if !IsNil(o.DecisionModel) {
		toSerialize["decision_model"] = o.DecisionModel
	}
	if !IsNil(o.DecisionToken) {
		toSerialize["decision_token"] = o.DecisionToken
	}
	if !IsNil(o.DeviceData) {
		toSerialize["device_data"] = o.DeviceData
	}
	toSerialize["e_disclosure_accepted_at"] = o.EDisclosureAcceptedAt
	if !IsNil(o.ErrorDetails) {
		toSerialize["error_details"] = o.ErrorDetails
	}
	if !IsNil(o.MetaData) {
		toSerialize["meta_data"] = o.MetaData
	}
	if !IsNil(o.MonthlyMortgageOrRent) {
		toSerialize["monthly_mortgage_or_rent"] = o.MonthlyMortgageOrRent
	}
	if !IsNil(o.OfferId) {
		toSerialize["offer_id"] = o.OfferId
	}
	if !IsNil(o.PrequalifiedOfferPreTermsAcceptedAt) {
		toSerialize["prequalified_offer_pre_terms_accepted_at"] = o.PrequalifiedOfferPreTermsAcceptedAt
	}
	if !IsNil(o.PrimaryIncomeSource) {
		toSerialize["primary_income_source"] = o.PrimaryIncomeSource
	}
	toSerialize["privacy_policy_accepted_at"] = o.PrivacyPolicyAcceptedAt
	if !IsNil(o.ResidenceType) {
		toSerialize["residence_type"] = o.ResidenceType
	}
	if !IsNil(o.RewardsDisclosurePostTermsAcceptedAt) {
		toSerialize["rewards_disclosure_post_terms_accepted_at"] = o.RewardsDisclosurePostTermsAcceptedAt
	}
	toSerialize["rewards_disclosure_pre_terms_accepted_at"] = o.RewardsDisclosurePreTermsAcceptedAt
	toSerialize["soct_accepted_at"] = o.SoctAcceptedAt
	toSerialize["state"] = o.State
	if !IsNil(o.TermScheduleInformationAcceptedAt) {
		toSerialize["term_schedule_information_accepted_at"] = o.TermScheduleInformationAcceptedAt
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.TotalAnnualIncome) {
		toSerialize["total_annual_income"] = o.TotalAnnualIncome
	}
	toSerialize["type"] = o.Type
	toSerialize["updated_date"] = o.UpdatedDate
	toSerialize["user_token"] = o.UserToken
	return toSerialize, nil
}

type NullableApplicationsResponse struct {
	value *ApplicationsResponse
	isSet bool
}

func (v NullableApplicationsResponse) Get() *ApplicationsResponse {
	return v.value
}

func (v *NullableApplicationsResponse) Set(val *ApplicationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationsResponse(val *ApplicationsResponse) *NullableApplicationsResponse {
	return &NullableApplicationsResponse{value: val, isSet: true}
}

func (v NullableApplicationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


