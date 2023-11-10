# ApplicationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | Unique identifier of the credit account for which the user is applying.  Returned when retrieving an application in the &#x60;APPROVED&#x60; state. | [optional] 
**AnyNonTaxableIncome** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that the user has a non-taxable income source. | [optional] 
**ApplicationAcceptedAt** | Pointer to **time.Time** | Date and time when the application was accepted on the Marqeta platform, in UTC.  Returned if the user accepted their approved application. | [optional] 
**ApplicationCanceledAt** | Pointer to **time.Time** | Date and time when the application was canceled on the Marqeta platform, in UTC. | [optional] 
**BenefitsDisclosureAcceptedAt** | Pointer to **time.Time** | Date and time when Marqeta accepted the Benefits Disclosure, in UTC.  Returned if the user accepted their approved application. | [optional] 
**BundleToken** | **string** | Unique identifier of the bundle associated with the application. | 
**CardMembershipAgreementAcceptedAt** | Pointer to **time.Time** | Date and time when Marqeta accepted the Card Membership Agreement, in UTC.  Returned if the user accepted their approved application. | [optional] 
**CreatedDate** | **time.Time** | Date and time when the application was created on the Marqeta platform, in UTC. | 
**DecisionModel** | Pointer to [**DecisionsResponse**](DecisionsResponse.md) |  | [optional] 
**DecisionToken** | Pointer to **string** | Unique identifier of the decision made on the application. | [optional] 
**DeviceData** | Pointer to [**DeviceData**](DeviceData.md) |  | [optional] 
**EDisclosureAcceptedAt** | **time.Time** | Date and time when Marqeta accepted the e-Disclosure, in UTC.  Returned if the user accepted their approved application. | 
**ErrorDetails** | Pointer to [**ErrorDetailsResponse**](ErrorDetailsResponse.md) |  | [optional] 
**MetaData** | Pointer to **map[string]interface{}** | Customer-defined additional information about the application. | [optional] 
**MonthlyMortgageOrRent** | Pointer to **float32** | Monthly amount of the mortgage or rent that the user currently pays. | [optional] 
**OfferId** | Pointer to **string** | Unique identifier of the offer for a pre-screened applicant. | [optional] 
**PrequalifiedOfferPreTermsAcceptedAt** | Pointer to **time.Time** | Date and time when Marqeta accepted the Pre-qualified Offer Pre-terms, in UTC.  Returned if the user accepted their approved application. | [optional] 
**PrimaryIncomeSource** | Pointer to **string** | Whether the primary income source comes from the user being employed, unemployed, self-employment, or another situation. | [optional] 
**PrivacyPolicyAcceptedAt** | **time.Time** | Date and time when Marqeta accepted the Privacy Policy, in UTC.  Returned if the user accepted their approved application. | 
**ResidenceType** | Pointer to **string** | Whether the user owns or rents their residence, or has another situation. | [optional] 
**RewardsDisclosurePostTermsAcceptedAt** | Pointer to **time.Time** | Date and time when Marqeta accepted the Rewards Disclosure, in UTC.  Returned if the user accepted their approved application. | [optional] 
**RewardsDisclosurePreTermsAcceptedAt** | **time.Time** | Date and time when Marqeta accepted the Rewards Disclosure, in UTC.  Returned if the user accepted their approved application. | 
**SoctAcceptedAt** | **time.Time** | Date and time when Marqeta accepted the Summary of Credit Terms (SOCT), in UTC.  Returned if the user accepted their approved application. | 
**State** | [**ApplicationResourceState**](ApplicationResourceState.md) |  | 
**TermScheduleInformationAcceptedAt** | Pointer to **time.Time** | Date and time when Marqeta accepted the Terms Schedule, in UTC.  Returned if the user accepted their approved application. | [optional] 
**Token** | **string** | Unique identifier of the application. | 
**TotalAnnualIncome** | Pointer to **float32** | The total amount of the user&#39;s annual income. | [optional] 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 
**UpdatedDate** | **time.Time** | Date and time when the application was last updated on the Marqeta platform, in UTC. | 
**UserToken** | **string** | Unique identifier of the applicant, the user applying for a credit account. | 

## Methods

### NewApplicationsResponse

`func NewApplicationsResponse(bundleToken string, createdDate time.Time, eDisclosureAcceptedAt time.Time, privacyPolicyAcceptedAt time.Time, rewardsDisclosurePreTermsAcceptedAt time.Time, soctAcceptedAt time.Time, state ApplicationResourceState, token string, type_ ApplicationType, updatedDate time.Time, userToken string, ) *ApplicationsResponse`

NewApplicationsResponse instantiates a new ApplicationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsResponseWithDefaults

`func NewApplicationsResponseWithDefaults() *ApplicationsResponse`

NewApplicationsResponseWithDefaults instantiates a new ApplicationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *ApplicationsResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *ApplicationsResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *ApplicationsResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *ApplicationsResponse) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetAnyNonTaxableIncome

`func (o *ApplicationsResponse) GetAnyNonTaxableIncome() bool`

GetAnyNonTaxableIncome returns the AnyNonTaxableIncome field if non-nil, zero value otherwise.

### GetAnyNonTaxableIncomeOk

`func (o *ApplicationsResponse) GetAnyNonTaxableIncomeOk() (*bool, bool)`

GetAnyNonTaxableIncomeOk returns a tuple with the AnyNonTaxableIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyNonTaxableIncome

`func (o *ApplicationsResponse) SetAnyNonTaxableIncome(v bool)`

SetAnyNonTaxableIncome sets AnyNonTaxableIncome field to given value.

### HasAnyNonTaxableIncome

`func (o *ApplicationsResponse) HasAnyNonTaxableIncome() bool`

HasAnyNonTaxableIncome returns a boolean if a field has been set.

### GetApplicationAcceptedAt

`func (o *ApplicationsResponse) GetApplicationAcceptedAt() time.Time`

GetApplicationAcceptedAt returns the ApplicationAcceptedAt field if non-nil, zero value otherwise.

### GetApplicationAcceptedAtOk

`func (o *ApplicationsResponse) GetApplicationAcceptedAtOk() (*time.Time, bool)`

GetApplicationAcceptedAtOk returns a tuple with the ApplicationAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAcceptedAt

`func (o *ApplicationsResponse) SetApplicationAcceptedAt(v time.Time)`

SetApplicationAcceptedAt sets ApplicationAcceptedAt field to given value.

### HasApplicationAcceptedAt

`func (o *ApplicationsResponse) HasApplicationAcceptedAt() bool`

HasApplicationAcceptedAt returns a boolean if a field has been set.

### GetApplicationCanceledAt

`func (o *ApplicationsResponse) GetApplicationCanceledAt() time.Time`

GetApplicationCanceledAt returns the ApplicationCanceledAt field if non-nil, zero value otherwise.

### GetApplicationCanceledAtOk

`func (o *ApplicationsResponse) GetApplicationCanceledAtOk() (*time.Time, bool)`

GetApplicationCanceledAtOk returns a tuple with the ApplicationCanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCanceledAt

`func (o *ApplicationsResponse) SetApplicationCanceledAt(v time.Time)`

SetApplicationCanceledAt sets ApplicationCanceledAt field to given value.

### HasApplicationCanceledAt

`func (o *ApplicationsResponse) HasApplicationCanceledAt() bool`

HasApplicationCanceledAt returns a boolean if a field has been set.

### GetBenefitsDisclosureAcceptedAt

`func (o *ApplicationsResponse) GetBenefitsDisclosureAcceptedAt() time.Time`

GetBenefitsDisclosureAcceptedAt returns the BenefitsDisclosureAcceptedAt field if non-nil, zero value otherwise.

### GetBenefitsDisclosureAcceptedAtOk

`func (o *ApplicationsResponse) GetBenefitsDisclosureAcceptedAtOk() (*time.Time, bool)`

GetBenefitsDisclosureAcceptedAtOk returns a tuple with the BenefitsDisclosureAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitsDisclosureAcceptedAt

`func (o *ApplicationsResponse) SetBenefitsDisclosureAcceptedAt(v time.Time)`

SetBenefitsDisclosureAcceptedAt sets BenefitsDisclosureAcceptedAt field to given value.

### HasBenefitsDisclosureAcceptedAt

`func (o *ApplicationsResponse) HasBenefitsDisclosureAcceptedAt() bool`

HasBenefitsDisclosureAcceptedAt returns a boolean if a field has been set.

### GetBundleToken

`func (o *ApplicationsResponse) GetBundleToken() string`

GetBundleToken returns the BundleToken field if non-nil, zero value otherwise.

### GetBundleTokenOk

`func (o *ApplicationsResponse) GetBundleTokenOk() (*string, bool)`

GetBundleTokenOk returns a tuple with the BundleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleToken

`func (o *ApplicationsResponse) SetBundleToken(v string)`

SetBundleToken sets BundleToken field to given value.


### GetCardMembershipAgreementAcceptedAt

`func (o *ApplicationsResponse) GetCardMembershipAgreementAcceptedAt() time.Time`

GetCardMembershipAgreementAcceptedAt returns the CardMembershipAgreementAcceptedAt field if non-nil, zero value otherwise.

### GetCardMembershipAgreementAcceptedAtOk

`func (o *ApplicationsResponse) GetCardMembershipAgreementAcceptedAtOk() (*time.Time, bool)`

GetCardMembershipAgreementAcceptedAtOk returns a tuple with the CardMembershipAgreementAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardMembershipAgreementAcceptedAt

`func (o *ApplicationsResponse) SetCardMembershipAgreementAcceptedAt(v time.Time)`

SetCardMembershipAgreementAcceptedAt sets CardMembershipAgreementAcceptedAt field to given value.

### HasCardMembershipAgreementAcceptedAt

`func (o *ApplicationsResponse) HasCardMembershipAgreementAcceptedAt() bool`

HasCardMembershipAgreementAcceptedAt returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ApplicationsResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApplicationsResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApplicationsResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetDecisionModel

`func (o *ApplicationsResponse) GetDecisionModel() DecisionsResponse`

GetDecisionModel returns the DecisionModel field if non-nil, zero value otherwise.

### GetDecisionModelOk

`func (o *ApplicationsResponse) GetDecisionModelOk() (*DecisionsResponse, bool)`

GetDecisionModelOk returns a tuple with the DecisionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionModel

`func (o *ApplicationsResponse) SetDecisionModel(v DecisionsResponse)`

SetDecisionModel sets DecisionModel field to given value.

### HasDecisionModel

`func (o *ApplicationsResponse) HasDecisionModel() bool`

HasDecisionModel returns a boolean if a field has been set.

### GetDecisionToken

`func (o *ApplicationsResponse) GetDecisionToken() string`

GetDecisionToken returns the DecisionToken field if non-nil, zero value otherwise.

### GetDecisionTokenOk

`func (o *ApplicationsResponse) GetDecisionTokenOk() (*string, bool)`

GetDecisionTokenOk returns a tuple with the DecisionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionToken

`func (o *ApplicationsResponse) SetDecisionToken(v string)`

SetDecisionToken sets DecisionToken field to given value.

### HasDecisionToken

`func (o *ApplicationsResponse) HasDecisionToken() bool`

HasDecisionToken returns a boolean if a field has been set.

### GetDeviceData

`func (o *ApplicationsResponse) GetDeviceData() DeviceData`

GetDeviceData returns the DeviceData field if non-nil, zero value otherwise.

### GetDeviceDataOk

`func (o *ApplicationsResponse) GetDeviceDataOk() (*DeviceData, bool)`

GetDeviceDataOk returns a tuple with the DeviceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceData

`func (o *ApplicationsResponse) SetDeviceData(v DeviceData)`

SetDeviceData sets DeviceData field to given value.

### HasDeviceData

`func (o *ApplicationsResponse) HasDeviceData() bool`

HasDeviceData returns a boolean if a field has been set.

### GetEDisclosureAcceptedAt

`func (o *ApplicationsResponse) GetEDisclosureAcceptedAt() time.Time`

GetEDisclosureAcceptedAt returns the EDisclosureAcceptedAt field if non-nil, zero value otherwise.

### GetEDisclosureAcceptedAtOk

`func (o *ApplicationsResponse) GetEDisclosureAcceptedAtOk() (*time.Time, bool)`

GetEDisclosureAcceptedAtOk returns a tuple with the EDisclosureAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDisclosureAcceptedAt

`func (o *ApplicationsResponse) SetEDisclosureAcceptedAt(v time.Time)`

SetEDisclosureAcceptedAt sets EDisclosureAcceptedAt field to given value.


### GetErrorDetails

`func (o *ApplicationsResponse) GetErrorDetails() ErrorDetailsResponse`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ApplicationsResponse) GetErrorDetailsOk() (*ErrorDetailsResponse, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ApplicationsResponse) SetErrorDetails(v ErrorDetailsResponse)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ApplicationsResponse) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetMetaData

`func (o *ApplicationsResponse) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ApplicationsResponse) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ApplicationsResponse) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ApplicationsResponse) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMonthlyMortgageOrRent

`func (o *ApplicationsResponse) GetMonthlyMortgageOrRent() float32`

GetMonthlyMortgageOrRent returns the MonthlyMortgageOrRent field if non-nil, zero value otherwise.

### GetMonthlyMortgageOrRentOk

`func (o *ApplicationsResponse) GetMonthlyMortgageOrRentOk() (*float32, bool)`

GetMonthlyMortgageOrRentOk returns a tuple with the MonthlyMortgageOrRent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyMortgageOrRent

`func (o *ApplicationsResponse) SetMonthlyMortgageOrRent(v float32)`

SetMonthlyMortgageOrRent sets MonthlyMortgageOrRent field to given value.

### HasMonthlyMortgageOrRent

`func (o *ApplicationsResponse) HasMonthlyMortgageOrRent() bool`

HasMonthlyMortgageOrRent returns a boolean if a field has been set.

### GetOfferId

`func (o *ApplicationsResponse) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *ApplicationsResponse) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *ApplicationsResponse) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *ApplicationsResponse) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetPrequalifiedOfferPreTermsAcceptedAt

`func (o *ApplicationsResponse) GetPrequalifiedOfferPreTermsAcceptedAt() time.Time`

GetPrequalifiedOfferPreTermsAcceptedAt returns the PrequalifiedOfferPreTermsAcceptedAt field if non-nil, zero value otherwise.

### GetPrequalifiedOfferPreTermsAcceptedAtOk

`func (o *ApplicationsResponse) GetPrequalifiedOfferPreTermsAcceptedAtOk() (*time.Time, bool)`

GetPrequalifiedOfferPreTermsAcceptedAtOk returns a tuple with the PrequalifiedOfferPreTermsAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrequalifiedOfferPreTermsAcceptedAt

`func (o *ApplicationsResponse) SetPrequalifiedOfferPreTermsAcceptedAt(v time.Time)`

SetPrequalifiedOfferPreTermsAcceptedAt sets PrequalifiedOfferPreTermsAcceptedAt field to given value.

### HasPrequalifiedOfferPreTermsAcceptedAt

`func (o *ApplicationsResponse) HasPrequalifiedOfferPreTermsAcceptedAt() bool`

HasPrequalifiedOfferPreTermsAcceptedAt returns a boolean if a field has been set.

### GetPrimaryIncomeSource

`func (o *ApplicationsResponse) GetPrimaryIncomeSource() string`

GetPrimaryIncomeSource returns the PrimaryIncomeSource field if non-nil, zero value otherwise.

### GetPrimaryIncomeSourceOk

`func (o *ApplicationsResponse) GetPrimaryIncomeSourceOk() (*string, bool)`

GetPrimaryIncomeSourceOk returns a tuple with the PrimaryIncomeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIncomeSource

`func (o *ApplicationsResponse) SetPrimaryIncomeSource(v string)`

SetPrimaryIncomeSource sets PrimaryIncomeSource field to given value.

### HasPrimaryIncomeSource

`func (o *ApplicationsResponse) HasPrimaryIncomeSource() bool`

HasPrimaryIncomeSource returns a boolean if a field has been set.

### GetPrivacyPolicyAcceptedAt

`func (o *ApplicationsResponse) GetPrivacyPolicyAcceptedAt() time.Time`

GetPrivacyPolicyAcceptedAt returns the PrivacyPolicyAcceptedAt field if non-nil, zero value otherwise.

### GetPrivacyPolicyAcceptedAtOk

`func (o *ApplicationsResponse) GetPrivacyPolicyAcceptedAtOk() (*time.Time, bool)`

GetPrivacyPolicyAcceptedAtOk returns a tuple with the PrivacyPolicyAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyAcceptedAt

`func (o *ApplicationsResponse) SetPrivacyPolicyAcceptedAt(v time.Time)`

SetPrivacyPolicyAcceptedAt sets PrivacyPolicyAcceptedAt field to given value.


### GetResidenceType

`func (o *ApplicationsResponse) GetResidenceType() string`

GetResidenceType returns the ResidenceType field if non-nil, zero value otherwise.

### GetResidenceTypeOk

`func (o *ApplicationsResponse) GetResidenceTypeOk() (*string, bool)`

GetResidenceTypeOk returns a tuple with the ResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceType

`func (o *ApplicationsResponse) SetResidenceType(v string)`

SetResidenceType sets ResidenceType field to given value.

### HasResidenceType

`func (o *ApplicationsResponse) HasResidenceType() bool`

HasResidenceType returns a boolean if a field has been set.

### GetRewardsDisclosurePostTermsAcceptedAt

`func (o *ApplicationsResponse) GetRewardsDisclosurePostTermsAcceptedAt() time.Time`

GetRewardsDisclosurePostTermsAcceptedAt returns the RewardsDisclosurePostTermsAcceptedAt field if non-nil, zero value otherwise.

### GetRewardsDisclosurePostTermsAcceptedAtOk

`func (o *ApplicationsResponse) GetRewardsDisclosurePostTermsAcceptedAtOk() (*time.Time, bool)`

GetRewardsDisclosurePostTermsAcceptedAtOk returns a tuple with the RewardsDisclosurePostTermsAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsDisclosurePostTermsAcceptedAt

`func (o *ApplicationsResponse) SetRewardsDisclosurePostTermsAcceptedAt(v time.Time)`

SetRewardsDisclosurePostTermsAcceptedAt sets RewardsDisclosurePostTermsAcceptedAt field to given value.

### HasRewardsDisclosurePostTermsAcceptedAt

`func (o *ApplicationsResponse) HasRewardsDisclosurePostTermsAcceptedAt() bool`

HasRewardsDisclosurePostTermsAcceptedAt returns a boolean if a field has been set.

### GetRewardsDisclosurePreTermsAcceptedAt

`func (o *ApplicationsResponse) GetRewardsDisclosurePreTermsAcceptedAt() time.Time`

GetRewardsDisclosurePreTermsAcceptedAt returns the RewardsDisclosurePreTermsAcceptedAt field if non-nil, zero value otherwise.

### GetRewardsDisclosurePreTermsAcceptedAtOk

`func (o *ApplicationsResponse) GetRewardsDisclosurePreTermsAcceptedAtOk() (*time.Time, bool)`

GetRewardsDisclosurePreTermsAcceptedAtOk returns a tuple with the RewardsDisclosurePreTermsAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsDisclosurePreTermsAcceptedAt

`func (o *ApplicationsResponse) SetRewardsDisclosurePreTermsAcceptedAt(v time.Time)`

SetRewardsDisclosurePreTermsAcceptedAt sets RewardsDisclosurePreTermsAcceptedAt field to given value.


### GetSoctAcceptedAt

`func (o *ApplicationsResponse) GetSoctAcceptedAt() time.Time`

GetSoctAcceptedAt returns the SoctAcceptedAt field if non-nil, zero value otherwise.

### GetSoctAcceptedAtOk

`func (o *ApplicationsResponse) GetSoctAcceptedAtOk() (*time.Time, bool)`

GetSoctAcceptedAtOk returns a tuple with the SoctAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoctAcceptedAt

`func (o *ApplicationsResponse) SetSoctAcceptedAt(v time.Time)`

SetSoctAcceptedAt sets SoctAcceptedAt field to given value.


### GetState

`func (o *ApplicationsResponse) GetState() ApplicationResourceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationsResponse) GetStateOk() (*ApplicationResourceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApplicationsResponse) SetState(v ApplicationResourceState)`

SetState sets State field to given value.


### GetTermScheduleInformationAcceptedAt

`func (o *ApplicationsResponse) GetTermScheduleInformationAcceptedAt() time.Time`

GetTermScheduleInformationAcceptedAt returns the TermScheduleInformationAcceptedAt field if non-nil, zero value otherwise.

### GetTermScheduleInformationAcceptedAtOk

`func (o *ApplicationsResponse) GetTermScheduleInformationAcceptedAtOk() (*time.Time, bool)`

GetTermScheduleInformationAcceptedAtOk returns a tuple with the TermScheduleInformationAcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermScheduleInformationAcceptedAt

`func (o *ApplicationsResponse) SetTermScheduleInformationAcceptedAt(v time.Time)`

SetTermScheduleInformationAcceptedAt sets TermScheduleInformationAcceptedAt field to given value.

### HasTermScheduleInformationAcceptedAt

`func (o *ApplicationsResponse) HasTermScheduleInformationAcceptedAt() bool`

HasTermScheduleInformationAcceptedAt returns a boolean if a field has been set.

### GetToken

`func (o *ApplicationsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApplicationsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApplicationsResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTotalAnnualIncome

`func (o *ApplicationsResponse) GetTotalAnnualIncome() float32`

GetTotalAnnualIncome returns the TotalAnnualIncome field if non-nil, zero value otherwise.

### GetTotalAnnualIncomeOk

`func (o *ApplicationsResponse) GetTotalAnnualIncomeOk() (*float32, bool)`

GetTotalAnnualIncomeOk returns a tuple with the TotalAnnualIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnnualIncome

`func (o *ApplicationsResponse) SetTotalAnnualIncome(v float32)`

SetTotalAnnualIncome sets TotalAnnualIncome field to given value.

### HasTotalAnnualIncome

`func (o *ApplicationsResponse) HasTotalAnnualIncome() bool`

HasTotalAnnualIncome returns a boolean if a field has been set.

### GetType

`func (o *ApplicationsResponse) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationsResponse) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationsResponse) SetType(v ApplicationType)`

SetType sets Type field to given value.


### GetUpdatedDate

`func (o *ApplicationsResponse) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ApplicationsResponse) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ApplicationsResponse) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.


### GetUserToken

`func (o *ApplicationsResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *ApplicationsResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *ApplicationsResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


