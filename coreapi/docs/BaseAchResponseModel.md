# BaseAchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSuffix** | **string** | ACH account identifier appended to the bank account number. | 
**AccountType** | **string** | Type of account. | 
**Active** | **bool** | Specifies whether the account is active. | [default to false]
**BankName** | Pointer to **string** | Name of the bank holding the account. This field is returned if it exists in the resource. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**DateSentForVerification** | Pointer to **time.Time** | Date and time when the account was sent for verification, in UTC. This field is returned if it exists in the resource. | [optional] 
**DateVerified** | Pointer to **time.Time** | Date and time when the account was verified, in UTC.  This field is returned if it exists in the resource. | [optional] 
**IsDefaultAccount** | Pointer to **bool** | If there are multiple funding sources, this field specifies which source is used by default in funding calls. If there is only one funding source, the system ignores this field and always uses that source. | [optional] [default to false]
**LastModifiedTime** | **time.Time** | Date and time when the resource was last modified, in UTC. | 
**NameOnAccount** | **string** | Name on the ACH account. | 
**Partner** | Pointer to **string** |  | [optional] 
**PartnerAccountLinkReferenceToken** | Pointer to **string** |  | [optional] 
**Token** | **string** | The unique identifier of the funding source. | 
**VerificationNotes** | Pointer to **string** | Free-form text field for holding notes about verification. This field is returned only if &#x60;verification_override &#x3D; true&#x60;. | [optional] 
**VerificationOverride** | Pointer to **bool** | Allows the ACH funding source to be used regardless of its verification status. This field is returned if it exists in the resource. | [optional] [default to false]
**VerificationStatus** | Pointer to **string** | Account verification status. This field is returned if it exists in the resource. | [optional] 

## Methods

### NewBaseAchResponseModel

`func NewBaseAchResponseModel(accountSuffix string, accountType string, active bool, createdTime time.Time, lastModifiedTime time.Time, nameOnAccount string, token string, ) *BaseAchResponseModel`

NewBaseAchResponseModel instantiates a new BaseAchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAchResponseModelWithDefaults

`func NewBaseAchResponseModelWithDefaults() *BaseAchResponseModel`

NewBaseAchResponseModelWithDefaults instantiates a new BaseAchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSuffix

`func (o *BaseAchResponseModel) GetAccountSuffix() string`

GetAccountSuffix returns the AccountSuffix field if non-nil, zero value otherwise.

### GetAccountSuffixOk

`func (o *BaseAchResponseModel) GetAccountSuffixOk() (*string, bool)`

GetAccountSuffixOk returns a tuple with the AccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSuffix

`func (o *BaseAchResponseModel) SetAccountSuffix(v string)`

SetAccountSuffix sets AccountSuffix field to given value.


### GetAccountType

`func (o *BaseAchResponseModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BaseAchResponseModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BaseAchResponseModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetActive

`func (o *BaseAchResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BaseAchResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BaseAchResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBankName

`func (o *BaseAchResponseModel) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BaseAchResponseModel) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BaseAchResponseModel) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *BaseAchResponseModel) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BaseAchResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BaseAchResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BaseAchResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDateSentForVerification

`func (o *BaseAchResponseModel) GetDateSentForVerification() time.Time`

GetDateSentForVerification returns the DateSentForVerification field if non-nil, zero value otherwise.

### GetDateSentForVerificationOk

`func (o *BaseAchResponseModel) GetDateSentForVerificationOk() (*time.Time, bool)`

GetDateSentForVerificationOk returns a tuple with the DateSentForVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSentForVerification

`func (o *BaseAchResponseModel) SetDateSentForVerification(v time.Time)`

SetDateSentForVerification sets DateSentForVerification field to given value.

### HasDateSentForVerification

`func (o *BaseAchResponseModel) HasDateSentForVerification() bool`

HasDateSentForVerification returns a boolean if a field has been set.

### GetDateVerified

`func (o *BaseAchResponseModel) GetDateVerified() time.Time`

GetDateVerified returns the DateVerified field if non-nil, zero value otherwise.

### GetDateVerifiedOk

`func (o *BaseAchResponseModel) GetDateVerifiedOk() (*time.Time, bool)`

GetDateVerifiedOk returns a tuple with the DateVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateVerified

`func (o *BaseAchResponseModel) SetDateVerified(v time.Time)`

SetDateVerified sets DateVerified field to given value.

### HasDateVerified

`func (o *BaseAchResponseModel) HasDateVerified() bool`

HasDateVerified returns a boolean if a field has been set.

### GetIsDefaultAccount

`func (o *BaseAchResponseModel) GetIsDefaultAccount() bool`

GetIsDefaultAccount returns the IsDefaultAccount field if non-nil, zero value otherwise.

### GetIsDefaultAccountOk

`func (o *BaseAchResponseModel) GetIsDefaultAccountOk() (*bool, bool)`

GetIsDefaultAccountOk returns a tuple with the IsDefaultAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccount

`func (o *BaseAchResponseModel) SetIsDefaultAccount(v bool)`

SetIsDefaultAccount sets IsDefaultAccount field to given value.

### HasIsDefaultAccount

`func (o *BaseAchResponseModel) HasIsDefaultAccount() bool`

HasIsDefaultAccount returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BaseAchResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BaseAchResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BaseAchResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetNameOnAccount

`func (o *BaseAchResponseModel) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *BaseAchResponseModel) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *BaseAchResponseModel) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.


### GetPartner

`func (o *BaseAchResponseModel) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *BaseAchResponseModel) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *BaseAchResponseModel) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *BaseAchResponseModel) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPartnerAccountLinkReferenceToken

`func (o *BaseAchResponseModel) GetPartnerAccountLinkReferenceToken() string`

GetPartnerAccountLinkReferenceToken returns the PartnerAccountLinkReferenceToken field if non-nil, zero value otherwise.

### GetPartnerAccountLinkReferenceTokenOk

`func (o *BaseAchResponseModel) GetPartnerAccountLinkReferenceTokenOk() (*string, bool)`

GetPartnerAccountLinkReferenceTokenOk returns a tuple with the PartnerAccountLinkReferenceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAccountLinkReferenceToken

`func (o *BaseAchResponseModel) SetPartnerAccountLinkReferenceToken(v string)`

SetPartnerAccountLinkReferenceToken sets PartnerAccountLinkReferenceToken field to given value.

### HasPartnerAccountLinkReferenceToken

`func (o *BaseAchResponseModel) HasPartnerAccountLinkReferenceToken() bool`

HasPartnerAccountLinkReferenceToken returns a boolean if a field has been set.

### GetToken

`func (o *BaseAchResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BaseAchResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BaseAchResponseModel) SetToken(v string)`

SetToken sets Token field to given value.


### GetVerificationNotes

`func (o *BaseAchResponseModel) GetVerificationNotes() string`

GetVerificationNotes returns the VerificationNotes field if non-nil, zero value otherwise.

### GetVerificationNotesOk

`func (o *BaseAchResponseModel) GetVerificationNotesOk() (*string, bool)`

GetVerificationNotesOk returns a tuple with the VerificationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNotes

`func (o *BaseAchResponseModel) SetVerificationNotes(v string)`

SetVerificationNotes sets VerificationNotes field to given value.

### HasVerificationNotes

`func (o *BaseAchResponseModel) HasVerificationNotes() bool`

HasVerificationNotes returns a boolean if a field has been set.

### GetVerificationOverride

`func (o *BaseAchResponseModel) GetVerificationOverride() bool`

GetVerificationOverride returns the VerificationOverride field if non-nil, zero value otherwise.

### GetVerificationOverrideOk

`func (o *BaseAchResponseModel) GetVerificationOverrideOk() (*bool, bool)`

GetVerificationOverrideOk returns a tuple with the VerificationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationOverride

`func (o *BaseAchResponseModel) SetVerificationOverride(v bool)`

SetVerificationOverride sets VerificationOverride field to given value.

### HasVerificationOverride

`func (o *BaseAchResponseModel) HasVerificationOverride() bool`

HasVerificationOverride returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *BaseAchResponseModel) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *BaseAchResponseModel) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *BaseAchResponseModel) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *BaseAchResponseModel) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


