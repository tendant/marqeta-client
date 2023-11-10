# CreateApplicationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnyNonTaxableIncome** | **bool** | A value of &#x60;true&#x60; indicates that the user has a non-taxable income source.  Required when retrieving an application. | 
**BundleToken** | **string** | Unique identifier of the bundle associated with the application. | 
**DeviceData** | Pointer to [**DeviceData**](DeviceData.md) |  | [optional] 
**EDisclosureTrackingToken** | **string** | The tracking token of the eDisclosure.  This is the &#x60;tracking_token&#x60; returned in the &#x60;E_DISCLOSURE&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before a decision on the application is made. | 
**MetaData** | Pointer to **map[string]interface{}** | Customer-defined additional information about the application. | [optional] 
**MonthlyMortgageOrRent** | **float32** | Monthly amount of the mortgage or rent that the user currently pays.  Required when retrieving an application. | 
**OfferId** | Pointer to **string** | Unique identifier of the offer for a pre-screened applicant. | [optional] 
**PrequalifiedOfferPreTermsTrackingToken** | Pointer to **string** | The tracking token of the Pre-qualified Offer Pre-Terms.  This is the &#x60;tracking_token&#x60; returned in the &#x60;PREQUALIFICATION_PRE_TERMS&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before a decision on the application is made. | [optional] 
**PrimaryIncomeSource** | **string** | Whether the primary income source comes from the user being employed, unemployed, self-employment, or another situation.  Required when retrieving an application. | 
**PrivacyPolicyTrackingToken** | **string** | The tracking token of the Privacy Policy.  This is the &#x60;tracking_token&#x60; returned in the &#x60;PRIVACY_POLICY&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before a decision on the application is made. | 
**ResidenceType** | **string** | Whether the user owns or rents their residence, or has another situation.  Required when retrieving an application. | 
**RewardsDisclosurePreTermsTrackingToken** | Pointer to **string** | The tracking token of the Rewards Disclosure Pre-terms.  This is the &#x60;tracking_token&#x60; returned in the &#x60;REWARDS_DISCLOSURE_PRE_TERMS&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before a decision on the application is made. | [optional] 
**SoctTrackingToken** | **string** | The tracking token of the Summary of Credit Terms (SOCT).  This is the &#x60;tracking_token&#x60; returned in the &#x60;SOCT&#x60; object when sending a &#x60;GET&#x60; request to the &#x60;/credit/applications/files&#x60; endpoint before a decision on the application is made. | 
**Token** | Pointer to **string** | Unique identifier of the application. | [optional] 
**TotalAnnualIncome** | **float32** | The total amount of the user&#39;s annual income.  Required when retrieving an application. | 
**Type** | Pointer to [**ApplicationType**](ApplicationType.md) |  | [optional] 
**UserToken** | **string** | Unique identifier of the applicant, the user applying for a credit account. | 

## Methods

### NewCreateApplicationsRequest

`func NewCreateApplicationsRequest(anyNonTaxableIncome bool, bundleToken string, eDisclosureTrackingToken string, monthlyMortgageOrRent float32, primaryIncomeSource string, privacyPolicyTrackingToken string, residenceType string, soctTrackingToken string, totalAnnualIncome float32, userToken string, ) *CreateApplicationsRequest`

NewCreateApplicationsRequest instantiates a new CreateApplicationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationsRequestWithDefaults

`func NewCreateApplicationsRequestWithDefaults() *CreateApplicationsRequest`

NewCreateApplicationsRequestWithDefaults instantiates a new CreateApplicationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnyNonTaxableIncome

`func (o *CreateApplicationsRequest) GetAnyNonTaxableIncome() bool`

GetAnyNonTaxableIncome returns the AnyNonTaxableIncome field if non-nil, zero value otherwise.

### GetAnyNonTaxableIncomeOk

`func (o *CreateApplicationsRequest) GetAnyNonTaxableIncomeOk() (*bool, bool)`

GetAnyNonTaxableIncomeOk returns a tuple with the AnyNonTaxableIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyNonTaxableIncome

`func (o *CreateApplicationsRequest) SetAnyNonTaxableIncome(v bool)`

SetAnyNonTaxableIncome sets AnyNonTaxableIncome field to given value.


### GetBundleToken

`func (o *CreateApplicationsRequest) GetBundleToken() string`

GetBundleToken returns the BundleToken field if non-nil, zero value otherwise.

### GetBundleTokenOk

`func (o *CreateApplicationsRequest) GetBundleTokenOk() (*string, bool)`

GetBundleTokenOk returns a tuple with the BundleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleToken

`func (o *CreateApplicationsRequest) SetBundleToken(v string)`

SetBundleToken sets BundleToken field to given value.


### GetDeviceData

`func (o *CreateApplicationsRequest) GetDeviceData() DeviceData`

GetDeviceData returns the DeviceData field if non-nil, zero value otherwise.

### GetDeviceDataOk

`func (o *CreateApplicationsRequest) GetDeviceDataOk() (*DeviceData, bool)`

GetDeviceDataOk returns a tuple with the DeviceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceData

`func (o *CreateApplicationsRequest) SetDeviceData(v DeviceData)`

SetDeviceData sets DeviceData field to given value.

### HasDeviceData

`func (o *CreateApplicationsRequest) HasDeviceData() bool`

HasDeviceData returns a boolean if a field has been set.

### GetEDisclosureTrackingToken

`func (o *CreateApplicationsRequest) GetEDisclosureTrackingToken() string`

GetEDisclosureTrackingToken returns the EDisclosureTrackingToken field if non-nil, zero value otherwise.

### GetEDisclosureTrackingTokenOk

`func (o *CreateApplicationsRequest) GetEDisclosureTrackingTokenOk() (*string, bool)`

GetEDisclosureTrackingTokenOk returns a tuple with the EDisclosureTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDisclosureTrackingToken

`func (o *CreateApplicationsRequest) SetEDisclosureTrackingToken(v string)`

SetEDisclosureTrackingToken sets EDisclosureTrackingToken field to given value.


### GetMetaData

`func (o *CreateApplicationsRequest) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CreateApplicationsRequest) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CreateApplicationsRequest) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *CreateApplicationsRequest) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMonthlyMortgageOrRent

`func (o *CreateApplicationsRequest) GetMonthlyMortgageOrRent() float32`

GetMonthlyMortgageOrRent returns the MonthlyMortgageOrRent field if non-nil, zero value otherwise.

### GetMonthlyMortgageOrRentOk

`func (o *CreateApplicationsRequest) GetMonthlyMortgageOrRentOk() (*float32, bool)`

GetMonthlyMortgageOrRentOk returns a tuple with the MonthlyMortgageOrRent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyMortgageOrRent

`func (o *CreateApplicationsRequest) SetMonthlyMortgageOrRent(v float32)`

SetMonthlyMortgageOrRent sets MonthlyMortgageOrRent field to given value.


### GetOfferId

`func (o *CreateApplicationsRequest) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *CreateApplicationsRequest) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *CreateApplicationsRequest) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *CreateApplicationsRequest) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetPrequalifiedOfferPreTermsTrackingToken

`func (o *CreateApplicationsRequest) GetPrequalifiedOfferPreTermsTrackingToken() string`

GetPrequalifiedOfferPreTermsTrackingToken returns the PrequalifiedOfferPreTermsTrackingToken field if non-nil, zero value otherwise.

### GetPrequalifiedOfferPreTermsTrackingTokenOk

`func (o *CreateApplicationsRequest) GetPrequalifiedOfferPreTermsTrackingTokenOk() (*string, bool)`

GetPrequalifiedOfferPreTermsTrackingTokenOk returns a tuple with the PrequalifiedOfferPreTermsTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrequalifiedOfferPreTermsTrackingToken

`func (o *CreateApplicationsRequest) SetPrequalifiedOfferPreTermsTrackingToken(v string)`

SetPrequalifiedOfferPreTermsTrackingToken sets PrequalifiedOfferPreTermsTrackingToken field to given value.

### HasPrequalifiedOfferPreTermsTrackingToken

`func (o *CreateApplicationsRequest) HasPrequalifiedOfferPreTermsTrackingToken() bool`

HasPrequalifiedOfferPreTermsTrackingToken returns a boolean if a field has been set.

### GetPrimaryIncomeSource

`func (o *CreateApplicationsRequest) GetPrimaryIncomeSource() string`

GetPrimaryIncomeSource returns the PrimaryIncomeSource field if non-nil, zero value otherwise.

### GetPrimaryIncomeSourceOk

`func (o *CreateApplicationsRequest) GetPrimaryIncomeSourceOk() (*string, bool)`

GetPrimaryIncomeSourceOk returns a tuple with the PrimaryIncomeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIncomeSource

`func (o *CreateApplicationsRequest) SetPrimaryIncomeSource(v string)`

SetPrimaryIncomeSource sets PrimaryIncomeSource field to given value.


### GetPrivacyPolicyTrackingToken

`func (o *CreateApplicationsRequest) GetPrivacyPolicyTrackingToken() string`

GetPrivacyPolicyTrackingToken returns the PrivacyPolicyTrackingToken field if non-nil, zero value otherwise.

### GetPrivacyPolicyTrackingTokenOk

`func (o *CreateApplicationsRequest) GetPrivacyPolicyTrackingTokenOk() (*string, bool)`

GetPrivacyPolicyTrackingTokenOk returns a tuple with the PrivacyPolicyTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyTrackingToken

`func (o *CreateApplicationsRequest) SetPrivacyPolicyTrackingToken(v string)`

SetPrivacyPolicyTrackingToken sets PrivacyPolicyTrackingToken field to given value.


### GetResidenceType

`func (o *CreateApplicationsRequest) GetResidenceType() string`

GetResidenceType returns the ResidenceType field if non-nil, zero value otherwise.

### GetResidenceTypeOk

`func (o *CreateApplicationsRequest) GetResidenceTypeOk() (*string, bool)`

GetResidenceTypeOk returns a tuple with the ResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceType

`func (o *CreateApplicationsRequest) SetResidenceType(v string)`

SetResidenceType sets ResidenceType field to given value.


### GetRewardsDisclosurePreTermsTrackingToken

`func (o *CreateApplicationsRequest) GetRewardsDisclosurePreTermsTrackingToken() string`

GetRewardsDisclosurePreTermsTrackingToken returns the RewardsDisclosurePreTermsTrackingToken field if non-nil, zero value otherwise.

### GetRewardsDisclosurePreTermsTrackingTokenOk

`func (o *CreateApplicationsRequest) GetRewardsDisclosurePreTermsTrackingTokenOk() (*string, bool)`

GetRewardsDisclosurePreTermsTrackingTokenOk returns a tuple with the RewardsDisclosurePreTermsTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsDisclosurePreTermsTrackingToken

`func (o *CreateApplicationsRequest) SetRewardsDisclosurePreTermsTrackingToken(v string)`

SetRewardsDisclosurePreTermsTrackingToken sets RewardsDisclosurePreTermsTrackingToken field to given value.

### HasRewardsDisclosurePreTermsTrackingToken

`func (o *CreateApplicationsRequest) HasRewardsDisclosurePreTermsTrackingToken() bool`

HasRewardsDisclosurePreTermsTrackingToken returns a boolean if a field has been set.

### GetSoctTrackingToken

`func (o *CreateApplicationsRequest) GetSoctTrackingToken() string`

GetSoctTrackingToken returns the SoctTrackingToken field if non-nil, zero value otherwise.

### GetSoctTrackingTokenOk

`func (o *CreateApplicationsRequest) GetSoctTrackingTokenOk() (*string, bool)`

GetSoctTrackingTokenOk returns a tuple with the SoctTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoctTrackingToken

`func (o *CreateApplicationsRequest) SetSoctTrackingToken(v string)`

SetSoctTrackingToken sets SoctTrackingToken field to given value.


### GetToken

`func (o *CreateApplicationsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateApplicationsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateApplicationsRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateApplicationsRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTotalAnnualIncome

`func (o *CreateApplicationsRequest) GetTotalAnnualIncome() float32`

GetTotalAnnualIncome returns the TotalAnnualIncome field if non-nil, zero value otherwise.

### GetTotalAnnualIncomeOk

`func (o *CreateApplicationsRequest) GetTotalAnnualIncomeOk() (*float32, bool)`

GetTotalAnnualIncomeOk returns a tuple with the TotalAnnualIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnnualIncome

`func (o *CreateApplicationsRequest) SetTotalAnnualIncome(v float32)`

SetTotalAnnualIncome sets TotalAnnualIncome field to given value.


### GetType

`func (o *CreateApplicationsRequest) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApplicationsRequest) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApplicationsRequest) SetType(v ApplicationType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateApplicationsRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserToken

`func (o *CreateApplicationsRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CreateApplicationsRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CreateApplicationsRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


