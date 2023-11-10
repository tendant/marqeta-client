# DigitalWalletSamsungPayProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**DeviceId** | **string** | User&#39;s Samsung device unique identifier, as provided by Samsung during the provisioning process. | 
**DeviceType** | **string** | Type of device into which the digital wallet token will be provisioned. | 
**ProvisioningAppVersion** | **string** | Version of the application making the provisioning request. Used for debugging and fraud prevention. | 
**WalletUserId** | **string** | User&#39;s Samsung Wallet account ID, as provided by Samsung during the provisioning process. | 

## Methods

### NewDigitalWalletSamsungPayProvisionRequest

`func NewDigitalWalletSamsungPayProvisionRequest(cardToken string, deviceId string, deviceType string, provisioningAppVersion string, walletUserId string, ) *DigitalWalletSamsungPayProvisionRequest`

NewDigitalWalletSamsungPayProvisionRequest instantiates a new DigitalWalletSamsungPayProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletSamsungPayProvisionRequestWithDefaults

`func NewDigitalWalletSamsungPayProvisionRequestWithDefaults() *DigitalWalletSamsungPayProvisionRequest`

NewDigitalWalletSamsungPayProvisionRequestWithDefaults instantiates a new DigitalWalletSamsungPayProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *DigitalWalletSamsungPayProvisionRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletSamsungPayProvisionRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletSamsungPayProvisionRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetDeviceId

`func (o *DigitalWalletSamsungPayProvisionRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DigitalWalletSamsungPayProvisionRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DigitalWalletSamsungPayProvisionRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceType

`func (o *DigitalWalletSamsungPayProvisionRequest) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DigitalWalletSamsungPayProvisionRequest) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DigitalWalletSamsungPayProvisionRequest) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetProvisioningAppVersion

`func (o *DigitalWalletSamsungPayProvisionRequest) GetProvisioningAppVersion() string`

GetProvisioningAppVersion returns the ProvisioningAppVersion field if non-nil, zero value otherwise.

### GetProvisioningAppVersionOk

`func (o *DigitalWalletSamsungPayProvisionRequest) GetProvisioningAppVersionOk() (*string, bool)`

GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAppVersion

`func (o *DigitalWalletSamsungPayProvisionRequest) SetProvisioningAppVersion(v string)`

SetProvisioningAppVersion sets ProvisioningAppVersion field to given value.


### GetWalletUserId

`func (o *DigitalWalletSamsungPayProvisionRequest) GetWalletUserId() string`

GetWalletUserId returns the WalletUserId field if non-nil, zero value otherwise.

### GetWalletUserIdOk

`func (o *DigitalWalletSamsungPayProvisionRequest) GetWalletUserIdOk() (*string, bool)`

GetWalletUserIdOk returns a tuple with the WalletUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletUserId

`func (o *DigitalWalletSamsungPayProvisionRequest) SetWalletUserId(v string)`

SetWalletUserId sets WalletUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


