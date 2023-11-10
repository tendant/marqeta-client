# FeeTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | **string** | Specifies the business account holder to which the fee applies. | 
**CreatedTime** | **time.Time** | Date and time when the &#x60;fee_charge&#x60; object was created, in UTC. | 
**Fees** | [**[]FeeDetail**](FeeDetail.md) | Contains attributes that define characteristics of one or more fees. | 
**Tags** | Pointer to **string** | Metadata about the fee charge.  This field is returned if it exists in the resource. | [optional] 
**Token** | **string** | Unique identifier of the fee transfer. | 
**UserToken** | **string** | Specifies the user account holder to which the fee applies. | 

## Methods

### NewFeeTransferResponse

`func NewFeeTransferResponse(businessToken string, createdTime time.Time, fees []FeeDetail, token string, userToken string, ) *FeeTransferResponse`

NewFeeTransferResponse instantiates a new FeeTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTransferResponseWithDefaults

`func NewFeeTransferResponseWithDefaults() *FeeTransferResponse`

NewFeeTransferResponseWithDefaults instantiates a new FeeTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *FeeTransferResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *FeeTransferResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *FeeTransferResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.


### GetCreatedTime

`func (o *FeeTransferResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FeeTransferResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FeeTransferResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetFees

`func (o *FeeTransferResponse) GetFees() []FeeDetail`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *FeeTransferResponse) GetFeesOk() (*[]FeeDetail, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *FeeTransferResponse) SetFees(v []FeeDetail)`

SetFees sets Fees field to given value.


### GetTags

`func (o *FeeTransferResponse) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeeTransferResponse) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeeTransferResponse) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeeTransferResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *FeeTransferResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FeeTransferResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FeeTransferResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *FeeTransferResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *FeeTransferResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *FeeTransferResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


