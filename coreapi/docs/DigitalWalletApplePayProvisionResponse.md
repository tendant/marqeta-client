# DigitalWalletApplePayProvisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationData** | **string** | Cryptographic one-time passcode conforming to the payment network operator or service provider specifications. | 
**CardToken** | **string** | Unique identifier of the card resource to use for the provisioning request. | 
**CreatedTime** | **time.Time** | Date and time when the digital wallet provisioning request was created, in UTC. | 
**EncryptedPassData** | **string** | Payload encrypted with a shared key derived from the Apple Public Certificates and the generated ephemeral private key. | 
**EphemeralPublicKey** | **string** | Ephemeral public key used for the provisioning attempt. | 
**LastModifiedTime** | **time.Time** | Date and time when the digital wallet token provisioning request was last updated, in UTC. | 

## Methods

### NewDigitalWalletApplePayProvisionResponse

`func NewDigitalWalletApplePayProvisionResponse(activationData string, cardToken string, createdTime time.Time, encryptedPassData string, ephemeralPublicKey string, lastModifiedTime time.Time, ) *DigitalWalletApplePayProvisionResponse`

NewDigitalWalletApplePayProvisionResponse instantiates a new DigitalWalletApplePayProvisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletApplePayProvisionResponseWithDefaults

`func NewDigitalWalletApplePayProvisionResponseWithDefaults() *DigitalWalletApplePayProvisionResponse`

NewDigitalWalletApplePayProvisionResponseWithDefaults instantiates a new DigitalWalletApplePayProvisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationData

`func (o *DigitalWalletApplePayProvisionResponse) GetActivationData() string`

GetActivationData returns the ActivationData field if non-nil, zero value otherwise.

### GetActivationDataOk

`func (o *DigitalWalletApplePayProvisionResponse) GetActivationDataOk() (*string, bool)`

GetActivationDataOk returns a tuple with the ActivationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationData

`func (o *DigitalWalletApplePayProvisionResponse) SetActivationData(v string)`

SetActivationData sets ActivationData field to given value.


### GetCardToken

`func (o *DigitalWalletApplePayProvisionResponse) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletApplePayProvisionResponse) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletApplePayProvisionResponse) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetCreatedTime

`func (o *DigitalWalletApplePayProvisionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DigitalWalletApplePayProvisionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DigitalWalletApplePayProvisionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEncryptedPassData

`func (o *DigitalWalletApplePayProvisionResponse) GetEncryptedPassData() string`

GetEncryptedPassData returns the EncryptedPassData field if non-nil, zero value otherwise.

### GetEncryptedPassDataOk

`func (o *DigitalWalletApplePayProvisionResponse) GetEncryptedPassDataOk() (*string, bool)`

GetEncryptedPassDataOk returns a tuple with the EncryptedPassData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassData

`func (o *DigitalWalletApplePayProvisionResponse) SetEncryptedPassData(v string)`

SetEncryptedPassData sets EncryptedPassData field to given value.


### GetEphemeralPublicKey

`func (o *DigitalWalletApplePayProvisionResponse) GetEphemeralPublicKey() string`

GetEphemeralPublicKey returns the EphemeralPublicKey field if non-nil, zero value otherwise.

### GetEphemeralPublicKeyOk

`func (o *DigitalWalletApplePayProvisionResponse) GetEphemeralPublicKeyOk() (*string, bool)`

GetEphemeralPublicKeyOk returns a tuple with the EphemeralPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralPublicKey

`func (o *DigitalWalletApplePayProvisionResponse) SetEphemeralPublicKey(v string)`

SetEphemeralPublicKey sets EphemeralPublicKey field to given value.


### GetLastModifiedTime

`func (o *DigitalWalletApplePayProvisionResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DigitalWalletApplePayProvisionResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DigitalWalletApplePayProvisionResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


