# BusinessProprietorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternativeNames** | Pointer to **string** | Alternate names of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**Dob** | Pointer to **time.Time** | Business proprietor or officer&#39;s date of birth.  This field is returned if it exists in the resource. | [optional] 
**Email** | Pointer to **string** | Email address of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**FirstName** | Pointer to **string** | First name of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**Home** | Pointer to [**AddressResponseModel**](AddressResponseModel.md) |  | [optional] 
**Identifications** | Pointer to [**[]IdentificationResponseModel**](IdentificationResponseModel.md) | One or more objects containing personal identifications of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**LastName** | Pointer to **string** | Last name of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**MiddleName** | Pointer to **string** | Middle name of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**Phone** | Pointer to **string** | Telephone number of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**Ssn** | Pointer to **string** | Social Security Number (SSN) of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 
**Title** | Pointer to **string** | Title of the business proprietor or officer.  This field is returned if it exists in the resource. | [optional] 

## Methods

### NewBusinessProprietorResponse

`func NewBusinessProprietorResponse() *BusinessProprietorResponse`

NewBusinessProprietorResponse instantiates a new BusinessProprietorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessProprietorResponseWithDefaults

`func NewBusinessProprietorResponseWithDefaults() *BusinessProprietorResponse`

NewBusinessProprietorResponseWithDefaults instantiates a new BusinessProprietorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternativeNames

`func (o *BusinessProprietorResponse) GetAlternativeNames() string`

GetAlternativeNames returns the AlternativeNames field if non-nil, zero value otherwise.

### GetAlternativeNamesOk

`func (o *BusinessProprietorResponse) GetAlternativeNamesOk() (*string, bool)`

GetAlternativeNamesOk returns a tuple with the AlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeNames

`func (o *BusinessProprietorResponse) SetAlternativeNames(v string)`

SetAlternativeNames sets AlternativeNames field to given value.

### HasAlternativeNames

`func (o *BusinessProprietorResponse) HasAlternativeNames() bool`

HasAlternativeNames returns a boolean if a field has been set.

### GetDob

`func (o *BusinessProprietorResponse) GetDob() time.Time`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *BusinessProprietorResponse) GetDobOk() (*time.Time, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *BusinessProprietorResponse) SetDob(v time.Time)`

SetDob sets Dob field to given value.

### HasDob

`func (o *BusinessProprietorResponse) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetEmail

`func (o *BusinessProprietorResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BusinessProprietorResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BusinessProprietorResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BusinessProprietorResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *BusinessProprietorResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BusinessProprietorResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BusinessProprietorResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BusinessProprietorResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHome

`func (o *BusinessProprietorResponse) GetHome() AddressResponseModel`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *BusinessProprietorResponse) GetHomeOk() (*AddressResponseModel, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *BusinessProprietorResponse) SetHome(v AddressResponseModel)`

SetHome sets Home field to given value.

### HasHome

`func (o *BusinessProprietorResponse) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIdentifications

`func (o *BusinessProprietorResponse) GetIdentifications() []IdentificationResponseModel`

GetIdentifications returns the Identifications field if non-nil, zero value otherwise.

### GetIdentificationsOk

`func (o *BusinessProprietorResponse) GetIdentificationsOk() (*[]IdentificationResponseModel, bool)`

GetIdentificationsOk returns a tuple with the Identifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifications

`func (o *BusinessProprietorResponse) SetIdentifications(v []IdentificationResponseModel)`

SetIdentifications sets Identifications field to given value.

### HasIdentifications

`func (o *BusinessProprietorResponse) HasIdentifications() bool`

HasIdentifications returns a boolean if a field has been set.

### GetLastName

`func (o *BusinessProprietorResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BusinessProprietorResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BusinessProprietorResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BusinessProprietorResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *BusinessProprietorResponse) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BusinessProprietorResponse) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BusinessProprietorResponse) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BusinessProprietorResponse) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhone

`func (o *BusinessProprietorResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BusinessProprietorResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BusinessProprietorResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BusinessProprietorResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSsn

`func (o *BusinessProprietorResponse) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *BusinessProprietorResponse) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *BusinessProprietorResponse) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *BusinessProprietorResponse) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetTitle

`func (o *BusinessProprietorResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BusinessProprietorResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BusinessProprietorResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BusinessProprietorResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


