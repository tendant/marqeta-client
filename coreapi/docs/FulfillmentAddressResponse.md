# FulfillmentAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | Number and street of the address.  This field is returned if it exists in the resource. | [optional] 
**Address2** | Pointer to **string** | Additional address information.  This field is returned if it exists in the resource. | [optional] 
**City** | Pointer to **string** | City of the address.  This field is returned if it exists in the resource. | [optional] 
**Country** | Pointer to **string** | Country of the address.  This field is returned if it exists in the resource. | [optional] 
**FirstName** | Pointer to **string** | First name of the addressee.  This field is returned if it exists in the resource. | [optional] 
**LastName** | Pointer to **string** | Last name of the addressee.  This field is returned if it exists in the resource. | [optional] 
**MiddleName** | Pointer to **string** | Middle name of the addressee.  This field is returned if it exists in the resource. | [optional] 
**Phone** | Pointer to **string** | Telephone number of the addressee.  This field is returned if it exists in the resource. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the address.  This field is returned if it exists in the resource. | [optional] 
**State** | Pointer to **string** | State or province of the address.  This field is returned if it exists in the resource. | [optional] 
**Zip** | Pointer to **string** | United States ZIP code of the address.  This field is returned if it exists in the resource. | [optional] 

## Methods

### NewFulfillmentAddressResponse

`func NewFulfillmentAddressResponse() *FulfillmentAddressResponse`

NewFulfillmentAddressResponse instantiates a new FulfillmentAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentAddressResponseWithDefaults

`func NewFulfillmentAddressResponseWithDefaults() *FulfillmentAddressResponse`

NewFulfillmentAddressResponseWithDefaults instantiates a new FulfillmentAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *FulfillmentAddressResponse) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *FulfillmentAddressResponse) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *FulfillmentAddressResponse) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *FulfillmentAddressResponse) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *FulfillmentAddressResponse) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *FulfillmentAddressResponse) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *FulfillmentAddressResponse) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *FulfillmentAddressResponse) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *FulfillmentAddressResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FulfillmentAddressResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FulfillmentAddressResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *FulfillmentAddressResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *FulfillmentAddressResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FulfillmentAddressResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FulfillmentAddressResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FulfillmentAddressResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFirstName

`func (o *FulfillmentAddressResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FulfillmentAddressResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FulfillmentAddressResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *FulfillmentAddressResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *FulfillmentAddressResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FulfillmentAddressResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FulfillmentAddressResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *FulfillmentAddressResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *FulfillmentAddressResponse) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *FulfillmentAddressResponse) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *FulfillmentAddressResponse) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *FulfillmentAddressResponse) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetPhone

`func (o *FulfillmentAddressResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *FulfillmentAddressResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *FulfillmentAddressResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *FulfillmentAddressResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *FulfillmentAddressResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *FulfillmentAddressResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *FulfillmentAddressResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *FulfillmentAddressResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *FulfillmentAddressResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FulfillmentAddressResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FulfillmentAddressResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FulfillmentAddressResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *FulfillmentAddressResponse) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *FulfillmentAddressResponse) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *FulfillmentAddressResponse) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *FulfillmentAddressResponse) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


