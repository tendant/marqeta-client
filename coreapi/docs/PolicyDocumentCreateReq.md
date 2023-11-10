# PolicyDocumentCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountStatement** | [**PolicyDocumentTemplateReq**](PolicyDocumentTemplateReq.md) |  | 
**BenefitsDisclosurePremium** | [**PolicyDocumentAssetReq**](PolicyDocumentAssetReq.md) |  | 
**BenefitsDisclosureTraditional** | [**PolicyDocumentAssetReq**](PolicyDocumentAssetReq.md) |  | 
**CardMemberAgreement** | [**PolicyDocumentAssetReq**](PolicyDocumentAssetReq.md) |  | 
**EDisclosure** | [**PolicyDocumentAssetReq**](PolicyDocumentAssetReq.md) |  | 
**Name** | **string** | Name of the document policy. | 
**NoaaMultipleReasonWithDoddFrank** | [**PolicyDocumentTemplateReq**](PolicyDocumentTemplateReq.md) |  | 
**NoaaSingleReason** | [**PolicyDocumentTemplateReq**](PolicyDocumentTemplateReq.md) |  | 
**NoaaSingleReasonWithDoddFrank** | [**PolicyDocumentTemplateReq**](PolicyDocumentTemplateReq.md) |  | 
**PrivacyPolicy** | [**PolicyDocumentAssetReq**](PolicyDocumentAssetReq.md) |  | 
**RewardsDisclosure** | Pointer to [**PolicyDocumentAssetAndTemplateReq**](PolicyDocumentAssetAndTemplateReq.md) |  | [optional] 
**SummaryOfCreditTerms** | [**PolicyDocumentAssetAndTemplateReq**](PolicyDocumentAssetAndTemplateReq.md) |  | 
**TermsSchedule** | [**PolicyDocumentTemplateReq**](PolicyDocumentTemplateReq.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the document policy. | [optional] 

## Methods

### NewPolicyDocumentCreateReq

`func NewPolicyDocumentCreateReq(accountStatement PolicyDocumentTemplateReq, benefitsDisclosurePremium PolicyDocumentAssetReq, benefitsDisclosureTraditional PolicyDocumentAssetReq, cardMemberAgreement PolicyDocumentAssetReq, eDisclosure PolicyDocumentAssetReq, name string, noaaMultipleReasonWithDoddFrank PolicyDocumentTemplateReq, noaaSingleReason PolicyDocumentTemplateReq, noaaSingleReasonWithDoddFrank PolicyDocumentTemplateReq, privacyPolicy PolicyDocumentAssetReq, summaryOfCreditTerms PolicyDocumentAssetAndTemplateReq, termsSchedule PolicyDocumentTemplateReq, ) *PolicyDocumentCreateReq`

NewPolicyDocumentCreateReq instantiates a new PolicyDocumentCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentCreateReqWithDefaults

`func NewPolicyDocumentCreateReqWithDefaults() *PolicyDocumentCreateReq`

NewPolicyDocumentCreateReqWithDefaults instantiates a new PolicyDocumentCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountStatement

`func (o *PolicyDocumentCreateReq) GetAccountStatement() PolicyDocumentTemplateReq`

GetAccountStatement returns the AccountStatement field if non-nil, zero value otherwise.

### GetAccountStatementOk

`func (o *PolicyDocumentCreateReq) GetAccountStatementOk() (*PolicyDocumentTemplateReq, bool)`

GetAccountStatementOk returns a tuple with the AccountStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatement

`func (o *PolicyDocumentCreateReq) SetAccountStatement(v PolicyDocumentTemplateReq)`

SetAccountStatement sets AccountStatement field to given value.


### GetBenefitsDisclosurePremium

`func (o *PolicyDocumentCreateReq) GetBenefitsDisclosurePremium() PolicyDocumentAssetReq`

GetBenefitsDisclosurePremium returns the BenefitsDisclosurePremium field if non-nil, zero value otherwise.

### GetBenefitsDisclosurePremiumOk

`func (o *PolicyDocumentCreateReq) GetBenefitsDisclosurePremiumOk() (*PolicyDocumentAssetReq, bool)`

GetBenefitsDisclosurePremiumOk returns a tuple with the BenefitsDisclosurePremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitsDisclosurePremium

`func (o *PolicyDocumentCreateReq) SetBenefitsDisclosurePremium(v PolicyDocumentAssetReq)`

SetBenefitsDisclosurePremium sets BenefitsDisclosurePremium field to given value.


### GetBenefitsDisclosureTraditional

`func (o *PolicyDocumentCreateReq) GetBenefitsDisclosureTraditional() PolicyDocumentAssetReq`

GetBenefitsDisclosureTraditional returns the BenefitsDisclosureTraditional field if non-nil, zero value otherwise.

### GetBenefitsDisclosureTraditionalOk

`func (o *PolicyDocumentCreateReq) GetBenefitsDisclosureTraditionalOk() (*PolicyDocumentAssetReq, bool)`

GetBenefitsDisclosureTraditionalOk returns a tuple with the BenefitsDisclosureTraditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitsDisclosureTraditional

`func (o *PolicyDocumentCreateReq) SetBenefitsDisclosureTraditional(v PolicyDocumentAssetReq)`

SetBenefitsDisclosureTraditional sets BenefitsDisclosureTraditional field to given value.


### GetCardMemberAgreement

`func (o *PolicyDocumentCreateReq) GetCardMemberAgreement() PolicyDocumentAssetReq`

GetCardMemberAgreement returns the CardMemberAgreement field if non-nil, zero value otherwise.

### GetCardMemberAgreementOk

`func (o *PolicyDocumentCreateReq) GetCardMemberAgreementOk() (*PolicyDocumentAssetReq, bool)`

GetCardMemberAgreementOk returns a tuple with the CardMemberAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardMemberAgreement

`func (o *PolicyDocumentCreateReq) SetCardMemberAgreement(v PolicyDocumentAssetReq)`

SetCardMemberAgreement sets CardMemberAgreement field to given value.


### GetEDisclosure

`func (o *PolicyDocumentCreateReq) GetEDisclosure() PolicyDocumentAssetReq`

GetEDisclosure returns the EDisclosure field if non-nil, zero value otherwise.

### GetEDisclosureOk

`func (o *PolicyDocumentCreateReq) GetEDisclosureOk() (*PolicyDocumentAssetReq, bool)`

GetEDisclosureOk returns a tuple with the EDisclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDisclosure

`func (o *PolicyDocumentCreateReq) SetEDisclosure(v PolicyDocumentAssetReq)`

SetEDisclosure sets EDisclosure field to given value.


### GetName

`func (o *PolicyDocumentCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyDocumentCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyDocumentCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetNoaaMultipleReasonWithDoddFrank

`func (o *PolicyDocumentCreateReq) GetNoaaMultipleReasonWithDoddFrank() PolicyDocumentTemplateReq`

GetNoaaMultipleReasonWithDoddFrank returns the NoaaMultipleReasonWithDoddFrank field if non-nil, zero value otherwise.

### GetNoaaMultipleReasonWithDoddFrankOk

`func (o *PolicyDocumentCreateReq) GetNoaaMultipleReasonWithDoddFrankOk() (*PolicyDocumentTemplateReq, bool)`

GetNoaaMultipleReasonWithDoddFrankOk returns a tuple with the NoaaMultipleReasonWithDoddFrank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoaaMultipleReasonWithDoddFrank

`func (o *PolicyDocumentCreateReq) SetNoaaMultipleReasonWithDoddFrank(v PolicyDocumentTemplateReq)`

SetNoaaMultipleReasonWithDoddFrank sets NoaaMultipleReasonWithDoddFrank field to given value.


### GetNoaaSingleReason

`func (o *PolicyDocumentCreateReq) GetNoaaSingleReason() PolicyDocumentTemplateReq`

GetNoaaSingleReason returns the NoaaSingleReason field if non-nil, zero value otherwise.

### GetNoaaSingleReasonOk

`func (o *PolicyDocumentCreateReq) GetNoaaSingleReasonOk() (*PolicyDocumentTemplateReq, bool)`

GetNoaaSingleReasonOk returns a tuple with the NoaaSingleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoaaSingleReason

`func (o *PolicyDocumentCreateReq) SetNoaaSingleReason(v PolicyDocumentTemplateReq)`

SetNoaaSingleReason sets NoaaSingleReason field to given value.


### GetNoaaSingleReasonWithDoddFrank

`func (o *PolicyDocumentCreateReq) GetNoaaSingleReasonWithDoddFrank() PolicyDocumentTemplateReq`

GetNoaaSingleReasonWithDoddFrank returns the NoaaSingleReasonWithDoddFrank field if non-nil, zero value otherwise.

### GetNoaaSingleReasonWithDoddFrankOk

`func (o *PolicyDocumentCreateReq) GetNoaaSingleReasonWithDoddFrankOk() (*PolicyDocumentTemplateReq, bool)`

GetNoaaSingleReasonWithDoddFrankOk returns a tuple with the NoaaSingleReasonWithDoddFrank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoaaSingleReasonWithDoddFrank

`func (o *PolicyDocumentCreateReq) SetNoaaSingleReasonWithDoddFrank(v PolicyDocumentTemplateReq)`

SetNoaaSingleReasonWithDoddFrank sets NoaaSingleReasonWithDoddFrank field to given value.


### GetPrivacyPolicy

`func (o *PolicyDocumentCreateReq) GetPrivacyPolicy() PolicyDocumentAssetReq`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *PolicyDocumentCreateReq) GetPrivacyPolicyOk() (*PolicyDocumentAssetReq, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *PolicyDocumentCreateReq) SetPrivacyPolicy(v PolicyDocumentAssetReq)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.


### GetRewardsDisclosure

`func (o *PolicyDocumentCreateReq) GetRewardsDisclosure() PolicyDocumentAssetAndTemplateReq`

GetRewardsDisclosure returns the RewardsDisclosure field if non-nil, zero value otherwise.

### GetRewardsDisclosureOk

`func (o *PolicyDocumentCreateReq) GetRewardsDisclosureOk() (*PolicyDocumentAssetAndTemplateReq, bool)`

GetRewardsDisclosureOk returns a tuple with the RewardsDisclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsDisclosure

`func (o *PolicyDocumentCreateReq) SetRewardsDisclosure(v PolicyDocumentAssetAndTemplateReq)`

SetRewardsDisclosure sets RewardsDisclosure field to given value.

### HasRewardsDisclosure

`func (o *PolicyDocumentCreateReq) HasRewardsDisclosure() bool`

HasRewardsDisclosure returns a boolean if a field has been set.

### GetSummaryOfCreditTerms

`func (o *PolicyDocumentCreateReq) GetSummaryOfCreditTerms() PolicyDocumentAssetAndTemplateReq`

GetSummaryOfCreditTerms returns the SummaryOfCreditTerms field if non-nil, zero value otherwise.

### GetSummaryOfCreditTermsOk

`func (o *PolicyDocumentCreateReq) GetSummaryOfCreditTermsOk() (*PolicyDocumentAssetAndTemplateReq, bool)`

GetSummaryOfCreditTermsOk returns a tuple with the SummaryOfCreditTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryOfCreditTerms

`func (o *PolicyDocumentCreateReq) SetSummaryOfCreditTerms(v PolicyDocumentAssetAndTemplateReq)`

SetSummaryOfCreditTerms sets SummaryOfCreditTerms field to given value.


### GetTermsSchedule

`func (o *PolicyDocumentCreateReq) GetTermsSchedule() PolicyDocumentTemplateReq`

GetTermsSchedule returns the TermsSchedule field if non-nil, zero value otherwise.

### GetTermsScheduleOk

`func (o *PolicyDocumentCreateReq) GetTermsScheduleOk() (*PolicyDocumentTemplateReq, bool)`

GetTermsScheduleOk returns a tuple with the TermsSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsSchedule

`func (o *PolicyDocumentCreateReq) SetTermsSchedule(v PolicyDocumentTemplateReq)`

SetTermsSchedule sets TermsSchedule field to given value.


### GetToken

`func (o *PolicyDocumentCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyDocumentCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyDocumentCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyDocumentCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


