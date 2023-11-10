# FundingAccountResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSuffix** | Pointer to **string** | Account identifier appended to the bank account number. This field is returned if it exists in the resource. | [optional] 
**AccountType** | Pointer to **string** | Type of account. This field is returned if it exists in the resource. | [optional] 
**Active** | Pointer to **bool** | Specifies whether the account is active. This field is returned if it exists in the resource. | [optional] [default to false]
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder. This token is returned if a &#x60;user_token&#x60; is not specified. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**DateSentForVerification** | Pointer to **time.Time** | Date and time when the account was sent for verification, in UTC. This field is returned if it exists in the resource. | [optional] 
**DateVerified** | Pointer to **time.Time** | Date and time when the account was verified, in UTC. This field is returned if it exists in the resource. | [optional] 
**ExpDate** | Pointer to **string** | Payment card expiration date. This field is returned if it exists in the resource. | [optional] 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source.  This field is returned if it exists in the resource. | [optional] [default to false]
**LastModifiedTime** | **time.Time** | Date and time when the resource was last modified, in UTC. | 
**LinkPartnerAccountReferenceToken** | Pointer to **string** |  | [optional] 
**NameOnAccount** | Pointer to **string** | Name on the account. This field is returned if it exists in the resource. | [optional] 
**Partner** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the funding source. This field is returned if it exists in the resource. | [optional] 
**Type** | Pointer to **string** | Funding source type. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder. This token is returned if a &#x60;business_token&#x60; is not specified. | [optional] 
**VerificationNotes** | Pointer to **string** | Free-form text field for holding notes about verification. This field is returned only if &#x60;verification_override &#x3D; true&#x60;. | [optional] 
**VerificationOverride** | Pointer to **bool** | Allows the ACH funding source to be used regardless of its verification status. | [optional] [default to false]
**VerificationStatus** | Pointer to **string** | Account verification status. This field is returned if it exists in the resource. | [optional] 

## Methods

### NewFundingAccountResponseModel

`func NewFundingAccountResponseModel(createdTime time.Time, lastModifiedTime time.Time, ) *FundingAccountResponseModel`

NewFundingAccountResponseModel instantiates a new FundingAccountResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingAccountResponseModelWithDefaults

`func NewFundingAccountResponseModelWithDefaults() *FundingAccountResponseModel`

NewFundingAccountResponseModelWithDefaults instantiates a new FundingAccountResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSuffix

`func (o *FundingAccountResponseModel) GetAccountSuffix() string`

GetAccountSuffix returns the AccountSuffix field if non-nil, zero value otherwise.

### GetAccountSuffixOk

`func (o *FundingAccountResponseModel) GetAccountSuffixOk() (*string, bool)`

GetAccountSuffixOk returns a tuple with the AccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSuffix

`func (o *FundingAccountResponseModel) SetAccountSuffix(v string)`

SetAccountSuffix sets AccountSuffix field to given value.

### HasAccountSuffix

`func (o *FundingAccountResponseModel) HasAccountSuffix() bool`

HasAccountSuffix returns a boolean if a field has been set.

### GetAccountType

`func (o *FundingAccountResponseModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *FundingAccountResponseModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *FundingAccountResponseModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *FundingAccountResponseModel) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetActive

`func (o *FundingAccountResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FundingAccountResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FundingAccountResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FundingAccountResponseModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBusinessToken

`func (o *FundingAccountResponseModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *FundingAccountResponseModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *FundingAccountResponseModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *FundingAccountResponseModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *FundingAccountResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FundingAccountResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FundingAccountResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDateSentForVerification

`func (o *FundingAccountResponseModel) GetDateSentForVerification() time.Time`

GetDateSentForVerification returns the DateSentForVerification field if non-nil, zero value otherwise.

### GetDateSentForVerificationOk

`func (o *FundingAccountResponseModel) GetDateSentForVerificationOk() (*time.Time, bool)`

GetDateSentForVerificationOk returns a tuple with the DateSentForVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSentForVerification

`func (o *FundingAccountResponseModel) SetDateSentForVerification(v time.Time)`

SetDateSentForVerification sets DateSentForVerification field to given value.

### HasDateSentForVerification

`func (o *FundingAccountResponseModel) HasDateSentForVerification() bool`

HasDateSentForVerification returns a boolean if a field has been set.

### GetDateVerified

`func (o *FundingAccountResponseModel) GetDateVerified() time.Time`

GetDateVerified returns the DateVerified field if non-nil, zero value otherwise.

### GetDateVerifiedOk

`func (o *FundingAccountResponseModel) GetDateVerifiedOk() (*time.Time, bool)`

GetDateVerifiedOk returns a tuple with the DateVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateVerified

`func (o *FundingAccountResponseModel) SetDateVerified(v time.Time)`

SetDateVerified sets DateVerified field to given value.

### HasDateVerified

`func (o *FundingAccountResponseModel) HasDateVerified() bool`

HasDateVerified returns a boolean if a field has been set.

### GetExpDate

`func (o *FundingAccountResponseModel) GetExpDate() string`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *FundingAccountResponseModel) GetExpDateOk() (*string, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *FundingAccountResponseModel) SetExpDate(v string)`

SetExpDate sets ExpDate field to given value.

### HasExpDate

`func (o *FundingAccountResponseModel) HasExpDate() bool`

HasExpDate returns a boolean if a field has been set.

### GetIsDefaultAccount

`func (o *FundingAccountResponseModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *FundingAccountResponseModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *FundingAccountResponseModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *FundingAccountResponseModel) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *FundingAccountResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *FundingAccountResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *FundingAccountResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLinkPartnerAccountReferenceToken

`func (o *FundingAccountResponseModel) GetLinkPartnerAccountReferenceToken() string`

GetLinkPartnerAccountReferenceToken returns the LinkPartnerAccountReferenceToken field if non-nil, zero value otherwise.

### GetLinkPartnerAccountReferenceTokenOk

`func (o *FundingAccountResponseModel) GetLinkPartnerAccountReferenceTokenOk() (*string, bool)`

GetLinkPartnerAccountReferenceTokenOk returns a tuple with the LinkPartnerAccountReferenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPartnerAccountReferenceToken

`func (o *FundingAccountResponseModel) SetLinkPartnerAccountReferenceToken(v string)`

SetLinkPartnerAccountReferenceToken sets LinkPartnerAccountReferenceToken field to given value.

### HasLinkPartnerAccountReferenceToken

`func (o *FundingAccountResponseModel) HasLinkPartnerAccountReferenceToken() bool`

HasLinkPartnerAccountReferenceToken returns a boolean if a field has been set.

### GetNameOnAccount

`func (o *FundingAccountResponseModel) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *FundingAccountResponseModel) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *FundingAccountResponseModel) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.

### HasNameOnAccount

`func (o *FundingAccountResponseModel) HasNameOnAccount() bool`

HasNameOnAccount returns a boolean if a field has been set.

### GetPartner

`func (o *FundingAccountResponseModel) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *FundingAccountResponseModel) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *FundingAccountResponseModel) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *FundingAccountResponseModel) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetToken

`func (o *FundingAccountResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FundingAccountResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FundingAccountResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FundingAccountResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *FundingAccountResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FundingAccountResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FundingAccountResponseModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FundingAccountResponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserToken

`func (o *FundingAccountResponseModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *FundingAccountResponseModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *FundingAccountResponseModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *FundingAccountResponseModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetVerificationNotes

`func (o *FundingAccountResponseModel) GetVerificationNotes() string`

GetVerificationNotes returns the VerificationNotes field if non-nil, zero value otherwise.

### GetVerificationNotesOk

`func (o *FundingAccountResponseModel) GetVerificationNotesOk() (*string, bool)`

GetVerificationNotesOk returns a tuple with the VerificationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNotes

`func (o *FundingAccountResponseModel) SetVerificationNotes(v string)`

SetVerificationNotes sets VerificationNotes field to given value.

### HasVerificationNotes

`func (o *FundingAccountResponseModel) HasVerificationNotes() bool`

HasVerificationNotes returns a boolean if a field has been set.

### GetVerificationOverride

`func (o *FundingAccountResponseModel) GetVerificationOverride() bool`

GetVerificationOverride returns the VerificationOverride field if non-nil, zero value otherwise.

### GetVerificationOverrideOk

`func (o *FundingAccountResponseModel) GetVerificationOverrideOk() (*bool, bool)`

GetVerificationOverrideOk returns a tuple with the VerificationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationOverride

`func (o *FundingAccountResponseModel) SetVerificationOverride(v bool)`

SetVerificationOverride sets VerificationOverride field to given value.

### HasVerificationOverride

`func (o *FundingAccountResponseModel) HasVerificationOverride() bool`

HasVerificationOverride returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *FundingAccountResponseModel) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *FundingAccountResponseModel) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *FundingAccountResponseModel) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *FundingAccountResponseModel) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


