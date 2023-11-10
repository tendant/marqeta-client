# BeneficialOwnerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name of the beneficial owner.  This field is returned if it exists in the resource. | [optional] 
**Getdob** | Pointer to **time.Time** | Date of birth of the beneficial owner.  This field is returned if it exists in the resource. | [optional] 
**Home** | Pointer to [**AddressResponseModel**](AddressResponseModel.md) |  | [optional] 
**LastName** | Pointer to **string** | Last name of the beneficial owner.  This field is returned if it exists in the resource. | [optional] 
**MiddleName** | Pointer to **string** | Middle name of the beneficial owner.  This field is returned if it exists in the resource. | [optional] 
**Phone** | Pointer to **string** | Ten-digit phone number of the beneficial owner.  This field is returned if it exists in the resource. | [optional] 
**Title** | Pointer to **string** | Title of the beneficial owner.  This field is returned if it exists in the resource. | [optional] 

## Methods

### NewBeneficialOwnerResponse

`func NewBeneficialOwnerResponse() *BeneficialOwnerResponse`

NewBeneficialOwnerResponse instantiates a new BeneficialOwnerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeneficialOwnerResponseWithDefaults

`func NewBeneficialOwnerResponseWithDefaults() *BeneficialOwnerResponse`

NewBeneficialOwnerResponseWithDefaults instantiates a new BeneficialOwnerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *BeneficialOwnerResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BeneficialOwnerResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BeneficialOwnerResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BeneficialOwnerResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGetdob

`func (o *BeneficialOwnerResponse) GetGetdob() time.Time`

GetGetdob returns the Getdob field if non-nil, zero value otherwise.

### GetGetdobOk

`func (o *BeneficialOwnerResponse) GetGetdobOk() (*time.Time, bool)`

GetGetdobOk returns a tuple with the Getdob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetdob

`func (o *BeneficialOwnerResponse) SetGetdob(v time.Time)`

SetGetdob sets Getdob field to given value.

### HasGetdob

`func (o *BeneficialOwnerResponse) HasGetdob() bool`

HasGetdob returns a boolean if a field has been set.

### GetHome

`func (o *BeneficialOwnerResponse) GetHome() AddressResponseModel`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *BeneficialOwnerResponse) GetHomeOk() (*AddressResponseModel, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *BeneficialOwnerResponse) SetHome(v AddressResponseModel)`

SetHome sets Home field to given value.

### HasHome

`func (o *BeneficialOwnerResponse) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetLastName

`func (o *BeneficialOwnerResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BeneficialOwnerResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BeneficialOwnerResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BeneficialOwnerResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *BeneficialOwnerResponse) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BeneficialOwnerResponse) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BeneficialOwnerResponse) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BeneficialOwnerResponse) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhone

`func (o *BeneficialOwnerResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BeneficialOwnerResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BeneficialOwnerResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BeneficialOwnerResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTitle

`func (o *BeneficialOwnerResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BeneficialOwnerResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BeneficialOwnerResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BeneficialOwnerResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


