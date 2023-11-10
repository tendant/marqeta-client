# PolicyDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountStatement** | Pointer to [**PolicyDocumentTemplateResponse**](PolicyDocumentTemplateResponse.md) |  | [optional] 
**BenefitsDisclosurePremium** | Pointer to [**PolicyDocumentAssetResponse**](PolicyDocumentAssetResponse.md) |  | [optional] 
**BenefitsDisclosureTraditional** | Pointer to [**PolicyDocumentAssetResponse**](PolicyDocumentAssetResponse.md) |  | [optional] 
**CardMemberAgreement** | Pointer to [**PolicyDocumentAssetResponse**](PolicyDocumentAssetResponse.md) |  | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the document policy was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**EDisclosure** | Pointer to [**PolicyDocumentAssetResponse**](PolicyDocumentAssetResponse.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the document policy. | [optional] 
**NoaaMultipleReasonWithDoddFrank** | Pointer to [**PolicyDocumentTemplateResponse**](PolicyDocumentTemplateResponse.md) |  | [optional] 
**NoaaSingleReason** | Pointer to [**PolicyDocumentTemplateResponse**](PolicyDocumentTemplateResponse.md) |  | [optional] 
**NoaaSingleReasonWithDoddFrank** | Pointer to [**PolicyDocumentTemplateResponse**](PolicyDocumentTemplateResponse.md) |  | [optional] 
**PrivacyPolicy** | Pointer to [**PolicyDocumentAssetResponse**](PolicyDocumentAssetResponse.md) |  | [optional] 
**RewardsDisclosure** | Pointer to [**PolicyDocumentAssetAndTemplateResponse**](PolicyDocumentAssetAndTemplateResponse.md) |  | [optional] 
**SummaryOfCreditTerms** | Pointer to [**PolicyDocumentAssetAndTemplateResponse**](PolicyDocumentAssetAndTemplateResponse.md) |  | [optional] 
**TermsSchedule** | Pointer to [**PolicyDocumentTemplateResponse**](PolicyDocumentTemplateResponse.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the document policy. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the document policy was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewPolicyDocumentResponse

`func NewPolicyDocumentResponse() *PolicyDocumentResponse`

NewPolicyDocumentResponse instantiates a new PolicyDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentResponseWithDefaults

`func NewPolicyDocumentResponseWithDefaults() *PolicyDocumentResponse`

NewPolicyDocumentResponseWithDefaults instantiates a new PolicyDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountStatement

`func (o *PolicyDocumentResponse) GetAccountStatement() PolicyDocumentTemplateResponse`

GetAccountStatement returns the AccountStatement field if non-nil, zero value otherwise.

### GetAccountStatementOk

`func (o *PolicyDocumentResponse) GetAccountStatementOk() (*PolicyDocumentTemplateResponse, bool)`

GetAccountStatementOk returns a tuple with the AccountStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatement

`func (o *PolicyDocumentResponse) SetAccountStatement(v PolicyDocumentTemplateResponse)`

SetAccountStatement sets AccountStatement field to given value.

### HasAccountStatement

`func (o *PolicyDocumentResponse) HasAccountStatement() bool`

HasAccountStatement returns a boolean if a field has been set.

### GetBenefitsDisclosurePremium

`func (o *PolicyDocumentResponse) GetBenefitsDisclosurePremium() PolicyDocumentAssetResponse`

GetBenefitsDisclosurePremium returns the BenefitsDisclosurePremium field if non-nil, zero value otherwise.

### GetBenefitsDisclosurePremiumOk

`func (o *PolicyDocumentResponse) GetBenefitsDisclosurePremiumOk() (*PolicyDocumentAssetResponse, bool)`

GetBenefitsDisclosurePremiumOk returns a tuple with the BenefitsDisclosurePremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitsDisclosurePremium

`func (o *PolicyDocumentResponse) SetBenefitsDisclosurePremium(v PolicyDocumentAssetResponse)`

SetBenefitsDisclosurePremium sets BenefitsDisclosurePremium field to given value.

### HasBenefitsDisclosurePremium

`func (o *PolicyDocumentResponse) HasBenefitsDisclosurePremium() bool`

HasBenefitsDisclosurePremium returns a boolean if a field has been set.

### GetBenefitsDisclosureTraditional

`func (o *PolicyDocumentResponse) GetBenefitsDisclosureTraditional() PolicyDocumentAssetResponse`

GetBenefitsDisclosureTraditional returns the BenefitsDisclosureTraditional field if non-nil, zero value otherwise.

### GetBenefitsDisclosureTraditionalOk

`func (o *PolicyDocumentResponse) GetBenefitsDisclosureTraditionalOk() (*PolicyDocumentAssetResponse, bool)`

GetBenefitsDisclosureTraditionalOk returns a tuple with the BenefitsDisclosureTraditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitsDisclosureTraditional

`func (o *PolicyDocumentResponse) SetBenefitsDisclosureTraditional(v PolicyDocumentAssetResponse)`

SetBenefitsDisclosureTraditional sets BenefitsDisclosureTraditional field to given value.

### HasBenefitsDisclosureTraditional

`func (o *PolicyDocumentResponse) HasBenefitsDisclosureTraditional() bool`

HasBenefitsDisclosureTraditional returns a boolean if a field has been set.

### GetCardMemberAgreement

`func (o *PolicyDocumentResponse) GetCardMemberAgreement() PolicyDocumentAssetResponse`

GetCardMemberAgreement returns the CardMemberAgreement field if non-nil, zero value otherwise.

### GetCardMemberAgreementOk

`func (o *PolicyDocumentResponse) GetCardMemberAgreementOk() (*PolicyDocumentAssetResponse, bool)`

GetCardMemberAgreementOk returns a tuple with the CardMemberAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardMemberAgreement

`func (o *PolicyDocumentResponse) SetCardMemberAgreement(v PolicyDocumentAssetResponse)`

SetCardMemberAgreement sets CardMemberAgreement field to given value.

### HasCardMemberAgreement

`func (o *PolicyDocumentResponse) HasCardMemberAgreement() bool`

HasCardMemberAgreement returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PolicyDocumentResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PolicyDocumentResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PolicyDocumentResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PolicyDocumentResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEDisclosure

`func (o *PolicyDocumentResponse) GetEDisclosure() PolicyDocumentAssetResponse`

GetEDisclosure returns the EDisclosure field if non-nil, zero value otherwise.

### GetEDisclosureOk

`func (o *PolicyDocumentResponse) GetEDisclosureOk() (*PolicyDocumentAssetResponse, bool)`

GetEDisclosureOk returns a tuple with the EDisclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDisclosure

`func (o *PolicyDocumentResponse) SetEDisclosure(v PolicyDocumentAssetResponse)`

SetEDisclosure sets EDisclosure field to given value.

### HasEDisclosure

`func (o *PolicyDocumentResponse) HasEDisclosure() bool`

HasEDisclosure returns a boolean if a field has been set.

### GetName

`func (o *PolicyDocumentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyDocumentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyDocumentResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyDocumentResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNoaaMultipleReasonWithDoddFrank

`func (o *PolicyDocumentResponse) GetNoaaMultipleReasonWithDoddFrank() PolicyDocumentTemplateResponse`

GetNoaaMultipleReasonWithDoddFrank returns the NoaaMultipleReasonWithDoddFrank field if non-nil, zero value otherwise.

### GetNoaaMultipleReasonWithDoddFrankOk

`func (o *PolicyDocumentResponse) GetNoaaMultipleReasonWithDoddFrankOk() (*PolicyDocumentTemplateResponse, bool)`

GetNoaaMultipleReasonWithDoddFrankOk returns a tuple with the NoaaMultipleReasonWithDoddFrank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoaaMultipleReasonWithDoddFrank

`func (o *PolicyDocumentResponse) SetNoaaMultipleReasonWithDoddFrank(v PolicyDocumentTemplateResponse)`

SetNoaaMultipleReasonWithDoddFrank sets NoaaMultipleReasonWithDoddFrank field to given value.

### HasNoaaMultipleReasonWithDoddFrank

`func (o *PolicyDocumentResponse) HasNoaaMultipleReasonWithDoddFrank() bool`

HasNoaaMultipleReasonWithDoddFrank returns a boolean if a field has been set.

### GetNoaaSingleReason

`func (o *PolicyDocumentResponse) GetNoaaSingleReason() PolicyDocumentTemplateResponse`

GetNoaaSingleReason returns the NoaaSingleReason field if non-nil, zero value otherwise.

### GetNoaaSingleReasonOk

`func (o *PolicyDocumentResponse) GetNoaaSingleReasonOk() (*PolicyDocumentTemplateResponse, bool)`

GetNoaaSingleReasonOk returns a tuple with the NoaaSingleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoaaSingleReason

`func (o *PolicyDocumentResponse) SetNoaaSingleReason(v PolicyDocumentTemplateResponse)`

SetNoaaSingleReason sets NoaaSingleReason field to given value.

### HasNoaaSingleReason

`func (o *PolicyDocumentResponse) HasNoaaSingleReason() bool`

HasNoaaSingleReason returns a boolean if a field has been set.

### GetNoaaSingleReasonWithDoddFrank

`func (o *PolicyDocumentResponse) GetNoaaSingleReasonWithDoddFrank() PolicyDocumentTemplateResponse`

GetNoaaSingleReasonWithDoddFrank returns the NoaaSingleReasonWithDoddFrank field if non-nil, zero value otherwise.

### GetNoaaSingleReasonWithDoddFrankOk

`func (o *PolicyDocumentResponse) GetNoaaSingleReasonWithDoddFrankOk() (*PolicyDocumentTemplateResponse, bool)`

GetNoaaSingleReasonWithDoddFrankOk returns a tuple with the NoaaSingleReasonWithDoddFrank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoaaSingleReasonWithDoddFrank

`func (o *PolicyDocumentResponse) SetNoaaSingleReasonWithDoddFrank(v PolicyDocumentTemplateResponse)`

SetNoaaSingleReasonWithDoddFrank sets NoaaSingleReasonWithDoddFrank field to given value.

### HasNoaaSingleReasonWithDoddFrank

`func (o *PolicyDocumentResponse) HasNoaaSingleReasonWithDoddFrank() bool`

HasNoaaSingleReasonWithDoddFrank returns a boolean if a field has been set.

### GetPrivacyPolicy

`func (o *PolicyDocumentResponse) GetPrivacyPolicy() PolicyDocumentAssetResponse`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *PolicyDocumentResponse) GetPrivacyPolicyOk() (*PolicyDocumentAssetResponse, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *PolicyDocumentResponse) SetPrivacyPolicy(v PolicyDocumentAssetResponse)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *PolicyDocumentResponse) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.

### GetRewardsDisclosure

`func (o *PolicyDocumentResponse) GetRewardsDisclosure() PolicyDocumentAssetAndTemplateResponse`

GetRewardsDisclosure returns the RewardsDisclosure field if non-nil, zero value otherwise.

### GetRewardsDisclosureOk

`func (o *PolicyDocumentResponse) GetRewardsDisclosureOk() (*PolicyDocumentAssetAndTemplateResponse, bool)`

GetRewardsDisclosureOk returns a tuple with the RewardsDisclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsDisclosure

`func (o *PolicyDocumentResponse) SetRewardsDisclosure(v PolicyDocumentAssetAndTemplateResponse)`

SetRewardsDisclosure sets RewardsDisclosure field to given value.

### HasRewardsDisclosure

`func (o *PolicyDocumentResponse) HasRewardsDisclosure() bool`

HasRewardsDisclosure returns a boolean if a field has been set.

### GetSummaryOfCreditTerms

`func (o *PolicyDocumentResponse) GetSummaryOfCreditTerms() PolicyDocumentAssetAndTemplateResponse`

GetSummaryOfCreditTerms returns the SummaryOfCreditTerms field if non-nil, zero value otherwise.

### GetSummaryOfCreditTermsOk

`func (o *PolicyDocumentResponse) GetSummaryOfCreditTermsOk() (*PolicyDocumentAssetAndTemplateResponse, bool)`

GetSummaryOfCreditTermsOk returns a tuple with the SummaryOfCreditTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryOfCreditTerms

`func (o *PolicyDocumentResponse) SetSummaryOfCreditTerms(v PolicyDocumentAssetAndTemplateResponse)`

SetSummaryOfCreditTerms sets SummaryOfCreditTerms field to given value.

### HasSummaryOfCreditTerms

`func (o *PolicyDocumentResponse) HasSummaryOfCreditTerms() bool`

HasSummaryOfCreditTerms returns a boolean if a field has been set.

### GetTermsSchedule

`func (o *PolicyDocumentResponse) GetTermsSchedule() PolicyDocumentTemplateResponse`

GetTermsSchedule returns the TermsSchedule field if non-nil, zero value otherwise.

### GetTermsScheduleOk

`func (o *PolicyDocumentResponse) GetTermsScheduleOk() (*PolicyDocumentTemplateResponse, bool)`

GetTermsScheduleOk returns a tuple with the TermsSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsSchedule

`func (o *PolicyDocumentResponse) SetTermsSchedule(v PolicyDocumentTemplateResponse)`

SetTermsSchedule sets TermsSchedule field to given value.

### HasTermsSchedule

`func (o *PolicyDocumentResponse) HasTermsSchedule() bool`

HasTermsSchedule returns a boolean if a field has been set.

### GetToken

`func (o *PolicyDocumentResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyDocumentResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyDocumentResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyDocumentResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *PolicyDocumentResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PolicyDocumentResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PolicyDocumentResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PolicyDocumentResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


