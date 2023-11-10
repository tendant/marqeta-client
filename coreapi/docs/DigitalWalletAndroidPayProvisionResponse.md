# DigitalWalletAndroidPayProvisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**CreatedTime** | **time.Time** | Date and time when the digital wallet provisioning request was created, in UTC. | 
**LastModifiedTime** | **time.Time** | Date and time when the digital wallet token provisioning request was last updated, in UTC. | 
**PushTokenizeRequestData** | [**AndroidPushTokenizeRequestData**](AndroidPushTokenizeRequestData.md) |  | 

## Methods

### NewDigitalWalletAndroidPayProvisionResponse

`func NewDigitalWalletAndroidPayProvisionResponse(cardToken string, createdTime time.Time, lastModifiedTime time.Time, pushTokenizeRequestData AndroidPushTokenizeRequestData, ) *DigitalWalletAndroidPayProvisionResponse`

NewDigitalWalletAndroidPayProvisionResponse instantiates a new DigitalWalletAndroidPayProvisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletAndroidPayProvisionResponseWithDefaults

`func NewDigitalWalletAndroidPayProvisionResponseWithDefaults() *DigitalWalletAndroidPayProvisionResponse`

NewDigitalWalletAndroidPayProvisionResponseWithDefaults instantiates a new DigitalWalletAndroidPayProvisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *DigitalWalletAndroidPayProvisionResponse) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletAndroidPayProvisionResponse) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletAndroidPayProvisionResponse) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetCreatedTime

`func (o *DigitalWalletAndroidPayProvisionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DigitalWalletAndroidPayProvisionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DigitalWalletAndroidPayProvisionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastModifiedTime

`func (o *DigitalWalletAndroidPayProvisionResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DigitalWalletAndroidPayProvisionResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DigitalWalletAndroidPayProvisionResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetPushTokenizeRequestData

`func (o *DigitalWalletAndroidPayProvisionResponse) GetPushTokenizeRequestData() AndroidPushTokenizeRequestData`

GetPushTokenizeRequestData returns the PushTokenizeRequestData field if non-nil, zero value otherwise.

### GetPushTokenizeRequestDataOk

`func (o *DigitalWalletAndroidPayProvisionResponse) GetPushTokenizeRequestDataOk() (*AndroidPushTokenizeRequestData, bool)`

GetPushTokenizeRequestDataOk returns a tuple with the PushTokenizeRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTokenizeRequestData

`func (o *DigitalWalletAndroidPayProvisionResponse) SetPushTokenizeRequestData(v AndroidPushTokenizeRequestData)`

SetPushTokenizeRequestData sets PushTokenizeRequestData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


