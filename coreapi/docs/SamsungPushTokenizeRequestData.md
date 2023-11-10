# SamsungPushTokenizeRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardType** | Pointer to **string** | Specifies the type of card. | [optional] 
**DisplayName** | Pointer to **string** | Name of the card as displayed in the digital wallet, typically showing the card brand and last four digits of the primary account number (PAN). &#x60;Visa 5678&#x60;, for example. | [optional] 
**ExtraProvisionPayload** | Pointer to **string** | Encrypted card or cardholder details. | [optional] 
**LastDigits** | Pointer to **string** | Last four digits of the primary account number of the physical or virtual card. | [optional] 
**Network** | Pointer to **string** | Specifies the card network of the physical or virtual card. | [optional] 
**TokenServiceProvider** | Pointer to **string** | Specifies the network that provides the digital wallet token service, as determined by the Samsung Wallet library. | [optional] 

## Methods

### NewSamsungPushTokenizeRequestData

`func NewSamsungPushTokenizeRequestData() *SamsungPushTokenizeRequestData`

NewSamsungPushTokenizeRequestData instantiates a new SamsungPushTokenizeRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamsungPushTokenizeRequestDataWithDefaults

`func NewSamsungPushTokenizeRequestDataWithDefaults() *SamsungPushTokenizeRequestData`

NewSamsungPushTokenizeRequestDataWithDefaults instantiates a new SamsungPushTokenizeRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardType

`func (o *SamsungPushTokenizeRequestData) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *SamsungPushTokenizeRequestData) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *SamsungPushTokenizeRequestData) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *SamsungPushTokenizeRequestData) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetDisplayName

`func (o *SamsungPushTokenizeRequestData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SamsungPushTokenizeRequestData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SamsungPushTokenizeRequestData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SamsungPushTokenizeRequestData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExtraProvisionPayload

`func (o *SamsungPushTokenizeRequestData) GetExtraProvisionPayload() string`

GetExtraProvisionPayload returns the ExtraProvisionPayload field if non-nil, zero value otherwise.

### GetExtraProvisionPayloadOk

`func (o *SamsungPushTokenizeRequestData) GetExtraProvisionPayloadOk() (*string, bool)`

GetExtraProvisionPayloadOk returns a tuple with the ExtraProvisionPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProvisionPayload

`func (o *SamsungPushTokenizeRequestData) SetExtraProvisionPayload(v string)`

SetExtraProvisionPayload sets ExtraProvisionPayload field to given value.

### HasExtraProvisionPayload

`func (o *SamsungPushTokenizeRequestData) HasExtraProvisionPayload() bool`

HasExtraProvisionPayload returns a boolean if a field has been set.

### GetLastDigits

`func (o *SamsungPushTokenizeRequestData) GetLastDigits() string`

GetLastDigits returns the LastDigits field if non-nil, zero value otherwise.

### GetLastDigitsOk

`func (o *SamsungPushTokenizeRequestData) GetLastDigitsOk() (*string, bool)`

GetLastDigitsOk returns a tuple with the LastDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDigits

`func (o *SamsungPushTokenizeRequestData) SetLastDigits(v string)`

SetLastDigits sets LastDigits field to given value.

### HasLastDigits

`func (o *SamsungPushTokenizeRequestData) HasLastDigits() bool`

HasLastDigits returns a boolean if a field has been set.

### GetNetwork

`func (o *SamsungPushTokenizeRequestData) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SamsungPushTokenizeRequestData) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SamsungPushTokenizeRequestData) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SamsungPushTokenizeRequestData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTokenServiceProvider

`func (o *SamsungPushTokenizeRequestData) GetTokenServiceProvider() string`

GetTokenServiceProvider returns the TokenServiceProvider field if non-nil, zero value otherwise.

### GetTokenServiceProviderOk

`func (o *SamsungPushTokenizeRequestData) GetTokenServiceProviderOk() (*string, bool)`

GetTokenServiceProviderOk returns a tuple with the TokenServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenServiceProvider

`func (o *SamsungPushTokenizeRequestData) SetTokenServiceProvider(v string)`

SetTokenServiceProvider sets TokenServiceProvider field to given value.

### HasTokenServiceProvider

`func (o *SamsungPushTokenizeRequestData) HasTokenServiceProvider() bool`

HasTokenServiceProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


