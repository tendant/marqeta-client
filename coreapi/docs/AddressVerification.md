# AddressVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the cardholder. | [optional] 
**PostalCode** | Pointer to **string** | Postal code. | [optional] 
**StreetAddress** | Pointer to **string** | Street address provided by the cardholder. | [optional] 
**Zip** | Pointer to **string** | United States ZIP code. | [optional] 

## Methods

### NewAddressVerification

`func NewAddressVerification() *AddressVerification`

NewAddressVerification instantiates a new AddressVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressVerificationWithDefaults

`func NewAddressVerificationWithDefaults() *AddressVerification`

NewAddressVerificationWithDefaults instantiates a new AddressVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddressVerification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressVerification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressVerification) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressVerification) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressVerification) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressVerification) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressVerification) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressVerification) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *AddressVerification) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *AddressVerification) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *AddressVerification) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *AddressVerification) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetZip

`func (o *AddressVerification) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddressVerification) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddressVerification) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AddressVerification) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


