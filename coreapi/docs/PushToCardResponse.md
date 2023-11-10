# PushToCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**ExpDate** | Pointer to **string** |  | [optional] 
**FastFundTransferEligible** | Pointer to **bool** |  | [optional] [default to false]
**GamblingFundTransferEligible** | Pointer to **bool** |  | [optional] [default to false]
**LastFour** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**LastName** | Pointer to **string** |  | [optional] [readonly] 
**NameOnCard** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewPushToCardResponse

`func NewPushToCardResponse(createdTime time.Time, lastModifiedTime time.Time, ) *PushToCardResponse`

NewPushToCardResponse instantiates a new PushToCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushToCardResponseWithDefaults

`func NewPushToCardResponseWithDefaults() *PushToCardResponse`

NewPushToCardResponseWithDefaults instantiates a new PushToCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *PushToCardResponse) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *PushToCardResponse) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *PushToCardResponse) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *PushToCardResponse) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *PushToCardResponse) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *PushToCardResponse) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *PushToCardResponse) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *PushToCardResponse) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *PushToCardResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PushToCardResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PushToCardResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PushToCardResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *PushToCardResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PushToCardResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PushToCardResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PushToCardResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PushToCardResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PushToCardResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PushToCardResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetExpDate

`func (o *PushToCardResponse) GetExpDate() string`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *PushToCardResponse) GetExpDateOk() (*string, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *PushToCardResponse) SetExpDate(v string)`

SetExpDate sets ExpDate field to given value.

### HasExpDate

`func (o *PushToCardResponse) HasExpDate() bool`

HasExpDate returns a boolean if a field has been set.

### GetFastFundTransferEligible

`func (o *PushToCardResponse) GetFastFundTransferEligible() bool`

GetFastFundTransferEligible returns the FastFundTransferEligible field if non-nil, zero value otherwise.

### GetFastFundTransferEligibleOk

`func (o *PushToCardResponse) GetFastFundTransferEligibleOk() (*bool, bool)`

GetFastFundTransferEligibleOk returns a tuple with the FastFundTransferEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastFundTransferEligible

`func (o *PushToCardResponse) SetFastFundTransferEligible(v bool)`

SetFastFundTransferEligible sets FastFundTransferEligible field to given value.

### HasFastFundTransferEligible

`func (o *PushToCardResponse) HasFastFundTransferEligible() bool`

HasFastFundTransferEligible returns a boolean if a field has been set.

### GetGamblingFundTransferEligible

`func (o *PushToCardResponse) GetGamblingFundTransferEligible() bool`

GetGamblingFundTransferEligible returns the GamblingFundTransferEligible field if non-nil, zero value otherwise.

### GetGamblingFundTransferEligibleOk

`func (o *PushToCardResponse) GetGamblingFundTransferEligibleOk() (*bool, bool)`

GetGamblingFundTransferEligibleOk returns a tuple with the GamblingFundTransferEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamblingFundTransferEligible

`func (o *PushToCardResponse) SetGamblingFundTransferEligible(v bool)`

SetGamblingFundTransferEligible sets GamblingFundTransferEligible field to given value.

### HasGamblingFundTransferEligible

`func (o *PushToCardResponse) HasGamblingFundTransferEligible() bool`

HasGamblingFundTransferEligible returns a boolean if a field has been set.

### GetLastFour

`func (o *PushToCardResponse) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *PushToCardResponse) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *PushToCardResponse) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *PushToCardResponse) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *PushToCardResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *PushToCardResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *PushToCardResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetLastName

`func (o *PushToCardResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PushToCardResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PushToCardResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PushToCardResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetNameOnCard

`func (o *PushToCardResponse) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *PushToCardResponse) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *PushToCardResponse) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *PushToCardResponse) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetPostalCode

`func (o *PushToCardResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PushToCardResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PushToCardResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PushToCardResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *PushToCardResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PushToCardResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PushToCardResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PushToCardResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *PushToCardResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushToCardResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushToCardResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PushToCardResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


