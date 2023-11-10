# CardHolderAddressUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the address is active. | [optional] [default to true]
**Address1** | Pointer to **string** | Street name and number of the address. | [optional] 
**Address2** | Pointer to **string** | Additional address information. | [optional] 
**City** | Pointer to **string** | City of the address. | [optional] 
**Country** | Pointer to **string** | Country of the address. | [optional] 
**FirstName** | Pointer to **string** | First name or given name of the account holder. | [optional] 
**IsDefaultAddress** | Pointer to **bool** | A value of &#x60;true&#x60; specifies that this address is the default address used by the account holder&#39;s funding source. If this is the account holder&#39;s only address, it is used as the default regardless of this field&#39;s setting. | [optional] [default to false]
**LastName** | Pointer to **string** | Last name or family name of the account holder. | [optional] 
**Phone** | Pointer to **string** | Telephone number of the account holder. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the address. | [optional] 
**State** | Pointer to **string** | Two-character state, province, or territorial abbreviation.  For a complete list, see &lt;&lt;/core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, Valid state, provincial, and territorial abbreviations&gt;&gt;. | [optional] 
**Zip** | Pointer to **string** | United States ZIP code of the address. | [optional] 

## Methods

### NewCardHolderAddressUpdateModel

`func NewCardHolderAddressUpdateModel() *CardHolderAddressUpdateModel`

NewCardHolderAddressUpdateModel instantiates a new CardHolderAddressUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardHolderAddressUpdateModelWithDefaults

`func NewCardHolderAddressUpdateModelWithDefaults() *CardHolderAddressUpdateModel`

NewCardHolderAddressUpdateModelWithDefaults instantiates a new CardHolderAddressUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardHolderAddressUpdateModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardHolderAddressUpdateModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardHolderAddressUpdateModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardHolderAddressUpdateModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress1

`func (o *CardHolderAddressUpdateModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CardHolderAddressUpdateModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CardHolderAddressUpdateModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *CardHolderAddressUpdateModel) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *CardHolderAddressUpdateModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CardHolderAddressUpdateModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CardHolderAddressUpdateModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CardHolderAddressUpdateModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *CardHolderAddressUpdateModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CardHolderAddressUpdateModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CardHolderAddressUpdateModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CardHolderAddressUpdateModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CardHolderAddressUpdateModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CardHolderAddressUpdateModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CardHolderAddressUpdateModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CardHolderAddressUpdateModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFirstName

`func (o *CardHolderAddressUpdateModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CardHolderAddressUpdateModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CardHolderAddressUpdateModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CardHolderAddressUpdateModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetIsDefaultAddress

`func (o *CardHolderAddressUpdateModel) GetIsDefaultAddress() bool`

GetIsDefaultAddress returns the IsDefaultAddress field if non-nil, zero value otherwise.

### GetIsDefaultAddressOk

`func (o *CardHolderAddressUpdateModel) GetIsDefaultAddressOk() (*bool, bool)`

GetIsDefaultAddressOk returns a tuple with the IsDefaultAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAddress

`func (o *CardHolderAddressUpdateModel) SetIsDefaultAddress(v bool)`

SetIsDefaultAddress sets IsDefaultAddress field to given value.

### HasIsDefaultAddress

`func (o *CardHolderAddressUpdateModel) HasIsDefaultAddress() bool`

HasIsDefaultAddress returns a boolean if a field has been set.

### GetLastName

`func (o *CardHolderAddressUpdateModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CardHolderAddressUpdateModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CardHolderAddressUpdateModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CardHolderAddressUpdateModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *CardHolderAddressUpdateModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CardHolderAddressUpdateModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CardHolderAddressUpdateModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CardHolderAddressUpdateModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *CardHolderAddressUpdateModel) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CardHolderAddressUpdateModel) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CardHolderAddressUpdateModel) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CardHolderAddressUpdateModel) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *CardHolderAddressUpdateModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardHolderAddressUpdateModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardHolderAddressUpdateModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CardHolderAddressUpdateModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *CardHolderAddressUpdateModel) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CardHolderAddressUpdateModel) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CardHolderAddressUpdateModel) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CardHolderAddressUpdateModel) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


