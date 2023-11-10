# CardholderAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the address is active. | [optional] [default to false]
**Address1** | **string** | Street address of the funding source. | 
**Address2** | Pointer to **string** | Additional address information for the funding source. This field is returned if it exists in the resource. | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder associated with the address. This field is returned if it exists in the resource. It is required if &#39;user_token&#39; is not specified. | [optional] 
**City** | **string** | City of the funding source. | 
**Country** | **string** | Country of the funding source. | 
**CreatedTime** | **time.Time** | Date and time when the address was created, in UTC. | 
**FirstName** | **string** | First name of the account holder associated with the funding source. | 
**IsDefaultAddress** | Pointer to **bool** | Whether this address is the default address used by the funding source. | [optional] [default to false]
**LastModifiedTime** | **time.Time** | Date and time when the address was last modified, in UTC. This field is returned if it exists in the resource. | 
**LastName** | **string** | Last name of the account holder associated with the funding source. | 
**Phone** | Pointer to **string** | Phone number of the funding source. This field is returned if it exists in the resource. | [optional] 
**PostalCode** | **string** | Postal code of the funding source. | 
**State** | **string** | Two-character state, province, or territorial abbreviation. For the complete list, see &lt;&lt;/core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, Valid state, provincial, and territorial abbreviations&gt;&gt;. | 
**Token** | **string** | Unique identifier of the &#x60;funding_source_address&#x60; object. | 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder associated with the address. This field is returned if it exists in the resource. It is required if &#39;business_token&#39; is not specified. | [optional] 
**Zip** | **string** | United States ZIP code of the funding source. | 

## Methods

### NewCardholderAddressResponse

`func NewCardholderAddressResponse(address1 string, city string, country string, createdTime time.Time, firstName string, lastModifiedTime time.Time, lastName string, postalCode string, state string, token string, zip string, ) *CardholderAddressResponse`

NewCardholderAddressResponse instantiates a new CardholderAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardholderAddressResponseWithDefaults

`func NewCardholderAddressResponseWithDefaults() *CardholderAddressResponse`

NewCardholderAddressResponseWithDefaults instantiates a new CardholderAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardholderAddressResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardholderAddressResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardholderAddressResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardholderAddressResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress1

`func (o *CardholderAddressResponse) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CardholderAddressResponse) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CardholderAddressResponse) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *CardholderAddressResponse) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CardholderAddressResponse) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CardholderAddressResponse) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CardholderAddressResponse) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetBusinessToken

`func (o *CardholderAddressResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *CardholderAddressResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *CardholderAddressResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *CardholderAddressResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCity

`func (o *CardholderAddressResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CardholderAddressResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CardholderAddressResponse) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *CardholderAddressResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CardholderAddressResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CardholderAddressResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCreatedTime

`func (o *CardholderAddressResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CardholderAddressResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CardholderAddressResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetFirstName

`func (o *CardholderAddressResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CardholderAddressResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CardholderAddressResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetIsDefaultAddress

`func (o *CardholderAddressResponse) GetIsDefaultAddress() bool`

GetIsDefaultAddress returns the IsDefaultAddress field if non-nil, zero value otherwise.

### GetIsDefaultAddressOk

`func (o *CardholderAddressResponse) GetIsDefaultAddressOk() (*bool, bool)`

GetIsDefaultAddressOk returns a tuple with the IsDefaultAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAddress

`func (o *CardholderAddressResponse) SetIsDefaultAddress(v bool)`

SetIsDefaultAddress sets IsDefaultAddress field to given value.

### HasIsDefaultAddress

`func (o *CardholderAddressResponse) HasIsDefaultAddress() bool`

HasIsDefaultAddress returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardholderAddressResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardholderAddressResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardholderAddressResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastName

`func (o *CardholderAddressResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CardholderAddressResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CardholderAddressResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPhone

`func (o *CardholderAddressResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CardholderAddressResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CardholderAddressResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CardholderAddressResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *CardholderAddressResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CardholderAddressResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CardholderAddressResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetState

`func (o *CardholderAddressResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardholderAddressResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardholderAddressResponse) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *CardholderAddressResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardholderAddressResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardholderAddressResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *CardholderAddressResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardholderAddressResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardholderAddressResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *CardholderAddressResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetZip

`func (o *CardholderAddressResponse) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CardholderAddressResponse) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CardholderAddressResponse) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


