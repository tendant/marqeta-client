# AchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSuffix** | **string** | ACH account identifier appended to the bank account number. | 
**AccountType** | **string** | Type of account. | 
**Active** | **bool** | Specifies whether the account is active. | [default to false]
**BankName** | Pointer to **string** | Name of the bank holding the account. This field is returned if it exists in the resource. | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder. This token is returned if a &#x60;user_token&#x60; is not specified. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**DateSentForVerification** | Pointer to **time.Time** | Date and time when the account was sent for verification, in UTC.  This field is returned if it exists in the resource. | [optional] 
**DateVerified** | Pointer to **time.Time** | Date and time when the account was verified, in UTC.  This field is returned if it exists in the resource. | [optional] 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source.  This field is returned if it exists in the resource. | [optional] [default to false]
**LastModifiedTime** | **time.Time** | Date and time when the resource was last modified, in UTC. | 
**NameOnAccount** | **string** | Name on the ACH account. | 
**Partner** | Pointer to **string** |  | [optional] 
**PartnerAccountLinkReferenceToken** | Pointer to **string** |  | [optional] 
**Token** | **string** | Unique identifier of the funding source. | 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder. This token is returned if a &#x60;business_token&#x60; is not specified. | [optional] 
**VerificationNotes** | Pointer to **string** | Free-form text field for holding notes about verification. This field is returned only if &#x60;verification_override &#x3D; true&#x60;. | [optional] 
**VerificationOverride** | Pointer to **bool** | Allows the ACH funding source to be used regardless of its verification status. This field is returned if it exists in the resource. | [optional] [default to false]
**VerificationStatus** | Pointer to **string** | Account verification status. This field is returned if it exists in the resource. | [optional] 

## Methods

### NewAchResponseModel

`func NewAchResponseModel(accountSuffix string, accountType string, active bool, createdTime time.Time, lastModifiedTime time.Time, nameOnAccount string, token string, ) *AchResponseModel`

NewAchResponseModel instantiates a new AchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchResponseModelWithDefaults

`func NewAchResponseModelWithDefaults() *AchResponseModel`

NewAchResponseModelWithDefaults instantiates a new AchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSuffix

`func (o *AchResponseModel) GetAccountSuffix() string`

GetAccountSuffix returns the AccountSuffix field if non-nil, zero value otherwise.

### GetAccountSuffixOk

`func (o *AchResponseModel) GetAccountSuffixOk() (*string, bool)`

GetAccountSuffixOk returns a tuple with the AccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSuffix

`func (o *AchResponseModel) SetAccountSuffix(v string)`

SetAccountSuffix sets AccountSuffix field to given value.


### GetAccountType

`func (o *AchResponseModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AchResponseModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AchResponseModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetActive

`func (o *AchResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AchResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AchResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBankName

`func (o *AchResponseModel) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *AchResponseModel) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *AchResponseModel) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *AchResponseModel) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBusinessToken

`func (o *AchResponseModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *AchResponseModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *AchResponseModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *AchResponseModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AchResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AchResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AchResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDateSentForVerification

`func (o *AchResponseModel) GetDateSentForVerification() time.Time`

GetDateSentForVerification returns the DateSentForVerification field if non-nil, zero value otherwise.

### GetDateSentForVerificationOk

`func (o *AchResponseModel) GetDateSentForVerificationOk() (*time.Time, bool)`

GetDateSentForVerificationOk returns a tuple with the DateSentForVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSentForVerification

`func (o *AchResponseModel) SetDateSentForVerification(v time.Time)`

SetDateSentForVerification sets DateSentForVerification field to given value.

### HasDateSentForVerification

`func (o *AchResponseModel) HasDateSentForVerification() bool`

HasDateSentForVerification returns a boolean if a field has been set.

### GetDateVerified

`func (o *AchResponseModel) GetDateVerified() time.Time`

GetDateVerified returns the DateVerified field if non-nil, zero value otherwise.

### GetDateVerifiedOk

`func (o *AchResponseModel) GetDateVerifiedOk() (*time.Time, bool)`

GetDateVerifiedOk returns a tuple with the DateVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateVerified

`func (o *AchResponseModel) SetDateVerified(v time.Time)`

SetDateVerified sets DateVerified field to given value.

### HasDateVerified

`func (o *AchResponseModel) HasDateVerified() bool`

HasDateVerified returns a boolean if a field has been set.

### GetIsDefaultAccount

`func (o *AchResponseModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *AchResponseModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *AchResponseModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *AchResponseModel) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *AchResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *AchResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *AchResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetNameOnAccount

`func (o *AchResponseModel) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *AchResponseModel) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *AchResponseModel) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.


### GetPartner

`func (o *AchResponseModel) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *AchResponseModel) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *AchResponseModel) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *AchResponseModel) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPartnerAccountLinkReferenceToken

`func (o *AchResponseModel) GetPartnerAccountLinkReferenceToken() string`

GetPartnerAccountLinkReferenceToken returns the PartnerAccountLinkReferenceToken field if non-nil, zero value otherwise.

### GetPartnerAccountLinkReferenceTokenOk

`func (o *AchResponseModel) GetPartnerAccountLinkReferenceTokenOk() (*string, bool)`

GetPartnerAccountLinkReferenceTokenOk returns a tuple with the PartnerAccountLinkReferenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAccountLinkReferenceToken

`func (o *AchResponseModel) SetPartnerAccountLinkReferenceToken(v string)`

SetPartnerAccountLinkReferenceToken sets PartnerAccountLinkReferenceToken field to given value.

### HasPartnerAccountLinkReferenceToken

`func (o *AchResponseModel) HasPartnerAccountLinkReferenceToken() bool`

HasPartnerAccountLinkReferenceToken returns a boolean if a field has been set.

### GetToken

`func (o *AchResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AchResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AchResponseModel) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *AchResponseModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AchResponseModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AchResponseModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *AchResponseModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetVerificationNotes

`func (o *AchResponseModel) GetVerificationNotes() string`

GetVerificationNotes returns the VerificationNotes field if non-nil, zero value otherwise.

### GetVerificationNotesOk

`func (o *AchResponseModel) GetVerificationNotesOk() (*string, bool)`

GetVerificationNotesOk returns a tuple with the VerificationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNotes

`func (o *AchResponseModel) SetVerificationNotes(v string)`

SetVerificationNotes sets VerificationNotes field to given value.

### HasVerificationNotes

`func (o *AchResponseModel) HasVerificationNotes() bool`

HasVerificationNotes returns a boolean if a field has been set.

### GetVerificationOverride

`func (o *AchResponseModel) GetVerificationOverride() bool`

GetVerificationOverride returns the VerificationOverride field if non-nil, zero value otherwise.

### GetVerificationOverrideOk

`func (o *AchResponseModel) GetVerificationOverrideOk() (*bool, bool)`

GetVerificationOverrideOk returns a tuple with the VerificationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationOverride

`func (o *AchResponseModel) SetVerificationOverride(v bool)`

SetVerificationOverride sets VerificationOverride field to given value.

### HasVerificationOverride

`func (o *AchResponseModel) HasVerificationOverride() bool`

HasVerificationOverride returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *AchResponseModel) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *AchResponseModel) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *AchResponseModel) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *AchResponseModel) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


