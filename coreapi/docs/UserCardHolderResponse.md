# UserCardHolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderGroupToken** | Pointer to **string** | Associates the specified account holder group with the cardholder. | [optional] 
**Active** | Pointer to **bool** | Specifies if the cardholder is in the &#x60;ACTIVE&#x60; state on the Marqeta platform. | [optional] [default to false]
**Address1** | Pointer to **string** | Cardholder&#39;s address. | [optional] 
**Address2** | Pointer to **string** | Additional address information for the cardholder. | [optional] 
**Authentication** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**BirthDate** | Pointer to **string** | Cardholder&#39;s date of birth. | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the business resource. | [optional] 
**City** | Pointer to **string** | City where the cardholder resides. | [optional] 
**Company** | Pointer to **string** | Company name. | [optional] 
**CorporateCardHolder** | Pointer to **bool** | Specifies if the cardholder holds a corporate card. | [optional] [default to false]
**Country** | Pointer to **string** | Country where the cardholder resides. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**Email** | Pointer to **string** | Valid email address of the cardholder. | [optional] 
**FirstName** | Pointer to **string** | Cardholder&#39;s first name. | [optional] 
**Gender** | Pointer to **string** | Gender of the cardholder. | [optional] 
**Honorific** | Pointer to **string** | Cardholder&#39;s title or prefix: Dr., Miss, Mr., Ms., and so on. | [optional] 
**IdCardExpirationDate** | Pointer to **string** | Expiration date of the cardholder&#39;s identification. | [optional] [readonly] 
**IdCardNumber** | Pointer to **string** | Cardholder&#39;s identification card number. | [optional] 
**Identifications** | Pointer to [**[]IdentificationResponseModel**](IdentificationResponseModel.md) | One or more objects containing identifications associated with the cardholder. | [optional] 
**IpAddress** | Pointer to **string** | Cardholder&#39;s IP address. | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the resource was last updated, in UTC. | 
**LastName** | Pointer to **string** | Cardholder&#39;s last name. | [optional] 
**Metadata** | Pointer to **map[string]string** | Associates any additional metadata you provide with the cardholder. | [optional] 
**MiddleName** | Pointer to **string** | Cardholder&#39;s middle name. | [optional] 
**Nationality** | Pointer to **string** | Cardholder&#39;s nationality. | [optional] 
**Notes** | Pointer to **string** | Any additional information pertaining to the cardholder. | [optional] 
**ParentToken** | Pointer to **string** | Unique identifier of the parent user or business resource. | [optional] 
**PassportExpirationDate** | Pointer to **string** | Expiration date of the cardholder&#39;s passport. | [optional] [readonly] 
**PassportNumber** | Pointer to **string** | Cardholder&#39;s passport number. | [optional] 
**Password** | Pointer to **string** | Password to the cardholder&#39;s user account on the Marqeta platform. | [optional] 
**Phone** | Pointer to **string** | Cardholder&#39;s telephone number. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the cardholder&#39;s address. | [optional] 
**Ssn** | Pointer to **string** | Cardholder&#39;s Social Security Number (SSN). | [optional] 
**State** | Pointer to **string** | State or province where the cardholder resides. | [optional] 
**Status** | Pointer to **string** | Specifies the status of the cardholder on the Marqeta platform. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the cardholder. | [optional] 
**UsesParentAccount** | Pointer to **bool** | Indicates whether the child shares balances with the parent (&#x60;true&#x60;), or the child&#39;s balances are independent of the parent (&#x60;false&#x60;). | [optional] [default to false]
**Zip** | Pointer to **string** | United States ZIP code of the cardholder&#39;s address. | [optional] 

## Methods

### NewUserCardHolderResponse

`func NewUserCardHolderResponse(createdTime time.Time, lastModifiedTime time.Time, ) *UserCardHolderResponse`

NewUserCardHolderResponse instantiates a new UserCardHolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCardHolderResponseWithDefaults

`func NewUserCardHolderResponseWithDefaults() *UserCardHolderResponse`

NewUserCardHolderResponseWithDefaults instantiates a new UserCardHolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderGroupToken

`func (o *UserCardHolderResponse) GetAccountHolderGroupToken() string`

GetAccountHolderGroupToken returns the AccountHolderGroupToken field if non-nil, zero value otherwise.

### GetAccountHolderGroupTokenOk

`func (o *UserCardHolderResponse) GetAccountHolderGroupTokenOk() (*string, bool)`

GetAccountHolderGroupTokenOk returns a tuple with the AccountHolderGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderGroupToken

`func (o *UserCardHolderResponse) SetAccountHolderGroupToken(v string)`

SetAccountHolderGroupToken sets AccountHolderGroupToken field to given value.

### HasAccountHolderGroupToken

`func (o *UserCardHolderResponse) HasAccountHolderGroupToken() bool`

HasAccountHolderGroupToken returns a boolean if a field has been set.

### GetActive

`func (o *UserCardHolderResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserCardHolderResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserCardHolderResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserCardHolderResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress1

`func (o *UserCardHolderResponse) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *UserCardHolderResponse) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *UserCardHolderResponse) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *UserCardHolderResponse) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *UserCardHolderResponse) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *UserCardHolderResponse) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *UserCardHolderResponse) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *UserCardHolderResponse) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAuthentication

`func (o *UserCardHolderResponse) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *UserCardHolderResponse) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *UserCardHolderResponse) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *UserCardHolderResponse) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetBirthDate

`func (o *UserCardHolderResponse) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *UserCardHolderResponse) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *UserCardHolderResponse) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *UserCardHolderResponse) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetBusinessToken

`func (o *UserCardHolderResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *UserCardHolderResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *UserCardHolderResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *UserCardHolderResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCity

`func (o *UserCardHolderResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserCardHolderResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserCardHolderResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserCardHolderResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompany

`func (o *UserCardHolderResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserCardHolderResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserCardHolderResponse) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UserCardHolderResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCorporateCardHolder

`func (o *UserCardHolderResponse) GetCorporateCardHolder() bool`

GetCorporateCardHolder returns the CorporateCardHolder field if non-nil, zero value otherwise.

### GetCorporateCardHolderOk

`func (o *UserCardHolderResponse) GetCorporateCardHolderOk() (*bool, bool)`

GetCorporateCardHolderOk returns a tuple with the CorporateCardHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateCardHolder

`func (o *UserCardHolderResponse) SetCorporateCardHolder(v bool)`

SetCorporateCardHolder sets CorporateCardHolder field to given value.

### HasCorporateCardHolder

`func (o *UserCardHolderResponse) HasCorporateCardHolder() bool`

HasCorporateCardHolder returns a boolean if a field has been set.

### GetCountry

`func (o *UserCardHolderResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserCardHolderResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserCardHolderResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserCardHolderResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedTime

`func (o *UserCardHolderResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *UserCardHolderResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *UserCardHolderResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEmail

`func (o *UserCardHolderResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCardHolderResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCardHolderResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCardHolderResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UserCardHolderResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCardHolderResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCardHolderResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCardHolderResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGender

`func (o *UserCardHolderResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UserCardHolderResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UserCardHolderResponse) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UserCardHolderResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetHonorific

`func (o *UserCardHolderResponse) GetHonorific() string`

GetHonorific returns the Honorific field if non-nil, zero value otherwise.

### GetHonorificOk

`func (o *UserCardHolderResponse) GetHonorificOk() (*string, bool)`

GetHonorificOk returns a tuple with the Honorific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorific

`func (o *UserCardHolderResponse) SetHonorific(v string)`

SetHonorific sets Honorific field to given value.

### HasHonorific

`func (o *UserCardHolderResponse) HasHonorific() bool`

HasHonorific returns a boolean if a field has been set.

### GetIdCardExpirationDate

`func (o *UserCardHolderResponse) GetIdCardExpirationDate() string`

GetIdCardExpirationDate returns the IdCardExpirationDate field if non-nil, zero value otherwise.

### GetIdCardExpirationDateOk

`func (o *UserCardHolderResponse) GetIdCardExpirationDateOk() (*string, bool)`

GetIdCardExpirationDateOk returns a tuple with the IdCardExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardExpirationDate

`func (o *UserCardHolderResponse) SetIdCardExpirationDate(v string)`

SetIdCardExpirationDate sets IdCardExpirationDate field to given value.

### HasIdCardExpirationDate

`func (o *UserCardHolderResponse) HasIdCardExpirationDate() bool`

HasIdCardExpirationDate returns a boolean if a field has been set.

### GetIdCardNumber

`func (o *UserCardHolderResponse) GetIdCardNumber() string`

GetIdCardNumber returns the IdCardNumber field if non-nil, zero value otherwise.

### GetIdCardNumberOk

`func (o *UserCardHolderResponse) GetIdCardNumberOk() (*string, bool)`

GetIdCardNumberOk returns a tuple with the IdCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardNumber

`func (o *UserCardHolderResponse) SetIdCardNumber(v string)`

SetIdCardNumber sets IdCardNumber field to given value.

### HasIdCardNumber

`func (o *UserCardHolderResponse) HasIdCardNumber() bool`

HasIdCardNumber returns a boolean if a field has been set.

### GetIdentifications

`func (o *UserCardHolderResponse) GetIdentifications() []IdentificationResponseModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *UserCardHolderResponse) GetIdentificationsOk() (*[]IdentificationResponseModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *UserCardHolderResponse) SetIdentifications(v []IdentificationResponseModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *UserCardHolderResponse) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetIpAddress

`func (o *UserCardHolderResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserCardHolderResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserCardHolderResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UserCardHolderResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *UserCardHolderResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *UserCardHolderResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *UserCardHolderResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastName

`func (o *UserCardHolderResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCardHolderResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCardHolderResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCardHolderResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMetadata

`func (o *UserCardHolderResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserCardHolderResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserCardHolderResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UserCardHolderResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMiddleName

`func (o *UserCardHolderResponse) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UserCardHolderResponse) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UserCardHolderResponse) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UserCardHolderResponse) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNationality

`func (o *UserCardHolderResponse) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *UserCardHolderResponse) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *UserCardHolderResponse) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *UserCardHolderResponse) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetNotes

`func (o *UserCardHolderResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UserCardHolderResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UserCardHolderResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UserCardHolderResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetParentToken

`func (o *UserCardHolderResponse) GetParentToken() string`

GetParentToken returns the ParentToken field if non-nil, zero value otherwise.

### GetParentTokenOk

`func (o *UserCardHolderResponse) GetParentTokenOk() (*string, bool)`

GetParentTokenOk returns a tuple with the ParentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentToken

`func (o *UserCardHolderResponse) SetParentToken(v string)`

SetParentToken sets ParentToken field to given value.

### HasParentToken

`func (o *UserCardHolderResponse) HasParentToken() bool`

HasParentToken returns a boolean if a field has been set.

### GetPassportExpirationDate

`func (o *UserCardHolderResponse) GetPassportExpirationDate() string`

GetPassportExpirationDate returns the PassportExpirationDate field if non-nil, zero value otherwise.

### GetPassportExpirationDateOk

`func (o *UserCardHolderResponse) GetPassportExpirationDateOk() (*string, bool)`

GetPassportExpirationDateOk returns a tuple with the PassportExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportExpirationDate

`func (o *UserCardHolderResponse) SetPassportExpirationDate(v string)`

SetPassportExpirationDate sets PassportExpirationDate field to given value.

### HasPassportExpirationDate

`func (o *UserCardHolderResponse) HasPassportExpirationDate() bool`

HasPassportExpirationDate returns a boolean if a field has been set.

### GetPassportNumber

`func (o *UserCardHolderResponse) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *UserCardHolderResponse) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *UserCardHolderResponse) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *UserCardHolderResponse) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### GetPassword

`func (o *UserCardHolderResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCardHolderResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCardHolderResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCardHolderResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhone

`func (o *UserCardHolderResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserCardHolderResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserCardHolderResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserCardHolderResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *UserCardHolderResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserCardHolderResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserCardHolderResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UserCardHolderResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSsn

`func (o *UserCardHolderResponse) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *UserCardHolderResponse) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *UserCardHolderResponse) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *UserCardHolderResponse) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetState

`func (o *UserCardHolderResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserCardHolderResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserCardHolderResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserCardHolderResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *UserCardHolderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserCardHolderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserCardHolderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserCardHolderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *UserCardHolderResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserCardHolderResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserCardHolderResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserCardHolderResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsesParentAccount

`func (o *UserCardHolderResponse) GetUsesParentAccount() bool`

GetUsesParentAccount returns the UsesParentAccount field if non-nil, zero value otherwise.

### GetUsesParentAccountOk

`func (o *UserCardHolderResponse) GetUsesParentAccountOk() (*bool, bool)`

GetUsesParentAccountOk returns a tuple with the UsesParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesParentAccount

`func (o *UserCardHolderResponse) SetUsesParentAccount(v bool)`

SetUsesParentAccount sets UsesParentAccount field to given value.

### HasUsesParentAccount

`func (o *UserCardHolderResponse) HasUsesParentAccount() bool`

HasUsesParentAccount returns a boolean if a field has been set.

### GetZip

`func (o *UserCardHolderResponse) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *UserCardHolderResponse) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *UserCardHolderResponse) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *UserCardHolderResponse) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


