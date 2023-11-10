# UserCardHolderUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderGroupToken** | Pointer to **string** | Associates the specified account holder group with the cardholder. Send a &#x60;GET&#x60; request to &#x60;/accountholdergroups&#x60; to retrieve account holder group tokens. | [optional] 
**Address1** | Pointer to **string** | Cardholder address.  *NOTE:* Required for KYC verification (US-based cardholders only). Cannot perform KYC if set to a PO Box. | [optional] 
**Address2** | Pointer to **string** | Additional address information for the cardholder.  *NOTE:* Cannot perform KYC if set to a PO Box. | [optional] 
**BirthDate** | Pointer to **string** | Cardholder date of birth.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**City** | Pointer to **string** | The city that corresponds to the address.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Company** | Pointer to **string** | Company name. | [optional] 
**CorporateCardHolder** | Pointer to **bool** | Specifies if the cardholder holds a corporate card. | [optional] [default to false]
**Country** | Pointer to **string** | Country in which the cardholder resides.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Email** | Pointer to **string** | Valid email address for the cardholder.  This value must be unique among cardholders. | [optional] 
**FirstName** | Pointer to **string** | Cardholder first name.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Gender** | Pointer to **string** | Gender of the cardholder. | [optional] 
**Honorific** | Pointer to **string** | Cardholder title or prefix: Ms., Mr., Miss, Mrs. | [optional] 
**IdCardExpirationDate** | Pointer to **string** | Expiration date of the cardholder&#39;s identification card. | [optional] 
**IdCardNumber** | Pointer to **string** | Cardholder&#39;s identification card number. | [optional] 
**Identifications** | Pointer to [**[]IdentificationRequestModel**](IdentificationRequestModel.md) | One or more objects containing identifications associated with the cardholder. | [optional] 
**IpAddress** | Pointer to **string** | Cardholder IP address. | [optional] 
**LastName** | Pointer to **string** | Cardholder last name.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Metadata** | Pointer to **map[string]string** | Associates any additional metadata you provide with the cardholder. | [optional] 
**MiddleName** | Pointer to **string** | Cardholder middle name. | [optional] 
**Nationality** | Pointer to **string** | Cardholder nationality. | [optional] 
**Notes** | Pointer to **string** | Any additional information pertaining to the cardholder. | [optional] 
**ParentToken** | Pointer to **string** | Unique identifier of an existing user or business resource.  Required if &#x60;uses_parent_account &#x3D; true&#x60;. This account holder is configured as the parent of the current cardholder.  To unlink a child account from a parent account, update this field to a null value. | [optional] 
**PassportExpirationDate** | Pointer to **string** | Expiration date of the cardholder&#39;s passport. | [optional] 
**PassportNumber** | Pointer to **string** | Cardholder passport number. | [optional] 
**Password** | Pointer to **string** | Cardholder&#39;s user account password on the Marqeta platform. | [optional] 
**Phone** | Pointer to **string** | Cardholder telephone number (including area code), prepended by the &#x60;+&#x60; symbol and the 1- to 3-digit country calling code. Do not include hyphens, spaces, or parentheses. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the cardholder&#39;s address.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Ssn** | Pointer to **string** | Cardholder&#39;s Social Security Number (SSN). | [optional] 
**State** | Pointer to **string** | State where the cardholder resides.  *NOTE:* &lt;&lt;/core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, Valid two-character abbreviation&gt;&gt; required for KYC verification (US-based cardholders only). | [optional] 
**Token** | Pointer to **string** | Unique identifier of the cardholder. | [optional] 
**UsesParentAccount** | Pointer to **bool** | Indicates whether the child shares balances with the parent (&#x60;true&#x60;), or the child&#39;s balances are independent of the parent (&#x60;false&#x60;).  If set to &#x60;true&#x60;, you must also include a &#x60;parent_token&#x60; in the request. This value cannot be updated. | [optional] [default to false]

## Methods

### NewUserCardHolderUpdateModel

`func NewUserCardHolderUpdateModel() *UserCardHolderUpdateModel`

NewUserCardHolderUpdateModel instantiates a new UserCardHolderUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCardHolderUpdateModelWithDefaults

`func NewUserCardHolderUpdateModelWithDefaults() *UserCardHolderUpdateModel`

NewUserCardHolderUpdateModelWithDefaults instantiates a new UserCardHolderUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderGroupToken

`func (o *UserCardHolderUpdateModel) GetAccountHolderGroupToken() string`

GetAccountHolderGroupToken returns the AccountHolderGroupToken field if non-nil, zero value otherwise.

### GetAccountHolderGroupTokenOk

`func (o *UserCardHolderUpdateModel) GetAccountHolderGroupTokenOk() (*string, bool)`

GetAccountHolderGroupTokenOk returns a tuple with the AccountHolderGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderGroupToken

`func (o *UserCardHolderUpdateModel) SetAccountHolderGroupToken(v string)`

SetAccountHolderGroupToken sets AccountHolderGroupToken field to given value.

### HasAccountHolderGroupToken

`func (o *UserCardHolderUpdateModel) HasAccountHolderGroupToken() bool`

HasAccountHolderGroupToken returns a boolean if a field has been set.

### GetAddress1

`func (o *UserCardHolderUpdateModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *UserCardHolderUpdateModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *UserCardHolderUpdateModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *UserCardHolderUpdateModel) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *UserCardHolderUpdateModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *UserCardHolderUpdateModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *UserCardHolderUpdateModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *UserCardHolderUpdateModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetBirthDate

`func (o *UserCardHolderUpdateModel) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *UserCardHolderUpdateModel) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *UserCardHolderUpdateModel) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *UserCardHolderUpdateModel) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetCity

`func (o *UserCardHolderUpdateModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserCardHolderUpdateModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserCardHolderUpdateModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserCardHolderUpdateModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompany

`func (o *UserCardHolderUpdateModel) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserCardHolderUpdateModel) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserCardHolderUpdateModel) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UserCardHolderUpdateModel) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCorporateCardHolder

`func (o *UserCardHolderUpdateModel) GetCorporateCardHolder() bool`

GetCorporateCardHolder returns the CorporateCardHolder field if non-nil, zero value otherwise.

### GetCorporateCardHolderOk

`func (o *UserCardHolderUpdateModel) GetCorporateCardHolderOk() (*bool, bool)`

GetCorporateCardHolderOk returns a tuple with the CorporateCardHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCardHolder

`func (o *UserCardHolderUpdateModel) SetCorporateCardHolder(v bool)`

SetCorporateCardHolder sets CorporateCardHolder field to given value.

### HasCorporateCardHolder

`func (o *UserCardHolderUpdateModel) HasCorporateCardHolder() bool`

HasCorporateCardHolder returns a boolean if a field has been set.

### GetCountry

`func (o *UserCardHolderUpdateModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserCardHolderUpdateModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserCardHolderUpdateModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserCardHolderUpdateModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *UserCardHolderUpdateModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCardHolderUpdateModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCardHolderUpdateModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCardHolderUpdateModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UserCardHolderUpdateModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCardHolderUpdateModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCardHolderUpdateModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCardHolderUpdateModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGender

`func (o *UserCardHolderUpdateModel) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UserCardHolderUpdateModel) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UserCardHolderUpdateModel) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UserCardHolderUpdateModel) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetHonorific

`func (o *UserCardHolderUpdateModel) GetHonorific() string`

GetHonorific returns the Honorific field if non-nil, zero value otherwise.

### GetHonorificOk

`func (o *UserCardHolderUpdateModel) GetHonorificOk() (*string, bool)`

GetHonorificOk returns a tuple with the Honorific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorific

`func (o *UserCardHolderUpdateModel) SetHonorific(v string)`

SetHonorific sets Honorific field to given value.

### HasHonorific

`func (o *UserCardHolderUpdateModel) HasHonorific() bool`

HasHonorific returns a boolean if a field has been set.

### GetIdCardExpirationDate

`func (o *UserCardHolderUpdateModel) GetIdCardExpirationDate() string`

GetIdCardExpirationDate returns the IdCardExpirationDate field if non-nil, zero value otherwise.

### GetIdCardExpirationDateOk

`func (o *UserCardHolderUpdateModel) GetIdCardExpirationDateOk() (*string, bool)`

GetIdCardExpirationDateOk returns a tuple with the IdCardExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardExpirationDate

`func (o *UserCardHolderUpdateModel) SetIdCardExpirationDate(v string)`

SetIdCardExpirationDate sets IdCardExpirationDate field to given value.

### HasIdCardExpirationDate

`func (o *UserCardHolderUpdateModel) HasIdCardExpirationDate() bool`

HasIdCardExpirationDate returns a boolean if a field has been set.

### GetIdCardNumber

`func (o *UserCardHolderUpdateModel) GetIdCardNumber() string`

GetIdCardNumber returns the IdCardNumber field if non-nil, zero value otherwise.

### GetIdCardNumberOk

`func (o *UserCardHolderUpdateModel) GetIdCardNumberOk() (*string, bool)`

GetIdCardNumberOk returns a tuple with the IdCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardNumber

`func (o *UserCardHolderUpdateModel) SetIdCardNumber(v string)`

SetIdCardNumber sets IdCardNumber field to given value.

### HasIdCardNumber

`func (o *UserCardHolderUpdateModel) HasIdCardNumber() bool`

HasIdCardNumber returns a boolean if a field has been set.

### GetIdentifications

`func (o *UserCardHolderUpdateModel) GetIdentifications() []IdentificationRequestModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *UserCardHolderUpdateModel) GetIdentificationsOk() (*[]IdentificationRequestModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *UserCardHolderUpdateModel) SetIdentifications(v []IdentificationRequestModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *UserCardHolderUpdateModel) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetIpAddress

`func (o *UserCardHolderUpdateModel) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserCardHolderUpdateModel) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserCardHolderUpdateModel) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UserCardHolderUpdateModel) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastName

`func (o *UserCardHolderUpdateModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCardHolderUpdateModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCardHolderUpdateModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCardHolderUpdateModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *UserCardHolderUpdateModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserCardHolderUpdateModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserCardHolderUpdateModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UserCardHolderUpdateModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *UserCardHolderUpdateModel) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UserCardHolderUpdateModel) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UserCardHolderUpdateModel) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UserCardHolderUpdateModel) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNationality

`func (o *UserCardHolderUpdateModel) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UserCardHolderUpdateModel) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UserCardHolderUpdateModel) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *UserCardHolderUpdateModel) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetNotes

`func (o *UserCardHolderUpdateModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UserCardHolderUpdateModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UserCardHolderUpdateModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UserCardHolderUpdateModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetParentToken

`func (o *UserCardHolderUpdateModel) GetParentToken() string`

GetParentToken returns the ParentToken field if non-nil, zero value otherwise.

### GetParentTokenOk

`func (o *UserCardHolderUpdateModel) GetParentTokenOk() (*string, bool)`

GetParentTokenOk returns a tuple with the ParentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentToken

`func (o *UserCardHolderUpdateModel) SetParentToken(v string)`

SetParentToken sets ParentToken field to given value.

### HasParentToken

`func (o *UserCardHolderUpdateModel) HasParentToken() bool`

HasParentToken returns a boolean if a field has been set.

### GetPassportExpirationDate

`func (o *UserCardHolderUpdateModel) GetPassportExpirationDate() string`

GetPassportExpirationDate returns the PassportExpirationDate field if non-nil, zero value otherwise.

### GetPassportExpirationDateOk

`func (o *UserCardHolderUpdateModel) GetPassportExpirationDateOk() (*string, bool)`

GetPassportExpirationDateOk returns a tuple with the PassportExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportExpirationDate

`func (o *UserCardHolderUpdateModel) SetPassportExpirationDate(v string)`

SetPassportExpirationDate sets PassportExpirationDate field to given value.

### HasPassportExpirationDate

`func (o *UserCardHolderUpdateModel) HasPassportExpirationDate() bool`

HasPassportExpirationDate returns a boolean if a field has been set.

### GetPassportNumber

`func (o *UserCardHolderUpdateModel) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *UserCardHolderUpdateModel) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *UserCardHolderUpdateModel) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *UserCardHolderUpdateModel) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### GetPassword

`func (o *UserCardHolderUpdateModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCardHolderUpdateModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCardHolderUpdateModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCardHolderUpdateModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *UserCardHolderUpdateModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserCardHolderUpdateModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserCardHolderUpdateModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserCardHolderUpdateModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *UserCardHolderUpdateModel) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserCardHolderUpdateModel) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserCardHolderUpdateModel) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UserCardHolderUpdateModel) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSsn

`func (o *UserCardHolderUpdateModel) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *UserCardHolderUpdateModel) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *UserCardHolderUpdateModel) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *UserCardHolderUpdateModel) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetState

`func (o *UserCardHolderUpdateModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserCardHolderUpdateModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserCardHolderUpdateModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserCardHolderUpdateModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *UserCardHolderUpdateModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserCardHolderUpdateModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserCardHolderUpdateModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserCardHolderUpdateModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsesParentAccount

`func (o *UserCardHolderUpdateModel) GetUsesParentAccount() bool`

GetUsesParentAccount returns the UsesParentAccount field if non-nil, zero value otherwise.

### GetUsesParentAccountOk

`func (o *UserCardHolderUpdateModel) GetUsesParentAccountOk() (*bool, bool)`

GetUsesParentAccountOk returns a tuple with the UsesParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesParentAccount

`func (o *UserCardHolderUpdateModel) SetUsesParentAccount(v bool)`

SetUsesParentAccount sets UsesParentAccount field to given value.

### HasUsesParentAccount

`func (o *UserCardHolderUpdateModel) HasUsesParentAccount() bool`

HasUsesParentAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


