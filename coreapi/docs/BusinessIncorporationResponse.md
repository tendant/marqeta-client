# BusinessIncorporationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressRegisteredUnder** | Pointer to [**AddressResponseModel**](AddressResponseModel.md) |  | [optional] 
**IncorporationType** | Pointer to **string** | Organizational structure of the business (corporation or sole proprietorship, for example).  This field is returned if it exists in the resource. | [optional] 
**IsPublic** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that the business is publicly held.  This field is returned if it exists in the resource. | [optional] [default to false]
**NameRegisteredUnder** | Pointer to **string** | Name under which the business is registered.  This field is returned if it exists in the resource. | [optional] 
**StateOfIncorporation** | Pointer to **string** | State where the business is incorporated.  This field is returned if it exists in the resource. | [optional] 
**StockSymbol** | Pointer to **string** | Stock symbol associated with the business.  This field is returned if it exists in the resource. | [optional] 

## Methods

### NewBusinessIncorporationResponse

`func NewBusinessIncorporationResponse() *BusinessIncorporationResponse`

NewBusinessIncorporationResponse instantiates a new BusinessIncorporationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessIncorporationResponseWithDefaults

`func NewBusinessIncorporationResponseWithDefaults() *BusinessIncorporationResponse`

NewBusinessIncorporationResponseWithDefaults instantiates a new BusinessIncorporationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressRegisteredUnder

`func (o *BusinessIncorporationResponse) GetAddressRegisteredUnder() AddressResponseModel`

GetAddressRegisteredUnder returns the AddressRegisteredUnder field if non-nil, zero value otherwise.

### GetAddressRegisteredUnderOk

`func (o *BusinessIncorporationResponse) GetAddressRegisteredUnderOk() (*AddressResponseModel, bool)`

GetAddressRegisteredUnderOk returns a tuple with the AddressRegisteredUnder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRegisteredUnder

`func (o *BusinessIncorporationResponse) SetAddressRegisteredUnder(v AddressResponseModel)`

SetAddressRegisteredUnder sets AddressRegisteredUnder field to given value.

### HasAddressRegisteredUnder

`func (o *BusinessIncorporationResponse) HasAddressRegisteredUnder() bool`

HasAddressRegisteredUnder returns a boolean if a field has been set.

### GetIncorporationType

`func (o *BusinessIncorporationResponse) GetIncorporationType() string`

GetIncorporationType returns the IncorporationType field if non-nil, zero value otherwise.

### GetIncorporationTypeOk

`func (o *BusinessIncorporationResponse) GetIncorporationTypeOk() (*string, bool)`

GetIncorporationTypeOk returns a tuple with the IncorporationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationType

`func (o *BusinessIncorporationResponse) SetIncorporationType(v string)`

SetIncorporationType sets IncorporationType field to given value.

### HasIncorporationType

`func (o *BusinessIncorporationResponse) HasIncorporationType() bool`

HasIncorporationType returns a boolean if a field has been set.

### GetIsPublic

`func (o *BusinessIncorporationResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *BusinessIncorporationResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *BusinessIncorporationResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *BusinessIncorporationResponse) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetNameRegisteredUnder

`func (o *BusinessIncorporationResponse) GetNameRegisteredUnder() string`

GetNameRegisteredUnder returns the NameRegisteredUnder field if non-nil, zero value otherwise.

### GetNameRegisteredUnderOk

`func (o *BusinessIncorporationResponse) GetNameRegisteredUnderOk() (*string, bool)`

GetNameRegisteredUnderOk returns a tuple with the NameRegisteredUnder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRegisteredUnder

`func (o *BusinessIncorporationResponse) SetNameRegisteredUnder(v string)`

SetNameRegisteredUnder sets NameRegisteredUnder field to given value.

### HasNameRegisteredUnder

`func (o *BusinessIncorporationResponse) HasNameRegisteredUnder() bool`

HasNameRegisteredUnder returns a boolean if a field has been set.

### GetStateOfIncorporation

`func (o *BusinessIncorporationResponse) GetStateOfIncorporation() string`

GetStateOfIncorporation returns the StateOfIncorporation field if non-nil, zero value otherwise.

### GetStateOfIncorporationOk

`func (o *BusinessIncorporationResponse) GetStateOfIncorporationOk() (*string, bool)`

GetStateOfIncorporationOk returns a tuple with the StateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfIncorporation

`func (o *BusinessIncorporationResponse) SetStateOfIncorporation(v string)`

SetStateOfIncorporation sets StateOfIncorporation field to given value.

### HasStateOfIncorporation

`func (o *BusinessIncorporationResponse) HasStateOfIncorporation() bool`

HasStateOfIncorporation returns a boolean if a field has been set.

### GetStockSymbol

`func (o *BusinessIncorporationResponse) GetStockSymbol() string`

GetStockSymbol returns the StockSymbol field if non-nil, zero value otherwise.

### GetStockSymbolOk

`func (o *BusinessIncorporationResponse) GetStockSymbolOk() (*string, bool)`

GetStockSymbolOk returns a tuple with the StockSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockSymbol

`func (o *BusinessIncorporationResponse) SetStockSymbol(v string)`

SetStockSymbol sets StockSymbol field to given value.

### HasStockSymbol

`func (o *BusinessIncorporationResponse) HasStockSymbol() bool`

HasStockSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


