# AvsInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostalCode** | Pointer to **string** | Postal code of the address. | [optional] 
**StreetAddress** | Pointer to **string** | Street name and number of the address. | [optional] 
**Zip** | Pointer to **string** | United States ZIP code of the address. | [optional] 

## Methods

### NewAvsInformation

`func NewAvsInformation() *AvsInformation`

NewAvsInformation instantiates a new AvsInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvsInformationWithDefaults

`func NewAvsInformationWithDefaults() *AvsInformation`

NewAvsInformationWithDefaults instantiates a new AvsInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostalCode

`func (o *AvsInformation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AvsInformation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AvsInformation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AvsInformation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *AvsInformation) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *AvsInformation) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *AvsInformation) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *AvsInformation) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetZip

`func (o *AvsInformation) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AvsInformation) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AvsInformation) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AvsInformation) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


