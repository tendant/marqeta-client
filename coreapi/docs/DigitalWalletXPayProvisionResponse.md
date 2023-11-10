# DigitalWalletXPayProvisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**CreatedTime** | **time.Time** | Date and time when the digital wallet provisioning request was created, in UTC. | 
**LastModifiedTime** | **time.Time** | Date and time when the digital wallet token provisioning request was last updated, in UTC. | 
**PushTokenizeRequestData** | [**XpayPushTokenizeRequestData**](XpayPushTokenizeRequestData.md) |  | 

## Methods

### NewDigitalWalletXPayProvisionResponse

`func NewDigitalWalletXPayProvisionResponse(cardToken string, createdTime time.Time, lastModifiedTime time.Time, pushTokenizeRequestData XpayPushTokenizeRequestData, ) *DigitalWalletXPayProvisionResponse`

NewDigitalWalletXPayProvisionResponse instantiates a new DigitalWalletXPayProvisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletXPayProvisionResponseWithDefaults

`func NewDigitalWalletXPayProvisionResponseWithDefaults() *DigitalWalletXPayProvisionResponse`

NewDigitalWalletXPayProvisionResponseWithDefaults instantiates a new DigitalWalletXPayProvisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *DigitalWalletXPayProvisionResponse) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletXPayProvisionResponse) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletXPayProvisionResponse) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetCreatedTime

`func (o *DigitalWalletXPayProvisionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DigitalWalletXPayProvisionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DigitalWalletXPayProvisionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastModifiedTime

`func (o *DigitalWalletXPayProvisionResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DigitalWalletXPayProvisionResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DigitalWalletXPayProvisionResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetPushTokenizeRequestData

`func (o *DigitalWalletXPayProvisionResponse) GetPushTokenizeRequestData() XpayPushTokenizeRequestData`

GetPushTokenizeRequestData returns the PushTokenizeRequestData field if non-nil, zero value otherwise.

### GetPushTokenizeRequestDataOk

`func (o *DigitalWalletXPayProvisionResponse) GetPushTokenizeRequestDataOk() (*XpayPushTokenizeRequestData, bool)`

GetPushTokenizeRequestDataOk returns a tuple with the PushTokenizeRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTokenizeRequestData

`func (o *DigitalWalletXPayProvisionResponse) SetPushTokenizeRequestData(v XpayPushTokenizeRequestData)`

SetPushTokenizeRequestData sets PushTokenizeRequestData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


