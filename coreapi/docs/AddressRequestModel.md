# AddressRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | Street name and number of the address.  This field is required for KYC verification (US-based accounts only). Cannot perform KYC if set to a PO Box. | [optional] 
**Address2** | Pointer to **string** | Additional address information. | [optional] 
**City** | Pointer to **string** | City of the address.  This field is required for KYC verification (US-based accounts only). | [optional] 
**Country** | Pointer to **string** | Country of the address. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of the address.  This field is required for KYC verification (US-based accounts only). | [optional] 
**State** | Pointer to **string** | State of the address.  This field is required for KYC verification (US-based accounts only). | [optional] 
**Zip** | Pointer to **string** | United States ZIP code of the address. | [optional] 

## Methods

### NewAddressRequestModel

`func NewAddressRequestModel() *AddressRequestModel`

NewAddressRequestModel instantiates a new AddressRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressRequestModelWithDefaults

`func NewAddressRequestModelWithDefaults() *AddressRequestModel`

NewAddressRequestModelWithDefaults instantiates a new AddressRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *AddressRequestModel) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *AddressRequestModel) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *AddressRequestModel) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *AddressRequestModel) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *AddressRequestModel) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *AddressRequestModel) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *AddressRequestModel) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *AddressRequestModel) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *AddressRequestModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressRequestModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressRequestModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressRequestModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *AddressRequestModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressRequestModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressRequestModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AddressRequestModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressRequestModel) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressRequestModel) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressRequestModel) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressRequestModel) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *AddressRequestModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressRequestModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressRequestModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressRequestModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *AddressRequestModel) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddressRequestModel) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddressRequestModel) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AddressRequestModel) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


