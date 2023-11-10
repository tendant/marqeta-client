# DigitalWalletXPayProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**DeviceId** | **string** | Unique identifier of the user&#39;s XPay device, as provided by XPay during the provisioning process. | 
**DeviceType** | **string** | Type of device into which the digital wallet token will be provisioned. | 
**ProvisioningAppVersion** | **string** | Version of the application making the provisioning request. Used for debugging and fraud prevention. | 
**TokenRequestorId** | **string** | Unique numerical identifier of the digital wallet token requestor within the card network. These ID numbers map to &#x60;token_requestor_name&#x60; field values as follows:  *Mastercard*  * 50110030273 – &#x60;APPLE_PAY&#x60; * 50120834693 – &#x60;ANDROID_PAY&#x60; * 50139059239 – &#x60;SAMSUNG_PAY&#x60;  *Visa*  * 40010030273 – &#x60;APPLE_PAY&#x60; * 40010075001 – &#x60;ANDROID_PAY&#x60; * 40010043095 – &#x60;SAMSUNG_PAY&#x60; * 40010075196 – &#x60;MICROSOFT_PAY&#x60; * 40010075338 – &#x60;VISA_CHECKOUT&#x60; * 40010075449 – &#x60;FACEBOOK&#x60; * 40010075839 – &#x60;NETFLIX&#x60; * 40010077056 – &#x60;FITBIT_PAY&#x60; * 40010069887 – &#x60;GARMIN_PAY&#x60; | 
**WalletAccountId** | **string** | User&#39;s XPay account identifier, as provided by XPay during the provisioning process. | 

## Methods

### NewDigitalWalletXPayProvisionRequest

`func NewDigitalWalletXPayProvisionRequest(cardToken string, deviceId string, deviceType string, provisioningAppVersion string, tokenRequestorId string, walletAccountId string, ) *DigitalWalletXPayProvisionRequest`

NewDigitalWalletXPayProvisionRequest instantiates a new DigitalWalletXPayProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletXPayProvisionRequestWithDefaults

`func NewDigitalWalletXPayProvisionRequestWithDefaults() *DigitalWalletXPayProvisionRequest`

NewDigitalWalletXPayProvisionRequestWithDefaults instantiates a new DigitalWalletXPayProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *DigitalWalletXPayProvisionRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletXPayProvisionRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletXPayProvisionRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetDeviceId

`func (o *DigitalWalletXPayProvisionRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DigitalWalletXPayProvisionRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DigitalWalletXPayProvisionRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceType

`func (o *DigitalWalletXPayProvisionRequest) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DigitalWalletXPayProvisionRequest) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DigitalWalletXPayProvisionRequest) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetProvisioningAppVersion

`func (o *DigitalWalletXPayProvisionRequest) GetProvisioningAppVersion() string`

GetProvisioningAppVersion returns the ProvisioningAppVersion field if non-nil, zero value otherwise.

### GetProvisioningAppVersionOk

`func (o *DigitalWalletXPayProvisionRequest) GetProvisioningAppVersionOk() (*string, bool)`

GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAppVersion

`func (o *DigitalWalletXPayProvisionRequest) SetProvisioningAppVersion(v string)`

SetProvisioningAppVersion sets ProvisioningAppVersion field to given value.


### GetTokenRequestorId

`func (o *DigitalWalletXPayProvisionRequest) GetTokenRequestorId() string`

GetTokenRequestorId returns the TokenRequestorId field if non-nil, zero value otherwise.

### GetTokenRequestorIdOk

`func (o *DigitalWalletXPayProvisionRequest) GetTokenRequestorIdOk() (*string, bool)`

GetTokenRequestorIdOk returns a tuple with the TokenRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestorId

`func (o *DigitalWalletXPayProvisionRequest) SetTokenRequestorId(v string)`

SetTokenRequestorId sets TokenRequestorId field to given value.


### GetWalletAccountId

`func (o *DigitalWalletXPayProvisionRequest) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *DigitalWalletXPayProvisionRequest) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *DigitalWalletXPayProvisionRequest) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


