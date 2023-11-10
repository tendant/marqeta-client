# CardHolderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderGroupToken** | Pointer to **string** | Associates the specified account holder group with the cardholder.  Send a &#x60;GET&#x60; request to &#x60;/accountholdergroups&#x60; to retrieve account holder group tokens. | [optional] 
**Active** | Pointer to **bool** | Specifies if the cardholder is in the &#x60;ACTIVE&#x60; state on the Marqeta platform.  *NOTE:* Do not set the value of the &#x60;user.active&#x60; field directly. Instead, use the &#x60;/usertransitions&#x60; endpoints to transition user resources between statuses. For more information on status changes, see &lt;&lt;/core-api/user-transitions#postUsertransitions, Create User Transition&gt;&gt;. | [optional] [default to true]
**Address1** | Pointer to **string** | Cardholder&#39;s address.  *NOTE:* Required for KYC verification (US-based cardholders only). Cannot perform KYC if set to a PO Box. | [optional] 
**Address2** | Pointer to **string** | Additional address information for the cardholder.  *NOTE:* Cannot perform KYC if set to a PO Box. | [optional] 
**BirthDate** | Pointer to **string** | Cardholder&#39;s date of birth.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**City** | Pointer to **string** | City where the cardholder resides.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Company** | Pointer to **string** | Company name. | [optional] 
**CorporateCardHolder** | Pointer to **bool** | Specifies if the cardholder holds a corporate card. | [optional] [default to false]
**Country** | Pointer to **string** | Country where the cardholder resides.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Email** | Pointer to **string** | Valid email address of the cardholder.  This value must be unique among users.  *NOTE:* Required for KYC verification by certain banks (US-based cardholders only). To determine if you must provide an email address, contact your Marqeta representative. | [optional] 
**FirstName** | Pointer to **string** | Cardholder&#39;s first name.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Gender** | Pointer to **string** | Gender of the cardholder. | [optional] 
**Honorific** | Pointer to **string** | Cardholder&#39;s title or prefix: Dr., Miss, Mr., Ms., and so on. | [optional] 
**IdCardExpirationDate** | Pointer to **string** | Expiration date of the cardholder&#39;s identification card. | [optional] 
**IdCardNumber** | Pointer to **string** | Cardholder&#39;s identification card number. | [optional] 
**Identifications** | Pointer to [**[]IdentificationRequestModel**](IdentificationRequestModel.md) | One or more objects containing identifications associated with the cardholder. | [optional] 
**IpAddress** | Pointer to **string** | Cardholder&#39;s IP address. | [optional] 
**LastName** | Pointer to **string** | Cardholder&#39;s last name.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Metadata** | Pointer to **map[string]string** | Associates any additional metadata you provide with the cardholder. | [optional] 
**MiddleName** | Pointer to **string** | Cardholder&#39;s middle name. | [optional] 
**Nationality** | Pointer to **string** | Cardholder&#39;s nationality. | [optional] 
**Notes** | Pointer to **string** | Any additional information pertaining to the cardholder. | [optional] 
**ParentToken** | Pointer to **string** | Unique identifier of the parent user or business resource. Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user resource tokens or to &#x60;/businesses&#x60; to retrieve business resource tokens.  Required if &#x60;uses_parent_account &#x3D; true&#x60;. This user or business is configured as the parent of the current user. | [optional] 
**PassportExpirationDate** | Pointer to **string** | Expiration date of the cardholder&#39;s passport. | [optional] 
**PassportNumber** | Pointer to **string** | Cardholder&#39;s passport number. | [optional] 
**Password** | Pointer to **string** | Password to the cardholder&#39;s user account on the Marqeta platform.  * Must contain at least one numeral + * Must contain at least one lowercase letter + * Must contain at least one uppercase letter + * Must contain at least one of these symbols: &#x60;@ # $ % ! ^ &amp; * ( ) \\ _ + ~ &#x60; - &#x3D; [ ] { } , ; : &#39; \&quot; , . / &lt; &gt; ?&#x60; | [optional] 
**Phone** | Pointer to **string** | Telephone number of the cardholder (including area code), prepended by the &#x60;+&#x60; symbol and the 1- to 3-digit country calling code. Do not include hyphens, spaces, or parentheses.  *NOTE:* Required for KYC verification by certain banks (US-based cardholders only). To determine if you must provide a phone number, contact your Marqeta representative. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the cardholder&#39;s address.  *NOTE:* Required for KYC verification (US-based cardholders only). | [optional] 
**Ssn** | Pointer to **string** | Cardholder&#39;s Social Security Number (SSN). | [optional] 
**State** | Pointer to **string** | State or province where the cardholder resides.  *NOTE:* &lt;&lt;/core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, Valid two-character abbreviation&gt;&gt; required for KYC verification (US-based cardholders only). | [optional] 
**Token** | Pointer to **string** | Unique identifier of the cardholder. If you do not include a token, the system generates one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UsesParentAccount** | Pointer to **bool** | Indicates whether the child shares balances with the parent (&#x60;true&#x60;), or the child&#39;s balances are independent of the parent (&#x60;false&#x60;).  If set to &#x60;true&#x60;, you must also include a &#x60;parent_token&#x60; in the request. This value cannot be updated. | [optional] [default to false]

## Methods

### NewCardHolderModel

`func NewCardHolderModel() *CardHolderModel`

NewCardHolderModel instantiates a new CardHolderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardHolderModelWithDefaults

`func NewCardHolderModelWithDefaults() *CardHolderModel`

NewCardHolderModelWithDefaults instantiates a new CardHolderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderGroupToken

`func (o *CardHolderModel) GetAccountHolderGroupToken() string`

GetAccountHolderGroupToken returns the AccountHolderGroupToken field if non-nil, zero value otherwise.

### GetAccountHolderGroupTokenOk

`func (o *CardHolderModel) GetAccountHolderGroupTokenOk() (*string, bool)`

GetAccountHolderGroupTokenOk returns a tuple with the AccountHolderGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderGroupToken

`func (o *CardHolderModel) SetAccountHolderGroupToken(v string)`

SetAccountHolderGroupToken sets AccountHolderGroupToken field to given value.

### HasAccountHolderGroupToken

`func (o *CardHolderModel) HasAccountHolderGroupToken() bool`

HasAccountHolderGroupToken returns a boolean if a field has been set.

### GetActive

`func (o *CardHolderModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardHolderModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardHolderModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardHolderModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress1

`func (o *CardHolderModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CardHolderModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CardHolderModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *CardHolderModel) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *CardHolderModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CardHolderModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CardHolderModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CardHolderModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetBirthDate

`func (o *CardHolderModel) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *CardHolderModel) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *CardHolderModel) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *CardHolderModel) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetCity

`func (o *CardHolderModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CardHolderModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CardHolderModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CardHolderModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompany

`func (o *CardHolderModel) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CardHolderModel) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CardHolderModel) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CardHolderModel) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCorporateCardHolder

`func (o *CardHolderModel) GetCorporateCardHolder() bool`

GetCorporateCardHolder returns the CorporateCardHolder field if non-nil, zero value otherwise.

### GetCorporateCardHolderOk

`func (o *CardHolderModel) GetCorporateCardHolderOk() (*bool, bool)`

GetCorporateCardHolderOk returns a tuple with the CorporateCardHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCardHolder

`func (o *CardHolderModel) SetCorporateCardHolder(v bool)`

SetCorporateCardHolder sets CorporateCardHolder field to given value.

### HasCorporateCardHolder

`func (o *CardHolderModel) HasCorporateCardHolder() bool`

HasCorporateCardHolder returns a boolean if a field has been set.

### GetCountry

`func (o *CardHolderModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CardHolderModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CardHolderModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CardHolderModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *CardHolderModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CardHolderModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CardHolderModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CardHolderModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *CardHolderModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CardHolderModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CardHolderModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CardHolderModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGender

`func (o *CardHolderModel) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CardHolderModel) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CardHolderModel) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CardHolderModel) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetHonorific

`func (o *CardHolderModel) GetHonorific() string`

GetHonorific returns the Honorific field if non-nil, zero value otherwise.

### GetHonorificOk

`func (o *CardHolderModel) GetHonorificOk() (*string, bool)`

GetHonorificOk returns a tuple with the Honorific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorific

`func (o *CardHolderModel) SetHonorific(v string)`

SetHonorific sets Honorific field to given value.

### HasHonorific

`func (o *CardHolderModel) HasHonorific() bool`

HasHonorific returns a boolean if a field has been set.

### GetIdCardExpirationDate

`func (o *CardHolderModel) GetIdCardExpirationDate() string`

GetIdCardExpirationDate returns the IdCardExpirationDate field if non-nil, zero value otherwise.

### GetIdCardExpirationDateOk

`func (o *CardHolderModel) GetIdCardExpirationDateOk() (*string, bool)`

GetIdCardExpirationDateOk returns a tuple with the IdCardExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardExpirationDate

`func (o *CardHolderModel) SetIdCardExpirationDate(v string)`

SetIdCardExpirationDate sets IdCardExpirationDate field to given value.

### HasIdCardExpirationDate

`func (o *CardHolderModel) HasIdCardExpirationDate() bool`

HasIdCardExpirationDate returns a boolean if a field has been set.

### GetIdCardNumber

`func (o *CardHolderModel) GetIdCardNumber() string`

GetIdCardNumber returns the IdCardNumber field if non-nil, zero value otherwise.

### GetIdCardNumberOk

`func (o *CardHolderModel) GetIdCardNumberOk() (*string, bool)`

GetIdCardNumberOk returns a tuple with the IdCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardNumber

`func (o *CardHolderModel) SetIdCardNumber(v string)`

SetIdCardNumber sets IdCardNumber field to given value.

### HasIdCardNumber

`func (o *CardHolderModel) HasIdCardNumber() bool`

HasIdCardNumber returns a boolean if a field has been set.

### GetIdentifications

`func (o *CardHolderModel) GetIdentifications() []IdentificationRequestModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *CardHolderModel) GetIdentificationsOk() (*[]IdentificationRequestModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *CardHolderModel) SetIdentifications(v []IdentificationRequestModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *CardHolderModel) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetIpAddress

`func (o *CardHolderModel) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CardHolderModel) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CardHolderModel) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CardHolderModel) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastName

`func (o *CardHolderModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CardHolderModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CardHolderModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CardHolderModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *CardHolderModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CardHolderModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CardHolderModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CardHolderModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *CardHolderModel) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *CardHolderModel) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *CardHolderModel) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *CardHolderModel) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNationality

`func (o *CardHolderModel) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *CardHolderModel) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *CardHolderModel) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *CardHolderModel) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetNotes

`func (o *CardHolderModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CardHolderModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CardHolderModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CardHolderModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetParentToken

`func (o *CardHolderModel) GetParentToken() string`

GetParentToken returns the ParentToken field if non-nil, zero value otherwise.

### GetParentTokenOk

`func (o *CardHolderModel) GetParentTokenOk() (*string, bool)`

GetParentTokenOk returns a tuple with the ParentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentToken

`func (o *CardHolderModel) SetParentToken(v string)`

SetParentToken sets ParentToken field to given value.

### HasParentToken

`func (o *CardHolderModel) HasParentToken() bool`

HasParentToken returns a boolean if a field has been set.

### GetPassportExpirationDate

`func (o *CardHolderModel) GetPassportExpirationDate() string`

GetPassportExpirationDate returns the PassportExpirationDate field if non-nil, zero value otherwise.

### GetPassportExpirationDateOk

`func (o *CardHolderModel) GetPassportExpirationDateOk() (*string, bool)`

GetPassportExpirationDateOk returns a tuple with the PassportExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportExpirationDate

`func (o *CardHolderModel) SetPassportExpirationDate(v string)`

SetPassportExpirationDate sets PassportExpirationDate field to given value.

### HasPassportExpirationDate

`func (o *CardHolderModel) HasPassportExpirationDate() bool`

HasPassportExpirationDate returns a boolean if a field has been set.

### GetPassportNumber

`func (o *CardHolderModel) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *CardHolderModel) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *CardHolderModel) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *CardHolderModel) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### GetPassword

`func (o *CardHolderModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CardHolderModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CardHolderModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CardHolderModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *CardHolderModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CardHolderModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CardHolderModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CardHolderModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *CardHolderModel) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CardHolderModel) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CardHolderModel) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CardHolderModel) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSsn

`func (o *CardHolderModel) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *CardHolderModel) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *CardHolderModel) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *CardHolderModel) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetState

`func (o *CardHolderModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardHolderModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardHolderModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CardHolderModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *CardHolderModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardHolderModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardHolderModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardHolderModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsesParentAccount

`func (o *CardHolderModel) GetUsesParentAccount() bool`

GetUsesParentAccount returns the UsesParentAccount field if non-nil, zero value otherwise.

### GetUsesParentAccountOk

`func (o *CardHolderModel) GetUsesParentAccountOk() (*bool, bool)`

GetUsesParentAccountOk returns a tuple with the UsesParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesParentAccount

`func (o *CardHolderModel) SetUsesParentAccount(v bool)`

SetUsesParentAccount sets UsesParentAccount field to given value.

### HasUsesParentAccount

`func (o *CardHolderModel) HasUsesParentAccount() bool`

HasUsesParentAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


