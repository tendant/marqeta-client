# BeneficialOwnerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dob** | Pointer to **time.Time** | Date of birth of the beneficial owner. | [optional] 
**FirstName** | Pointer to **string** | First name of the beneficial owner. | [optional] 
**Home** | Pointer to [**AddressRequestModel**](AddressRequestModel.md) |  | [optional] 
**LastName** | Pointer to **string** | Last name of the beneficial owner. | [optional] 
**MiddleName** | Pointer to **string** | Middle name of the beneficial owner. | [optional] 
**Phone** | Pointer to **string** | Ten-digit phone number of the beneficial owner. | [optional] 
**Ssn** | Pointer to **string** | Nine-digit Social Security Number (SSN) of the beneficial owner. | [optional] 
**Title** | Pointer to **string** | Title of the beneficial owner. | [optional] 

## Methods

### NewBeneficialOwnerRequest

`func NewBeneficialOwnerRequest() *BeneficialOwnerRequest`

NewBeneficialOwnerRequest instantiates a new BeneficialOwnerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeneficialOwnerRequestWithDefaults

`func NewBeneficialOwnerRequestWithDefaults() *BeneficialOwnerRequest`

NewBeneficialOwnerRequestWithDefaults instantiates a new BeneficialOwnerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDob

`func (o *BeneficialOwnerRequest) GetDob() time.Time`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *BeneficialOwnerRequest) GetDobOk() (*time.Time, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *BeneficialOwnerRequest) SetDob(v time.Time)`

SetDob sets Dob field to given value.

### HasDob

`func (o *BeneficialOwnerRequest) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetFirstName

`func (o *BeneficialOwnerRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BeneficialOwnerRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BeneficialOwnerRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BeneficialOwnerRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHome

`func (o *BeneficialOwnerRequest) GetHome() AddressRequestModel`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *BeneficialOwnerRequest) GetHomeOk() (*AddressRequestModel, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *BeneficialOwnerRequest) SetHome(v AddressRequestModel)`

SetHome sets Home field to given value.

### HasHome

`func (o *BeneficialOwnerRequest) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetLastName

`func (o *BeneficialOwnerRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BeneficialOwnerRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BeneficialOwnerRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BeneficialOwnerRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *BeneficialOwnerRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BeneficialOwnerRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BeneficialOwnerRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BeneficialOwnerRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhone

`func (o *BeneficialOwnerRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BeneficialOwnerRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BeneficialOwnerRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BeneficialOwnerRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSsn

`func (o *BeneficialOwnerRequest) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *BeneficialOwnerRequest) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *BeneficialOwnerRequest) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *BeneficialOwnerRequest) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetTitle

`func (o *BeneficialOwnerRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BeneficialOwnerRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BeneficialOwnerRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BeneficialOwnerRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


