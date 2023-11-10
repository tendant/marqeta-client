# DigitalWalletApplePayProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**Certificates** | **[]string** | Base64-encoded leaf and sub-CA certificates provided by Apple.  The first element of the array should be the leaf certificate, followed by the sub-CA. | 
**DeviceType** | **string** | Type of device into which the digital wallet token will be provisioned. | 
**Nonce** | **string** | One-time-use nonce provided by Apple for security purposes. | 
**NonceSignature** | **string** | Apple-provided signature to the nonce. | 
**ProvisioningAppVersion** | **string** | Version of the application making the provisioning request. Used for debugging and fraud prevention. | 

## Methods

### NewDigitalWalletApplePayProvisionRequest

`func NewDigitalWalletApplePayProvisionRequest(cardToken string, certificates []string, deviceType string, nonce string, nonceSignature string, provisioningAppVersion string, ) *DigitalWalletApplePayProvisionRequest`

NewDigitalWalletApplePayProvisionRequest instantiates a new DigitalWalletApplePayProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletApplePayProvisionRequestWithDefaults

`func NewDigitalWalletApplePayProvisionRequestWithDefaults() *DigitalWalletApplePayProvisionRequest`

NewDigitalWalletApplePayProvisionRequestWithDefaults instantiates a new DigitalWalletApplePayProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *DigitalWalletApplePayProvisionRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletApplePayProvisionRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletApplePayProvisionRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetCertificates

`func (o *DigitalWalletApplePayProvisionRequest) GetCertificates() []string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *DigitalWalletApplePayProvisionRequest) GetCertificatesOk() (*[]string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *DigitalWalletApplePayProvisionRequest) SetCertificates(v []string)`

SetCertificates sets Certificates field to given value.


### GetDeviceType

`func (o *DigitalWalletApplePayProvisionRequest) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DigitalWalletApplePayProvisionRequest) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DigitalWalletApplePayProvisionRequest) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetNonce

`func (o *DigitalWalletApplePayProvisionRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DigitalWalletApplePayProvisionRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DigitalWalletApplePayProvisionRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetNonceSignature

`func (o *DigitalWalletApplePayProvisionRequest) GetNonceSignature() string`

GetNonceSignature returns the NonceSignature field if non-nil, zero value otherwise.

### GetNonceSignatureOk

`func (o *DigitalWalletApplePayProvisionRequest) GetNonceSignatureOk() (*string, bool)`

GetNonceSignatureOk returns a tuple with the NonceSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonceSignature

`func (o *DigitalWalletApplePayProvisionRequest) SetNonceSignature(v string)`

SetNonceSignature sets NonceSignature field to given value.


### GetProvisioningAppVersion

`func (o *DigitalWalletApplePayProvisionRequest) GetProvisioningAppVersion() string`

GetProvisioningAppVersion returns the ProvisioningAppVersion field if non-nil, zero value otherwise.

### GetProvisioningAppVersionOk

`func (o *DigitalWalletApplePayProvisionRequest) GetProvisioningAppVersionOk() (*string, bool)`

GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAppVersion

`func (o *DigitalWalletApplePayProvisionRequest) SetProvisioningAppVersion(v string)`

SetProvisioningAppVersion sets ProvisioningAppVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


