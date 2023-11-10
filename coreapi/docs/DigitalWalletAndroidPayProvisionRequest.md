# DigitalWalletAndroidPayProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**DeviceId** | **string** | Unique identifier of the user&#39;s Google device, as provided by Google during the provisioning process. | 
**DeviceType** | **string** | Type of device into which the digital wallet token will be provisioned. | 
**ProvisioningAppVersion** | **string** | Version of the application making the provisioning request. Used for debugging and fraud prevention. | 
**WalletAccountId** | **string** | User&#39;s Google Wallet account ID, as provided by Google during the provisioning process. | 

## Methods

### NewDigitalWalletAndroidPayProvisionRequest

`func NewDigitalWalletAndroidPayProvisionRequest(cardToken string, deviceId string, deviceType string, provisioningAppVersion string, walletAccountId string, ) *DigitalWalletAndroidPayProvisionRequest`

NewDigitalWalletAndroidPayProvisionRequest instantiates a new DigitalWalletAndroidPayProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletAndroidPayProvisionRequestWithDefaults

`func NewDigitalWalletAndroidPayProvisionRequestWithDefaults() *DigitalWalletAndroidPayProvisionRequest`

NewDigitalWalletAndroidPayProvisionRequestWithDefaults instantiates a new DigitalWalletAndroidPayProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *DigitalWalletAndroidPayProvisionRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletAndroidPayProvisionRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletAndroidPayProvisionRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetDeviceId

`func (o *DigitalWalletAndroidPayProvisionRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DigitalWalletAndroidPayProvisionRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DigitalWalletAndroidPayProvisionRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceType

`func (o *DigitalWalletAndroidPayProvisionRequest) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DigitalWalletAndroidPayProvisionRequest) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DigitalWalletAndroidPayProvisionRequest) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetProvisioningAppVersion

`func (o *DigitalWalletAndroidPayProvisionRequest) GetProvisioningAppVersion() string`

GetProvisioningAppVersion returns the ProvisioningAppVersion field if non-nil, zero value otherwise.

### GetProvisioningAppVersionOk

`func (o *DigitalWalletAndroidPayProvisionRequest) GetProvisioningAppVersionOk() (*string, bool)`

GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAppVersion

`func (o *DigitalWalletAndroidPayProvisionRequest) SetProvisioningAppVersion(v string)`

SetProvisioningAppVersion sets ProvisioningAppVersion field to given value.


### GetWalletAccountId

`func (o *DigitalWalletAndroidPayProvisionRequest) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *DigitalWalletAndroidPayProvisionRequest) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *DigitalWalletAndroidPayProvisionRequest) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


