# CardHolderAddressModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the address is active. | [optional] [default to true]
**Address1** | **string** | Street name and number of the address. | 
**Address2** | Pointer to **string** | Additional address information. | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the business account holder. This token is required if a &#x60;user_token&#x60; is not specified. | [optional] 
**City** | **string** | City of the address. | 
**Country** | **string** | Country of the address. | 
**FirstName** | **string** | First name or given name of the account holder. | 
**IsDefaultAddress** | Pointer to **bool** | A value of &#x60;true&#x60; specifies that this address is the default address used by the account holder&#39;s funding source. If this is the account holder&#39;s only address, it is used as the default regardless of this field&#39;s setting. | [optional] [default to false]
**LastName** | **string** | Last name or family name of the account holder. | 
**Phone** | Pointer to **string** | Telephone number of the account holder. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the address. | [optional] 
**State** | **string** | Two-character state, province, or territorial abbreviation.  For a complete list of valid state and province abbreviations, see &lt;&lt;/core-api/kyc-verification#_valid_state_provincial_and_territorial_abbreviations, Valid state, provincial, and territorial abbreviations&gt;&gt;. | 
**Token** | Pointer to **string** | Unique identifier of the address. If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user account holder. This token is required if a &#x60;business_token&#x60; is not specified. | [optional] 
**Zip** | Pointer to **string** | United States ZIP code. This field is required if &#x60;postal_code&#x60; is not specified. | [optional] 

## Methods

### NewCardHolderAddressModel

`func NewCardHolderAddressModel(address1 string, city string, country string, firstName string, lastName string, state string, ) *CardHolderAddressModel`

NewCardHolderAddressModel instantiates a new CardHolderAddressModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardHolderAddressModelWithDefaults

`func NewCardHolderAddressModelWithDefaults() *CardHolderAddressModel`

NewCardHolderAddressModelWithDefaults instantiates a new CardHolderAddressModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardHolderAddressModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardHolderAddressModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardHolderAddressModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardHolderAddressModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddress1

`func (o *CardHolderAddressModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CardHolderAddressModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CardHolderAddressModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *CardHolderAddressModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CardHolderAddressModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CardHolderAddressModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CardHolderAddressModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetBusinessToken

`func (o *CardHolderAddressModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *CardHolderAddressModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *CardHolderAddressModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *CardHolderAddressModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCity

`func (o *CardHolderAddressModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CardHolderAddressModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CardHolderAddressModel) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *CardHolderAddressModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CardHolderAddressModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CardHolderAddressModel) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetFirstName

`func (o *CardHolderAddressModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CardHolderAddressModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CardHolderAddressModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetIsDefaultAddress

`func (o *CardHolderAddressModel) GetIsDefaultAddress() bool`

GetIsDefaultAddress returns the IsDefaultAddress field if non-nil, zero value otherwise.

### GetIsDefaultAddressOk

`func (o *CardHolderAddressModel) GetIsDefaultAddressOk() (*bool, bool)`

GetIsDefaultAddressOk returns a tuple with the IsDefaultAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAddress

`func (o *CardHolderAddressModel) SetIsDefaultAddress(v bool)`

SetIsDefaultAddress sets IsDefaultAddress field to given value.

### HasIsDefaultAddress

`func (o *CardHolderAddressModel) HasIsDefaultAddress() bool`

HasIsDefaultAddress returns a boolean if a field has been set.

### GetLastName

`func (o *CardHolderAddressModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CardHolderAddressModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CardHolderAddressModel) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPhone

`func (o *CardHolderAddressModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CardHolderAddressModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CardHolderAddressModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CardHolderAddressModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *CardHolderAddressModel) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CardHolderAddressModel) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CardHolderAddressModel) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CardHolderAddressModel) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *CardHolderAddressModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardHolderAddressModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardHolderAddressModel) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *CardHolderAddressModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardHolderAddressModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardHolderAddressModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardHolderAddressModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *CardHolderAddressModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardHolderAddressModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardHolderAddressModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *CardHolderAddressModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetZip

`func (o *CardHolderAddressModel) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CardHolderAddressModel) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CardHolderAddressModel) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CardHolderAddressModel) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


