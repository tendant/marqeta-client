# DigitalWalletSamsungPayProvisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**CreatedTime** | **time.Time** | Date and time when the digital wallet provisioning request was created, in UTC. | 
**LastModifiedTime** | **time.Time** | Date and time when the digital wallet token provisioning request was last updated, in UTC. | 
**PushTokenizeRequestData** | [**SamsungPushTokenizeRequestData**](SamsungPushTokenizeRequestData.md) |  | 

## Methods

### NewDigitalWalletSamsungPayProvisionResponse

`func NewDigitalWalletSamsungPayProvisionResponse(cardToken string, createdTime time.Time, lastModifiedTime time.Time, pushTokenizeRequestData SamsungPushTokenizeRequestData, ) *DigitalWalletSamsungPayProvisionResponse`

NewDigitalWalletSamsungPayProvisionResponse instantiates a new DigitalWalletSamsungPayProvisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletSamsungPayProvisionResponseWithDefaults

`func NewDigitalWalletSamsungPayProvisionResponseWithDefaults() *DigitalWalletSamsungPayProvisionResponse`

NewDigitalWalletSamsungPayProvisionResponseWithDefaults instantiates a new DigitalWalletSamsungPayProvisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *DigitalWalletSamsungPayProvisionResponse) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletSamsungPayProvisionResponse) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletSamsungPayProvisionResponse) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetCreatedTime

`func (o *DigitalWalletSamsungPayProvisionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DigitalWalletSamsungPayProvisionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DigitalWalletSamsungPayProvisionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastModifiedTime

`func (o *DigitalWalletSamsungPayProvisionResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DigitalWalletSamsungPayProvisionResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DigitalWalletSamsungPayProvisionResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetPushTokenizeRequestData

`func (o *DigitalWalletSamsungPayProvisionResponse) GetPushTokenizeRequestData() SamsungPushTokenizeRequestData`

GetPushTokenizeRequestData returns the PushTokenizeRequestData field if non-nil, zero value otherwise.

### GetPushTokenizeRequestDataOk

`func (o *DigitalWalletSamsungPayProvisionResponse) GetPushTokenizeRequestDataOk() (*SamsungPushTokenizeRequestData, bool)`

GetPushTokenizeRequestDataOk returns a tuple with the PushTokenizeRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTokenizeRequestData

`func (o *DigitalWalletSamsungPayProvisionResponse) SetPushTokenizeRequestData(v SamsungPushTokenizeRequestData)`

SetPushTokenizeRequestData sets PushTokenizeRequestData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


