# SendingProvisioningDataToGooglePayBackendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardSetting** | **int32** | Indicates if the Funding Primary Account Number (FPAN) will be attempted.  * *1* - FPAN save will be attempted. * *0* - FPAN save will not be attempted. | 
**CardToken** | **string** | Unique identifier of the card resource. | 
**ClientSessionId** | **string** | String provided by Google Wallet that identifies the client session. | 
**IntegratorId** | **string** | Google-assigned string that uniquely identifies both the integrator that is initiating the session and the issuer of the card. | 
**PublicDeviceId** | **string** | String provided by Google Wallet that identifies the Android device that will receive the digital wallet token. | 
**PublicWalletId** | **string** | String provided by Google Wallet that identifies the device-scoped wallet that receives the digital wallet token. | 
**ServerSessionId** | **string** | String provided by Google Wallet that identifies the backend session. | 
**TokenSetting** | **int32** | Indicates if tokenization will be attempted.  * *1* - Tokenization will be attempted. * *0* - Tokenization will not be attempted. | 

## Methods

### NewSendingProvisioningDataToGooglePayBackendRequest

`func NewSendingProvisioningDataToGooglePayBackendRequest(cardSetting int32, cardToken string, clientSessionId string, integratorId string, publicDeviceId string, publicWalletId string, serverSessionId string, tokenSetting int32, ) *SendingProvisioningDataToGooglePayBackendRequest`

NewSendingProvisioningDataToGooglePayBackendRequest instantiates a new SendingProvisioningDataToGooglePayBackendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingProvisioningDataToGooglePayBackendRequestWithDefaults

`func NewSendingProvisioningDataToGooglePayBackendRequestWithDefaults() *SendingProvisioningDataToGooglePayBackendRequest`

NewSendingProvisioningDataToGooglePayBackendRequestWithDefaults instantiates a new SendingProvisioningDataToGooglePayBackendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardSetting

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardSetting() int32`

GetCardSetting returns the CardSetting field if non-nil, zero value otherwise.

### GetCardSettingOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardSettingOk() (*int32, bool)`

GetCardSettingOk returns a tuple with the CardSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSetting

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetCardSetting(v int32)`

SetCardSetting sets CardSetting field to given value.


### GetCardToken

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetClientSessionId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetClientSessionId() string`

GetClientSessionId returns the ClientSessionId field if non-nil, zero value otherwise.

### GetClientSessionIdOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetClientSessionIdOk() (*string, bool)`

GetClientSessionIdOk returns a tuple with the ClientSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSessionId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetClientSessionId(v string)`

SetClientSessionId sets ClientSessionId field to given value.


### GetIntegratorId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetIntegratorId() string`

GetIntegratorId returns the IntegratorId field if non-nil, zero value otherwise.

### GetIntegratorIdOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetIntegratorIdOk() (*string, bool)`

GetIntegratorIdOk returns a tuple with the IntegratorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetIntegratorId(v string)`

SetIntegratorId sets IntegratorId field to given value.


### GetPublicDeviceId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicDeviceId() string`

GetPublicDeviceId returns the PublicDeviceId field if non-nil, zero value otherwise.

### GetPublicDeviceIdOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicDeviceIdOk() (*string, bool)`

GetPublicDeviceIdOk returns a tuple with the PublicDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDeviceId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetPublicDeviceId(v string)`

SetPublicDeviceId sets PublicDeviceId field to given value.


### GetPublicWalletId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicWalletId() string`

GetPublicWalletId returns the PublicWalletId field if non-nil, zero value otherwise.

### GetPublicWalletIdOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetPublicWalletIdOk() (*string, bool)`

GetPublicWalletIdOk returns a tuple with the PublicWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicWalletId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetPublicWalletId(v string)`

SetPublicWalletId sets PublicWalletId field to given value.


### GetServerSessionId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetServerSessionId() string`

GetServerSessionId returns the ServerSessionId field if non-nil, zero value otherwise.

### GetServerSessionIdOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetServerSessionIdOk() (*string, bool)`

GetServerSessionIdOk returns a tuple with the ServerSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSessionId

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetServerSessionId(v string)`

SetServerSessionId sets ServerSessionId field to given value.


### GetTokenSetting

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetTokenSetting() int32`

GetTokenSetting returns the TokenSetting field if non-nil, zero value otherwise.

### GetTokenSettingOk

`func (o *SendingProvisioningDataToGooglePayBackendRequest) GetTokenSettingOk() (*int32, bool)`

GetTokenSettingOk returns a tuple with the TokenSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSetting

`func (o *SendingProvisioningDataToGooglePayBackendRequest) SetTokenSetting(v int32)`

SetTokenSetting sets TokenSetting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


