# BusinessIncorporation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRegisteredUnder** | Pointer to [**AddressRequestModel**](AddressRequestModel.md) |  | [optional] 
**IncorporationType** | Pointer to **string** | Organizational structure of the business, such as corporation or sole proprietorship.  This field is required for KYC verification (US-based accounts only). | [optional] 
**IsPublic** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that the business is publicly held. | [optional] [default to false]
**NameRegisteredUnder** | Pointer to **string** | Name under which the business is registered. | [optional] 
**StateOfIncorporation** | Pointer to **string** | State where the business is incorporated.  This field is required for KYC verification (US-based accounts only). | [optional] 
**StockSymbol** | Pointer to **string** | Business stock symbol. | [optional] 

## Methods

### NewBusinessIncorporation

`func NewBusinessIncorporation() *BusinessIncorporation`

NewBusinessIncorporation instantiates a new BusinessIncorporation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessIncorporationWithDefaults

`func NewBusinessIncorporationWithDefaults() *BusinessIncorporation`

NewBusinessIncorporationWithDefaults instantiates a new BusinessIncorporation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRegisteredUnder

`func (o *BusinessIncorporation) GetAddressRegisteredUnder() AddressRequestModel`

GetAddressRegisteredUnder returns the AddressRegisteredUnder field if non-nil, zero value otherwise.

### GetAddressRegisteredUnderOk

`func (o *BusinessIncorporation) GetAddressRegisteredUnderOk() (*AddressRequestModel, bool)`

GetAddressRegisteredUnderOk returns a tuple with the AddressRegisteredUnder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRegisteredUnder

`func (o *BusinessIncorporation) SetAddressRegisteredUnder(v AddressRequestModel)`

SetAddressRegisteredUnder sets AddressRegisteredUnder field to given value.

### HasAddressRegisteredUnder

`func (o *BusinessIncorporation) HasAddressRegisteredUnder() bool`

HasAddressRegisteredUnder returns a boolean if a field has been set.

### GetIncorporationType

`func (o *BusinessIncorporation) GetIncorporationType() string`

GetIncorporationType returns the IncorporationType field if non-nil, zero value otherwise.

### GetIncorporationTypeOk

`func (o *BusinessIncorporation) GetIncorporationTypeOk() (*string, bool)`

GetIncorporationTypeOk returns a tuple with the IncorporationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationType

`func (o *BusinessIncorporation) SetIncorporationType(v string)`

SetIncorporationType sets IncorporationType field to given value.

### HasIncorporationType

`func (o *BusinessIncorporation) HasIncorporationType() bool`

HasIncorporationType returns a boolean if a field has been set.

### GetIsPublic

`func (o *BusinessIncorporation) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *BusinessIncorporation) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *BusinessIncorporation) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *BusinessIncorporation) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetNameRegisteredUnder

`func (o *BusinessIncorporation) GetNameRegisteredUnder() string`

GetNameRegisteredUnder returns the NameRegisteredUnder field if non-nil, zero value otherwise.

### GetNameRegisteredUnderOk

`func (o *BusinessIncorporation) GetNameRegisteredUnderOk() (*string, bool)`

GetNameRegisteredUnderOk returns a tuple with the NameRegisteredUnder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRegisteredUnder

`func (o *BusinessIncorporation) SetNameRegisteredUnder(v string)`

SetNameRegisteredUnder sets NameRegisteredUnder field to given value.

### HasNameRegisteredUnder

`func (o *BusinessIncorporation) HasNameRegisteredUnder() bool`

HasNameRegisteredUnder returns a boolean if a field has been set.

### GetStateOfIncorporation

`func (o *BusinessIncorporation) GetStateOfIncorporation() string`

GetStateOfIncorporation returns the StateOfIncorporation field if non-nil, zero value otherwise.

### GetStateOfIncorporationOk

`func (o *BusinessIncorporation) GetStateOfIncorporationOk() (*string, bool)`

GetStateOfIncorporationOk returns a tuple with the StateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfIncorporation

`func (o *BusinessIncorporation) SetStateOfIncorporation(v string)`

SetStateOfIncorporation sets StateOfIncorporation field to given value.

### HasStateOfIncorporation

`func (o *BusinessIncorporation) HasStateOfIncorporation() bool`

HasStateOfIncorporation returns a boolean if a field has been set.

### GetStockSymbol

`func (o *BusinessIncorporation) GetStockSymbol() string`

GetStockSymbol returns the StockSymbol field if non-nil, zero value otherwise.

### GetStockSymbolOk

`func (o *BusinessIncorporation) GetStockSymbolOk() (*string, bool)`

GetStockSymbolOk returns a tuple with the StockSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockSymbol

`func (o *BusinessIncorporation) SetStockSymbol(v string)`

SetStockSymbol sets StockSymbol field to given value.

### HasStockSymbol

`func (o *BusinessIncorporation) HasStockSymbol() bool`

HasStockSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


