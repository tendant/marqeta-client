# CreditBureauAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | **string** | Street address. | 
**Address2** | Pointer to **string** | Additional address information. | [optional] 
**City** | **string** | Address city. | 
**State** | **string** | Address state. | 
**Zip** | **string** | Address ZIP code. | 

## Methods

### NewCreditBureauAddress

`func NewCreditBureauAddress(address1 string, city string, state string, zip string, ) *CreditBureauAddress`

NewCreditBureauAddress instantiates a new CreditBureauAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBureauAddressWithDefaults

`func NewCreditBureauAddressWithDefaults() *CreditBureauAddress`

NewCreditBureauAddressWithDefaults instantiates a new CreditBureauAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *CreditBureauAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CreditBureauAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CreditBureauAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *CreditBureauAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CreditBureauAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CreditBureauAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CreditBureauAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *CreditBureauAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreditBureauAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreditBureauAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *CreditBureauAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreditBureauAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreditBureauAddress) SetState(v string)`

SetState sets State field to given value.


### GetZip

`func (o *CreditBureauAddress) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CreditBureauAddress) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CreditBureauAddress) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


