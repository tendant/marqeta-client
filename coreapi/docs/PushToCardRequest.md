# PushToCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | **string** |  | 
**Address2** | Pointer to **string** |  | [optional] 
**City** | **string** |  | 
**Country** | **string** |  | 
**Cvv** | **string** |  | 
**ExpDate** | **string** |  | 
**NameOnCard** | **string** |  | 
**Pan** | **string** |  | 
**PostalCode** | **string** |  | 
**State** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 
**UserToken** | **string** |  | 

## Methods

### NewPushToCardRequest

`func NewPushToCardRequest(address1 string, city string, country string, cvv string, expDate string, nameOnCard string, pan string, postalCode string, state string, userToken string, ) *PushToCardRequest`

NewPushToCardRequest instantiates a new PushToCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushToCardRequestWithDefaults

`func NewPushToCardRequestWithDefaults() *PushToCardRequest`

NewPushToCardRequestWithDefaults instantiates a new PushToCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *PushToCardRequest) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *PushToCardRequest) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *PushToCardRequest) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *PushToCardRequest) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *PushToCardRequest) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *PushToCardRequest) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *PushToCardRequest) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *PushToCardRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PushToCardRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PushToCardRequest) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *PushToCardRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PushToCardRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PushToCardRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCvv

`func (o *PushToCardRequest) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *PushToCardRequest) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *PushToCardRequest) SetCvv(v string)`

SetCvv sets Cvv field to given value.


### GetExpDate

`func (o *PushToCardRequest) GetExpDate() string`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *PushToCardRequest) GetExpDateOk() (*string, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *PushToCardRequest) SetExpDate(v string)`

SetExpDate sets ExpDate field to given value.


### GetNameOnCard

`func (o *PushToCardRequest) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *PushToCardRequest) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *PushToCardRequest) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.


### GetPan

`func (o *PushToCardRequest) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *PushToCardRequest) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *PushToCardRequest) SetPan(v string)`

SetPan sets Pan field to given value.


### GetPostalCode

`func (o *PushToCardRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PushToCardRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PushToCardRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetState

`func (o *PushToCardRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PushToCardRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PushToCardRequest) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *PushToCardRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushToCardRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushToCardRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PushToCardRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *PushToCardRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *PushToCardRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *PushToCardRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


