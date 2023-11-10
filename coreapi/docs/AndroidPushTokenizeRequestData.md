# AndroidPushTokenizeRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name of the card as displayed in the digital wallet, typically showing the card brand and last four digits of the primary account number (PAN). &#x60;Visa 5678&#x60;, for example. | [optional] 
**LastDigits** | Pointer to **string** | Last four digits of the primary account number of the physical or virtual card. | [optional] 
**Network** | Pointer to **string** | Specifies the card network of the physical or virtual card. | [optional] 
**OpaquePaymentCard** | Pointer to **string** | Encrypted data field created by the issuer and passed to Google Wallet during the push provisioning process. | [optional] 
**TokenServiceProvider** | Pointer to **string** | Specifies the network that provides the digital wallet token service. | [optional] 
**UserAddress** | Pointer to [**AndroidPushTokenRequestAddress**](AndroidPushTokenRequestAddress.md) |  | [optional] 

## Methods

### NewAndroidPushTokenizeRequestData

`func NewAndroidPushTokenizeRequestData() *AndroidPushTokenizeRequestData`

NewAndroidPushTokenizeRequestData instantiates a new AndroidPushTokenizeRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidPushTokenizeRequestDataWithDefaults

`func NewAndroidPushTokenizeRequestDataWithDefaults() *AndroidPushTokenizeRequestData`

NewAndroidPushTokenizeRequestDataWithDefaults instantiates a new AndroidPushTokenizeRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AndroidPushTokenizeRequestData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AndroidPushTokenizeRequestData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AndroidPushTokenizeRequestData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AndroidPushTokenizeRequestData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastDigits

`func (o *AndroidPushTokenizeRequestData) GetLastDigits() string`

GetLastDigits returns the LastDigits field if non-nil, zero value otherwise.

### GetLastDigitsOk

`func (o *AndroidPushTokenizeRequestData) GetLastDigitsOk() (*string, bool)`

GetLastDigitsOk returns a tuple with the LastDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDigits

`func (o *AndroidPushTokenizeRequestData) SetLastDigits(v string)`

SetLastDigits sets LastDigits field to given value.

### HasLastDigits

`func (o *AndroidPushTokenizeRequestData) HasLastDigits() bool`

HasLastDigits returns a boolean if a field has been set.

### GetNetwork

`func (o *AndroidPushTokenizeRequestData) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AndroidPushTokenizeRequestData) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AndroidPushTokenizeRequestData) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AndroidPushTokenizeRequestData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOpaquePaymentCard

`func (o *AndroidPushTokenizeRequestData) GetOpaquePaymentCard() string`

GetOpaquePaymentCard returns the OpaquePaymentCard field if non-nil, zero value otherwise.

### GetOpaquePaymentCardOk

`func (o *AndroidPushTokenizeRequestData) GetOpaquePaymentCardOk() (*string, bool)`

GetOpaquePaymentCardOk returns a tuple with the OpaquePaymentCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaquePaymentCard

`func (o *AndroidPushTokenizeRequestData) SetOpaquePaymentCard(v string)`

SetOpaquePaymentCard sets OpaquePaymentCard field to given value.

### HasOpaquePaymentCard

`func (o *AndroidPushTokenizeRequestData) HasOpaquePaymentCard() bool`

HasOpaquePaymentCard returns a boolean if a field has been set.

### GetTokenServiceProvider

`func (o *AndroidPushTokenizeRequestData) GetTokenServiceProvider() string`

GetTokenServiceProvider returns the TokenServiceProvider field if non-nil, zero value otherwise.

### GetTokenServiceProviderOk

`func (o *AndroidPushTokenizeRequestData) GetTokenServiceProviderOk() (*string, bool)`

GetTokenServiceProviderOk returns a tuple with the TokenServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenServiceProvider

`func (o *AndroidPushTokenizeRequestData) SetTokenServiceProvider(v string)`

SetTokenServiceProvider sets TokenServiceProvider field to given value.

### HasTokenServiceProvider

`func (o *AndroidPushTokenizeRequestData) HasTokenServiceProvider() bool`

HasTokenServiceProvider returns a boolean if a field has been set.

### GetUserAddress

`func (o *AndroidPushTokenizeRequestData) GetUserAddress() AndroidPushTokenRequestAddress`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *AndroidPushTokenizeRequestData) GetUserAddressOk() (*AndroidPushTokenRequestAddress, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *AndroidPushTokenizeRequestData) SetUserAddress(v AndroidPushTokenRequestAddress)`

SetUserAddress sets UserAddress field to given value.

### HasUserAddress

`func (o *AndroidPushTokenizeRequestData) HasUserAddress() bool`

HasUserAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


