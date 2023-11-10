# PTCAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** |  | 
**Country** | Pointer to **string** |  | [optional] 
**County** | **string** |  | 
**Line1** | **string** |  | 
**Line2** | Pointer to **string** |  | [optional] 
**PostalCode** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewPTCAddress

`func NewPTCAddress(city string, county string, line1 string, postalCode string, state string, ) *PTCAddress`

NewPTCAddress instantiates a new PTCAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPTCAddressWithDefaults

`func NewPTCAddressWithDefaults() *PTCAddress`

NewPTCAddressWithDefaults instantiates a new PTCAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *PTCAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PTCAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PTCAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *PTCAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PTCAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PTCAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PTCAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *PTCAddress) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *PTCAddress) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *PTCAddress) SetCounty(v string)`

SetCounty sets County field to given value.


### GetLine1

`func (o *PTCAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *PTCAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *PTCAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *PTCAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *PTCAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *PTCAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *PTCAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetPostalCode

`func (o *PTCAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PTCAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PTCAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetState

`func (o *PTCAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PTCAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PTCAddress) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


