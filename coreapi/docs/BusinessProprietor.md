# BusinessProprietor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternativeNames** | Pointer to **string** | Alternate names of the business proprietor or officer. | [optional] 
**Dob** | Pointer to **time.Time** | Business proprietor or officer&#39;s date of birth.  This field is required for KYC verification (US-based accounts only). | [optional] 
**Email** | Pointer to **string** | Email address of the business proprietor or officer. | [optional] 
**FirstName** | **string** | First name of business proprietor or officer. | 
**Home** | Pointer to [**AddressRequestModel**](AddressRequestModel.md) |  | [optional] 
**Identifications** | Pointer to [**[]IdentificationRequestModel**](IdentificationRequestModel.md) | One or more objects containing personal identifications of the business proprietor or officer. | [optional] 
**LastName** | **string** | Last name of business proprietor or officer. | 
**MiddleName** | Pointer to **string** | Middle name of business proprietor or officer. | [optional] 
**Phone** | Pointer to **string** | Telephone number of the business proprietor or officer. | [optional] 
**Ssn** | Pointer to **string** | Social Security Number (SSN) of the business proprietor or officer. | [optional] 
**Title** | Pointer to **string** | Title of business proprietor or officer. | [optional] 

## Methods

### NewBusinessProprietor

`func NewBusinessProprietor(firstName string, lastName string, ) *BusinessProprietor`

NewBusinessProprietor instantiates a new BusinessProprietor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessProprietorWithDefaults

`func NewBusinessProprietorWithDefaults() *BusinessProprietor`

NewBusinessProprietorWithDefaults instantiates a new BusinessProprietor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternativeNames

`func (o *BusinessProprietor) GetAlternativeNames() string`

GetAlternativeNames returns the AlternativeNames field if non-nil, zero value otherwise.

### GetAlternativeNamesOk

`func (o *BusinessProprietor) GetAlternativeNamesOk() (*string, bool)`

GetAlternativeNamesOk returns a tuple with the AlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeNames

`func (o *BusinessProprietor) SetAlternativeNames(v string)`

SetAlternativeNames sets AlternativeNames field to given value.

### HasAlternativeNames

`func (o *BusinessProprietor) HasAlternativeNames() bool`

HasAlternativeNames returns a boolean if a field has been set.

### GetDob

`func (o *BusinessProprietor) GetDob() time.Time`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *BusinessProprietor) GetDobOk() (*time.Time, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *BusinessProprietor) SetDob(v time.Time)`

SetDob sets Dob field to given value.

### HasDob

`func (o *BusinessProprietor) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetEmail

`func (o *BusinessProprietor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BusinessProprietor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BusinessProprietor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BusinessProprietor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *BusinessProprietor) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BusinessProprietor) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BusinessProprietor) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetHome

`func (o *BusinessProprietor) GetHome() AddressRequestModel`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *BusinessProprietor) GetHomeOk() (*AddressRequestModel, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *BusinessProprietor) SetHome(v AddressRequestModel)`

SetHome sets Home field to given value.

### HasHome

`func (o *BusinessProprietor) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIdentifications

`func (o *BusinessProprietor) GetIdentifications() []IdentificationRequestModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *BusinessProprietor) GetIdentificationsOk() (*[]IdentificationRequestModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *BusinessProprietor) SetIdentifications(v []IdentificationRequestModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *BusinessProprietor) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetLastName

`func (o *BusinessProprietor) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BusinessProprietor) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BusinessProprietor) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMiddleName

`func (o *BusinessProprietor) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BusinessProprietor) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BusinessProprietor) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BusinessProprietor) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhone

`func (o *BusinessProprietor) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BusinessProprietor) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BusinessProprietor) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BusinessProprietor) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSsn

`func (o *BusinessProprietor) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *BusinessProprietor) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *BusinessProprietor) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *BusinessProprietor) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetTitle

`func (o *BusinessProprietor) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BusinessProprietor) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BusinessProprietor) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BusinessProprietor) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


